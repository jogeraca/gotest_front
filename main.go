package main

import (
	"encoding/json"
	"fmt"
	"github.com/bmizerany/pat"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type Users struct {
	Username  string
	Password  string
	Email     string
	Telephone string
	Country   string
	City      string
	Name      string
	Address   string
}

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(index))
	mux.Post("/", http.HandlerFunc(send))
	//mux.Get("/confirmation", http.HandlerFunc(confirmation))

	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	// Validate form
	// Send message in an email
	// Redirect to confirmation page
	msg := &Message{
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
		Password:  r.FormValue("password"),
		Telephone: r.FormValue("telephone"),
		Country:   r.FormValue("country"),
		City:      r.FormValue("city"),
		Name:      r.FormValue("name"),
		Address:   r.FormValue("address"),
	}

	if msg.Validate() == false {
		render(w, "templates/index.html", msg)
		return
	}
	hash, _ := HashPassword(msg.Password)
	user := new(Users)
	user.Username = msg.Username
	user.Password = hash
	user.Email = msg.Email
	user.Name = msg.Name
	user.Telephone = msg.Telephone
	user.Country = msg.Country
	user.City = msg.City
	user.Address = msg.Address
	msg_json, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := prepare_message(string(msg_json))
	if err != nil {
		panic(err)
	}
	//	result := Result{Response: res}

	//fmt.Println(result.Response)
	///  			_, err := xhr.Send("POST", "/foo", string(fooJson))
	success := Message{
		Result: res,
	}
	render(w, "templates/index.html", success)
	//http.Redirect(w, r, "/", http.StatusSeeOther)

}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//func CheckPasswordHash(password, hash string) bool {
//err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
//return err == nil
//}
