# Loan Billing System

## Description
The **Loan Engine System** is a backend service built to manage customer loans, repayment schedules, and payments for amartha finance
It provides:
- Loan creation and retrieval
- Automatic calculation of outstanding balances
- Detection of delinquent loans

## Assumption
The system assume or out of scope to be build :
- customer did not need to use credential
- loan can be weekly or monthly
- for repayment schedule it only have due date not invoice date
- payment can be made at any time (without penalty fee)
- delinquent status happened when there is two record for certain loan in repayment_schedule that already pass the due date
- outstanding amount calculate based on unpaid data in repayment_schedule

---

## Tech Stack
- **Language:** Go (Golang)
- **Framework:** Gin (HTTP router)
- **ORM:** GORM
- **Database:** PostgreSQL
- **Containerization:** Docker & Docker Compose
- **Testing:** GoMock & Testify
- **Dependency Management:** Go Modules

---

## How to Run (Docker)
### 1. Clone the repository
```bash
git clone https://github.com/yourusername/loan-billing-system.git
cd loan-billing-system
```

### 2. Build & Run with Docker Compose
```bash
docker-compose up --build
```

This will:
- Start **PostgreSQL** on port `5432`
- Start **Loan Engine** on port `8080`

## System Design

```mermaid
erDiagram
    CUSTOMER {
        uint id PK
        string name
        string email
        string phone
        time created_at
    }

    LOAN {
        uint id PK
        uint customer_id FK
        uint loan_prodocut_id FK
        int64 principal
        float64 interest_rate
        int64 total_amount
        int repayment_count
        string repayment_frequency
        time start_date
        time created_at
    }

    REPAYMENT_SCHEDULE {
        uint id PK
        uint loan_id FK
        int sequence
        int64 amount
        bool is_paid
        time due_date
        time paid_at
        time created_at
    }

    PAYMENT {
        uint id PK
        uint loan_id FK
        int64 amount
        string status
        string payment_method
        time paid_at
        time created_at
    }

    LOAN_PRODUCT {
        uint id PK
        string name
        string description
        int64 principal_amount
        float64 interest_rate
        int repayment_count
        string repayment_frequency
        bool is_active
        time created_at
    }

    CUSTOMER ||--o{ LOAN : "has"
    LOAN }o--|| LOAN_PRODUCT : "belongs to"
    LOAN ||--o{ REPAYMENT_SCHEDULE : "contains"
    LOAN ||--o{ PAYMENT : "records"
```
---

## 📬 API Endpoints
| Method | Endpoint                 | Description                                                                    |
|--------|--------------------------|--------------------------------------------------------------------------------|
| POST   | `/customers`             | Create new customers                                                           |
| GET    | `/customers`             | Get all customer with filter                                                   |
| GET    | `/customers/:id`         | Get customer with full loan details and also outstanding and delinquent status |
| POST   | `/loan-products`         | Create new loan product                                                        |
| GET    | `/loan-products`         | Get all loan products with filter                                              |
| GET    | `/loan-products/:id`     | Get loan products by id                                                        |
| POST   | `/loans`                 | Create loan for specific user                                                  |
| POST   | `/loans/:id/pay`         | Pay loan                                                                       |
| GET    | `/loans/:id/outstanding` | Get outstanding amount for a loan                                              |
| POST   | `/loans/:id/delinquent`  | Get delinquent status of a loan                                                |
| GET    | `/loans/:id/outstanding` | Get total outstanding amount by customer                                       |
