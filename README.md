# Tutor - Educational Slides Generator

A terminal-based presentation system designed for educators to create and display code-focused lessons with syntax highlighting and interactive navigation.

## Overview

This application transforms educational content into an engaging terminal-based presentation experience. Perfect for programming courses, technical workshops, and code reviews, it allows educators to present code snippets, data structures, and technical concepts in a visually appealing format with color-coded elements.

## Features

### 🎯 **Educational Focus**
- **Code-Centric Design**: Optimized for displaying programming concepts and code examples
- **Syntax Highlighting**: Color-coded code blocks for better readability
- **Structured Content**: Hierarchical organization with slides, tabs, and nested content

### 🖥️ **Terminal Interface**
- **Modern TUI**: Clean, professional terminal user interface using tview
- **Cross-Platform**: Works on Windows, macOS, and Linux terminals
- **Responsive Layout**: Dynamic layout with flexible proportions

### 📊 **Content Types**
- **Code Display**: Syntax-highlighted code blocks with customizable formatting
- **Data Tables**: Formatted tables with color-coded headers and cells
- **Multi-Tab Navigation**: Tabbed interface within slides for organized content
- **Comments & Annotations**: Support for explanatory text and notes

### ⌨️ **Keyboard Navigation**
- **Ctrl+N / Ctrl+P**: Navigate between slides
- **Alt+N / Alt+P**: Navigate between tabs within a slide
- **Alt+Left / Alt+Right**: Switch focus between content areas
- **Mouse Support**: Full mouse interaction support

## Quick Start

### Prerequisites
- Go 1.24.2 or later
- SQLite support (included via modernc.org/sqlite)

### Installation

1. **Clone or download the project**:
   ```bash
   # If using git
   git clone <repository-url>
   cd tutor
   
   # Or download and extract the project files
   ```

2. **Build the application**:
   ```bash
   go build -o tutor
   ```

3. **Initialize the database**:
   ```bash
   # The application will automatically create the database on first run
   ./tutor "your_item_name"
   ```

### Database Structure

The application automatically creates and manages a SQLite database with the following structure:

- **PROJECT**: Project definitions and metadata
- **ITEM**: Individual lessons or topics within projects  
- **SLIDE**: Individual presentation slides with code content
- **TAB**: Tabbed content within slides for organized information

### Usage

1. **Start the application**:
   ```bash
   ./tutor "lesson_name"
   ```

2. **Navigate**:
   - Use Ctrl+N/Ctrl+P to move between slides
   - Use Alt+N/Alt+P to switch between tabs within slides
   - Use Alt+Left/Alt+Right to change focus areas

## Content Creation

### Database Population

The application expects content to be populated in the SQLite database. Here's the basic structure:

```sql
-- Example content structure
INSERT INTO PROJECT (name) VALUES ('Programming Course');

INSERT INTO ITEM (id_project, name, comment) 
VALUES (1, 'Introduction to Go', 'Basic Go programming concepts');

INSERT INTO SLIDE (id_item, num, name, content, content_type, direct)
VALUES (1, 1, 'Hello World Example', 'package main\n\nimport "fmt"\n\nfunc main() {\n    fmt.Println("Hello, World!")\n}', 'code', 'column');

INSERT INTO TAB (id_slide, num, name, content, content_type)
VALUES (1, 1, 'Explanation', 'This is a simple Go program that prints Hello, World!', 'text');
```

### Content Types

1. **Code Content**: Set `content_type = 'code'` for syntax-highlighted display
2. **Table Content**: Set `content_type = 'table'` with pipe-separated values
3. **Text Content**: Set `content_type = 'text'` for regular text with formatting

### Layout Options

- **Direct**: 'column' for vertical layout, 'row' for horizontal layout
- **Content Proportion**: Integer value (default: 1) for content area size
- **Page Proportion**: Integer value (default: 2) for tab area size

## Architecture

### Core Components

- **main.go**: Application entry point and argument processing
- **database.go**: Database initialization and connection management
- **app.go**: TUI application setup and global shortcuts
- **page_main.go**: Main slide navigation interface
- **page_demo.go**: Individual slide rendering and tab management
- **slide.go & tab.go**: Data structure definitions

### Key Technologies

- **tview**: Modern terminal UI framework
- **tcell**: Terminal cell handling and color support
- **modernc.org/sqlite**: Pure Go SQLite implementation

## Configuration

### Database Location
- Database files are stored in the `DBs/` directory
- Automatic directory creation if not present
- Default database name: `DB.db`

### Logging
- Application logs are written to `app.log`
- Includes error tracking and debugging information

## Keyboard Shortcuts

| Shortcut | Action |
|----------|--------|
| `Ctrl+N` | Next slide |
| `Ctrl+P` | Previous slide |
| `Alt+N` | Next tab (within current slide) |
| `Alt+P` | Previous tab (within current slide) |
| `Alt+Left` | Focus content area |
| `Alt+Right` | Focus tabs area |
| Mouse Click | Navigate and interact |

## Educational Use Cases

### Programming Courses
- **Language Syntax**: Display code examples with proper formatting
- **Algorithm Visualization**: Step-by-step code walkthrough
- **Best Practices**: Highlighted code patterns and conventions

### Technical Workshops
- **API Documentation**: Formatted endpoint examples
- **Configuration Examples**: Structured config file displays
- **Troubleshooting**: Error handling and debugging scenarios

### Code Reviews
- **Before/After Comparisons**: Tab-based code comparison
- **Refactoring Examples**: Progressive code improvements
- **Performance Analysis**: Data structure and algorithm comparisons

## Development

### Building from Source
```bash
# Development build
go build -o tutor

# Release build with optimizations
go build -ldflags="-s -w" -o tutor
```

### Dependencies
- `github.com/rivo/tview`: Terminal UI framework
- `github.com/gdamore/tcell/v2`: Terminal cell handling
- `modernc.org/sqlite`: SQLite database driver

## Troubleshooting

### Common Issues

1. **Database Not Found**: 
   - The application automatically creates the database on first run
   - Ensure write permissions in the project directory

2. **No Content Displayed**:
   - Verify the item name exists in the database
   - Check that slides and tabs are properly linked

3. **Terminal Display Issues**:
   - Ensure your terminal supports true color
   - Check terminal size (minimum recommended: 80x24)

### Debug Mode
- Check `app.log` for detailed error information
- Run with verbose output to see database initialization messages

## License

This project is designed for educational purposes. Please ensure compliance with your institution's software usage policies.

## Support

For educational use cases, consider extending the application with:
- Content import/export features
- Presenter notes and timing
- Interactive code execution
- Multiple presentation themes

---

**Built with ❤️ for educators and students**