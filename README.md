# Snippetbox

Snippet sharing web application built with Go and PostgreSQL.  
Based on "Let's Go" by Alex Edwards.

## Requirements

- Go 1.22+
- PostgreSQL 15+
- golang-migrate CLI: `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
- make

## 1. Clone the project

```bash
git clone https://github.com/shodruzhosimzoda/Snippetbox.git
cd Snippetbox
```

## 2. Set up the Database
```sql
    CREATE DATABASE snippetbox;
        CREATE USER snippets_user WITH PASSWORD 'your_password_here';
        GRANT ALL PRIVILEGES ON DATABASE snippetbox TO snippets_user;
        \c snippetbox
        GRANT ALL ON SCHEMA public TO snippets_user;
        GRANT ALL ON ALL TABLES IN SCHEMA public TO snippets_user;
        GRANT ALL ON ALL SEQUENCES IN SCHEMA public TO snippets_user;
```
## Configure your .env file 
``` 
   MIGRATIONS_PATH=./db/migrations
   DB_URL=postgres://user:password@host:port/db?sslmode=disable
 
```

## 5. Apply Migrations

``` 
    make migrate-up
```

## 6. Run server
### You can run server via Makefile or via commandline:
``` 
    go run ./cmd/web
```