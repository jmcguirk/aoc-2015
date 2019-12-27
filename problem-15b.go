package main

type Problem15B struct {

}

func (this *Problem15B) Solve() {

	Log.Info("Starting Problem 15 B");
	kitchen := &KitchenOptimization{};
	err := kitchen.Init("source-data/input-day-15b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	best := kitchen.OptimizeWithCalories(100, 500);
	Log.Info("Best score is %d", best);
	for k, v := range kitchen.BestAllocation{
		Log.Info("%s - %d", k, v);
	}
}