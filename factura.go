package main

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
	Total    float64
}

func Factura() {
	// Ingresar los datos del cliente
	var NombreCliente string
	var Productos []*Producto

	fmt.Print("Ingrese el nombre del cliente: ")
	fmt.Scanln(&NombreCliente)

	Fecha := time.Now().Format("02/01/2006") // Formato de fecha día/mes/año

	for {
		var producto Producto
		fmt.Print("Ingrese el nombre del producto (o 'fin' para terminar): ")
		fmt.Scanln(&producto.Nombre)
		if producto.Nombre == "fin" {
			break
		}
		fmt.Print("Ingrese el precio unitario: ")
		fmt.Scanln(&producto.Precio)
		fmt.Print("Ingrese la cantidad: ")
		fmt.Scanln(&producto.Cantidad)
		producto.Total = producto.Precio * float64(producto.Cantidad)
		Productos = append(Productos, &producto)

	}

	// Crear el PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Título de la factura
	pdf.Cell(190, 10, "Factura") // Centrado
	pdf.Ln(10)

	// Información del cliente y fecha
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(100, 10, fmt.Sprintf("Cliente: %s", NombreCliente))
	pdf.Cell(90, 10, fmt.Sprintf("Fecha: %s", Fecha))
	pdf.Ln(15)

	// Encabezado de la tabla con bordes
	pdf.SetFont("Arial", "B", 12)
	pdf.SetFillColor(200, 200, 200)
	pdf.CellFormat(50, 10, "Producto", "1", 0, "C", false, 0, "")
	pdf.CellFormat(40, 10, "Precio Unitario", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Cantidad", "1", 0, "C", false, 0, "")
	pdf.CellFormat(30, 10, "Total", "1", 1, "C", false, 0, "") // '1' al final hace que avance de línea

	// Listado de productos con bordes
	pdf.SetFont("Arial", "", 12)
	var ValorTotal float64
	fmt.Println("\nResumen dela factura:")
	fmt.Printf("Cliente: %s\n", NombreCliente)
	fmt.Printf("Fecha: %s\n", Fecha)
	fmt.Println("Producto:",)
	for _, producto := range Productos {
		pdf.CellFormat(50, 10, producto.Nombre, "1", 0, "L", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("%.2f", producto.Precio), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%d", producto.Cantidad), "1", 0, "C", false, 0, "")
		pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", producto.Total), "1", 1, "C", false, 0, "") // Nueva línea
		ValorTotal += producto.Total
	}

	// Total de la factura
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(120, 10, "Valor Total", "1", 0, "R", false, 0, "")
	pdf.CellFormat(30, 10, fmt.Sprintf("%.2f", ValorTotal), "1", 1, "C", false, 0, "")

	// Guardar el PDF
	err := pdf.OutputFileAndClose("Factura.pdf")
	if err != nil {
		fmt.Println("Error al crear el PDF", err)
	} else {
		fmt.Println("Factura.pdf generada exitosamente!")
	}

}

func main() {
	Factura()
}
