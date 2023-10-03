package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator/v10"
)

var v = validator.New()

type Dog struct {
	Name      string `json:"name" validate:"required"`
	Birthday  string `json:"birthday" validate:"numeric,len=10"`
	isGoodBoy bool   `json:"is_good_boy" validate:"required,istrue"`
}

func validateIsGoodBoy(fl validator.FieldLevel) bool {
	//fmt.Println(fl.Field().String())
	//return fl.Field().Bool()
	return false
}

func main() {
	body, _ := ioutil.ReadFile("dog_invalid.json")
	var pet Dog

	err := json.Unmarshal(body, &pet)
	if err != nil {
		fmt.Printf("Could not unmarshal json: %s", err)
		os.Exit(1)
	}

	v.RegisterValidation("istrue", validateIsGoodBoy)
	err = v.Struct(&pet)
	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Printf("Error validating field %s. Expected %s to equal %s, got value %s\n", err.Field(), err.Tag(), err.Param(), err.Value())
		}
	} else {
		fmt.Println("Validation succeeded")
	}

}
