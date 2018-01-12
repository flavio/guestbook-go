FROM golang:1.8-alpine

COPY . /go/src/github.com/flavio/guestbook-go
WORKDIR /go/src/github.com/flavio/guestbook-go
RUN go build


FROM alpine
WORKDIR /app
RUN adduser -h /app -D web
COPY --from=0 /go/src/github.com/flavio/guestbook-go/guestbook-go /app/guestbook
COPY --from=0 /go/src/github.com/flavio/guestbook-go/public /app/public

## Cannot use the --chown option of COPY because it's not supported by
## Docker Hub's automated builds :/
RUN chown -R web:web *
USER web
ENTRYPOINT ["./guestbook"]
EXPOSE 4000
