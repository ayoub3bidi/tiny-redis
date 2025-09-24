# Tiny Redis

A toy Redis clone written in Go. Supports:
- RESP protocol (parsing and writing).
- Commands: `PING`, `SET`, `GET`.
- In-memory datastore.
- Append-Only File (AOF) persistence.

---

## 📌 Run the Server
```bash
go run .
````

Server runs on port **6379** (default Redis port).

---

## 📌 Test with redis-cli

Open another terminal:

```bash
redis-cli
```

### Example:

```
127.0.0.1:6379> PING
PONG

127.0.0.1:6379> SET foo bar
OK

127.0.0.1:6379> GET foo
"bar"
```

---

## 📂 Code Overview

* **main.go** → Starts TCP server, accepts connections.
* **resp.go** → Parses and serializes RESP protocol messages.
* **handler.go** → Handles commands (`PING`, `SET`, `GET`) and manages datastore.
* **aof.go** → Simple Append-Only File persistence.

---

## 📖 Learning Goals

This project teaches you:

* How Redis speaks with clients using RESP.
* How to implement a TCP server in Go.
* How to parse and generate protocols.
* Basics of persistence with append-only logs.

---

## ⚠️ Limitations

* Only `PING`, `SET`, `GET` supported.
* No expiry (like `SETEX`).
* AOF is append-only, not replayed at startup.
* Single-threaded datastore with global locks.
