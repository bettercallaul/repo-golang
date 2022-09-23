package main

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var empployee Employee
	empployee.name = "Sultan Aulia"
	empployee.age = 23
	empployee.division = "Divisi Humas"

	fmt.println("Namanya: ", empployee.name)
	fmt.println("Umurnya: ", empployee.age)
	fmt.println("Divisinya: ", empployee.division)
}
