# Formus

A high-performance, self-hosted backend service for handling HTML form submissions. Built with Go and SQLite.

![Go Version](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-black?style=flat-square)

---

## Overview

Formus provides a centralized endpoint for static websites to collect and manage user submissions without third-party dependencies. It is designed for developers who prioritize data privacy and system transparency.

## Core Features

- **Dynamic Schema**: Accepts any JSON payload without prior configuration.
- **Persistent Storage**: Uses SQLite with optimized WAL mode for concurrent writes.
- **SSR Dashboard**: A lightweight administrative interface for real-time submission monitoring.
- **Security**: Isolated execution environment via Docker with environment-based authentication.

## Architecture

The system follows a minimalist design to ensure low memory footprint and high throughput.

- **Language**: Go 1.26
- **Database**: SQLite (CGO-free driver)
- **Delivery**: Multi-stage Docker build (Alpine Linux)
- **Rendering**: Standard Go `html/template` engine (Zero JS overhead)

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repository:

```bash
git clone https://github.com/whitekotan1/formus
```

2. Navigate into the project directory:

```bash
cd formus
```

3. Create a `.env` file with your password:

```
FORMUS_PASSWORD=your_secure_password
```

4. Build and start the service:

```bash
docker compose up -d --build
```

---

## Usage

Once running, Formus exposes:

- A submission endpoint for HTML forms
- An SSR dashboard for monitoring submissions

Example form:

```html
<form action="http://localhost:5000/submit" method="POST">
  <input type="text" name="email" />
  <textarea name="message"></textarea>
  <button type="submit">Send</button>
</form>
```

---

