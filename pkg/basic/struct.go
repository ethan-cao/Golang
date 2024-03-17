package basic

import "fmt"

type Human struct {
	ID   int
	Name string
}

type User struct {
	Human //  struct embedding, field in Human promoted to User
	Email string
}

// receiver function
// define func on struct, Java class method
func (user User) Details() string {
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s", user.ID, user.Name, user.Email)
}

func TestStruct() {
	user := User{
		Human: Human{
			ID:   1,
			Name: "ethan",
		},
		Email: "email",
	}

	fmt.Println(user.Name) // ethan
	fmt.Println(user.Details())
}
