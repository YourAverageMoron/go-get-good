package amazeing

import "fmt"

type Coordinate struct {
	X, Y int
}

type Maze struct {
	mazeMap              [][]rune
	start, end, position Coordinate
}

func (maze *Maze) MoveUp() bool {
	if maze.position.Y == 0 || maze.mazeMap[maze.position.Y-1][maze.position.X] == 35 {
		return false
	}
	maze.position.Y--
	return true
}

func (maze *Maze) MoveDown() bool {
	if maze.position.Y == len(maze.mazeMap)-1 || maze.mazeMap[maze.position.Y+1][maze.position.X] == 35 {
		return false
	}
	maze.position.Y++
	return true
}

func (maze *Maze) MoveRight() bool {
	if maze.position.X == 0 || maze.mazeMap[maze.position.Y][maze.position.X+1] == 35 {
		return false
	}
	maze.position.X++
	return true
}

func (maze *Maze) MoveLeft() bool {
	if maze.position.X == 1 || maze.mazeMap[maze.position.Y][maze.position.X-1] == 35 {
		return false
	}
	maze.position.X--
	return true
}

func (maze *Maze) MoveToCoordinate(c Coordinate) bool {
	if c.Y >= len(maze.mazeMap) ||
		c.X >= len(maze.mazeMap) ||
		maze.mazeMap[c.Y][c.X] == 35 {
		return false
	}
	maze.position = c
	return true

}

func (maze *Maze) Display() {
	for x := 0; x < len(maze.mazeMap); x++ {
		for y := 0; y < len(maze.mazeMap[x]); y++ {
			if x == maze.position.Y && y == maze.position.X {
				fmt.Print("X")
				continue
			}
			fmt.Print(string(maze.mazeMap[x][y]))
		}
		fmt.Println()
	}
}

func CreateMaze() Maze {
	start := Coordinate{1, 5}
	return Maze{
		mazeMap: [][]rune{
			{'#', '#', '#', '#', 'E', '#'},
			{'#', ' ', ' ', '#', ' ', '#'},
			{'#', ' ', '#', '#', ' ', '#'},
			{'#', ' ', '#', '#', ' ', '#'},
			{'#', ' ', ' ', ' ', ' ', '#'},
			{'#', 'S', '#', '#', '#', '#'},
		},
		start:    start,
		end:      Coordinate{4, 0},
		position: start,
	}
}
