# syntax=docker/dockerfile:1
FROM golang:1.16.5-alpine
WORKDIR /server
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -o /server/build/server ./main
# RUN /server/build/server

CMD /server/build/server
