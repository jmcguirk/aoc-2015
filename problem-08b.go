package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem8B struct {

}

func (this *Problem8B) Solve() {
	Log.Info("Problem 8B solver beginning!")


	file, err := os.Open("source-data/input-day-08b.txt");
	if err != nil {
		Log.FatalError(err);
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalDiskRep := 0;
	totalMemoryRep := 0;
	for scanner.Scan() {             // internally, it advances token based on sperator
		line := strings.TrimSpace(scanner.Text());
		diskLen := len(line);
		memoryRep := len(SpecialEscape(line));
		Log.Info(line + " " + SpecialEscape(line) + " : %d, %d", diskLen, memoryRep);
		totalDiskRep += diskLen;
		totalMemoryRep += memoryRep;
	}
	Log.Info("Finished parsing file %d on disk, %d in memory. Delta %d", totalDiskRep, totalMemoryRep, totalMemoryRep - totalDiskRep);
}

func SpecialEscape(str string) string {
	return strconv.Quote(str);
	//return stripSlash;
}

