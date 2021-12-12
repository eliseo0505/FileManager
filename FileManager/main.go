package main

import (
	"fmt"
	"io/fs"
	"os"
)

var directory = "C:/Users/roxpl/Desktop/Mega/MEGAsync/Go/FileManager" // Directorio
var file = "Index.html"                                               // Archivo que se busca

func main() {
	openDir()
}

// 1. Leer Directorio.
// 2. Retorna un slice con los archivos(slFile).
// 3. Es llamada la siguiente función(sFileV).
func openDir() []fs.DirEntry {
	dir, err := os.ReadDir(directory)
	slFile := []string{}
	if err != nil {
		fmt.Printf("Ha ocurrido un error al intentar abrir el directorio %v", err)
	} else {
		for _, fD := range dir {
			slFile = append(slFile, fD.Name())
		}
		sFileV(slFile)
	}
	return dir
}

// 1. Agrupa los archivos que NO coinciden con la busqueda en un slice(sFileF).
// 2. Renombra una variable con el encontrado(sFileT).
// 3. Es llamada la siguiente función(resultSearch).
// 4. Retorna la variable con el archivo encontrado(sFileT).
func sFileV(slFile []string) string {
	sFileF := []string{}
	sFileT := ""
	for _, nF := range slFile {
		if file != nF {
			sFileF = append(sFileF, nF)
		} else {
			sFileT = file
		}
	}
	resultSearch(sFileT)
	return sFileT
}

// 1. Recorre el slice recibido.
// 2. Si el contenido de 'sFileT' es igual a la variable global 'file' llama la función readFile()
//    sino, llama la función FileNotExist()
func resultSearch(sFileT string) {
	if sFileT == file {
		fmt.Println("Archivo encontrado:")
		readFile()
	} else {
		fmt.Println("1. Archivo no encontrado...")
		fileNotExist()
	}
}

// 1. Crea el archivo
// 2. Llama a la función readFile()
func fileNotExist() {
	nF, err := os.Create(file)
	if err != nil {
		fmt.Println("No se ha podido crear el archivo.")
	} else {
		nF.Write([]byte("<html>\n<head>\n\n</head>\n<body>\n<h1>Hola Mundo!</h1>\n</body>\n</html>"))
		fmt.Println("2. Archivo creado:")
		readFile()

	}
	return
}

// 1. Se crean las variables para concatenar el directorio completo
// 2. Se leé el archivo y se lo imprime.
func readFile() {
	i := string(directory)
	o := "/"
	p := string(file)
	a := i + o + p
	index, err := os.ReadFile(a)
	if err != nil {
		fmt.Println("El archivo no pudo ser leido.")
	} else {
		fmt.Printf(string(index))
	}
}
