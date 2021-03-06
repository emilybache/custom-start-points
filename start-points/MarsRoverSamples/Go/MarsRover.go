package marsrover

type Coordinates struct {
    x int
    y int
}

type Direction int

const (
    N Direction = iota
    E
    S
    W
)
func (d Direction) String() string {
    return [...]string{"N", "E", "S", "W"}[d]
}

type Command int

const (
    B Command = iota
    F
    L
    R
)

func (c Command) String() string {
    return [...]string{"B", "F", "L", "R"}[c]
}

type Obstacle struct {
    position Coordinates
}

type Plateau struct {
    maxX int
    maxY int
    obstacles []Obstacle
}

type Status int

const (
    OK Status = iota
    NOK
)
func (s Status) String() string {
    return [...]string{"OK", "NOK"}[s]
}

type MarsRover struct {
    plateau Plateau
    heading Direction
    position Coordinates
    status Status
}

func (r MarsRover) turnLeft() {

}

func (r MarsRover) currentLocation() interface{} {
    return ""
}

func (r MarsRover) acceptCommands(commands []Command) {

}

func (r MarsRover) coordinates() Coordinates {
    return Coordinates{0, 0}
}

func (r MarsRover) forward() {

}

func (r MarsRover) backward() {

}

func (r MarsRover) turnRight() {

}