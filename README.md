# Introduction 
This API functions as a tool to use to retrieve the exchange rates between a crypto symbol and another symbol (crypto or fiat). 
It uses the Binance market exchange rates so, there might be some exchange rates that does exist but isn't covered by the API. 
All exchange rates are fetched from the [CryptoCompareAPI](https://min-api.cryptocompare.com/documentation). 
Because of rate-limiting policies by crypto-compare, the data in the database is only updated once in 24 hours)


# Endpoints 
Postman [documentation can be found here](https://google.com)
- /v1/rates - fetches all rates 
- /v1/rates/:crypto-symbol - fetches all exchange rates for the specified symbol
- /v1/rates/:crypto-symbol/:fiat - fetches the exchange rate between the specified symbol and fiat 
- /v1/rates/history/:crypto-symbol/:fiat - retrieves the exchange rate between the symbol and fiat for the past 24 hours
- /v1/balance/:address - retrieves the current balance of a given ETH address


# Local Development
## Requirements
- Go >v1.19.x
## Setup 
- Clone the project and run `go mod tidy` in the project root folder to install packages 
- Create a .env file in the project root folder by copying the .env.example file and rename it to have a base env to work with. In the project root folder run `cp .env.example .env`
- Add your values to the .env file and you should be good to go. 
- To start up the api, you can run the api directly by running `APP_ENV=dev go run ./cmd/api` at the project root folder or build it first then run the build using the following commands
  - `go build -o=./bin/api ./cmd/api`
  - `APP_ENV=dev ./bin/api`

# Project structure 
## CMD/API 
This folder contains everything about the about, handlers, helpers and e2e tests for the API 
## Services 
This is folder where external services that can be re-used by different processes are created along with their tests 
## db 
This folder contains the database integrations. The reason for separating the db this way is to allow for the easy integration of other databases or overriding

# Testing 
- The project contains e2e tests that can be found in /cmp/api/test folder to test the APIs
- Tests for db can be found in the /db folder 
- Tests for each service can be found in /services folder 
- All test uses an in-memory database [memongo](https://github.com/tryvium-travels/memongo) so as to relieve the stress of having to create a separate test database
