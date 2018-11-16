package lucky

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/antchfx/htmlquery"
)

const urlSource = "http://www.flalottery.com/site/winningNumberSearch?searchTypeIn=number&gameNameIn=LUCKYMONEY&"

// Winner are the results of a set of numbers.
type Winner struct {
	Date    string
	Numbers [5]int
	Winners string
	Prize   string
}

// Results queries and extracts the results for the given numbers from the lotto site.
func Results(numbers [5]int) []Winner {
	completeURL := urlSource + formatQueryString(numbers)
	doc, err := htmlquery.LoadURL(completeURL)
	if err != nil {
		panic(err)
	}
	winnersSlice := make([]Winner, 0)
	// Find all news item. > TBODY > TR
	for _, n := range htmlquery.Find(doc, "//tbody/tr") {
		date := htmlquery.FindOne(n, "//td[1]")
		numbers := htmlquery.FindOne(n, "//td[2]")
		winners := htmlquery.FindOne(n, "//td[3]")
		prize := htmlquery.FindOne(n, "//td[4]")
		winner := Winner{
			htmlquery.InnerText(date),
			extractNumbers(htmlquery.InnerText(numbers)),
			htmlquery.InnerText(winners),
			htmlquery.InnerText(prize)}
		winnersSlice = append(winnersSlice, winner)
	}
	return winnersSlice
}

func formatQueryString(numbers [5]int) string {
	return fmt.Sprintf("n1In=%v&n2In=%v&n3In=%v&n4In=%v&n5In=&n6In=&pbIn=&mbIn=&lbIn=%v", numbers[0], numbers[1], numbers[2], numbers[3], numbers[4])
}

func extractNumbers(numbers string) [5]int {
	re := regexp.MustCompile("[0-9]+")
	var intarray [5]int
	for index, num := range re.FindAllString(numbers, -1) {
		value, _ := strconv.Atoi(num)
		intarray[index] = value
	}
	return intarray
}
