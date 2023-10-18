package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre   string   `json:"Nombre"`
	Apellido string   `json:"Apellido"`
	Edad     int      `json:"Edad"`
	Dichos   []string `json:"Dichos"`
}

func main() {
	s := `[{"Nombre":"Eduar","Apellido":"Tua","Edad":32,"Dichos":["Cachicamo diciéndole a morrocoy conchudo","La mona, aunque se vista de seda, mona se queda","Poco a poco se anda lejos"]},{"Nombre":"Carlos","Apellido":"Pérez","Edad":27,"Dichos":["Camarón que se duerme se lo lleva la corriente","A ponerse las alpargatas que lo que viene es joropo","No gastes pólvora en zamuro"]},{"Nombre":"M","Apellido":"Hmmmm","Edad":54,"Dichos":["Ni lava ni presta la batea","Hijo de gato, caza ratón","Más vale pájaro en mano que cien volando"]}]`
	jsonBytes := []byte(s)
	p1 := []Persona{}
	err := json.Unmarshal(jsonBytes, &p1)

	if err != nil {
		fmt.Println(err)
	}

	for _, el := range p1 {
		fmt.Println(el.Nombre)
		fmt.Println(el.Apellido)
		fmt.Println(el.Edad)
		fmt.Println("Dichos:")
		for _, dicho := range el.Dichos {
			fmt.Println(dicho)
		}
	}

}
