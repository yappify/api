# Setting up development environment

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
