package main

import (
	"fmt"

	"github.com/rthornton128/goncurses"
)

func (g *GameState) drawWindow() {
	// Draw level
	for y := 0; y < LevelHeight; y++ {
		for x := 0; x < LevelWidth; x++ {
			var ch rune
			var attr goncurses.Char

			switch CellType(g.Level[y][x]) {
			case CellBlank:
				ch = ' '
				g.Win.ColorOn(int16(ColorNormal))
			case CellWall:
				ch = ' '
				g.Win.ColorOn(int16(ColorWall))
			case CellPellet:
				ch = '.'
				g.Win.ColorOn(int16(ColorPellet))
			case CellPowerup:
				ch = '*'
				g.Win.ColorOn(int16(ColorPowerup))
				attr = goncurses.A_BOLD
			case CellGhostWall:
				ch = ' '
				g.Win.ColorOn(int16(ColorGhostWall))
			}
			g.Win.MoveAddChar(y, x, goncurses.Char(ch)|attr)
		}
	}

	g.updateStatus()

	// Draw ghosts
	if !g.Invincible {
		for i := range g.Ghosts {
			g.Win.ColorOn(g.Ghosts[i].Color)
			g.Win.MoveAddChar(g.Ghosts[i].Pos.Y, g.Ghosts[i].Pos.X, goncurses.Char(g.Ghosts[i].Char))
		}
	} else {
		// Draw vulnerable ghosts
		g.Win.ColorOn(int16(ColorBlueGhost))
		timeChar := goncurses.Char('0' + rune(g.TimeLeft))
		for i := range g.Ghosts {
			g.Win.MoveAddChar(g.Ghosts[i].Pos.Y, g.Ghosts[i].Pos.X, timeChar)
		}
	}

	// Draw Pacman
	g.Win.ColorOn(g.Pacman.Color)
	g.Win.MoveAddChar(g.Pacman.Pos.Y, g.Pacman.Pos.X, goncurses.Char(g.Pacman.Char))

	g.Win.Refresh()
}

func (g *GameState) updateStatus() {
	g.Status.Move(0, 0)
	g.Status.ClearToEOL()

	// Lives and score on same line
	g.Status.ColorOn(int16(ColorPacman))
	for i := 0; i < g.Lives; i++ {
		g.Status.Print("P ")
	}
	g.Status.ColorOn(int16(ColorNormal))
	g.Status.Print(fmt.Sprintf("     Score: %d", g.Points))
	g.Status.Refresh()
}

func (g *GameState) clearStatus() {
	g.Status.Move(0, 0)
	g.Status.ClearToEOL()
	g.updateStatus()
}
