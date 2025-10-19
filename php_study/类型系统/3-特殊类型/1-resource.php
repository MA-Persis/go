<?php
// 文件资源
$file = fopen('test.txt', 'r');
var_dump($file); // resource(5) of type (stream)

// 数据库连接
$conn = mysqli_connect('localhost', 'user', 'pass', 'db');

// 释放资源
fclose($file);
mysqli_close($conn);

// 注意事项：资源不能被序列化
?>