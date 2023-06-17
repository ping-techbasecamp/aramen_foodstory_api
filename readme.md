### Install Guide
1. Set GOPATH
```bash
export GOPATH=<path to project>
```
2. Setup
```bash
go mod tidy
go generate ./...
```
3. Run go
```bash
go run cmd/api/main.go
```

### Testing
```bash
go test ./...
```
