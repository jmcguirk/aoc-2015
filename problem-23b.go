package main

type Problem23B struct {

}

func (this *Problem23B) Solve() {

	Log.Info("Starting Problem 23B");
	sys := &IntcodeMachine{};
	err := sys.Init("source-data/input-day-23b.txt");
	sys.SetRegisterValue(int('a'), 1);
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