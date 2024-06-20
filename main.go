package main

import (
	"fmt"

	"github.com/skyrocketOoO/GoUtils/StructMapper"
)

type PBPerson struct {
	Name string
	Age  int64
	// Option string
	Number string
}

type BankCard struct {
	Number string
}

type CustomPerson struct {
	Name   string
	Age    string
	Option string
	BankCard
}

func main() {
	// _ := &pb.Person{
	// 	Name:              "Alice",
	// 	Age:               30,
	// 	Height:            170,
	// 	Weight:            65,
	// 	NetWorth:          1000000,
	// 	Temperature:       37,
	// 	Distance:          10000,
	// 	FixedSalary:       50000,
	// 	FixedAssets:       200000,
	// 	SFixedBonus:       10000,
	// 	SFixedLiabilities: 5000,
	// 	IsEmployed:        true,
	// 	Gpa:               3.75,
	// 	Accuracy:          0.99,
	// 	ProfilePicture:    []byte("picture_data"),
	// 	Hobbies:           []string{"reading", "hiking"},
	// 	Contacts: map[string]string{
	// 		"email": "alice@example.com",
	// 		"phone": "123-456-7890",
	// 	},
	// 	EmploymentStatus: &pb.Person_JobTitle{JobTitle: "Software Engineer"},
	// }
	// Create an in
	pbPerson := &PBPerson{Name: "Alice", Age: 30, Number: "eafawef"}
	// Create an instance of CustomPerson
	customPerson := &CustomPerson{}

	fmt.Printf("%+v\n", customPerson) // Output: &{Name:Alice Age:30}
	// Map values from PBPerson to CustomPerson
	err := StructMapper.MapStructToStruct(pbPerson, customPerson)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print the mapped custom struct
	fmt.Printf("%+v\n", customPerson) // Output: &{Name:Alice Age:30}

	// StructMapper.PrintStructInfo(PBPerson{})
}
