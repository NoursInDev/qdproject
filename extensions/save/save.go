package save

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func SaveFloorToFile(floorContent [][]int) error {
	// date & heure actuelle
	currentTime := time.Now()
	fileName := fmt.Sprintf("%d-%02d-%02d-%02d-%02d.txt", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute())

	filePath := fmt.Sprintf("../floor-files/%s", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range floorContent {
		rowStr := strings.Trim(fmt.Sprint(strings.ReplaceAll(fmt.Sprint(row), " ", "")), "[]") // Convertit la ligne en une chaîne et supprime les crochets
		_, err := fmt.Fprintln(file, rowStr)
		if err != nil {
			return err
		}
	}

	return nil
}
