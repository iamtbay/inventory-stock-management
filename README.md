# 📦 Inventory & Order Management System

A robust, high-performance RESTful API for managing products and orders, built with **Go (Golang)** using **Clean Architecture** principles.

## 🚀 Overview

This project simulates a backend system for an e-commerce platform. It handles product inventory management and order processing with strict concurrency controls and data integrity.

It is designed to demonstrate advanced backend concepts such as **Dependency Injection**, **Graceful Shutdown**, **Database Migrations**, and **Test-Driven Development (TDD)**.

## 🛠️ Tech Stack

* **Language:** Go (1.23)
* **Database:** PostgreSQL (with `pgxpool` for connection pooling)
* **Architecture:** Clean Architecture (Ports & Adapters)
* **Infrastructure:** Docker & Docker Compose
* **Documentation:** Swagger (OpenAPI 2.0)
* **Migrations:** Golang-Migrate
* **Testing:** Go Standard Library & Mocking

## 🏗️ Architecture

The project follows the **Clean Architecture** pattern to ensure separation of concerns and testability.

```text
├── cmd/             # Entry point (Main)
├── internal/
│   ├── domain/      # Core business logic (Entities & Interfaces)
│   ├── service/     # Business logic implementation
│   └── adapters/    # Database, HTTP Handlers, Router
├── migrations/      # SQL Migration files
└── docs/            # Swagger documentation
```
## Key Features
* Atomic Transactions: Ensures stock is reserved only when an order is successfully created.
* Concurrency Safe: Prevents race conditions during high-load stock updates using pgxpool.
* Graceful Shutdown: Ensures ongoing requests are completed before the server stops to prevent data loss.
* Context Propagation: Efficiently manages request timeouts and cancellation across layers.
* Database Migrations: Version control for database schema changes using golang-migrate.
* API Documentation: Interactive API documentation served via Swagger UI.

## ⚙️ How to Run
### Prerequisites
Docker & Docker Compose

### Quick Start
Clone the repository:

```bash
git clone [https://github.com/YOUR_USERNAME/inventory-system.git](https://github.com/YOUR_USERNAME/inventory-system.git)
cd inventory-system
```
Start the application (DB + App + Migrations):

```bash
docker-compose up --build
```
Access the API:

API: http://localhost:8080

Swagger Docs: http://localhost:8080/swagger/index.html

## 🧪 Running Tests
Unit tests use Mocks to isolate the business logic from the database, ensuring fast and reliable testing.

```bash
go test ./internal/service/... -v
```

## 🧠 Challenges & Learnings
Challenge: Implementing Clean Architecture was initially challenging, particularly strictly separating the concerns between layers (Service, Repository, Handler).

Solution: I used Dependency Injection and defined Interfaces in the domain layer. This decoupled the business logic from the database implementation. This approach provided me with a solid framework for building scalable and maintainable systems, proving that good architecture is crucial for larger projects.