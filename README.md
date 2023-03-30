# Pandaroll - Rolling Your Migrations Made Easy
Are you tired of manually managing database migrations? Do you want a tool that makes the process as smooth as a panda rolling downhill? Look no further than Pandaroll!

Pandaroll is a powerful database migration tool that simplifies the process of managing database changes. With Pandaroll, you can roll your migrations with ease, whether you're migrating to a new database or updating an existing one.

The best part? Pandaroll is designed to be user-friendly and intuitive. You don't need to be a database expert to use it. In fact, even pandas could use it!

## Features
* Easy to use: With a simple command-line interface, Pandaroll is straightforward and easy to use.
* Flexible: Pandaroll can be used with a variety of databases, including MySQL, PostgreSQL, SQLite, and more. (soon)
* Safe: Pandaroll ensures that your data remains safe during migrations, with rollback functionality in case of any issues.

While we're only early days, the feature list is quite bare but we're working on getting it buffed up!

| Feature                            | Supported                    |
|------------------------------------|------------------------------|
| Run migrations                     | Yes                          |
| Create migrations                  | Yes                          |
| Rollback                           | Not yet                      |
| Multiple DBMS'                     | Not yet, just Postgres today |
| Disabling transaction in migration | Not yet                      |

## How to Use
Using Pandaroll is easy. Simply install it on your system, configure it with your database settings, and run the migration command. Pandaroll will handle the rest!

Here's an example of how to use Pandaroll:

```bash
$ pandaroll migrate
```
And that's it!

So, what are you waiting for? Don't let database migrations be a pain in the neck. Roll with Pandaroll, and let those pandas roll with you!

## Install

There are (soon to be) a number of ways you can install Pandaroll!

### Direct

You can directly download any version from the [Releases](https://github.com/BlobDevelopment/pandaroll/releases) page, we provide binaries for a lot of common platforms and architectures. This is the quickest way to get up and running.

### Homebrew (soon)

If you have [Homebrew](https://brew.sh/), you can install with

```bash
$ brew install ...
```

### Winget (soon)

Are you a Windows user and looking to use [Winget](https://learn.microsoft.com/en-us/windows/package-manager/winget/)? Well lucky you, we got a package there too!

```bash
$ winget ...
```

### Docker

We can't forget good old [Docker](https://www.docker.com/). See also [Running] for putting it into `docker-compose` or just running the image directly.

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
    volumes:
      # Mount the local "migrations" folder
      - ./migrations:/migrations
```
