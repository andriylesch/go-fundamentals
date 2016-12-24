package main

import (
	"fmt"
)

type UserRole struct{
    RoleName string    
}

type User struct{
    FirstName string
    LastName string
    Age byte
    UserRole // Anonymous field
}

func main() {

    var users = make([]User,0)

    /*
    
    WRONG initialization of variable

    user1 := User {FirstName: "FirstName1", LastName: "LastName1", Age: 30, RoleName: "Administrator"} // 
    
    */

    // first solution initialization
    user1 := User {
            FirstName: "FirstName1", 
            LastName: "LastName1", 
            Age: 30,
            UserRole : UserRole {RoleName : "Administrator"},
        }

    // second solution initialization
    var user2 User
    user2.FirstName = "FirstName2"
    user2.LastName = "LastName2"
    user2.Age = 25
    
    // equivalent set value
    user2.RoleName = "Manager"
    user2.UserRole.RoleName = "Manager"

    users = append(users, user1)
    users = append(users, user2)

    for _,item := range users {
        fmt.Println("-------------------")
        fmt.Println("User FullName : ", item.FirstName, item.LastName)
        fmt.Println("User Age : ", item.Age)
        fmt.Println("User Role : ", item.RoleName)
    }
}