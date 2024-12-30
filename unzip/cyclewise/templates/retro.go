package templates

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strings"

    "github.com/yuin/goldmark"
    "github.com/yuin/goldmark/ast"
    "github.com/yuin/goldmark/parser"
    "github.com/yuin/goldmark/text"
    "github.com/yuin/goldmark/util"
)

type Task struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Status   string `json:"status"`
    Project  string `json:"project"`
}

type Retro struct {
    Week      string  `json:"week"`
    Tasks     []Task  `json:"tasks"`
    Projects  []string `json:"projects"`
}

func GenerateRetro(file string) {
    // Read the Markdown file
    content, err := ioutil.ReadFile(file)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Parse Markdown using goldmark
    md := goldmark.New()
    reader := text.NewReader(content)
    doc := md.Parser().Parse(reader)

    retro := Retro{}
    retro.Tasks = []Task{}
    retro.Projects = []string{}

    // Traverse the Markdown AST to extract tasks
    ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
        if !entering {
            return ast.WalkContinue, nil
        }

        if n.Kind() == ast.KindHeading {
            heading := n.(*ast.Heading)
            if heading.Level == 2 {
                text := string(heading.Text(content))
                retro.Projects = append(retro.Projects, text)
            }
        }

        if n.Kind() == ast.KindList {
            list := n.(*ast.List)
            for li := list.FirstChild(); li != nil; li = li.NextSibling() {
                if li.Kind() == ast.KindListItem {
                    item := li.(*ast.ListItem)
                    taskText := strings.TrimSpace(string(item.Text(content)))
                    task := Task{
                        ID:     "", // Generate task ID later if needed
                        Name:   taskText,
                        Status: "Pending",
                    }
                    retro.Tasks = append(retro.Tasks, task)
                }
            }
        }
        return ast.WalkContinue, nil
    })

    retroJSON, err := json.MarshalIndent(retro, "", "  ")
    if err != nil {
        fmt.Println("Error creating JSON:", err)
        return
    }

    // Write the JSON to a file
    jsonFile := strings.Replace(file, ".md", ".json", 1)
    if err := ioutil.WriteFile(jsonFile, retroJSON, 0644); err != nil {
        fmt.Println("Error writing JSON file:", err)
        return
    }

    fmt.Printf("Retro JSON created: %s\n", jsonFile)
}
