package main

type Problem19B struct {

}

func (this *Problem19B) Solve() {

	Log.Info("Starting Problem 19B");
	sys := &ChemicalReactionSystem{};
	//err := sys.Init("source-data/input-day-19a.txt");
	err := sys.Init("source-data/input-day-19b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	Log.Info("Parsed reaction system - %d rules", len(sys.Rules));
	Log.Info("Initial state %s", sys.InitialState);
	steps := sys.CalculateMinStepsToGenerateDesiredCompound();
	Log.Info("Steps required %d", steps);
}