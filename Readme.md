## Web Service in GO

## Steps to run

1.  First start the postgres service

```bash
make dev
```

2. Create `.env` file based on `.env-example` file.

3. To run the project build the project using following command

```bash
go build
```

4. Now run the executable created

```bash
./lets-go
```

5. To stop the postgres container from running.

```bash
make stop
```

#### Without database connection : [See Here](https://github.com/0xmatriksh/lets-go/tree/88ef82436ce614de6d2c107c58612d4d99dcd512)
