package main

import (
	"fmt"
)

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista( tl *task) {
	t.tasks = append(t.tasks, tl)
}

func ( t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index + 1:]...)
}

func (t *taskList) imprimirLista(){
	for _, tarea := range t.tasks {
		fmt.Println("Nombre:", tarea.nombre)
		fmt.Println("Descripcion:", tarea.descripcion)
		fmt.Println("Completado:", tarea.completado)
	}
}

func(t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre Tarea:", tarea.nombre)
			fmt.Println("Descripcion Tarea:", tarea.descripcion)
		}
	}
}

type task struct{
	nombre string
	descripcion string
	completado bool
}

func (t *task) marcarCompleta(){
	t.completado = true
}

func (t *task) actualizaDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizaNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre: 	 "Completar mi curso de Go::",
		descripcion: "Completare mi curso de Go de Platzi en esta semana::",
	}

	t2 := &task{
		nombre: 	 "Completar mi curso de JavaScript->",
		descripcion: "Completare mi curso de JavaScript de Platzi en esta semana->",
	}

	t3 := &task{
		nombre: 	 "Completar mi curso de NodeJS..",
		descripcion: "Completare mi curso de NodeJS de Platzi en esta semana..",
	}

	lista := &taskList {
		tasks: []*task{
			t1, t2,
		},
	}

	lista.agregarALista(t3)
	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("*** tareas completadas ***")
	lista.imprimirListaCompletados()

/* 	for i:= 0; i < len(lista.tasks); i++ {
		fmt.Println("index->", i, "nombre->", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks {
		fmt.Println("index:", index, "nombre:", tarea.nombre)
	}

	for i:= 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("contador:", i)
	} */
}