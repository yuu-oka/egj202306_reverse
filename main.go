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

// 駅情報
type Station struct {
	Name             string
	NameEn           string
	NextStationIdOut int
	RequireTimeOut   int

	NextStationIdIn int
	RequireTimeIn   int
}

func init() {
	Stations := [...]Station{
		Station{"新宿",             "shinjuku",         shinOkubo,       2, yoyogi,          2},
		Station{"新大久保",         "shin-okubo",       takadanobaba,    2, shinjuku,        2},
		Station{"高田馬場",         "takadanobaba",     mejiro,          2, shinOkubo,       2},
		Station{"目白",             "mejiro",           ikebukuro,       3, takadanobaba,    2},
		Station{"池袋",             "ikebukuro",        otsuka,          2, mejiro,          2},
		Station{"大塚",             "otsuka",           sugamo,          2, ikebukuro,       2},
		Station{"巣鴨",             "sugamo",           komagome,        2, otsuka,          2},
		Station{"駒込",             "komagome",         tabata,          2, sugamo,          2},
		Station{"田端",             "tabata",           nishiNippori,    2, komagome,        2},
		Station{"西日暮里",         "nishi-nippori",    nippori,         2, tabata,          2},
		Station{"日暮里",           "nippori",          uguisudani,      2, nishiNippori,    2},
		Station{"鶯谷",             "uguisudani",       ueno,            2, nippori,         2},
		Station{"上野",             "ueno",             okachimachi,     2, uguisudani,      2},
		Station{"御徒町",           "okachimachi",      akihabara,       2, ueno,            2},
		Station{"秋葉原",           "akihabara",        kanda,           2, okachimachi,     2},
		Station{"神田",             "kanda",            tokyo,           2, akihabara,       2},
		Station{"東京",             "tokyo",            yurakucho,       2, kanda,           2},
		Station{"有楽町",           "yurakucho",        shinbashi,       2, tokyo,           2},
		Station{"新橋",             "shinbashi",        hamamatsucho,    2, yurakucho,       2},
		Station{"浜松町",           "hamamatsucho",     tamachi,         2, shinbashi,       2},
		Station{"田町",             "tamachi",          takanawaGateway, 2, hamamatsucho,    2},
		Station{"高輪ゲートウェイ", "takanawa-gateway", shinagawa,       2, tamachi,         2},
		Station{"品川",             "shinagawa",        osaki,           2, takanawaGateway, 2},
		Station{"大崎",             "osaki",            gotanda,         2, shinagawa,       2},
		Station{"五反田",           "gotanda",          meguro,          2, osaki,           2},
		Station{"目黒",             "meguro",           ebisu,           2, gotanda,         2},
		Station{"恵比寿",           "ebisu",            shibuya,         2, meguro,          2},
		Station{"渋谷",             "shibuya",          harajuku,        2, ebisu,           2},
		Station{"原宿",             "harajuku",         yoyogi,          2, shibuya,         2},
		Station{"代々木",           "yoyogi",           shinjuku,        2, harajuku,        2},
	}

	// sinjuku -> ueno
	var sumIn int = 0
	var sumOut int = 0
	start := shinjuku
	goal := ueno

	// In
	tmpStation := Stations[start]
	for {
		sumIn += tmpStation.RequireTimeIn
		tmpStation = Stations[tmpStation.NextStationIdIn]

		if tmpStation.Name == Stations[goal].Name {
			break
		}
	}

	// Out
	tmpStation = Stations[start]
	for {
		sumOut += tmpStation.RequireTimeOut
		tmpStation = Stations[tmpStation.NextStationIdOut]

		if tmpStation.Name == Stations[goal].Name {
			break
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
