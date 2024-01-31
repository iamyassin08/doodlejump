package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	platformWidth  = 20
	platformHeight = 1
	maxPlatforms   = 10
	screenWidth    = 40
	screenHeight   = 20
	doodleWidth    = 1
	doodleHeight   = 1
	jumpHeight     = 2
)

type Platform struct {
	x, y int
}

type Doodle struct {
	x, y     int
	velocity int
}

var platforms []Platform
var doodle Doodle
var gameOver bool

func main() {
	rand.Seed(time.Now().UnixNano())

	initializeGame()
	for !gameOver {
		renderGame()
		processInput()
		updateGame()
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Game Over! Your score:", doodle.y)
}

func initializeGame() {
	platforms = make([]Platform, maxPlatforms)
	for i := 0; i < maxPlatforms; i++ {
		platforms[i] = Platform{x: rand.Intn(screenWidth - platformWidth), y: rand.Intn(screenHeight)}
	}

	doodle = Doodle{x: screenWidth / 2, y: screenHeight - doodleHeight - 1, velocity: 0}
	gameOver = false
}

func renderGame() {
	fmt.Print("\033[H\033[2J") // Clear the screen
	for y := 0; y < screenHeight; y++ {
		for x := 0; x < screenWidth; x++ {
			isDoodleHere := x >= doodle.x && x < doodle.x+doodleWidth && y == doodle.y
			isPlatformHere := false
			for _, platform := range platforms {
				if x >= platform.x && x < platform.x+platformWidth && y >= platform.y && y < platform.y+platformHeight {
					isPlatformHere = true
					break
				}
			}
			if isDoodleHere {
				fmt.Print("D")
			} else if isPlatformHere {
				fmt.Print("=")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func processInput() {
	// No user input in this simplified example
}

func updateGame() {
	doodle.y += doodle.velocity
	doodle.velocity += 1

	if doodle.y >= screenHeight {
		gameOver = true
	}

	if doodle.velocity < 0 {
		for i := range platforms {
			platform := &platforms[i]
			if doodle.x >= platform.x && doodle.x < platform.x+platformWidth && doodle.y == platform.y {
				doodle.velocity = jumpHeight
				break
			}
		}
	}

	if rand.Float64() < 0.2 && len(platforms) < maxPlatforms {
		platforms = append(platforms, Platform{x: rand.Intn(screenWidth - platformWidth), y: 0})
	}

	for i := range platforms {
		platform := &platforms[i]
		platform.y += 1
		if platform.y > screenHeight {
			copy(platforms[i:], platforms[i+1:])
			platforms = platforms[:len(platforms)-1]
		}
	}
}
