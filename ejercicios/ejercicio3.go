package ejercicios

import "sort"

type Objeto struct {
	Nombre string
	Peso   int
	Valor  int
}

func (o *Objeto) ValorPorUnidadDePeso() float32 {
	return float32(o.Valor) / float32(o.Peso)
}

func Ejercicio3(objetos []Objeto, capacidad int) map[Objeto]float32 {
	sort.Slice(objetos, func(i, j int) bool {
		return objetos[i].ValorPorUnidadDePeso() > objetos[j].ValorPorUnidadDePeso()
	})

	res, cap := make(map[Objeto]float32), 0
	for _, obj := range objetos {
		if cap+obj.Peso <= capacidad {
			res[obj] = 1
			cap += obj.Peso
		} else {
			pesoRestante := capacidad - cap
			res[obj] = float32(pesoRestante) / float32(obj.Peso)
			break
		}
	}
	return res
}
