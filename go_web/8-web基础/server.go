package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// 用户结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 模拟数据库
var users = []User{
	{ID: 1, Name: "张三", Age: 25},
	{ID: 2, Name: "李四", Age: 30},
}

func main() {
	fmt.Println("启动 HTTP 服务器在 :8080 端口...")

	// 注册路由处理函数
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", usersHandler)
	// curl http://localhost:8080/users
	http.HandleFunc("/users/", userDetailHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/health", healthHandler)

	// 启动服务器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("服务器启动失败: %v\n", err)
	}
}

// 首页处理
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Go HTTP 服务器</title>
		<style>
			body { font-family: Arial, sans-serif; margin: 40px; }
			.container { max-width: 800px; margin: 0 auto; }
			.endpoint { background: #f5f5f5; padding: 10px; margin: 10px 0; }
		</style>
	</head>
	<body>
		<div class="container">
			<h1>🎉 Go HTTP 服务器运行中</h1>
			<p>可用的 API 端点：</p>
			<div class="endpoint">
				<strong>GET /</strong> - 首页 (返回 HTML)
			</div>
			<div class="endpoint">
				<strong>GET /users</strong> - 获取所有用户列表 (JSON)
			</div>
			<div class="endpoint">
				<strong>GET /users/{id}</strong> - 获取特定用户信息 (JSON)
			</div>
			<div class="endpoint">
				<strong>GET /time</strong> - 获取服务器当前时间 (JSON)
			</div>
			<div class="endpoint">
				<strong>GET /health</strong> - 健康检查 (JSON)
			</div>
		</div>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

// 获取所有用户
func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	case "POST":
		// 简单的 POST 处理
		var newUser User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "无效的 JSON 数据", http.StatusBadRequest)
			return
		}

		// 生成新 ID
		newUser.ID = len(users) + 1
		users = append(users, newUser)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	default:
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
	}
}

// 获取特定用户
func userDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	// 从 URL 中提取用户 ID
	idStr := r.URL.Path[len("/users/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "无效的用户 ID", http.StatusBadRequest)
		return
	}

	// 查找用户
	for _, user := range users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.NotFound(w, r)
}

// 获取服务器时间
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]interface{}{
		"timestamp": time.Now().Unix(),
		"datetime":  time.Now().Format("2006-01-02 15:04:05"),
		"timezone":  time.Now().Location().String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 健康检查
func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	response := map[string]string{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "go-http-server",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
