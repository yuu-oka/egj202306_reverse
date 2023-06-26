package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	_ "image/png"
	"fmt"
	"log"
	"math/rand"
	"time"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	title    = "Mountain Hand Line Game"
	debug    = true
	screenX  = 320
	screenY  = 480

	count    = 5
)

const (
	// game mode
	modeTitle    = iota
	modeGame     = iota
	modeResult   = iota
	modeHelp     = iota
)

var (
	stations [30]Station
	mplusNormalFont font.Face
)

//go:embed resources/images/yamanote_train.png
var byteTrainImg []byte
//go:embed resources/images/yamanote.png
var byteYamanoteImg []byte
//go:embed resources/images/reverce.png
var byteReverceImg []byte
//go:embed resources/images/kanban.png
var byteKanbanImg []byte

var (
	trainImg *ebiten.Image
	yamanoteImg *ebiten.Image
	reverceImg *ebiten.Image
	kanbanImg *ebiten.Image

	ptime time.Time
	atime time.Time
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

func _makeImg(byteImg []byte) *ebiten.Image{
	img, _, err := image.Decode(bytes.NewReader(byteImg))
	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func init() {
	rand.Seed(time.Now().UnixNano())


	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}
	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    20,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	trainImg    = _makeImg(byteTrainImg)
	yamanoteImg = _makeImg(byteYamanoteImg)
	reverceImg  = _makeImg(byteReverceImg)
	kanbanImg   = _makeImg(byteKanbanImg)

	stations = [...]Station{
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
}

// Game Main
type Game struct {
	mode int
	counter int
	
	direction int

	from int
	to int

	inSum int
	outSum int

	state int
}

func (g *Game) Init() {
	g.inSum   = 0
	g.outSum  = 0
	g.counter = 0
	g.state   = 0
}

func (g *Game) Update() error {

	switch g.mode {
	case modeTitle:
		if g.isKeyJustPressed(ebiten.KeySpace) {
			g.mode = modeGame
			g.Init()

		} else if g.isKeyJustPressed(ebiten.KeyH) {
			g.mode = modeHelp
		}
	case modeGame:

		if g.state == 0 {
			g.from = rand.Intn(len(stations))
			g.to   = rand.Intn(len(stations))

			if g.isKeyJustPressed(ebiten.KeySpace) {
				g.state = 1
				ptime = time.Now()

				// In
				tmpStation := stations[g.from]
				for {
					g.inSum += tmpStation.RequireTimeIn
					tmpStation = stations[tmpStation.NextStationIdIn]

					if tmpStation.Name == stations[g.to].Name {
						break
					}
				}

				// Out
				tmpStation = stations[g.from]
				for {
					g.outSum += tmpStation.RequireTimeOut
					tmpStation = stations[tmpStation.NextStationIdOut]

					if tmpStation.Name == stations[g.to].Name {
						break
					}
				}
			}

		} else {
			atime = time.Now()
			g.counter = int((atime.UnixNano() - ptime.UnixNano()) / int64(time.Second))
			if g.counter > 5 {
				g.mode = modeResult
			}
		}

	case modeResult:
		if g.isKeyJustPressed(ebiten.KeyEscape) {
			g.mode = modeTitle
		}

	case modeHelp:
		if g.isKeyJustPressed(ebiten.KeyEscape) {
			g.mode = modeTitle
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear

	// タイトル画面と結果画面での表示
	switch g.mode {
	case modeTitle:
		screen.DrawImage(reverceImg, op)

	case modeGame:
		op.GeoM.Translate(0, 80)
		screen.DrawImage(kanbanImg, op)

		op.GeoM.Translate(170, 0)
		screen.DrawImage(kanbanImg, op)

		text.Draw(screen, stations[g.from].Name, mplusNormalFont, 20, 120, color.Black)
		text.Draw(screen, stations[g.to].Name,   mplusNormalFont, 200, 120, color.Black)

		op = &ebiten.DrawImageOptions{}
		op.Filter = ebiten.FilterLinear

		op.GeoM.Translate(150, 250)
		screen.DrawImage(trainImg, op)

		text.Draw(screen, strconv.Itoa((count - g.counter)),  mplusNormalFont, 0, 240, color.White)
	case modeResult:
		
	case modeHelp:
		op.GeoM.Translate(0, 0)
		screen.DrawImage(yamanoteImg, op)

		op.GeoM.Translate(0, 200)
		screen.DrawImage(trainImg, op)

		text.Draw(screen, "hoge", mplusNormalFont, 0, 80, color.White)
	}


	if debug {
		ebitenutil.DebugPrint(screen, fmt.Sprintf(
			"GameMode: %d\n",
			g.mode,
		))

		text.Draw(screen, strconv.Itoa(g.inSum), mplusNormalFont, 100, 80, color.White)
		text.Draw(screen, strconv.Itoa(g.outSum),   mplusNormalFont, 100, 160, color.White)
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
