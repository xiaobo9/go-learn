package phoneArea

import (
	"os"
	"testing"
)

func TestBuildArea(*testing.T) {
	name := "areaForIgnore.csv"
	os.Remove(name)
	var areas = BuildArea()
	ToFile(areas, name)
}
