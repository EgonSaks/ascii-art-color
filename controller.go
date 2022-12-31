package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type letter struct {
	Letter string
	Ascii  int
	Table  []string
}

func mapFont(font string) []letter {
	readFile, err := os.ReadFile(font)
	if err != nil {
		fmt.Printf("could not read the content in the file: %v", err)
	}
	slice := strings.Split(string(readFile), "\n")

	letters := make([]letter, 0)
	index := 31

	for _, row := range slice {

		if row == "" {
			index++
			continue
		}

		grid := make([]string, 0)

		for _, char := range row {
			grid = append(grid, string(char))
		}

		letters = append(letters, letter{
			Letter: string(index),
			Ascii:  index,
			Table:  grid,
		})
	}
	return letters
}

func mapInput(input string, letters []letter) []letter {
	lines := strings.Split(input, "\\n")
	output := make([]letter, 0)

	for _, line := range lines {
		characters := []rune(line)

		if line != "" {
			for _, ch := range characters {
				if ch >= 32 && ch <= 126 {
					for _, v := range letters {
						if rune(v.Ascii) == ch {
							output = append(output, v)
						}
					}
				} else {
					fmt.Println("\ninput includes non ascii character(s), please use ascii character(s)")
					os.Exit(0)
				}
			}
		} else {
			output = append(output, letter{})
		}
	}
	return output
}

func printInput(str string, letter []letter, args []string) []string {
	var pixels []string

	for row := 0; row < 8; row++ {
		pixel := ""

		for letterIndex := 0; letterIndex < len(str); letterIndex++ {
			letter := mapInput(string(str[letterIndex]), letter)

			for column := 0; column < len(letter[row].Table); column++ {
				coloredPixel := color(letter[row].Table[column])

				// Check if the arguments count is between the range, same or equal to length of word
				if len(args[1:]) <= len(str) {

					// Check if nr of arguments is bigger or equal to 2 including word and not more than word
					if (len(args) >= 2) && (len(args) <= len(str)+1) {

						// Check if the second argument is a single letter
						if len(args[1]) == 1 {

							if len(args[1]) == 1 {

								// Check if the argument is :
								if args[1] == ":" {
									pixel = pixel + coloredPixel
								} else {
									matches := false
									// Iterate over the arguments and check if letter matches the argument

									for _, arg := range args[1:] {
										if letter[row].Letter == arg {
											matches = true
											break
										}
									}
									// If the current letter matches any of the arguments, color the letter
									if matches {
										pixel = pixel + coloredPixel
									} else {
										pixel = pixel + letter[row].Table[column]
									}
								}
							}

						} else if len(args[1]) == 2 {

							rangeParts := strings.Split(args[1], "")

							if rangeParts[0] != ":" {
								// Parse the start argument, e.g. n:
								startIndex, _ := strconv.Atoi(rangeParts[0])
								// Check if the current letter index is within the specified range
								if startIndex <= len(str) {
									if letterIndex >= startIndex-1 {
										pixel = pixel + coloredPixel
									} else {
										pixel = pixel + letter[row].Table[column]
									}
								} else {
									fmt.Println("\nStarting position is bigger than length of the word")
									os.Exit(0)
								}
							} else {
								// Parse the end argument, e.g. :n
								endIndex, _ := strconv.Atoi(rangeParts[1])
								// Check if the current letter index is within the specified range
								if endIndex <= len(str)  {
									if letterIndex <= endIndex-1 {
										pixel = pixel + coloredPixel
									} else {
										pixel = pixel + letter[row].Table[column]
									}
								} else {
									fmt.Println("\nEnding position is bigger than length of the word")
									os.Exit(0)
								}
							}

						} else {

							// Parse the start and end indices from the argument
							rangeParts := strings.Split(args[1], ":")
							startIndex, _ := strconv.Atoi(rangeParts[0])
							endIndex, _ := strconv.Atoi(rangeParts[1])

							// Check if the current letter index is within the specified range
							if (startIndex < endIndex) && (startIndex <= len(str)) && (endIndex <= len(str)) || startIndex == endIndex {
								if (letterIndex >= startIndex-1) && (letterIndex <= endIndex-1) {
									pixel = pixel + coloredPixel
								} else {
									pixel = pixel + letter[row].Table[column]
								}
							} else {
								fmt.Println("\nStarting position has to be smaller than ending position.\nStarting and/or ending positions can't be bigger than length of the word.\ne.g. 2:3")
								os.Exit(0)
							}

						}

					} else {

						// If there is no second argument, add the colored pixel to the string
						pixel = pixel + coloredPixel
					}

				} else {
					fmt.Println("\nArguments count is greater than length of the word.")
					os.Exit(0)
				}

			}

		}
		pixels = append(pixels, pixel)

	}

	return pixels
}

func colorPrint(output []string) {
	for _, v := range output {
		fmt.Println(string(v))
	}
}
