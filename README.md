# ServerPulse

ServerPulse is a lightweight, local-first server monitoring agent written in Go.

It is designed to be simple, low-overhead, and easy to extend. The initial focus is on collecting basic system metrics such as CPU, memory, and disk usage, with a clean architecture suitable for future expansion (server ingestion, alerts, dashboards).

---

## ✨ Features

- 📊 Collects system metrics
  - CPU usage
  - Memory usage
  - Disk usage
- ⏱ Configurable collection interval
- 🧵 Graceful shutdown using context and OS signals
- 🧪 Testable, dependency-injected design
- 🧹 Strict linting with `golangci-lint`
- ⚙️ Makefile-driven development & CI
- 🖥 Runs locally with no external dependencies

---

## 📁 Project Structure

```
serverpulse/
├── cmd/
│   ├── agent/          # Monitoring agent binary
│   ├── server/         # (future) backend server
│   └── cli/            # (future) CLI tool
├── internal/
│   └── metrics/        # System metrics collectors
├── .github/
│   └── workflows/      # CI/CD workflows
├── build/              # Build artifacts (ignored)
├── Makefile
├── go.mod
├── go.sum
├── .golangci.yaml      # Linter configuration
└── README.md
```

Each folder under `cmd/` represents a standalone binary.

---

## 🚀 Getting Started

### Prerequisites

- Go 1.25+
- Linux or macOS (Linux recommended for production)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/RaunakPrakash/serverpulse.git
   cd serverpulse
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

### Building

Build the agent binary:
```bash
make build
```

This creates binaries in the `build/` folder (e.g., `build/agent`).

### Running

Run the agent locally with a custom interval:
```bash
go run ./cmd/agent -interval=5
```

Or using the built binary:
```bash
./build/agent -interval=5
```

Available flags:
- `-interval`: Metrics collection interval in seconds (default: 60)
- `-endpoint`: Server endpoint URL (for future server integration)
- `-apikey`: API key for authentication (for future server integration)
- `-debug`: Enable debug logging
- `-disk-path`: Disk path to monitor (default: "/")

### Testing

Run unit tests:
```bash
make test
```

Run tests with race detection:
```bash
make race
```

### Linting

Run the linter:
```bash
make lint
```

Format code:
```bash
golangci-lint fmt
```

### CI

The project uses GitHub Actions for CI. On every push and pull request to `main`, it runs:
- Linting
- Tests
- Race detection
- Coverage
- Build

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature`
3. Make your changes
4. Run tests and linting: `make ci`
5. Commit your changes: `git commit -am 'Add some feature'`
6. Push to the branch: `git push origin feature/your-feature`
7. Submit a pull request

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🔮 Future Plans

- Backend server for metrics ingestion
- Alerting system
- Web dashboard
- CLI tool for configuration
- Support for more metrics (network, processes, etc.)
- Docker containerization
