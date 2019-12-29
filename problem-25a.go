package main

type Problem25A struct {

}

func (this *Problem25A) Solve() {

	Log.Info("Starting Problem 25A");

	desiredRow := 2981;
	desiredCol := 3075;

	anchorCol := desiredCol + desiredRow - 1;
	Log.Info("Diag starts at col %d", anchorCol)
	anchorColVal := TriangularTerm(anchorCol);
	Log.Info("Calculated anchor col term %d", anchorColVal)
	termNum := anchorColVal - desiredRow + 1;
	Log.Info("row %d, col %d is term number %d", desiredRow, desiredCol, termNum);

	start := 20151125;

	next := start;
	for i := 1; i < termNum; i++{
		prod := (next * 252533)/33554393;
		next = (next * 252533) - (prod * 33554393);
	}
	Log.Info("Calculated row %d, col %d, value - %d", desiredRow, desiredCol, next);
}