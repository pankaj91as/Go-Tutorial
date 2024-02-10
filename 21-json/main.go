package main

import (
	"encoding/json"
	"fmt"
)

type Users struct {
	Name     string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Status   bool     `json:"status"`
	Optin    []string `json:"courses"`
}

func main() {
	// create user data
	usersData := []Users{
		{"Username 1", "username1@gmail.com", "abc123", true, []string{"WordPress", "Joomla"}},
		{"Username 2", "username2@gmail.com", "abc123", true, []string{"ModX", "Magento"}},
		{"Username 3", "username3@gmail.com", "abc123", true, []string{"OpenCart"}},
		{"Username 4", "username4@gmail.com", "abc123", true, nil},
	}

	ConvertToJSON(usersData)
	ConvertToFormatedJSON(usersData)

	// JSON Byte Data
	jsonData := []byte(`{
		"username": "Username 3",
		"email": "username3@gmail.com",
		"status": true,
		"courses": [
				"OpenCart"
		]
	}`)
	parseJSONData(jsonData)

	parseJSONWithoutStruct(jsonData)
}

func ConvertToJSON(usersData []Users) {
	// parse data into JSON
	parsedData, err := json.Marshal(usersData)
	if err != nil {
		panic(err)
	}

	// print unformated JSON
	fmt.Printf("%s\n", parsedData)
}

func ConvertToFormatedJSON(usersData []Users) {
	// print formated JSON
	parsedFormatedData, err := json.MarshalIndent(usersData, "", "\t")
	if err != nil {
		panic(err)
	}

	// print unformated JSON
	fmt.Printf("%s\n", parsedFormatedData)
}

func parseJSONData(body []byte) {
	var usersData Users
	isValidJson := json.Valid(body)
	if isValidJson {
		fmt.Println("JSON is valid & i am going to unmarshal it!")
		json.Unmarshal(body, &usersData)
	} else {
		fmt.Println("JSON is not valid")
	}

	// print valid json
	fmt.Println(usersData)

	// print json with key value
	fmt.Printf("%#v\n", usersData)
}

func parseJSONWithoutStruct(body []byte) {
	var myUsersJson map[string]interface{}
	json.Unmarshal(body, &myUsersJson)

	fmt.Println(myUsersJson)
	for k, v := range myUsersJson {
		fmt.Printf("Key: %v Value: %v Type: %T\n", k, v, v)
	}
}
