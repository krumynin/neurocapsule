_COMPOSE=docker-compose -f deployments/docker-compose.yaml --project-name neurocapsule
_COMPOSE_BUILD=${_COMPOSE} build --build-arg VERSION=${VERSION} --build-arg BASE_VERSION=${BASE_VERSION}

dev-up-app:
	${_COMPOSE_BUILD}
	BASE_VERSION=${BASE_VERSION} ${_COMPOSE} --profile app up -d $(NAMESPACE)-app

dev-down-app:
	${_COMPOSE} down -d $(NAMESPACE)-app

dev-down:
	${_COMPOSE} down --remove-orphans

dev-clean:
	${_COMPOSE} down --remove-orphans -v --rmi all

dev-restart: dev-down dev-up-app

dev-local-env:
	${_COMPOSE} --profile dependencies up -d

dev-migrate:
	${_COMPOSE} run --rm --service-ports $(NAMESPACE)-migrate

