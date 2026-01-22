/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2026-01-13
 * Fileoverview: This program uses functions and if statements to act as a calculator.
 */

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main () {
	// constants
	const ADD  = "a";
  const SUBTRACT  = "b";
  const MULTIPLY  = "c";
  const DIVIDE  = "d";
  const ABSOLUTE  = "e";
  const ROUND  = "f";
  const EXPONENT  = "g";
  const SQRT  = "h";

	// variables
	var operation string
  var number1String string
  var number2String string
  var number1 float64
  var number2 float64
  var answer float64

	reader := bufio.NewReader(os.Stdin)

	// introduction
	fmt.Println("Welcome to my calculator program.")
  fmt.Println("Which operation would you like to perform today? (Select by typing the letter in front of the operation.)")
  fmt.Println("a. add")
  fmt.Println("b. subtract")
  fmt.Println("c. multiply")
  fmt.Println("d. divide")
  fmt.Println("e. absolute value")
  fmt.Println("f. round")
  fmt.Println("g. raise to an exponent")
  fmt.Println("h. square root")

	// prompt
	fmt.Print("What operation do you want to choose: ")
	operation, _ = reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	// if statements
	if operation == ADD {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		fmt.Printf("Enter the second value: ")
		number2String, _ = reader.ReadString('\n')
		number2String = strings.TrimSpace(number2String)
		number2, _ = strconv.ParseFloat(number2String, 64)

		answer = number1+number2
		fmt.Printf("%f + %f = %f\n", number1, number2, answer)
	} else if  operation == SUBTRACT {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		fmt.Printf("Enter the second value: ")
		number2String, _ = reader.ReadString('\n')
		number2String = strings.TrimSpace(number2String)
		number2, _ = strconv.ParseFloat(number2String, 64)

		answer = number1-number2
		fmt.Printf("%f - %f = %f\n", number1, number2, answer)
	} else if  operation == MULTIPLY {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		fmt.Printf("Enter the second value: ")
		number2String, _ = reader.ReadString('\n')
		number2String = strings.TrimSpace(number2String)
		number2, _ = strconv.ParseFloat(number2String, 64)

		answer = number1*number2
		fmt.Printf("%f * %f = %f\n", number1, number2, answer)
	} else if  operation == DIVIDE {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		fmt.Printf("Enter the second value: ")
		number2String, _ = reader.ReadString('\n')
		number2String = strings.TrimSpace(number2String)
		number2, _ = strconv.ParseFloat(number2String, 64)

		answer = number1/number2
		fmt.Printf("%f / %f = %f\n", number1, number2, answer)
	} else if  operation == ABSOLUTE {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		answer = math.Abs(number1)
		fmt.Printf("The absolute value of %f = %f\n", number1, answer)
	} else if  operation == ROUND {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		answer = math.Round(number1)
		fmt.Printf("%f rounded = %f\n", number1, answer)
	} else if  operation == EXPONENT {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		fmt.Printf("Enter the second value: ")
		number2String, _ = reader.ReadString('\n')
		number2String = strings.TrimSpace(number2String)
		number2, _ = strconv.ParseFloat(number2String, 64)

		answer = math.Pow(number1, number2)
		fmt.Printf("%f to the power of %f = %f\n", number1, number2, answer)
	}else if  operation == SQRT {
		fmt.Printf("Enter the first value: ")
		number1String, _ = reader.ReadString('\n')
		number1String = strings.TrimSpace(number1String)
		number1, _ = strconv.ParseFloat(number1String, 64)

		answer = math.Sqrt(number1)
		fmt.Printf("The square root of %.2f = %.2f\n", number1, answer)
	} else {
		fmt.Println("You have entered an invalid input.")
	}

	fmt.Println("\nDone.")
}
