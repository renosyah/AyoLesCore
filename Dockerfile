# docker file for ayolescore app
# build golang api
FROM golang:latest as builder
ADD . /go/src/github.com/renosyah/AyoLesCore
WORKDIR /go/src/github.com/renosyah/AyoLesCore
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/main /bin/main
# COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/.server.toml .
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/.heroku.toml .
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/sql /sql
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/template /template
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/files /files
EXPOSE 8000
CMD /bin/main
MAINTAINER syahputrareno975@gmail.com