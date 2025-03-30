package CbGen

import (
	"AutomatedTestCaseGeneration/Generator"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"sort"
)

const DefaultCandidateIdx = 0

type CbGen struct {
	numPara      int
	ways         int
	parameters   [][]any
	combinations []string
}

func NewCbGen(parameters [][]any, ways int) (*CbGen, error) {
	numPara := len(parameters)
	if numPara < ways {
		return nil, fmt.Errorf("number of parameters < ways")
	}

	var newParameters [][]any
	for i, parameter := range parameters {
		if len(parameter) == 0 {
			return nil, fmt.Errorf("no candidates for parameter %d", i)
		}
		newParameters = append(newParameters, parameter)
	}

	newCb := &CbGen{
		numPara:    len(parameters),
		ways:       ways,
		parameters: newParameters,
	}
	newCb.genCombinations()

	return newCb, nil
}

func (cb *CbGen) genCombinations() {
	paraSelections := combin.Combinations(cb.numPara, cb.ways)

	for _, selection := range paraSelections {
		selectionMap := make(map[int]bool)
		for _, s := range selection {
			selectionMap[s] = true
		}
		ctsPds := cb.cartesianProduct(selectionMap)
		for _, ctsPd := range ctsPds {
			ctsPdStr, _ := Generator.SliceToString(ctsPd)
			cb.combinations = append(cb.combinations, ctsPdStr)
		}
	}
}

func (cb *CbGen) cartesianProduct(selection map[int]bool) [][]any {
	if len(selection) == 0 {
		return [][]any{}
	}
	result := [][]any{{}}
	for i := 0; i < cb.numPara; i++ {
		var temp [][]any
		for _, prefix := range result {
			if selection[i] {
				for _, v := range cb.parameters[i] {
					newCombo := append([]any{}, prefix...)
					newCombo = append(newCombo, v)
					temp = append(temp, newCombo)
				}
			} else {
				newCombo := append([]any{}, prefix...)
				newCombo = append(newCombo, cb.parameters[i][DefaultCandidateIdx])
				temp = append(temp, newCombo)
			}
		}
		result = temp
	}

	return result
}

func (cb *CbGen) GetCombinations() []string {
	sort.Strings(cb.combinations)
	return cb.combinations
}
