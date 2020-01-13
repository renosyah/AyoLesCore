# docker file for ayolescore app
# build golang api
FROM golang:latest as builder
ADD . /go/src/github.com/renosyah/AyoLesCore
WORKDIR /go/src/github.com/renosyah/AyoLesCore
RUN go get -u github.com/golang/dep/cmd/dep
COPY . .
RUN dep ensure -v
RUN CGO_ENABLED=0 GOOS=linux go build -o main .
RUN rm -rf /api
RUN rm -rf /auth
RUN rm -rf /cmd
RUN rm -rf /model
RUN rm -rf /router
RUN rm -rf /vendor
RUN rm -rf /util
RUN rm .dockerignore
RUN rm .gitignore
RUN rm .server.toml
RUN rm Dockerfile
RUN rm Gopkg.lock
RUN rm Gopkg.toml
RUN rm README.md
RUN rm heroku.yml
RUN rm main.go


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ADD . /sql
ADD . /template
ADD . /files
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/main .
# COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/.server.toml .
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/.heroku.toml .
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/sql /sql
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/template /template
COPY --from=builder /go/src/github.com/renosyah/AyoLesCore/files /files
EXPOSE 8000
EXPOSE 80
CMD ./main --config=.heroku.toml
MAINTAINER syahputrareno975@gmail.com