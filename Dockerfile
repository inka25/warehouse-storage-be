### STAGE 1 : Build the go source code into binary
FROM golang:1.14 as builder
EXPOSE 8080:8080

ENV APP_DIR /warehouse-storage-be
#ENV GOFLAGS -mod=vendor

## Copy source code from local machine into container
RUN mkdir -p ${APP_DIR}
COPY . ${APP_DIR}

# Compile the binary and statically link
RUN cd $APP_DIR && CGO_ENABLED=0 go build -ldflags '-d -w -s'

