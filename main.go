package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	title   = "Mountain Hand Line Game"
	debug   = true
	screenX = 320
	screenY = 480
)

const (
	// game mode
	modeTitle    = iota
	modeGame     = iota
	modeGameover = iota
	modeHelp     = iota
)

/*
const (
	shinjuku        = iota
	shinOkubo       = iota
	takadanobaba    = iota
	mejiro          = iota
	ikebukuro       = iota
	otsuka          = iota
	sugamo          = iota
	komagome        = iota
	tabata          = iota
	nishiNippori    = iota
	nippori         = iota
	uguisudani      = iota
	ueno            = iota
	okachimachi     = iota
	akihabara       = iota
	kanda           = iota
	tokyo           = iota
	yurakucho       = iota
	shinbashi       = iota
	hamamatsucho    = iota
	tamachi         = iota
	takanawaGateway = iota
	shinagawa       = iota
	osaki           = iota
	gotanda         = iota
	meguro          = iota
	ebisu           = iota
	shibuya         = iota
	harajuku        = iota
	yoyogi          = iota
)
*/

// 駅情報
type Station struct {
	Name string
	NextStationIn string
	RequireTimeIn  int

	NextStationOut string
	RequireTimeOut int
}


func init() {
	Stations := map[string]Station{
		"sinjuku"  : Station{"sinjuku"  , "tokyo"    , 1, "akihabara", 5},
		"tokyo"    : Station{"tokyo"    , "ueno"     , 2, "sinjuku"  , 6},
		"ueno"     : Station{"ueno"     , "akihabara", 3, "tokyo"    , 7},
		"akihabara": Station{"akihabara", "sinjuku"  , 4, "ueno"     , 8},
	}

	fmt.Println(Stations["sinjuku"].NextStationIn)
	fmt.Printf("%#v", Stations)

	// sinjuku -> ueno
	var sumIn  int = 0
	var sumOut int = 0
	start := "sinjuku"
	goal  := "ueno"

	// In
	tmpStation := Stations[start]
	for {
		sumIn += tmpStation.RequireTimeIn
		tmpStation = Stations[tmpStation.NextStationIn]

		if tmpStation.Name == goal {
			break;
		}
	}

	// Out
	tmpStation = Stations[start]
	for {
		sumOut += tmpStation.RequireTimeOut
		tmpStation = Stations[tmpStation.NextStationOut]

		if tmpStation.Name == goal {
			break;
		}
	}


	fmt.Println(sumIn)
	fmt.Println(sumOut)
}

// Game Main
type Game struct {
	mode int
}

func (g *Game) Update() error {

	switch g.mode {
	case modeTitle:
		if g.isKeyJustPressed(ebiten.KeySpace) {
			g.mode = modeGame
		}
	case modeGame:

	case modeGameover:

	case modeHelp:

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf(
			"GameMode: %d\n",
			g.mode,
		))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func (g *Game) isKeyJustPressed(key ebiten.Key) bool {
	if inpututil.IsKeyJustPressed(key) {
		return true
	}
	return false
}

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
