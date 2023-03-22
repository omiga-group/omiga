# Omiga
Omiga Crypto Product Finder

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

# Database Structure

## Products Table

| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | integer | primary key, auto increment |
| name | varchar(255) | not null |
| description | text | |
| type | varchar(50) | not null |
| yield | decimal(5,2) | |
| risk | decimal(5,2) | |
| liquidity | decimal(5,2) | |
| fees | decimal(5,2) | |

## Cryptocurrencies Table

| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | integer | primary key, auto increment |
| name | varchar(255) | not null |
| symbol | varchar(10) | not null |
| supported_products | text | |

## Exchanges Table

| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | integer | primary key, auto increment |
| name | varchar(255) | not null |
| country | varchar(50) | |
| supported_cryptocurrencies | text | |
| url | varchar(255) | not null |

## Users Table

| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | integer | primary key, auto increment |
| name | varchar(255) | not null |
| email | varchar(255) | not null |
| investment_amount | decimal(10,2) | |
| risk_tolerant | boolean | |
| experience | varchar(50) | |
| country | varchar(50) | |

## User Responses Table

| Column Name | Data Type | Constraints |
| --- | --- | --- |
| id | integer | primary key, auto increment |
| user_id | integer | foreign key to Users Table |
| question | varchar(255) | |
| answer | text | |

# Questions

Can we use GraphQL Nexus or GraphQL Yoga or Hasura to auto-generates a GraphQL API based on this database schema?