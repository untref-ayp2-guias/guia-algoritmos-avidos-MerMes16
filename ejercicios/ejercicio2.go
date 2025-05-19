package ejercicios

type Tarea struct {
	Nombre string
	Tiempo int
}

func Ejercicio2(tareas []Tarea) {
	if len(tareas) < 2 {
		return
	}
	for i := 1; i < len(tareas)-1; i++ {
		for j := 0; j < len(tareas)-i; j++ {
			if tareas[j].Tiempo > tareas[j+1].Tiempo {
				tareas[j], tareas[j+1] = tareas[j+1], tareas[j]
			}
		}
	}
}
