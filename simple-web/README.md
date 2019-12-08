# Forum Mirror

SYSU 2019 Service Computing [Homework](https://pmlpml.github.io/ServiceComputingOnCloud/ex-services)

A simple web application that craws and displays a part of [Emacs China](https://emacs-china.org)

## Build & Run

### Docker

#### Prerequisites

+ docker (19.03.4)
+ docker-compose (1.25.0)

Run the following command and visit `localhost:3000` on browser:

```bash
docker-compose up -d
```

### Without Docker

#### Prerequisites

+ node (10.16.0)
+ yarn (1.19.1)
+ go (1.13)
+ postgres (10)

Backend's  configuration is `config/config.go`

Default connection to postgres will assume:

+ host: localhost:5432 
+ username: postgres
+ password: postgres

#### Run backend

Let server listen on `localhost:8080`

```bash
go run main.go
```

#### Run frontend

Launch React app on `localhost:3000` 

```bash
yarn
yarn start
```

The you can visit the application on browser
