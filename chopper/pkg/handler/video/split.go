package video

import (
	"fmt"
	"path/filepath"
	"strconv"

	sh "github.com/codeskyblue/go-sh"
)

func Split(inputVideo string, outputDir string, duration int) {
	extension := filepath.Ext(inputVideo)
	out, err := sh.Command(
		"ffmpeg",
		"-hide_banner",
		"-loglevel",
		"error",
		"-i",
		inputVideo,
		"-map",
		"0:v",
		"-f",
		"segment",
		"-segment_time",
		strconv.Itoa(duration),
		"-c:v",
		"copy",
		"-reset_timestamps",
		"1",
		outputDir+"%5d"+extension,
	).Output()
	if err != nil {
		fmt.Printf("ERROR, %v", err)
	}
	fmt.Printf("%v", out)
}

func GetVideo(inputVideo string, outputDir string) string {
	extension := filepath.Ext(inputVideo)
	basename := filepath.Base(inputVideo)
	outputFile := outputDir + basename + "_video" + extension
	out, err := sh.Command(
		"ffmpeg",
		"-hide_banner",
		"-loglevel",
		"error",
		"-i",
		inputVideo,
		"-map",
		"0:v",
		"-c:v",
		"copy",
		outputFile,
	).Output()
	if err != nil {
		fmt.Printf("ERROR, %v", err)
	}
	fmt.Printf("%v", out)
	return outputFile
}

func GetAudio(inputVideo string, outputDir string) string {
	extension := filepath.Ext(inputVideo)
	basename := filepath.Base(inputVideo)
	outputFile := outputDir + basename + "_audio" + extension
	out, err := sh.Command(
		"ffmpeg",
		"-hide_banner",
		"-loglevel",
		"error",
		"-i",
		inputVideo,
		"-map",
		"0:a",
		"-c:a",
		"copy",
		outputFile,
	).Output()
	if err != nil {
		fmt.Printf("ERROR, %v", err)
	}
	fmt.Printf("%v", out)
	return outputFile
}
