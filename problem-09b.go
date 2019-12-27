package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem9B struct {

}

func (this *Problem9B) Solve() {
	Log.Info("Problem 9B solver beginning!")


	file, err := os.Open("source-data/input-day-09b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	graph := &UndirectedGraph{};
	graph.Init();

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		if(line != ""){
			parts := strings.Split(line, "=");
			weight, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64);
			if(err != nil){
				Log.FatalError(err);
			}
			routeParts := strings.Split(parts[0], " to ");
			fromLabel := strings.TrimSpace(routeParts[0]);
			toLabel := strings.TrimSpace(routeParts[1]);

			fromNode := graph.GetOrCreateNode(fromLabel);
			toNode := graph.GetOrCreateNode(toLabel);
			graph.CreateEdgeWithWeight(fromNode, toNode, int(weight));
		}
	}

	Log.Info("Finished parsing graph - %d total nodes, %d total edges", len(graph.Nodes), len(graph.Edges));

	cycle := graph.LongestCycle();
	if(cycle == nil){
		Log.Fatal("Failed to find a cycle");
	}
	Log.Info("Best cycle is of length %d", cycle.BestSoFar)

	//destinationOfInterest := ""
}
