package functions

import (
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"
)

func EncryptDesEncrypt(origin string) {
	fmt.Println("Original:", origin)
	crypt := Crypt(origin)
	fmt.Println("Criptografada:", crypt)
	decrypt := Decrypt(crypt)
	fmt.Println("Descriptografada:", decrypt)
	fmt.Printf("A descriptografia ocorreu ok? %v", origin == decrypt)
}
func Crypt(text string) string {
	l := newListCrypt(text)
	s := len(l) - 1
	f := -1
	var c []string
	if len(l)%2 == 0 {
		s = len(l) / 2
		f = 1
	}
	for i := 0; i < len(l); i++ {
		c = append(c, l[s])
		s = (s + f) % len(l)
	}
	return strings.Join(c, "")
}
func newListCrypt(text string) []string {
	t2 := strings.ToUpper(text)
	var l []string
	for i := 0; i < len(t2); i++ {
		upper := string(text[i]) == strings.ToUpper(string(text[i]))
		if (i+int(t2[i]))%2 == 0 {
			switch int(t2[i]) {
			case 65:
				if upper {
					l = append(l, string(rune(85)))
				} else {
					l = append(l, string(rune(117)))
				}
			case 69:
				if upper {
					l = append(l, string(rune(79)))
				} else {
					l = append(l, string(rune(111)))
				}
			case 73:
				if upper {
					l = append(l, string(rune(73)))
				} else {
					l = append(l, string(rune(105)))
				}
			case 79:
				if upper {
					l = append(l, string(rune(69)))
				} else {
					l = append(l, string(rune(101)))
				}
			case 85:
				if upper {
					l = append(l, string(rune(65)))
				} else {
					l = append(l, string(rune(97)))
				}
			default:
				if !upper {
					l = append(l, strings.ToLower(string(t2[i])))
				} else {
					l = append(l, string(t2[i]))
				}
			}
		} else {
			if !upper {
				l = append(l, strings.ToLower(string(t2[i])))
			} else {
				l = append(l, string(t2[i]))
			}
		}
	}
	return l
}
func Decrypt(text string) string {
	s := len(text) - 1
	f := -1
	if len(text)%2 == 0 {
		s = len(text) / 2
		f = 1
	}
	var nl []string
	for i := 0; i < len(text); i++ {
		nl = append(nl, string(text[s]))
		s = (s + f) % len(text)
	}
	nl = newListCrypt(strings.Join(nl, ""))
	return strings.Join(nl, "")
}
func HorseRide(input string) {
	baseString := "d4 f5 d6 e8 c7 a8 b6 a4 b2 d1 f2 h1 g3 h5 g7 e6 f8 d7 b8 a6 b4 a2 c1 e2 g1 h3 f4 d3 c5 e4 c3 d5 e3 c4 e5 c6 d8 b7 a5 b3 a1 c2 e1 g2 h4 g6 h8 f7 h6 g4 h2 f1 d2 b1 a3 b5 a7 c8 e7 g8 f6 h7 g5 f3"
	baseSplit := strings.Split(baseString, " ")
	size := len(baseSplit)
	ponto := ""
	start := 0
	for i := 0; i < size; i++ {
		ponto = baseSplit[i]
		if strings.EqualFold(input, ponto) {
			start = i
			break
		}
		if i == 63 {
			log.Fatal("Coordenada inexistente")
		}
	}
	matriz := NovaMatriz("__")
	for i := 1; i <= size; i++ {
		time.Sleep(time.Millisecond * 800)
		x, y := Coordinate(baseSplit[start])
		matriz[x][y] = strconv.Itoa(i)
		fmt.Print("\n\n")
		ShowMatriz(matriz)
		start = (start + 1) % size
	}
}
func ShowMatriz(matriz [][]string) {
	letra := "A"[0]
	fmt.Printf("   ")
	for c := 0; c < len(matriz); c++ {
		fmt.Printf("% 2d ", c+1)
	}
	fmt.Println()
	for i := 0; i < len(matriz); i++ {
		fmt.Printf("%v  ", string(letra))
		letra++
		for j := 0; j < len(matriz[i]); j++ {
			fmt.Printf("%02v ", matriz[i][j])
		}
		fmt.Println()
	}
}
func NovaMatriz(element string) [][]string {
	var m [][]string
	for i := 0; i < 8; i++ {
		var n []string
		for j := 0; j < 8; j++ {
			n = append(n, element)
		}
		m = append(m, n)
	}
	return m
}
func Coordinate(coordinate string) (int, int) {
	coordinate = strings.ToUpper(coordinate)
	linha := int(coordinate[0] - 65)
	var coll, _ = strconv.Atoi(string(coordinate[1]))
	return linha, coll - 1
}
func FactInteractive(n int) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("Invalid number")
	}
	b := big.NewInt(1)
	for i := 1; i <= n; i++ {
		b.Mul(b, big.NewInt(int64(i)))
	}
	return b, nil
}
func FactoriaRestrito(n int) uint64 {
	var a uint64 = 1
	for i := 1; i <= n; i++ {
		a *= uint64(i)
	}
	return a

}
func FibonacciRecursivo(n int) int64 {
	if n > 1 {
		return FibonacciRecursivo(n-1) + FibonacciRecursivo(n-2)
	}
	return int64(n)
}
func FibonacciIterativo(n int) int64 {
	var a, b int64 = 0, 1
	for i := 1; i < n; i++ {
		b += a
		a = b - a
	}
	return b
}
