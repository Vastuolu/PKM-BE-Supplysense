package helper

import (
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

func MapValidationErr(err error)map[string]string{
	var errMap = make(map[string]string)
	validationErrors, ok := err.(validator.ValidationErrors); 
	if !ok {
		log.Fatalf("error: MapValidationErr is Error: %v", err)
		return nil
	}
	for _,err := range validationErrors{
		var errMessage string
		fieldName := err.Field()
		tag := err.Tag()
		
		switch tag{
		case "required":
			errMessage = fieldName+" is Required"
		case "required_if":
			errMessage = fieldName+" is Required"
		case "email":
			errMessage = fieldName+" is not a valid Email"
		case "max":
			errMessage = fieldName+" is too much, max "+ err.Param()
		case "url":
			errMessage = fieldName+" is not a valid url"
		default:
			errMessage = fieldName+" is not valid, "
		}

		errMap[strings.ToLower(fieldName)] = errMessage

	}
	return errMap
}