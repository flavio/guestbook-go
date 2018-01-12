Simple guestbook web application built using Go and Vue.js.

Uses Redis to store the data.

Meant to be used as a demo of a simple containerized workload for Kubernetes.

## Environment variables

  * `GUESTBOOK_REDIS_SOCKET`: path to the unix socket used by the Redis database.
  * `GUESTBOOK_REDIS_HOST`: host running the Redis database.
  * `GUESTBOOK_REDIS_PORT`: port used by the Redis database.
  * `GUESTBOOK_PORT`: port used by the guestbook web application.

By default the guestbook will listen on port 4000 and will attempt to connect to
a Redis database running on its standard port on localhost.
