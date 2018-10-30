package main

import (
	"fmt"
	"github.com/angus-g/go-obj/obj"
	"os"
)

func main() {
	fmt.Println("Test of obj library")
	vert, norm, err := obj.Parse("stall.obj")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Vertices:", len(vert))
	for i := 0; i < len(vert); i += 3 {
		fmt.Println("Vertex", i/3, vert[i], vert[i+1], vert[i+2])
	}

	fmt.Println("Normals:", len(norm))
	for i := 0; i < len(norm); i += 3 {
		fmt.Println("Normal", i/3, norm[i], norm[i+1], norm[i+2])
	}
}
