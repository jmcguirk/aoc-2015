package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem6B struct {

}

func (this *Problem6B) Solve() {
	Log.Info("Problem 6B solver beginning!")

	grid := IntegerGrid2D{};
	grid.Init();
	gridSize := 1000;

	for i := 0; i < gridSize; i++{
		for j := 0; j < gridSize; j++{
			grid.SetValue(i, j, 0);
		}
	}

	file, err := os.Open("source-data/input-day-06a.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//sum := int64(0);
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			lineParts := strings.Split(line, " ");
			op := lineParts[0];
			pivot := 0;
			if(op != "toggle"){
				pivot++;
			}
			start := lineParts[1 + pivot];
			end := lineParts[1 + pivot+2];
			//Log.Info("%s, start -%s, end %s", op, start, end);

			startParts := strings.Split(start, ",");
			startX, _ := strconv.ParseInt(startParts[0], 10, 64);
			startY, _ := strconv.ParseInt(startParts[1], 10, 64);

			endParts := strings.Split(end, ",");
			endX, _ := strconv.ParseInt(endParts[0], 10, 64);
			endY, _ := strconv.ParseInt(endParts[1], 10, 64);

			for i := int(startX); i <= int(endX); i++{
				for j := int(startY); j <= int(endY); j++{
					curr := grid.GetValue(i, j);
					if(op == "toggle"){
						grid.SetValue(i, j, curr+2)
					} else if(op == "turn"){
						if(lineParts[1] == "on"){
							grid.SetValue(i, j, curr+1);
						} else if(lineParts[1] == "off"){
							if(curr >= 1){
								grid.SetValue(i, j, curr-1)
							}
						}
					}
				}
			}
		}
	}

	minX := grid.MinX();
	maxX := grid.MaxX();

	minY := grid.MinY();
	maxY := grid.MaxY();

	lights := 0;
	for i:= minX; i <= maxX; i++{
		for j:= minY; j <= maxY; j++{
			lights += grid.GetValue(i, j);
		}
	}
	Log.Info("Christmas light grid has total %d brightness", lights);
}