# wiki-races

This is a complete rewrite of my wiki races game, based on what
I wish I did for my first design.

> First Design: <https://github.com/ElderINTERalliance/WikiRaces2021>

Given that this project uses Go and Vuejs, it's a little less
beginner friendly than the last version, which used only Javascript.

Notes:

- [server.go](main/server.go) serves the frontend in [frontend](frontend)

Todo:

- fix encoding issues in page generation
- add page viewer
- Cache files
- Add pin number for letting people into games
- add leaderboard
- add time sync for all clients
