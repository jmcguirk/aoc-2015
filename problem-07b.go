package main

type Problem7B struct {

}

func (this *Problem7B) Solve() {
	Log.Info("Problem 7B solver beginning!")

	circuit := &CircuitSystem{};
	err := circuit.Init("source-data/input-day-07b.txt");
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