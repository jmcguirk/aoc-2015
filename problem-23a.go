package main

type Problem23A struct {

}

func (this *Problem23A) Solve() {

	Log.Info("Starting Problem 23A");
	sys := &IntcodeMachine{};
	err := sys.Init("source-data/input-day-23a.txt");
	if(err != nil){
		Log.FatalError(err);
	}
	err = sys.Execute();
	if(err != nil){
		Log.FatalError(err);
	}
	registerOfInterest := 'b';
	Log.Info("Value of register %c is %d", registerOfInterest, sys.GetRegisterValue(int(registerOfInterest)));
}