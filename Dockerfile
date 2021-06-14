# build
FROM golang:1.16.5-alpine3.13 AS build
COPY ./src /src
WORKDIR /src
ENV "GOPROXY" "https://goproxy.cn,direct"
RUN go build -o /build/app

# iamge
FROM alpine:latest
COPY --from=build /build/app /bin/app
RUN mkdir /env
WORKDIR /env
ENTRYPOINT [ "/bin/app" ]