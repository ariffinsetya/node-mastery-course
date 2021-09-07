package main

import "github.com/ariffinsetya/micro/chopper/pkg/handler/video"

func main() {
	outputDir := "./tmp/"
	videoOnly := video.GetVideo("tsukimichi.mkv", outputDir)
	video.GetAudio("tsukimichi.mkv", outputDir)
	video.Split(videoOnly, outputDir, 300)
}
