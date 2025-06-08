package main

import "fmt"

type person struct {
	age int 
	name string
}

func (p *person) changeUser() {
	p.age = 20
	p.name = "Bobby"
}

func main(){
	user := person{35,"Fred"}
	user.changeUser()
	showUser(&user)
}

func showUser(p *person){
	//p.age = 35
	//p.name = "Steve"
	fmt.Println(*p)
}
