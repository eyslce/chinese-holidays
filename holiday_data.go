package chinese_holidays

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	HolidayStatusNormal = iota
	HolidayStatusWeekend
	HolidayStatusWorkday
	HolidayStatusHoliday
)

type holidayData struct {
	Date   string `json:"date"`
	Year   int    `json:"year"`
	Month  int    `json:"month"`
	Day    int    `json:"day"`
	Status int    `json:"status"`
}

var LoadHolidayDataByYear = func(savePath string, year int) (map[string]int, error) {
	path := filepath.Join(savePath, fmt.Sprintf("%d.json", year))
	file, err := os.Open(path)
	if err == nil {
		data, err := readFromFile(file)
		if err != nil {
			return nil, err
		}
		return parseHolidayData(data)
	}
	if !os.IsNotExist(err) {
		return nil, err
	}
	bytes, err := downloadHolidayData(year)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path, bytes, fs.ModePerm)
	if err != nil {
		return nil, err
	}
	return parseHolidayData(bytes)
}

func parseHolidayData(data []byte) (map[string]int, error) {
	var holidayData []holidayData
	err := json.Unmarshal(data, &holidayData)
	if err != nil {
		return nil, err
	}
	result := make(map[string]int)
	for _, holiday := range holidayData {
		result[holiday.Date] = holiday.Status
	}
	return result, nil
}

func readFromFile(file *os.File) ([]byte, error) {
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func downloadHolidayData(year int) ([]byte, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := fmt.Sprintf("http://api.haoshenqi.top/holiday?date=%d", year)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}
