package main

type Problem19A struct {

}

func (this *Problem19A) Solve() {

	Log.Info("Starting Problem 19A");
	sys := &ChemicalReactionSystem{};
	//err := sys.Init("source-data/input-day-19a.txt");
	err := sys.Init("source-data/input-day-19a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	Log.Info("Parsed reaction system - %d rules", len(sys.Rules));
	Log.Info("Initial state %s", sys.InitialState);
	Log.Info("Unique products %d", sys.CountUniqueOutputs());
}