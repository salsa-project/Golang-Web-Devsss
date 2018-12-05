package main

import "fmt"

var c int
var d = 43

type person struct{
	first string
	age int
	saying string
}

//func receiver identifier(params) returns {code}

func (p person) speak(){
	fmt.Println(p.first, "says", p.saying)
}

type secretagent struct{
	person
	ltk bool
}

func (sa secretagent) speak(){
	fmt.Println(sa.first, "says even more ", sa.saying)
}

type human interface {
	speak()
}

func foo(h human){
	h.speak()
}

func main() {

	a := "james"
	fmt.Println("a")
	fmt.Printf("%T\n", a)

	b := fmt.Sprint("Hello", a)
	fmt.Println(b)

	c = 42
	fmt.Println("c:", c)

	d =43
	fmt.Println("d:", d)
	fmt.Printf("d type: %T\n", d)


			//Slice
			ee := []int{1,2,3,4,5}
			fmt.Println(ee)

			for i, v := range ee {
			fmt.Println(i, v)

			//Map 
			f := map[string]int{"salsa":30, "project":27}
			fmt.Println(f)

			for k, v := range f {
				fmt.Println(k, v)
			}

			//Struct
			p1 := person{
				first: "SalSa",
				age: 26,
				saying: "doing is better than staying in place",
			}
			p2 := person{
				first: "MrSalSa",
				age: 27,
				saying: "I tried till i couldn't find what to choose",
			}

			xp := []person{p1, p2}
			fmt.Println(xp)
			for i2, v2 := range xp{
				fmt.Println(i2, v2)
			}

			//loop init, cond, post
			for j := 0; j < 10; j++{
				fmt.Println(j)
			}

			sa1 := secretagent{
				person: person{
					first: "jack",
					age: 29,
					saying: "nice when you work to make it nice!",
				},
				ltk: true,
			}
			fmt.Println(sa1)
			fmt.Println(sa1.first)
			fmt.Println(sa1.person.first)

			p1.speak()
			p2.speak()
			sa1.speak()
			fmt.Println("--")
			foo(p1)
			foo(p2)
			foo(sa1)
		}
	}
