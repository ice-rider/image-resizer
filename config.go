package main

import "flag"

type Config struct {
	InputPath   string
	OutputPath  string
	FromWidth   int
	FromHeight  int
	ToWidth     int
	ToHeight    int
	Recursive   bool
	TUI         bool
}

func parseFlags() Config {
	var (
		from       int
		to         int
		fromWidth  int
		fromHeight int
		toWidth    int
		toHeight   int
		inputPath  string
		outputPath string
		recursive  bool
		tui        bool
	)

	flag.StringVar(&inputPath, "input", "", "input directory or file")
	flag.StringVar(&outputPath, "output", "", "output directory")
	flag.IntVar(&fromWidth, "from-width", 0, "required source width (0 = skip check)")
	flag.IntVar(&fromHeight, "from-height", 0, "required source height (0 = skip check)")
	flag.IntVar(&toWidth, "to-width", 16, "target width")
	flag.IntVar(&toHeight, "to-height", 16, "target height")
	flag.IntVar(&from, "from", 0, "source square size (overrides from-width/from-height if they are 0)")
	flag.IntVar(&to, "to", 0, "target square size (overrides to-width/to-height if they are 0)")
	flag.BoolVar(&recursive, "recursive", false, "process subdirectories recursively")
	flag.BoolVar(&tui, "tui", false, "force TUI mode")
	flag.Parse()

	if from != 0 {
		if fromWidth == 0 {
			fromWidth = from
		}
		if fromHeight == 0 {
			fromHeight = from
		}
	}
	if to != 0 {
		if toWidth == 0 {
			toWidth = to
		}
		if toHeight == 0 {
			toHeight = to
		}
	}

	return Config{
		InputPath:   inputPath,
		OutputPath:  outputPath,
		FromWidth:   fromWidth,
		FromHeight:  fromHeight,
		ToWidth:     toWidth,
		ToHeight:    toHeight,
		Recursive:   recursive,
		TUI:         tui,
	}
}