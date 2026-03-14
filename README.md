# Batch Image Resizer

[English](#english) | [Русский](#русский)

---

## English

A powerful command-line tool and TUI (Text User Interface) application for batch resizing images. Built with Go and the Bubble Tea framework.

### Features

- 🖼️ **Batch Processing** — Resize multiple images at once
- 📁 **Recursive Mode** — Process subdirectories automatically
- 🎯 **Dimension Filtering** — Filter images by source width/height before resizing
- ⚡ **Concurrent Processing** — Multi-threaded image processing for better performance
- 🖥️ **TUI Mode** — Interactive terminal interface with real-time logs
- 📦 **Cross-Platform** — Windows, Linux, macOS (amd64 & arm64)
- 📝 **Multiple Formats** — Support for JPEG, PNG, GIF, BMP

### Installation

```bash
go build -o resizer .
```

Or download a pre-built binary from the releases page.

### Usage

#### CLI Mode

```bash
# Basic usage
./resizer -input ./images -output ./resized -to-width 800 -to-height 600

# Square resize
./resizer -input ./images -output ./resized -to 512

# With dimension filtering (only resize images that are exactly 1920x1080)
./resizer -input ./images -output ./resized -from-width 1920 -from-height 1080 -to-width 800 -to-height 600

# Recursive mode (process subdirectories)
./resizer -input ./images -output ./resized -to 1024 -recursive

# Interactive TUI mode
./resizer -tui
```

#### Command Line Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-input` | Input directory or file | (required) |
| `-output` | Output directory | (required) |
| `-from-width` | Required source width (0 = skip check) | 0 |
| `-from-height` | Required source height (0 = skip check) | 0 |
| `-to-width` | Target width | 16 |
| `-to-height` | Target height | 16 |
| `-from` | Source square size (overrides from-width/from-height) | 0 |
| `-to` | Target square size (overrides to-width/to-height) | 0 |
| `-recursive` | Process subdirectories recursively | false |
| `-tui` | Force TUI mode | false |

#### TUI Mode

The interactive terminal interface allows you to:
- Set input/output paths
- Configure resize dimensions
- Toggle recursive mode
- View real-time processing logs

**Controls:**
- `Tab` / `Shift+Tab` — Navigate between fields
- `Enter` — Start resizing
- `R` — Toggle recursive mode
- `Ctrl+C` / `Q` — Quit

### Building for All Platforms

```bash
# Build for current platform
make build

# Build for all platforms (Windows, Linux, macOS + amd64 & arm64)
make build-all

# Build for specific platform
make build-windows
make build-linux
make build-darwin
make build-windows-arm64
make build-linux-arm64
make build-darwin-arm64

# Clean build directory
make clean
```

### Project Structure

```
resizer/
├── main.go      # Entry point, CLI/TUI mode selection
├── config.go    # Configuration and flag parsing
├── core.go      # Image processor logic
├── resize.go    # Image resizing functions
├── tui.go       # TUI components and logic
├── Makefile     # Cross-platform build system
└── README.md    # Documentation
```

### License

MIT

---

## Русский

Мощный инструмент командной строки и TUI (текстовый пользовательский интерфейс) для пакетного изменения размера изображений. Создан на Go с использованием фреймворка Bubble Tea.

### Возможности

- 🖼️ **Пакетная обработка** — Изменение размера нескольких изображений одновременно
- 📁 **Рекурсивный режим** — Автоматическая обработка подкаталогов
- 🎯 **Фильтрация по размерам** — Фильтрация изображений по ширине/высоте перед изменением
- ⚡ **Многопоточность** — Параллельная обработка изображений для лучшей производительности
- 🖥️ **TUI режим** — Интерактивный терминальный интерфейс с логами в реальном времени
- 📦 **Кроссплатформенность** — Windows, Linux, macOS (amd64 & arm64)
- 📝 **Множество форматов** — Поддержка JPEG, PNG, GIF, BMP

### Установка

```bash
go build -o resizer .
```

Или скачайте готовый бинарный файл на странице релизов.

### Использование

#### Режим командной строки

```bash
# Базовое использование
./resizer -input ./images -output ./resized -to-width 800 -to-height 600

# Квадратный ресайз
./resizer -input ./images -output ./resized -to 512

# С фильтрацией по размерам (только изображения 1920x1080)
./resizer -input ./images -output ./resized -from-width 1920 -from-height 1080 -to-width 800 -to-height 600

# Рекурсивный режим (обработка подкаталогов)
./resizer -input ./images -output ./resized -to 1024 -recursive

# Интерактивный TUI режим
./resizer -tui
```

#### Флаги командной строки

| Флаг | Описание | По умолчанию |
|------|----------|--------------|
| `-input` | Входная директория или файл | (требуется) |
| `-output` | Выходная директория | (требуется) |
| `-from-width` | Требуемая ширина исходника (0 = без проверки) | 0 |
| `-from-height` | Требуемая высота исходника (0 = без проверки) | 0 |
| `-to-width` | Целевая ширина | 16 |
| `-to-height` | Целевая высота | 16 |
| `-from` | Размер квадрата исходника (переопределяет from-width/from-height) | 0 |
| `-to` | Размер квадрата результата (переопределяет to-width/to-height) | 0 |
| `-recursive` | Рекурсивная обработка подкаталогов | false |
| `-tui` | Принудительный TUI режим | false |

#### TUI Режим

Интерактивный терминальный интерфейс позволяет:
- Устанавливать пути ввода/вывода
- Настраивать размеры для изменения
- Переключать рекурсивный режим
- Просматривать логи обработки в реальном времени

**Управление:**
- `Tab` / `Shift+Tab` — Навигация между полями
- `Enter` — Запустить изменение размера
- `R` — Переключить рекурсивный режим
- `Ctrl+C` / `Q` — Выход

### Сборка для всех платформ

```bash
# Сборка для текущей платформы
make build

# Сборка для всех платформ (Windows, Linux, macOS + amd64 & arm64)
make build-all

# Сборка для конкретной платформы
make build-windows
make build-linux
make build-darwin
make build-windows-arm64
make build-linux-arm64
make build-darwin-arm64

# Очистка директории сборки
make clean
```

### Структура проекта

```
resizer/
├── main.go      # Точка входа, выбор режима CLI/TUI
├── config.go    # Конфигурация и парсинг флагов
├── core.go      # Логика процессора изображений
├── resize.go    # Функции изменения размера изображений
├── tui.go       # Компоненты и логика TUI
├── Makefile     # Кроссплатформенная система сборки
└── README.md    # Документация
```

### Лицензия

MIT
