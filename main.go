package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	filename := "ejemplo.bin"
	numBytes := 223 // Bytes aleatorios a escribir

	// Creando archivo
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creando archivo:", err)
		return
	}
	defer file.Close()

	// Generando bytes aleatorios
	randomBytes := make([]byte, numBytes)
	_, err = rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generando bytes aleatorios:", err)
		return
	}

	// Escribiendo bytes aleatorios al archivo
	_, err = file.Write(randomBytes)
	if err != nil {
		fmt.Println("Error escribiendo al archivo:", err)
		return
	}

	fmt.Printf("Escritos exitosamente %d bytes aleatorios a %s\n", numBytes, filename)
}
