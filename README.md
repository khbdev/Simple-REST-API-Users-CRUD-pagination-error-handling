

# Mini Task: Go + Gin Users CRUD with Pagination

This project demonstrates a simple REST API for managing users with Go and Gin framework.
It includes **CRUD operations**, **cursor-based pagination**, and **standard error handling**.

---

## 📌 Features

* Create a new user (`POST /v1/users`)
* List users with **limit + after_id pagination** (`GET /v1/users`)
* Update a user (`PUT /v1/users/:id`)
* Delete a user (`DELETE /v1/users/:id`)
* Proper **error responses** for invalid input or missing resources

---

## 🛠 Technologies

* **Go**
* **Gin Web Framework** (`github.com/gin-gonic/gin`)
* In-memory storage (slice)
* Cursor-based pagination (`after_id` + `limit`)

---

## 🚀 API Endpoints

### 1️⃣ Create User

* **POST** `/v1/users`
* **Body JSON:**

```json
{
  "name": "Ali",
  "email": "ali@mail.com",
  "age": 20
}
```

* **Response:**

```json
{
  "id": 1,
  "name": "Ali",
  "email": "ali@mail.com",
  "age": 20
}
```

---

### 2️⃣ Get Users (Pagination)

* **GET** `/v1/users?limit=5&after_id=0`

* **Query Parameters:**

  * `limit` (optional, default 5) – number of users to return
  * `after_id` (optional, default 0) – fetch users **after this ID**

* **Response:**

```json
{
  "users": [
    {"id":1,"name":"Ali","email":"ali@mail.com","age":20},
    {"id":2,"name":"Bob","email":"bob@mail.com","age":21}
  ],
  "next_after": 2
}
```

---

### 3️⃣ Update User

* **PUT** `/v1/users/:id`
* **Body JSON:**

```json
{
  "name": "Ali Updated",
  "email": "ali2@mail.com",
  "age": 21
}
```

* **Response:**

```json
{
  "id":1,
  "name":"Ali Updated",
  "email":"ali2@mail.com",
  "age":21
}
```

* **Error if user not found:**

```json
{
  "error": "user not found"
}
```

---

### 4️⃣ Delete User

* **DELETE** `/v1/users/:id`
* **Response:**

```json
{
  "message": "user deleted"
}
```

* **Error if user not found:**

```json
{
  "error": "user not found"
}
```

---

## 📂 Project Structure

```
mini-task/
├── main.go           # Server entry point
├── models/
│   └── user.go       # User model
├── routes/
│   └── users.go      # CRUD handlers
└── storage/
    └── users.go      # In-memory storage + pagination logic
```



