package main

import (
	"bufio"
	"os"
	"strings"
)

type Problem3B struct {

}

func (this *Problem3B) Solve() {
	Log.Info("Problem 3B solver beginning!")

	grid := &IntegerGrid2D{};
	grid.Init();

	curPos := &IntVec2{};
	roboPos := &IntVec2{};

	file, err := os.Open("source-data/input-day-03b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currVal int = 1;
	grid.SetValue(curPos.X, curPos.Y, 1);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		for i, c := range line{
			agentPos := curPos;
			if(i % 2 == 1){
				agentPos = roboPos;
			}
			switch(int(c)){
			case AsciiNorth:
				agentPos.Y--;
				break;
			case AsciiSouth:
				agentPos.Y++;
				break;
			case AsciiWest:
				agentPos.X--;
				break;
			case AsciiEast:
				agentPos.X++;
				break;
			}
			if(!grid.IsVisited(agentPos.X, agentPos.Y)){
				grid.SetValue(agentPos.X, agentPos.Y, 1);
				currVal++;
			}
		}
	}
	Log.Info("Finished simulation, santa visited %d houses", currVal);
}
