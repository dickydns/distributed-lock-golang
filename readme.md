# Go Idempotency Helper with Redis

A simple reusable Idempotency Helper written in Go using Redis.

This project demonstrates how to prevent duplicate requests using Redis, `SETNX`, and response caching.

## Features

- Redis Connection
- Idempotency Key
- Distributed Lock (SETNX)
- Response Cache
- TTL Support
- Reusable Helper
- JSON Response
- Environment Configuration (.env)

---

## Project Structure

```text
go-idempotency-redis/
│
├── helper/
│   ├── env.go
│   ├── redis.go
│   ├── response.go
│   └── idempotency.go
│
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
└── README.md
```

---

## Installation

Clone repository

```bash
git clone https://github.com/dickydns/distributed-lock-golang
```

Go to project

```bash
cd distributed-lock-golang
```

Install dependencies

```bash
go mod tidy
```

---

## Environment

Create a `.env` file.

```env
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
```

---

## Run

```bash
go run main.go
```

---

## Example

```go
key_from_frontend := "user:register:123"
result, err := idempotency.Execute(
    key_from_frontend,
    30*time.Second,
    func() (interface{}, error) {

        fmt.Println("Creating User...")

        return helper.SuccessResponse(map[string]interface{}{
            "id":   1,
            "name": "User 123",
        }), nil
    },
)
```

---

## Flow

```text
                Client
                   │
                   ▼
         Check Cached Response
                   │
          ┌────────┴────────┐
          │                 │
        HIT               MISS
          │                 │
          ▼                 ▼
 Return Cached       Acquire Lock (SETNX)
  Response                │
                    ┌─────┴─────┐
                    │           │
              Lock Failed   Lock Success
                    │           │
                    ▼           ▼
           Wait Response    Execute Callback
                    │           │
                    ▼           ▼
          Read Cached     Save Response
             Response          │
                    └──────────┘
                         │
                         ▼
                   Return Response
```

---

## How It Works

### First Request

1. Check cached response.
2. Cache not found.
3. Acquire Redis lock using `SETNX`.
4. Execute business logic.
5. Store response in Redis.
6. Release lock.
7. Return response.

---

### Duplicate Request

1. Check cached response.
2. Cache found.
3. Return cached response immediately.

---

## Why Use Idempotency?

Without Idempotency

```
POST /payment

↓

Payment Created

↓

POST /payment

↓

Payment Created Again 
```

With Idempotency

```
POST /payment

↓

Payment Created

↓

POST /payment

↓

Return Existing Response 
```

---

## Common Use Cases

- Payment
- Checkout
- Order
- Registration
- Top Up
- Reward Claim
- Voucher Redemption

---

## Technologies

- Golang
- Redis
- SETNX
- JSON
- Environment Variables

---

## Author

Dicky Perdian

GitHub

https://github.com/dickydns