package main

type Problem17B struct {

}

func (this *Problem17B) Solve() {

	Log.Info("Starting Problem 17B");
	change := &ChangeMakingSystem{};
	change.Init();
	err := change.LoadDenominations("source-data/input-day-17b.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	Log.Info("Loaded %d containers", len(change.Denominations));
	amountDesired := 150;

	uniqueWays := change.FindMinWaysToMakeChange(amountDesired);

	Log.Info("Counted %d min ways to make %d", uniqueWays, amountDesired);
}