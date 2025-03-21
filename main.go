package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var PLAYER_VERTICAL_SPEED float32 = 200.0
var PADDLE_HEIGHT float32 = 80.0
var PADDLE_WIDTH float32 = 10.0

type Paddle struct {
	Width  float32
	Height float32
	Color  color.RGBA
}

type Ball struct {
	Radius float32
	Color  color.RGBA
}

type Player struct {
	Position rl.Vector2
	Speed    int
}

func (p *Player) Update(deltaTime float32) {
	// If UP_ARROW is pressed, move player up
	if rl.IsKeyDown(rl.KeyUp) {
		p.Position.Y -= (PLAYER_VERTICAL_SPEED * deltaTime)
	}

	// If DOWN_ARROW is pressed, move player down
	if rl.IsKeyDown(rl.KeyDown) {
		p.Position.Y += (PLAYER_VERTICAL_SPEED * deltaTime)
	}

	// Keep paddle movement boundaries within screen height
	screenHeight := float32(rl.GetScreenHeight())
	if p.Position.Y < 1 {
		p.Position.Y = 1
	}
	if p.Position.Y > screenHeight-PADDLE_HEIGHT-1 {
		p.Position.Y = screenHeight - PADDLE_HEIGHT - 1
	}
}

func main() {
	var screenWidth, screenHeight int32 = 800, 450

	rl.InitWindow(screenWidth, screenHeight, "Pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Set paddle config
	paddle := Paddle{
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
		Color:  rl.Black,
	}
	ball := Ball{Radius: 7.5, Color: rl.Red}

	// Setup player A
	// Default position is at the center of the screen
	// on the player side
	// Default speed = 0
	playerA := Player{}
	playerA.Position = rl.Vector2{
		X: 5.0,
		Y: float32(screenHeight)/2.0 - float32(paddle.Height)/2.0,
	}

	for !rl.WindowShouldClose() {
		deltaTime := rl.GetFrameTime()

		playerA.Update(deltaTime)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangleV(
			playerA.Position,
			rl.Vector2{
				X: paddle.Width,
				Y: paddle.Height,
			},
			paddle.Color,
		)

		rl.DrawRectangleV(
			rl.Vector2{
				X: float32(screenWidth) - float32(paddle.Width) - 5.0,
				Y: float32(screenHeight)/2.0 - float32(paddle.Height)/2.0,
			},
			rl.Vector2{
				X: paddle.Width,
				Y: paddle.Height,
			},
			paddle.Color,
		)

		rl.DrawCircleV(
			rl.Vector2{
				X: float32(screenWidth)/2.0 - ball.Radius,
				Y: float32(screenHeight)/2.0 - ball.Radius,
			},
			ball.Radius,
			ball.Color,
		)

		rl.EndDrawing()
	}
}
