package main

import (
  "errors"
  "fmt"
)

type Invite struct {
  ID int
  Email
}

type Email struct {
  DisplayName string
  Username    string
  Domain      string
}

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

// Declare an interface called `Addressable` that defines the `Address` method
// from the Email struct above

type Addressable interface {
  Address() string
}

// Change the signature of the `SendEmail` function to take
// an `Addressable` interface, instead of the concrete type `Invite`
func SendEmail(a Addressable) error {
  // Do some work here
  address := a.Address()
  if address == "" {
    return errors.New("no email address provided")
  }
  return nil
}