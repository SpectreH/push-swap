package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		os.Exit(0)
	}
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	userInput := strings.Split(os.Args[1], " ")
	var aStack []int
	var bStack []int

	err := AppendNumbers(userInput, &aStack)
	if err {
		return
	}

	var allInstructions []string = []string{"pa", "pb", "sa", "sb", "ss", "ra", "rb", "rr", "rra", "rrb", "rrr"}
	var instructionsToUse []string

	errorDetected := GetInstructions(allInstructions, &instructionsToUse)
	if errorDetected {
		return
	}

	for i := 0; i < len(instructionsToUse); i++ {
		if instructionsToUse[i] == "pa" {
			PushTop(&bStack, &aStack)
		} else if instructionsToUse[i] == "pb" {
			PushTop(&aStack, &bStack)
		} else if instructionsToUse[i] == "sa" {
			Swap(&aStack)
		} else if instructionsToUse[i] == "sb" {
			Swap(&bStack)
		} else if instructionsToUse[i] == "ss" {
			SwapBoth(&aStack, &bStack)
		} else if instructionsToUse[i] == "ra" {
			Rotate(&aStack)
		} else if instructionsToUse[i] == "rb" {
			Rotate(&bStack)
		} else if instructionsToUse[i] == "rr" {
			RotateBoth(&aStack, &bStack)
		} else if instructionsToUse[i] == "rra" {
			ReverseRotate(&aStack)
		} else if instructionsToUse[i] == "rrb" {
			ReverseRotate(&bStack)
		} else {
			ReverseRotateBoth(&aStack, &bStack)
		}
	}

	if sort.IntsAreSorted(aStack) && len(bStack) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// Reads instructions on the standard input and saves them. If gets non-existing insturction - returns
func GetInstructions(allInstructions []string, instructionsToUse *[]string) bool {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for i, k := range allInstructions {
			if scanner.Text() == "" {
				continue
			}

			if k == scanner.Text() {
				*instructionsToUse = append(*instructionsToUse, scanner.Text())
				break
			}

			if i == len(allInstructions)-1 {
				fmt.Println("Error")
				return true
			}
		}
	}
	return false
}

// Append all number to A stack
func AppendNumbers(userInput []string, aStack *[]int) bool {
	for i := 0; i < len(userInput); i++ {
		numToAppend, err := strconv.Atoi(userInput[i])
		if err != nil {
			fmt.Println("Error")
			return true
		}

		// Check for duplicates
		if i != 0 {
			for k := 0; k < len(*aStack); k++ {
				if (*aStack)[k] == numToAppend {
					fmt.Println("Error")
					return true
				}
			}
		}

		(*aStack) = append((*aStack), numToAppend)
	}
	return false
}

// Pushes the top first element of one stack to another (pa, pb)
func PushTop(exStack *[]int, impStack *[]int) {
	if len(*exStack) == 0 {
		// fmt.Println("Not enought elements")
		return
	}
	tempVar := (*exStack)[0]
	*exStack = append((*exStack)[:0], (*exStack)[0+1:]...)

	if len(*impStack) == 0 {
		*impStack = append(*impStack, tempVar)
		return
	}

	*impStack = append((*impStack)[:0+1], (*impStack)[0:]...)
	(*impStack)[0] = tempVar
}

// Swap first 2 elements of stack (sb,sa)
func Swap(stack *[]int) {
	if len(*stack) < 2 {
		// fmt.Println("Not enought elements")
		return
	}

	tempVar := (*stack)[0]
	(*stack)[0] = (*stack)[1]
	(*stack)[1] = tempVar
}

// Shifts up all elements of stack by 1 (ra, rb)
func Rotate(stack *[]int) {
	if len(*stack) < 2 {
		return
	}

	newStack := make([]int, len(*stack))

	newStack[len(*stack)-1] = (*stack)[0]

	for i := 1; i < len(*stack); i++ {
		newStack[i-1] = (*stack)[i]
	}

	*stack = newStack
}

// Shifts down all elements of stack by 1 (rra, rrb)
func ReverseRotate(stack *[]int) {
	if len(*stack) < 2 {
		return
	}

	newStack := make([]int, len(*stack))

	newStack[0] = (*stack)[len(*stack)-1]

	for i := 1; i < len(*stack); i++ {
		newStack[i] = (*stack)[i-1]
	}

	*stack = newStack
}

// Executes swap function for both stacks (ss)
func SwapBoth(A_StackTable *[]int, B_StackTable *[]int) {
	Swap(A_StackTable)
	Swap(B_StackTable)
}

// Executes rotate function for both stacks (rr)
func RotateBoth(A_StackTable *[]int, B_StackTable *[]int) {
	Rotate(A_StackTable)
	Rotate(B_StackTable)
}

// Executes reverse rotate function for both stacks (rrr)
func ReverseRotateBoth(A_StackTable *[]int, B_StackTable *[]int) {
	ReverseRotate(A_StackTable)
	ReverseRotate(B_StackTable)
}
