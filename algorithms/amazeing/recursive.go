package amazeing

import "fmt"

type MazeResult struct {
	maze          *Maze
	seenLocations map[Coordinate]bool
}

func RecurseMaze(maze *Maze, seenLocations map[Coordinate]bool) *Maze {
    fmt.Println(seenLocations)
	maze.Display()
	seenLocations[maze.position] = true
	if maze.end.X == maze.position.X && maze.end.Y == maze.position.Y {
		return maze
	}
	if (!seenLocations[Coordinate{X: maze.position.X - 1, Y: maze.position.Y}] && maze.MoveLeft()) ||
		(!seenLocations[Coordinate{X: maze.position.X, Y: maze.position.Y - 1}] && maze.MoveUp()) ||
		(!seenLocations[Coordinate{X: maze.position.X + 1, Y: maze.position.Y}] && maze.MoveRight()) ||
		(!seenLocations[Coordinate{X: maze.position.X, Y: maze.position.Y + 1}] && maze.MoveDown()) {
		return RecurseMaze(maze, seenLocations)
	}
	panic("something went wrong")
}

func SolveMaze(maze *Maze) {
	RecurseMaze(maze, make(map[Coordinate]bool))
}
