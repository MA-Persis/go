# Go的net/http包详细介绍

Go语言的`net/http`包提供了构建HTTP客户端和服务器的完整功能，是Go语言网络编程的核心包之一。下面通过表格概览主要组件：

| 组件类型 | 核心结构/函数 | 主要作用 |
|---------|-------------|---------|
| **服务端** | `http.Server` | HTTP服务器主结构 |
| | `http.Handler` | 请求处理接口 |
| | `http.ServeMux` | 路由复用器 |
| | `http.ResponseWriter` | 响应写入接口 |
| | `http.Request` | 请求信息结构 |
| **客户端** | `http.Client` | HTTP客户端 |
| | `http.Request` | 客户端请求 |
| | `http.Response` | 服务器响应 |

## 🏗️ 服务端核心组件

### 1. 创建基础HTTP服务器

```go
package main

import (
    "fmt"
    "net/http"
)

// 方法1: 使用HandlerFunc
func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// 方法2: 实现Handler接口
type welcomeHandler struct{}

func (h *welcomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to Go HTTP Server!")
}

func main() {
    // 使用默认多路复用器
    http.HandleFunc("/hello", helloHandler)
    http.Handle("/welcome", &welcomeHandler{})
    
    // 自定义路由
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Root path")
    })
    
    // 启动服务器
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", mux)
}
```

### 2. 中间件(Middleware)模式

```go
// 日志中间件
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("%s %s\n", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// 认证中间件
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token != "secret-token" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Protected data")
    })
    
    // 应用中间件链
    handler := loggingMiddleware(authMiddleware(mux))
    
    http.ListenAndServe(":8080", handler)
}
```

### 3. 高级服务器配置

```go
func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
    
    server := &http.Server{
        Addr:         ":8080",
        Handler:      mux,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }
    
    // 优雅关闭
    go func() {
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("Server error: %v\n", err)
        }
    }()
    
    // 处理中断信号
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, os.Interrupt)
    <-quit
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := server.Shutdown(ctx); err != nil {
        fmt.Printf("Server shutdown error: %v\n", err)
    }
}
```

## 🌐 客户端核心功能

### 1. 基础HTTP请求

```go
package main

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strings"
    "time"
)

func main() {
    // GET请求
    resp, err := http.Get("https://httpbin.org/get")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("GET Response: %s\n", string(body))
    
    // POST请求 - JSON数据
    jsonData := `{"name":"John","age":30}`
    resp, err = http.Post("https://httpbin.org/post", "application/json", 
        strings.NewReader(jsonData))
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    // POST请求 - 表单数据
    formData := url.Values{}
    formData.Add("username", "john")
    formData.Add("password", "secret")
    
    resp, err = http.PostForm("https://httpbin.org/post", formData)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
}
```

### 2. 自定义客户端和请求

```go
func advancedClient() {
    // 创建自定义客户端
    client := &http.Client{
        Timeout: 10 * time.Second,
        Transport: &http.Transport{
            MaxIdleConns:        100,
            IdleConnTimeout:     90 * time.Second,
            TLSHandshakeTimeout: 10 * time.Second,
        },
    }
    
    // 创建自定义请求
    req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
    if err != nil {
        panic(err)
    }
    
    // 设置请求头
    req.Header.Set("User-Agent", "MyGoClient/1.0")
    req.Header.Set("Authorization", "Bearer token123")
    
    // 执行请求
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    fmt.Printf("Status: %s\n", resp.Status)
    for k, v := range resp.Header {
        fmt.Printf("%s: %s\n", k, v)
    }
}
```

### 3. 处理JSON响应

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func jsonRequests() {
    // 发送JSON请求
    user := User{ID: 1, Name: "John Doe", Email: "john@example.com"}
    jsonData, _ := json.Marshal(user)
    
    req, err := http.NewRequest("POST", "https://httpbin.org/post", 
        bytes.NewBuffer(jsonData))
    if err != nil {
        panic(err)
    }
    req.Header.Set("Content-Type", "application/json")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    // 解析JSON响应
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)
    fmt.Printf("Response: %+v\n", result)
}
```

## 🔧 实用功能详解

### 1. 文件服务器

```go
func fileServer() {
    // 静态文件服务
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // 单文件服务
    http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./files/document.pdf")
    })
    
    http.ListenAndServe(":8080", nil)
}
```

### 2. 处理不同Content-Type

```go
func contentHandlers() {
    http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{"message": "Hello JSON"})
    })
    
    http.HandleFunc("/html", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, "<h1>Hello HTML</h1>")
    })
    
    http.HandleFunc("/xml", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/xml")
        fmt.Fprintf(w, "<message>Hello XML</message>")
    })
}
```

### 3. 请求参数处理

```go
func parameterHandlers() {
    http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
        // 查询参数
        name := r.URL.Query().Get("name")
        age := r.URL.Query().Get("age")
        fmt.Fprintf(w, "Name: %s, Age: %s", name, age)
    })
    
    http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
        // 表单参数
        if err := r.ParseForm(); err != nil {
            http.Error(w, "Bad Request", http.StatusBadRequest)
            return
        }
        username := r.FormValue("username")
        password := r.FormValue("password")
        fmt.Fprintf(w, "Username: %s, Password: %s", username, password)
    })
}
```

## ⚡ 性能优化技巧

### 1. 连接池配置

```go
func createOptimizedClient() *http.Client {
    return &http.Client{
        Transport: &http.Transport{
            MaxIdleConns:        100,
            MaxIdleConnsPerHost: 10,
            IdleConnTimeout:     90 * time.Second,
            DisableCompression:  false, // 启用gzip压缩
        },
        Timeout: 30 * time.Second,
    }
}
```

### 2. 响应处理最佳实践

```go
func efficientResponseHandling() {
    resp, err := http.Get("https://example.com/large-file")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    
    // 流式处理大文件
    file, err := os.Create("downloaded.file")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    // 使用io.Copy避免内存爆炸
    _, err = io.Copy(file, resp.Body)
    if err != nil {
        panic(err)
    }
}
```

## 🛡️ 错误处理与调试

```go
func robustHTTPCall() {
    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    req, err := http.NewRequest("GET", "https://httpbin.org/delay/5", nil)
    if err != nil {
        fmt.Printf("Request creation error: %v\n", err)
        return
    }
    
    resp, err := client.Do(req)
    if err != nil {
        // 处理各种网络错误
        if urlErr, ok := err.(*url.Error); ok {
            if urlErr.Timeout() {
                fmt.Println("Request timeout")
            } else {
                fmt.Printf("URL error: %v\n", urlErr)
            }
        } else {
            fmt.Printf("Other error: %v\n", err)
        }
        return
    }
    defer resp.Body.Close()
    
    // 检查HTTP状态码
    if resp.StatusCode != http.StatusOK {
        fmt.Printf("HTTP error: %s\n", resp.Status)
        return
    }
    
    // 成功处理响应
    body, _ := io.ReadAll(resp.Body)
    fmt.Printf("Success: %s\n", string(body[:100]))
}
```

## 📊 实际应用场景

### 1. RESTful API服务器

```go
package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    
    "github.com/gorilla/mux"
)

type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

var products = []Product{
    {1, "Laptop", 999.99},
    {2, "Mouse", 29.99},
}

func main() {
    r := mux.NewRouter()
    
    r.HandleFunc("/products", getProducts).Methods("GET")
    r.HandleFunc("/products/{id}", getProduct).Methods("GET")
    r.HandleFunc("/products", createProduct).Methods("POST")
    
    http.ListenAndServe(":8080", r)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    
    for _, p := range products {
        if p.ID == id {
            json.NewEncoder(w).Encode(p)
            return
        }
    }
    http.Error(w, "Product not found", http.StatusNotFound)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
    var product Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    products = append(products, product)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}
```

## 总结

Go的`net/http`包提供了强大而灵活的HTTP编程能力：

- **服务端**：通过`Handler`接口、中间件模式和路由管理，可以构建高性能的Web服务器
- **客户端**：支持各种HTTP方法、请求定制和连接池优化
- **生产就绪**：包含超时控制、优雅关闭、错误处理等企业级特性
- **扩展性强**：可以轻松集成第三方路由库、模板引擎等

这个包的设计体现了Go语言的简洁哲学，既提供了基础构建块，又保持了高度的可扩展性，是构建现代Web应用的理想选择。