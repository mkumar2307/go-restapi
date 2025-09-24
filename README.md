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

---

## ⚙️ Setup     

### 1️⃣ Clone the repository        

```bash
git clone https://github.com/mkumar2307/go-restapi.git
cd go-restapi
```

### 2️⃣ Install dependencies       

```go mod tidy```       

### 3️⃣ Set environment variables        

PostgreSQL example         

```bash
export DB_TYPE=postgres
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=secret
export DB_NAME=mydb
```        

MySQL example        

```bash
export DB_TYPE=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=secret
export DB_NAME=mydb
```        

SQLite example         

```bash
export DB_TYPE=sqlite
export DB_NAME=test.db
```      

## ▶️ Run the server       

```go run main.go```       

Server will start at:         
👉 http://localhost:8080        

## 📌 API Endpoints      

| Method | Endpoint       | Description        |
|--------|----------------|--------------------|
| GET    | /users         | Get all users      |
| GET    | /users/:id     | Get user by ID     |
| POST   | /users         | Create new user    |
| PUT    | /users/:id     | Update user by ID  |
| DELETE | /users/:id     | Delete user by ID  |


## 📖 Usage Examples      

### Create User        

```bash
curl -X POST http://localhost:8080/users/ \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice", "email": "alice@example.com"}'
```

### Get All Users         

```curl http://localhost:8080/users/```

### Get Single User       

```curl http://localhost:8080/users/1```

### Update User         

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Alice Smith", "email": "alice.smith@example.com"}'
```

### Delete User         
```curl -X DELETE http://localhost:8080/users/1```