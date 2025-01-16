package chinese_holidays_test

import (
	"chinese-holidays"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// 模拟 holidayDataMap
type MockHolidayData struct {
	mock.Mock
}

func (m *MockHolidayData) loadHolidayDataByYear(savePath string, year int) (map[string]int, error) {
	args := m.Called(savePath, year)
	return args.Get(0).(map[string]int), args.Error(1)
}

// 为 IsWeekend 编写测试用例
func TestIsWeekend(t *testing.T) {
	h := &chinese_holidays.Holiday{
		Cfg: &chinese_holidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chinese_holidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-02": chinese_holidays.HolidayStatusNormal,
		"2025-01-04": chinese_holidays.HolidayStatusWeekend,
	}, nil)

	// 测试正常工作日
	isWeekend, err := h.IsWeekend("2025-01-02")
	assert.NoError(t, err)
	assert.False(t, isWeekend)

	// 测试周末
	isWeekend, err = h.IsWeekend("2025-01-04")
	assert.NoError(t, err)
	assert.True(t, isWeekend)

	// 测试格式错误的日期
	isWeekend, err = h.IsWeekend("2025-01-077")
	assert.Error(t, err)
	assert.False(t, isWeekend)

	// 测试不存在的日期
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2024).Return(map[string]int{}, nil)
	isWeekend, err = h.IsWeekend("2025-02-31")
	assert.Error(t, err)
	assert.False(t, isWeekend)
}

// 为 IsPublicHoliday 编写测试用例
func TestIsPublicHoliday(t *testing.T) {
	h := &chinese_holidays.Holiday{
		Cfg: &chinese_holidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chinese_holidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-01": chinese_holidays.HolidayStatusHoliday,
		"2025-01-02": chinese_holidays.HolidayStatusNormal,
	}, nil)

	// 测试法定假期
	isPublicHoliday, err := h.IsPublicHoliday("2025-01-01")
	assert.NoError(t, err)
	assert.True(t, isPublicHoliday)

	// 测试普通工作日
	isPublicHoliday, err = h.IsPublicHoliday("2025-01-02")
	assert.NoError(t, err)
	assert.False(t, isPublicHoliday)

	// 测试格式错误的日期
	isPublicHoliday, err = h.IsPublicHoliday("2025-01-022")
	assert.Error(t, err)
	assert.False(t, isPublicHoliday)

	// 测试不存在的日期
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2024).Return(map[string]int{}, nil)
	isPublicHoliday, err = h.IsPublicHoliday("2025-02-31")
	assert.Error(t, err)
	assert.False(t, isPublicHoliday)
}

// 为 IsHoliday 编写测试用例
func TestIsHoliday(t *testing.T) {
	h := &chinese_holidays.Holiday{
		Cfg: &chinese_holidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chinese_holidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-01": chinese_holidays.HolidayStatusHoliday,
		"2025-01-02": chinese_holidays.HolidayStatusNormal,
		"2025-01-04": chinese_holidays.HolidayStatusWeekend,
	}, nil)

	// 测试法定假期
	isHoliday, err := h.IsHoliday("2025-01-01")
	assert.NoError(t, err)
	assert.True(t, isHoliday)

	// 测试普通工作日
	isHoliday, err = h.IsHoliday("2025-01-02")
	assert.NoError(t, err)
	assert.False(t, isHoliday)

	// 测试周末
	isHoliday, err = h.IsHoliday("2025-01-04")
	assert.NoError(t, err)
	assert.True(t, isHoliday)

	// 测试格式错误的日期
	isHoliday, err = h.IsHoliday("2025-01-077")
	assert.Error(t, err)
	assert.False(t, isHoliday)

	// 测试不存在的日期
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2024).Return(map[string]int{}, nil)
	isHoliday, err = h.IsHoliday("2025-02-31")
	assert.Error(t, err)
	assert.False(t, isHoliday)
}

// 为 IsWeekday 编写测试用例
func TestIsWeekday(t *testing.T) {
	h := &chinese_holidays.Holiday{
		Cfg: &chinese_holidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chinese_holidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-01": chinese_holidays.HolidayStatusHoliday,
		"2025-01-02": chinese_holidays.HolidayStatusNormal,
		"2025-01-26": chinese_holidays.HolidayStatusWorkday,
		"2025-01-04": chinese_holidays.HolidayStatusWeekend,
	}, nil)

	// 测试法定假期
	isWeekday, err := h.IsWeekday("2025-01-01")
	assert.NoError(t, err)
	assert.False(t, isWeekday)

	// 测试普通工作日
	isWeekday, err = h.IsWeekday("2025-01-02")
	assert.NoError(t, err)
	assert.True(t, isWeekday)

	// 测试补班
	isWeekday, err = h.IsWeekday("2025-01-26")
	assert.NoError(t, err)
	assert.True(t, isWeekday)

	// 测试周末
	isWeekday, err = h.IsWeekday("2025-01-04")
	assert.NoError(t, err)
	assert.False(t, isWeekday)

	// 测试格式错误的日期
	isWeekday, err = h.IsWeekday("2025-01-077")
	assert.Error(t, err)
	assert.False(t, isWeekday)

	// 测试不存在的日期
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2024).Return(map[string]int{}, nil)
	isWeekday, err = h.IsWeekday("2025-02-31")
	assert.Error(t, err)
	assert.False(t, isWeekday)
}

func TestNewHoliday(t *testing.T) {
	h := chinese_holidays.NewHoliday(chinese_holidays.WithSavePath("/tmp/"))
	isHoliday, err := h.IsHoliday("2025-01-01")
	assert.NoError(t, err)
	assert.True(t, isHoliday)
	isWeekday, err := h.IsWeekday("2025-01-26")
	assert.NoError(t, err)
	assert.True(t, isWeekday)
}
