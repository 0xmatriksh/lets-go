<div align="center">

## Web Service in GO

<img src="images/go.png" height="50">
<img src="https://img.shields.io/github/go-mod/go-version/Dreamacro/clash?style=flat-square">
&emsp;
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/Postgresql_elephant.svg/1985px-Postgresql_elephant.svg.png" height="50">

</div>

## Steps to run

1. Create `.env` file based on `env.example` file.

2. First, start the Postgres service with following command.

```
make dev
```

3. To run the project build the project using following command

```
go build
```

4. Now, run the executable created with following command

```
./lets-go
```

5. To stop the postgres container from running.

```
make stop
```

#### Without database connection : [See Here](https://github.com/0xmatriksh/lets-go/tree/88ef82436ce614de6d2c107c58612d4d99dcd512)
