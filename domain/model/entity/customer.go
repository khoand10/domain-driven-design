package entity

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

//type Address struct {
//	Street  string
//	City    string
//	State   string
//	ZipCode string
//}
