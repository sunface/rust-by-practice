# Rust Authentication Server Using Warp + JWT

![Rust Logo](https://www.rust-lang.org/static/images/rust-logo-blk.svg)

Welcome to the Rust Authentication Server using Warp and JWT! This README will guide you through setting up and using this authentication server built with Rust.
Base on example guide on Youtube: https://www.youtube.com/watch?v=6oMoHZZeyb0

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Server](#running-the-server)
- [API Endpoints](#api-endpoints)


## Introduction

This project provides a secure authentication server written in Rust using the [Warp](https://github.com/seanmonstar/warp) web framework and JSON Web Tokens (JWT) for user authentication. It's designed to be easy to set up, highly performant, and secure. You can use it as a standalone authentication server or integrate it into your Rust web applications.

## Features

- User registration with password hashing.
- User login with JWT generation and validation.
- Token refresh mechanism for maintaining user sessions.
- Secure password reset functionality.
- CORS support for easy integration with front-end applications.

## Prerequisites

Before you get started, make sure you have the following prerequisites:

- Rust and Cargo: Install Rust and Cargo from [https://www.rust-lang.org/tools/install](https://www.rust-lang.org/tools/install).

## Getting Started

### Installation

1. Clone the repository to your local machine:


Change into the project directory:
```bash
cd rust-authentication-server
```
Build the project:
```bash
cargo build --release
```
Configuration
Configuration for the authentication server is managed through environment variables. Create a .env file in the project root or set these environment variables in your deployment environment.

Here are the required environment variables:

Running the Server
Once you have configured the environment variables, you can start the authentication server using the following command:

```bash
  cargo run --release
```
Your authentication server should now be running at http://localhost:3000

To use the authentication server, you can make HTTP requests to the provided API endpoints.

API Endpoints

POST /login: Authenticate a user and receive a JWT token. Requires a JSON body with email and pw fields. Returns a JWT token.
GET /user: Retrieve user information using a valid token. Requires a valid JWT token passed in the Authorization header. Returns user information.
GET /admin: Retrieve admin information using a valid token. Requires a valid JWT token passed in the Authorization header. Returns admin information.

Example
POST /admin:
```bash
  curl http://localhost:8000/login -d '{"email": "admin@adminaty.com", "pw": "4321"}' -H 'Content-Type: application/json'

  {"token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIyIiwicm9sZSI6IkFkbWluIiwiZXhwIjoxNjAzMTM0MjA1fQ.uYglVKRvb3h0bDC0Uz8FwGTu4v__Rl3toVI9fMI4_IT8keKde_SZRFQ4ii_PKzI4wjmDsZlnpULe6Tg0vWfEnw"}

  curl http://localhost:8000/admin -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIyIiwicm9sZSI6IkFkbWluIiwiZXhwIjoxNjAzMTM0MjA1fQ.uYglVKRvb3h0bDC0Uz8FwGTu4v__Rl3toVI9fMI4_IT8keKde_SZRFQ4ii_PKzI4wjmDsZlnpULe6Tg0vWfEnw' -H 'Content-Type: application/json'

  Hello Admin 2

  curl http://localhost:8000/user -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIyIiwicm9sZSI6IkFkbWluIiwiZXhwIjoxNjAzMTM0MjA1fQ.uYglVKRvb3h0bDC0Uz8FwGTu4v__Rl3toVI9fMI4_IT8keKde_SZRFQ4ii_PKzI4wjmDsZlnpULe6Tg0vWfEnw' -H 'Content-Type: application/json'

  Hello User 2
```
POST /user:
```bash
  curl http://localhost:8000/user -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwicm9sZSI6IlVzZXIiLCJleHAiOjE2MDMxMzQwODl9.dWnt5vfcGdwypEQUr3bLMrZYfdyxj3v6-io6VREWHXebMUCKBddf9xGcz4vHrCXruzx42zrS3Kygiqw3xV8W-A' -H 'Content-Type: application/json'

  Hello User 1

  curl http://localhost:8000/admin -H 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJzdWIiOiIxIiwicm9sZSI6IlVzZXIiLCJleHAiOjE2MDMxMzQwODl9.dWnt5vfcGdwypEQUr3bLMrZYfdyxj3v6-io6VREWHXebMUCKBddf9xGcz4vHrCXruzx42zrS3Kygiqw3xV8W-A' -H 'Content-Type: application/json'

  {"message":"no permission","status":"401 Unauthorized"}
```
