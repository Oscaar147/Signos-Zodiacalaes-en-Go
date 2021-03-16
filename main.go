package main

import "fmt"

func main() {
	fmt.Print("\n\nBienvenido al calendario de Signos zodiacales\n")

	var mes int
	var dia int

	fmt.Print("Que dia naciste?: ")
	fmt.Scan(&dia)
	//	if !IsValidDay(dia) {
	//	fmt.Print("Invalid day")
	//	}
	fmt.Print("De que mes?: ")
	fmt.Scan(&mes)
	//	if !IsValidMonth(mes) {
	//	fmt.Print("Invalid month")
	//	}

	signoZodiacal := calcularSigno(mes, dia)

	fmt.Print("Tu signo del zodiaco es " + signoZodiacal + "\n\n")

}

func calcularSigno(mes int, dia int) string {
	switch mes {
	case int(1):
		if dia < 20 {
			return "Capricornio"
		}
		return "Acuario"
	case int(2):
		if dia < 19 {
			return "Acuario"
		}
		return "Piscis"
	case int(3):
		if dia < 21 {
			return "Piscis"
		}
		return "Aries"
	case int(4):
		if dia < 20 {
			return "Aries"
		}
		return "Tauro"
	case int(5):
		if dia < 21 {
			return "Tauro"
		}
		return "Geminis"
	case int(6):
		if dia < 21 {
			return "Geminis"
		}
		return "Cancer"
	case int(7):
		if dia < 23 {
			return "Cancer"
		}
		return "Leo"

	case int(8):
		if dia < 23 {
			return "Leo"
		}
		return "Virgo"
	case int(9):
		if dia < 23 {
			return "Virgo"
		}
		return "Libra"
	case int(10):
		if dia < 23 {
			return "Libra"
		}
		return "Escorpio"
	case int(11):
		if dia < 22 {
			return "Escorpio"
		}
		return "Sagitario"
	case int(12):
		if dia < 22 {
			return "Sagitario"
		}
		return "Capricornio"
	}
	return "error, intenta de nuevo!"
}
