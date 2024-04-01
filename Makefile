#
# Generic makefile for golang + gogi projects
#
# Usage:
#
#   Tasks for CI:
#
#	@build  - build production image
#	@push   - push production image
#	@stub   - create .artifact file containing current git version
#
#   Tasks for local development
#
#	setup       - install dependencies and ci tools
#   build       - create binary executable
#   run-server  - run server with config
#	lint        - run golangci linter inside ci image
#	test        - run go test inside ci image
#
#   Extra params for @test and @lint
#   - CI_FORCE_REBUILD=true   - Disables checking for existing ci image, which will force rebuilding ci image
#   - CI_MOUNT_LOCAL_FS=false - Disables mounting local fs for @test and @lint targets.
#                               Used for running targets over code contained inside image
#
# Extending and configuration
#
#   Provide correct values for variables:
#   - GOGI_VERSION - gogi image version
#   - GOCI_VERSION - goci toolkit version
#   - BASE_VERSION - basemod version
#
#   Add supplementary targets inside extra.mk, you can overwrite existing targets to provide extra steps
#

BASE_VERSION=1.21
GOCI_VERSION=v1.54.2

DOCKER_SOURCE_DIR=/go/src/stash.lamoda.ru/neurocapsule/neurocapsule
DOCKER_REGISTRY?=neurocapsule.docker.lamoda.ru

VERSION ?= $(shell git describe --tags || git rev-parse --short HEAD)
NAMESPACE = neurocapsule

CI_PUSH_LATEST?=false
CI_MOUNT_LOCAL_FS?=true
CI_IMAGE_EXISTS:=$(shell docker image inspect ${DOCKER_REGISTRY}/${NAMESPACE}/ci:${VERSION} 2>/dev/null 1>/dev/null && echo "false" || echo "true")
CI_FORCE_REBUILD?=${CI_IMAGE_EXISTS}

#
# Extra targets for repository
#
-include deployments/extra.mk

#
# Local targets
#
setup: GOLANGCI_INSTALLER=https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh
setup:
	curl -sSfL ${GOLANGCI_INSTALLER} | sh -s -- -b ${GOPATH}/bin ${GOCI_VERSION}

test:
	CGO_ENABLED=1 go test -cover -race -coverprofile=coverage.out -v ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run

build: LDFLAGS="-s -X stash.lamoda.ru/neurocapsule/neurocapsule/cmd.Version=$(VERSION)"
build:
	GOPRIVATE=stash.lamoda.ru CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags ${LDFLAGS} -v -a -o output main.go

run-server:
	NEUROCAPSULE_CONF=deployments/local.dev.env go run main.go application

#
# Docker targets
#
docker-build-ci-image:
ifeq (true,${CI_FORCE_REBUILD})
	@docker build \
		--build-arg VERSION=${VERSION} \
		--build-arg BASE_VERSION=${BASE_VERSION} \
		--tag ${DOCKER_REGISTRY}/${NAMESPACE}/ci:latest \
		--tag ${DOCKER_REGISTRY}/${NAMESPACE}/ci:${VERSION} \
		--file deployments/Dockerfile \
		--target ci \
		.
else
	$(info [MAKE] Skipping CI image build. Pass CI_FORCE_REBUILD=true if rebuild required)
endif

docker-prepare-test-env:
	$(info Nothing to prepare. Skipping target)

docker-remove-test-env:
	$(info Nothing to remove. Skipping target)

#
# Mandatory targets for CI/CD
#
@build:
	$(info [MAKE] Bulding production image)
	@docker build \
		--build-arg VERSION=${VERSION} \
		--build-arg BASE_VERSION=${BASE_VERSION} \
		--tag ${DOCKER_REGISTRY}/${NAMESPACE}/app:latest \
		--tag ${DOCKER_REGISTRY}/${NAMESPACE}/app:${VERSION} \
		--file deployments/Dockerfile \
		.

@push:
	$(info [MAKE] Pushing production image)
ifeq (true,${CI_PUSH_LATEST})
	@docker push ${DOCKER_REGISTRY}/${NAMESPACE}/app:${VERSION} \
				 ${DOCKER_REGISTRY}/${NAMESPACE}/app:latest
else
	@docker push ${DOCKER_REGISTRY}/${NAMESPACE}/app:${VERSION}
endif

@lint: docker-build-ci-image
	$(info [MAKE] Running linters)
ifeq (true,${CI_MOUNT_LOCAL_FS})
	@docker run --rm \
		--volume $(PWD):$(DOCKER_SOURCE_DIR) \
		--workdir "$(DOCKER_SOURCE_DIR)" \
		${DOCKER_REGISTRY}/${NAMESPACE}/ci:${VERSION} \
		make lint
else
	@docker run --rm \
		${DOCKER_REGISTRY}/${NAMESPACE}/ci:${VERSION} \
		make lint
endif

@test: docker-build-ci-image docker-prepare-test-env
	$(info [MAKE] Running tests)
ifeq (true,${CI_MOUNT_LOCAL_FS})
	@docker run --rm \
		--volume $(PWD):$(DOCKER_SOURCE_DIR) \
		--workdir "$(DOCKER_SOURCE_DIR)" \
		${DOCKER_REGISTRY}/${NAMESPACE}/ci:${VERSION} \
		make test
else
	@docker run --rm \
		${DOCKER_REGISTRY}/${NAMESPACE}/ci:${VERSION} \
		make test
endif

@clean: docker-remove-test-env
	$(info [MAKE] Cleaning docker images)
	@docker image rm --force \
		${DOCKER_REGISTRY}/${NAMESPACE}/ci:${VERSION} \
		${DOCKER_REGISTRY}/${NAMESPACE}/ci:latest \
		${DOCKER_REGISTRY}/${NAMESPACE}/app:${VERSION} \
		${DOCKER_REGISTRY}/${NAMESPACE}/app:latest \
	|| true

@stub:
	@echo "VERSION=${VERSION}" > .artifact
	@echo "LAST_BUILD_VERSION=$(VERSION)" > .release
