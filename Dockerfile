# docker file for ayolescore app
FROM golang:latest as builder
ADD . /go/src/github.com/renosyah/AyoLesCore
WORKDIR /go/src/github.com/renosyah/AyoLesCore
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN rm -rf /vendor
CMD ./main --config=.heroku.toml
MAINTAINER syahputrareno975@gmail.com