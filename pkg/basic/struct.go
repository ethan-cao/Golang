package basic

import "fmt"

// a struct is collection of different data types
// similar to Java calss
type User struct {
	ID    int
	Name  string
	Email string
}

// define func on struct
// the receiver u is a pointer to the User struct that this method will be called on.
func (u *User) Details() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", u.ID, u.Name, u.Email)
}

func test() {
	user := User{
		ID:    1,
		Name:  "ethan",
		Email: "test",
	}

	fmt.Println(user.Name) // ethan
}
