/**
* @author Rabira Motuma
* @version 1.0.0
* @date 2026-01-13
* @fileoverview This program uses functions and if statements to act as a calculator.
*/

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
let operation: string;
let number1String: string;
let number2String: string;
let number1: number;
let number2: number;
let answer: number

// introduction
console.log("Welcome to my calculator program.")
console.log("Which operation would you like to perform today? (Select by typing the letter in front of the operation.)")
console.log("a. add");
console.log("b. subtract");
console.log("c. multiply");
console.log("d. divide");
console.log("e. absolute value");
console.log("f. round");
console.log("g. raise to an exponent");
console.log("h. square root");

// calculation
operation = prompt("What operation do you want to choose:") || "0";

// if
if (operation == ADD) {
  number1String = prompt("Enter first value:") || "0";
  number2String = prompt("Enter second value:") || "0";
  number1 = parseFloat(number1String);
  number2 = parseFloat(number2String);
  answer = number1 + number2;
  console.log(`${number1} + ${number2} = ${answer}`);
}
else if (operation == SUBTRACT) {
  number1String = prompt("Enter first value:") || "0";
  number2String = prompt("Enter second value:") || "0";
  number1 = parseFloat(number1String);
  number2 = parseFloat(number2String);
  answer = number1 - number2;
  console.log(`${number1} - ${number2} = ${answer}`);
}
else if (operation == MULTIPLY) {
  number1String = prompt("Enter first value:") || "0";
  number2String = prompt("Enter second value:") || "0";
  number1 = parseFloat(number1String);
  number2 = parseFloat(number2String);
  answer = number1 * number2;
  console.log(`${number1} * ${number2} = ${answer}`);
}
else if (operation == DIVIDE) {
  number1String = prompt("Enter first value:") || "0";
  number2String = prompt("Enter second value:") || "0";
  number1 = parseFloat(number1String);
  number2 = parseFloat(number2String);
  answer = number1 / number2;
  console.log(`${number1} / ${number2} = ${answer}`);
}
else if (operation == ABSOLUTE) {
  number1String = prompt("Enter first value:") || "0";
  number1 = parseFloat(number1String);
  answer = Math.abs(number1)
  console.log(`The absolute value of ${number1} is ${answer}`);
}
else if (operation == ROUND) {
  number1String = prompt("Enter first value:") || "0";
  number1 = parseFloat(number1String);
  answer = Math.round(number1)
  console.log(`${number1} rounded is ${answer}`);
}
else if (operation == EXPONENT) {
  number1String = prompt("Enter first value:") || "0";
  number2String = prompt("Enter second value:") || "0";
  number1 = parseFloat(number1String);
  number2 = parseFloat(number2String);
  answer = Math.pow(number1, number2)
  console.log(`${number1} raised to the power of ${number2} = ${answer}`);
}
else if (operation == SQRT) {
  number1String = prompt("Enter first value:") || "0";
  number1 = parseFloat(number1String);
  answer = Math.sqrt(number1)
  console.log(`The square root of ${number1} = ${answer}`);
}
else {
  console.log("You have entered an invalid input in a step.")
}

console.log("\nDone.")
