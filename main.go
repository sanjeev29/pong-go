package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Paddle struct {
	Width  int
	Height int
	Color  color.RGBA
}

type Ball struct {
	Radius float32
	Color  color.RGBA
}

func main() {
	var screenWidth, screenHeight int32 = 800, 450

	rl.InitWindow(screenWidth, screenHeight, "Pong")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Set paddle config
	paddle := Paddle{Width: 10, Height: 80, Color: rl.Black}
	ball := Ball{Radius: 7.5, Color: rl.Red}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawRectangleV(
			rl.Vector2{
				X: 5.0,
				Y: float32(screenHeight)/2.0 - float32(paddle.Height)/2.0,
			},
			rl.Vector2{
				X: float32(paddle.Width),
				Y: float32(paddle.Height),
			},
			paddle.Color,
		)

		rl.DrawRectangleV(
			rl.Vector2{
				X: float32(screenWidth) - float32(paddle.Width) - 5.0,
				Y: float32(screenHeight)/2.0 - float32(paddle.Height)/2.0,
			},
			rl.Vector2{
				X: float32(paddle.Width),
				Y: float32(paddle.Height),
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
