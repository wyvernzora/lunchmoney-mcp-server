# lunchmoney-mcp-server

## Overview

`lunchmoney-mcp-server` is an [Model Context Protocol (MCP)](https://modelcontextprotocol.io) server built
using [mcp-go](https://github.com/mark3labs/mcp-go) to provide LLMs ability to ingest category, transaction
and budget data from [Lunch Money](https://lunchmoney.app/).

## Configuration
You can configure the server using the following environment variables

| Name               | Default   | Description                               |
| ------------------ | --------- | ----------------------------------------- |
| `BIND_ADDRESS`     | `0.0.0.0` | The IP address that the server listens on |
| `PORT`             | `3000`    | The port that the server listens on       |
| `LUNCHMONEY_TOKEN` |           | The Lunch Money authorization token       |

## Usage
```
$ export LUNCHMONEY_TOKEN='your-token-here'
$ docker run -p 3000:3000 -e LUNCHMONEY_TOKEN ghcr.wyvernzora.io/lunchmoney-mcp-server:latest
```

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
