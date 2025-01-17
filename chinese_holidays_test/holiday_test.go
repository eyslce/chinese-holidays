package chinese_holidays_test

import (
	"testing"

	chineseholidays "github.com/eyslce/chinese-holidays"
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
	h := &chineseholidays.Holiday{
		Cfg: &chineseholidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chineseholidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-02": chineseholidays.HolidayStatusNormal,
		"2025-01-04": chineseholidays.HolidayStatusWeekend,
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
	h := &chineseholidays.Holiday{
		Cfg: &chineseholidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chineseholidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-01": chineseholidays.HolidayStatusHoliday,
		"2025-01-02": chineseholidays.HolidayStatusNormal,
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
	h := &chineseholidays.Holiday{
		Cfg: &chineseholidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chineseholidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-01": chineseholidays.HolidayStatusHoliday,
		"2025-01-02": chineseholidays.HolidayStatusNormal,
		"2025-01-04": chineseholidays.HolidayStatusWeekend,
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
	h := &chineseholidays.Holiday{
		Cfg: &chineseholidays.Config{
			SavePath: "mock_path",
		},
	}

	// 创建 mock 对象
	mockHolidayData := new(MockHolidayData)
	// 替换 loadHolidayDataByYear 方法
	chineseholidays.LoadHolidayDataByYear = mockHolidayData.loadHolidayDataByYear

	// 设置预期返回值
	mockHolidayData.On("loadHolidayDataByYear", "mock_path", 2025).Return(map[string]int{
		"2025-01-01": chineseholidays.HolidayStatusHoliday,
		"2025-01-02": chineseholidays.HolidayStatusNormal,
		"2025-01-26": chineseholidays.HolidayStatusWorkday,
		"2025-01-04": chineseholidays.HolidayStatusWeekend,
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

//func TestNewHoliday(t *testing.T) {
//	h := chinese_holidays.NewHoliday(chinese_holidays.WithSavePath("/tmp/"))
//	isHoliday, err := h.IsHoliday("2025-01-01")
//	assert.NoError(t, err)
//	assert.True(t, isHoliday)
//	isWeekday, err := h.IsWeekday("2025-01-26")
//	assert.NoError(t, err)
//	assert.True(t, isWeekday)
//}
