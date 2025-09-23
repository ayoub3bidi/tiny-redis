# Tiny Redis (Redis Clone in Go)

A minimal Redis-like in-memory database written in Go.  
Supports a subset of Redis commands (`PING`, `SET`, `GET`) and communicates using the [RESP protocol](https://redis.io/docs/latest/develop/reference/protocol-spec/).

---

## 🚀 Features
- TCP server listening on port **6379** (default Redis port).
- RESP serialization/deserialization implemented from scratch.
- Basic commands:
  - `PING` → replies with `PONG`
  - `SET key value` → store a key-value pair
  - `GET key` → retrieve a value by key
- In-memory store using Go maps.
- Lightweight, under 300 lines of Go code.

---

## 📂 Project Structure
```

tiny-redis/
├── main.go      # TCP server & connection handler
├── resp.go      # RESP reader/writer
├── handler.go   # Command dispatcher (PING, SET, GET)
└── go.mod

````

*(A future `aof.go` file will add persistence.)*

---

## 🛠️ Installation & Running

Clone the repo and run:

```bash
git clone https://github.com/your-username/tiny-redis.git
cd tiny-redis
go run .
````

Server will start on:

```
TinyRedis listening on :6379
```

---

## 🧪 Testing with redis-cli

You can connect using the standard Redis CLI:

```bash
redis-cli ping
# -> PONG

redis-cli set mykey hello
# -> OK

redis-cli get mykey
# -> "hello"
```
