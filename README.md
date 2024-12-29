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
  - Quickly generate templates, update task statuses, and add new projects via CLI.

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

### Update Retrospective
Update the status of tasks from the prior week:
```bash
./cyclewise retro --file weekly-plan-2024-01-01.md
```

### Add a New Project
Add a project to the configuration file:
```bash
./cyclewise add-project --name "Node Visualizer" --roadmap "/roadmaps/node-visualizer.md"
```

### Export JSON
Export the weekly retrospective as JSON:
```bash
./cyclewise export-json --file weekly-retro-2024-01-01.md
```

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
- [x] Generate Markdown templates for plans and retrospectives.
- [x] CLI for adding projects and updating tasks.
- [ ] JSON export for analytics.

### Future Enhancements
- [ ] Web-based dashboard for progress visualization.
- [ ] Integration with third-party tools like Trello or Notion.
- [ ] Centralized JSON storage for long-term task tracking.

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

---

## License

CycleWise is licensed under the [MIT License](LICENSE).

---

## Contact

For questions or feedback, please contact cyclewise@gmailisprofessional.com
