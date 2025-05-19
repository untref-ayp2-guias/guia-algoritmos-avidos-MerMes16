package ejercicios

var billetes = []int{10000, 2000, 1000, 500, 200, 100, 50, 20, 10, 5, 2, 1}

func Cambiar(cantidad int) map[int]int {
	cambio := make(map[int]int)
	for _, denominacion := range billetes {
		if cantidad >= denominacion {

			cantidadBilletes := cantidad / denominacion
			cambio[denominacion] = cantidadBilletes
			cantidad %= denominacion

			if cantidad > 0 {
				for k, v := range Cambiar(cantidad) {
					cambio[k] += v
				}
			}
		}
	}
	return cambio
}
