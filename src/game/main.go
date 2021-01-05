package main

import (
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// GimulType
type GimulType int

// 보드 상수 선언
const (
	GimulTypeNone      GimulType = -1
	GimulTypeGreenWang GimulType = 0
	GimulTypeGreenJa   GimulType = 1
	GimulTypeGreenJang GimulType = 2
	GimulTypeGreenSang GimulType = 3
	GimulTypeRedWang   GimulType = 4
	GimulTypeRedJa     GimulType = 5
	GimulTypeRedJang   GimulType = 6
	GimulTypeRedSang   GimulType = 7
	GimulTypeMax       GimulType = 8
)

//
const (
	ScreenWidth  = 500
	ScreenHeight = 400
	BoardWidth   = 4
	BoardHeight  = 3
	GimulStartX  = 20
	GimulStartY  = 23
	GridWidth    = 116
	GridHeight   = 116
)

// Game implements ebiten.Game interface.
type Game struct{}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	if gameover {
		return nil
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	screen.DrawImage(bgimg, nil)

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/GridWidth, y/GridHeight

		if i >= 0 && i < GridWidth && j >= 0 && j < GridHeight {
			if !selected {
				if board[i][j] != GimulTypeNone && currentTeam == GetTeamType(board[i][j]) {
					selected = true
					selectedX, selectedY = i, j
				}
			} else {
				if i == selectedX && j == selectedY {
					selected = false
				} else {
					moveGimul(selectedX, selectedY, i, j)
				}
			}
		}
	}

	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+GridHeight*j))

			switch board[i][j] {
			case GimulTypeGreenWang:
				screen.DrawImage(gimulImgs[GimulTypeGreenWang], op)
			case GimulTypeGreenJa:
				screen.DrawImage(gimulImgs[GimulTypeGreenJa], op)
			case GimulTypeGreenJang:
				screen.DrawImage(gimulImgs[GimulTypeGreenJang], op)
			case GimulTypeGreenSang:
				screen.DrawImage(gimulImgs[GimulTypeGreenSang], op)
			case GimulTypeRedWang:
				screen.DrawImage(gimulImgs[GimulTypeRedWang], op)
			case GimulTypeRedJa:
				screen.DrawImage(gimulImgs[GimulTypeRedJa], op)
			case GimulTypeRedJang:
				screen.DrawImage(gimulImgs[GimulTypeRedJang], op)
			case GimulTypeRedSang:
				screen.DrawImage(gimulImgs[GimulTypeRedSang], op)
			}
		}
	}

	if selected {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(GimulStartX+GridWidth*selectedX), float64(GimulStartY+GridHeight*selectedY))
		screen.DrawImage(selectedImg, opts)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// 팀
type TeamType int

// 팀 구분
const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

var (
	board       [BoardWidth][BoardHeight]GimulType
	bgimg       *ebiten.Image
	selectedImg *ebiten.Image
	gimulImgs   [GimulTypeMax]*ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType = TeamGreen
	gameover    bool
)

func GetTeamType(gimulType GimulType) TeamType {
	if gimulType == GimulTypeGreenJa ||
		gimulType == GimulTypeGreenJang ||
		gimulType == GimulTypeGreenSang ||
		gimulType == GimulTypeGreenWang {
		return TeamGreen
	}

	if gimulType == GimulTypeRedJa ||
		gimulType == GimulTypeRedJang ||
		gimulType == GimulTypeRedSang ||
		gimulType == GimulTypeRedWang {
		return TeamGreen
	}

	return TeamNone
}

func moveGimul(prevX, prevY, tarX, tarY int) {
	if isMovable(prevX, prevY, tarX, tarY) {
		OnDie(board[tarX][tarY])
		board[prevX][prevY], board[tarX][tarY] = GimulTypeNone, board[prevX][prevY]
		selected = false

		if currentTeam == TeamGreen {
			currentTeam = TeamRed
		} else {
			currentTeam = TeamGreen
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isMovable(prevX, prevY, tarX, tarY int) bool {
	if GetTeamType(board[prevX][prevY]) == GetTeamType(board[tarX][tarY]) {
		return false
	}
	switch board[prevX][prevY] {
	case GimulTypeGreenJa:
		return prevY == tarY && prevX+1 == tarX
	case GimulTypeRedJa:
		return prevY == tarY && prevX-1 == tarX
	case GimulTypeGreenJang, GimulTypeRedJang:
		return abs(prevX-tarX)+abs(prevY-tarY) == 1
	case GimulTypeGreenSang, GimulTypeRedSang:
		return (abs(prevX-tarX) == 1 && abs(prevY-tarY) == 1)
	case GimulTypeGreenWang, GimulTypeRedWang:
		return (abs(prevX-tarX) == 1 || abs(prevY-tarY) == 1)
	}
	return true
}

// OnDie calls when gimul is died
func OnDie(gimulType GimulType) {
	if gimulType == GimulTypeGreenWang ||
		gimulType == GimulTypeRedWang {
		gameover = true
	}
}

// main
func main() {

	// 배경 이미지
	reader, err := os.Open("./images/bgimg.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bgimg = ebiten.NewImageFromImage(m)

	// Draw GimulTypeGreenWang
	reader, err = os.Open("./images/green_wang.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeGreenWang] = ebiten.NewImageFromImage(m)

	// Draw GimulTypeGreenJa
	reader, err = os.Open("./images/green_ja.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeGreenJa] = ebiten.NewImageFromImage(m)

	// Draw GimulTypeGreenJang
	reader, err = os.Open("./images/green_jang.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeGreenJang] = ebiten.NewImageFromImage(m)

	// Draw GimulTypeGreenSang
	reader, err = os.Open("./images/green_sang.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeGreenSang] = ebiten.NewImageFromImage(m)

	// Draw GimulTypeRedWang
	reader, err = os.Open("./images/red_wang.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeRedWang] = ebiten.NewImageFromImage(m)

	// Draw GimulTypeRedJa
	reader, err = os.Open("./images/red_ja.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeRedJa] = ebiten.NewImageFromImage(m)

	// Draw GimulTypeRedJang
	reader, err = os.Open("./images/red_jang.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeRedJang] = ebiten.NewImageFromImage(m)

	// Draw GimulTypeRedSang
	reader, err = os.Open("./images/red_sang.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	gimulImgs[GimulTypeRedSang] = ebiten.NewImageFromImage(m)

	// Draw Selected Image
	reader, err = os.Open("./images/selected.png")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err = image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	selectedImg = ebiten.NewImageFromImage(m)

	// 이미지 배열 초기화
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulTypeNone
		}
	}

	board[0][0] = GimulTypeGreenSang
	board[0][1] = GimulTypeGreenWang
	board[0][2] = GimulTypeGreenJang
	board[1][1] = GimulTypeGreenJa

	board[3][0] = GimulTypeRedSang
	board[3][1] = GimulTypeRedWang
	board[3][2] = GimulTypeRedJang
	board[2][1] = GimulTypeRedJa

	game := &Game{}
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("12 janggi")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
