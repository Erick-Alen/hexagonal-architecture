Updating go libraries:
```bash
  docker compose build
  docker compose up -d
  docker exec -it <container-id> /bin/bash
```
installing mockgen:
```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

Updating go libraries:
```bash
go mod tidy
```

```bash
mockgen -source=application/product.go -destination=application/mocks/product_mock.go application

```
