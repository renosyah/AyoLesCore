FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY . .
EXPOSE 8000
EXPOSE 80
CMD ./main --config=.heroku.toml
MAINTAINER syahputrareno975@gmail.com