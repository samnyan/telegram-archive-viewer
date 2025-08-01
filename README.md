# Telegram Archive Viewer

A Telegram Takeout viewer and hosting server.
Import all your chat history into a searchable database.

# Features

## Frontend

- [x] Load and view json format takeout data.
- [x] Basic text message styling support.
- [ ] Media preview.
- [ ] Json message uploader.

## Server

- [ ] Import json format takeout data.
- [ ] Message viewer frontend.
- [ ] Media file deduplication.
- [ ] Fuzzy search history.

# Development

## Requirements

- Install [Taskfile](https://taskfile.dev/installation/)

## Server

Golang API server.

### Run it

Install [Go](https://go.dev/dl/).

Copy `conf/config.example.yaml` to `conf/config.yaml` and edit it.

Run server:

```shell
task server
```

## Frontend

Vue.js frontend.

### Run it

Install [Node.js](https://nodejs.org/en/download/), [pnpm](https://pnpm.io/installation)

Then run

```shell
task web
```

or check `web/package.json` for npm scripts.

## Generate openapi code

After editing `api/openapi.yaml`, run

```shell
task server-oapigen
```

to generate server code.

```shell
task web-oapigen
```

to generate frontend code.
