package main

func main() {
	Log.Init();
	Log.Info("Starting up AOC 2015");

	solver := Problem10B{};

	solver.Solve();
	Log.Info("Solver complete - exiting");
}
