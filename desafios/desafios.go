package functions

import (
	fn "aleatorios/functions"
	"fmt"
	"os"
	"strconv"
)

func Desafios(option int) {

	switch option {
	case 1:
		origin := "lata amarelinha"
		if len(os.Args) > 1 {
			origin = os.Args[1]
		}
		fn.EncryptDesEncrypt(origin)
	case 2:
		input := "b7"
		if len(os.Args) > 1 {
			input = os.Args[1]
		}
		fn.HorseRide(input)
	case 3:
		input := "45"
		if len(os.Args) > 1 {
			input = os.Args[1]
		}
		n, _ := strconv.Atoi(input)
		fibIte := fn.FibonacciIterativo(n)
		fmt.Printf("O %vº número de fibonacci iterativo é %v\n", n, fibIte)
		fibRec := fn.FibonacciRecursivo(n)
		fmt.Printf("O %vº número de fibonacci recursivo é %v\n", n, fibRec)
	case 4:
		input := "20"
		if len(os.Args) > 1 {
			input = os.Args[1]
		}
		f, _ := strconv.Atoi(input)
		fact, err := fn.FactInteractive(f)
		if err != nil {
			fmt.Printf(err.Error())
		}
		//x := fn.FactoriaRestrito(f)
		fmt.Printf("O fatorial de %v é %v\n", f, fact)
		//fmt.Printf("O fatorial de %v é %v\n", f, x)
	default:
		println("Nenhum desafio numerado com", option)
	}
}
