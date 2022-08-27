package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(genPass("fazztract", "Low"))
	fmt.Println(genPass("fazztract", "Medium"))
	fmt.Println(genPass("fazztract", "Strong"))

}

func genPass(password string, level string) string {
	var randomChar = strconv.Itoa(rand.Intn(1000) + 100)

	if len(password) >= 6 {
		if level == "Low" {
			password += randomChar

			return password

		} else if level == "Medium" {
			password += randomChar

			slice := make([]string, len(password))
			for i := range password {
				if rand.Int()%2 == 0 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			return result

		} else if level == "Strong" {
			slice := make([]string, len(password))
			for i := range password {
				if rand.Int()%2 == 0 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			specialChars := "!@#$%&*"
			for i := 0; i < 2; i++ {
				result += string(specialChars[rand.Intn(len(specialChars))])
			}

			result += randomChar

			return result

		} else {
			return "level selection is't correct"
		}
	} else {
		return "Password min 6 characters"
	}
}
