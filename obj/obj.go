// Package obj allows simple parsing of Wavefront OBJ files
// into slices of vertices, normals and elements.
package obj

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Parse a vertex (normal) string, a list of whitespace-separated
// floating point numbers.
func parseVertex(t []string) []float32 {
	x, _ := strconv.ParseFloat(t[0], 32)
	y, _ := strconv.ParseFloat(t[1], 32)
	z, _ := strconv.ParseFloat(t[2], 32)

	return []float32{float32(x), float32(y), float32(z)}
}

// Parse an element string, a list of whitespace-separated elements.
// Elements are of the form "<vi>/<ti>/<ni>" where indices are the
// vertex, texture coordinate and normal, respectively.
func parseElement(t []string) []int32 {
	e := make([]int32, len(t))

	for i := 0; i < len(t); i++ {
		f := strings.Split(t[i], "/")
		// for now, just grab the vertex index
		x, _ := strconv.ParseInt(f[0], 10, 32)
		e[i] = int32(x) - 1 // convert to 0-indexing
	}

	// convert quads to triangles
	if len(t) > 3 {
		e = append(e, e[0], e[2])
	}

	return e
}

func Parse(filename string) ([]float32, []int32) {
	fp, _ := os.Open(filename)
	scanner := bufio.NewScanner(fp)

	vertices := []float32{}
	normals := []float32{}
	elements := []int32{}

	for scanner.Scan() {
		toks := strings.Fields(strings.TrimSpace(scanner.Text()))

		switch toks[0] {
		case "v":
			vertices = append(vertices, parseVertex(toks[1:])...)
		case "vn":
			normals = append(normals, parseVertex(toks[1:])...)
		case "f":
			elements = append(elements, parseElement(toks[1:])...)
		}
	}

	return vertices, elements
}
