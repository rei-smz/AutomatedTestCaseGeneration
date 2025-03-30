package Examples

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestQuadraticSolver_SolveQuadratic(t *testing.T) {

	type args struct {
		a float64
		b float64
		c float64
	}
	var tests []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}

	const testName string = "quadratic_cb-2way"
	inputFile, _ := os.Open("../" + testName + "-result")
	outputFile, _ := os.Open("../" + testName + "-out")
	defer inputFile.Close()
	defer outputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)
	outputScanner := bufio.NewScanner(outputFile)

	for inputScanner.Scan() {
		line := strings.TrimSpace(inputScanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		parts := strings.Split(line, ",")
		a, _ := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		b, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		c, _ := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
		test := struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{
			name: line,
			args: args{
				a: a,
				b: b,
				c: c,
			},
		}
		tests = append(tests, test)
	}

	var currentLines []string
	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			if len(currentLines) > 0 {
				combined := strings.Join(currentLines, "\n")
				if strings.EqualFold(strings.TrimSpace(combined), "not enough precision") {
					tests[testIdx].want = ""
					tests[testIdx].wantErr = true
				} else {
					tests[testIdx].want = combined
					tests[testIdx].wantErr = false
				}
				testIdx++
				currentLines = nil
			}
		} else {
			currentLines = append(currentLines, line)
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := NewQuadraticSolver()
			got, err := qs.SolveQuadratic(tt.args.a, tt.args.b, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("SolveQuadratic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SolveQuadratic() got = %v, want %v", got, tt.want)
			}
		})
	}
}
