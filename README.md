## Description
ORM created using Go. It's convert prisma schema to migrations file.

## Usage
```bash
# Create migrations
go-orm run schema.prisma
# Apply migrations
go-orm apply
```

## Features
- [x] Convert Prisma schema to json file
- [x] Generate migration files
- [x] Apply migrations to database
- [ ] Rollback migrations
- [ ] Diff migrations
