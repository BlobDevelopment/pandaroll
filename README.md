# Pandaroll

Rollforwards or rollbackwards with Pandaroll! The easy to use database migration tool. Are you tired of writing a new script every time you need to setup DB migrations or using tools that are old, slow and just not nice to work with? Pandaroll is here for you :)

The goal of this application is to make database migrations easy, no fuss, just get up and go. No matter if it's testing locally, running in CI or in production.

## Features

While we're only early days, the feature list is quite bare but we're working on getting it buffed up!

| Feature                            | Supported                    |
|------------------------------------|------------------------------|
| Run migrations                     | Yes                          |
| Create migrations                  | Yes                          |
| Rollback                           | Not yet                      |
| Multiple DBMS'                     | Not yet, just Postgres today |
| Disabling transaction in migration | Not yet                      |

## Install

There are (soon to be) a number of ways you can install Pandaroll!

### Direct

You can directly download any version from the [Releases] page, we provide binaries for a lot of common platforms and architectures. This is the quickest way to get up and running.

### Homebrew (soon)

If you have Homebrew, you can install with

```bash
$ brew install pandaroll
```

### Winget (soon)

Are you a Windows user and looking to use [Winget]()? Well lucky you, we got a package there too!

```bash
$ winget ...
```

### Docker

We can't forget good old [Docker](). See also [Running] for putting it into `docker-compose` or just running the image directly.

```bash
$ docker pull blobdev/pandaroll:1
```

## Running

### Docker

Pandaroll is published on [Docker Hub](https://hub.docker.com/r/blobdev/pandaroll) for easy use.

Running standalone:

```bash
$ docker run blobdev/pandaroll:1 migrate
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
    image: blobdev/pandaroll:1
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