## Simpre JWT auth with Fiber

#### to run the project you need to run a postgres db and set the credentials y .env file

```bash
docker run --name postgres_db -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=super_db -p 5432:5432 -d postgres
mv .env.example .env
go run main.go
```

