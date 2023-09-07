package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"kc.com/kc/characters"
	"kc.com/kc/db"
	"kc.com/kc/helpers"
	"kc.com/kc/managers"
	"kc.com/kc/maps"
	"kc.com/kc/objects"
	"kc.com/kc/spawners"
)

var ()

func init() {

	err := maps.LoadMap("./maps/test.json")

	if err != nil {
		print("here")
		println(err.Error())
	}

	err = db.NewItemDb()

	if err != nil {
		print("here")
		println(err.Error())
	}

	err = db.NewPlantDb()

	if err != nil {
		print("here")
		println(err.Error())
	}

	err = helpers.NewNumbers("./numbers.png")

	if err != nil {
		print("here")
		println(err.Error())
	}

	managers.NewEventManager()
	managers.NewImageManager()
	managers.NewWeatherManager()
	managers.NewDebugManager()
	managers.NewCollisionManager(640)
	managers.NewKeyPressManager()
	managers.NewCharacterManager()
	managers.NewObjectManager()

	err = characters.SpawnFarmer()

	if err != nil {
		print("here")
		println(err.Error())
	}

	err = spawners.ObjectSpawner.Spawn("plant", "tomato", 150, 150)

	if err != nil {
		print("here")
		println(err.Error())
	}

	err = spawners.ObjectSpawner.Spawn("plant", "tomato", 220, 150)

	if err != nil {
		print("here")
		println(err.Error())
	}

	objects.SpawnWall("w_1", "./wall.png", 270, 270)
	objects.SpawnWall("w_2", "./wall.png", 270, 202)
	objects.SpawnWall("w_3", "./wall.png", 270, 234)
	objects.SpawnWall("w_4", "./wall.png", 270, 266)
	objects.SpawnWall("w_5", "./wall.png", 125, 66)
}

type Game struct {
	keys []ebiten.Key
}

func (g *Game) Update() error {
	managers.WeatherManager.OnUpdate()
	managers.CharacterManager.OnUpdate()
	managers.ObjectManager.OnUpdate()

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	managers.KeypressManager.OnUpdate(g.keys)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	maps.CurrMap.Draw(screen)
	// managers.DebugManager.Draw(screen)
	managers.CharacterManager.Draw(screen)
	managers.ObjectManager.Draw(screen)
	managers.WeatherManager.DrawLighting(screen)
	managers.WeatherManager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 320
}

func main() {
	ebiten.SetWindowSize(960, 960)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
