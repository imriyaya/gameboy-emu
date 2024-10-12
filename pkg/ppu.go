package pkg

const VramBegin uint16 = 0x8000
const VramEnd uint16 = 0x9FFF
const VramSize uint16 = VramEnd - VramBegin + 1

type Pixel uint8

const (
	Black     Pixel = iota
	LightGray Pixel = iota
	DarkGray  Pixel = iota
	White     Pixel = iota
)

type Tile [8][8]Pixel

func EmptyTile() *Tile {
	tile := &Tile{}
	return tile
}

type PPU struct {
	Vram [VramSize]uint8
}

func (gameboy *Gameboy) readTile(address) Tile {

}
