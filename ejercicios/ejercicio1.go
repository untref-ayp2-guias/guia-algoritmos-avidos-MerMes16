package ejercicios

type Actividad struct {
	Nombre string
	Inicio int
	Fin    int
}

func SelectorRecursivo(actividades []Actividad) []Actividad {
	seleccionadas := []Actividad{}
	if len(actividades) == 0 {
		return seleccionadas
	}
	actual, seleccionadas := actividades[0], append(seleccionadas, actividades[0])
	for i, actividad := range actividades {
		if actividad.Inicio > actual.Fin {
			return append(seleccionadas, SelectorRecursivo(actividades[i:])...)
		}
	}
	return seleccionadas
}
