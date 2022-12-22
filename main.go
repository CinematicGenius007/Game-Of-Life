package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findLife(mat [][]int, i int, j int, n int) int {
	sum := mat[(i+1)%n][j] + mat[(n+i-1)%n][j] + mat[i][(j+1)%n] + mat[i][(n+j-1)%n] + mat[(i+1)%n][(j+1)%n] + mat[(i+1)%n][(n+j-1)%n] + mat[(n+i-1)%n][(n+j-1)%n] + mat[(n+i-1)%n][(j+1)%n]
	if (mat[i][j] == 0 && sum == 3) || (mat[i][j] == 1 && (sum == 2 || sum == 3)) {
		return 1
	}
	return 0
}

func printGame(mat [][]int, generation, alive int) {
	fmt.Printf("Generation #%d\n", generation)
	fmt.Printf("Alive: %d\n", alive)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	n, evolution := 0, 9
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}

	current := make([][]int, n)

	firstGenAlive := 0

	for i := 0; i < n; i++ {
		current[i] = make([]int, n)
		for j := 0; j < n; j++ {
			current[i][j] = rand.Intn(2)
			if current[i][j] == 1 {
				firstGenAlive++
			}
		}
	}

	printGame(current, 1, firstGenAlive)

	generation := 2

	for i := 0; i < evolution; i++ {
		time.Sleep(500 * time.Millisecond)
		next := make([][]int, n)
		alive := 0
		for j := 0; j < n; j++ {
			next[j] = make([]int, n)
			for k := 0; k < n; k++ {
				next[j][k] = findLife(current, j, k, n)
				if next[j][k] == 1 {
					alive++
				}
			}
		}
		current = next
		printGame(current, generation, alive)
		generation++
	}

	/*for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if current[i][j] == 1 {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}*/
}
