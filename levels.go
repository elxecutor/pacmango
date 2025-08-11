package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (g *GameState) loadLevel(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("loading level %s: %w", filename, err)
	}
	defer file.Close()

	// Reset game state
	g.Food = 0
	g.PendingDir = Direction{}

	// Initialize Pacman
	g.Pacman.Char = 'C'
	g.Pacman.Color = int16(ColorPacman)
	g.Pacman.Dir = Direction{0, -1}

	// Initialize ghosts
	ghostChars := []rune{'&', '&', '&', '&'}
	ghostColors := []int16{
		int16(ColorGhost1), int16(ColorGhost2),
		int16(ColorGhost3), int16(ColorGhost4),
	}

	for i := range g.Ghosts {
		g.Ghosts[i].Char = ghostChars[i]
		g.Ghosts[i].Color = ghostColors[i]
	}

	scanner := bufio.NewScanner(file)

	for y := 0; y < LevelHeight && scanner.Scan(); y++ {
		line := scanner.Text()
		fields := strings.Fields(line)

		for x := 0; x < LevelWidth && x < len(fields); x++ {
			val, err := strconv.Atoi(fields[x])
			if err != nil {
				continue
			}

			g.Level[y][x] = val

			switch CellType(val) {
			case CellPellet:
				g.Food++
			case CellBlinky:
				g.Ghosts[0].Pos = Position{y, x}
				g.Ghosts[0].Dir = Direction{1, 0}
				g.Level[y][x] = int(CellBlank)
			case CellInkey:
				g.Ghosts[1].Pos = Position{y, x}
				g.Ghosts[1].Dir = Direction{-1, 0}
				g.Level[y][x] = int(CellBlank)
			case CellClyde:
				g.Ghosts[2].Pos = Position{y, x}
				g.Ghosts[2].Dir = Direction{0, -1}
				g.Level[y][x] = int(CellBlank)
			case CellPinky:
				g.Ghosts[3].Pos = Position{y, x}
				g.Ghosts[3].Dir = Direction{0, 1}
				g.Level[y][x] = int(CellBlank)
			case CellPacmanStart:
				g.Pacman.Pos = Position{y, x}
				g.Level[y][x] = int(CellBlank)
			}
		}
	}

	// Read level number
	if scanner.Scan() {
		if num, err := strconv.Atoi(strings.TrimSpace(scanner.Text())); err == nil {
			g.LevelNumber = num
		}
	}

	// Save starting positions
	g.Pacman.StartPos = g.Pacman.Pos
	for i := range g.Ghosts {
		g.Ghosts[i].StartPos = g.Ghosts[i].Pos
	}

	return nil
}
