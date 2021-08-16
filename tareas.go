package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) agregarALista(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) eliminarDeLista(i int) {
	tl.tasks = append(tl.tasks[:i], tl.tasks[i+1:]...)
}

func (tl taskList) imprimirLista() {
	for _, task := range tl.tasks {
		fmt.Println("Nombre:", task.nombre)
		fmt.Println("Descripcion:", task.descripcion)
	}
}

func (tl taskList) imprimirCompletadas() {
	for _, task := range tl.tasks {
		if task.completado {
			fmt.Println("Nombre:", task.nombre)
			fmt.Println("Descripcion:", task.descripcion)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(description string) {
	t.descripcion = description
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de Platzi en esta semana",
	}

	t2 := task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de Platzi en esta semana",
	}

	t3 := task{
		nombre:      "Completar mi curso de NodeJS",
		descripcion: "Completar mi curso de NodeJS de Platzi en esta semana",
	}

	lista := taskList{
		tasks: []*task{
			&t1, &t2,
		},
	}
	lista.agregarALista(&t3)

	lista.imprimirLista()
	lista.imprimirCompletadas()

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i, "nombre", lista.tasks[i].nombre)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
