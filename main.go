package main

import (
	"fmt"
	"strconv"
	"strings"
)

var SglDigit, DblDigit, TensPlace []string

func PriceInWords(price string) string {
	SglDigit = append(SglDigit, "Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine")
	DblDigit = append(DblDigit, "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen")
	TensPlace = append(TensPlace, "", "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety")

	words := []string{}
	digit := 0
	nxtDigit := 0
	priceLen := len(price)
	intp, _ := strconv.Atoi(price)
	if intp > 0 && priceLen <= 10 {

		for i := priceLen - 1; i >= 0; i-- {

			n := price[i]
			digit, _ = strconv.Atoi(string(n))
			if i > 0 {
				x := price[i-1]
				nxtDigit, _ = strconv.Atoi(string(x))
			} else {
				nxtDigit = 0
			}

			switch priceLen - i - 1 {

			case 0:
				a, b := value2(digit, nxtDigit, "")
				words = append(words, b, a)

			case 1:
				n := price[i+1]
				b, _ := strconv.Atoi(string(n))
				a := value(digit, b)
				words = append(words, a)

			case 2:

				if digit != 0 {
					a := SglDigit[digit]
					words = append(words, "Hundred", a)
					n := price[i+1]
					b, _ := strconv.Atoi(string(n))
					m := price[i+2]
					c, _ := strconv.Atoi(string(m))
					if b != 0 && c != 0 {
						words = append(words, "and")
					} else {
						words = append(words, "")
					}
				} else {
					words = append(words, "")
				}

			case 3:
				a, b := value2(digit, nxtDigit, "Thousand")
				words = append(words, b, a)

			case 4:
				n := price[i+1]
				b, _ := strconv.Atoi(string(n))
				a := value(digit, b)
				words = append(words, a)

			case 5:
				a, b := value2(digit, nxtDigit, "Lakh")
				words = append(words, b, a)

			case 6:
				n := price[i+1]
				b, _ := strconv.Atoi(string(n))
				a := value(digit, b)
				words = append(words, a)

			case 7:
				a, b := value2(digit, nxtDigit, "Crore")
				words = append(words, b, a)

			case 8:
				n := price[i+1]
				b, _ := strconv.Atoi(string(n))
				a := value(digit, b)
				words = append(words, a)

			case 9:
				if digit != 0 {
					w := " " + " Hundred" + SglDigit[digit]
					if price[i+1] != 0 || price[i+2] != 0 {
						w += " and"
					} else {
						w += " Crore"
					}
					words = append(words, w)
				} else {
					words = append(words, "")
				}
			}

		}
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return strings.Join(words, " ")
	}
	return strings.Join(words, " ")

}

func value(dgt, prevDgt int) string {
	var w string
	if dgt == 0 {
		w = ""
		return w
	} else {
		w = " "
		if dgt == 1 {
			w += DblDigit[prevDgt]
		} else {
			w += TensPlace[dgt]
		}
		return w
	}
}

func value2(dgt, nxtDgt int, denom string) (string, string) {
	if (dgt != 0 && nxtDgt != 1) && (nxtDgt != 0 || dgt > 0) {
		a := "" + SglDigit[dgt]
		b := "" + denom
		return a, b
	} else {
		return "", ""
	}
}

func main() {

	number := 9865
	numberString := strconv.Itoa(number) //number converted to string
	priceInWords := PriceInWords(numberString)
	fmt.Println("Number IN string =", numberString)
	fmt.Println("price in words :", priceInWords)//Nine Thousand and Eight Hundred  Sixty Five 
}
