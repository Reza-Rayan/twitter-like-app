# twitter-like-app
A minimalistic Twitter-like REST API built with Golang and Gin. This backend application allows users to register, log in, create posts, follow/unfollow other users, and view profiles with follower/following details. It’s designed as a learning project to explore building scalable and modular APIs in Go using best practices.

## 🔧 Technologies Used
- **Golang** – Core language powering the backend
- **Gin** – Lightweight and fast web framework for building RESTful APIs
- **Viper** – Configuration management (YAML-based)
- **SQLite3** – Simple and portable embedded SQL database
- **JWT** – Authentication with JSON Web Tokens
- **bcrypt** – Password hashing and validation

## 📦 Features

- User authentication (sign up, login)
- Create, read, update, delete (CRUD) posts
- Follow and unfollow users
- User profile with followers & following
- Secure password storage (bcrypt)
- Configurable via `config.yaml` using Viper
- Modular folder structure (routes, models, db, middlewares)

## 📁 Folder Structure
twitter-like-app/
├── config/ # Application configuration using Viper
│ └── config.go
├── db/ # Database initialization and schema
│ └── db.go
├── middlewares/ # Custom middleware (e.g., JWT auth)
│ └── auth.go
├── models/ # Database models and logic
│ ├── user.go
│ ├── post.go
│ └── follow.go
├── routes/ # API route handlers
│ └── routes.go
├── config.yaml # YAML-based application configuration
├── go.mod
├── go.sum
└── main.go # App entry point


## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/Reza-Rayan/twitter-like-app.git
   cd twitter-like-app