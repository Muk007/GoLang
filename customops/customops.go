package customops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// This function writes data/balance to a file.
// To be exported...
func WriteToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

// This function reads data/balance from the file.
// To be exported...
func ReadFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("couldn't find file")
	}
	stringConvertedData := string(data)
	float64ConvertedData, err := strconv.ParseFloat(stringConvertedData, 64)
	if err != nil {
		return 0, errors.New("failed to convert into float64")
	}
	return float64ConvertedData, nil
}
