# CI Tools
ARG BASE_VERSION
FROM gotools.docker.lamoda.ru/base-mod:${BASE_VERSION} AS ci
ENV GOPRIVATE stash.lamoda.ru
ENV PROJECTDIR /usr/local/go/src/stash.lamoda.ru/neurocapsule/neurocapsule
WORKDIR ${PROJECTDIR}
COPY go.mod go.sum ${PROJECTDIR}/
RUN  go mod download
COPY . ${PROJECTDIR}/
RUN make setup

# Binary build
FROM ci AS build
ARG VERSION
RUN make VERSION=${VERSION} build

# Production image
FROM alpine:3
# Adding extra utils to image by OPS request: telnet, curl
# TODO: Move this installation to base alpine from OPS
RUN apk add --no-cache busybox-extras curl
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/go/src/stash.lamoda.ru/neurocapsule/neurocapsule/output /bin/output
CMD /bin/output
