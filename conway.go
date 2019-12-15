package main

import (
    "fmt"
    "time"
)

// TODO: make size of board dynamic

type cell bool

func (c cell) String() string {
	if c {
		return fmt.Sprintf("✮") // alive ■ ✮
	} else {
		return fmt.Sprintf("∙") // dead
	}
}

func update(curr [19][19]cell) [19][19]cell {
	var next [19][19]cell

	// function to check neighbours of cell at x,y
	check := func(x, y int) cell {
		var count = 0
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				// make sure we're on the board properly, we're not counting ourselves, and increment counter if this neighbour is alive
				if(x+i > -1 && x+i < 19 && y+j > -1 && y+j < 19 && !(i==0 && j==0) && curr[x+i][y+j]){
					count++
				}
				
			}
		}
		if curr[x][y] {
			return count == 2 || count == 3 // the living need 2 or 3 neighbours to stay alive
		} else {
			return count == 3 // the dead need 3 neighbours to be brought to life
		}
	}

	// for each cell, we update the value according to conway rules
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			next[i][j] = check(i, j)
		}
	}

	return next
}

func main() {
	// initialize board and its values. 
	// each cell is 'dead' or 'alive,' and its next state is determined by its current state and its neighbour's states.
	var board [19][19]cell

	// top left 
	board[3][5] = true
	board[3][6] = true
	board[3][7] = true

	board[8][5] = true
	board[8][6] = true
	board[8][7] = true

	board[7][8] = true
	
	board[5][3] = true
	board[6][3] = true
	board[7][3] = true

	// top right
	board[5][10] = true
	board[6][10] = true
	board[7][10] = true

	board[5][15] = true
	board[6][15] = true


	//bottom left
	board[11][8] = true
	board[12][8] = true
	board[13][8] = true
	
	board[11][3] = true
	board[13][3] = true	

	board[10][5] = true
	board[10][6] = true
	board[10][7] = true

	board[15][5] = true
	board[15][6] = true
	board[15][7] = true

	// bottom right
	board[11][10] = true
	board[12][10] = true
	board[13][10] = true
		

	board[10][11] = true
	board[10][12] = true
	board[10][13] = true

	board[15][11] = true
	board[15][12] = true
	board[15][13] = true

	var step = 1 // initializing step
	
	// run through steps
	for {
		fmt.Printf("Step %d:\n", step) // step number
		step++ // increment step

		// print the board
		for i := 0; i < 19; i++ {
			fmt.Println(board[i])
		}
		fmt.Println() // spacer
		
		board = update(board) // get new board values

		time.Sleep(200*time.Millisecond) // wait some time between steps
	}

}
