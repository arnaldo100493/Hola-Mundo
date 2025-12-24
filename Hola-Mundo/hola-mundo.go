package main

import (
	"fmt"

	"rsc.io/quote"
)

//var firstName, lastName, age = "Alex", "Roel", 27

// Declaración de constantes
const Pi float32 = 3.14

const (
	X = 100
	Y = 0b1010 //Binario
	Z = 0o12   //Octal
	W = 0xFF   //Hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola, Mundo!")
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	var name string
	var age int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name)
	//fmt.Scanln(&name, &name2)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)

	fmt.Printf("El tipo de name es: %T\n", name)
	fmt.Printf("El tipo de name es: %T\n", age)




	/*name := "Alex"
	//age := 28

	//fmt.Printf("Hola, me llamo %s %s y tengo %d años.\n", name, name2, age)

	//greeting := fmt.Sprintf("Hola, me llamo %s %s y tengo %d años.\n", name, name2, age)

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.\n", name, age)

	fmt.Println(greeting)

	//fmt.Print("Otro mensaje")

	//var integer16 int16 = 50
	//var integer32 int32 = 100

	s := "100"
	i, _ := strconv.Atoi(s)

	fmt.Println(i + i)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s + s)

	//fmt.Println(int32(integer16) + integer32)

	var (
		defaultInt    uint
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	fullName := "Alex Roel \t(alias \"roelcode\")\n"
	fmt.Println(fullName)

	var a byte = 'a'
	fmt.Println(a)

	s := "hola"
	fmt.Println(s[0])

	var r rune = '❤'
	fmt.Println(r)

	//Pi = 3.1415 //Error

	//Declaración de variables

	//var firstName, lastName string
	//var age int

	var (
		firstName string = "Alex"
		lastName  string = "Roel"
		age       int    = 27
	)

	//var firstName, lastName, age = "Alex", "Roel", 27

	//firstName, lastName, age := "Alex", "Roel", 27

	firstName, lastName := "Alex", "Roel" //Esta declaración de variables con := solo funcionan dentro de las funciones.
	age := 27
	age = 30

	firstName = "Alex"
	lastName = "Roel"
	age = 27

	fmt.Println(firstName, lastName, age)
	fmt.Println(Pi)
	fmt.Println(X, Y, Z, W)
	fmt.Println(Viernes)*/
}
