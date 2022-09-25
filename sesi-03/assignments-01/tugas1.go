package main

//import package
import (
	"fmt"
	"os"
	"strconv"
)

// struct person

type Person struct {
	// struct innitialization
	Number int
	Name string
	Address string
	Occupation string
	Reasons string
}

// function main
func main() {
var partisipan = []Person {
	{Number: 1, Name: ": Thor Odinson", Address: ": Valhalla", Occupation: ": God of Thunder", Reasons: ": Awesome Mantap Keren Abiez"},
	{Number: 2, Name: ": Bruce Banner", Address: ": Boston", Occupation: ": Scientist", Reasons: ": Ijo Ijo Lucu"},
	{Number: 3, Name: ": Peter Parker", Address: ": Queens", Occupation: ": Man of Spider", Reasons: ": Swinging is Cool"},
	{Number: 4, Name: ": Stephen Strange", Address: ": New York Besar", Occupation: ": Doctor", Reasons: ": Saving people is my passion"},
	{Number: 5, Name: ": Clint Barton", Address: ": Wyoming", Occupation: ": Archery Instructor", Reasons: ": No reason to live beside my family"},
	{Number: 6, Name: ": Scott Lang", Address: ": South Hampton", Occupation: ": Cybersecurity", Reasons: ": Hacking the rich's money so I can give it away to the poor"},
	{Number: 7, Name: ": Matt Murdock", Address: ": Queens", Occupation: ": Lawyer", Reasons: ": At age 8, I lost my eyes and since then i decided to become a lawyer"},
}



var argsRaw = os.Args
// inputting Argsraw to array[1]
var args = argsRaw[1] 
// string to integer convert
num, err := strconv.Atoi(args)
// using err to cointain zero value
_ = err
// generateParticipants as function with partisipan & num as parameter
generateParticipants(partisipan, num)
}

func generateParticipants(p []Person, num int) {
	if num <= 7 {
		fmt.Println("Person Data")
		fmt.Println("Number", p[num-1].Number)
		fmt.Println("Name", p[num-1].Name)
		fmt.Println("Address", p[num-1].Address)
		fmt.Println("Occupation", p[num-1].Occupation)
		fmt.Println("Reasons", p[num-1].Reasons)
	} else {
		fmt.Println("there's no data to that person")
	}
}

