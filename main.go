package main

// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details

type User struct {
	Val int
	Info
}

type Info struct {
	Number string
}

func main() {
	user := User{}
	user.Number = "abc"
}
