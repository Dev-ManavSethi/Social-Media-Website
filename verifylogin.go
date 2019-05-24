package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func verify(w http.ResponseWriter, r *http.Request) {

	MongoDBclient, error01 := mongo.Connect(context.Background(), "mongodb://localhost:27017")
	if error01 != nil {
		//handle the error
	}

	email := r.FormValue("email")
	passString := r.FormValue("password")
	//loggedin := r.FormValue("loggedin") == "on"

	result := MongoDBclient.Database("facebook").Collection("users").FindOne(context.Background(), bson.M{"email": email, "password": []byte(passString)})

	var user user

	result.Decode(&user)

	error02 := bcrypt.CompareHashAndPassword(user.Password, []byte(passString))

	if error02 != nil || email != user.Email {

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		fmt.Fprintf(w, "Invalid username or password!")
		//invalid password
	} else {

		// set cookie , name: facebook-session,  value : user.uuid

		http.SetCookie(w, &http.Cookie{
			Name:  "facebook-session",
			Value: user.UUID,
		})

		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}

}
