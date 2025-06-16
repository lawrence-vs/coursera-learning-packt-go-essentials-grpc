package adult

import (
	"errors"
	"fmt"
	// "log"
)

type Adult struct {
	Age  int
	Name string
}

// encapsulating - setter data validated
func (a *Adult) GetAge() int {
	return a.Age
}

func (a *Adult) SetAge(newAge int) error {
	if newAge < 1 {
		return errors.New("invalid age")
	}
	a.Age = newAge
	return nil
}

func (a *Adult) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old (an Adult).", a.Name, a.Age)
}
