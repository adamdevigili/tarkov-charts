<h1>tarkovammo3D</h1>

![Vercel](https://therealsujitk-vercel-badge.vercel.app/?app=tarkov-ammo-3d)

[**tarkovammo3D**](https://www.tarkovammo3d.com/) is a web-based, 3-dimensional scatter plot that relates live data about ammounition types in the game [Escape From Tarkov](https://www.escapefromtarkov.com/). 

This data is the [penetration, damage](https://escapefromtarkov.fandom.com/wiki/Ballistics#Armor_penetration_tables), and the cost of a single round type. These rounds are grouped by caliber, as all rounds in each caliber will have a similar "slope", usually the round with the highest penetration doing the least amount of damage, while also being the most expensive.

There are plenty of ammo charts that currently exist ([tarkov-tools](https://tarkov-tools.com/ammo/) being my personal favorite), however none exist that relate this data to the price of the ammo, one of the most important data points, at least for newer players (or at the beginning of a "wipe", where all players start from the beginning), since it's unlikely they will be able to afford the highest-tier ammo, and will need to find the best "bang-for-buck" ammo.

<h2>Tech</h2>

- [Vercel](https://vercel.com/) to for hosting/secret management
- [React](https://reactjs.org/)/[PlotlyJS for React](https://plotly.com/javascript/react/)
- [Go](https://golang.org/) for lambda function to maintain data store
- [tarkov-tools.com](https://tarkov-tools.com/) GraphQL API for internal ammo IDs (so we don't have to hardcode & update when new ammo types get added)
- [tarkov-market.com](https://tarkov-market.com/) REST API for price data
- [jsonbin](https://jsonbin.io/) to store JSON data
- [GitHub Actions](https://github.com/features/actions) to periodically call Go lambda function to update data store