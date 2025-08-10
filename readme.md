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

![img.png](img.png)

[Mermaid Js](https://mermaid.live/edit#pako:eNqlVNuO2jAQ_ZXIzwEFcgHytoJUlbosiMtDK6TIaw9gNbFTx6lKgX9fO2FZSGhZCT8lM2fsMzNnZo-IoIBCBHLE8EbidMUtfYbL-WIyjmbWvvo3p2BcWYxa028ftlxJxjcWxyk0jJBiljSs2VbwC6xiKVhEAlZAY6wqx3HFq4_nydPLPQqljRS5EinIWDu_1J2JwDzOpKCCFKqO0IDAszLNjbAMX_Bd6zDj0gCQkKtYao71OCUUTmKcioKrK58lIcO7FLiKybXzVIcP_1rCrwI42dWqkissVUyvXv1_uWbR9On7OHpZxPPh12i0fI4-VbyyPs2yWHnFq5F0Pd1XIRKL5XGGGa1RpQXcysAgz_Tv53XK6pFkbtA-NUKXWRV5U6en7qSgtoI-xt_IOJ7OJqPlcPHIRFHIiWSZYoL_U8ANMd7X8WNafe8-Jor9_rRUzwvmcGi1xL6a9NBaoS3OV-hi-o-i1TocrktocK-QCL7J9QRewU_X3ZgEE0QEV5jx_FbIu8gMTgIRkmoYspHeKnqRUb0jy86tkNqC7g0yOIrlT3PXUeOKzAg9okwJicI1TnKwES6UmO84ORsq1GnXnq1QRo2rVVxuZBtlmP8QQmOULKD8ReEe_UFhJ_Da_W63N3C9Tsfvej3PRjsUdoO2N-j5rh947sB3naB_tNHf8gan3Xdcr9_ru57jOG7gD2wkRbHZnglspMmwekkCpyCHRgn6Mb_nH98AA4HbPg)

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
