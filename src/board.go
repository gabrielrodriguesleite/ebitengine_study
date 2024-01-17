package game

import "github.com/hajimehoshi/ebiten/v2"

type Board struct {
	size  int
	tiles map[*Tile]struct{}
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	boardImage.Fill(frameColor)

	nonAnimatingTiles := map[*Tile]struct{}{}
	for t := range b.tiles {
		nonAnimatingTiles[t] = struct{}{}
	}

	for t := range nonAnimatingTiles {
		t.Draw(boardImage)
	}

}

func (b *Board) Update() error {
	return nil
}

func (b *Board) Size() (int, int) {
	x := b.size*tileSize + (b.size+1)*tileMargin
	y := x
	return x, y
}

func NewBoard(size int) (*Board, error) {
	b := &Board{
		size:  size,
		tiles: map[*Tile]struct{}{},
	}

	for i := 0; i < 2; i++ {
		if err := addRandomTile(b.tiles, b.size); err != nil {
			return nil, err
		}
	}

	return b, nil
}
