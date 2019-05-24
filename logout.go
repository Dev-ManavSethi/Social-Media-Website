package main

import "net/http"

func logout(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "facebook-session",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
