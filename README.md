# Yappify API

[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/yappify/template/blob/main/LICENSE) ![ci-badge](https://github.com/yappify/api/actions/workflows/cicd.yml/badge.svg) ![Go](https://img.shields.io/badge/Go-blue.svg?style=flat&logo=go&logoColor=white)

Lead maintainer: [@abyanmajid](https://github.com/abyanmajid) \
Documentation: [github.com/yappify/api/tree/main/docs](https://github.com/yappify/api/tree/main/docs)

This source code makes up the RESTful API that serves the backend server of *Yappify*. it is a monolith written in *Go*, with the *chi* router, *PostgreSQL* database, and *sqlc* ORM.

## Contributing

1. Quickly set up development environment by running the following
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

You can alternatively run the server in a docker container. To do this, use:
1. `make image` to build a docker image of the server
2. `make image-up` to start the server in a docker container
3. `make image-down` to stop the container.

You can now start making modifications. For more information, please consult the [documentation.](https://github.com/yappify/api/tree/main/docs)
