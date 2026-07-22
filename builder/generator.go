package builder

import (
	"bytes"
	"context"
	"cv-website/utils"
	"cv-website/view/homepage"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func RunStaticGenerator() {
	if err := os.RemoveAll("dist"); err != nil {
		log.Printf("builder: failed to clean dist: %v", err)
		return
	}
	if err := os.MkdirAll("dist", 0755); err != nil {
		log.Printf("builder: failed to create dist: %v", err)
		return
	}

	copyDir("assets", filepath.Join("dist", "assets"))

	data := utils.CVData
	data.Year = strconv.Itoa(time.Now().Year())

	var buf bytes.Buffer
	if err := homepage.Show(data).Render(context.Background(), &buf); err != nil {
		log.Printf("builder: render failed: %v", err)
		return
	}

	if err := os.WriteFile(filepath.Join("dist", "index.html"), buf.Bytes(), 0644); err != nil {
		log.Printf("builder: failed to write index.html: %v", err)
		return
	}

	log.Println("builder: static site generated in dist/")
}

func copyDir(src, dst string) {
	entries, err := os.ReadDir(src)
	if err != nil {
		return
	}
	if err := os.MkdirAll(dst, 0755); err != nil {
		return
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			copyDir(srcPath, dstPath)
		} else {
			srcFile, err := os.Open(srcPath)
			if err != nil {
				continue
			}
			dstFile, err := os.Create(dstPath)
			if err != nil {
				srcFile.Close()
				continue
			}
			io.Copy(dstFile, srcFile)
			srcFile.Close()
			dstFile.Close()
		}
	}
}
