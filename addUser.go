package main

import (
	"context"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	MongoDBclient, error0 := mongo.Connect(context.Background(), "mongodb://localhost:27017")
	if error0 != nil {
		//handle error
	}

	error01 := r.ParseForm()
	if error01 != nil {
		//handle error
	}

	fname := r.FormValue("firstname")
	lname := r.FormValue("lastname")
	age := r.FormValue("age")
	ageInt, _ := strconv.Atoi(age)
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	gender := r.FormValue("gender")
	uuid, error01 := uuid.NewV4()
	if error01 != nil {
		//handle error
	}
	uuidString := uuid.String()

	passEnc, error2 := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if error2 != nil {
		//handle error
	}

	LastLogin := time.Now().Unix()

	var user user
	user.LastLogin = LastLogin
	user.Password = passEnc
	user.FirstName = fname
	user.LastName = lname
	user.Email = email
	user.Gender = gender
	user.UUID = uuidString
	user.Age = ageInt

	//add user to mongo
	_, error02 := MongoDBclient.Database("facebook").Collection("users").InsertOne(context.TODO(), user)
	if error02 != nil {
		//handle error
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "facebook-session",
		Value: uuidString,
	})

	tpl, error03 := template.ParseFiles("templates/completeProfile.gohtml")
	if error03 != nil {
		//handle error
	}

	error4 := tpl.Execute(w, nil)
	if error4 != nil {
		//handle error
	}

}
