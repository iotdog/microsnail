package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// diff_binary_0.dat
	f, err := os.Open("diff_binary_0.dat")
	check(err)
	defer f.Close()
	tmp := make([]byte, 4)
	_, err = f.Read(tmp)
	check(err)
	pointNumber := binary.LittleEndian.Uint32(tmp)
	fmt.Printf("get point number: %d\n", pointNumber)
	_, err = f.Read(tmp)
	pointDimension := binary.LittleEndian.Uint32(tmp)
	fmt.Printf("get point dimension: %d\n", pointDimension)
	dataSize := pointNumber * pointDimension
	data := make([]float32, dataSize)
	var i uint32
	for i = 0; i < dataSize; i++ {
		tmp = make([]byte, 4)
		_, err = f.Read(tmp)
		data[i] = math.Float32frombits(binary.LittleEndian.Uint32(tmp))
		fmt.Printf("get point value: %1.10f\n", data[i])
	}
}
