FROM opensuse:latest

RUN useradd -d /app web
USER web
WORKDIR /app

COPY --chown=web:users guestbook-go /app/guestbook
COPY --chown=web:users public /app/public

EXPOSE 4000

ENTRYPOINT ["./guestbook"]
