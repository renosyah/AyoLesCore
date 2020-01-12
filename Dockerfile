# docker file for ayolescore app
FROM golang:latest
ADD . /go/src/github.com/renosyah/AyoLesCore
WORKDIR /go/src/github.com/renosyah/AyoLesCore
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure -v
RUN go install
EXPOSE 8000
CMD /go/bin/AyoLesCore
MAINTAINER syahputrareno975@gmail.com