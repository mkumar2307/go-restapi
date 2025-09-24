# Go REST API with Gin and GORM    

A simple **REST API in Go** using the [Gin framework](https://github.com/gin-gonic/gin) and [GORM](https://gorm.io/) ORM.          
This project demonstrates **CRUD operations** on a `User` model, with support for multiple SQL databases (Postgres, MySQL, SQLite, SQL Server).

---

## 🚀 Features
- ✅ Clean project structure (`config/`, `models/`, `controllers/`, `routes/`)
- ✅ Database connection with **GORM**
- ✅ Supports **Postgres, MySQL, SQLite, SQL Server**
- ✅ Full **CRUD API** for `User`
- ✅ Auto migration of database schema
- ✅ JSON-based REST responses

---

## Project Structure    

go-restapi/
│── go.mod
│── go.sum
│── main.go
│
├── config/
│ └── sqldb.go # Database connection logic
│
├── models/
│ └── user.go # User model
│
├── controllers/
│ └── user_controller.go # CRUD logic
│
└── routes/
└── user_routes.go # API route definitions         


