package ello

import (
	"database/sql"
	"time"

	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Database connection
var DB *runner.DB

func init() {
	db, err := sql.Open("postgres", "dbname=dd8a20nnsdelm7 host=ec2-54-197-226-173.compute-1.amazonaws.com port=5432 user=uc7s3bo18tu0ap password=p8umei0b8m2mcod6oiemcegj9g1 sslmode=require")
	if err != nil {
		panic(err)
	}

	runner.MustPing(db)

	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	dat.EnableInterpolation = true

	runner.LogQueriesThreshold = 10 * time.Millisecond

	DB = runner.NewDB(db, "postgres")
}

func Users() ([]*User, error) {

	var users []*User
	err := DB.SQL(`SELECT id, username, email FROM users limit 100000`).QueryStructs(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UserForUsername(username string) (User, error) {
	var user User
	err := DB.SQL(`SELECT id, username, email FROM users WHERE username = $1 `, username).QueryStruct(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
