package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2015");

	solver := Problem20A{};

	solver.Solve();
	Log.Info("Solver complete - exiting");
}
