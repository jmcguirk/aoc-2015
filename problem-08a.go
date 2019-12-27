package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Problem8A struct {

}

func (this *Problem8A) Solve() {
	Log.Info("Problem 8A solver beginning!")


	file, err := os.Open("source-data/input-day-08a.txt");
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
		memoryRep := len(SpecialUnescape(line));
		Log.Info(line + " " + SpecialUnescape(line) + " : %d, %d", diskLen, memoryRep);
		totalDiskRep += diskLen;
		totalMemoryRep += memoryRep;
	}
	Log.Info("Finished parsing file %d on disk, %d in memory. Delta %d", totalDiskRep, totalMemoryRep, totalDiskRep - totalMemoryRep);
}

func SpecialUnescape(str string) string {
	stripSlash := str[1:len(str) - 1];
	// We replace with surrogate characters so they don't get shuffled in subsequent runs
	stripSlash = strings.Replace(stripSlash, "\\\\", "S", -1)
	stripSlash = strings.Replace(stripSlash, "\\\"", "Q", -1)
	// Courtesy of https://stackoverflow.com/questions/24656624/golang-display-character-not-ascii-like-not-0026
	stripSlash, _ = strconv.Unquote(strings.Replace(strconv.Quote(string(stripSlash)), `\\x`, `\x`, -1))
	return stripSlash;
}

