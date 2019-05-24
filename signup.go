package main

import (
	"net/http"
	"text/template"
)

func signup(w http.ResponseWriter, r *http.Request) {

	//check cookie session
	_, error0 := r.Cookie("facebook-session")

	if error0 == http.ErrNoCookie {

		tpl, error1 := template.ParseFiles("templates/signup.gohtml")
		if error1 != nil {

		}

		error2 := tpl.Execute(w, nil)
		if error2 != nil {

		}

	} else if error0 != http.ErrNoCookie && error0 != nil {

		tpl, error1 := template.ParseFiles("templates/signup.gohtml")
		if error1 != nil {

		}
		//generate cookie

		error2 := tpl.Execute(w, nil)
		if error2 != nil {

		}

	} else {

		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}

}
