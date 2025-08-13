package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func (g *GameState) loadLevel(levelNum int) error {
	filename := fmt.Sprintf("levels/level%d.txt", levelNum)

	// Use embedded file system
	data, err := levelsFS.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("loading level %d: %w", levelNum, err)
	}

	return g.parseLevel(string(data))
}

func (g *GameState) parseLevel(levelData string) error {
	reader := strings.NewReader(levelData)
	scanner := bufio.NewScanner(reader)

	// Reset game state for new level
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

	// Parse level data
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
		} else {
			g.LevelNumber = g.CurrentLevel
		}
	} else {
		g.LevelNumber = g.CurrentLevel
	}

	// Save starting positions
	g.Pacman.StartPos = g.Pacman.Pos
	for i := range g.Ghosts {
		g.Ghosts[i].StartPos = g.Ghosts[i].Pos
	}

	return nil
}
