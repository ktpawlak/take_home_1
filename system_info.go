package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const printJSON = false

func parseOutput(str string) map[string]string {
	result := make(map[string]string)

	re := regexp.MustCompile(`([0-9\.s]+)\s\((kernel|userspace)\)`)
	match := re.FindAllStringSubmatch(str, -1)

	if len(match) != 0 {
		for _, el := range match {
			result[el[2]] = el[1]
		}
	}
	return result
}

func printOutput(w http.ResponseWriter, in map[string]string) {
	if printJSON {
		j, err := json.Marshal(in)
		if err == nil {
			fmt.Fprintf(w, "%s\n", string(j))
		}
	} else {
		for key := range in {
			fmt.Fprintf(w, "%s: %s\n", key, in[key])
		}
	}
}

func durationHandler(w http.ResponseWriter, req *http.Request) {
	result := make(map[string]string)

	cmd := exec.Command("systemd-analyze")
	out, err := cmd.CombinedOutput()

	if err != nil {
		result["error"] = "Failed to handle request"
	} else {
		result = parseOutput(string(out))
		if len(result) != 2 {
			result["error"] = "Failed to parse startup timings"
		}

		total := 0.0
		for key := range result {
			val, _ := strconv.ParseFloat(strings.TrimSuffix(result[key], "s"), 64)
			total = total + val
		}
		result["total"] = fmt.Sprintf("%.3fs", total)
	}
	printOutput(w, result)
}

func helpHandler(w http.ResponseWriter, req *http.Request) {
	var help = map[string]string{"/version": "returns server version", "/duration": "returns startup duration"}

	printOutput(w, help)
}

func versionHandler(w http.ResponseWriter, req *http.Request) {
	var version = map[string]string{"version": "1.0.0"}

	printOutput(w, version)
}

func main() {
	http.HandleFunc("/", helpHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/duration", durationHandler)

	fmt.Println("Server ready, endpoints: /version and /duration")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("server failure: " + err.Error())
	}
}
