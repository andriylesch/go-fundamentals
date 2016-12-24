package main

import "fmt"

type User struct {
	FirstName string
	LastName  string
}

type Users []User

// void method which Print full username
func UserPrint(user User) {
	fmt.Println("UserPrint method : ", user.FirstName, user.LastName)
}

// function with return result
func BuildUserFullName(user User) string {
	return user.FirstName + " " + user.LastName
}

// function with multiple returns
func GetUserData(user User) (fname, lname string) {
	fname, lname = user.FirstName, user.LastName
	return
}

// variadic function with list of params
func VariadicFunc(args ...string) string {

	var res string
	for _, value := range args {
		res += " " + value
	}

	return res
}

// function with passing function as parameter
func Execute(a, b int, do func(int, int) int) {

	res := do(a, b)
	fmt.Println("Execute method : result = ", res)
}

// function with function type as parameter
func ExecuteFunc(a, b int, do DoSum) {

	res := do(a, b)
	fmt.Println("ExecuteFunc method : result = ", res)
}

func Calc(a, b int) int {
	return a + b
}

//Declare function type - for .Net developers it's like delegate
type DoSum func(int, int) int

//Extentions method
func (users Users) PrintUsersInfo(isFormat bool) {

	fmt.Println("PrintUsersInfo :")

	for _, user := range users {
		if isFormat {
			fmt.Println("-------------")
			fmt.Println("User info ")
			fmt.Println("FirstName : ", user.FirstName)
			fmt.Println("LastName : ", user.LastName)
			fmt.Println("-------------")
		} else {
			fmt.Println(user.FirstName + " _ " + user.LastName)
		}
	}
}

func (user *User) Rename(name string) {
	user.FirstName = name
}

func main() {
	user := User{FirstName: "Andriy", LastName: "Lesch"}

	UserPrint(user)

	fullname := BuildUserFullName(user)
	fmt.Println("BuildUserFullName method : ", fullname)

	fname, lname := GetUserData(user)
	fmt.Println("GetUserData method : FirstName = ", fname, " LastName =", lname)

	variadicRes := VariadicFunc("param1", "param2", "param3", "param4", "param5")
	fmt.Println("VariadicFunc method : result = ", variadicRes)

	Execute(15, 14, Calc) // using exists function 'Calc'

	Execute(15, 20, func(a, b int) int { return a + b }) // using anonymous function

	ExecuteFunc(15, 20, Calc) // using delegate as func parameter

	users := Users{
		{"FN1", "LN1"},
		User{"FN2", "LN2"},
		User{FirstName: "FN3", LastName: "LN3"},
	}

	users.PrintUsersInfo(true)

	users[0].Rename("firstName")
	users.PrintUsersInfo(false)

}
