package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem3A struct {

}

func (this *Problem3A) Solve() {
	Log.Info("Problem 3A solver beginning!")

	grid := &IntegerGrid2D{};
	grid.Init();

	curPos := IntVec2{};

	file, err := os.Open("source-data/input-day-03a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currVal int = 1;
	grid.SetValue(curPos.X, curPos.Y, 1);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		for _, c := range line{
			switch(int(c)){
				case AsciiNorth:
					curPos.Y--;
					break;
				case AsciiSouth:
					curPos.Y++;
					break;
				case AsciiWest:
					curPos.X--;
					break;
				case AsciiEast:
					curPos.X++;
					break;
			}
			if(!grid.IsVisited(curPos.X, curPos.Y)){
				grid.SetValue(curPos.X, curPos.Y, 1);
				currVal++;
			}
		}
	}
	Log.Info("Finished simulation, santa visited %d houses", currVal);
}
