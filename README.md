# OPWEB
A web server designed to be run alongside openpilot.

# Build Release
```bash
go build -ldflags="-extldflags=-static -s -w"
```
