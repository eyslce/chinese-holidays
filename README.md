<div align="center">

[![CI](https://github.com/ppdxzz/go-holiday/actions/workflows/ci.yml/badge.svg)](https://github.com/ppdxzz/go-holiday/actions/workflows/ci.yml)
![Release](https://img.shields.io/github/v/release/eyslce/chinese-holidays?label=Release&color=blue)
![GitHub repo size](https://img.shields.io/github/repo-size/eyslce/chinese-holidays?label=Size&color=green)
</div>


### 📖 项目介绍
获取中国节假日信息
是一个golang版本的获取中国节假日信息的sdk,可以很容易判断一个日期是否是工作日或者节假日


### 🚀 快速开始
```shell
go get -u github.com/eyslce/chinese-holidays
```
#### 判断节假日以及工作日

```go
package main

import (
	"fmt"

	chineseholidays "github.com/eyslce/chinese-holidays"
)

func main() {
	h := chineseholidays.NewHoliday(chineseholidays.WithSavePath("/tmp/"))
	//IsHoliday 判断是否节假日（包括法定假期和周末），入参表示`2006-01-02`的日期字符串
	isHoliday, err := h.IsHoliday("2025-01-01")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("日期%s是否为节假日(包含双休)：%t\n", "2025-01-01", isHoliday)

	// IsWeekday 判断是否工作日（包括普通工作日和补班），入参表示`2006-01-02`的日期字符串
	isWeekday, err := h.IsWeekday("2025-01-26")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("日期%s是否为工作日（包括普通工作日和补班）：%t\n", "2025-01-26", isWeekday)

	// IsWeekend 判断是否周末，入参表示`2006-01-02`的日期字符串
	isWeekend, err := h.IsWeekend("2025-01-04")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("日期%s是否为周末：%t\n", "2025-01-04", isWeekend)
}
```
_运行结果：_
```markdown
日期2025-01-01是否为节假日(包含双休)：true
日期2025-01-26是否为工作日（包括普通工作日和补班）：true
日期2025-01-04是否为周末：true
```

#### 更多用法请参考
[holiday_test.go](chinese_holidays_test/holiday_test.go)

### ⚠️ 注意事项
- 目前只支持2025年以后的数据，之前的数据暂时无法获取。
- 数据源来自第三方网站，后期会考虑增加更多数据源。