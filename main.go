package main

import (
	"context"
	"log"
	"net/http"
	"text/template"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/mongodb/mongo-go-driver/mongo"
)

type user struct {
	FirstName string
	LastName  string
	Email     string
	Password  []byte
	Age       int
	Gender    string

	LastLogin   int64
	UUID        string
	FriendsList []string
	LikedPages  []string
}

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/verify", verify)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/AddUser", AddUser)
	http.HandleFunc("/completeProfile", completeProfile)
	http.HandleFunc("/home", home)
	// http.HandleFunc("/profile", profile)

	http.HandleFunc("/logout", logout)

	log.Println("Server running")
	error1 := http.ListenAndServe(":8090", nil)
	if error1 != nil {
		for error1 != nil {
			error1 = http.ListenAndServe(":8090", nil)

		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	//check cookie session
	_, error0 := r.Cookie("facebook-session")

	if error0 == http.ErrNoCookie {

		tpl, error1 := template.ParseFiles("templates/index.gohtml")
		if error1 != nil {

		}
		error2 := tpl.Execute(w, nil)
		if error2 != nil {

		}

	} else if error0 != http.ErrNoCookie && error0 != nil {

		tpl, error1 := template.ParseFiles("templates/index.gohtml")
		if error1 != nil {

		}

		error2 := tpl.Execute(w, nil)
		if error2 != nil {

		}

	} else {

		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}

}

func home(w http.ResponseWriter, r *http.Request) {

	MongoDBclient, error0 := mongo.Connect(context.Background(), "mongodb://localhost:27017")

	if error0 != nil {
		//handle error
	}

	cookie, error01 := r.Cookie("facebook-session")
	if error01 == http.ErrNoCookie {

		tpl, error2 := template.ParseFiles("templates/login.gohtml")
		if error2 != nil {
			//handle error
		}
		error03 := tpl.Execute(w, nil)

		if error03 != nil {
			//handle error
		}

	}

	uuid := cookie.Value

	//get user from uuid from db

	result := MongoDBclient.Database("facebook").Collection("users").FindOne(context.TODO(), bson.M{"uuid": uuid})

	var user user
	error02 := result.Decode(&user)
	if error02 != nil {
		//handle error
	}

	tpl, error1 := template.ParseFiles("templates/home.gohtml")
	if error1 != nil {

	}

	error2 := tpl.Execute(w, user)
	if error2 != nil {

	}

}
