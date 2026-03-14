package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Processor struct {
	Config
	Total   int
	Success int
	Errors  []string
}

func NewProcessor(cfg Config) *Processor {
	return &Processor{Config: cfg}
}

func (p *Processor) Process(logCh chan<- string) {
	defer close(logCh)

	var files []string
	err := filepath.Walk(p.InputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if !p.Recursive && path != p.InputPath {
				return filepath.SkipDir
			}
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if !supportedExt[ext] {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		p.Errors = append(p.Errors, fmt.Sprintf("walk error: %v", err))
		if logCh != nil {
			logCh <- fmt.Sprintf("ERROR: %v", err)
		}
		return
	}

	p.Total = len(files)
	if logCh != nil {
		logCh <- fmt.Sprintf("Found %d images", p.Total)
	}

	var wg sync.WaitGroup
	workers := 4
	sem := make(chan struct{}, workers)

	for _, f := range files {
		wg.Add(1)
		go func(filePath string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			fi, err := os.Open(filePath)
			if err != nil {
				p.Errors = append(p.Errors, fmt.Sprintf("%s: open: %v", filePath, err))
				if logCh != nil {
					logCh <- fmt.Sprintf("ERROR %s: open failed", filePath)
				}
				return
			}
			defer fi.Close()

			img, _, err := image.Decode(fi)
			if err != nil {
				p.Errors = append(p.Errors, fmt.Sprintf("%s: decode: %v", filePath, err))
				if logCh != nil {
					logCh <- fmt.Sprintf("ERROR %s: decode failed", filePath)
				}
				return
			}

			if p.FromWidth > 0 || p.FromHeight > 0 {
				bounds := img.Bounds()
				w, h := bounds.Dx(), bounds.Dy()
				if (p.FromWidth > 0 && w != p.FromWidth) || (p.FromHeight > 0 && h != p.FromHeight) {
					if logCh != nil {
						logCh <- fmt.Sprintf("SKIP %s: dimensions %dx%d don't match required %dx%d",
							filePath, w, h, p.FromWidth, p.FromHeight)
					}
					return
				}
			}

			resized := resizeImage(img, p.ToWidth, p.ToHeight)

			rel, err := filepath.Rel(p.InputPath, filePath)
			if err != nil {
				rel = filepath.Base(filePath)
			}
			outPath := filepath.Join(p.OutputPath, rel)
			outPath = strings.TrimSuffix(outPath, filepath.Ext(outPath)) + ".png"

			if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
				p.Errors = append(p.Errors, fmt.Sprintf("%s: mkdir: %v", filePath, err))
				if logCh != nil {
					logCh <- fmt.Sprintf("ERROR %s: cannot create output dir", filePath)
				}
				return
			}

			fout, err := os.Create(outPath)
			if err != nil {
				p.Errors = append(p.Errors, fmt.Sprintf("%s: create: %v", filePath, err))
				if logCh != nil {
					logCh <- fmt.Sprintf("ERROR %s: cannot create output file", filePath)
				}
				return
			}
			defer fout.Close()

			if err := png.Encode(fout, resized); err != nil {
				p.Errors = append(p.Errors, fmt.Sprintf("%s: encode: %v", filePath, err))
				if logCh != nil {
					logCh <- fmt.Sprintf("ERROR %s: encode failed", filePath)
				}
				return
			}

			p.Success++
			if logCh != nil {
				logCh <- fmt.Sprintf("OK %s -> %s", filePath, outPath)
			}
		}(f)
	}

	wg.Wait()
	if logCh != nil {
		logCh <- fmt.Sprintf("Done. Success: %d/%d, Errors: %d", p.Success, p.Total, len(p.Errors))
	}
}