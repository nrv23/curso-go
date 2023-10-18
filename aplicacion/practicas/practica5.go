package main

import (
	"fmt"
	"sort"
)

type usuario struct {
	Nombre   string
	Apellido string
	Edad     int
	Dichos   []string
}

type ByAge []usuario
type ByName []usuario
type ByComments []usuario

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

// ordenar por nombre

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Nombre < a[j].Nombre }

func main() {
	u1 := usuario{
		Nombre:   "Eduar",
		Apellido: "Tua",
		Edad:     32,
		Dichos: []string{
			"Cachicamo diciéndole a morrocoy conchudo",
			"La mona, aunque se vista de seda, mona se queda",
			"Poco a poco se anda lejos",
		},
	}

	u2 := usuario{
		Nombre:   "Carlos",
		Apellido: "Pérez",
		Edad:     27,
		Dichos: []string{
			"Camarón que se duerme se lo lleva la corriente",
			"A ponerse las alpargatas que lo que viene es joropo",
			"No gastes pólvora en zamuro",
		},
	}

	u3 := usuario{
		Nombre:   "Che",
		Apellido: "López",
		Edad:     54,
		Dichos: []string{
			"Ni lava ni presta la batea",
			"Hijo de gato, caza ratón",
			"Más vale pájaro en mano que cien volando",
		},
	}

	usuarios := []usuario{u1, u2, u3}

	// ordenar los dichos del usuario
	for _, v := range usuarios {
		fmt.Println(v.Nombre, v.Apellido, v.Edad)
		sort.Strings(v.Dichos)
		for _, v := range v.Dichos {
			fmt.Println("\t", v)
		}
	}

	fmt.Println(usuarios)
	// ordenar por nopmbrec
	sort.Sort(ByName(usuarios))
	fmt.Println(usuarios)

	//
	fmt.Println("")

	// ordenado por edad
	sort.Sort(ByAge(usuarios))
	fmt.Println(usuarios)

}
