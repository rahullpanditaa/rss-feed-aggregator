# Gator - RSS-Feed-Aggregator (CLI)

Command-line RSS feed aggregator written in Go. Register/login, add and follow feeds, browse posts. A background worker continuously fetches and stores posts in PostgreSQL.

## Prerequisites
- **Go** 1.25+ (tested on 1.25.3)
- **PostgreSQL** 16.10
- **Goose** (DB migrations)
  - go install github.com/pressly/goose/v3/cmd/goose@latest

## Install
- go install github.com/rahullpanditaa/rss-feed-aggregator@latest
