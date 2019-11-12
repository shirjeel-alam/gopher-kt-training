package main

import (
  "errors"
  "fmt"
)

// Given the following two structs, embed an email into the Invite struct
// to make the `main` program below work.
type Invite struct {
  ID int
  Email
}

type Email struct {
  DisplayName string
  Username    string
  Domain      string
}

// Create a method on Email called `Address` that takes no arguments
// It should return a string in the format:
// `DisplayName <Username@Domain>`
// You can use `fmt.Sprintf` to create the return string
func (e Email) Address() string {
  return fmt.Sprintf("%s <%s@%s>", e.DisplayName, e.Username, e.Domain)
}

func main() {
  invite := Invite{
    ID: 1,
    Email: Email{
      DisplayName: "John Smith",
      Username:    "jsmith",
      Domain:      "yahoo.com",
    },
  }
  if err := SendEmail(invite); err != nil {
    fmt.Printf("error sending email: %s", err)
    return
  }
  fmt.Printf("sent email to: %s", invite.Address())
}

func SendEmail(i Invite) error {
  // Do some work here
  address := i.Address()
  if address == "" {
    return errors.New("no email address provided")
  }
  return nil
}