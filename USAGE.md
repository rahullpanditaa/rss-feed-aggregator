# USAGE.md

This document provides practical examples of how to use the **RSS Feed Aggregator** from the command line. It covers all supported commands with explanations and real-world usage patterns.

The aggregator is designed to be simple, predictable, and easy to work with in a terminal environment.

---

# 🧑‍💻 Prerequisites

Before using the CLI, ensure:

* PostgreSQL is running
* Your config file exists at:

```
~/.rssfeedconfig.json
```

* The `db_url` field points to your PostgreSQL instance
* Migrations have been applied:

```sh
goose up
```

---

# 🚀 Basic Workflow Overview

A typical interaction looks like:

1. Register a user
2. Log in
3. Add an RSS feed (auto-followed)
4. Follow/unfollow additional feeds
5. Start the aggregator loop
6. Browse posts in another terminal

---

# 📌 Command Reference & Examples

## 1. **Register a New User**

```sh
go run . register <username>
```

Example:

```sh
go run . register rahul
```

Output:

```
Created user: rahul
User ID: ...
Username: rahul
```

---

## 2. **Log In**

```sh
go run . login <username>
```

Example:

```sh
go run . login rahul
```

Output:

```
app state has been set to given user: rahul
```

---

## 3. **List All Users**

```sh
go run . users
```

Shows every user registered in the system.

---

## 4. **Add a New Feed**

```sh
go run . addfeed <name> <url>
```

Example:

```sh
go run . addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
```

This:

* Inserts the feed into the database
* Automatically follows the feed for the current user

---

## 5. **List All Feeds**

```sh
go run . feeds
```

Shows every feed known to the system.

---

## 6. **Follow a Feed**

```sh
go run . follow <url>
```

Example:

```sh
go run . follow https://blog.boot.dev/index.xml
```

Follows an existing feed by URL.

---

## 7. **Unfollow a Feed**

```sh
go run . unfollow <url>
```

Example:

```sh
go run . unfollow https://blog.boot.dev/index.xml
```

---

## 8. **See Feeds You Are Following**

```sh
go run . following
```

Displays the list of feeds associated with your user.

---

## 9. **Browse Recent Posts**

```sh
go run . browse <limit>
```

Example:

```sh
go run . browse 10
```

Output example:

```
2025-02-10 12:15:03 | Intro to Rust
https://blog.boot.dev/rust-intro

2025-02-09 18:20:12 | How Go Interfaces Work
https://blog.boot.dev/go-interfaces
```

---

## 10. **Start the Aggregator Loop**

```sh
go run . agg <duration>
```

Example:

```sh
go run . agg 30s
```

This starts a long-running process that:

* Selects the next feed to fetch
* Downloads and parses RSS XML
* Saves posts into the database
* Repeats every N seconds

You typically run this in **Terminal 1**.

---

## 11. **Browse While Aggregator Runs**

In **Terminal 2**:

```sh
go run . browse 10
```

This lets you see newly scraped posts as they appear.

---

## 12. **Reset the Database**

```sh
go run . reset
```

Drops all data in:

* users
* feeds
* feed_follows
* posts

Use with caution.

---

## 13. **Help Command**

If implemented:

```sh
go run . help
```

Shows all available commands.

---

# 📝 Example Full Session

A realistic usage flow:

```sh
# Register and log in
go run . register rahul
go run . login rahul

# Add and follow a feed
go run . addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml

# Start aggregator loop (terminal 1)
go run . agg 15s

# Browse posts (terminal 2)
go run . browse 5

# Follow another feed
go run . follow https://example.com/rss.xml

# See feeds you follow
go run . following

# Unfollow
go run . unfollow https://example.com/rss.xml
```

---

# 🎉 Summary

This usage guide provides everything needed to:

* create users
* manage feeds
* follow/unfollow feeds
* scrape posts automatically
* browse aggregated content conveniently

The CLI is designed to be ergonomic, predictable, and easily extendable.
