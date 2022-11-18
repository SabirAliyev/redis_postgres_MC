# Caching in Go using Redis server.

## Preparations
Download the required packages (if first time start):

```
go mod download
```

Up the Docker Compose from the project directory:

```
docker-compose up
```

Configure your postgres (if first time start). User name, password, create database.
To enter to psql from CLI run:

```
psql -h 172.17.0.2 -p 5432 -U postgres -W
```
then enter user password.


----------------------
If you don't know postgres user password, then enter to the postgres container using docker:

```
docker exec -it postgres-demon bash 

psql -U postgres
```

To reset the password if you have forgotten:

```
ALTER USER user_name WITH PASSWORD 'new_password';
```
-----------------------

Get postgres container ID:

```
~$ docker container inspect <container-name> | grep -i IPAddress
```

## Running the app

Run `server.go` with the PostgreSQL and Redis configuration:

```
DATABASE_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable \
REDIS=localhost:6379 \
  go run .
```

To GET an exist entity from CLI:

```
curl -X GET http://0.0.0.0:8080/names/nm0000069
```

To check an entity in redis DB:

```
docker exec -it <container ID> redis-cli
```

Then:

```
keys *
```

To get decoded entity:

```
get nm0000069
```
