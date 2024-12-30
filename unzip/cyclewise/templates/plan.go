package templates

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
    "time"
)

type Task struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Status   string `json:"status"`
    LastStatus string `json:"last_status"` // Status from the retro
}

type Project struct {
    ID                string  `json:"id"`
    Name              string  `json:"name"`
    Roadmap           string  `json:"roadmap"`
    LastUpdated       string  `json:"last_updated"`
    IncludedInCurrentPlan bool `json:"included_in_current_plan"`
    Tasks             []Task  `json:"tasks"`
}

func GeneratePlan() {
    // Load the previous week's JSON export
    prevWeekFile := "weekly-plan.json"
    content, err := ioutil.ReadFile(prevWeekFile)
    if err != nil {
        fmt.Println("Error reading previous week's JSON:", err)
        return
    }

    var projects []Project
    if err := json.Unmarshal(content, &projects); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    // Display projects and prompt for selection
    fmt.Println("Which projects would you like to track this week?")
    for i, project := range projects {
        fmt.Printf("[%d] %s (Last updated: %s)\n", i+1, project.Name, project.LastUpdated)
    }
    fmt.Print("Enter the numbers (comma-separated): ")

    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    selectedIndexes := strings.Split(input, ",")

    selectedProjects := []Project{}
    for _, indexStr := range selectedIndexes {
        index, err := strconv.Atoi(indexStr)
        if err != nil || index < 1 || index > len(projects) {
            fmt.Println("Invalid selection:", indexStr)
            continue
        }
        project := projects[index-1]
        project.IncludedInCurrentPlan = true
        selectedProjects = append(selectedProjects, project)
    }

    // Update tasks for the new plan
    for i := range selectedProjects {
        updatedTasks := []Task{}
        for _, task := range selectedProjects[i].Tasks {
            if task.Status != "Complete" {
                task.LastStatus = task.Status
                task.Status = "Pending"
                updatedTasks = append(updatedTasks, task)
            }
        }
        selectedProjects[i].Tasks = updatedTasks
    }

    // Generate new weekly plan
    date := time.Now().Format("2006-01-02")
    filename := fmt.Sprintf("weekly-plan-%s.md", date)
    file, err := os.Create(filename)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    file.WriteString(fmt.Sprintf("# Weekly Planning: %s\n\n", date))
    for _, project := range selectedProjects {
        file.WriteString(fmt.Sprintf("## Project: %s\n", project.Name))
        file.WriteString(fmt.Sprintf("**Project ID:** `%s`\n", project.ID))
        if project.Roadmap != "" {
            file.WriteString(fmt.Sprintf("### Roadmap: [Link to Roadmap](%s)\n", project.Roadmap))
        }
        file.WriteString("### Goals for the Week:\n")
        file.WriteString("| Task ID | Task                         | Status    |\n")
        file.WriteString("|---------|------------------------------|-----------|\n")
        for _, task := range project.Tasks {
            file.WriteString(fmt.Sprintf("| %s     | %s | %s |\n", task.ID, task.Name, task.Status))
        }
        file.WriteString("\n---\n\n")
    }

    fmt.Printf("New weekly plan created: %s\n", filename)

    // Export updated JSON
    newJSONFile := fmt.Sprintf("weekly-plan-%s.json", date)
    updatedJSON, err := json.MarshalIndent(selectedProjects, "", "  ")
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }
    if err := ioutil.WriteFile(newJSONFile, updatedJSON, 0644); err != nil {
        fmt.Println("Error writing JSON file:", err)
        return
    }
    fmt.Printf("Updated JSON export: %s\n", newJSONFile)
}
