package main

type Problem17A struct {

}

func (this *Problem17A) Solve() {

	Log.Info("Starting Problem 17A");
	change := &ChangeMakingSystem{};
	change.Init();
	err := change.LoadDenominations("source-data/input-day-17a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	Log.Info("Loaded %d containers", len(change.Denominations));
	amountDesired := 150;

	uniqueWays := change.CountWaysToMakeChange(amountDesired);

	Log.Info("Counted %d ways to make %d", uniqueWays, amountDesired);
}