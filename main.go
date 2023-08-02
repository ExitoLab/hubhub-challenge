package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calculateStrength([]string{"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8",
		"addx 13", "addx 4", "noop", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1",
		"addx 5", "addx -1", "addx 5", "addx -1", "addx -35", "addx 1", "addx 24", "addx -19", "addx 1", "addx 16", "addx -11", "noop", "noop", "addx 21", "addx -15",
		"noop", "noop", "addx -3", "addx 9", "addx 1", "addx -3", "addx 8", "addx 1", "addx 5", "noop", "noop", "noop", "noop", "noop", "addx -36", "noop", "addx 1", "addx 7", "noop", "noop", "noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop",
		"noop", "addx 1", "noop", "noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13",
		"addx 7", "noop", "addx 1", "addx -33", "noop", "noop", "noop", "addx 2", "noop",
		"noop", "noop", "addx 8", "noop", "addx -1", "addx 2", "addx 1", "noop", "addx 17",
		"addx -9", "addx 1", "addx 1", "addx -3", "addx 11", "noop", "noop", "addx 1",
		"noop", "addx 1", "noop", "noop", "addx -13", "addx -19", "addx 1", "addx 3", "addx 26", "addx -30", "addx 12", "addx -1", "addx 3", "addx 1", "noop", "noop", "noop",
		"addx -9", "addx 18", "addx 1", "addx 2", "noop", "noop", "addx 9", "noop", "noop", "noop", "addx -1", "addx 2", "addx -37", "addx 1", "addx 3", "noop", "addx 15", "addx -21", "addx 22", "addx -6", "addx 1", "noop", "addx 2", "addx 1", "noop", "addx -10", "noop", "noop", "addx 20", "addx 1", "addx 2", "addx 2", "addx -6", "addx -11", "noop", "noop", "noop"}))
}

func processSignal(instructions []string) map[int]int {
	cycle := 0
	regX := 1
	instructionIndex := 0
	cycleRegValues := map[int]int{}
	for instructionIndex < len(instructions) {
		if instructions[instructionIndex] == "noop" {
			for i := 0; i < 1; i++ {
				cycleRegValues[cycle] = regX
				cycle++
			}
			instructionIndex++
		} else if strings.HasPrefix(instructions[instructionIndex], "addx") {
			for i := 0; i < 2; i++ {
				cycleRegValues[cycle] = regX
				cycle++
			}
			val, _ := strconv.Atoi(strings.Split(instructions[instructionIndex], " ")[1])
			regX += val
			cycleRegValues[cycle] = regX
			cycle++
			instructionIndex++
		}
	}
	return cycleRegValues
}

func calculateStrength(instructions []string) int {
	cyclesRegValues := processSignal(instructions)
	return cyclesRegValues[20]*21 + cyclesRegValues[60]*19 + cyclesRegValues[100]*18 + cyclesRegValues[140]*21 + cyclesRegValues[180]*16 + cyclesRegValues[140]*18
}

//TODO:
//Need to check if all the index exist so the code doesnt break
