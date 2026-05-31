package main

import "fmt"

func main() {

	//?VARIABLE AND TYPES
	// in GO, there is different types of variables, for example:
	//int- //store integers(whole numbers) such as 123 or -123,
	//float32- // stores floating point number with decimals such as 19.99,
	//string - // stire text, such as "Hello world". String values are surrounded by double quotes.
	//bool - //stores values with two state; true or false

	//!DECLARING VARIABLES.
	//in go we declare variables in two ways.
	//? 01- Use the var keyword, followed by variable name and type

	var school string = "minnesota schools"
	//!NOTE: You always have to specify either type or value(or both).
	fmt.Println(school)

	//? 02 - With the := sign
	//Use the := sign, followed by the variable value.

	company := "apilox"
	fmt.Println(company)
	//!NOTE: In this case, the type of the variable is inferred from the value(means that the compiler decides the type of the variable based on the value). ELSE - it is not possible to declare a variable using:= without assigning a value to it.

	//? VARIABLE DECLARATION WITH INITIAL VALUE.
	/* if the value of a variable is known from the start, you can declare variable and assign a value to it in one line */

	var student1 string = "John" //type is string
	var student2 = "Jane"        // type is inferred.
	x := 2                       // type is inferred.

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	//? VARIABLE DECLARATION WITHOUT INITIAL VALUES.
	/* in Go, all variables are initialized. So, if you declare a variable withouy an initial value, its value will be set to the default value of its type: */

	/* e.g's */
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//! NOTES:- It is possible to declare variable and initialize it (assign value to it later) by normal method. BUT using := you can't declare variable without initializing it.

	//? DIFFERENT BETWEEN var AND :=;
	//VAR:
	/*      01- Can be used inside function and outside.
			02 - variable declaration and value assignment can be done separately.


	:=
			01- can only be used inside functions
			02 - variable declaration and value assignment cannot be done separetely (must be done in the same line)
	*/

	//? MULTIPLE VALUE DECLARATION.
	/* In Go, it is possible to declare multiple variables on the same line */
	var v, d, f, g, h int = 1, 2, 3, 4, 5

	fmt.Println(v)
	fmt.Println(d)
	fmt.Println(f)

	fmt.Println(g)
	fmt.Println(h)
	// fmt.Println(c)

	//! NOTE: If you use TYPE keyword, it's only possible to declare one type of variable per line and VICE-VERSA

	//? GO VARIABLE DECLARATION IN A BLOCK
	/* Multiple variable declarations can also be grouped together into a block for a greater readability */

	var (
		m int
		y int    = 1
		j string = "hello"
	)

	fmt.Println(m)
	fmt.Println(y)
	fmt.Println(j)

	//? GO VARIABLE NAMING RULES
	/* A variable can have a short name like x and y or more descriptive name (age, price, carname, etc)
	There is rules in naming Go variables:
			01- A variable must start with a letter or underscore character(_)
			02- A variable cannot start with a digit
			03- A variable can only conatin alpha numeric caharacters and underscore(A-Z, a-z, 0-9, and _)
			04 - Variables names are case sensitive. (age, Age, AGE are three different variables)
			05 - There is no limit of length of variable name.
			06 - a variable name cannot contain a space
			07 - a variable name cannot be any Go keyword.

	*/

	//? TOPIC 07.
	//? Go CONSTANTS

	/* If variable should have a fixed value that cannot be changed, you can use CONST keyword.
	The CONST keyword declares the variables as "constant" which means that it is unchangable and read-only. */

	const CONSTNAME string = "FunMeter"

	//!NOTE: The value of constant must be assigned when you declare it.

	//?DECLARING CONSTANT

	const PI = 3.14

	//CONSTANT RULES.
	/*    01- Constant name follow the same naming rules as variables
	02- constants names are usually written in UPPERCASE letters (for easy identification and differentiation from variables)
	03 - Costants can be declared both inside and outside of a function
	*/

	//? CONSTANT TYPES
	/* There are two types of constants
	01-typed constants
	02-Untyped constants
	*/

	//? TYPED CONSTANTS.
	//Typed constants are declared with defined type:
	const A int = 1

	//? UNTYPED CONSTANTS
	//Untyped constants are declared without defined tye.
	const B = 1

	//!NOTE:  01- In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value)
	//!       02- When constant is declared, it is not possible to change the value later.

	//? MULTIPKE CONSTANTS DECLARATION
	/* Multiple constants can be grouped together into a block for readbility */

	const (
		C int = 1
		D     = 3.13
		Y     = "Hi!"
	)

}
