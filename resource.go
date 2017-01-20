package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Resource struct {
	Valid    bool
	Resource string `json:"resource"`
}

func isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func resourceParse(resource string) Resource {
	fmt.Println(resource)
	validator := Resource{
		Valid:    false,
		Resource: resource,
	}
	if isValidUUID(validator.Resource) {
		validator.Valid = true
	}
	return validator
}

func Uuid() string {
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		panic(err)
	}
	return strings.Replace(string(uuid), "\n", "", -1)
}
