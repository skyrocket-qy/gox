package main

import (
	"fmt"
	"reflect"
)

type PBPerson struct {
	Name string
	Age  int32
}

type CustomPerson struct {
	Name string
	Age  int32
}

// Function to map values from PBPerson to CustomPerson
func mapPBToCustom(pb interface{}, custom interface{}) error {
	pbVal := reflect.ValueOf(pb).Elem()
	customVal := reflect.ValueOf(custom).Elem()

	for i := 0; i < pbVal.NumField(); i++ {
		pbField := pbVal.Type().Field(i)
		customField := customVal.FieldByName(pbField.Name)

		// Check if the field exists in the custom struct and the types are assignable
		if customField.IsValid() && customField.CanSet() &&
			customField.Type() == pbField.Type {
			customField.Set(pbVal.Field(i))
		}
	}

	return nil
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
	pbPerson := &PBPerson{Name: "Alice", Age: 30}

	// Create an instance of CustomPerson
	customPerson := &CustomPerson{}

	// Map values from PBPerson to CustomPerson
	err := mapPBToCustom(pbPerson, customPerson)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print the mapped custom struct
	fmt.Printf("%+v\n", customPerson) // Output: &{Name:Alice Age:30}
}
