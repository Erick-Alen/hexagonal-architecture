## Hexagonal Architecture/Ports and Adapters
This repository was intended to exemplify an implmentation of an hexagonal architecture, compounded by ports and adapters.
Therefore, here's some commands you need to know to run this aplication locally.

### Running into a Docker container
For this, you need to install docker in your environment and then, run the following commands:
```bash
  docker compose build
  docker compose up -d
```
After this, you need to get the id of your container to execute in the integrated terminal with these commands:
```bash
  docker container ls
  docker exec -it <container-id> /bin/bash
```
### Updating go libraries:
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

### Inserting data through sqliteDB
```bash
touch db.sqlite
sqlite3 db.sqlite
```
Then, just execute your sql commands
```sql
create table products(id string, name string, price float, status string);
```


<!-- ### initializing cobra-cli
Check if there are a cobra-init module into your `go/bin` folder, and then run the following command to initialize the cli into the project:
```bash
cobra-cli init
```
Now you can use the command `cobra-cli -h` to explore the
```bash
cobra-cli init
``` -->

### using `go run main.go` command with cobra-cli
Here you can use the `go run main.go` command and choose which command you want to run (the commands available recognized are the files inside the `cmd folder`). The root is the default command.
To know more details about the commands available and the command's details:
```bash
go run main.go -h
```
```bash
go run main.go <your-command> -h
```
Now you can use the command `go run main.go` to test the application or run the command as an API server as well
```bash
go run main.go http
```
