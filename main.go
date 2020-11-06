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
	Feature             map[string]string
	LicensesInformation []licenseInformation
	ClientsInformation  []clientInformation
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

// splitFeature - split Feture Info from another stuff
func splitFeature(s string) (featreInfo, additionalInfo string) {
	i1 := strings.Index(s, licDelimeter)
	i2 := strings.Index(s, clientDelimeter)

	if i1 >= 0 && i2 >= 0 && i1 < i2 {
		featreInfo = s[:i1]
		additionalInfo = s[i1:]
	}
	if i1 >= 0 && i2 >= 0 && i1 > i2 {
		featreInfo = s[:i2]
		additionalInfo = s[i2:]
	}
	if i1 < 0 && i2 > 0 {
		featreInfo = s[:i2]
		additionalInfo = s[i2:]
	}
	if i1 > 0 && i2 < 0 {
		featreInfo = s[:i1]
		additionalInfo = s[i1:]
	}
	if i1 < 0 && i2 < 0 {
		featreInfo = s
	}
	return featreInfo, additionalInfo
}

// textToMap - Convert lsmon output blocks to map
func textToMap(slice []string) (blockMap map[string]string) {
	blockMap = make(map[string]string)
	for _, v := range slice {
		v = strings.Trim(v, " |-\n\t")
		if len(v) == 0 {
			continue
		}
		if strings.Index(v, ":") == -1 {
			continue
		}
		v = strings.ReplaceAll(v, "\\", "\\\\")
		v = strings.ReplaceAll(v, "\"", "")
		v = strings.ReplaceAll(v, ": ", ":")
		v = "{\"" + v + "\"}"
		v = strings.ReplaceAll(v, "  ", "")
		v = strings.Replace(v, ":", "\":\"", 1)
		v = strings.Replace(v, " \":", "\":", 1)

		var stringMap map[string]interface{}
		if err := json.Unmarshal([]byte(v), &stringMap); err != nil {
			panic(err)
		}

		for i, element := range stringMap {
			e, ok := element.(string)
			if ok {
				e = strings.Trim(e, " \t\n")
				blockMap[i] = e
			}
		}
	}
	return blockMap
}

// getFeturesInfo - get Feature Information
func getFeturesInfo(lsmonOut string) fetures {
	var feats fetures
	// feturesInfo - Slice of features. Text splitted by features blocks
	var feturesInfo []string
	var featureInfo, additionalStuff string
	lsmonOut = trimBorders(lsmonOut)
	feturesInfo = splitdata(lsmonOut, featureDelimeter)

	// Split data for features blocks
	for i, data := range feturesInfo {
		featureInfo, additionalStuff = splitFeature(data)
		feats.Features = append(feats.Features, parseFeature(featureInfo))
		if len(additionalStuff) > 0 {
			// TODO get all License Information
			feats.Features[i].LicensesInformation = getLicenseInformation(additionalStuff)
			// TODO get all client Information
			feats.Features[i].ClientsInformation = getClientInformation(additionalStuff)
		}
	}

	return feats
}

// parseFeature - parse Feature Information
func parseFeature(s string) feature {
	var feat feature
	var slice []string

	s = strings.Trim(s, " |-\n\t")
	slice = strings.Split(s, "\n   |- ")
	feat.Feature = textToMap(slice)
	return feat
}

// getLicenseInformation - get License Information
func getLicenseInformation(s string) []licenseInformation {
	var lics []licenseInformation
	var licsSlice, licSlice []string

	licsSlice = strings.Split(s, licDelimeter)
	for _, v := range licsSlice {
		i1 := strings.Index(v, clientDelimeter)
		if i1 >= 0 {
			v = v[:i1]
		}
		v = strings.Trim(v, " |-\n\t")
		if len(v) == 0 {
			continue
		}
		licSlice = strings.Split(v, "\n     |- ")

		lics = append(lics, licenseInformation{LicenseInformation: textToMap(licSlice)})
	}

	return lics
}

// getClientInformation - get Client Information
func getClientInformation(s string) []clientInformation {
	var clientsInfo []clientInformation
	var clientsSlice, clSlice []string

	clientsSlice = strings.Split(s, clientDelimeter)
	for _, v := range clientsSlice {
		i1 := strings.Index(v, licDelimeter)
		if i1 >= 0 {
			v = v[:i1]
		}
		v = strings.Trim(v, " |-\n\t")
		if len(v) == 0 {
			continue
		}
		clSlice = strings.Split(v, "\n     |- ")
		clientsInfo = append(clientsInfo, clientInformation{ClientInformation: textToMap(clSlice)})
	}

	return clientsInfo
}
