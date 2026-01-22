## Description

ORM created using Go. It's convert prisma schema to migrations file.

## Usage

```bash
# Create migrations
go-orm run schema.prisma

# Apply migrations
# Set DATABASE_URL environment variable
export DATABASE_URL="postgres://postgres:root@localhost:5432/testdemo?sslmode=disable"
# Set DATABASE_URL environment variable (Windows)
$env:DATABASE_URL="postgres://postgres:root@localhost:5432/testdemo?sslmode=disable"

# Run apply command
go-orm apply
```

## Features

- [x] Convert Prisma schema to json file
- [x] Generate migration files
- [x] Apply migrations to database (prisma schema order have to maintain before generate)
- [ ] Rollback migrations
- [ ] Diff migrations
