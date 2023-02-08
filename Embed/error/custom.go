package main

import (
    "fmt"
)

type UserNotAuthorizedError struct {
    message string
}

func (e *UserNotAuthorizedError) Error() string {
    return e.message
}

func someFunction() error {
    return &UserNotAuthorizedError{message: "User is not authorized to perform this action"}
}

func main() {
    err := someFunction()
    if err != nil {
        fmt.Println(err)
        // Output: User is not authorized to perform this action
    }
}
