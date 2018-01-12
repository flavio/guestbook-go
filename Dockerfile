FROM golang:1.8

COPY . /go/src/github.com/flavio/guestbook-go

WORKDIR /go/src/github.com/flavio/guestbook-go

RUN go build

RUN useradd -d /go/src/github.com/flavio/guestbook-go web

# Cannot use the --chown option of COPY because it's not supported by
# Docker Hub's automated builds :/
RUN chown -R web:users *

USER web

EXPOSE 4000

ENTRYPOINT ["./guestbook-go"]
