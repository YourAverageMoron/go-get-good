package amazeing

import "fmt"


type MazeResult struct {
	maze *Maze
	path []Coordinate
}

func RecurseMaze(result *MazeResult, seenLocations map[Coordinate]bool) *MazeResult {
	maze := result.maze
	maze.Display()
	seenLocations[maze.position] = true
	if maze.end.X == maze.position.X && maze.end.Y == maze.position.Y {
		return result
	}
	if (!seenLocations[Coordinate{X: maze.position.X - 1, Y: maze.position.Y}] && maze.MoveLeft()) ||
		(!seenLocations[Coordinate{X: maze.position.X, Y: maze.position.Y - 1}] && maze.MoveUp()) ||
		(!seenLocations[Coordinate{X: maze.position.X + 1, Y: maze.position.Y}] && maze.MoveRight()) ||
		(!seenLocations[Coordinate{X: maze.position.X, Y: maze.position.Y + 1}] && maze.MoveDown()) {
		result.path = append(result.path, maze.position)
		return RecurseMaze(result, seenLocations)
	}
	var coordinate Coordinate
	coordinate, result.path = result.path[len(result.path)-1], result.path[:len(result.path)-1]
	result.maze.MoveToCoordinate(coordinate)
	return RecurseMaze(result, seenLocations)
}

func SolveMaze(maze *Maze) {
    result := RecurseMaze(&MazeResult{maze: maze, path: []Coordinate{}}, make(map[Coordinate]bool))
    fmt.Println(result.path)
}
