# Go 语言 `strconv` 包详解

`strconv` 包是 Go 语言中用于**字符串和基本数据类型之间转换**的核心包。它提供了各种函数来实现字符串与整数、浮点数、布尔值等基本类型之间的相互转换。

## 📋 strconv 包核心功能概览

| 功能类别 | 主要函数 | 描述 |
| :--- | :--- | :--- |
| **字符串 ↔ 整数** | `Atoi`, `Itoa` | 简单整数转换 |
| | `ParseInt`, `FormatInt` | 带进制控制的整数转换 |
| | `ParseUint`, `FormatUint` | 无符号整数转换 |
| **字符串 ↔ 浮点数** | `ParseFloat`, `FormatFloat` | 浮点数转换 |
| **字符串 ↔ 布尔值** | `ParseBool`, `FormatBool` | 布尔值转换 |
| **引用转换** | `Quote`, `Unquote` | 带引号的字符串处理 |
| | `QuoteToASCII`, `QuoteRune` | 特殊字符处理 |
| **追加操作** | `AppendInt`, `AppendFloat` | 向字节切片追加转换结果 |

## 🛠️ 详细功能说明

### 1. 字符串与整数转换

#### 简单转换：`Atoi` 和 `Itoa`
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    // 字符串转整数 (ASCII to Integer)
    num, err := strconv.Atoi("123")
    if err != nil {
        fmt.Println("转换失败:", err)
    } else {
        fmt.Println("转换结果:", num) // 输出: 转换结果: 123
    }
    
    // 整数转字符串 (Integer to ASCII)
    str := strconv.Itoa(456)
    fmt.Println("字符串结果:", str) // 输出: 字符串结果: 456
    
    // 错误处理示例
    _, err = strconv.Atoi("abc")
    if err != nil {
        fmt.Println("错误:", err) // 输出: 错误: strconv.Atoi: parsing "abc": invalid syntax
    }
}
```

#### 带进制控制的转换：`ParseInt` 和 `FormatInt`
```go
func advancedIntConversion() {
    // 解析不同进制的字符串
    num, err := strconv.ParseInt("1010", 2, 64) // 二进制
    if err == nil {
        fmt.Println("二进制 1010 =", num) // 输出: 二进制 1010 = 10
    }
    
    num, err = strconv.ParseInt("FF", 16, 64) // 十六进制
    if err == nil {
        fmt.Println("十六进制 FF =", num) // 输出: 十六进制 FF = 255
    }
    
    num, err = strconv.ParseInt("0755", 8, 64) // 八进制
    if err == nil {
        fmt.Println("八进制 0755 =", num) // 输出: 八进制 0755 = 493
    }
    
    // 格式化不同进制的字符串
    fmt.Println("10的二进制:", strconv.FormatInt(10, 2))   // 输出: 1010
    fmt.Println("255的十六进制:", strconv.FormatInt(255, 16)) // 输出: ff
    fmt.Println("493的八进制:", strconv.FormatInt(493, 8))   // 输出: 755
    
    // 无符号整数转换
    unum, err := strconv.ParseUint("4294967295", 10, 64)
    if err == nil {
        fmt.Println("无符号整数:", unum) // 输出: 无符号整数: 4294967295
    }
    
    fmt.Println("格式化无符号:", strconv.FormatUint(4294967295, 10)) // 输出: 4294967295
}
```

### 2. 字符串与浮点数转换

```go
func floatConversion() {
    // 字符串转浮点数
    f, err := strconv.ParseFloat("3.14159", 64) // 64位精度
    if err == nil {
        fmt.Println("浮点数:", f) // 输出: 浮点数: 3.14159
    }
    
    // 特殊值处理
    f, err = strconv.ParseFloat("NaN", 64)
    if err == nil {
        fmt.Println("NaN:", strconv.IsNaN(f)) // 输出: NaN: true
    }
    
    f, err = strconv.ParseFloat("Inf", 64)
    if err == nil {
        fmt.Println("无穷大:", strconv.IsInf(f, 1)) // 输出: 无穷大: true
    }
    
    // 浮点数转字符串
    fmt.Println("格式化为字符串:", strconv.FormatFloat(3.14159, 'f', 2, 64))
    // 输出: 格式化为字符串: 3.14
    
    fmt.Println("科学计数法:", strconv.FormatFloat(123456789.0, 'e', 2, 64))
    // 输出: 科学计数法: 1.23e+08
}
```

#### FormatFloat 的格式化选项
```go
func formatFloatOptions() {
    f := 123.456
    
    // 'f': 普通小数格式，prec 控制小数位数
    fmt.Println(strconv.FormatFloat(f, 'f', 2, 64))  // 输出: 123.46
    fmt.Println(strconv.FormatFloat(f, 'f', -1, 64)) // 输出: 123.456
    
    // 'e': 科学计数法，prec 控制小数位数
    fmt.Println(strconv.FormatFloat(f, 'e', 2, 64))  // 输出: 1.23e+02
    
    // 'g': 自动选择最紧凑的表示法
    fmt.Println(strconv.FormatFloat(f, 'g', 4, 64))  // 输出: 123.5
    fmt.Println(strconv.FormatFloat(123456789.0, 'g', -1, 64)) // 输出: 1.23456789e+08
}
```

### 3. 字符串与布尔值转换

```go
func boolConversion() {
    // 字符串转布尔值
    b, err := strconv.ParseBool("true")
    if err == nil {
        fmt.Println("布尔值:", b) // 输出: 布尔值: true
    }
    
    b, err = strconv.ParseBool("1")     // true
    b, err = strconv.ParseBool("t")     // true
    b, err = strconv.ParseBool("TRUE")  // true
    b, err = strconv.ParseBool("false") // false
    b, err = strconv.ParseBool("0")     // false
    b, err = strconv.ParseBool("f")     // false
    b, err = strconv.ParseBool("FALSE") // false
    
    // 布尔值转字符串
    fmt.Println("true 转为字符串:", strconv.FormatBool(true))  // 输出: true
    fmt.Println("false 转为字符串:", strconv.FormatBool(false)) // 输出: false
}
```

### 4. 引用转换（Quoting）

```go
func quotingExamples() {
    // 引用字符串（添加引号和转义）
    quoted := strconv.Quote("Hello, 世界!\n")
    fmt.Println("引用后的字符串:", quoted) // 输出: "Hello, 世界!\n"
    
    // 取消引用
    unquoted, err := strconv.Unquote(quoted)
    if err == nil {
        fmt.Println("取消引用后的字符串:", unquoted) // 输出: Hello, 世界!（换行）
    }
    
    // ASCII 引用（非ASCII字符会被转义）
    asciiQuoted := strconv.QuoteToASCII("Hello, 世界!")
    fmt.Println("ASCII引用:", asciiQuoted) // 输出: "Hello, \u4e16\u754c!"
    
    // 引用单个字符
    runeQuoted := strconv.QuoteRune('世')
    fmt.Println("字符引用:", runeQuoted) // 输出: '世'
}
```

### 5. 追加操作（Append）

追加函数将转换结果追加到字节切片中，适用于高性能场景。

```go
func appendExamples() {
    // 创建初始字节切片
    buf := make([]byte, 0, 20)
    
    // 追加整数
    buf = strconv.AppendInt(buf, 123, 10)
    buf = append(buf, ' ') // 添加空格
    
    // 追加浮点数
    buf = strconv.AppendFloat(buf, 3.14159, 'f', 2, 64)
    buf = append(buf, ' ')
    
    // 追加布尔值
    buf = strconv.AppendBool(buf, true)
    buf = append(buf, ' ')
    
    // 追加引用字符串
    buf = strconv.AppendQuote(buf, "hello")
    
    fmt.Println("结果:", string(buf)) // 输出: 结果: 123 3.14 true "hello"
}
```

## 💡 高级用法和最佳实践

### 1. 错误处理模式

```go
func errorHandlingPatterns() {
    // 使用多个返回值处理错误
    value, err := strconv.Atoi("not_a_number")
    if err != nil {
        if numError, ok := err.(*strconv.NumError); ok {
            fmt.Printf("错误类型: %s, 输入: %s, 原因: %s\n",
                numError.Func, numError.Num, numError.Err)
            // 输出: 错误类型: Atoi, 输入: not_a_number, 原因: invalid syntax
        }
    }
    
    // 忽略错误（不推荐，除非你确定输入有效）
    value, _ = strconv.Atoi("42") // 谨慎使用！
    fmt.Println("值:", value)
}
```

### 2. 进制转换工具

```go
func baseConversionUtils() {
    // 创建进制转换工具函数
    convertBase := func(numStr string, fromBase, toBase int) (string, error) {
        // 先解析为int64
        num, err := strconv.ParseInt(numStr, fromBase, 64)
        if err != nil {
            return "", err
        }
        // 再格式化为目标进制
        return strconv.FormatInt(num, toBase), nil
    }
    
    // 二进制转十六进制
    result, err := convertBase("101010", 2, 16)
    if err == nil {
        fmt.Println("二进制 101010 = 十六进制", result) // 输出: 二进制 101010 = 十六进制 2a
    }
    
    // 十六进制转十进制
    result, err = convertBase("FF", 16, 10)
    if err == nil {
        fmt.Println("十六进制 FF = 十进制", result) // 输出: 十六进制 FF = 十进制 255
    }
}
```

### 3. 高性能数值处理

```go
func highPerformanceProcessing() {
    // 在处理大量数值转换时，使用Append系列函数提高性能
    data := []string{"123", "456", "789", "101112"}
    
    // 传统方式（每次分配新字符串）
    var sum1 int
    for _, s := range data {
        num, _ := strconv.Atoi(s)
        sum1 += num
    }
    
    // 高性能方式（避免中间字符串分配）
    var sum2 int
    buf := make([]byte, 0, 10)
    for _, s := range data {
        // 重用缓冲区
        buf = buf[:0]
        buf = append(buf, s...)
        num, _ := strconv.ParseInt(string(buf), 10, 64)
        sum2 += int(num)
    }
    
    fmt.Println("总和:", sum1, sum2) // 输出: 总和: 103470 103470
}
```

## ⚠️ 注意事项和常见问题

### 1. 数值范围问题

```go
func rangeIssues() {
    // 整数溢出检查
    _, err := strconv.Atoi("99999999999999999999")
    if err != nil {
        fmt.Println("溢出错误:", err) // 输出: 溢出错误: strconv.Atoi: parsing "99999999999999999999": value out of range
    }
    
    // 使用ParseInt指定位宽
    num, err := strconv.ParseInt("32768", 10, 16) // 16位有符号整数最大32767
    if err != nil {
        fmt.Println("16位溢出:", err) // 输出: 16位溢出: strconv.ParseInt: parsing "32768": value out of range
    }
    
    // 正确处理位宽参数
    num, err = strconv.ParseInt("32767", 10, 16)
    if err == nil {
        fmt.Println("16位最大值:", num) // 输出: 16位最大值: 32767
    }
}
```

### 2. 浮点数精度问题

```go
func floatPrecision() {
    // 浮点数精度损失
    f, _ := strconv.ParseFloat("0.1", 64)
    fmt.Printf("0.1 的64位表示: %.20f\n", f) // 输出: 0.1 的64位表示: 0.10000000000000000555
    
    // 格式化时控制精度
    fmt.Println("保留2位小数:", strconv.FormatFloat(1.23456789, 'f', 2, 64)) // 输出: 1.23
    fmt.Println("保留6位小数:", strconv.FormatFloat(1.23456789, 'f', 6, 64)) // 输出: 1.234568
}
```

### 3. 特殊值处理

```go
func specialValues() {
    // 处理NaN和无穷大
    nan, _ := strconv.ParseFloat("NaN", 64)
    posInf, _ := strconv.ParseFloat("+Inf", 64)
    negInf, _ := strconv.ParseFloat("-Inf", 64)
    
    fmt.Println("IsNaN:", strconv.IsNaN(nan))           // 输出: true
    fmt.Println("IsPositiveInf:", strconv.IsInf(posInf, 1)) // 输出: true
    fmt.Println("IsNegativeInf:", strconv.IsInf(negInf, -1)) // 输出: true
    
    // 格式化特殊值
    fmt.Println("NaN格式化为字符串:", strconv.FormatFloat(nan, 'f', -1, 64)) // 输出: NaN
    fmt.Println("Inf格式化为字符串:", strconv.FormatFloat(posInf, 'f', -1, 64)) // 输出: +Inf
}
```

## 🔄 与其他包配合使用

### 1. 与 `fmt` 包对比

```go
func compareWithFmt() {
    num := 123
    
    // fmt.Sprintf 更灵活但性能较低
    str1 := fmt.Sprintf("%d", num)
    fmt.Println("fmt结果:", str1) // 输出: 123
    
    // strconv.Itoa 性能更高
    str2 := strconv.Itoa(num)
    fmt.Println("strconv结果:", str2) // 输出: 123
    
    // 对于复杂格式化，使用fmt
    str3 := fmt.Sprintf("数值: %05d", num) // 输出: 数值: 00123
    fmt.Println("格式化结果:", str3)
    
    // 对于简单转换，使用strconv
    str4 := strconv.FormatInt(int64(num), 10) // 简单但冗长
    fmt.Println("简单转换:", str4) // 输出: 123
}
```

### 2. 与 `strings` 包配合

```go
func withStringsPackage() {
    // 处理包含数字的字符串
    input := "id:123,name:456,value:789"
    
    // 先分割字符串
    parts := strings.Split(input, ",")
    for _, part := range parts {
        pair := strings.Split(part, ":")
        if len(pair) == 2 {
            key := pair[0]
            valueStr := pair[1]
            
            // 转换数值部分
            value, err := strconv.Atoi(valueStr)
            if err == nil {
                fmt.Printf("%s: %d (类型: %T)\n", key, value, value)
            } else {
                fmt.Printf("%s: %s (字符串)\n", key, valueStr)
            }
        }
    }
}
```

## 📊 最佳实践总结

1. **选择正确的函数**：
   - 简单整数转换：`Atoi` / `Itoa`
   - 需要控制进制：`ParseInt` / `FormatInt`
   - 高性能场景：`Append` 系列函数

2. **始终处理错误**：
   ```go
   value, err := strconv.Atoi(input)
   if err != nil {
       // 正确处理错误
       return
   }
   ```

3. **注意数值范围**：
   - 使用合适的位宽参数
   - 检查溢出错误

4. **考虑性能**：
   - 大量转换时使用 `Append` 系列函数
   - 重用缓冲区和变量

5. **正确处理特殊值**：
   - NaN、无穷大等特殊浮点数值
   - 空字符串和无效输入

`strconv` 包是 Go 语言中处理基本数据类型转换的基石，掌握它的使用对于编写健壮和高性能的 Go 程序至关重要。