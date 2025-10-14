package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTP 客户端示例
func main() {
	fmt.Println("=== Go HTTP 客户端示例 ===\n")

	// 1. 健康检查
	fmt.Println("1. 健康检查:")
	healthCheck()

	// 2. 获取服务器时间
	fmt.Println("\n2. 服务器时间:")
	getServerTime()

	// 3. 获取所有用户
	fmt.Println("\n3. 所有用户:")
	getAllUsers()

	// 4. 获取特定用户
	fmt.Println("\n4. 用户详情:")
	getUserByID(1)
	getUserByID(999) // 不存在的用户

	// 5. 创建新用户
	fmt.Println("\n5. 创建新用户:")
	createUser()

	// 6. 再次获取所有用户查看结果
	fmt.Println("\n6. 更新后的用户列表:")
	getAllUsers()
}

// 健康检查
func healthCheck() {
	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("状态码: %d\n", resp.StatusCode)
	fmt.Printf("响应: %s\n", string(body))
}

// 获取服务器时间
func getServerTime() {
	resp, err := http.Get("http://localhost:8080/time")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Printf("服务器时间: %s\n", result["datetime"])
	fmt.Printf("时间戳: %.0f\n", result["timestamp"])
}

// 获取所有用户
func getAllUsers() {
	resp, err := http.Get("http://localhost:8080/users")
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var users []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&users)

	for _, user := range users {
		fmt.Printf("ID: %.0f, 姓名: %s, 年龄: %.0f\n",
			user["id"], user["name"], user["age"])
	}
}

// 根据ID获取用户
func getUserByID(id int) {
	url := fmt.Sprintf("http://localhost:8080/users/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("请求失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		fmt.Printf("用户 ID=%d 不存在\n", id)
		return
	}

	var user map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&user)

	fmt.Printf("ID: %.0f, 姓名: %s, 年龄: %.0f\n",
		user["id"], user["name"], user["age"])
}

// 创建新用户
func createUser() {
	newUser := map[string]interface{}{
		"name": "王五",
		"age":  28,
	}

	jsonData, _ := json.Marshal(newUser)

	// 创建带超时的客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Post(
		"http://localhost:8080/users",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Printf("创建用户失败: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		var createdUser map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&createdUser)
		fmt.Printf("创建用户成功: ID=%.0f, 姓名=%s\n",
			createdUser["id"], createdUser["name"])
	} else {
		fmt.Printf("创建用户失败，状态码: %d\n", resp.StatusCode)
	}
}
