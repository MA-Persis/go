<?php
// $_SERVER - 服务器和执行环境信息
echo "服务器软件: " . ($_SERVER['SERVER_SOFTWARE'] ?? '未知') . "\n";
echo "请求方法: " . ($_SERVER['REQUEST_METHOD'] ?? '未知') . "\n";
echo "用户IP: " . ($_SERVER['REMOTE_ADDR'] ?? '未知') . "\n";
echo "脚本路径: " . ($_SERVER['SCRIPT_FILENAME'] ?? '未知') . "\n";

// $_GET - GET请求参数
// 访问: http://example.com?name=John&age=25
if (isset($_GET['name'])) {
    echo "姓名: " . htmlspecialchars($_GET['name']) . "\n";
}

// $_POST - POST请求参数
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $username = $_POST['username'] ?? '未提供';
    $password = $_POST['password'] ?? '未提供';
}

// $_REQUEST - GET, POST, COOKIE 的合并（不推荐使用）
$action = $_REQUEST['action'] ?? 'default';

// $_SESSION - 会话变量
session_start();
$_SESSION['user_id'] = 123;
$_SESSION['username'] = 'john_doe';

// $_COOKIE - HTTP Cookies
$theme = $_COOKIE['theme'] ?? 'light';

// $_FILES - 文件上传
if (isset($_FILES['avatar'])) {
    $file_name = $_FILES['avatar']['name'];
    $file_tmp = $_FILES['avatar']['tmp_name'];
}

// $_ENV - 环境变量
$db_host = $_ENV['DB_HOST'] ?? 'localhost';

// $GLOBALS - 引用全局作用域中的所有变量
function test_global() {
    $GLOBALS['global_var'] = "我是全局变量";
}
?>