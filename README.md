## Packages

### Step 1 : Create the project
```bash
go mod init github.com/maintainerd/auth
```

### Step 2 : Install dependencies
```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get gorm.io/datatypes
go get github.com/joho/godotenv
go get github.com/google/uuid
go get -u github.com/golang-jwt/jwt/v5
```

### Step 3: Install CLI
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### Commands
```bash
# Create migration files
goose -dir db/migration/v1/core create create_tablename_table sql
goose -dir db/migration/v1/auth create create_tablename_table sql

# Run app
go run cmd/main.go
make run

# Run migrations
make migrate-up
make migrate-down

# Run seeder
make seed-local
```

