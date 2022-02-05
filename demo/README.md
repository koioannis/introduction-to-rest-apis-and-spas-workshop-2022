# Installation

### 1. Install git
Install [git](https://git-scm.com/downloads). Verify installation:
```bash
$ git version
git version 2.25.1 // you should see something similiar
```

----
### 2. Clone the project and cd into the folder
```bash
$ git clone https://github.com/koioannis/introduction-to-rest-apis-and-spas-workshop-2022.git
$ cd introduction-to-rest-apis-and-spas-workshop-2022 
```

----
### 3.1 Docker way (recommented)
* Install [Docker](https://www.docker.com/get-started). Verify installation:
```bash
$ docker -v
Docker version 20.10.8, build 3967b7d // you should see something similiar
```
* Install [Docker Compose](https://docs.docker.com/compose/install/). Verify installation by running:

```bash
$ docker-compose -v
docker-compose version 1.29.1, build c34c88b2 // you should see something similiar
```
* Build and run the api and client

```bash
$ cd demo
$ docker-compose up
```
<i>Docker will download the required images, so if it takes a bit don't worry. </i>

----

### 3.2 Manual way
* Install [Go](https://go.dev/dl/). Verify installation:
```bash
$ go version
go version go1.16.5 linux/amd64 // you should see something similiar
```

* Install [Node](https://nodejs.org/en/). Verify installation:
```bash
$ npm --version 
8.1.3 // you should see something similiar
```

* Start the api server by running the follow commands:
```bash
$ cd demo/api
$ go run .
```
* Install dependencies and start the client server
```bash
$ cd demo/frontend
$ npm install
$ npm run serve
```

* API should be listening on [localhost:3000](`http://localhost:3000`)
* Client should be listening on [localhost:8080](`http://localhost:8080`)
