package main

type Problem24B struct {

}

func (this *Problem24B) Solve() {

	Log.Info("Starting Problem 24B");
	sys := &WeightOptimizationProblem{};
	err := sys.Init("source-data/input-day-24b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	sys.OptimizeWithKnownGood(true, 5, 224551227);
}