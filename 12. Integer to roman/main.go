package main

import (
	"fmt"
	"strconv"
)

func main() {
var num int = 3490

var res string = convToRoman(num)
fmt.Println(res)
}

func convToRoman(num int) string {
	var romanNumber string = ""
	numberStr := strconv.Itoa(num)

	switch(len(numberStr)) {
		case 1:
				fmt.Println("единицы")
				romanNumber = units(num)
		case 2:
				fmt.Println("десятки")
				var remains int = num % 10
				romanNumber = dozens(num-remains)

				if remains > 0 {
					romanNumber += units(remains)
				}
		case 3:
				fmt.Println("сотни")
				var remainsWeaving int = num % 100 // остаток от 100
				romanNumber = weaving(num-remainsWeaving)
				var remainsDozeng int = num % 10 // остаток от 10
				romanNumber += dozens(remainsWeaving - remainsDozeng)

				if remainsDozeng % 10 > 0 {
				romanNumber += units(remainsDozeng % 10)
				}

				case 4:
				fmt.Println("тысячи")
				var remainsThousandths int = num % 1000 // остаток от 1000
				romanNumber = thousandths(num - remainsThousandths)
				var remainsWeaving int = remainsThousandths % 100 // остаток от 100
				romanNumber += weaving(remainsThousandths-remainsWeaving)
				var remainsDozeng int = num % 10 // остаток от 10
				romanNumber += dozens(remainsWeaving - remainsDozeng)

				if remainsDozeng % 10 > 0 {
				romanNumber += units(remainsDozeng % 10)
				}
}

return romanNumber

}

func units(num int) string {
	var result string

	// проверка на римскую спец комбинацию
	isSpeciaCombine, resultSpecialCombine := isSpecialCombine(num)
	if(isSpeciaCombine) {
			result += resultSpecialCombine
			return result
	}

	//  перебираем самые примитивные варианты которые присутсвую в римском алфавите
	isPrimitiveCombine, resultPrimitive := isPrimitiveVariant(num)
	if(isPrimitiveCombine) {
		result = resultPrimitive
		return result
	}

	// если num > 5
	if num > 5 {
		result = getRomanABC(5)
		var remains int = num % 5

		// проверяем на римскую спец комбинацию и возвращаем результат
		isSpeciaCombine, resultSpecialCombine := isSpecialCombine(remains)
		if(isSpeciaCombine) {
				result += resultSpecialCombine
				return result
		}

		for i:= 0; i < remains; i++ {
			result += getRomanABC(1)
		}

		return result
	}

	for i:= 0; i < num; i++ {
		result += getRomanABC(1)
	}
	return result
}

func getRomanABC(num int) string {
var romanAbc map[int]string = map[int]string{
		1: "I",
		5: "V",
		10: "X",
		50: "L",
		100: "C",
		500: "D",
		1000: "M",
	}

	return romanAbc[num]
}

func isSpecialCombine(num int) (bool, string) {
	var isSpecialSymbol bool = false
	var result string = ""

	switch(num){
	case 4, 9, 40, 90, 400, 900:
		isSpecialSymbol = true
	}

	switch(num){
	case 4:
		result = getRomanABC(1)
		result += getRomanABC(5)
	case 9:
		result = getRomanABC(1)
		result += getRomanABC(10)
	case 40:
		result = getRomanABC(10)
		result += getRomanABC(50)
	case 90:
		result = getRomanABC(10)
		result += getRomanABC(100)
	case 400:
		result = getRomanABC(100)
		result += getRomanABC(500)
	case 900:
		result = getRomanABC(100)
		result += getRomanABC(1000)
}

	return isSpecialSymbol, result
}

// работа с десятками
func dozens(num int) string {
	var result string = ""

	// проверяем на римские спец комбинации
	isSpeciaCombine, resultSpecialCombine := isSpecialCombine(num)
	if(isSpeciaCombine) {
			result += resultSpecialCombine
			return result
	}

	//  перебираем самые примитивные варианты которые присутсвую в римском алфавите
	isPrimitiveCombine, resultPrimitive := isPrimitiveVariant(num)
	if(isPrimitiveCombine) {
		result = resultPrimitive
		return result
	}

	// если num > 50
	if num > 50 {
		result = getRomanABC(50)
		var remains int = num % 50

		// проверяем на римскую спец комбинацию и возвращаем результат
		isSpeciaCombine, resultSpecialCombine := isSpecialCombine(remains)
		if(isSpeciaCombine) {
				result += resultSpecialCombine
				return result
		}

		for i:= 0; i < remains; i += 10 {
			result += getRomanABC(10)
		}

		return result
	}

	for i:= 0; i < num; i += 10 {
		result += getRomanABC(10)
	}
	return result
}

func isPrimitiveVariant(num int) (bool, string){
	var isPrimitive bool = false
	var result string = ""

	switch(num){
	case 1,5,10,50,100,500,1000:
		result = getRomanABC(num)
		isPrimitive = true
	}

	return isPrimitive, result
}

// работа с сотнями
func weaving(num int) string {
	var result string = ""

	// проверяем на римские спец комбинации
	isSpeciaCombine, resultSpecialCombine := isSpecialCombine(num)
	if(isSpeciaCombine) {
			result += resultSpecialCombine
			return result
	}

	//  перебираем самые примитивные варианты которые присутсвую в римском алфавите
	isPrimitiveCombine, resultPrimitive := isPrimitiveVariant(num)
	if(isPrimitiveCombine) {
		result = resultPrimitive
		return result
	}

	// если num > 50
	if num > 500 {
		result = getRomanABC(500)
		var remains int = num % 500

		// проверяем на римскую спец комбинацию и возвращаем результат
		isSpeciaCombine, resultSpecialCombine := isSpecialCombine(remains)
		if(isSpeciaCombine) {
				result += resultSpecialCombine
				return result
		}

		for i:= 0; i < remains; i += 100 {
			result += getRomanABC(100)
		}

		return result
	}

	for i:= 0; i < num; i += 100 {
		result += getRomanABC(100)
	}
	return result
}

func thousandths(num int) string {
	var result string = ""

	// проверяем на римские спец комбинации
	isSpeciaCombine, resultSpecialCombine := isSpecialCombine(num)
	if(isSpeciaCombine) {
			result += resultSpecialCombine
			return result
	}

	//  перебираем самые примитивные варианты которые присутсвую в римском алфавите
	isPrimitiveCombine, resultPrimitive := isPrimitiveVariant(num)
	if(isPrimitiveCombine) {
		result = resultPrimitive
		return result
	}

		for i:= 0; i < num; i += 1000 {
			result += getRomanABC(1000)
		}

	return result
}
