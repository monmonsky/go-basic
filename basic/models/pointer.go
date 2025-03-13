package models

import "fmt"

func RunPointer() {
	umur := 20
	umurB := &umur

	fmt.Println(umur, *umurB)
	umur = 10
	fmt.Println(umur, *umurB)
	*umurB = 30
	fmt.Println(umur, *umurB)

	var user UserWithAddress = UserWithAddress{
		Id: 10,
		Email:	"monsky2@k3monspace.com",
	}

	// SetNameNative("test2@k3monspace.com", user)
	// fmt.Println(user)

	SetNamePointer("test2@k3monspace.com", &user)
	fmt.Println(user)
}

func SetNamePointer(newName string, user *UserWithAddress) {
	user.Email = newName
}

func SetNameNative(newName string, user UserWithAddress) {
	user.Email = newName
}

func SetName(newName string) string {
	return newName
}