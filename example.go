package main

import (
	"fmt"
)

type vertex struct {
	x, y int
}

func main() {
	fmt.Printf("Variables Example results = %d\n", variables())
	pointers()
	slices()
	maps()
}

func variables() (variableResult int) {
	//Explicit type declartion
	var var1, var2, var3 int = 1, 2, 3

	//The type is implied
	var4, var5, var6 := 4, 5, 6

	//variableResult is a named result parameter
	variableResult = var1 + var2 + var3 + var4 + var5 + var6

	//Don't need to return anything since variableResult was assigned above.
	return
}

func pointers() {
	vertexOrignal := vertex{1, 2}

	//This is going to make a copy of the orginal.
	vertexCopy := vertexOrignal

	vertexCopy.x = 3
	vertexCopy.y = 4

	fmt.Println("Pointer Example")
	fmt.Printf("Copy variable changed Original x = %d y = %d\n", vertexOrignal.x, vertexOrignal.y)

	//This will give you a reference to the pointer
	vertexReference := &vertexOrignal

	vertexReference.x = 3
	vertexReference.y = 4

	fmt.Printf("Pointer variable changed Original x = %d y = %d\n", vertexOrignal.x, vertexOrignal.y)

	//Create a pointer
	vertexPointer := &vertex{5, 6}

	vertexPointerRef := vertexPointer
	vertexPointerRef.x++
	vertexPointerRef.y++

	//This will give you the value at the pointer
	vertexValue := *vertexPointer
	fmt.Printf("Vertex Pointer change x = %d y = %d\n", vertexValue.x, vertexValue.y)
	return
}

func slices() {
	sliceLength()
	sliceNil()
	mutatingSlices()
}

func sliceLength() {
	//Create array with length of five
	sliceWithLength := make([]vertex, 5)

	for i := 0; i < len(sliceWithLength); i++ {
		sliceWithLength[i] = vertex{i, i + 1}
	}
	fmt.Println("Slice Example")
	fmt.Printf("Slice with length = %v\n", sliceWithLength)
}

func sliceNil() {
	//Create a nil slice.
	var sliceWithNil []vertex

	for i := 0; i < 5; i++ {
		//Must append to array since the slice is nil.
		//This will be costly to performance.
		sliceWithNil = append(sliceWithNil, vertex{i, i + 1})
	}

	fmt.Printf("Slice nil = %v\n", sliceWithNil)
}

func mutatingSlices() {

	sliceVertices := make([]vertex, 5)

	for i := 0; i < len(sliceVertices); i++ {
		sliceVertices[i] = vertex{i, i + 1}
	}

	//Changing values won't change the slice
	changeValues(sliceVertices)
	fmt.Printf("After change values  = %v\n", sliceVertices)

	//Change the references will change the slice
	changeReferences(sliceVertices)
	fmt.Printf("After change references  = %v\n", sliceVertices)

	//Make a slice of pointers so you always have a reference
	slicePointers := make([]*vertex, 5)

	for i := 0; i < len(slicePointers); i++ {
		slicePointers[i] = &vertex{i, i + 1}
	}

	//Change the array of pointers.
	changePointers(slicePointers)

	//Create an array of values
	sliceValues := make([]vertex, 5)
	for i := 0; i < len(slicePointers); i++ {
		//Gets the value
		sliceValues[i] = *slicePointers[i]
	}

	fmt.Printf("After slice of pointers change  = %v\n", sliceValues)
}

func changeValues(vertices []vertex) {
	for _, vertex := range vertices {
		vertex.x++
		vertex.y++
	}

	return
}

func changeReferences(vertices []vertex) {
	for i := 0; i < len(vertices); i++ {
		vertex := &vertices[i]
		vertex.x++
		vertex.y++
	}

	return
}

func changePointers(vertices []*vertex) {
	for _, vertex := range vertices {
		vertex.x++
		vertex.y++
	}

	return
}

func maps() {
	mapValues := make(map[int]vertex)

	for i := 0; i < 5; i++ {
		mapValues[i] = vertex{i, i + 1}
	}

	fmt.Printf("Original Map  = %v\n", mapValues)

	mutatingMaps()
}

func mutatingMaps() {
	mapPointers := make(map[int]*vertex)
	for i := 0; i < 5; i++ {
		mapPointers[i] = &vertex{i, i + 1}
	}

	for i := 0; i < 5; i++ {
		mapPointers[i].x++
		mapPointers[i].y++
	}

	//Get the values from the map of pointers.
	mapValues := make(map[int]vertex)
	for i := 0; i < 5; i++ {
		mapValues[i] = *mapPointers[i]
	}

	fmt.Printf("Mutated Map  = %v\n", mapValues)

}
