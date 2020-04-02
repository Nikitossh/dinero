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

var validCost = regexp.MustCompile(`^(?P<Date>(19|20)\d\d-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01]))\s+(?P<Category>\w+)\s+(?P<Value>[0-9]{1,9})\s+(?P<Comment>...{1,255}?)$`)
var db *gorm.DB

// Important note: field names must be capital
type Cost struct {
	gorm.Model
	Date     time.Time
	Category string
	Value    int
	Comment  string
}

func NewCost(date time.Time, category string, value int, comment string) *Cost {
	return &Cost{Date: date, Category: category, Value: value, Comment: comment}
}

func (c *Cost) String() string {
	d := c.Date.Format("2006-01-02")
	return fmt.Sprintf("%v %s %d %s", d, c.Category, c.Value, c.Comment)
}

// panic with error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// return empty string if user Input is empty
func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter cost like `2000-01-31 Category 995 Comment not more than 255 vars`")
		fmt.Println()
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
	if err != nil {
		fmt.Println("Error with opening file. Does this file exists and have the correct permissions?")
	}
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
	// YYYY-MM-DD Category Value Comment
	trimmed := strings.TrimSpace(s)

	if validCost.MatchString(trimmed) {
		return true
	} else {
		return false
	}
}

func SaveCostsToDB(file string) {
	// Create slice of Costs from file
	costs := make([]*Cost, 0)
	costs = CostsFromFile(file)
	for _, v := range costs {
		fmt.Println(v)
		// Write costs to database
		db.Create(v)
	}
}

func SaveCostToDB(c *Cost) {
	db.Create(c)
}

func GetAllCosts(db *gorm.DB) []*Cost {
	var c []*Cost
	db.Find(&c)
	return c
}

func GetCategories(db *gorm.DB) []string {
	var categories []string
	db.Table("costs").Pluck("distinct(category)", &categories)
	return categories
}

func main() {
	// connect to database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=dinero dbname=dinero password=TodayIsTheBestDay sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// create table schema from struct if not created
	db.AutoMigrate(&Cost{})

	var c []*Cost
	c = GetAllCosts(db)
	for _, cb := range c {
		if cb.Category == "car" {
			fmt.Println(cb)
		}
	}

	categories := GetCategories(db)
	for _, cat := range categories {
		fmt.Println(cat)
	}

/// Save from file to database
	file := "months/march"
	SaveCostsToDB(file)

	//// Create cost by hand
	//c := CostFromTerminal()
	//SaveCostToDB(c)
}
