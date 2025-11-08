# 🧠 learning-Go  
*A hands-on journey to master Go, one small program at a time.*

---

### 📌 Overview  
This repo documents my **progressive learning path in Go (Golang)** — moving from basics to real-world backend patterns.  
Each folder contains a self-contained Go module (`go mod init` per folder) that can be executed independently.

Think of it as a **developer’s playground** — not just code dumps, but small, well-documented examples that print both results **and** reasoning.

---

### 🗂️ Folder Structure  

| Folder | Topic | Description |
|---------|--------|-------------|
| `01_hello_world` | Hello World | Basic Go setup and first program |
| `02_variables` | Variables & Constants | `var`, `const`, and `:=` usage with examples and reasoning |
| `03_functions` | Functions | Basic, multiple returns, named returns, variadic, closures, defer, and error handling |
| `04_conditionals` | (coming soon) | `if`, `switch`, scope, and idiomatic branching |
| `05_loops` | (coming soon) | Using `for` as the universal looping construct |
| `06_arrays_slices` | (planned) | Array vs slice behavior and memory implications |
| `07_maps` | (planned) | Key-value structures and idiomatic iteration |
| `15_goroutines`+ | (future) | Concurrency, channels, and async processing |
| `19_rest_api` | (planned) | Basic REST API using `net/http` |
| `30_blog_api` | (future) | Full CRUD backend with GORM & SQLite |

---

### ⚙️ Getting Started  

#### 1️⃣ Clone the Repo  
```bash
git clone https://github.com/soummyaanon/learning-Go.git
cd learning-Go

2️⃣ Run Any Example

Each folder is an independent Go module. For example:

cd 02_variables
go run main.go

You’ll see output along with reasoning for every concept.

🧭 Learning Philosophy

“Don’t just read Go — run it until it feels obvious.”

Each example is designed to be:

Self-explanatory → prints both the result and why it works that way.

Incremental → builds on the previous topic.

Practical → mirrors real Go patterns you’ll use in backends, CLI tools, or agents.

🧩 How to Contribute / Extend

If you want to extend this repo or learn alongside:

Fork this repo.

Create a new folder with the next sequential number (e.g. 04_conditionals).

Add your main.go with concise examples and explanations.

Run:

go mod init learning-go-04
go run main.go


Keep examples clean, runnable, and reason-oriented.

🧠 Future Topics

error wrapping and custom error types

Goroutines & Channels deep dive

REST API patterns (net/http, chi, gin)

ORM with gorm and SQLite/Postgres

Unit testing with testing package

Dockerizing Go apps

Building CLI tools

🧰 Quick Info






📜 License

MIT License © 2025 Soumyaranjan

💬 Personal Note

This repo is not about speed — it’s about clarity.
Each folder represents a mental milestone — from

“What does := even mean?”
to
“How do goroutines talk to each other?”
