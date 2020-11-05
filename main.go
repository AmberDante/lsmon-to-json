package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	exitOK           = 0
	exitFail         = 1
	introduction     = "[Contacting Sentinel RMS Development Kit"
	ending           = "Press Enter to continue . . ."
	featureDelimeter = " |- Feature Information"
	licDelimeter     = "   |- License Information"
	clientDelimeter  = "   |- Client Information"
)

type fetures struct {
	Features []feature
}

type feature struct {
	feature            map[string]string
	LicenseInformation []licenseInformation
	ClientInformation  []clientInformation
}

type licenseInformation struct {
	LicenseInformation map[string]string
}

type clientInformation struct {
	ClientInformation map[string]string
}

func main() {
	err := run(os.Stdin, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
	os.Exit(exitOK)
}

func run(stdin io.Reader, stdout io.Writer) error {
	scanner := bufio.NewScanner(stdin)
	var lsmonOut string
	var preJSON fetures
	for scanner.Scan() {
		lsmonOut = lsmonOut + scanner.Text() + "\n"
	}
	preJSON = getFeturesInfo(lsmonOut)
	// TODO !!! marshal JSON struct to JSON object
	JSONtoOUT, err := createJSON(preJSON)
	if err != nil {
		return err
	}
	fmt.Fprint(stdout, string(JSONtoOUT))

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func createJSON(s fetures) (string, error) {
	jsonM, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonM), nil
}

// trimBorders - trim unnecessary information
func trimBorders(s string) string {
	i1 := strings.Index(s, introduction)
	if i1 >= 0 {
		i1 = i1 + len(introduction)
		i1 = strings.Index(s[i1:], "\n\n\n") + i1 + 3
		s = s[i1:]
	}
	i1 = strings.LastIndex(s, ending)
	if i1 >= 0 {
		s = s[:i1]
	}
	s = strings.Trim(s, "\n ")
	return s
}

func splitdata(s string, sep string) []string {
	if sep == "" {
		var onestring []string
		onestring = append(onestring, s)
		return onestring
	}
	return strings.Split(s, sep)
}

// getFeturesInfo - get Feature Information
func getFeturesInfo(lsmonOut string) fetures {
	var feats fetures
	// var feturesInfo []string
	lsmonOut = trimBorders(lsmonOut)
	// feturesInfo = splitdata(lsmonOut, featureDelimeter)

	// Split data for features blocks
	// for i, data := range feturesInfo {

	// }

	return feats
}

// parseFeature - parse Feature Information
func parseFeature(s string) feature {
	var feat feature

	return feat
}

// getLicenseInformation - get License Information
func getLicenseInformation(s string) licenseInformation {
	var lics licenseInformation

	return lics
}

// getClientInformation - get Client Information
func getClientInformation(s string) clientInformation {
	var clinfo clientInformation

	return clinfo
}
