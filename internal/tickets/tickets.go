package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	madrugada = "madrugada"
	manana    = "mañana"
	tarde     = "tarde"
	noche     = "noche"
)

type Ticket struct {
	id                 string
	name               string
	email              string
	destinationCountry string
	flightTime         string
	price              string
}

var tickets []Ticket

// ejemplo 1
func GetTotalTickets(destination string) (int, error) {

	err := errors.New("Se ingresó un una cadena vacía")
	if destination == "" {
		return 0, err
	}
	var count int
	for _, ticket := range tickets {
		if ticket.destinationCountry == destination {
			count++
		}
	}
	return count, nil
}

// // ejemplo 2
func GetEarlyMornings() (int, error) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var count int
	for _, ticket := range tickets {
		flightTime := strings.Split(ticket.flightTime, ":")
		hour, _ := strconv.Atoi(flightTime[0])
		if hour >= 0 && hour <= 6 {
			count++
		}
	}
	return count, nil
}

func GetMornings() (int, error) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var count int
	for _, ticket := range tickets {
		flightTime := strings.Split(ticket.flightTime, ":")
		hour, _ := strconv.Atoi(flightTime[0])
		if hour >= 7 && hour <= 12 {
			count++
		}
	}
	return count, nil
}
func GetAfternoons() (int, error) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var count int
	for _, ticket := range tickets {
		flightTime := strings.Split(ticket.flightTime, ":")
		hour, _ := strconv.Atoi(flightTime[0])
		if hour >= 13 && hour <= 19 {
			count++
		}
	}
	return count, nil
}
func GetNights() (int, error) {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var count int
	for _, ticket := range tickets {
		flightTime := strings.Split(ticket.flightTime, ":")
		hour, _ := strconv.Atoi(flightTime[0])
		if hour >= 20 && hour <= 23 {
			count++
		}
	}
	return count, nil
}

func TotalTimeSlot(timeSlot string) (func() (int, error), error) {
	switch timeSlot {
	case madrugada:
		return GetEarlyMornings, nil

	case manana:
		return GetMornings, nil

	case tarde:
		return GetAfternoons, nil

	case noche:
		return GetNights, nil
	default:
		return nil, errors.New("La opción ingresada no es válida")
	}
}

// // ejemplo 3
func AverageDestination(destination string, total int) (float64, error) {

	var totalFlights float64 = float64(total * 1.0)
	var average float64
	totalTickets, err := GetTotalTickets(destination)
	if err != nil {
		return 0, err
	}
	if total < 0 {
		return 0, errors.New("El total de pasajes diarios debe ser positivo")
	}
	average = float64(totalTickets) / totalFlights
	return average, nil
}

func GetData() {
	file, err := os.Open("tickets.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		user, err := reader.Read()
		if err != nil {
			break
		}
		usr := Ticket{
			id:                 user[0],
			name:               user[1],
			email:              user[2],
			destinationCountry: user[3],
			flightTime:         user[4],
			price:              user[5],
		}
		tickets = append(tickets, usr)
	}
}
