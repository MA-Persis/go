# go-sql-driver/mysql 介绍

`go-sql-driver/mysql` 是一个纯 Go 语言实现的 MySQL 驱动程序，用于 Go 语言的 `database/sql` 包。

## 主要特性

### 1. **纯 Go 实现**
- 无需 CGO 编译
- 跨平台支持
- 静态链接，部署简单

### 2. **轻量高效**
- 代码简洁，性能优秀
- 内存占用低
- 连接池管理

### 3. **功能完整**
- 支持 MySQL 4.1+ 和 5.x、8.x
- 支持 SSL/TLS 加密连接
- 支持预处理语句
- 支持事务操作

## 安装

```bash
go get -u github.com/go-sql-driver/mysql
```

## 基本用法

### 1. 导入包
```go
import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
```

### 2. 连接数据库
```go
// 连接格式: "用户名:密码@协议(地址:端口)/数据库名?参数"
db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
if err != nil {
    panic(err.Error())
}
defer db.Close()

// 验证连接
err = db.Ping()
if err != nil {
    panic(err.Error())
}
```

### 3. 常用连接参数

| 参数 | 说明 |
|------|------|
| `charset=utf8mb4` | 字符集设置 |
| `parseTime=True` | 将 DATETIME 解析为 time.Time |
| `loc=Local` | 时区设置 |
| `timeout=10s` | 连接超时 |
| `readTimeout=30s` | 读超时 |
| `writeTimeout=30s` | 写超时 |

## 核心功能示例

### 1. 查询操作
```go
type User struct {
    UID        int
    Username   string
    Departname string
    Created    time.Time
}

// 单行查询
var user User
err := db.QueryRow("SELECT uid, username, departname, created FROM user_info WHERE uid = ?", 1).Scan(
    &user.UID, &user.Username, &user.Departname, &user.Created)
if err != nil {
    log.Fatal(err)
}

// 多行查询
rows, err := db.Query("SELECT uid, username, departname, created FROM user_info WHERE departname = ?", "技术部")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
    var user User
    err := rows.Scan(&user.UID, &user.Username, &user.Departname, &user.Created)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(user)
}
```

### 2. 插入操作
```go
// 单条插入
result, err := db.Exec("INSERT INTO user_info (username, departname, created) VALUES (?, ?, ?)",
    "王五", "市场部", time.Now())
if err != nil {
    log.Fatal(err)
}

lastInsertID, err := result.LastInsertId()
rowsAffected, err := result.RowsAffected()
```

### 3. 预处理语句
```go
stmt, err := db.Prepare("INSERT INTO user_info (username, departname, created) VALUES (?, ?, ?)")
if err != nil {
    log.Fatal(err)
}
defer stmt.Close()

// 批量插入
users := []struct {
    username   string
    departname string
}{
    {"张三", "技术部"},
    {"李四", "产品部"},
}

for _, user := range users {
    _, err := stmt.Exec(user.username, user.departname, time.Now())
    if err != nil {
        log.Fatal(err)
    }
}
```

### 4. 事务处理
```go
tx, err := db.Begin()
if err != nil {
    log.Fatal(err)
}

// 在事务中执行多个操作
_, err = tx.Exec("UPDATE user_info SET departname = ? WHERE uid = ?", "运维部", 1)
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

_, err = tx.Exec("INSERT INTO user_info (username, departname) VALUES (?, ?)", "赵六", "测试部")
if err != nil {
    tx.Rollback()
    log.Fatal(err)
}

err = tx.Commit()
if err != nil {
    log.Fatal(err)
}
```

## 连接池配置

```go
db.SetMaxOpenConns(25)      // 最大打开连接数
db.SetMaxIdleConns(5)       // 最大空闲连接数
db.SetConnMaxLifetime(5 * time.Minute)  // 连接最大存活时间
```

## 高级特性

### 1. 支持上下文
```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

rows, err := db.QueryContext(ctx, "SELECT * FROM user_info")
if err != nil {
    log.Fatal(err)
}
defer rows.Close()
```

### 2. NULL 值处理
```go
var name sql.NullString
var number sql.NullInt64

err := db.QueryRow("SELECT username, phone FROM user_info WHERE uid = ?", 1).Scan(&name, &number)
if err != nil {
    log.Fatal(err)
}

if name.Valid {
    fmt.Println("Name:", name.String)
}
```

## 优势总结

1. **性能优秀**：纯 Go 实现，无 CGO 开销
2. **稳定可靠**：经过大量生产环境验证
3. **功能丰富**：支持 MySQL 大部分特性
4. **社区活跃**：维护积极，文档完善
5. **标准兼容**：完美集成 Go 的 database/sql 接口

这是 Go 语言中最流行、最稳定的 MySQL 驱动程序，被广泛应用于各种 Go 项目中。


# 使用 Go Modules 运行上述例子的完整步骤：

## 1. 初始化项目

```bash
# 创建项目目录
mkdir mysql-go-example
cd mysql-go-example

# 初始化 Go Modules
go mod init mysql-go-example
```

## 2. 项目结构
```
mysql-go-example/
├── go.mod
├── go.sum
├── main.go
└── models/
    └── user.go
```

## 3. 编写代码

### main.go
```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	UID        int
	Username   string
	Departname string
	Created    time.Time
}

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "ms:your_password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	defer db.Close()
	
	// 配置连接池
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	
	// 验证连接
	err = db.Ping()
	if err != nil {
		log.Fatal("数据库连接验证失败:", err)
	}
	fmt.Println("数据库连接成功!")
	
	// 创建表
	err = createTable(db)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	
	// 插入数据
	err = insertUser(db, "张三", "技术部")
	if err != nil {
		log.Fatal("插入数据失败:", err)
	}
	
	// 查询数据
	users, err := queryUsers(db)
	if err != nil {
		log.Fatal("查询数据失败:", err)
	}
	
	fmt.Println("查询结果:")
	for _, user := range users {
		fmt.Printf("UID: %d, 用户名: %s, 部门: %s, 创建时间: %s\n", 
			user.UID, user.Username, user.Departname, user.Created.Format("2006-01-02"))
	}
	
	// 事务示例
	err = transactionExample(db)
	if err != nil {
		log.Fatal("事务执行失败:", err)
	}
}

func createTable(db *sql.DB) error {
	createSQL := `
	CREATE TABLE IF NOT EXISTS user_info (
		uid INT(10) NOT NULL AUTO_INCREMENT,
		username VARCHAR(64) NULL DEFAULT NULL,
		departname VARCHAR(64) NULL DEFAULT NULL,
		created DATE NULL DEFAULT NULL,
		PRIMARY KEY (uid)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`
	
	_, err := db.Exec(createSQL)
	return err
}

func insertUser(db *sql.DB, username, departname string) error {
	result, err := db.Exec(
		"INSERT INTO user_info (username, departname, created) VALUES (?, ?, ?)",
		username, departname, time.Now(),
	)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	fmt.Printf("插入成功，ID: %d\n", id)
	return nil
}

func queryUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT uid, username, departname, created FROM user_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UID, &user.Username, &user.Departname, &user.Created)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

func transactionExample(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	
	// 更新操作
	_, err = tx.Exec("UPDATE user_info SET departname = ? WHERE username = ?", "运维部", "张三")
	if err != nil {
		return err
	}
	
	// 插入操作
	_, err = tx.Exec(
		"INSERT INTO user_info (username, departname, created) VALUES (?, ?, ?)",
		"李四", "产品部", time.Now(),
	)
	if err != nil {
		return err
	}
	
	fmt.Println("事务执行成功!")
	return nil
}
```

### models/user.go（可选，用于模块化）
```go
package models

import (
	"database/sql"
	"time"
)

type User struct {
	UID        int       `json:"uid"`
	Username   string    `json:"username"`
	Departname string    `json:"departname"`
	Created    time.Time `json:"created"`
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Create(user *User) error {
	result, err := r.DB.Exec(
		"INSERT INTO user_info (username, departname, created) VALUES (?, ?, ?)",
		user.Username, user.Departname, user.Created,
	)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	user.UID = int(id)
	return nil
}

func (r *UserRepository) FindByID(uid int) (*User, error) {
	var user User
	err := r.DB.QueryRow(
		"SELECT uid, username, departname, created FROM user_info WHERE uid = ?",
		uid,
	).Scan(&user.UID, &user.Username, &user.Departname, &user.Created)
	
	if err != nil {
		return nil, err
	}
	
	return &user, nil
}

func (r *UserRepository) FindAll() ([]User, error) {
	rows, err := r.DB.Query("SELECT uid, username, departname, created FROM user_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UID, &user.Username, &user.Departname, &user.Created)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, nil
}
```

## 4. 下载依赖

```bash
# 下载 go-sql-driver/mysql
go mod tidy

# 或者手动下载特定版本
go get github.com/go-sql-driver/mysql@v1.7.1
```

## 5. 运行项目

```bash
# 运行主程序
go run main.go

# 或者编译后运行
go build -o app
./app
```

## 6. 环境配置

### 创建配置文件 config.go
```go
package main

import (
	"os"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func LoadConfig() *Config {
	return &Config{
		DBUsername: getEnv("DB_USERNAME", "ms"),
		DBPassword: getEnv("DB_PASSWORD", "your_password"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBName:     getEnv("DB_NAME", "test"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}
```

## 7. 使用环境变量

创建 `.env` 文件（需要安装相关包）或使用系统环境变量：

```bash
# 设置环境变量
export DB_USERNAME=ms
export DB_PASSWORD=your_secure_password
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_NAME=test

# 然后运行程序
go run main.go
```

## 8. 测试连接

在运行前确保：
1. MySQL 服务正在运行
2. 数据库 `test` 已创建
3. 用户 `ms` 有访问权限
4. 连接字符串中的密码正确

## 9. 依赖管理

查看当前的依赖：
```bash
go list -m all
go mod graph
```

更新依赖：
```bash
go get -u github.com/go-sql-driver/mysql
go mod tidy
```

这样你就可以使用 Go Modules 来管理依赖并运行 MySQL 示例程序了。