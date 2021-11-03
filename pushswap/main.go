package main

import (
	"fmt"
	"os"
	"reflect"
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
	if os.Args[1] == "" {
		return
	}

	userInput := strings.Split(os.Args[1], " ")
	var aStack []int
	var bStack []int
	var sortedNumbers []int

	err := AppendNumbers(userInput, &aStack)
	if err {
		return
	}

	if sort.IntsAreSorted(aStack) {
		return
	}

	_ = bStack
	sortedNumbers = append(sortedNumbers, aStack...)
	sort.Ints(sortedNumbers)

	if len(aStack) == 1 {
		return
	} else if len(aStack) == 2 {
		if (aStack)[0] > (aStack)[1] {
			Swap(&aStack, "a")
			return
		}
	} else if len(aStack) == 3 {
		SmallSortA_Stack(&aStack)
		return
	}

	if len(aStack) <= 6 {
		min, max, _, _ := FindValues(&aStack)

		aStackSize := len(aStack) / 2
		for i := 0; i < aStackSize; i++ {
			PushTop(&aStack, &bStack, "b")
		}

		SmallSortA_Stack(&aStack)
		SmallSortB_Stack(&bStack)
		MergeSmallStacks(&aStack, &bStack, max, min)
		return
	}
}

func FindValues(Stack *[]int) (int, int, int, int) {
	var sortedStack []int
	sortedStack = append(sortedStack, *Stack...)
	sort.Ints(sortedStack)

	var min int
	var max int
	var stackLenght int = len(*Stack)
	var median int = (sortedStack)[stackLenght/2]
	for i := 0; i < len(*Stack); i++ {
		if i == 0 {
			min = (*Stack)[i]
			max = (*Stack)[i]
			continue
		}

		if min > (*Stack)[i] {
			min = (*Stack)[i]
		}

		if max < (*Stack)[i] {
			max = (*Stack)[i]
		}
	}

	return min, max, stackLenght, median
}

func SmallSortA_Stack(Stack *[]int) {
	min, max, stackLenght, median := FindValues(Stack)
	if sort.IntsAreSorted(*Stack) {
		return
	}

	if stackLenght == 2 {
		if (*Stack)[0] > (*Stack)[1] {
			Swap(Stack, "a")
			return
		}
	}

	if (*Stack)[0] == max {
		Rotate(Stack, "a")
	}

	if (*Stack)[0] == median {
		if (*Stack)[1] == min {
			Swap(Stack, "a")
		} else {
			ReverseRotate(Stack, "a")
		}
		return
	}

	if (*Stack)[1] == max {
		ReverseRotate(Stack, "a")
		Swap(Stack, "a")
	}
}

func SmallSortB_Stack(Stack *[]int) {
	min, max, stackLenght, _ := FindValues(Stack)

	var sortedValues []int
	sortedValues = append(sortedValues, *Stack...)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedValues)))

	if reflect.DeepEqual(*Stack, sortedValues) {
		return
	}

	if stackLenght == 2 {
		if (*Stack)[0] < (*Stack)[1] {
			Swap(Stack, "b")
			return
		}
	}

	if (*Stack)[2] == max {
		if (*Stack)[1] == min {
			ReverseRotate(Stack, "b")
		} else {
			Rotate(Stack, "b")
			Swap(Stack, "b")
		}

		return
	}

	if (*Stack)[1] == max {
		if (*Stack)[0] == min {
			Swap(Stack, "b")
		} else {
			Rotate(Stack, "b")
		}

		return
	}

	ReverseRotate(Stack, "b")
	Swap(Stack, "b")
}

func MergeSmallStacks(A_StackTable *[]int, B_StackTable *[]int, max int, min int) {
	for {
		if len(*B_StackTable) == 0 {
			break
		}

		var counter int
		if (*B_StackTable)[0] == max {
			PushTop(B_StackTable, A_StackTable, "a")
			Rotate(A_StackTable, "a")
			continue
		} else if (*B_StackTable)[0] == min {
			PushTop(B_StackTable, A_StackTable, "a")
			continue
		}

		if (*B_StackTable)[0] < (*A_StackTable)[0] {
			PushTop(B_StackTable, A_StackTable, "a")
		} else {
			for i := 0; i < len(*A_StackTable); i++ {
				if i == 0 {
					if (*A_StackTable)[0] > (*B_StackTable)[0] && (*A_StackTable)[len(*A_StackTable)-1] < (*B_StackTable)[0] {
						counter = i + 1
						break
					}
				} else if i == len(*A_StackTable)-1 {
					if (*A_StackTable)[0] < (*B_StackTable)[0] && (*A_StackTable)[len(*A_StackTable)-1] > (*B_StackTable)[0] {
						counter = i + 1
						break
					}
				} else if (*A_StackTable)[i] < (*B_StackTable)[0] && (*A_StackTable)[i+1] > (*B_StackTable)[0] {
					counter = i + 1
					break
				} else if (*A_StackTable)[i-1] < (*B_StackTable)[0] && (*A_StackTable)[i] > (*B_StackTable)[0] {
					counter = i - 1
					break
				}
			}

			for a := 1; a < len(*A_StackTable); a++ {
				if (counter)-a == 0 || counter == 0 {
					for b := 0; b < a; b++ {
						Rotate(A_StackTable, "a")
					}

					PushTop(B_StackTable, A_StackTable, "a")

					if len(*B_StackTable) != 0 {
						if (*B_StackTable)[0] < (*A_StackTable)[0] && (*B_StackTable)[0] > (*A_StackTable)[len(*A_StackTable)-1] {
							PushTop(B_StackTable, A_StackTable, "a")

						}
					}

					for b := 0; b < a; b++ {
						ReverseRotate(A_StackTable, "a")
					}
					break
				} else if counter+a == len(*A_StackTable) || counter == len(*A_StackTable) {
					for b := 0; b < a; b++ {
						ReverseRotate(A_StackTable, "a")
					}

					PushTop(B_StackTable, A_StackTable, "a")

					if len(*B_StackTable) != 0 {
						if (*B_StackTable)[0] < (*A_StackTable)[0] && (*B_StackTable)[0] > (*A_StackTable)[len(*A_StackTable)-1] {
							PushTop(B_StackTable, A_StackTable, "a")
							a++
						}
					}

					for b := 0; b < a+1; b++ {
						Rotate(A_StackTable, "a")
					}
					break
				}
			}
		}
	}
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
func PushTop(exStack *[]int, impStack *[]int, name string) {
	if len(*exStack) == 0 {
		fmt.Println("Not enought elements")
		return
	}
	fmt.Println("p" + name)
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
func Swap(stack *[]int, name string) {
	if len(*stack) < 2 {
		fmt.Println("Not enought elements")
		return
	}

	fmt.Println("s" + name)
	tempVar := (*stack)[0]
	(*stack)[0] = (*stack)[1]
	(*stack)[1] = tempVar
}

// Shifts up all elements of stack by 1 (ra, rb)
func Rotate(stack *[]int, name string) {
	fmt.Println("r" + name)
	newStack := make([]int, len(*stack))

	newStack[len(*stack)-1] = (*stack)[0]

	for i := 1; i < len(*stack); i++ {
		newStack[i-1] = (*stack)[i]
	}

	*stack = newStack
}

// Shifts down all elements of stack by 1 (rra, rrb)
func ReverseRotate(stack *[]int, name string) {
	fmt.Println("rr" + name)
	newStack := make([]int, len(*stack))

	newStack[0] = (*stack)[len(*stack)-1]

	for i := 1; i < len(*stack); i++ {
		newStack[i] = (*stack)[i-1]
	}

	*stack = newStack
}

// // Executes swap function for both stacks (ss)
// func SwapBoth(A_StackTable *[]int, B_StackTable *[]int) {
// 	Swap(A_StackTable)
// 	Swap(B_StackTable)
// }

// // Executes rotate function for both stacks (rr)
// func RotateBoth(A_StackTable *[]int, B_StackTable *[]int) {
// 	Rotate(A_StackTable)
// 	Rotate(B_StackTable)
// }

// // Executes reverse rotate function for both stacks (rrr)
// func ReverseRotateBoth(A_StackTable *[]int, B_StackTable *[]int) {
// 	ReverseRotate(A_StackTable)
// 	ReverseRotate(B_StackTable)
// }

// Checks if stack A sortend and stack B is clear
func CheckSort(sortedNumbers []int, A_StackTable *[]int, B_StackTable *[]int) bool {
	return reflect.DeepEqual(sortedNumbers, *A_StackTable) && len(*B_StackTable) == 0
}
