package ello

import (
	"database/sql"
	"log"
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
	db, err := sql.Open("postgres", "dbname=dfk5ej0olqe2rh host=ec2-54-243-192-203.compute-1.amazonaws.com port=5642 user=uf9tappjauoom password=p3jqdbbljvmoii510p374c4hu13 sslmode=require")
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
	err := DB.SQL(`SELECT id, username, email FROM users limit 10 `).QueryStructs(&users)
	if err != nil {
		return nil, err
	}
	log.Println(users)
	return users, nil
}

func UserForUsername(username string) (User, error) {
	var user User
	err := DB.SQL(`SELECT id, username, email FROM users WHERE username = $1 `, username).QueryStruct(&user)
	if err != nil {
		return User{}, err
	}
	log.Println(user)
	return user, nil
}
