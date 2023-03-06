#!/bin/bash

go build -o ./bin/pandaroll

DBMS=postgres \
    DB_HOST=0.0.0.0 \
    DB_PORT=5432 \
    DB_USERNAME=postgres \
    DB_PASSWORD=password \
    DB_DATABASE=test \
    ./bin/pandaroll $*