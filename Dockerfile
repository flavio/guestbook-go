FROM opensuse:latest

RUN useradd -d /app web
WORKDIR /app

COPY guestbook-go /app/guestbook
COPY public /app/public

# Cannot use the --chown option of COPY because it's not supported by
# Docker Hub's automated builds :/
RUN chown -R web:users *

USER web

EXPOSE 4000

ENTRYPOINT ["./guestbook"]
