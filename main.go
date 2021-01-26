package main

import ("fmt")
import ("net/http")
import ("html/template")

type User struct {
	name string 
	age uint16
	money int16
	avg_grades, happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. Age of %d. Her" + 
	" money: %d", u.name, u.age, u.money)
}

func (u *User) setNewName(newName string) {
	u.name = newName
} 

func home_page(w http.ResponseWriter, r *http.Request) {
	mari := User{"Mari", 20, 250, 3.8, 0.95 }
	// fmt.Fprintf(w, "<b>Main Text </b>")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, mari)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Our contacts www.")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	//var mari User = ...
	//mari := User{name: "Mari", age: 20, money: 250, avg_grades: 3.8, happiness: 0.95}
	handleRequest()
}