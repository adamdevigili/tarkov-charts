<h1>tarkov-charts</h1>

![Vercel](https://therealsujitk-vercel-badge.vercel.app/?app=tarkov-ammo-3d)

[**tarkov-charts**](https://www.tarkov-charts.com/) is a web application that provides a variety of graphs that relates live data about ammunition types, attachments, and other items in the game [Escape From Tarkov](https://www.escapefromtarkov.com/). 

Currently, the only graph is the [penetration, damage](https://escapefromtarkov.fandom.com/wiki/Ballistics#Armor_penetration_tables), and the cost of a single round type. These rounds are grouped by caliber, as all rounds in each caliber will have a similar "slope", usually the round with the highest penetration doing the least amount of damage, while also being the most expensive.

There are plenty of ammo charts that currently exist ([tarkov-tools](https://tarkov-tools.com/ammo/) being my personal favorite), however none exist that relate this data to the price of the ammo, one of the most important data points, at least for newer players (or at the beginning of a "wipe", where all players start from the beginning), since it's unlikely they will be able to afford the highest-tier ammo, and will need to find the best "bang-for-buck" ammo.

**NOTE**: As of 12.12 (December 12, 2021), many rounds have been removed from the flea market, leaving only trader prices, and eliminating a lot of the value of a graph like this. The future of this graph is undetermined at the moment.

<h2>Tech</h2>

- [Vercel](https://vercel.com/) for hosting/CD/secret management
- [React](https://reactjs.org/) + [PlotlyJS for React](https://plotly.com/javascript/react/)
- [Go](https://golang.org/) for lambda function to maintain data store and expose succinct API for web app
- [tarkov-tools.com](https://tarkov-tools.com/) GraphQL API for internal IDs and prices
- [tarkovdata](https://github.com/TarkovTracker/tarkovdata/) for item details 
- ~~[tarkov-market.com](https://tarkov-market.com/) REST API for price data~~ (deprecated)
- ~~[jsonbin.io](https://jsonbin.io/) to store JSON data~~ (deprecated)
- [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) to store aggregated data
- [GitHub Actions](https://github.com/features/actions) to periodically call Go API to update data store

[Lucidchart architecture diagram](https://lucid.app/lucidchart/invitations/accept/inv_30a42228-7a51-46c3-b983-e0d3dabc045a?viewport_loc=-89%2C-87%2C1558%2C1360%2C0_0)
