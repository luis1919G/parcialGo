package main

import (
	"fmt"

	"github.com/luis1919G/parcialGo/parcial/internal/tickets"
)

const (
	brasil   = "Brazil"
	china    = "China"
	mongolia = "Mongolia"

	madrugada = "madrugada"
	manana    = "mañana"
	tarde     = "tarde"
	noche     = "noche"
)

func main() {
	tickets.GetData()

	// *** Requerimiento 1 ***

	total, err := tickets.GetTotalTickets(brasil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)

	// *** Requerimiento 2 ***

	totalMadrugada, err := tickets.TotalTimeSlot(madrugada)
	if err != nil {
		fmt.Println(err)
	}
	tMadrugada, err := totalMadrugada()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tMadrugada)

	totalManana, err := tickets.TotalTimeSlot(manana)
	if err != nil {
		fmt.Println(err)
	}
	tManana, err := totalManana()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tManana)

	totalTarde, err := tickets.TotalTimeSlot(tarde)
	if err != nil {
		fmt.Println(err)
	}
	tTarde, err := totalTarde()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tTarde)

	totalNoche, err := tickets.TotalTimeSlot(noche)
	if err != nil {
		fmt.Println(err)
	}
	tNoche, err := totalNoche()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tNoche)

	// *** Requerimiento 3 ***

	average, err := tickets.AverageDestination(brasil, -10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("En promedio los viajes diarios a %s es de %0.2f ", brasil, average)
}
