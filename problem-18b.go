package main

type Problem18B struct {

}

func (this *Problem18B) Solve() {
	Log.Info("Problem 18B solver beginning!")

	grid := IntegerGrid2D{};
	grid.Init();
	err := grid.ParseAsciiGrid("source-data/input-day-18b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	const LightOn = int('#');
	const LightOff = int('.');

	maxGen := 100;

	minX := grid.MinX();
	maxX := grid.MaxX();

	minY := grid.MinY();
	maxY := grid.MaxY();

	grid.SetValue(minX, minY, LightOn);
	grid.SetValue(maxX, minY, LightOn);
	grid.SetValue(minX, maxY, LightOn);
	grid.SetValue(maxX, maxY, LightOn);

	currGen := 0;
	for {
		if(currGen >= maxGen){
			break;
		}
		prevGen := grid.Clone();
		for i:= minX; i <= maxX; i++{
			for j:= minY; j <= maxY; j++{
				if(i == minX && j == minY){
					continue;
				}
				if(i == maxX && j == minY){
					continue;
				}
				if(i == minX && j == maxY){
					continue;
				}
				if(i == maxX && j == maxY){
					continue;
				}

				neighbors := prevGen.Count8DirNeighborsMatching(i, j, LightOn);
				if(prevGen.GetValue(i, j) == LightOn){
					if(neighbors != 2 && neighbors != 3){
						grid.SetValue(i, j, LightOff);
					}
				} else{
					if(neighbors == 3){
						grid.SetValue(i, j, LightOn);
					}
				}
			}
		}
		currGen++;
	}

	lights := 0;
	for i:= minX; i <= maxX; i++{
		for j:= minY; j <= maxY; j++{
			if(grid.GetValue(i, j) == LightOn){
				lights++;
			}
		}
	}
	Log.Info("Light count is %d after %d steps", lights, maxGen);

}