```bash
# Start a PostgreSQL server using Docker
docker network create psql
docker run --name pgsrv --network psql -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres:alpine

# Start a PostgreSQL client to connect to the server
docker run --name pgsql -it --rm --network psql postgres:alpine psql -h pgsrv -U postgres

# Start a PostgreSQL client to connect to the server/database heroes
docker run --name pgsql -it --rm --network psql postgres:alpine psql -h pgsrv -U postgres --dbname=heroes

export POSTGRES_DSN=postgres://postgres:mysecretpassword@localhost/heroes?sslmode=disable
```

```sql
create database heroes;
\c heroes;
exit;
```

Connection string: `postgres://postgres:mysecretpassword@localhost/heroes?sslmode=disable`

Tool for migrations: [`migrate`](https://github.com/golang-migrate/migrate)

```bash
# Create SQL scripts for first migration
migrate create -seq -ext=.sql -dir=./migrations create_heroes_table

# Create SQL script for second migration
migrate create -seq -ext=.sql -dir=./migrations heroes_check_constraint
```