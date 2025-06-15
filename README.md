# Template Scaffold

A modern, flexible project scaffolding tool that supports both interactive and non-interactive modes. Create new projects from templates with ease!

## Features

- 🎯 Multiple template support
- 🔄 Interactive mode with prompts
- 📝 Non-interactive mode using YAML/JSON config
- 🛠️ Go templates with variable substitution
- 🔍 Dry-run mode with diff support
- 📦 Local and remote template sources

## Installation

```bash
# Clone the repository
git clone https://github.com/khareyash05/template-scaffold.git

# Navigate to the project directory
cd template-scaffold

# Build the project
go build

# Install globally (optional)
go install
```

## Usage

### Basic Commands

```bash
# Show help
scaffold --help

# List available templates
scaffold list

# Create a new project (interactive mode)
scaffold init <template-name>

# Create a new project (non-interactive mode)
scaffold init <template-name> --config config.yaml
```

### Command Options

```bash
# Specify template source
scaffold init <template-name> --source ./templates

# Set output directory
scaffold init <template-name> --output ./my-project

# Dry run (show what would be created)
scaffold init <template-name> --dry-run

# Show diff in dry run mode
scaffold init <template-name> --dry-run --diff
```

## Configuration

### Interactive Mode

When using interactive mode, you'll be prompted for each variable defined in the template's `questions.yaml`:

```yaml
- name: ProjectName
  prompt: "Project name"
  default: "my-app"
- name: UseDocker
  prompt: "Include Dockerfile? (y/n)"
  default: false
```

### Non-Interactive Mode

Create a YAML or JSON config file with your variables:

```yaml
# config.yaml
variables:
  ProjectName: my-awesome-app
  UseDocker: true
```

```json
// config.json
{
  "variables": {
    "ProjectName": "my-awesome-app",
    "UseDocker": true
  }
}
```

## Available Templates

### Go Empty

A basic Go project template with:
- Go module setup
- Git initialization
- README
- .gitignore

```bash
# Interactive mode
scaffold init go-empty

# Non-interactive mode
scaffold init go-empty --config examples/go-empty-config.yaml
```

### Go Cobra

A CLI application template using Cobra:
- Cobra command structure
- Viper config support
- Version command
- README with usage
- .gitignore

```bash
# Interactive mode
scaffold init go-cobra

# Non-interactive mode
scaffold init go-cobra --config examples/go-cobra-config.yaml
```

## Creating Your Own Templates

1. Create a new directory in `templates/`
2. Add a `questions.yaml` file defining your variables
3. Add your template files (use `.tmpl` extension for clarity)
4. Use Go template syntax for variable substitution

Example template structure:
```
templates/
  my-template/
    questions.yaml
    README.md.tmpl
    main.go.tmpl
    .gitignore.tmpl
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT © Yash Khare 