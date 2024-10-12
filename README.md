# go-http

Just having fun during free time with go.

# How to build or run

```
# Build with docker just use docker-compose
docker-compose up

# To build
make run
```

# Configuration

Sample configuration located in `config/config.yml`

```yaml
server:
  Addr: 0.0.0.0
  Port: 5001

database:
  Host: "db"
  Port: "5432"
  User: "postgres"
  Pass: "postgres"
  DBName: "kamil"
  DBFlags: "sslmode=disable"
  Type: "postgresql" # PostgreSQL will be default DB, support MySQL
```
