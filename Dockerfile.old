FROM golang:1.16 AS build
ADD . /src
WORKDIR /src
RUN go build -v -o /usr/local/bin/api
RUN chmod +x /usr/local/bin/api
ENTRYPOINT /usr/local/bin/api --port 8080 --host="0.0.0.0"
EXPOSE 8080
