# Distributed Lock Helper in Golang using Redis

A simple implementation of **Distributed Lock** in Golang using **Redis SETNX**.

This project demonstrates how to ensure that only one process can execute a specific task at a time, preventing race conditions in distributed systems.

---

##  Features

- Distributed Lock using Redis
- Atomic lock acquisition with `SETNX`
- Automatic lock expiration using TTL
- Simple callback-based API
- Easy to integrate into any Go project

---

##  Tech Stack

- Golang
- Redis
- go-redis/v9

---

## Project Structure

```text
distributed-lock-golang/
│
├── .env
├── go.mod
├── go.sum
├── main.go
│
└── helper/
    ├── distributed_lock.go
    ├── redis.go
    ├── response.go
    └── env.go
```

---

## ⚙️ Environment Variables

Create a `.env` file.

```env
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
```

---

## 📥 Installation

Clone the repository.

```bash
git clone https://github.com/dickydns/distributed-lock-golang
```

Go to the project.

```bash
cd distributed-lock-golang
```

Install dependencies.

```bash
go mod tidy
```

Run the application.

```bash
go run .
```

---

## 💻 Example

```go
lock := helper.NewDistributedLock(redis)
event:= "create:voucer"
result, err := lock.Execute(
    event,
    30*time.Second,
    func() (interface{}, error) {

        fmt.Println("Generating Event...")

        time.Sleep(3 * time.Second)

        fmt.Println("Event Generated")

        return "SUCCESS", nil
    },
)

if err != nil {
    fmt.Println(err)
    return
}

fmt.Println(result)
```

---

## 🔄 How It Works

```text
Request
    │
    ▼
SETNX
    │
 ┌──┴─────────┐
 │            │
Success     Failed
 │            │
 ▼            ▼
Execute    Return Error
 │
 ▼
Release Lock
```

---

## 🧠 Why Use SETNX?

`SETNX` stands for:

```
SET if Not Exists
```

Unlike using `GET` followed by `SET`, `SETNX` performs the check and write in a **single atomic operation**, preventing race conditions.

---

## ⏳ Why Use TTL?

Every lock has a Time-To-Live (TTL).

```go
30 * time.Second
```

The TTL acts as a safety mechanism.

If the application crashes before releasing the lock, Redis will automatically remove it after the TTL expires.

Without a TTL, a stale lock could block future requests indefinitely.

---

## Use Cases

- Payment Processing
- Inventory Updates
- Flash Sale
- Voucher Generation
- Scheduled Jobs
- Cron Jobs
- Data Synchronization
- Report Generation

---

## Related Articles

- Retry Pattern
- Circuit Breaker Pattern
- Idempotency
- Distributed Lock
- Worker Pool *(Coming Soon)*

---

##  Support

If you find this project useful, consider giving it a ⭐ on GitHub.

Feedback and contributions are always welcome!
