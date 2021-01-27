package main

import ("fmt")
import ("net/http")
import ("html/template")

type User struct {
	Name string 
	Age uint16
	Money int16
	Avg_grades, Happiness float64
	Hobbies []string 
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s. Age of %d. Her" + 
	" money: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
} 

func home_page(w http.ResponseWriter, r *http.Request) {
	mari := User{"Mari", 20, 25, 3.8, 0.9, []string{"Languages", "Books", "Travels"}}
	// fmt.Fprintf(w, "<b>Main Text </b>")
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, mari)
}