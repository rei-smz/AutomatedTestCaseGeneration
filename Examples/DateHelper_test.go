package Examples

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

var dh = NewDateHelper()

func TestDateHelper_GetDateAndTime(t *testing.T) {
	type args struct {
		timeMillis int64
	}
	var tests []struct {
		name string
		args args
		want string
	}
	const testName string = "cb-1way"
	inputFile, _ := os.Open("../timemillis_" + testName + "-result")
	outputFile, _ := os.Open("../getdateandtime_" + testName + "-out")
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
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		test := struct {
			name string
			args args
			want string
		}{
			name: line,
			args: args{
				a,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		tests[testIdx].want = line
		testIdx++
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dh.GetDateAndTime(tt.args.timeMillis); got != tt.want {
				t.Errorf("GetDateAndTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_GetDateOnly(t *testing.T) {
	type args struct {
		dateStr string
	}
	var tests []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}

	const testName string = "cb-1way"
	inputFile, _ := os.Open("../datestr_" + testName + "-result")
	outputFile, _ := os.Open("../getdateonly_" + testName + "-out")
	defer inputFile.Close()
	defer outputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)
	outputScanner := bufio.NewScanner(outputFile)

	for inputScanner.Scan() {
		line := strings.TrimSpace(inputScanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[\"")
		line = strings.TrimSuffix(line, "\"]")
		a := strings.TrimSpace(line)
		test := struct {
			name    string
			args    args
			want    int64
			wantErr bool
		}{
			name: line,
			args: args{
				a,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		if strings.Contains(line, "err:") {
			tests[testIdx].wantErr = true
		} else {
			tests[testIdx].want, _ = strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		}
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dh.GetDateOnly(tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDateOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDateOnly() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_GetDateOnlyFromTime(t *testing.T) {
	type args struct {
		timeMillis int64
	}
	var tests []struct {
		name string
		args args
		want string
	}

	const testName string = "cb-1way"
	inputFile, _ := os.Open("../timemillis_" + testName + "-result")
	outputFile, _ := os.Open("../getdateonlyfromtime_" + testName + "-out")
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
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		test := struct {
			name string
			args args
			want string
		}{
			name: line,
			args: args{
				a,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		tests[testIdx].want = line
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dh.GetDateOnlyFromTime(tt.args.timeMillis); got != tt.want {
				t.Errorf("GetDateOnlyFromTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_GetDaysBetweenTwoDate(t *testing.T) {
	type args struct {
		oldStr string
		newStr string
		df     string
	}
	var tests []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}

	const testName string = "cb-2way"
	inputFile, _ := os.Open("../ond_" + testName + "-result")
	outputFile, _ := os.Open("../getdaysbetweentwodate_" + testName + "-out")
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
		a := strings.TrimSpace(parts[0])
		b := strings.TrimSpace(parts[1])
		c := strings.TrimSpace(parts[2])
		a = strings.Trim(a, "\"")
		b = strings.Trim(b, "\"")
		c = strings.Trim(c, "\"")
		test := struct {
			name    string
			args    args
			want    int64
			wantErr bool
		}{
			name: line,
			args: args{
				a,
				b,
				c,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		if strings.Contains(line, "err:") {
			tests[testIdx].wantErr = true
		} else {
			tests[testIdx].want, _ = strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		}
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dh.GetDaysBetweenTwoDate(tt.args.oldStr, tt.args.newStr, tt.args.df)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDaysBetweenTwoDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDaysBetweenTwoDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestDateHelper_GetDesiredFormat(t *testing.T) {
//	type args struct {
//		df string
//	}
//	var tests []struct {
//		name string
//		args args
//		want string
//	}
//
//	const testName string = "cb-1way"
//	inputFile, _ := os.Open("../df_" + testName + "-result")
//	outputFile, _ := os.Open("../getdesiredformat_" + testName + "-out")
//	defer inputFile.Close()
//	defer outputFile.Close()
//	inputScanner := bufio.NewScanner(inputFile)
//	outputScanner := bufio.NewScanner(outputFile)
//
//	for inputScanner.Scan() {
//		line := strings.TrimSpace(inputScanner.Text())
//		if line == "" {
//			continue
//		}
//		line = strings.TrimPrefix(line, "[\"")
//		line = strings.TrimSuffix(line, "\"]")
//		a := strings.TrimSpace(line)
//		test := struct {
//			name string
//			args args
//			want string
//		}{
//			name: line,
//			args: args{
//				a,
//			},
//		}
//		tests = append(tests, test)
//	}
//
//	testIdx := 0
//	for outputScanner.Scan() {
//		line := strings.TrimSpace(outputScanner.Text())
//		if line == "" {
//			continue
//		}
//		tests[testIdx].want = line
//		testIdx++
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := dh.GetDesiredFormat(tt.args.df); got != tt.want {
//				t.Errorf("GetDesiredFormat() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestDateHelper_GetDesiredFormatFromTime(t *testing.T) {
	type args struct {
		df         string
		timeMillis int64
	}
	var tests []struct {
		name string
		args args
		want string
	}

	const testName string = "cb-2way"
	inputFile, _ := os.Open("../dftm_" + testName + "-result")
	outputFile, _ := os.Open("../getdesiredformatfromtime_" + testName + "-out")
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
		a := strings.TrimSpace(parts[0])
		a = strings.Trim(a, "\"")
		b, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		test := struct {
			name string
			args args
			want string
		}{
			name: line,
			args: args{
				a,
				b,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		tests[testIdx].want = line
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dh.GetDesiredFormatFromTime(tt.args.df, tt.args.timeMillis); got != tt.want {
				t.Errorf("GetDesiredFormatFromTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_GetHoursBetweenTwoDate(t *testing.T) {
	type args struct {
		oldStr string
		newStr string
		df     string
	}
	var tests []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}

	const testName string = "cb-2way"
	inputFile, _ := os.Open("../ond_" + testName + "-result")
	outputFile, _ := os.Open("../gethoursbetweentwodate_" + testName + "-out")
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
		a := strings.TrimSpace(parts[0])
		b := strings.TrimSpace(parts[1])
		c := strings.TrimSpace(parts[2])
		a = strings.Trim(a, "\"")
		b = strings.Trim(b, "\"")
		c = strings.Trim(c, "\"")
		test := struct {
			name    string
			args    args
			want    int64
			wantErr bool
		}{
			name: line,
			args: args{
				a,
				b,
				c,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		if strings.Contains(line, "err:") {
			tests[testIdx].wantErr = true
		} else {
			tests[testIdx].want, _ = strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		}
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dh.GetHoursBetweenTwoDate(tt.args.oldStr, tt.args.newStr, tt.args.df)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHoursBetweenTwoDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetHoursBetweenTwoDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_GetMinutesBetweenTwoDates(t *testing.T) {
	type args struct {
		oldStr string
		newStr string
		df     string
	}
	var tests []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}

	const testName string = "cb-2way"
	inputFile, _ := os.Open("../ond_" + testName + "-result")
	outputFile, _ := os.Open("../getminutesbetweentwodate_" + testName + "-out")
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
		a := strings.TrimSpace(parts[0])
		b := strings.TrimSpace(parts[1])
		c := strings.TrimSpace(parts[2])
		a = strings.Trim(a, "\"")
		b = strings.Trim(b, "\"")
		c = strings.Trim(c, "\"")
		test := struct {
			name    string
			args    args
			want    int64
			wantErr bool
		}{
			name: line,
			args: args{
				a,
				b,
				c,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		if strings.Contains(line, "err:") {
			tests[testIdx].wantErr = true
		} else {
			tests[testIdx].want, _ = strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		}
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dh.GetMinutesBetweenTwoDates(tt.args.oldStr, tt.args.newStr, tt.args.df)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMinutesBetweenTwoDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMinutesBetweenTwoDates() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_GetTimeOnly(t *testing.T) {
	type args struct {
		timeMillis int64
	}
	var tests []struct {
		name string
		args args
		want string
	}

	const testName string = "cb-1way"
	inputFile, _ := os.Open("../timemillis_" + testName + "-result")
	outputFile, _ := os.Open("../gettimeonly_" + testName + "-out")
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
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		test := struct {
			name string
			args args
			want string
		}{
			name: line,
			args: args{
				a,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		tests[testIdx].want = line
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dh.GetTimeOnly(tt.args.timeMillis); got != tt.want {
				t.Errorf("GetTimeOnly() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_GetToday(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "today",
			want: "29/03/2025",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dh.GetToday(); got != tt.want {
				t.Errorf("GetToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestDateHelper_GetTodayWithTime(t *testing.T) {
//	tests := []struct {
//		name string
//		want string
//	}{
//		{
//			name: "today",
//			want: "26/03/2025",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := dh.GetTodayWithTime(); got != tt.want {
//				t.Errorf("GetTodayWithTime() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestDateHelper_GetTomorrow(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "tomorrow",
			want: "30/03/2025",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dh.GetTomorrow(); got != tt.want {
				t.Errorf("GetTomorrow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_ParseAnyDate(t *testing.T) {
	type args struct {
		dateStr string
	}
	var tests []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}

	const testName string = "cb-1way"
	inputFile, _ := os.Open("../datestr_" + testName + "-result")
	outputFile, _ := os.Open("../parseanydate_" + testName + "-out")
	defer inputFile.Close()
	defer outputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)
	outputScanner := bufio.NewScanner(outputFile)

	for inputScanner.Scan() {
		line := strings.TrimSpace(inputScanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[\"")
		line = strings.TrimSuffix(line, "\"]")
		a := strings.TrimSpace(line)
		test := struct {
			name    string
			args    args
			want    int64
			wantErr bool
		}{
			name: line,
			args: args{
				a,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		if strings.Contains(line, "err:") {
			tests[testIdx].wantErr = true
		} else {
			tests[testIdx].want, _ = strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		}
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dh.ParseAnyDate(tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAnyDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseAnyDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_ParseDate(t *testing.T) {
	type args struct {
		dateStr string
		df      string
	}
	var tests []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}

	const testName string = "cb-2way"
	inputFile, _ := os.Open("../dsdf_" + testName + "-result")
	outputFile, _ := os.Open("../parsedate_" + testName + "-out")
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
		a := strings.TrimSpace(parts[0])
		b := strings.TrimSpace(parts[0])
		a = strings.Trim(a, "\"")
		b = strings.Trim(b, "\"")
		test := struct {
			name    string
			args    args
			want    int64
			wantErr bool
		}{
			name: line,
			args: args{
				a,
				b,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		if strings.Contains(line, "err:") {
			tests[testIdx].wantErr = true
		} else {
			tests[testIdx].want, _ = strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		}
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dh.ParseDate(tt.args.dateStr, tt.args.df)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_PrettifyDate(t *testing.T) {
	type args struct {
		timeMillis int64
	}
	var tests []struct {
		name string
		args args
		want string
	}

	const testName string = "cb-1way"
	inputFile, _ := os.Open("../timemillis_" + testName + "-result")
	outputFile, _ := os.Open("../prettifydate_" + testName + "-out")
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
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		test := struct {
			name string
			args args
			want string
		}{
			name: line,
			args: args{
				a,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		tests[testIdx].want = line
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dh.PrettifyDate(tt.args.timeMillis); got != tt.want {
				t.Errorf("PrettifyDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateHelper_PrettifyDateFromString(t *testing.T) {
	type args struct {
		timestampStr string
	}
	var tests []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}

	const testName string = "cb-1way"
	inputFile, _ := os.Open("../timestampstr_" + testName + "-result")
	outputFile, _ := os.Open("../prettifydatefromstring_" + testName + "-out")
	defer inputFile.Close()
	defer outputFile.Close()
	inputScanner := bufio.NewScanner(inputFile)
	outputScanner := bufio.NewScanner(outputFile)

	for inputScanner.Scan() {
		line := strings.TrimSpace(inputScanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[\"")
		line = strings.TrimSuffix(line, "\"]")
		a := strings.TrimSpace(line)
		test := struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{
			name: line,
			args: args{
				a,
			},
		}
		tests = append(tests, test)
	}

	testIdx := 0
	for outputScanner.Scan() {
		line := strings.TrimSpace(outputScanner.Text())
		if line == "" {
			continue
		}
		if strings.Contains(line, "err:") {
			tests[testIdx].wantErr = true
		} else {
			tests[testIdx].want = line
		}
		testIdx++
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dh.PrettifyDateFromString(tt.args.timestampStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("PrettifyDateFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PrettifyDateFromString() got = %v, want %v", got, tt.want)
			}
		})
	}
}
