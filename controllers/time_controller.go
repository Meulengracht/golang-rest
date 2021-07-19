package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
)

// represents the available boot times we show on the endpoint
type bootTimes struct {
	Kernel    string `json:"kernel"`
	Userspace string `json:"userspace"`
}

// helper function to set the http code on errors and serialize to json
func writeResponse(w http.ResponseWriter, response interface{}, err error) {
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(response)
}

// helper function to run commands and return the stdout
func runCommand(cmd string) (string, error) {
	cmdResult, err := exec.Command(cmd).Output()
	return string(cmdResult), err
}

// helper function that parses the stdout from the command and
// runs it through our regex expr to extract boot times for kernel and userspace
func getBootTimes(output string) (bootTimes, error) {
	var bootTimes bootTimes
	var timesRegex = regexp.MustCompile(`([0-9\.]+)s\ \((kernel|userspace)\)`)

	// expect exactly two matches, otherwise something has gone wrong
	// maybe the regex expression is wrong then
	matches := timesRegex.FindAllStringSubmatch(output, -1)
	if len(matches) != 2 {
		return bootTimes, fmt.Errorf("Could not parse boot times, the number of matches were not as expected: " + output)
	}

	// index 0 is the full match, index 1 is the first capture group (time)
	// and index 2 is the second capture group (identifier)
	if matches[0][2] == "kernel" {
		bootTimes.Kernel = matches[0][1] + "s"    // kernel time
		bootTimes.Userspace = matches[1][1] + "s" // userspace time
	} else {
		bootTimes.Kernel = matches[1][1] + "s"    // kernel time
		bootTimes.Userspace = matches[0][1] + "s" // userspace time
	}
	return bootTimes, nil
}

func GetStartupTimingInfo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	output, err := runCommand("systemd-analyze")
	if err != nil {
		writeResponse(w, "Failed to extract boot times on the underlying OS", err)
		return
	}

	bootTimes, err := getBootTimes(strings.ToLower(output))
	if err != nil {
		writeResponse(w, "Failed to parse boot times", err)
		return
	}
	writeResponse(w, bootTimes, err)
}
