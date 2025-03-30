package main

import (
	"AutomatedTestCaseGeneration/Examples"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func runGetDateAndTime(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("timemillis_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("getdateandtime_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		res := dh.GetDateAndTime(a)
		writer.WriteString(fmt.Sprintln(res))
		writer.WriteString("\n")
	}
}

func runGetDateOnly(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("datestr_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("getdateonly_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[\"")
		line = strings.TrimSuffix(line, "\"]")
		a := strings.TrimSpace(line)
		res, err := dh.GetDateOnly(a)
		if err != nil {
			writer.WriteString("err: " + fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

func runGetDateOnlyFromTime(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("timemillis_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("getdateonlyfromtime_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		res := dh.GetDateOnlyFromTime(a)
		writer.WriteString(fmt.Sprintln(res))
		writer.WriteString("\n")
	}
}

func runGetDaysBetweenTwoDate(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("ond_cb-3way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("getdaysbetweentwodate_cb-3way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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
		res, err := dh.GetDaysBetweenTwoDate(a, b, c)
		if err != nil {
			writer.WriteString("err: " + fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

func runGetDesiredFormat(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("df_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("getdesiredformat_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[\"")
		line = strings.TrimSuffix(line, "\"]")
		a := strings.TrimSpace(line)
		res := dh.GetDesiredFormat(a)
		writer.WriteString(fmt.Sprintln(res))
		writer.WriteString("\n")
	}
}

func runGetDesiredFormatFromTime(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("dftm_cb-2way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("getdesiredformatfromtime_cb-2way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		parts := strings.Split(line, ",")
		a := strings.TrimSpace(parts[0])
		a = strings.Trim(a, "\"")
		b, _ := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		res := dh.GetDesiredFormatFromTime(a, b)
		writer.WriteString(fmt.Sprintln(res))
		writer.WriteString("\n")
	}
}

func runGetHoursBetweenTwoDate(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("ond_cb-3way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("gethoursbetweentwodate_cb-3way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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
		res, err := dh.GetHoursBetweenTwoDate(a, b, c)
		if err != nil {
			writer.WriteString("err: " + fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

func runGetMinutesBetweenTwoDates(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("ond_cb-3way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("getminutesbetweentwodate_cb-3way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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
		res, err := dh.GetMinutesBetweenTwoDates(a, b, c)
		if err != nil {
			writer.WriteString("err: " + fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

func runGetTimeOnly(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("timemillis_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("gettimeonly_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		res := dh.GetTimeOnly(a)
		writer.WriteString(fmt.Sprintln(res))
		writer.WriteString("\n")
	}
}

func runGetToday(dh *Examples.DateHelper) {
	fmt.Println("GetToday:", dh.GetToday())
}

func runGetTodayWithTime(dh *Examples.DateHelper) {
	fmt.Println("GetToday:", dh.GetTodayWithTime())
}

func runGetTomorrow(dh *Examples.DateHelper) {
	fmt.Println("GetToday:", dh.GetTomorrow())
}

func runParseAnyDate(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("datestr_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("parseanydate_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[\"")
		line = strings.TrimSuffix(line, "\"]")
		a := strings.TrimSpace(line)
		res, err := dh.ParseAnyDate(a)
		if err != nil {
			writer.WriteString("err: " + fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

func runParseDate(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("dsdf_cb-2way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("parsedate_cb-2way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
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
		res, err := dh.ParseDate(a, b)
		if err != nil {
			writer.WriteString("err: " + fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

func runPrettifyDate(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("timemillis_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("prettifydate_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		a, _ := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		res := dh.PrettifyDate(a)
		writer.WriteString(fmt.Sprintln(res))
		writer.WriteString("\n")
	}
}

func runPrettifyDateFromString(dh *Examples.DateHelper) {
	inputFile, _ := os.Open("timestampstr_cb-1way-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("prettifydatefromstring_cb-1way-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[\"")
		line = strings.TrimSuffix(line, "\"]")
		a := strings.TrimSpace(line)
		res, err := dh.PrettifyDateFromString(a)
		if err != nil {
			writer.WriteString("err: " + fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

func runDateHelper() {
	dh := Examples.NewDateHelper()
	runGetDateAndTime(dh)
	runGetDateOnly(dh)
	runGetDateOnlyFromTime(dh)
	runGetDaysBetweenTwoDate(dh)
	runGetDesiredFormat(dh)
	runGetDesiredFormatFromTime(dh)
	runGetHoursBetweenTwoDate(dh)
	runGetMinutesBetweenTwoDates(dh)
	runGetTimeOnly(dh)
	runGetToday(dh)
	runGetTodayWithTime(dh)
	runGetTomorrow(dh)
	runParseAnyDate(dh)
	runParseDate(dh)
	runPrettifyDate(dh)
	runPrettifyDateFromString(dh)
}

func runQuadraticSolver() {
	inputFile, _ := os.Open("quadratic_new_cg-cp-result")
	defer inputFile.Close()
	outputFile, _ := os.Create("quadratic_new_cg-cp-out")
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	quadratic := Examples.NewQuadraticSolver()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		line = strings.TrimPrefix(line, "[")
		line = strings.TrimSuffix(line, "]")
		parts := strings.Split(line, ",")
		a, _ := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
		b, _ := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
		c, _ := strconv.ParseFloat(strings.TrimSpace(parts[2]), 64)
		res, err := quadratic.SolveQuadratic(a, b, c)
		if err != nil {
			writer.WriteString(fmt.Sprintln(err))
		} else {
			writer.WriteString(fmt.Sprintln(res))
		}
		writer.WriteString("\n")
	}
}

//func main() {
//	//runQuadraticSolver()
//
//	runDateHelper()
//}

//func main() {
//	// 加拿大东部时间 (EDT, UTC-4)
//	loc, _ := time.LoadLocation("America/Toronto")
//	// 2025年3月26日 00:00:00 EDT
//	start := time.Date(2025, 3, 29, 0, 0, 0, 0, loc)
//	// 2025年3月26日 23:59:59.999 EDT
//	end := time.Date(2025, 3, 29, 23, 59, 59, 999*1e6, loc)
//	fmt.Println("Start UnixMilli:", start.UnixMilli())
//	fmt.Println("End UnixMilli:  ", end.UnixMilli())
//}
