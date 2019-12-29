package main

type Problem24A struct {

}

func (this *Problem24A) Solve() {

	Log.Info("Starting Problem 23A");
	sys := &WeightOptimizationProblem{};
	err := sys.Init("source-data/input-day-24a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	sys.Optimize(false);
}