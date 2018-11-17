package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {

	rayFile, err := os.Create("ray.ppm")
	if err != nil {
		log.Fatal("failed to create ray.ppm file, error: ", err)
	}

	defer rayFile.Close()

	var x, y int = 200, 100

	w := bufio.NewWriter(rayFile)

	w.WriteString("P3\n")
	w.WriteString(strconv.Itoa(x))
	w.WriteString(" ")
	w.WriteString(strconv.Itoa(y))
	w.WriteString("\n255\n")

	for i := y - 1; i >= 0; i-- {
		for j := 0; j < x; j++ {
			r := float32(j) / float32(x)
			g := float32(i) / float32(y)
			b := 0.2

			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)

			w.WriteString(strconv.Itoa(ir))
			w.WriteString(" ")
			w.WriteString(strconv.Itoa(ig))
			w.WriteString(" ")
			w.WriteString(strconv.Itoa(ib))
			w.WriteString("\n")
		}
	}

	w.Flush()
}
