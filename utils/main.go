// a program to add a new solution to Readme.md
package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/template"
)

type ReadMeStruct struct {
	ID               string
	Title            string
	Description      string
	IsPremium        string
	Difficulty       string
	SolutionLink     string
	AcceptanceRate   string
	Frequency        string
	URL              string
	DiscussCount     string
	Accepted         string
	Submissions      string
	Companies        string
	RelatedTopics    string
	Likes            string
	Dislikes         string
	Rating           string
	AskedByFaang     string
	SimilarQuestions string
}

func main() {
	//get the current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	// get the solution number
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter solution number: ")

	text, _, _ := reader.ReadLine()
	solutionNumber, err := strconv.Atoi(string(text))
	if err != nil {
		fmt.Println(err)
		return
	}
	// construct a link
	link := fmt.Sprintf("[visit](https://github.com/aslon1213/My-Leetcode-Solutions/tree/main/src/solution%%20%d)", solutionNumber)
	// construct a new line
	newLine := fmt.Sprintf("| %d | %s |%s |\n", solutionNumber, "go", link)
	// write the new line to the readme file
	f, err := os.OpenFile(filepath.Join(dir, "Readme.md"), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	if _, err = f.WriteString(newLine); err != nil {
		fmt.Println(err)
		return
	}

	// create a directory for the new solution
	newDir := filepath.Join(dir, fmt.Sprintf("src/solution %d", solutionNumber))
	if err := os.Mkdir(newDir, 0755); err != nil {
		fmt.Println(err)
		return
	}
	// create a main.go file
	mainFile, err := os.Create(filepath.Join(newDir, "main.go"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// write README.md file there with information about the problem and solution
	readmeFile, err := os.Create(filepath.Join(newDir, "README.md"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// read csv
	dataset_file, err := os.Open("data/leetcode_dataset.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	var readme ReadMeStruct

	readme_template := `# Problem {{.ID}}: {{.Title}}

## Description
{{.Description}}

## Difficulty
{{.Difficulty}}

## Solution
{{.URL}}

## Statistics
- **Acceptance Rate:** {{.AcceptanceRate}}%
- **Frequency:** {{.Frequency}}
- **Likes:** {{.Likes}}
- **Dislikes:** {{.Dislikes}}
- **Rating:** {{.Rating}}

## Company Tags
{{.Companies}}

## Related Topics
{{.RelatedTopics}}

## Similar Questions
{{.SimilarQuestions}}

## Discussion
- **Discuss Count:** {{.DiscussCount}}
- **Accepted:** {{.Accepted}}
- **Submissions:** {{.Submissions}}

## Asked by FAANG
{{.AskedByFaang}}

`

	csv_dataset := csv.NewReader(dataset_file)
	records, err := csv_dataset.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	// id,title,description,is_premium,difficulty,solution_link,acceptance_rate,frequency,url,discuss_count,accepted,submissions,companies,related_topics,likes,dislikes,rating,asked_by_faang,similar_questions

	for _, record := range records {
		readme.ID = record[0]
		readme.Title = record[1]
		readme.Description = record[2]
		readme.AcceptanceRate = record[5]
		readme.Frequency = record[6]
		readme.URL = link
		readme.DiscussCount = record[8]
		readme.Accepted = record[9]
		readme.Submissions = record[10]
		readme.Companies = record[11]
		readme.RelatedTopics = record[12]
		readme.Likes = record[13]
		readme.Dislikes = record[14]
		readme.Rating = record[15]
		readme.AskedByFaang = record[16]
		readme.SimilarQuestions = record[17]
	}

	templ, err := template.New("readme").Parse(readme_template)
	if err != nil {
		fmt.Println(err)
		return
	}
	templ.Execute(readmeFile, readme)

	defer mainFile.Close()
	defer readmeFile.Close()
	defer dataset_file.Close()

}
