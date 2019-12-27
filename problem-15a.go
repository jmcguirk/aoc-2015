package main

type Problem15A struct {

}

func (this *Problem15A) Solve() {

	Log.Info("Starting Problem 15 A");
	kitchen := &KitchenOptimization{};
	err := kitchen.Init("source-data/input-day-15a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	best := kitchen.Optimize(100);
	Log.Info("Best score is %d", best);
	for k, v := range kitchen.BestAllocation{
		Log.Info("%s - %d", k, v);
	}
}