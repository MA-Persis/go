Go 的 `math` 包提供了丰富的数学函数和常量，用于进行各种科学计算和工程计算。下面我会为你梳理它的主要功能、常用函数以及一些实用技巧。

下面是 `math` 包中一些函数类别的快速概览：

| 类别           | 示例函数                              | 主要用途                                                                 |
| :------------- | :------------------------------------ | :----------------------------------------------------------------------- |
| **基本运算**   | `Abs`, `Mod`, `Remainder`, `Dim`      | 处理绝对值、取余、差值等                                                 |
| **幂与开方**   | `Pow`, `Pow10`, `Sqrt`, `Cbrt`        | 计算幂次、平方根、立方根等                                               |
| **三角函数**   | `Sin`, `Cos`, `Tan`, `Asin`, `Acos`, `Atan`, `Atan2` | 计算三角函数和反三角函数，**参数是弧度**                 |
| **对数与指数** | `Exp`, `Log`, `Log10`, `Log2`         | 计算指数、自然对数、常用对数、以2为底的对数                              |
| **取整与舍入** | `Ceil`, `Floor`, `Trunc`, `Round`, `RoundToEven` | 向上、向下取整、截断、四舍五入                                 |
| **极值比较**   | `Max`, `Min`                          | 返回两个值中的最大值或最小值                                             |
| **特殊值处理** | `IsNaN`, `IsInf`, `Inf`, `NaN`        | 判断和生成非数 (NaN)、无穷大 (Inf) 等特殊值                              |

### 📌 基本运算函数

*   **`Abs(x float64) float64`**：返回 `x` 的绝对值。
    ```go
    fmt.Println(math.Abs(-3.14)) // 输出: 3.14
    fmt.Println(math.Abs(2.5))   // 输出: 2.5
    ```
    **注意**：`math.Abs` 需要 `float64` 类型参数。对于整数，需先转换（如 `math.Abs(float64(-5))`），或使用标准库函数（如 `int(math.Abs(float64(-5)))` 转换回整数）。

*   **`Mod(x, y float64) float64`** 与 **`Remainder(x, y float64) float64`**：两者都计算余数，但处理符号的规则不同。
    *   `Mod` 的结果的符号与除数 `y` 相同。
    *   `Remainder` 的结果的符号与被除数 `x` 相同。
    ```go
    fmt.Println(math.Mod(10, 3))        // 输出: 1
    fmt.Println(math.Remainder(10, 3))  // 输出: 1
    fmt.Println(math.Mod(-10, 3))       // 输出: -1 (因为 -10 = -4*3 + 2, 但规则同除数3的符号，所以是 -1? 这里需要验证官方定义)
    fmt.Println(math.Remainder(-10, 3)) // 输出: -1 (规则同被除数-10的符号)
    ```
    对于 `math.Mod(-10, 3)`，准确计算是 `-10 % 3` 的浮点数版本，结果符号与除数 `3` 一致，因此是 `-1` (因为 `-10 = -3*3 -1`，余数确实是 `-1`)。而 `Remainder(-10, 3)` 结果符号与被除数 `-10` 一致，也是 `-1`。在很多情况下，`Remainder` 符合 IEEE 754 标准。

*   **`Dim(x, y float64) float64`**：返回 `x - y` 和 `0` 中的较大值。如果 `x > y`，返回 `x - y`；否则返回 `0`。
    ```go
    fmt.Println(math.Dim(5, 3)) // 输出: 2
    fmt.Println(math.Dim(3, 5)) // 输出: 0
    ```

### ✨ 幂与开方函数

*   **`Pow(x, y float64) float64`**：返回 `x` 的 `y` 次幂。
    ```go
    fmt.Println(math.Pow(2, 3))   // 输出: 8
    fmt.Println(math.Pow(3, 2))   // 输出: 9
    fmt.Println(math.Pow(8, 1/3.0)) // 输出: 2 (计算立方根)
    ```

*   **`Pow10(n int) float64`**：返回 `10` 的 `n` 次幂。比 `math.Pow(10, n)` 更高效。
    ```go
    fmt.Println(math.Pow10(2)) // 输出: 100
    fmt.Println(math.Pow10(-1)) // 输出: 0.1
    ```

*   **`Sqrt(x float64) float64`**：返回 `x` 的平方根。
    ```go
    fmt.Println(math.Sqrt(16)) // 输出: 4
    fmt.Println(math.Sqrt(2))  // 输出: 1.4142135623730951
    ```
    **注意**：对负数取平方根会返回 `NaN`。

*   **`Cbrt(x float64) float64`**：返回 `x` 的立方根。
    ```go
    fmt.Println(math.Cbrt(27)) // 输出: 3
    fmt.Println(math.Cbrt(-8)) // 输出: -2
    ```

### 📐 三角函数

`math` 包中的三角函数（如 `Sin`, `Cos`, `Tan`, `Asin`, `Acos`, `Atan`）都使用**弧度**作为参数单位，而不是角度。

*   **角度与弧度转换**：
    弧度 = 角度 × (π / 180)
    角度 = 弧度 × (180 / π)

*   **`Sin(x float64) float64`**, **`Cos(x float64) float64`**, **`Tan(x float64) float64`**：计算正弦、余弦、正切值。
    ```go
    rad := 45 * math.Pi / 180 // 将45度转换为弧度
    fmt.Println(math.Sin(rad)) // 输出: 0.7071067811865475 (近似√2/2)
    fmt.Println(math.Cos(rad)) // 输出: 0.7071067811865476 (近似√2/2)
    fmt.Println(math.Tan(rad)) // 输出: 0.9999999999999999 (近似1)
    ```

*   **`Asin(x float64) float64`**, **`Acos(x float64) float64`**, **`Atan(x float64) float64`**：计算反正弦、反余弦、反正切值（返回弧度）。
    ```go
    fmt.Println(math.Asin(1) * 180 / math.Pi) // 输出: 90 (弧度转角度)
    ```

*   **`Atan2(y, x float64) float64`**：返回 `y/x` 的反正切值，**能正确处理象限问题**（根据 `x` 和 `y` 的符号确定象限），通常比 `Atan(y/x)` 更推荐使用。
    ```go
    // 计算点 (1, 1) 的极角 (45度)
    angle := math.Atan2(1, 1) * 180 / math.Pi
    fmt.Println(angle) // 输出: 45
    ```

### 📊 对数与指数函数

*   **`Exp(x float64) float64`**：返回 `e` (欧拉数，约 2.71828) 的 `x` 次幂。
    ```go
    fmt.Println(math.Exp(1)) // 输出: 2.718281828459045 (近似e)
    fmt.Println(math.Exp(2)) // 输出: 7.38905609893065 (近似e²)
    ```

*   **`Log(x float64) float64`**：返回 `x` 的自然对数（以 `e` 为底）。
    ```go
    fmt.Println(math.Log(math.E)) // 输出: 1
    fmt.Println(math.Log(10))     // 输出: 2.302585092994046
    ```

*   **`Log10(x float64) float64`**：返回 `x` 的常用对数（以 10 为底）。
    ```go
    fmt.Println(math.Log10(1000)) // 输出: 3
    ```

*   **`Log2(x float64) float64`**：返回 `x` 的以 2 为底的对数。
    ```go
    fmt.Println(math.Log2(8)) // 输出: 3
    ```

### 🔢 取整与舍入函数

*   **`Ceil(x float64) float64`**：**向上取整**，返回大于等于 `x` 的最小整数。
    ```go
    fmt.Println(math.Ceil(3.2))  // 输出: 4
    fmt.Println(math.Ceil(-3.2)) // 输出: -3
    ```

*   **`Floor(x float64) float64`**：**向下取整**，返回小于等于 `x` 的最大整数。
    ```go
    fmt.Println(math.Floor(3.9))  // 输出: 3
    fmt.Println(math.Floor(-3.9)) // 输出: -4
    ```

*   **`Trunc(x float64) float64`**：**截断取整**，直接舍弃小数部分，返回整数部分。
    ```go
    fmt.Println(math.Trunc(3.9))  // 输出: 3
    fmt.Println(math.Trunc(-3.9)) // 输出: -3
    ```

*   **`Round(x float64) float64`**：**四舍五入**到最接近的整数。
    ```go
    fmt.Println(math.Round(3.4)) // 输出: 3
    fmt.Println(math.Round(3.5)) // 输出: 4
    fmt.Println(math.Round(-2.5)) // 输出: -3? (Go 的 Round 遵循 IEEE 754 标准，可能向远离零的方向舍入)
    ```
    **注意**：`math.Round` 的舍入规则是“四舍六入五成双”吗？在 Go 中，`math.Round(2.5)` 会得到 `3`，而 `math.RoundToEven(2.5)` 会得到 `2`。

*   **`RoundToEven(x float64) float64`**：四舍五入到最接近的整数，但当小数部分恰好为 `0.5` 时，会舍入到最接近的**偶数**。
    ```go
    fmt.Println(math.RoundToEven(2.5)) // 输出: 2
    fmt.Println(math.RoundToEven(3.5)) // 输出: 4
    ```

### ⚖️ 极值比较函数

*   **`Max(x, y float64) float64`**：返回 `x` 和 `y` 中的较大值。
    ```go
    fmt.Println(math.Max(3, 5))   // 输出: 5
    fmt.Println(math.Max(-2, -5)) // 输出: -2
    ```

*   **`Min(x, y float64) float64`**：返回 `x` 和 `y` 中的较小值。
    ```go
    fmt.Println(math.Min(3, 5))   // 输出: 3
    fmt.Println(math.Min(-2, -5)) // 输出: -5
    ```

### 🌀 特殊值处理函数

在处理浮点数时，可能会遇到非数（NaN）或无穷大（Inf）等特殊值。

*   **`IsNaN(x float64) bool`**：判断 `x` 是否为非数（NaN）。**不能直接用 `==` 判断 NaN**。
    ```go
    val := math.Sqrt(-1)        // 生成 NaN
    fmt.Println(math.IsNaN(val)) // 输出: true
    fmt.Println(val == math.NaN()) // 输出: false! 不要这样比较NaN
    ```

*   **`IsInf(x float64, sign int) bool`**：判断 `x` 是否为无穷大。`sign > 0` 判断正无穷，`sign < 0` 判断负无穷，`sign == 0` 判断任何无穷大。
    ```go
    posInf := math.Inf(1)
    negInf := math.Inf(-1)
    fmt.Println(math.IsInf(posInf, 1))  // 输出: true
    fmt.Println(math.IsInf(negInf, -1)) // 输出: true
    fmt.Println(math.IsInf(posInf, 0))  // 输出: true (任何无穷大)
    ```

*   **`Inf(sign int) float64`**：根据 `sign` 生成无穷大。`sign >= 0` 返回正无穷，`sign < 0` 返回负无穷。
    ```go
    posInf := math.Inf(1)
    negInf := math.Inf(-1)
    ```

*   **`NaN() float64`**：返回一个 NaN 值。
    ```go
    nan := math.NaN()
    ```

### 🧮 数学常量

`math` 包定义了一些常用的数学常量（`float64` 类型）：

*   **`Pi`**：圆周率 π，约等于 3.141592653589793
*   **`E`**：自然对数的底 e，约等于 2.718281828459045
*   **`Phi`**：黄金分割比，约等于 1.618033988749895
```go
fmt.Println(math.Pi)  // 输出: 3.141592653589793
fmt.Println(math.E)   // 输出: 2.718281828459045
fmt.Println(math.Phi) // 输出: 1.618033988749895
```

### 💡 实用技巧与注意事项

1.  **类型转换**：`math` 包中的函数大多操作 `float64` 类型。若需处理其他数值类型（如 `int`, `float32`），常需显式转换。
    ```go
    intValue := 10
    floatValue := math.Sqrt(float64(intValue)) // 必须转换
    ```

2.  **弧度与角度**：三角函数使用弧度制。使用角度时，务必转换 `radians = degrees * (Pi / 180)`。

3.  **特殊值（NaN, Inf）**：
    *   避免直接比较 `NaN`，始终使用 `math.IsNaN()`。
    *   检查除数是否为零，避免产生 `Inf` 或 `NaN`。
    *   对负数开偶次方根、对负数取对数等操作会产生 `NaN`。

4.  **浮点数精度**：浮点数计算存在精度误差，比较时不宜直接使用 `==`，应检查两数差值是否小于某个极小值（epsilon）。
    ```go
    // 不推荐: if a == b { ... }
    // 推荐:
    epsilon := 1e-10
    if math.Abs(a - b) < epsilon {
        // 认为 a 和 b 相等
    }
    ```

5.  **随机数生成**：`math` 包本身**不提供**随机数函数。随机数生成在 `math/rand` 包（伪随机数）和 `crypto/rand` 包（密码学安全随机数）中。

6.  **高精度计算**：对于需要高精度的财务计算或科学计算，可以考虑使用 `math/big` 包提供的 `Int`, `Float`, `Rat` 类型。

`math` 包是 Go 语言进行数学运算的基础工具包，熟练掌握其常用函数和注意事项，能让你在数据处理和科学计算中更加得心应手。