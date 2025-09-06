# 🖥️ Simple System Monitor (Go)

A beginner-friendly Go project that monitors **CPU** and **Memory usage**  
and exposes them as JSON via an HTTP endpoint.  

This project is designed as a **first step into Go** and a **starter SRE tool**.  

---

## 🚀 Features
- Collects CPU usage %
- Collects Memory usage %
- Simple HTTP server exposing `/metrics` endpoint
- JSON output for easy integration with other tools

---

## 📂 Project Structure
```
simple-sys-monitor/
│── main.go       # Main program
│── go.mod        # Go module file
│── README.md     # Documentation
```

---

## 🛠️ Getting Started

### 1. Install Go
Download Go (>= 1.19) from [go.dev](https://go.dev/dl/).

### 2. Clone the Repository
```bash
git clone https://github.com/ahmed20amir23/simple-sys-monitor.git
cd simple-sys-monitor
```

### 3. Install Dependencies
```bash
go mod tidy
```

### 4. Run the Server
```bash
go run main.go
```

### 5. Check Metrics
Open [http://localhost:8080/metrics](http://localhost:8080/metrics)  

Example output:
```json
{
  "cpu_usage_percent": 12.5,
  "memory_usage_percent": 43.8
}
```

---

## 📌 Next Steps (Future Improvements)
- Add **alerts** (e.g., warning if CPU > 80%)
- Store metrics in **SQLite/Prometheus**
- Build a simple **web dashboard**
- Send alerts via **Slack/Email**

---

## 👨‍💻 Author
Made with ❤️ by [ENG: (Ahmed)](https://github.com/ahmed20amir23)
