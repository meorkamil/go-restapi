# go-http


Just having fun during free time with go.


# Pre

- PostgreSQL server require, tested with PG 16
- Restore SQL file in `db/employee.sql`

# How to Build

```
make build
```

# How to run

```
make run
```

# Configuration

Sample configuration located in `config/config.yml`

```yaml
server:
  Addr: 0.0.0.0
  Port: 5001

database:
  Host: "192.168.0.40"
  User: "postgres"
  Pass: "123"
  DBName: "kamil"
  DBFlags: "?sslmode=disable"
```
