package nrpefn

import (
	"fmt"
	"os"
)

const (
	OK       int = iota
	Warning
	Critical
	Unknown
)

func FloatOK(value float64) {
	fmt.Printf("OK - Value: %.2f\n", value)
	os.Exit(OK)
}

func FloatWarning(value float64) {
	fmt.Printf("WARNING - Value: %.2f\n", value)
	os.Exit(Warning)
}

func FloatCritical(value float64) {
	fmt.Printf("CRITICAL - Value: %.2f\n", value)
	os.Exit(Critical)
}

func StringOK(value string) {
	fmt.Printf("OK - %s\n", value)
	os.Exit(OK)
}

func StringWarning(value string) {
	fmt.Printf("WARNING - %s\n", value)
	os.Exit(Warning)
}

func StringCritical(value string) {
	fmt.Printf("CRITICAL - %s\n", value)
	os.Exit(Critical)
}

func StringUnknown(){
	fmt.Printf("Unknown - %s\n", value)
	os.Exit(Unknown)
}

