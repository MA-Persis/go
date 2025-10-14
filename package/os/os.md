# Go 语言 `os` 包详解

Go 语言的 `os` 包提供了与操作系统交互的各种功能，包括文件操作、进程管理、环境变量访问等。它是 Go 标准库中用于系统交互的核心包之一。

## 📋 os 包核心功能概览

| 功能类别 | 主要函数/类型 | 描述 |
| :--- | :--- | :--- |
| **文件操作** | `Open`, `Create`, `Read`, `Write`, `Close` | 文件的打开、创建、读写和关闭 |
| | `Stat`, `Lstat` | 获取文件信息 |
| | `Rename`, `Remove`, `Mkdir`, `MkdirAll` | 文件/目录的重命名、删除和创建 |
| | `Chmod`, `Chown` | 修改文件权限和所有者 |
| **目录操作** | `ReadDir` | 读取目录内容 |
| | `Getwd`, `Chdir` | 获取和改变当前工作目录 |
| **进程管理** | `Getpid`, `Getppid` | 获取进程ID |
| | `Exit` | 程序退出 |
| | `Args` | 获取命令行参数 |
| **环境变量** | `Getenv`, `Setenv`, `Unsetenv` | 环境变量的获取、设置和删除 |
| | `Environ` | 获取所有环境变量 |
| **用户信息** | `Getuid`, `Getgid`, `Geteuid` | 获取用户和组ID |
| | `Getgroups` | 获取用户所属组 |
| **系统信息** | `Hostname` | 获取主机名 |
| | `Getpagesize` | 获取系统内存页大小 |
| **文件描述符** | `Stdout`, `Stdin`, `Stderr` | 标准输入、输出和错误 |
| | `File` 类型 | 文件操作的核心类型 |

## 📁 文件操作

### 1. 打开和创建文件

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 打开文件（只读）
    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println("打开文件失败:", err)
        return
    }
    defer file.Close() // 确保文件关闭
    
    // 创建文件（如果存在则截断）
    newFile, err := os.Create("newfile.txt")
    if err != nil {
        fmt.Println("创建文件失败:", err)
        return
    }
    defer newFile.Close()
    
    fmt.Println("文件操作成功")
}
```

### 2. 读取文件内容

```go
func readFile() {
    // 一次性读取整个文件（小文件适用）
    data, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("读取文件失败:", err)
        return
    }
    fmt.Println("文件内容:", string(data))
    
    // 逐行或分块读取（大文件适用）
    file, err := os.Open("largefile.txt")
    if err != nil {
        fmt.Println("打开文件失败:", err)
        return
    }
    defer file.Close()
    
    buffer := make([]byte, 1024) // 1KB缓冲区
    for {
        n, err := file.Read(buffer)
        if err != nil && err != io.EOF {
            fmt.Println("读取错误:", err)
            break
        }
        if n == 0 {
            break
        }
        fmt.Print(string(buffer[:n]))
    }
}
```

### 3. 写入文件内容

```go
func writeFile() {
    // 一次性写入整个文件
    content := []byte("Hello, World!\n")
    err := os.WriteFile("output.txt", content, 0644)
    if err != nil {
        fmt.Println("写入文件失败:", err)
        return
    }
    
    // 追加写入
    file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("打开文件失败:", err)
        return
    }
    defer file.Close()
    
    _, err = file.WriteString("追加的内容\n")
    if err != nil {
        fmt.Println("写入失败:", err)
    }
}
```

### 4. 文件信息和操作

```go
func fileOperations() {
    // 获取文件信息
    fileInfo, err := os.Stat("test.txt")
    if err != nil {
        fmt.Println("获取文件信息失败:", err)
        return
    }
    
    fmt.Println("文件名:", fileInfo.Name())
    fmt.Println("文件大小:", fileInfo.Size(), "字节")
    fmt.Println("修改时间:", fileInfo.ModTime())
    fmt.Println("权限:", fileInfo.Mode())
    
    // 重命名文件
    err = os.Rename("oldname.txt", "newname.txt")
    if err != nil {
        fmt.Println("重命名失败:", err)
    }
    
    // 删除文件
    err = os.Remove("file_to_delete.txt")
    if err != nil {
        fmt.Println("删除失败:", err)
    }
    
    // 修改文件权限
    err = os.Chmod("test.txt", 0755)
    if err != nil {
        fmt.Println("修改权限失败:", err)
    }
}
```

## 📂 目录操作

### 1. 创建和删除目录

```go
func directoryOperations() {
    // 创建单个目录
    err := os.Mkdir("mydir", 0755)
    if err != nil {
        fmt.Println("创建目录失败:", err)
    }
    
    // 创建多级目录
    err = os.MkdirAll("parent/child/grandchild", 0755)
    if err != nil {
        fmt.Println("创建多级目录失败:", err)
    }
    
    // 删除目录（必须是空目录）
    err = os.Remove("mydir")
    if err != nil {
        fmt.Println("删除目录失败:", err)
    }
    
    // 删除目录及其所有内容
    err = os.RemoveAll("parent")
    if err != nil {
        fmt.Println("删除目录树失败:", err)
    }
}
```

### 2. 读取目录内容

```go
func readDirectory() {
    // 读取目录内容
    entries, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("读取目录失败:", err)
        return
    }
    
    fmt.Println("当前目录内容:")
    for _, entry := range entries {
        info, err := entry.Info()
        if err != nil {
            continue
        }
        
        if entry.IsDir() {
            fmt.Printf("[目录] %s (%d bytes)\n", entry.Name(), info.Size())
        } else {
            fmt.Printf("[文件] %s (%d bytes)\n", entry.Name(), info.Size())
        }
    }
}
```

### 3. 工作目录操作

```go
func workingDirectory() {
    // 获取当前工作目录
    wd, err := os.Getwd()
    if err != nil {
        fmt.Println("获取工作目录失败:", err)
        return
    }
    fmt.Println("当前工作目录:", wd)
    
    // 改变工作目录
    err = os.Chdir("/tmp")
    if err != nil {
        fmt.Println("改变目录失败:", err)
        return
    }
    
    // 验证目录已改变
    wd, err = os.Getwd()
    if err != nil {
        fmt.Println("获取工作目录失败:", err)
        return
    }
    fmt.Println("新工作目录:", wd)
}
```

## ⚙️ 进程和环境操作

### 1. 进程信息和控制

```go
func processOperations() {
    // 获取进程ID
    pid := os.Getpid()
    ppid := os.Getppid()
    fmt.Printf("进程ID: %d, 父进程ID: %d\n", pid, ppid)
    
    // 获取命令行参数
    args := os.Args
    fmt.Println("命令行参数:", args)
    
    if len(args) > 1 {
        fmt.Println("第一个参数:", args[1])
    }
    
    // 程序退出
    // os.Exit(0) // 正常退出
    // os.Exit(1) // 异常退出
}
```

### 2. 环境变量操作

```go
func environmentOperations() {
    // 获取环境变量
    goPath := os.Getenv("GOPATH")
    fmt.Println("GOPATH:", goPath)
    
    // 设置环境变量
    err := os.Setenv("MY_VAR", "my_value")
    if err != nil {
        fmt.Println("设置环境变量失败:", err)
    }
    
    // 检查环境变量
    value := os.Getenv("MY_VAR")
    fmt.Println("MY_VAR:", value)
    
    // 获取所有环境变量
    envVars := os.Environ()
    fmt.Println("环境变量:")
    for _, env := range envVars {
        fmt.Println(env)
    }
    
    // 删除环境变量
    err = os.Unsetenv("MY_VAR")
    if err != nil {
        fmt.Println("删除环境变量失败:", err)
    }
}
```

### 3. 用户和组信息

```go
func userOperations() {
    // 获取用户ID和组ID
    uid := os.Getuid()
    gid := os.Getgid()
    euid := os.Geteuid()
    fmt.Printf("用户ID: %d, 组ID: %d, 有效用户ID: %d\n", uid, gid, euid)
    
    // 获取用户所属的所有组
    groups, err := os.Getgroups()
    if err != nil {
        fmt.Println("获取组信息失败:", err)
        return
    }
    fmt.Println("所属组:", groups)
}
```

## 🌐 系统信息

```go
func systemInfo() {
    // 获取主机名
    hostname, err := os.Hostname()
    if err != nil {
        fmt.Println("获取主机名失败:", err)
        return
    }
    fmt.Println("主机名:", hostname)
    
    // 获取系统页大小
    pageSize := os.Getpagesize()
    fmt.Println("系统页大小:", pageSize, "字节")
    
    // 获取临时目录
    tempDir := os.TempDir()
    fmt.Println("临时目录:", tempDir)
}
```

## 🔧 标准输入、输出和错误

```go
func stdIO() {
    // 写入标准输出
    fmt.Fprintln(os.Stdout, "这是标准输出")
    
    // 写入标准错误
    fmt.Fprintln(os.Stderr, "这是标准错误")
    
    // 从标准输入读取
    fmt.Print("请输入内容: ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        input := scanner.Text()
        fmt.Println("你输入了:", input)
    }
}
```

## 💡 高级文件操作

### 1. 文件锁

```go
func fileLocking() {
    file, err := os.Create("lockfile.txt")
    if err != nil {
        fmt.Println("创建文件失败:", err)
        return
    }
    defer file.Close()
    
    // 尝试获取排他锁
    err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
    if err != nil {
        fmt.Println("获取文件锁失败:", err)
        return
    }
    defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN) // 释放锁
    
    fmt.Println("已获取文件锁，执行操作...")
    time.Sleep(5 * time.Second)
    fmt.Println("操作完成，释放文件锁")
}
```

### 2. 文件权限和所有权

```go
func filePermissions() {
    // 创建文件
    err := os.WriteFile("permission_test.txt", []byte("test"), 0644)
    if err != nil {
        fmt.Println("创建文件失败:", err)
        return
    }
    
    // 更改文件权限
    err = os.Chmod("permission_test.txt", 0600) // 只有所有者可读写
    if err != nil {
        fmt.Println("更改权限失败:", err)
    }
    
    // 更改文件所有者（需要相应权限）
    // err = os.Chown("permission_test.txt", 1000, 1000) // UID, GID
    // if err != nil {
    //     fmt.Println("更改所有者失败:", err)
    // }
}
```

## ⚠️ 错误处理和最佳实践

### 1. 错误处理模式

```go
func properErrorHandling() {
    // 使用 os.IsNotExist 检查文件不存在错误
    _, err := os.Stat("nonexistent.txt")
    if os.IsNotExist(err) {
        fmt.Println("文件不存在")
    } else if err != nil {
        fmt.Println("其他错误:", err)
    } else {
        fmt.Println("文件存在")
    }
    
    // 使用 os.IsPermission 检查权限错误
    file, err := os.Open("/root/protected.txt")
    if os.IsPermission(err) {
        fmt.Println("权限不足")
    } else if err != nil {
        fmt.Println("其他错误:", err)
    } else {
        file.Close()
        fmt.Println("文件打开成功")
    }
}
```

### 2. 文件操作最佳实践

```go
func fileBestPractices() {
    // 1. 总是检查错误
    file, err := os.Open("important.txt")
    if err != nil {
        fmt.Println("错误:", err)
        return
    }
    
    // 2. 使用 defer 确保资源释放
    defer file.Close()
    
    // 3. 使用缓冲区提高读写性能
    buffer := make([]byte, 4096) // 4KB缓冲区
    for {
        n, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println("读取错误:", err)
            return
        }
        
        // 处理数据
        fmt.Print(string(buffer[:n]))
    }
    
    // 4. 处理大文件时使用流式处理，避免一次性加载到内存
}
```

## 🔄 跨平台注意事项

```go
func crossPlatform() {
    // 使用 os.PathSeparator 而不是硬编码的斜杠
    path := string(os.PathSeparator) + "path" + string(os.PathSeparator) + "to" + string(os.PathSeparator) + "file"
    fmt.Println("路径:", path)
    
    // 使用 filepath 包处理路径（更推荐）
    path = filepath.Join("path", "to", "file")
    fmt.Println("规范化路径:", path)
    
    // 检查操作系统
    fmt.Println("操作系统:", runtime.GOOS)
    if runtime.GOOS == "windows" {
        fmt.Println("这是Windows系统")
    } else if runtime.GOOS == "linux" {
        fmt.Println("这是Linux系统")
    } else if runtime.GOOS == "darwin" {
        fmt.Println("这是macOS系统")
    }
}
```

Go 语言的 `os` 包提供了丰富的系统交互功能，是进行文件操作、进程管理和系统信息获取的基础。掌握这些功能对于开发系统工具、服务器应用和跨平台程序非常重要。在实际使用中，应该始终注意错误处理和资源管理，以确保程序的健壮性和可靠性。