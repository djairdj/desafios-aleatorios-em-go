package functions

import (
	fn "aleatorios/functions"
)

func Desafios(option int) {
	switch option {
	case 1:
		fn.EncryptDesEncrypt()
	case 2:
		fn.HorseRide()
	default:
		println("Nenhum desafio numerado com", option)
	}
}
