package main

import (
	"fmt"
	"time"
	"math"
)

// Definimos la estructura
type Prestaciones struct {
	Nombre       string
	FechaEntrada time.Time
	FechaSalida  time.Time
	Sueldo       float64
	SueldoDiario float64
}

// Constructor equivalente
func NewPrestaciones(nombre string, fechaEntrada, fechaSalida string, sueldo float64) Prestaciones {
	layout := "02-01-2006"
	entrada, _ := time.Parse(layout, fechaEntrada)
	salida, _ := time.Parse(layout, fechaSalida)
	sueldodiario := math.Round((sueldo/23.83)*100) / 100 // redondeo a 2 decimales

	return Prestaciones{
		Nombre:       nombre,
		FechaEntrada: entrada,
		FechaSalida:  salida,
		Sueldo:       sueldo,
		SueldoDiario: sueldodiario,
	}
}

// Calcular Preaviso
func (p Prestaciones) CalcularPreaviso() float64 {
	preaviso := 0.0

	anos := p.FechaSalida.Year() - p.FechaEntrada.Year()
	meses := int(p.FechaSalida.Month() - p.FechaEntrada.Month() + time.Month(anos*12))
	if p.FechaSalida.Day() < p.FechaEntrada.Day() {
    	meses--
	}

	ceroATresMeses := 0
	tresASeisMeses := 7
	seisADoceMeses := 14
	mayorADoceMeses := 28

	if meses < 3 {
		preaviso = p.SueldoDiario * float64(ceroATresMeses)
	} else if meses > 3 && meses <= 6 {
		preaviso = p.SueldoDiario * float64(tresASeisMeses)
	} else if meses > 6 && meses < 12 {
		preaviso = p.SueldoDiario * float64(seisADoceMeses)
	} else if meses >= 12 {
		preaviso = p.SueldoDiario * float64(mayorADoceMeses)
	}

	return math.Round(preaviso*100) / 100
}

// Calcular Cesantía
func (p Prestaciones) CalcularCesantia() float64 {
	cesantia := 0.0

	anos := p.FechaSalida.Year() - p.FechaEntrada.Year()
	meses := int(p.FechaSalida.Month() - p.FechaEntrada.Month() + time.Month(anos*12))
	if p.FechaSalida.Day() < p.FechaEntrada.Day() {
    	meses--
	}

	ceroATresMeses := 0
	tresASeisMeses := 6
	seisADoceMeses := 13
	unA5Anos := 21 * anos
	mayorA5Anos := 23 * anos

	if meses < 3 {
		cesantia = p.SueldoDiario * float64(ceroATresMeses)
	} else if meses > 3 && meses <= 6 {
		cesantia = p.SueldoDiario * float64(tresASeisMeses)
	} else if meses > 6 && meses < 12 {
		cesantia = p.SueldoDiario * float64(seisADoceMeses)
	} else if meses >= 12 && anos <= 5 {
		cesantia = p.SueldoDiario * float64(unA5Anos)
	} else if anos > 5 {
		cesantia = p.SueldoDiario * float64(mayorA5Anos)
	}

	return math.Round(cesantia*100) / 100
}

// Calcular Vacaciones
func (p Prestaciones) CalcularVacaciones() float64 {
	vacaciones := 0.0
	anos := p.FechaSalida.Year() - p.FechaEntrada.Year()
	meses := p.FechaSalida.Month() - p.FechaEntrada.Month()

	mascincomeses := 6
	masdeseis := 7
	masdesiete := 8
	masdeocho := 9
	masdenueve := 10
	masdediez := 11
	masdeonce := 12

	unoA5Anos :=  14
	mayorA5Anos := 18
	
	if meses == 5{
		vacaciones = p.SueldoDiario * float64(mascincomeses)
	} else if meses == 6{
		vacaciones = p.SueldoDiario * float64(masdeseis)
	} else if meses == 7{
		vacaciones = p.SueldoDiario * float64(masdesiete)
	} else if meses == 8{
		vacaciones = p.SueldoDiario * float64(masdeocho)
	} else if meses == 9{
		vacaciones = p.SueldoDiario * float64(masdenueve)
	} else if meses == 10{
		vacaciones = p.SueldoDiario * float64(masdediez)
	} else if meses == 11{
		vacaciones = p.SueldoDiario * float64(masdeonce)
	} else if anos <= 5 {
		vacaciones = p.SueldoDiario * float64(unoA5Anos)
	} else {
		vacaciones = p.SueldoDiario * float64(mayorA5Anos)
	}

	return math.Round(vacaciones*100) / 100
}

// Calcular Regalia
func (p Prestaciones) CalculoRegalia() float64 {
	var ano float64 = 12
	regalia := 0.0
	meses := p.FechaSalida.Month() - p.FechaEntrada.Month()
	dias := float64(p.FechaSalida.Day() -p.FechaEntrada.Day())+1

	valorfloat := float64(dias) / 30


	mesesdias := float64(meses) + valorfloat
	fmt.Println(valorfloat)

	if meses <= 12 {
		regalia = (float64(mesesdias) * (p.Sueldo / ano)) 
	}

	return math.Round(regalia*100) / 100
}

func main() {
	var name string
	var fe string
	var fs string
	var sueldo float64
	fmt.Println("Escribe tu nombre:")
	fmt.Scanln(&name)
	fmt.Println("Escribe tu Fecha de entrada dd-mm-aaaa:")
	fmt.Scanln(&fe)
	fmt.Println("Escribe tu Fecha de salida dd-mm-aaaa:")
	fmt.Scanln(&fs)
	fmt.Println("Escribe tu sueldo:")
	fmt.Scanln(&sueldo)

	info := NewPrestaciones(name, fe, fs, sueldo)
	//info := NewPrestaciones("manuel", "01-01-2025", "26-10-2025", 10000)
	preaviso := info.CalcularPreaviso()
	cesantia := info.CalcularCesantia()
	vacaciones := info.CalcularVacaciones()
	regalia := info.CalculoRegalia()
	total := preaviso + cesantia + vacaciones + regalia

	fmt.Println("=================================================")
	fmt.Println("Informacion del empleado")
	fmt.Println("=================================================")
	fmt.Printf("Nombre: %s\n", info.Nombre)
	fmt.Printf("Sueldo diario: %.2f\n", info.SueldoDiario)
	fmt.Println("=================================================")
	fmt.Printf("Preaviso a pagar: %.2f\n", preaviso)
	fmt.Printf("Cesantia a pagar: %.2f\n", cesantia)
	fmt.Printf("Vacaciones a pagar: %.2f\n", vacaciones)
	fmt.Printf("Regalia a pagar: %.2f\n", regalia)
	fmt.Println("=================================================")
	fmt.Printf("Total a pagar: %.2f\n", total)
	fmt.Println("=================================================")
}
