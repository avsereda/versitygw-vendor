FROM golang:1.24.6 AS build
WORKDIR /build
COPY . /build
RUN /bin/sh -c '(cd /build/cmd/versitygw && go build)'
FROM debian:latest

RUN apt -y update && apt -y upgrade && apt -y autoclean
WORKDIR /
COPY --from=build /build/cmd/versitygw/versitygw /versitygw
COPY . /usr/src/versitygw
COPY start.sh /
EXPOSE 7071/tcp
ENTRYPOINT [ "/start.sh" ] 