package main

type GoFooBar struct {
	ID	string	`gorm:"primarykey;size:16"`
	Foo	string 	`json:"foo"`
	Bar int 	`json:"bar"`
}
