Go 语言的 `fmt` 包是处理格式化输入输出的核心标准库，它功能丰富且实用。下面我会为你梳理它的主要功能、常用函数以及一些实用技巧。

### 📌 核心输出函数

`fmt` 包提供了多组输出函数，适用于不同的场景。

| 函数系列         | 核心函数      | 主要功能                                                                 | 示例                                                                                                | 适用场景                  |
| :--------------- | :------------ | :----------------------------------------------------------------------- | :-------------------------------------------------------------------------------------------------- | :------------------------ |
| **标准输出**     | `Print`       | 输出内容，不换行                                                           | `fmt.Print("Hello, ", "World!")` 输出 `Hello, World!`                                               | 控制台简单输出              |
|                  | `Println`     | 输出内容**并换行**，多个参数间自动加空格                                         | `fmt.Println("Hello", "World!")` 输出 `Hello World!\n`                                              | 控制台输出，自动格式化        |
|                  | `Printf`      | **格式化输出**，需指定格式动词                                                 | `fmt.Printf("Name: %s, Age: %d", "Bob", 25)` 输出 `Name: Bob, Age: 25`                              | 需控制格式的复杂输出          |
| **字符串格式化** | `Sprint`      | 将内容格式化为字符串**并返回**，不换行                                          | `s := fmt.Sprint("The answer is ", 42)` `s` 为 `"The answer is 42"`                                 | 拼接字符串                  |
|                  | `Sprintf`     | **格式化**内容为字符串并返回                                                   | `s := fmt.Sprintf("Value: %d", 100)` `s` 为 `"Value: 100"`                                          | 生成格式化字符串            |
|                  | `Sprintln`    | 将内容格式化为字符串并返回，**末尾添加换行符**                                     | `s := fmt.Sprintln("Hello")` `s` 为 `"Hello\n"`                                                     |                           |
| **定向输出**     | `Fprint`      | 将内容输出到实现了 `io.Writer` 接口的对象（如文件、网络连接）                          | `file, _ := os.Create("text.txt`) `fmt.Fprint(file, "Hello ")` 向文件写入 `Hello `                  | 写入文件或网络流            |
|                  | `Fprintf`     | **格式化**输出到 `io.Writer` 对象                                            | `fmt.Fprintf(file, "Name: %s", "Alice")`                                                            | 格式化写入文件或网络流        |
|                  | `Fprintln`    | 输出内容到 `io.Writer` 对象**并添加换行符**                                      | `fmt.Fprintln(file, "World!")` 向文件写入 `World!\n`                                                |                           |

### 📥 输入扫描函数

`fmt` 包同样提供了从不同输入源读取数据的函数。

| 函数系列        | 核心函数       | 主要输入源   | 扫描规则                                      | 示例                                                                                                                             | 返回值                                       |
| :-------------- | :------------- | :----------- | :-------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------- |
| **标准输入**    | `Scan`         | 标准输入     | 读取**空格分隔**的值                            | `var a int; var b string; fmt.Scan(&a, &b)` 输入 `10 ABC` 后，`a=10`, `b="ABC"`                                                  | `(n int, err error)` 成功扫描的项数和错误 |
|                 | `Scanf`        | 标准输入     | 按**格式字符串**读取                             | `var a int; fmt.Scanf("Number:%d", &a)` 输入 `Number:42` 后，`a=42`                                                              | `(n int, err error)` 成功扫描的项数和错误 |
|                 | `Scanln`       | 标准输入     | 读取到**换行符**停止                            | `var a int; var b string; fmt.Scanln(&a, &b)` 输入 `10 ABC` 后，`a=10`, `b="ABC"`                                                | `(n int, err error)` 成功扫描的项数和错误 |
| **字符串扫描**  | `Sscan`        | 字符串       | 从字符串中读取**空格分隔**的值                     | `str := "100 Golang"; var x int; var y string; fmt.Sscan(str, &x, &y)` `x=100`, `y="Golang"`                                     | `(n int, err error)` 成功扫描的项数和错误 |
|                 | `Sscanf`       | 字符串       | 从字符串中按**格式字符串**读取                      | `str := "ID:12345"; var id int; fmt.Sscanf(str, "ID:%d", &id)` `id=12345`                                                        | `(n int, err error)` 成功扫描的项数和错误 |
|                 | `Sscanln`      | 字符串       | 从字符串中读取直到**换行符**                       | `str := "100 Golang\n"; var x int; var y string; fmt.Sscanln(str, &x, &y)` `x=100`, `y="Golang"`                                 | `(n int, err error)` 成功扫描的项数和错误 |
| **流输入**      | `Fscan`        | `io.Reader` | 从 `io.Reader`（如文件）读取**空格分隔**的值        | `file, _ := os.Open("data.txt"); var x int; fmt.Fscan(file, &x)` 从文件读取                                                                 | `(n int, err error)` 成功扫描的项数和错误 |
|                 | `Fscanf`       | `io.Reader` | 从 `io.Reader` 按**格式字符串**读取                | `file, _ := os.Open("data.txt"); var x int; fmt.Fscanf(file, "Value:%d", &x)` 从文件读取                                            | `(n int, err error)` 成功扫描的项数和错误 |
|                 | `Fscanln`      | `io.Reader` | 从 `io.Reader` 读取直到**换行符**                  | `file, _ := os.Open("data.txt"); var x int; fmt.Fscanln(file, &x)` 从文件读取                                                       | `(n int, err error)` 成功扫描的项数和错误 |

💡 **处理输入含空格的情况**：`Scan` 系列函数遇到空格或换行会停止读取。若要读取整行（包括空格），可用 `bufio` 包：
```go
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n') // 读取直到换行符
input = strings.TrimSpace(input)    // 去除末尾换行符和空格
```

### 🔤 常用格式化动词（Verbs）

`Printf`、`Sprintf` 和 `Fprintf` 使用格式化动词来指定数据的显示格式。

| 类别       | 动词 | 说明                                               | 示例                                          | 输出示例                          |
| :--------- | :--- | :------------------------------------------------- | :-------------------------------------------- | :-------------------------------- |
| **通用**   | `%v` | 值的默认格式表示                                     | `fmt.Printf("%v", site)`                      | `{studygolang}`                   |
|            | `%+v` | 输出结构体时包含字段名                                | `fmt.Printf("%+v", site)`                     | `{Name:studygolang}`              |
|            | `%#v` | 值的 Go 语法表示                                    | `fmt.Printf("%#v", site)`                     | `main.Website{Name:"studygolang"}`|
|            | `%T` | 值的类型的 Go 语法表示                               | `fmt.Printf("%T", site)`                      | `main.Website`                    |
|            | `%%` | 百分号字面量                                         | `fmt.Printf("%%")`                            | `%`                               |
| **整数**   | `%d` | 十进制表示                                         | `fmt.Printf("%d", 0x12)`                      | `18`                              |
|            | `%b` | 二进制表示                                         | `fmt.Printf("%b", 5)`                         | `101`                             |
|            | `%x` | 十六进制表示（小写 a-f）                              | `fmt.Printf("%x", 13)`                        | `d`                               |
|            | `%X` | 十六进制表示（大写 A-F）                              | `fmt.Printf("%X", 13)`                        | `D`                               |
|            | `%o` | 八进制表示                                         | `fmt.Printf("%o", 10)`                        | `12`                              |
|            | `%c` | 对应 Unicode 码点的字符                              | `fmt.Printf("%c", 0x4E2D)`                    | `中`                              |
|            | `%U` | Unicode 格式（U+1234）                              | `fmt.Printf("%U", 0x4E2D)`                    | `U+4E2D`                          |
| **浮点数** | `%f` | 有小数部分，无指数（默认精度6）                           | `fmt.Printf("%f", 10.2)`                      | `10.200000`                       |
|            | `%e` | 科学计数法（e，如 1.2e+10）                           | `fmt.Printf("%e", 10.2)`                      | `1.020000e+01`                    |
|            | `%E` | 科学计数法（E，如 1.2E+10）                           | `fmt.Printf("%E", 10.2)`                      | `1.020000E+01`                    |
|            | `%g` | 根据情况选择 `%e` 或 `%f`（更紧凑，无末尾0）               | `fmt.Printf("%g", 10.20)`                     | `10.2`                            |
|            | `%G` | 根据情况选择 `%E` 或 `%f`（更紧凑，无末尾0）               | `fmt.Printf("%G", 10.20+2i)`                  | `(10.2+2i)`                       |
| **布尔值** | `%t` | 输出 `true` 或 `false`                             | `fmt.Printf("%t", true)`                      | `true`                            |
| **字符串和字节切片** | `%s` | 输出字符串或 `[]byte`                                | `fmt.Printf("%s", []byte("Go"))`              | `Go`                              |
|            | `%q` | 带双引号的字符串（Go 语法安全转义）                       | `fmt.Printf("%q", "Go\t")`                    | `"Go\\t"`                         |
|            | `%x` | 十六进制表示（小写，每字节两字符）                         | `fmt.Printf("%x", "golang")`                  | `676f6c616e67`                    |
|            | `%X` | 十六进制表示（大写，每字节两字符）                         | `fmt.Printf("%X", "golang")`                  | `676F6C616E67`                    |
| **指针**   | `%p` | 指针的十六进制表示（带前缀 0x）                          | `fmt.Printf("%p", &x)`                        | `0xc0000160d8`                    |

### ⚙️ 宽度、精度与对齐

可以在动词中指定宽度、精度和对齐方式以实现精细控制。

*   **宽度与精度**：通过在 `%` 后加上数字设置宽度，加上 `.` 和数字设置精度。
    *   `%5d`：最小宽度5（右对齐）
    *   `%-5d`：最小宽度5（**左对齐**）
    *   `%05d`：最小宽度5，用**0填充**
    *   `%.3f`：浮点数保留3位小数
    *   `%9.2f`：宽度9，精度2（右对齐）
    *   `%-9.2f`：宽度9，精度2（**左对齐**）

*   **其他标志（Flags）**：
    *   `'+'`：总是输出数值的正负号。
    *   `' '`（空格）：对正数，在输出前加一个空格。
    *   `'#'`：备用格式，例如为八进制加 `0`，为十六进制加 `0x`，为指针去掉 `0x` 前缀。

示例：
```go
fmt.Printf("|%6s|%-6s|%06d|\n", "Hi", "Go", 42) // 输出：|    Hi|Go    |000042|
fmt.Printf("%+d, %+d, % d\n", 10, -10, 10)      // 输出：+10, -10,  10
fmt.Printf("%#x, %#o, %#p\n", 127, 127, &x)     // 输出：0x7f, 0177, c0000160d8
```

### 🛠️ 错误处理与自定义格式化

*   **错误处理**：`Print`、`Scan` 等系列函数基本都会返回 `(n int, err error)`，表示已处理的字节数和可能遇到的错误。良好的实践是检查错误。
    ```go
    n, err := fmt.Printf("%d", "not-a-number")
    if err != nil {
        log.Fatal(err) // 例如：%!d(string=not-a-number)
    }
    ```

*   **自定义类型的格式化**：
    *   实现 `Stringer` 接口（`String() string` 方法）来自定义 `%v` 等动词的输出。
        ```go
        type User struct { Name string; Age int }
        func (u User) String() string {
            return fmt.Sprintf("%s (%d years old)", u.Name, u.Age)
        }
        u := User{"Alice", 30}
        fmt.Println(u) // 输出：Alice (30 years old)
        ```
    *   实现 `Formatter` 接口（`Format(f State, verb rune)` 方法）可以更精细地控制各种动词下的格式化行为。

*   **生成错误**：使用 `Errorf` 生成包含格式化信息的错误。
    ```go
    if id < 0 {
        return fmt.Errorf("invalid id: %d", id) // 返回一个错误
    }
    ```

### 💡 实用技巧与注意事项

1.  **性能考虑**：频繁进行字符串拼接时，`fmt.Sprintf` 可能不是最高效的选择，对于简单拼接，考虑使用 `strings.Builder` 或 `+` 操作符。
2.  **输入中的空格与换行**：使用 `Scan` 系列函数要注意它们对空格和换行符的处理规则。需要读取整行时，`bufio.Reader` 的 `ReadString` 或 `ReadBytes` 方法更合适。
3.  **动词的灵活性**：`%v` 是一个很有用的通用动词，尤其在调试不知道具体类型时。但明确类型时，使用特定动词（如 `%d`, `%s`）通常更清晰。
4.  **宽度与截断**：为字符串指定精度（例如 `%.5s`）可以限制最大输出字符数，超出部分会被截断。
5.  **检查扫描结果**：使用 `Scan` 系列函数时，务必检查返回的 `n`（成功扫描的项数）和 `err`，以防止输入与预期不符导致程序错误。

`fmt` 包是 Go 程序员最亲密的伙伴之一，花时间熟悉它会让你的编程体验更加顺畅。