package main

import (
	"bufio"
	"dithering/colorspace"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	multipliers := [4]int{180, 220, 255, 135}

	file, err := os.Open("colors.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {

		baseR, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		baseG, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		baseB, _ := strconv.Atoi(scanner.Text())

		for _, value := range multipliers {
			r := baseR * value / 255
			g := baseG * value / 255
			b := baseB * value / 255
			col := colorspace.StandardRGB{r, g, b}
			fmt.Println(col)
		}

	}

}
