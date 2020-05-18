package models

type Captain struct {
	Id   int    `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

func NewCaptain(age int, name string) *Captain {

	captain := Captain{Age: age, Name: name}
	return &captain
}
