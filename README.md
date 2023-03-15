# Pandaroll

TODO

## Running

### Docker

Pandaroll is published on [Docker Hub](https://hub.docker.com/r/blobdev/pandaroll) for easy use.

Running standalone:

```bash
$ docker run blobdev/pandaroll:v1 migrate
```

Running with `docker-compose.yml`

```yml
version: '3'
services:
  db:
    container_name: db
    image: postgres:14.7-alpine
    environment:
      - POSTGRES_DB=test_database
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=password
    ports:
      - '5432:5432'

  db-migrator:
    container_name: migrator
    image: blobdev/pandaroll:v1
    depends_on:
      - db
    environment:
      - DBMS=postgres
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USERNAME=postgres
      - DB_PASSWORD=password
      - DB_DATABASE=test_database
```