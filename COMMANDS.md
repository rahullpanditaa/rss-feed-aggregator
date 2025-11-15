# COMMANDS.md

This document lists all available commands for the **RSS Feed Aggregator**, including a short description and usage example for each. Use this as a quick reference guide while working with the CLI.

---

# 📌 User Commands

## **register**

Create a new user.

```
register <username>
```

Example:

```
go run . register rahul
```

---

## **login**

Log in as an existing user.

```
login <username>
```

Example:

```
go run . login rahul
```

---

## **users**

List all users stored in the database.

```
users
```

---

## **reset**

Reset/drop all application data. Use with caution.

```
reset
```

---

# 📡 Feed Commands

## **addfeed**

Add a new RSS feed (automatically followed by the current user).

```
addfeed <name> <url>
```

Example:

```
go run . addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
```

---

## **feeds**

List all feeds stored in the system.

```
feeds
```

---

## **follow**

Follow an existing feed by URL.

```
follow <url>
```

Example:

```
go run . follow https://blog.boot.dev/index.xml
```

---

## **unfollow**

Unfollow a feed by URL.

```
unfollow <url>
```

Example:

```
go run . unfollow https://blog.boot.dev/index.xml
```

---

## **following**

Display all feeds followed by the current user.

```
following
```

---

# 📰 Aggregation & Browsing

## **agg**

Run the background feed scraper on an interval.

```
agg <duration>
```

Example:

```
go run . agg 30s
```

---

## **browse**

Display recent posts from feeds the user follows.

```
browse <limit>
```

Example:

```
go run . browse 10
```

---

# 🆘 Help

## **help**

Show a summary of all available commands.

```
help
```

---

# ✔️ Summary

This document provides a quick overview of:

* user management commands
* feed management commands
* aggregation commands
* browsing commands

Use this as your command reference while exploring the RSS Feed Aggregator.
