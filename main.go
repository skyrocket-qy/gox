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
	// Create an instance of PBPerson
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
