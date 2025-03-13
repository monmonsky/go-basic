package models

// Exported variable
// var User string = "ini dari models"

// Unexported variable
// var data []string = []string{"monsky", "nobeeid", "nobi"}

// inalias type
// type Angka int

// var User = struct {
// 	Id int
//	Email string
// }{
// 	Id: 10,
//  Email: "monsky@k3monspace.com"
// }

type User struct {
	Id    *int   `json:"id, omitempty"`
	Email string `json:"email"`
}

type UserWithAddress struct {
	Id      int
	Email   string
	Address string
}

var User2 = UserWithAddress{
	Id:      20,
	Email:   "monsky@k3monspace.com",
	Address: "Jakarta",
}

var User3 = User{
	Email: "",
}
