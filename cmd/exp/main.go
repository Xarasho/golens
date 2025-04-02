package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
	// Meta UserMeta
}

// type UserMeta struct {
// 	Visits int
// }

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// Anonymous Struct
	// user := struct {
	// 	Name string
	// 	Age  int
	// }{Name: "Alex Rojas", Age: 111}

	user := User{
		Name: "Alex Rojas",
		Bio:  `<script>alert("Haha, you've been h4x0r3d!");</script>`,
		Age:  123,
		// Meta: UserMeta{
		// 	Visits: 4,
		// },
	}

	// This is how you process the template
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
