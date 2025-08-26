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

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/Reza-Rayan/twitter-like-app.git
   cd twitter-like-app

2. Install dependencies
   ```bash
   go mod tidy


3. Run Project 
   ```bash
   make run
