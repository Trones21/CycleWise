# CycleWise

**CycleWise** is a streamlined tool for managing weekly plans and retrospectives across multiple projects. Designed with simplicity and efficiency in mind, CycleWise enables you to track tasks, reflect on progress, and iterate toward success.

---

## Features

- **Weekly Plans and Retrospectives**:
  - Generate structured Markdown templates for weekly plans and retrospectives.
  - Track tasks with statuses like `Completed`, `Progress Made`, or `No Progress`.
  
- **Task Management**:
  - Assign unique IDs to tasks for persistence across weeks.
  - Automatically carry over unfinished tasks to the next week.

- **Roadmap Integration**:
  - Link project-specific roadmaps to weekly plans for alignment with long-term goals.

- **JSON Export**:
  - Export plans and retrospectives as JSON for analytics and integration with other tools.

- **Command-Line Interface (CLI)**:
  - Quickly generate templates and add new projects via CLI.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/cyclewise.git
   cd cyclewise
   ```

2. Install dependencies (if any):
   ```bash
   go mod tidy
   ```

3. Build the executable:
   ```bash
   go build -o cyclewise
   ```

---

## Usage

### Generate a Weekly Plan
Generate a new weekly plan with tasks from the previous week:
```bash
./cyclewise plan --date 2024-01-01
```
- `--date`: Specify the week (defaults to the current date).

### Retrospective Workflow
The CLI does not directly update the retro file. Instead, users edit the retrospective Markdown manually. CycleWise assists in the following ways:

#### Convert Plan File to Retro Template 
Take the current week's plan and add a `Retro Status` column to the task tables.
   - Status options: `Complete`, `Made Progress`, `Did Not Make Progress`.

#### Retro JSON Export
`goldmark` is used to parse the retro markdown file and create a JSON export.

- Parse the plan to retrieve existing project and task IDs.
  - Generate new IDs for projects and tasks only if they are missing.

#### Export Retrospective to JSON
```bash
./cyclewise retro --file weekly-retro-2024-01-01.md
```
- Parses the retrospective file and outputs a JSON representation.

### Generate a New Weekly Plan from Previous Week
The CLI prompts you to select which projects to include in the new week’s plan:
```bash
./cyclewise plan --generate-new
```
1. The tool displays a list of projects from the previous plan:
   ```
   Which projects would you like to track this week?
   [1] Project A (Last updated: 2024-01-01)
   [2] Project B (Last updated: 2024-01-01)
   [3] Project C (Last updated: 2024-01-01)
   Enter the numbers (comma-separated): 1,3
   ```
2. For each selected project:
   - Tasks marked `Complete` are excluded.
   - Tasks marked `Made Progress` or `Did Not Make Progress` are carried over with their statuses reset to `Pending`.
3. The new weekly plan includes metadata:
   - `lastUpdated`: The date of the last retro update.
   - `includedInCurrentPlan`: Boolean to indicate if the project is part of the current plan.

---

## Configuration

CycleWise uses a `config.json` file to store project metadata:

```json
{
  "projects": [
    {
      "name": "Node Visualizer",
      "roadmap_link": "/roadmaps/node-visualizer.md",
      "date_added": "2024-12-29"
    },
    {
      "name": "Link Extractor",
      "roadmap_link": "",
      "date_added": "2024-12-29"
    }
  ]
}
```

Edit the configuration file manually or use the `add-project` CLI command.

---

## Roadmap

### MVP Features
- [ ] Generate Markdown templates for plans and retrospectives.
- [ ] CLI for adding projects and updating tasks.
- [ ] JSON export for analytics.
- [ ] Retrospective JSON export with task IDs and statuses.
- [ ] Integration with `goldmark` for Markdown parsing.
- [ ] Prompt for project selection and carryover tasks.

### Future Enhancements
- TBD
