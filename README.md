# omiga
Omiga

# Code Coverage
[![codecov](https://codecov.io/gh/omiga-group/omiga/branch/main/graph/badge.svg?token=R1BCXGRWWM)](https://codecov.io/gh/omiga-group/omiga)

# How to re-generate code
```bash
cd ./src
make generate
```

# How to start

## Start everything

Please note that starting dependencies only won't create database and migrate the schema

```bash
./src/scripts/run-docker-compose.sh up --build
```

## Start dependencies only

```bash
./src/scripts/start-dependencies.sh up --build
```

## Start `venue` domain and dependencies only

```bash
./src/venue/scripts/run-docker-compose.sh up --build
```

## Start `order` domain and dependencies only

```bash
./src/order/scripts/run-docker-compose.sh up --build
```
