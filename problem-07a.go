package main

type Problem7A struct {

}

func (this *Problem7A) Solve() {
	Log.Info("Problem 7A solver beginning!")

	circuit := &CircuitSystem{};
	err := circuit.Init("source-data/input-day-07a.txt");
	if(err != nil){
		Log.FatalError(err);
	}

	//Log.Info("Initial layout\n" + circuit.PrintCircuitLayout());


	/*
	for label, _ := range circuit.Nodes{
		Log.Info("%s: %d", label, circuit.CalculateSignal(label));
	}*/

	label := "a";
	Log.Info("%s: %d", label, circuit.CalculateSignal(label));


	/*
	label := "x";




	//Log.Info("Christmas light grid has total %d brightness", lights);*/
}