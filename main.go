package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// CSV read tutorial and base code from https://gosamples.dev/read-csv/
	fileRead, err := os.Open("salaries.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(fileRead)

	total2024 := 0
	number2024 := 0
	salaryList2024 := make([]float64, 0)

	total2023 := 0
	number2023 := 0
	salaryList2023 := make([]float64, 0)

	totalMid2024 := 0
	numberMid2024 := 0
	salaryListMid2024 := make([]float64, 0)

	totalRemoteEntry2024 := 0
	numberRemoteEntry2024 := 0
	salaryListRemoteEntry2024 := make([]float64, 0)

	totalOnsiteEntry2024 := 0
	numberOnsiteEntry2024 := 0
	salaryListOnsiteEntry2024 := make([]float64, 0)

	totalRemoteAll2024 := 0
	numberRemoteAll2024 := 0
	salaryListRemoteAll2024 := make([]float64, 0)

	totalOnsiteAll2024 := 0
	numberOnsiteAll2024 := 0
	salaryListOnsiteAll2024 := make([]float64, 0)

	var jobTitle = make(map[string][]float64)

	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		salaryNum, err := strconv.Atoi(record[8])
		//Entry-level 2024
		if record[4] == "2024" && record[1] == "Entry-level" {
			number2024++
			total2024 += salaryNum
			//fmt.Println(record)
			salaryList2024 = append(salaryList2024, float64(salaryNum))
		}
		//Entry-level 2023
		if record[4] == "2023" && record[1] == "Entry-level" {
			number2023++
			total2023 += salaryNum
			//fmt.Println(record)
			salaryList2023 = append(salaryList2023, float64(salaryNum))
		}
		//Mid-level 2024
		if record[4] == "2024" && record[1] == "Mid-level" {
			numberMid2024++
			totalMid2024 += salaryNum
			//fmt.Println(record)
			salaryListMid2024 = append(salaryListMid2024, float64(salaryNum))
		}
		//Entry-level Remote 2024
		if record[4] == "2024" && record[1] == "Entry-level" && record[3] == "Remote" {
			numberRemoteEntry2024++
			totalRemoteEntry2024 += salaryNum
			//fmt.Println(record)
			salaryListRemoteEntry2024 = append(salaryListRemoteEntry2024, float64(salaryNum))
		}
		//Entry-level Onsite 2024
		if record[4] == "2024" && record[1] == "Entry-level" && record[3] == "On-site" {
			numberOnsiteEntry2024++
			totalOnsiteEntry2024 += salaryNum
			//fmt.Println(record)
			salaryListOnsiteEntry2024 = append(salaryListOnsiteEntry2024, float64(salaryNum))
		}
		//All-level Remote 2024
		if record[4] == "2024" && record[3] == "Remote" {
			numberRemoteAll2024++
			totalRemoteAll2024 += salaryNum
			//fmt.Println(record)
			salaryListRemoteAll2024 = append(salaryListRemoteAll2024, float64(salaryNum))
		}
		//All-level Onsite 2024
		if record[4] == "2024" && record[3] == "On-site" {
			numberOnsiteAll2024++
			totalOnsiteAll2024 += salaryNum
			//fmt.Println(record)
			salaryListOnsiteAll2024 = append(salaryListOnsiteAll2024, float64(salaryNum))
		}
		if !strings.Contains(record[0], "Research") && record[4] == "2024" && record[1] == "Entry-level" {
			jobTitle[record[0]] = append(jobTitle[record[0]], float64(salaryNum))
		}
	}
	average2024 := total2024 / number2024
	average2023 := total2023 / number2023

	//Median of entry-level 2024
	slices.Sort(salaryList2024)
	median2024 := float64(0)
	len2024 := len(salaryList2024)
	if (len2024 % 2) == 0 {
		median2024 = (salaryList2024[len2024/2-1] + salaryList2024[len2024/2]) / 2
	} else {
		median2024 = salaryList2024[int(math.Ceil(float64(len2024/2)))]
	}

	//Median of entry-level 2023
	slices.Sort(salaryList2023)
	median2023 := float64(0)
	len2023 := len(salaryList2023)
	if (len2023 % 2) == 0 {
		median2023 = (salaryList2023[len2023/2-1] + salaryList2023[len2023/2]) / 2
	} else {
		median2023 = salaryList2023[int(math.Ceil(float64(len2023/2)))]
	}

	//Median of mid-level 2024
	slices.Sort(salaryListMid2024)
	medianMid2024 := float64(0)
	lenMid2024 := len(salaryListMid2024)
	if (lenMid2024 % 2) == 0 {
		medianMid2024 = (salaryListMid2024[lenMid2024/2-1] + salaryListMid2024[lenMid2024/2]) / 2
	} else {
		medianMid2024 = salaryListMid2024[int(math.Ceil(float64(lenMid2024/2)))]
	}

	//Median of remote entry-level 2024
	slices.Sort(salaryListRemoteEntry2024)
	medianRemoteEntry2024 := float64(0)
	lenRemoteEntry2024 := len(salaryListRemoteEntry2024)
	if (lenRemoteEntry2024 % 2) == 0 {
		medianRemoteEntry2024 = (salaryListRemoteEntry2024[lenRemoteEntry2024/2-1] + salaryListRemoteEntry2024[lenRemoteEntry2024/2]) / 2
	} else {
		medianRemoteEntry2024 = salaryListRemoteEntry2024[int(math.Ceil(float64(lenRemoteEntry2024/2)))]
	}

	//Median of on-site entry-level 2024
	slices.Sort(salaryListOnsiteEntry2024)
	medianOnsiteEntry2024 := float64(0)
	lenOnsiteEntry2024 := len(salaryListOnsiteEntry2024)
	if (lenOnsiteEntry2024 % 2) == 0 {
		medianOnsiteEntry2024 = (salaryListOnsiteEntry2024[lenOnsiteEntry2024/2-1] + salaryListOnsiteEntry2024[lenOnsiteEntry2024/2]) / 2
	} else {
		medianOnsiteEntry2024 = salaryListOnsiteEntry2024[int(math.Ceil(float64(lenOnsiteEntry2024/2)))]
	}

	//Median of remote all-level 2024
	slices.Sort(salaryListRemoteAll2024)
	medianRemoteAll2024 := float64(0)
	lenRemoteAll2024 := len(salaryListRemoteAll2024)
	if (lenRemoteAll2024 % 2) == 0 {
		medianRemoteAll2024 = (salaryListRemoteAll2024[lenRemoteAll2024/2-1] + salaryListRemoteAll2024[lenRemoteAll2024/2]) / 2
	} else {
		medianRemoteAll2024 = salaryListRemoteAll2024[int(math.Ceil(float64(lenRemoteAll2024/2)))]
	}

	//Median of on-site all-level 2024
	slices.Sort(salaryListOnsiteAll2024)
	medianOnsiteAll2024 := float64(0)
	lenOnsiteAll2024 := len(salaryListOnsiteAll2024)
	if (lenOnsiteAll2024 % 2) == 0 {
		medianOnsiteAll2024 = (salaryListOnsiteAll2024[lenOnsiteAll2024/2-1] + salaryListOnsiteAll2024[lenOnsiteAll2024/2]) / 2
	} else {
		medianOnsiteAll2024 = salaryListOnsiteAll2024[int(math.Ceil(float64(lenOnsiteAll2024/2)))]
	}

	midEntryDifference2024 := medianMid2024 - median2024

	//Median of entry-level job titles from 2024
	jobMedian := make(map[string]float64)
	for titleName, salaryList := range jobTitle {
		medianJob := float64(0)
		slices.Sort(salaryList)
		lenJob := len(salaryList)
		if (lenJob % 2) == 0 {
			medianJob = (salaryList[lenJob/2-1] + salaryList[lenJob/2]) / 2
		} else {
			medianJob = salaryList[int(math.Ceil(float64(lenJob/2)))]
		}
		jobMedian[titleName] = medianJob
	}

	//Sorting a map code retrieved from https://forum.golangbridge.org/t/sort-a-map-by-value/27710/3
	keys := make([]string, 0, len(jobMedian))
	for key := range jobMedian {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return jobMedian[keys[i]] > jobMedian[keys[j]] })

	color.Blue("\n--- Job salary stats from salaries.csv ---")
	fmt.Println()
	if average2024 < average2023 {
		color.Red("The average salary of entry-level jobs from 2024 ($" + strconv.Itoa(average2024) + ") is less than the average salary of entry-level jobs from 2023 ($" + strconv.Itoa(average2023) + ")")
	} else {
		color.Green("The average salary of entry-level jobs from 2024 ($" + strconv.Itoa(average2024) + ") is greater than the average salary of entry-level jobs from 2023 ($" + strconv.Itoa(average2023) + ")")
	}
	fmt.Println()
	fmt.Print("The median salary of entry-level jobs from 2024 compared to the median salary of entry-level jobs from 2023 is: ")
	if median2024 < median2023 {
		color.Red("-$" + strconv.Itoa((int(median2024))-(int(median2023))))
	} else {
		color.Green("+$" + strconv.Itoa((int(median2024))-(int(median2023))))
	}
	fmt.Println()
	if float64(average2024) > median2024 {
		fmt.Print("The average salary of entry-level jobs from 2024 compared to the median salary of entry-level jobs from 2024 is: ")
		color.Green("+$" + strconv.Itoa((average2024)-(int(median2024))))
	} else {
		fmt.Print("The average salary of entry-level jobs from 2024 compared to the median salary of entry-level jobs from 2024 is: ")
		color.Red("-$" + strconv.Itoa((average2024)-(int(median2024))))
	}
	if float64(average2023) > median2023 {
		fmt.Print("The average salary of entry-level jobs from 2023 compared to the median salary of entry-level jobs from 2023 is: ")
		color.Green("+$" + strconv.Itoa((average2023)-(int(median2023))))
	} else {
		fmt.Print("The average salary of entry-level jobs from 2023 compared to the median salary of entry-level jobs from 2023 is: ")
		color.Red("-$" + strconv.Itoa((average2023)-(int(median2023))))
	}
	fmt.Println()
	fmt.Print("By moving from an entry-level job to a mid-level job in 2024, you could get a salary raise of ")
	color.Cyan("$" + strconv.Itoa(int(midEntryDifference2024)) + "\n")
	fmt.Println()
	if medianRemoteEntry2024 < medianOnsiteEntry2024 {
		color.Red("Entry-level median remote salary from 2024: $" + strconv.Itoa(int(medianRemoteEntry2024)) + "\n")
		color.Red("Entry-level median on-site salary from 2024: $" + strconv.Itoa(int(medianOnsiteEntry2024)) + "\n")
	} else {
		color.Green("Entry-level median remote salary from 2024: $" + strconv.Itoa(int(medianRemoteEntry2024)) + "\n")
		color.Green("Entry-level median on-site salary from 2024: $" + strconv.Itoa(int(medianOnsiteEntry2024)) + "\n")
	}
	fmt.Println()
	if medianRemoteAll2024 < medianOnsiteAll2024 {
		color.Red("All-level median remote salary from 2024: $" + strconv.Itoa(int(medianRemoteAll2024)) + "\n")
		color.Red("All-level median on-site salary from 2024: $" + strconv.Itoa(int(medianOnsiteAll2024)) + "\n")
	} else {
		color.Green("All-level median remote salary from 2024: $" + strconv.Itoa(int(medianRemoteAll2024)) + "\n")
		color.Green("All-level median on-site salary from 2024: $" + strconv.Itoa(int(medianOnsiteAll2024)) + "\n")
	}
	fmt.Println("\nTop 5 highest paying entry-level jobs from 2024:")
	i := 0
	for _, key := range keys {
		if i < 5 {
			if i == 0 {
				color.Yellow("#"+strconv.Itoa(i+1)+" - %s: $%d\n", key, int(jobMedian[key]))
			} else {
				color.Cyan("#"+strconv.Itoa(i+1)+" - %s: $%d\n", key, int(jobMedian[key]))
			}
			i++
		} else {
			break
		}
	}
}
