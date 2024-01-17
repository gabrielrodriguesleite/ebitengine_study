package game

import (
	"errors"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	tileSize   = 80
	tileMargin = 4
)

var (
	tileImage = ebiten.NewImage(tileSize, tileSize)
)

func init() {
	tileImage.Fill(color.White)
}

type TileData struct {
	value int
	x     int
	y     int
}
type Tile struct {
	current TileData
}

func (t *Tile) Draw(boardImage *ebiten.Image) {
	i, j := t.current.x, t.current.y
	v := t.current.value
	if v == 0 {
		return
	}

	op := &ebiten.DrawImageOptions{}
	x := i*tileSize + (i+1)*tileMargin
	y := j*tileSize + (i+1)*tileMargin

	op.GeoM.Translate(float64(x), float64(y))
	op.ColorScale.ScaleWithColor(titleBackgroundColor(v))
	boardImage.DrawImage(tileImage, op)
	// str := strconv.Itoa(v)
	// size := 48.0
	// switch {
	// case 3 < len(str):
	// 	size = 24
	// case 2 < len(str):
	// 	size = 32
	// }

}

func NewTile(value, x, y int) *Tile {

	return &Tile{
		current: TileData{
			value: value,
			x:     x,
			y:     y,
		},
	}
}

func addRandomTile(tiles map[*Tile]struct{}, size int) error {
	cells := make([]bool, size*size)

	for t := range tiles {
		i := t.current.x + t.current.y*size
		cells[i] = true
	}

	availableCells := []int{}
	for i, b := range cells {
		if b {
			continue
		}
		availableCells = append(availableCells, i)
	}

	if len(availableCells) == 0 {
		return errors.New("2048: there is no space to add a new tile")
	}

	c := availableCells[rand.Intn(len(availableCells))]

	v := 2
	if rand.Intn(10) == 0 {
		v = 4
	}

	x := c % size
	y := c / size
	t := NewTile(v, x, y)
	tiles[t] = struct{}{}
	return nil
}
