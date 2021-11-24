package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type ScoreData struct {
	StepsCounterInA  int
	StepsCounterInB  int
	TypeMoveInStackA string
	TypeMoveInStackB string
}

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

	min, max, stackSize, median := FindValues(&aStack)

	if len(aStack) <= 6 {
		if aStack[0] > aStack[1] && aStack[1] == min {
			Swap(&aStack, "a")
		}

		for i := 0; i < stackSize/2; i++ {
			PushTop(&aStack, &bStack, "b")
		}

		SmallSortA_Stack(&aStack)
		SmallSortB_Stack(&bStack)
		MergeSmallStacks(&aStack, &bStack, max, min)
		return
	}

	if len(aStack) > 6 {
		var index int = 0
		for {
			if aStack[index] != min && aStack[index] != max {
				PushTop(&aStack, &bStack, "b")

				if bStack[0] > median && len(bStack) != 1 {
					Rotate(&bStack, "b")
				}
			} else {
				Rotate(&aStack, "a")
			}

			if len(aStack) == 2 {
				break
			}
		}
		FullSort(&aStack, &bStack)
	}
}

func FindValues(Stack *[]int) (int, int, int, int) {
	var sortedStack []int
	sortedStack = append(sortedStack, *Stack...)
	sort.Ints(sortedStack)

	var min int = sortedStack[0]
	var max int = sortedStack[len(sortedStack)-1]
	var stackLenght int = len(*Stack)
	var median int = (sortedStack)[stackLenght/2]

	return min, max, stackLenght, median
}

func FullSort(A_StackTable *[]int, B_StackTable *[]int) {
	for {
		if len(*B_StackTable) == 0 {
			break
		}

		var scoreList []ScoreData = make([]ScoreData, len(*B_StackTable))

		var steps int
		var stackLenght int = len(scoreList)
		var changeDirection bool
		for i := 0; i < stackLenght; i++ {
			scoreList[i].StepsCounterInB = steps

			if stackLenght%2 == 0 {
				if stackLenght/2 > i {
					steps++
				} else {
					steps--
				}
			} else {
				if i == (stackLenght/2)+1 {
					scoreList[i].StepsCounterInB = steps - 1
					steps--
				}

				if stackLenght/2 >= i {
					steps++
				} else {
					steps--
				}
			}

			if (i == (stackLenght/2)+1 && stackLenght%2 == 1) || (i == (stackLenght/2) && stackLenght%2 == 0) {
				changeDirection = true
			}

			if !changeDirection {
				scoreList[i].TypeMoveInStackB = "r"
			} else {
				scoreList[i].TypeMoveInStackB = "rr"
			}
		}

		var reverseRotateResult, rotateResult, bestResult, bestResultIndex int
		for i := 0; i < len(*B_StackTable); i++ {
			rotateResult = CheckRotate(*A_StackTable, (*B_StackTable)[i], "r")
			reverseRotateResult = CheckRotate(*A_StackTable, (*B_StackTable)[i], "rr")

			if rotateResult <= reverseRotateResult {
				scoreList[i].StepsCounterInA = rotateResult
				scoreList[i].TypeMoveInStackA = "r"
			} else {
				scoreList[i].StepsCounterInA = reverseRotateResult
				scoreList[i].TypeMoveInStackA = "rr"
			}

			if rotateResult == reverseRotateResult {
				if (scoreList[i].TypeMoveInStackA != scoreList[i].TypeMoveInStackB) && scoreList[i].TypeMoveInStackB == "rr" {
					scoreList[i].StepsCounterInA = reverseRotateResult
					scoreList[i].TypeMoveInStackA = "rr"
				}
			}

			possibleBestResult := scoreList[i].StepsCounterInA + scoreList[i].StepsCounterInB
			if bestResult == 0 || (bestResult > possibleBestResult) {
				bestResult = scoreList[i].StepsCounterInA + scoreList[i].StepsCounterInB
				bestResultIndex = i
			}
		}

		var stepsLeftInA, stepsLeftInB int = scoreList[bestResultIndex].StepsCounterInA, scoreList[bestResultIndex].StepsCounterInB
		for {
			if stepsLeftInA == 0 && stepsLeftInB == 0 {
				PushTop(B_StackTable, A_StackTable, "a")
				break
			}

			if scoreList[bestResultIndex].TypeMoveInStackA == scoreList[bestResultIndex].TypeMoveInStackB {
				if stepsLeftInA != 0 && stepsLeftInB != 0 {
					if scoreList[bestResultIndex].TypeMoveInStackA == "r" {
						RotateBoth(A_StackTable, B_StackTable)
					} else {
						ReverseRotateBoth(A_StackTable, B_StackTable)
					}
					stepsLeftInA--
					stepsLeftInB--
					continue
				}
			}

			if stepsLeftInA != 0 {
				FullSortMove(scoreList[bestResultIndex].TypeMoveInStackA, A_StackTable, &stepsLeftInA, "a")
			}

			if stepsLeftInB != 0 {
				FullSortMove(scoreList[bestResultIndex].TypeMoveInStackB, B_StackTable, &stepsLeftInB, "b")
			}
		}
	}

	min, _, stackSize, _ := FindValues(A_StackTable)
	var RotateType string
	var stepsCount int
	for i := 0; i < stackSize; i++ {
		if (*A_StackTable)[i] == min {
			if stackSize%2 == 0 {
				if stackSize/2 <= i {
					RotateType = "rr"
					stepsCount = stackSize - i
				} else {
					RotateType = "r"
					stepsCount = i
				}
			} else {
				if stackSize/2 < i {
					RotateType = "rr"
					stepsCount = stackSize - i
				} else {
					RotateType = "r"
					stepsCount = i
				}
			}
			break
		}
	}

	if RotateType == "r" {
		for i := 0; i < stepsCount; i++ {
			Rotate(A_StackTable, "a")
		}
	} else {
		for i := 0; i < stepsCount; i++ {
			ReverseRotate(A_StackTable, "a")
		}
	}
}

func FullSortMove(typeMove string, stack *[]int, stepsLeft *int, name string) {
	if typeMove == "r" {
		Rotate(stack, name)
	} else {
		ReverseRotate(stack, name)
	}
	*stepsLeft--
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
	if name != "bb" {
		fmt.Println("r" + name)
	}

	newStack := make([]int, len(*stack))

	newStack[len(*stack)-1] = (*stack)[0]

	for i := 1; i < len(*stack); i++ {
		newStack[i-1] = (*stack)[i]
	}

	*stack = newStack
}

// Shifts down all elements of stack by 1 (rra, rrb)
func ReverseRotate(stack *[]int, name string) {
	if name != "bb" {
		fmt.Println("rr" + name)
	}

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

// Executes rotate function for both stacks (rr)
func RotateBoth(A_StackTable *[]int, B_StackTable *[]int) {
	fmt.Println("rr")
	Rotate(A_StackTable, "bb")
	Rotate(B_StackTable, "bb")
}

// Executes reverse rotate function for both stacks (rrr)
func ReverseRotateBoth(A_StackTable *[]int, B_StackTable *[]int) {
	fmt.Println("rrr")
	ReverseRotate(A_StackTable, "bb")
	ReverseRotate(B_StackTable, "bb")
}

func CheckRotate(stack []int, numberToAppend int, rotateType string) int {
	var counter int

	if len(stack) == 1 {
		return 0
	} else if stack[0] > numberToAppend && stack[len(stack)-1] < numberToAppend {
		return 0
	}

	var copyStack []int = make([]int, len(stack))
	copy(copyStack, stack)

	if rotateType == "r" {
		for {
			tempVar := copyStack[len(copyStack)-1]
			copyStack[len(copyStack)-1] = copyStack[0]

			for i := 1; i < len(copyStack)-1; i++ {
				copyStack[i-1] = copyStack[i]
			}

			copyStack[len(copyStack)-2] = tempVar
			counter++

			if copyStack[0] > numberToAppend && copyStack[len(copyStack)-1] < numberToAppend {
				break
			}
		}
	} else {
		for {
			tempVar := copyStack[0]
			copyStack[0] = copyStack[len(copyStack)-1]

			for i := len(copyStack) - 1; i > 1; i-- {
				copyStack[i] = copyStack[i-1]
			}

			copyStack[1] = tempVar
			counter++

			if copyStack[0] > numberToAppend && copyStack[len(copyStack)-1] < numberToAppend {
				break
			}
		}
	}

	return counter
}
