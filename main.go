package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fsnotify/fsnotify"
)

var (
	sourceDir        = ""
	destDirMusic     = ""
	destDirVideo     = ""
	destDirImage     = ""
	destDirDocuments = ""

	imageExtensions    = []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".bmp", ".svg", ".ico"}
	videoExtensions    = []string{".mp4", ".avi", ".mov", ".mkv", ".flv", ".webm"}
	audioExtensions    = []string{".mp3", ".wav", ".aac", ".flac", ".m4a"}
	documentExtensions = []string{".pdf", ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx"}
)

func MakeUnique(dst, name string) string {
	filename := strings.TrimPrefix(name, filepath.Ext(name))
	extension := filepath.Ext(name)
	coutner := 1

	for {
		newName := filename + "(" + strconv.Itoa(coutner) + ")" + extension
		if _, err := os.Stat(filepath.Join(dst, name)); os.IsNotExist(err) {
			return newName
		}
		coutner++
	}
}

func MoveFile(dest string, path string) {
	name := filepath.Base(path)
	destPath := filepath.Join(dest, name)
	if _, err := os.Stat(destPath); err == nil {
		name = MakeUnique(dest, name)
		destPath = filepath.Join(dest, name)
	}
	if err := os.Rename(path, destPath); err != nil {
		log.Printf("Error moving file %s: %v", name, err)
	} else {
		log.Printf("Moved file: %s", name)
	}
}

func checkAndMoveFile(path string) {
	name := filepath.Base(path)

	for _, ext := range videoExtensions {
		if strings.HasSuffix(strings.ToLower(name), ext) {
			MoveFile(destDirVideo, path)
		}
	}

	for _, ext := range audioExtensions {
		if strings.HasSuffix(strings.ToLower(name), ext) {
			MoveFile(destDirMusic, path)
		}
	}

	for _, ext := range documentExtensions {
		if strings.HasSuffix(strings.ToLower(name), ext) {
			MoveFile(destDirDocuments, path)
		}
	}

	for _, ext := range imageExtensions {
		if strings.HasSuffix(strings.ToLower(name), ext) {
			MoveFile(destDirImage, path)
		}
	}

}

func watchDirectory() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {

		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				checkAndMoveFile(event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Error:", err)
		}
	}()
	err = watcher.Add(sourceDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	log.Println("Starting file organizer...")
	watchDirectory()
}
