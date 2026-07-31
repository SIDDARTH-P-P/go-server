# Go CRUD Server

This project is a basic Go REST API with:

- CRUD operations for an in-memory item store
- Basic authentication
- Clean Go folder structure

## Project Structure

```text
cmd/
  server/
    main.go
internal/
  api/
    handlers.go
  auth/
    auth.go
  store/
    store.go
    store_test.go
```

## Run the Server

From the project root:

```bash
cd /home/Downloads/go-server
go run ./cmd/server
```

The server starts on:

```text
http://localhost:8080
```

## Authentication

All routes require Basic Auth.

- Username: `admin`
- Password: `password123`

## API Endpoints

### List items

```bash
curl -u admin:password123 http://localhost:8080/items
```

### Create item

```bash
curl -u admin:password123 -X POST http://localhost:8080/items \
  -H "Content-Type: application/json" \
  -d '{"title":"Demo Item"}'
```

### Get item

```bash
curl -u admin:password123 http://localhost:8080/items/1
```

### Update item

```bash
curl -u admin:password123 -X PUT http://localhost:8080/items/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Item"}'
```

### Delete item

```bash
curl -u admin:password123 -X DELETE http://localhost:8080/items/1
```

## Notes

- The data store is in-memory, so items are reset when the server restarts.
- The package entrypoint is `cmd/server`, not a root-level `main.go`.
