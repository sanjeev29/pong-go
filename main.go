package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth  = 800.0
	screenHeight = 600.0
	paddleWidth  = 10.0
	paddleHeight = 80.0
	paddleOffset = 5.0
	paddleSpeed  = 5.0
	ballRadius   = 7.5
	ballSpeed    = 5.0
)

var (
	// Right side
	playerPaddle = rl.Rectangle{
		X:      screenWidth - paddleWidth - paddleOffset,
		Y:      screenHeight/2.0 - paddleHeight/2.0,
		Width:  paddleWidth,
		Height: paddleHeight,
	}

	// Left side
	botPaddle = rl.Rectangle{
		X:      paddleOffset,
		Y:      screenHeight/2.0 - paddleHeight/2.0,
		Width:  paddleWidth,
		Height: paddleHeight,
	}

	ball = rl.Vector2{
		X: screenWidth/2.0 - ballRadius/2.0,
		Y: screenHeight/2.0 - ballRadius/2.0,
	}
	ballVelocity = rl.Vector2{
		X: ballSpeed,
		Y: ballSpeed,
	}
)

func main() {
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "Pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// Updates game elements
		updateGame()

		// Draw game elements
		drawGame()
	}
}

func updateGame() {
	updateBall()

	updatePlayerPaddle()

	updateBotPaddle()
}

func drawGame() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	// Draw playerPaddle
	rl.DrawRectangleRec(playerPaddle, rl.Blue)

	// Draw botPaddle
	rl.DrawRectangleRec(botPaddle, rl.Green)

	// Draw ball
	rl.DrawCircle(int32(ball.X), int32(ball.Y), ballRadius, rl.Red)

	// Draw center line
	rl.DrawLine(screenWidth/2.0, 0, screenWidth/2.0, screenHeight, rl.RayWhite)

	rl.EndDrawing()

}

func updateBall() {
	ball.X += ballVelocity.X
	ball.Y += ballVelocity.Y

	// Ball out of bounds
	if ball.X <= 0 || ball.X >= screenWidth {
		// Ball position and speed are reset
		ball.X = screenWidth/2.0 - ballRadius/2.0
		ball.Y = screenHeight/2.0 - ballRadius/2.0
		ballVelocity.X *= -1.0
	}

	// Ball must bounce off of top and bottom walls
	if ball.Y >= screenHeight-ballRadius ||
		ball.Y <= ballRadius {
		ballVelocity.Y *= -1.0
	}

	// Ball collision with player / bot paddle
	if rl.CheckCollisionCircleRec(ball, ballRadius, playerPaddle) ||
		rl.CheckCollisionCircleRec(ball, ballRadius, botPaddle) {
		ballVelocity.X *= -1.0
	}
}

func updatePlayerPaddle() {
	// If UP_ARROW is pressed, move player up
	if rl.IsKeyDown(rl.KeyUp) &&
		playerPaddle.Y > 0 {
		playerPaddle.Y -= (paddleSpeed)
	}

	// If DOWN_ARROW is pressed, move player down
	if rl.IsKeyDown(rl.KeyDown) &&
		playerPaddle.Y+paddleHeight < screenHeight {
		playerPaddle.Y += (paddleSpeed)
	}
}

// Simple bot movement based on ball.Y position
func updateBotPaddle() {
	if ball.Y > botPaddle.Y+(paddleHeight/2.0) {
		botPaddle.Y += (paddleSpeed)
	} else if ball.Y < botPaddle.Y+(paddleHeight/2.0) {
		botPaddle.Y -= (paddleSpeed)
	}

	if botPaddle.Y < 0 {
		botPaddle.Y = 0
	} else if botPaddle.Y+paddleHeight > screenHeight {
		botPaddle.Y = screenHeight - paddleHeight
	}
}
