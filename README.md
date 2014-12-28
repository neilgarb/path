# path

Calculate the shorted path between two points on a Cartesian plane, given a set
of walls or barriers, defined as line segments.

## Usage

```
import (
    "fmt"
    "github.com/NeilGarb/path"
)

s := path.NewSpace(nil)
s.AddWall(path.Wall{path.Point{0, 0}, path.Point{10, 10})
shortestPath, err := s.ShortestPath(path.Point{2, 5}, path.Point{5, 2})
if err != nil {
    panic(err)
}
fmt.Printf("%v\n", shortestPath) // which is a []path.Point
```

_NB:_ This project is still under development.  You may use what code is here
in any way you choose, but you do so at your own risk.
