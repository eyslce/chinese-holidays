package chinese_holidays

import (
	"errors"
	"time"
)

type Holiday struct {
	Cfg *Config
}

func NewHoliday(opts ...Option) *Holiday {
	cfg := &Config{
		SavePath: "",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &Holiday{Cfg: cfg}
}

// IsWeekend 判断是否周末，入参表示`2006-01-02`的日期字符串
func (h *Holiday) IsWeekend(date string) (bool, error) {
	status, err := h.getHolidayStatus(date)
	if err != nil {
		return false, err
	}
	if status == HolidayStatusWeekend {
		return true, nil
	}
	return false, nil
}

// IsPublicHoliday 判断是否法定假期，入参表示`2006-01-02`的日期字符串
func (h *Holiday) IsPublicHoliday(date string) (bool, error) {
	status, err := h.getHolidayStatus(date)
	if err != nil {
		return false, err
	}
	if status == HolidayStatusHoliday {
		return true, nil
	}
	return false, nil
}

// IsHoliday 判断是否节假日（包括法定假期和周末），入参表示`2006-01-02`的日期字符串
func (h *Holiday) IsHoliday(date string) (bool, error) {
	status, err := h.getHolidayStatus(date)
	if err != nil {
		return false, err
	}
	if status == HolidayStatusHoliday || status == HolidayStatusWeekend {
		return true, nil
	}
	return false, nil
}

// IsWeekday 判断是否工作日（包括普通工作日和补班），入参表示`2006-01-02`的日期字符串
func (h *Holiday) IsWeekday(date string) (bool, error) {
	status, err := h.getHolidayStatus(date)
	if err != nil {
		return false, err
	}
	if status == HolidayStatusNormal || status == HolidayStatusWorkday {
		return true, nil
	}
	return false, nil
}

func (h *Holiday) getHolidayStatus(date string) (int, error) {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		return 0, err
	}
	year := t.Year()
	holidayDataMap, err := LoadHolidayDataByYear(h.Cfg.SavePath, year)
	if err != nil {
		return 0, err
	}
	status, ok := holidayDataMap[date]
	if ok {
		return status, nil
	}
	return 0, errors.New("holiday data not found")
}
