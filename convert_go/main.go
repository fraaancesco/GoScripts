package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

)

func  downloadMp4Video(url string, output string ) (string, error) {
	downloadDir := "download"
	output = filepath.Join(downloadDir, output)
	cmd := exec.Command("yt-dlp", 
						"-f", 
						"bestvideo+bestaudio/best",
						"-o",
						output, url)
	err := os.Chmod(output, 0644)
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	return output, nil
}


func downloadMp3(url string, output string) (string, error) {
	cmd := exec.Command("yt-dlp", "--extract-audio", "--audio-format", "mp3", output, url)
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return output, nil
}

func convertToMP3(inputFile string, outputFile string) (string, error) {
	downloadDir := "download"
    if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
        if err := os.Mkdir(downloadDir, os.ModePerm); err != nil {
            return "", fmt.Errorf("Unable to create folder %s: %v", downloadDir, err)
        }
    }
	output := filepath.Join(downloadDir, outputFile + ".mp3")
	fmt.Printf("inputFile %s\n", inputFile)
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-q:a", "0", "-map", "a", output)
	outputBytes, err := cmd.CombinedOutput()
	err = os.Chmod(output, 0644)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", fmt.Errorf("ffmpeg error: %s", outputBytes)
	}
	return output, nil
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Usage: go run convert.go  ")
		return
	}
	var typeConvert string
	var url string 
	var outputNameFile string 
    fmt.Print("Press 1 to convert to MP3 otherwise press 2 (mp4) ")
    fmt.Scanln(&typeConvert)

		if typeConvert == "2" {
			fmt.Print("Enter url ")
			fmt.Scanln(&url)
			fmt.Print("Enter name output file")
			fmt.Scanln(&outputNameFile)
			fmt.Println("Downloading video...")
			videoFile, err := downloadMp4Video(url,outputNameFile)
			if err != nil {
				fmt.Printf("Failed to download video: %v\n", err)
				return
			}
			fmt.Println("Downloaded video %d, you can find it on download..",videoFile)
			
		} else if typeConvert == "1" {
			fmt.Print("Enter url ")
			fmt.Scanln(&url)
			fmt.Print("Enter nome file")
			fmt.Scanln(&outputNameFile)
			audioFile, err := downloadMp3(url,outputNameFile)
			if err != nil {
				fmt.Printf("Failed to download video: %v\n", err)
				return
			}
			fmt.Println("%d Scaricato, lo trovi su download..", audioFile)
		} else {
			return
		}
			
			
}