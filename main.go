package main

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var validCost = regexp.MustCompile(`^(?P<date>(19|20)\d\d-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01]))\s+(?P<category>\w+)\s+(?P<value>[0-9]{1,9})\s+(?P<comment>...{1,255}?)$`)

type Cost struct {
	gorm.Model
	date     time.Time
	category string
	value    int
	comment  string
}

func NewCost(date time.Time, category string, value int, comment string) *Cost {
	return &Cost{date: date, category: category, value: value, comment: comment}
}

func (c *Cost) String() string {
	d := c.date.Format("2006-01-02")
	return fmt.Sprintf("%v %s %d %s", d, c.category, c.value, c.comment)
}

// panic with error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// return empty string if input is empty
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter cost in 2000-01-31 category 995 comment not more than 255 vars")
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			return text
		} else {
			break
		}
	}
	return ""
}

func CostFromString(s string) *Cost {
	if isValidCost(s) {
		return ToCost(s)
	}
	return nil
}

func CostFromTerminal() *Cost {
	s := readInput()
	return CostFromString(s)
}

func CostsFromFile(filename string) []*Cost {
	var result = make([]*Cost, 0)
	skipped, added, total := 0, 0, 0
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cost := scanner.Text()
		total++
		if isValidCost(cost) {
			//fmt.Println(cost)
			result = append(result, ToCost(cost))
			added++
		} else {
			skipped++
		}
	}
	fmt.Printf("Total lines in file %d \t\t Correct: %d  \t\t Skipped: %d\n", total, added, skipped)
	return result
}

func ToCost(s string) *Cost {
	rs := validCost.FindStringSubmatch(s)
	date, err := time.Parse("2006-01-02", rs[1])
	check(err)
	category := rs[5]
	value, _ := strconv.Atoi(rs[6])
	comment := rs[7]
	return NewCost(date, category, value, comment)
}

func isValidCost(s string) bool {
	// valid cost by regex, must be smth like:
	// YYYY-MM-DD category value comment
	trimmed := strings.TrimSpace(s)

	if validCost.MatchString(trimmed) {
		return true
	} else {
		return false
	}
}

func database() {

}

func main() {
	costs := make([]*Cost, 0)
	costs = CostsFromFile("/tmp/costs")

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=dinero dbname=dinero password=TodayIsTheBestDay sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Cost{})

	for _, v := range costs {
		fmt.Println(v)
		db.Create(v)
	}

	// Create cost by hand
	//c := CostFromTerminal()
	//fmt.Println(c)
}
