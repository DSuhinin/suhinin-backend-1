#
# Build arguments.
#
ARG GO_SERVICE_IMPORT_PATH=unknown
ARG GOPROXY=$GOPROXY
ARG GO111MODULE=$GO111MODULE
ARG PATH_GO_SOURCES=/go/src/$GO_SERVICE_IMPORT_PATH

#
# Build Go binary inside base container.
#
FROM golang:1.14 as base_go_docker_image
# Stage arguments
ARG PATH_GO_SOURCES
ARG GO_SERVICE_IMPORT_PATH
# Env variables.
ENV GOPROXY=$GOPROXY
ENV GO111MODULE=$GO111MODULE

# Create sources directory inside the container and copy project files.
RUN mkdir -p $GO_SERVICE_IMPORT_PATH/
WORKDIR $PATH_GO_SOURCES
COPY . $PATH_GO_SOURCES
# Build
RUN make go_get
RUN make go_build

# Container settings.
EXPOSE 8080
ENTRYPOINT ["./test-task-service"]
