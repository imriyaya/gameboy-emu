package pkg

const VramBegin uint16 = 0x8000
const VramEnd uint16 = 0x9FFF
const VramSize uint16 = VramEnd - VramBegin + 1
const VramTileBegin uint16 = 0x8000
const VramTileEnd uint16 = 0x97FF
const VramTileSize uint16 = VramTileEnd - VramTileBegin + 1
const VramTileMap0Begin uint16 = 0x9800
const VramTileMap1Begin uint16 = 0x9C00

const LCDCAddress uint16 = 0xFF40
const LCDCEnableBytePosition = 7
const LCDCWindowTileMapAreaBytePosition = 6
const LCDCWindowEnableBytePosition = 5
const LCDCBGWindowTileDataAreaBytePosition = 4
const LCDCGBTileMapAreaBytePosition = 3
const LCDCOBJSizeBytePosition = 2
const LCDCOBJEnableBytePosition = 1
const LCDCBGAndWindowEnablePriorityBytePosition = 0

const SCYAddress uint16 = 0xFF42
const SCXAddress uint16 = 0xFF43

type Palett uint8

const (
	Black     Palett = iota
	LightGray Palett = iota
	DarkGray  Palett = iota
	White     Palett = iota
)

type Tile [8][8]Palett
type TileMap [32][32]Tile
type BG [20][18]Tile

type PPU struct {
	gameboy *Gameboy
	Vram    [VramSize]uint8
}

func (ppu *PPU) readWindowTileMapArea() bool {
	return ppu.gameboy.memory.GetValue(LCDCAddress)>>LCDCWindowTileMapAreaBytePosition&0b1 != 0
}

func (ppu *PPU) readBGWindowTileDataArea() bool {
	return ppu.gameboy.memory.GetValue(LCDCAddress)>>LCDCBGWindowTileDataAreaBytePosition&0b1 != 0
}

func (ppu *PPU) readBGTileMapArea() bool {
	return ppu.gameboy.memory.GetValue(LCDCAddress)>>LCDCGBTileMapAreaBytePosition&0b1 != 0
}

func (ppu *PPU) readTile(id uint8, addressMode bool) Tile {
	if addressMode {
		return ppu.readTileByAddress(VramTileBegin + uint16(id*0xF))
	} else {
		if id <= 127 {
			return ppu.readTileByAddress(VramTileBegin + uint16(id*0xF) + 0x1000)
		} else {
			return ppu.readTileByAddress(VramTileBegin + uint16(id*0xF))
		}
	}
}

func (ppu *PPU) readTileByAddress(address uint16) Tile {
	tile := Tile{}
	for x := 0; x < 8; x++ {
		bottomByte := ppu.gameboy.memory.GetValue(address)
		topByte := ppu.gameboy.memory.GetValue(address + 1)
		for y := 0; y < 8; y++ {
			bottomBit := (bottomByte >> (7 - x) & 0b1) != 0
			topBit := (topByte >> (7 - x) & 0b1) != 0
			switch true {
			case !topBit && !bottomBit:
				tile[x][y] = Black
			case topBit && !bottomBit:
				tile[x][y] = LightGray
			case !topBit && bottomBit:
				tile[x][y] = DarkGray
			case topBit && bottomBit:
				tile[x][y] = White
			}
		}
		address += 2
	}
	return tile
}

func (ppu *PPU) readTileMap(mapArea bool, tileDataArea bool) TileMap {
	address := VramTileMap0Begin
	if mapArea {
		address = VramTileMap1Begin
	}
	tileMap := TileMap{}

	for x := 0; x < 32; x++ {
		for y := 0; y < 32; y++ {
			tileMap[x][y] = ppu.readTile(ppu.gameboy.memory.GetValue(address), tileDataArea)
			address += 1
		}
	}

	return tileMap
}

func (ppu *PPU) readBG() BG {
	tileMap := ppu.readTileMap(ppu.readBGTileMapArea(), ppu.readBGWindowTileDataArea())
	yBegin := ppu.gameboy.memory.GetValue(SCYAddress)
	xBegin := ppu.gameboy.memory.GetValue(SCXAddress)
	bg := BG{}

	for x := 0; x < 20; x++ {
		for y := 0; y < 18; y++ {
			bg[uint8(x)][uint8(y)] = tileMap[xBegin+uint8(x)][yBegin+uint8(y)]
		}
	}
	return bg
}
