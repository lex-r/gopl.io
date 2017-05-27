package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	monthAgo := make([]*github.Issue, 0)
	yearAgo := make([]*github.Issue, 0)
	moreYearAgo := make([]*github.Issue, 0)
	for _, item := range result.Items {
		if lessThanMonth(item) {
			monthAgo = append(monthAgo, item)
		} else if lessThanYear(item) {
			yearAgo = append(yearAgo, item)
		} else {
			moreYearAgo = append(moreYearAgo, item)
		}
	}

	fmt.Printf("%d тем:\n", result.TotalCount)
	if len(monthAgo) > 0 {
		fmt.Println("за последний месяц")
		for _, item := range monthAgo {
			printIssue(item)
		}
	}
	if len(yearAgo) > 0 {
		fmt.Println("за последний год")
		for _, item := range yearAgo {
			printIssue(item)
		}
	}
	if len(moreYearAgo) > 0 {
		fmt.Println("больше года назад")
		for _, item := range moreYearAgo {
			printIssue(item)
		}
	}

}

func printIssue(issue *github.Issue) {
	fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
}

func lessThanMonth(issue *github.Issue) bool {
	return time.Since(issue.CreatedAt) < time.Hour*24*30
}

func lessThanYear(issue *github.Issue) bool {
	return time.Since(issue.CreatedAt) < time.Hour*24*30*12
}
