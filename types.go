package main

type Product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	UrlTitle string  `json:"url_title"`
	Images   string  `json:"images"`
	Descr    string  `json:"descr"`
	Specs    string  `json:"specs"`
	Price    int     `json:"price"`
	Weight   int     `json:"weight"`
	Length   int     `json:"length"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
}

  type User struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	StreetName  string `json:"street_name"`
	HouseNumber string `json:"house_number"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
}
