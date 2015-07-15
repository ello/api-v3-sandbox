# GO API

## Setup

I'd follow some variant of http://www.wolfe.id.au/2015/03/05/using-sublime-text-for-go-development/

Once go is setup:

`
go get
go build
./api-go
`

* localhost:8080/users will get the first 10 in the staging db
* localhost:8080/user/{username} will return that user if they exist
