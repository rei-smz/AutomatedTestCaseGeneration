package main

import (
	"AutomatedTestCaseGeneration/Generator"
	"AutomatedTestCaseGeneration/Generator/CbGen"
	"AutomatedTestCaseGeneration/Generator/CpGen"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	Name       string  `json:"name"`
	Parameters [][]any `json:"parameters"`
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go Main.go [config path] [mode]")
		os.Exit(1)
	}
	configPath := os.Args[1]
	mode := os.Args[2]
	parseInt := false
	if len(os.Args) > 3 && os.Args[3] == "--parse-int" {
		parseInt = true
	}

	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		panic("Can not open config file")
	}

	configData, err := io.ReadAll(configFile)
	if err != nil {
		panic(err.Error())
	}
	if len(configData) == 0 {
		panic("Empty config file")
	}

	var tempConfig Config
	err = json.Unmarshal(configData, &tempConfig)
	if err != nil {
		panic(err.Error())
	}
	config := Config{
		Name: tempConfig.Name,
	}
	for _, array := range tempConfig.Parameters {
		var temp []any
		for _, v := range array {
			if parseInt {
				switch v.(type) {
				case float64:
					temp = append(temp, func(fv float64) int64 { return int64(fv) }(v.(float64)))
				default:
					temp = append(temp, v)
				}
			} else {
				temp = append(temp, v)
			}
		}
		config.Parameters = append(config.Parameters, temp)
	}

	var result []string
	var generator Generator.Generator
	switch mode {
	case "cp":
		generator, err = CpGen.NewCpGen(config.Parameters)
		if err != nil {
			panic(err.Error())
		}
		result = generator.GetCombinations()
	case "2way":
		generator, err = CbGen.NewCbGen(config.Parameters, 2)
		if err != nil {
			panic(err.Error())
		}
		result = generator.GetCombinations()
	case "3way":
		generator, err = CbGen.NewCbGen(config.Parameters, 3)
		if err != nil {
			panic(err.Error())
		}
		result = generator.GetCombinations()
	default:
		generator, err = CbGen.NewCbGen(config.Parameters, 1)
		if err != nil {
			panic(err.Error())
		}
		result = generator.GetCombinations()
	}

	outputPath := fmt.Sprintf("%v-%v-result", config.Name, mode)
	outputFile, err := os.Create(outputPath)
	defer outputFile.Close()
	if err != nil {
		panic("Can not open output file")
	}

	for _, combination := range result {
		outputFile.WriteString(fmt.Sprintf("%v\n", combination))
	}
}
