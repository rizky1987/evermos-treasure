package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)


	obstaclePosition := [][]int{
		{0,0},{0,1},{0,2},{0,3},{0,4},{0,5},{0,6},{0,7},
		{1,0},{1,7},
		{2,0},{2,2},{2,3},{2,4},{2,7},
		{3,0},{3,4},{3,6},{3,7},
		{4,0},{4,2},{4,7},
		{5,0},{5,1},{5,2},{5,3},{5,4},{5,5},{5,6},{5,7},
	}

	message := ""
	possibilityTreasureMarking := "."
	playerLocation := [][]int{{4,1}}
	treasureLocation := [][]int{{1,6}}
	var clearPathLocation [][]int
	isShowClearPathLocation := false
	isShowTreasureLocation := false
	for {

		fmt.Println("=== Welcome To Treasure Hunt ===")
		fmt.Println("")
		fmt.Println("You must find the treasure between the clear path and you must avoid the obstacle")
		fmt.Println("")

		fmt.Println("On the map bellow you can see a symbol that representation to something")
		fmt.Println("`#` represents an obstacle")
		fmt.Println("`.` represents a clear path")
		fmt.Println("`x` represents the player’s starting position")
		fmt.Println("")

		fmt.Println("For your controller you can type and press 'Enter' :")
		fmt.Println("'A' for up")
		fmt.Println("'B' for right")
		fmt.Println("'C' for down")
		fmt.Println("'D' for left")
		fmt.Println("")

		fmt.Println("For Cheat you can type and press 'Enter' : ")
		fmt.Println("'L' for all the possibility of probable coordinate points where the treasure might be located")
		fmt.Println("'S' for change probable treasure locations marked with '$'")
		fmt.Println("'T' for see where is the treasure location and marked as 'T'")
		fmt.Println("")
		fmt.Println(message)

		for i:=0; i<6; i++ {
			for j:=0; j<8; j++ {

				if CheckingObstaclePosition(obstaclePosition, i,j){
					fmt.Print("#")
				} else {

					if GeneratePlayer(playerLocation, i, j) {
						fmt.Print("x")
					} else {

						currentLocation := []int{i,j}
						clearPathLocation = append(clearPathLocation, currentLocation)

						if isShowTreasureLocation && GenerateTreasure(treasureLocation, i, j){
							fmt.Print("T")
						} else {
							fmt.Print(possibilityTreasureMarking)
						}

					}
				}
			}
			fmt.Println("")
		}

		if isShowClearPathLocation {
			PrintClearPathLocation(clearPathLocation)
		}
		inputText, _ := reader.ReadString('\n')
		// convert CRLF to LF
		inputText = strings.TrimSpace(strings.ToLower(inputText))

		isShowClearPathLocation = false
		isShowTreasureLocation = false
		if inputText == "a"{
			playerLocation, message = Move(obstaclePosition, playerLocation, treasureLocation, "up")
		} else if inputText == "b" {
			playerLocation, message = Move(obstaclePosition, playerLocation, treasureLocation, "right")
		} else if inputText == "c" {
			playerLocation, message = Move(obstaclePosition, playerLocation, treasureLocation, "down")
		} else if inputText == "d" {
			playerLocation, message = Move(obstaclePosition, playerLocation, treasureLocation,"left")
		} else if inputText == "l" {
			isShowClearPathLocation = true
		} else if inputText == "s" {
			possibilityTreasureMarking = "$"
		} else if inputText == "t" {
			isShowTreasureLocation = true
		} else {
			message = "We can't recognized your input"
		}

		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

	}

}

func PrintClearPathLocation(clearPathLocation [][]int) {

	for _, row := range clearPathLocation{
		fmt.Println(row)
	}
}

func Move(obstacleLocation, playerLocation, treasureLocation [][]int, moveType string)([][]int, string){

	var i,j int
	for _, row := range playerLocation {
		i = row[0]
		j = row[1]
	}

	if moveType == "up" {
		i = i - 1
	} else if  moveType == "down" {
		i = i + 1
	} else if  moveType == "left" {
		j = j - 1
	} else {
		j = j + 1
	}

	if CheckingObstaclePosition(obstacleLocation, i, j){

		return playerLocation, "Oooops you run into an obstacle"
	} else {

		var ti,tj int
		for _, row := range treasureLocation {
			ti = row[0]
			tj = row[1]
		}

		playerLocation = [][]int{{i,j}}
		if ti == i && tj == j{

			return playerLocation, "Yeay.... You Found the treasure"
		}


		return playerLocation, ""
	}

}

func GenerateTreasure(treasureLocation [][]int, i,j int) bool {

	for _, row := range treasureLocation {
		if row[0]== i && row[1]== j {
			return true
		}

	}

	return false
}

func GeneratePlayer(playerLocation [][]int, i,j int) bool {

	for _, row := range playerLocation {
		if row[0]== i && row[1]== j {
			return true
		}

	}

	return false
}

func CheckingObstaclePosition(obstacleLocation [][]int, i,j int) bool {

	for _, row := range obstacleLocation {
		if row[0]== i && row[1]== j {
			return true
		}

	}

	return false
}
