package main

import "github.com/rthornton128/goncurses"

// Game constants
const (
	LevelHeight = 29
	LevelWidth  = 28
	MaxGhosts   = 4
)

// CellType represents the type of each cell in the game grid
type CellType int

const (
	CellBlank CellType = iota
	CellWall
	CellPellet
	CellPowerup
	CellGhostWall
	CellBlinky
	CellInkey
	CellClyde
	CellPinky
	CellPacmanStart
)

// Color constants
const (
	ColorNormal = iota + 1
	ColorWall
	ColorPellet
	ColorPowerup
	ColorGhostWall
	ColorGhost1
	ColorGhost2
	ColorGhost3
	ColorGhost4
	ColorBlueGhost
	ColorPacman
	ColorCursor
)

// Position represents a coord in the game grid
type Position struct {
	Y, X int
}

// Direction represents movement direction
type Direction struct {
	Y, X int
}

// Entity represents a game character (Pacman/Ghost)
type Entity struct {
	Pos      Position
	Dir      Direction
	StartPos Position
	Color    int16
	Char     rune
}

// GameState holds the entire state of the game
type GameState struct {
	Level      [LevelHeight][LevelWidth]int
	Pacman     Entity
	Ghosts     [MaxGhosts]Entity
	PendingDir Direction

	Invincible   bool
	Food         int
	LevelNumber  int
	CurrentLevel int
	GhostsInARow int
	TimeLeft     int
	Points       int
	Lives        int
	FreeLife     int
	HowSlow      int
	SpeedOfGame  int
	TickCounter  int

	Win    *goncurses.Window
	Status *goncurses.Window
	HighScore   int // Track high score
}
