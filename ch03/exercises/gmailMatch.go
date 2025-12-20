package main

import (
    "fmt"
    "regexp"
)

func isValidGmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9_-]+@gmail\.com$`)
    return re.MatchString(email)
}

func main() {
    emails := []string{
        "user@gmail.com",
        "user.name@gmail.com",
        "user+tag@gmail.com",
        "user_name@gmail.com",
        "user123@gmail.com",
        "invalid@yahoo.com",
        "@gmail.com",
        "user@gmail",
    }
    
    for _, email := range emails {
        fmt.Printf("%-25s -> %v\n", email, isValidGmail(email))
    }
}