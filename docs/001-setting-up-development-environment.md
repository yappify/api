# Setting up development environment

Table of contents:
- [Quickstart](#quickstart)
- [Running server in a docker container](#running-server-in-a-docker-container)
- [Viewing the database](#viewing-the-database)
- [A note on Makefile](#a-note-on-makefile)

## Quickstart

1. Set up development environment by running the following
```
mkdir yappify-api && cd yappify-api
git clone https://github.com/yappify/api.git .
cp .env.default .env && rm .env.default
```

2. Start up docker desktop and start a postgres service in the background
```
make db-up
```

3. Now, run the server
```
make run
```

4. To test that the server is up and running, you can perform a health check to `localhost:8000/health`
```
curl localhost:8000/health
```

## Running server in a docker container

First build the docker image for the server by running `make image`. Then, use `make image-up` to run the container, and `make image-down` to stop it.

## Viewing the database

### Using `psql`

You can use the `psql` command line utility to view and query the database. To install `psql`, run
```
sudo apt-get install postgresql-client
```

Then, open the database via `psql`
```
make db-psql
```

### Using database clients

You can use database clients such as [pgAdmin](https://www.pgadmin.org/download/) and [Beekeeper Studio](https://www.beekeeperstudio.io/get-community). Note that whenever you are asked to enter username, password, or database name, know that as per `docker/docker-compose.yml`, you should input:

- Username: `postgres`
- Password: `postgres`
- Database: `db`

#### Using pgAdmin as a docker service

Alternatively, one of the services in `docker/docker-compose.yml` is web-based pgAdmin which, after running `make db-up`, can be accessed at `localhost:80`. The login details are as follows:

- Email: `admin@admin.com`
- Password: `admin`

## A note on Makefile

There is likely a `make` script you can run to perform common commands. You can list all defined `make` scripts by running
```
make help
```

If you repeatedly had to use run a certain command that is not yet defined in the Makefile, please do add that command and update the `help` target.
