# Sony Level Up

## Software Engineer Technical Challenge

## Overview

This repository contains my solution to the Sony Playstation Software Engineer Technical Challenge. The goal was to build a system that fetches user data from a provided API, processes the data, and determines an account-wide achievement level based on game ownership and achievements.

The solution includes both a backend API service and a frontend UI that together provide a complete, user-friendly way to view each user’s achievement level.

## Solution Summary

_Backend_:

- Implemented using Go.
- Fetches user, game library, and achievement data from the provided Sony Playstation mock server.
- Processes the data to calculate achievement levels (Bronze, Silver, Gold, Platinum) based on specified criteria.
- Exposes RESTful endpoints to serve processed user achievement levels

_Frontend_:

- Built with React and TypeScript
- Fetches processed data from the backend API
- Displays a user-friendly interface listing users and their achievement levels
- Responsive design and simple UI for clear presentation of data

---

## Achievement Level Criteria

| Achievement Level | Criteria                                         |
| ----------------- | ------------------------------------------------ |
| Bronze            | Owns more than 10 games                          |
| Silver            | Owns 10+ games and has 75%+ achievements in each |
| Gold              | Owns 25+ games and has 80%+ achievements in each |
| Platinum          | Owns 50+ games and has 100% achievements in each |

---

## Getting Started

### Prerequisites

- Java 17+ (or your backend runtime environment)
- Node.js and npm (for frontend)
- Go

### Running the API Server (Sony Playstation API)

1. Extract the provided `.zip` containing `playstation-tech-test-api.jar`
2. Run the API server:
   ```bash
   java -jar playstation-tech-test-api.jar
   ```
3. API will be accessible at:
   ```text
   http://localhost:8080
   ```
4. OpenAPI documentation available at:
   ```text
   http://localhost:8080/swagger-ui/index.html
   ```

### Running the Backend

1. Run the backend go server at the root of the project:
   ```bash
   go run .
   ```
2. Backend will be available at:
   ```text
   http://localhost:5000
   ```

### Running the Frontend

1. Install dependecies by navigating to the ui directory and running:
   ```bash
   cd ui
   npm install
   ```
2. Run the frontend:
   ```bash
   npm run dev
   ```
3. Backend will be available at:
   ```text
   http://localhost:5173
   ```

### Usage

1. Open the frontend URL in your browser at:
   ```text
   http://localhost:5173
   ```
2. The app fetches all users from the backend.
3. View each user’s achievement level by click on their profile card.
