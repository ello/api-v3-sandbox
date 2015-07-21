# GO API

## Setup

I'd follow some variant of http://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/

Once go is setup:

```
go get
go build
./api-go
```

* localhost:8080/users will get the first 10 in the staging db
* localhost:8080/user/{username} will return that user if they exist

## Docker style

This assume you have boot2docker setup and running:

```
docker build -t <appname> .
docker run -t -i -p 8080:8080 <appname>
```

This will have the server running interactively (ctrl-c will kill it).

In another terminal:

```
curl `boot2docker ip`:8080/users
```
