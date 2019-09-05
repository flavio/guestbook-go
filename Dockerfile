FROM golang:1.13-alpine3.10

COPY . /code/guestbook-go
WORKDIR /code/guestbook-go
RUN CGO_ENABLED=0 go build -mod vendor


FROM alpine:3.10
WORKDIR /app
RUN adduser -h /app -D web
COPY --from=0 /code/guestbook-go/guestbook-go /app/guestbook
COPY --from=0 /code/guestbook-go/public /app/public

## Cannot use the --chown option of COPY because it's not supported by
## Docker Hub's automated builds :/
RUN chown -R web:web *
USER web
ENTRYPOINT ["./guestbook"]
EXPOSE 4000
