package models

type Store struct {
	Pants   string `json: "pants"`
	Shoes   string `json: "shoes"`
	TShirts string `json: "tshirts"`
	//DiorDress  []int64 `json: "dress"`
}

type Order struct {
	OrID    int64  `json: "id"`
	Pants   string `json: "pants"`
	Shoes   string `json: "shoes"`
	TShirts string `json: "tshirt"`
	//DiorDress  []int64 `json: "dress"`
}

type DiorDress struct {
}
