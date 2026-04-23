#!/bin/bash
# Creating a docker database

docker run -d 
    --name auth-db
    -p 5432:5432
    -e POSTGRES_USER= admin
    -e POSTGRES_PASSWORD= pass
    -e POSTGRES_DB= authdb
    postgres

PAUSE