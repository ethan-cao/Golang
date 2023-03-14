package basic

import "fmt"

// a struct is collection of different data types
// similar to Java calss
type User struct {
	ID    int
	Name  string
	Email string
}

// receiver function
// define func on struct, Java claass method
func (u User) Details() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", u.ID, u.Name, u.Email)
}

func TestStruct() {
	user := User{
		ID:    1,
		Name:  "ethan",
		Email: "test",
	}

	fmt.Println(user.Name) // ethan
	fmt.Println(user.Details())
}
