package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

const (
	EXPECTED_AMOUNT_OF_PARTS = 2
	PROPERTY_NAME_REGEX      = "^([a-zA-Z_$][a-zA-Z_$0-9]*)(\\.[a-zA-Z_$][a-zA-Z_$0-9]*)*$"
)

type propertyMap map[string]interface{}

func splitAtEqualSign(s string) []string {
	return strings.Split(s, "=")
}

func isValidPropertyName(propName string) bool {
	regex, _ := regexp.Compile(PROPERTY_NAME_REGEX)
	return regex.MatchString(propName)
}

func castToCorrectType(stringVal string) interface{} {
	stringVal = strings.TrimSpace(stringVal)

	if intVal, err := strconv.ParseInt(stringVal, 10, 64); err == nil {
		return intVal
	}
	if floatVal, err := strconv.ParseFloat(stringVal, 64); err == nil {
		return floatVal
	}
	if boolVal, err := strconv.ParseBool(stringVal); err == nil {
		return boolVal
	}

	return stringVal
}

func parseLine(line string, lineNumber int) (propertyName string, propertyValue interface{}, err error) {
	lineParts := splitAtEqualSign(line)
	if len(lineParts) != EXPECTED_AMOUNT_OF_PARTS {
		return "", "", fmt.Errorf("syntax error in line %d: unexpected amount of equal signs, should be exactly one", lineNumber)
	}

	propertyName = lineParts[0]
	if !isValidPropertyName(propertyName) {
		return "", "", fmt.Errorf("syntax error in line %d: invalid property name", lineNumber)
	}

	propertyValue = castToCorrectType(lineParts[1])

	return propertyName, propertyValue, nil
}

func main() {
	propertyMap := make(propertyMap)

	propertyFile, err := os.Open("resources/my.properties")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer propertyFile.Close()

	scanner := bufio.NewScanner(propertyFile)
	for lineNumber := 1; scanner.Scan(); lineNumber++ {
		currLine := scanner.Text()

		propertyName, propertyValue, err := parseLine(currLine, lineNumber)
		if err != nil {
			log.Fatal(err)
			return
		}

		propertyMap[propertyName] = propertyValue
	}

	// Test printing
	for k, v := range propertyMap {
		fmt.Println("key:", k, reflect.TypeOf(k), ":: value:", v, reflect.TypeOf(v))
	}
}
