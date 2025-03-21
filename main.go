package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// TODO: MOVE TO CONSTANTS FILE?
var PLAYER_VERTICAL_SPEED float32 = 200.0
var PLAYER_OFFSET float32 = 5.0

var PADDLE_HEIGHT float32 = 80.0
var PADDLE_WIDTH float32 = 10.0

var BALL_SPEED_X float32 = 4.0
var BALL_SPEED_Y float32 = 5.0
var BALL_RADIUS float32 = 7.5
var BALL_COLOR color.RGBA = rl.Red

type Paddle struct {
	Width  float32
	Height float32
	Color  color.RGBA
}

type Ball struct {
	Speed    rl.Vector2
	Position rl.Vector2
	Radius   float32
	Color    color.RGBA
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

func (b *Ball) Update() {
	b.Position.X += b.Speed.X
	b.Position.Y += b.Speed.Y

	screenHeight := float32(rl.GetScreenHeight())

	// Ball must bounce off of top and bottom walls
	if b.Position.Y >= screenHeight-BALL_RADIUS ||
		b.Position.Y <= BALL_RADIUS {
		b.Speed.Y *= -1.0
	}
}

func checkCollision(playerA *Player, ball *Ball) {
	screenWidth := float32(rl.GetScreenWidth())
	// screenHeight := float32(rl.GetScreenHeight())

	// Check collision with player A (left paddle))
	if ball.Position.X <= playerA.Position.X+PADDLE_WIDTH+PLAYER_OFFSET+BALL_RADIUS &&
		ball.Position.Y >= playerA.Position.Y &&
		ball.Position.Y <= playerA.Position.Y+PADDLE_HEIGHT {
		ball.Speed.X *= -1.0
	}

	if ball.Position.X >= screenWidth-BALL_RADIUS ||
		ball.Position.X <= BALL_RADIUS {
		ball.Speed.X *= -1.0
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

	// Set ball config
	ball := Ball{}
	ball.Speed = rl.Vector2{
		X: BALL_SPEED_X,
		Y: BALL_SPEED_Y,
	}
	ball.Position = rl.Vector2{
		X: float32(screenWidth/2.0) - BALL_RADIUS,
		Y: float32(screenHeight/2.0) - BALL_RADIUS,
	}
	ball.Radius = BALL_RADIUS
	ball.Color = BALL_COLOR

	// Setup player A
	// Default position is at the center of the screen
	// on the player side
	// Default speed = 0
	playerA := Player{}
	playerA.Position = rl.Vector2{
		X: PLAYER_OFFSET,
		Y: float32(screenHeight)/2.0 - float32(paddle.Height)/2.0,
	}

	for !rl.WindowShouldClose() {
		deltaTime := rl.GetFrameTime()

		// Update player position
		playerA.Update(deltaTime)

		// Update ball position
		ball.Update()

		// Check ball and player collision
		checkCollision(&playerA, &ball)

		// BEGIN DRAWING
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
			ball.Position,
			ball.Radius,
			ball.Color,
		)

		rl.EndDrawing()
	}
}
