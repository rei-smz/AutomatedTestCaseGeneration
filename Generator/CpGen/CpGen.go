package CpGen

import (
	"AutomatedTestCaseGeneration/Generator/CbGen"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"math"
	"sort"
)

type CpGen struct {
	boundaries [][]any
	testCases  [][]any
	cbGen      *CbGen.CbGen
}

func NewCpGen(boundaries [][]any) (*CpGen, error) {
	if len(boundaries) == 0 {
		return nil, fmt.Errorf("no boundaries provided")
	}
	newCpGen := CpGen{
		boundaries: boundaries,
	}
	for _, boundary := range boundaries {
		switch boundary[0].(type) {
		case int64:
			newCpGen.testCases = append(newCpGen.testCases, newCpGen.genCasesInt(boundary))
		case float64:
			newCpGen.testCases = append(newCpGen.testCases, newCpGen.genCasesFloat(boundary))
		case string:
			newCpGen.testCases = append(newCpGen.testCases, boundary)
		}
	}
	log.Println(newCpGen.testCases)
	var err error
	newCpGen.cbGen, err = CbGen.NewCbGen(newCpGen.testCases, len(boundaries))
	if err != nil {
		return nil, err
	}

	return &newCpGen, nil
}

func (cp *CpGen) genCasesInt(boundary []any) []any {
	sort.Slice(boundary, func(i, j int) bool {
		return boundary[i].(int64) < boundary[j].(int64)
	})
	set := mapset.NewSet[any]()
	for i, b := range boundary {
		temps := []int64{b.(int64), b.(int64) - 1, b.(int64) + 1}
		if i == 0 {
			temps = append(temps, b.(int64)-2)
		}
		for _, t := range temps {
			if i == 0 && i == len(boundary)-1 {
				set.Add(t)
				continue
			}
			if i == 0 && t < boundary[i+1].(int64) {
				set.Add(t)
				continue
			}
			if i == len(boundary)-1 && t > boundary[i-1].(int64)+1 {
				set.Add(t)
				continue
			}
			if i > 0 && i < len(boundary)-1 && t < boundary[i+1].(int64) && t > boundary[i-1].(int64)+1 {
				set.Add(t)
			}
		}
	}
	return set.ToSlice()
}

func (cp *CpGen) genCasesFloat(boundary []any) []any {
	const Accuracy = 1e-16
	sort.Slice(boundary, func(i, j int) bool {
		return boundary[i].(float64)-boundary[j].(float64) < Accuracy
	})
	log.Println("genCasesFloat input", boundary)
	var returns []any
	for i, b := range boundary {
		temps := []float64{b.(float64), b.(float64) - 10e-3, b.(float64) + 1}
		if i == 0 {
			temps = append(temps, b.(float64)-1)
		}
		for _, t := range temps {
			duplicate := false
			for _, v := range returns {
				if math.Abs(t-v.(float64)) < Accuracy {
					duplicate = true
					break
				}
			}
			if duplicate {
				continue
			}
			if i == 0 && i == len(boundary)-1 {
				returns = append(returns, t)
				continue
			}
			if i == 0 && t-boundary[i+1].(float64) < Accuracy {
				returns = append(returns, t)
				continue
			}
			if i == len(boundary)-1 && t-boundary[i-1].(float64)+1 > Accuracy {
				returns = append(returns, t)
				continue
			}
			if i > 0 && i < len(boundary)-1 && t-boundary[i+1].(float64) < Accuracy && t-boundary[i-1].(float64)+1 > Accuracy {
				returns = append(returns, t)
			}
		}
	}
	log.Println("genCasesFloat output", returns)
	return returns
}

func (cp *CpGen) GetCombinations() []string {
	return cp.cbGen.GetCombinations()
}
