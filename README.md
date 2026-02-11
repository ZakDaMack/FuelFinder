# FuelFinder

### Find the published site [here](https://fuelfinder.zakdowsett.co.uk/)

<p align="center">
  <img src="https://github.com/ZakDaMack/FuelFinder/blob/main/docs/logo.png?raw=true" alt="FuelFinder logo"/>
</p>

API and front end to show cheapest fuel in your area, provided by gov fuel price data. Note that this is an interim scheme and is limited, such as a delay between the setting of prices, the publication of price data, and the prices being available in third-party apps. Customers should always check the price displayed at the forecourt before purchasing road fuel.

The Department for Energy Security and Net Zero has stated it will publish its consultation on the end-state solution in March 2024. (Although I am unsure if this has come to fruition). The government eventually plans to make this madnatory for all brands as well as be a live data stream for developers to consume.

## Todo
- Add auth ready for mobile app
- Get averages in area
- Show station history
- Login screen

## Current Datasets
Datasets can be found on the government website [here](https://www.gov.uk/guidance/access-fuel-price-data). These are constantly updated and the station scraper script will collect any new entries that appear here.

## Local development
A docker compose file has been supplied in the `deployments/` folder. To run the project locally, run `docker compose -f deployments/docker-compose.yml up -d --build` and navigate to `localhost:8000`.

The react app can be found in the `web/` directory. To start up development, copy the `.env` file to connect to the docker stack for testing purposes. Run with `npm run start`.

The three backend microservices entrypoints can be found under the `cmd/` folder. These can be built into binaries using the `make build` command.

## Architecture
The project consists of a number of different projects, with the FuelFinder server connecting to a mongo server, reachable via gRPC. 
The scraper is timed to run every 15 minutes by default and will upload the data to the FuelFinder server, which can then be reachable at the web SPA via the api gateway service. (It would be nice to interact with the FuelFinder server directly at some point) 

![fuelfinder architecture](/docs/arch_drawing_dark.png)