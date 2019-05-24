package main

import (
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {

	//check cookie session
	_, error0 := r.Cookie("facebook-session")

	if error0 == http.ErrNoCookie {

		tpl, error1 := template.ParseFiles("templates/login.gohtml")
		if error1 != nil {

		}

		error2 := tpl.Execute(w, nil)
		if error2 != nil {

		}

	} else if error0 != http.ErrNoCookie && error0 != nil {

		tpl, error1 := template.ParseFiles("templates/login.gohtml")
		if error1 != nil {

		}

		error2 := tpl.Execute(w, nil)
		if error2 != nil {

		}

	} else {

		http.Redirect(w, r, "/home", http.StatusSeeOther)

	}

}
