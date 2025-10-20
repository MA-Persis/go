<?php
$server_info = [
    'PHP_SELF' => $_SERVER['PHP_SELF'] ?? '',           // 当前脚本路径
    'SERVER_NAME' => $_SERVER['SERVER_NAME'] ?? '',     // 服务器名
    'HTTP_HOST' => $_SERVER['HTTP_HOST'] ?? '',         // 主机名
    'HTTP_USER_AGENT' => $_SERVER['HTTP_USER_AGENT'] ?? '', // 用户代理
    'REMOTE_ADDR' => $_SERVER['REMOTE_ADDR'] ?? '',     // 客户端IP
    'REQUEST_URI' => $_SERVER['REQUEST_URI'] ?? '',     // 请求URI
    'REQUEST_METHOD' => $_SERVER['REQUEST_METHOD'] ?? '', // 请求方法
    'QUERY_STRING' => $_SERVER['QUERY_STRING'] ?? '',   // 查询字符串
    'DOCUMENT_ROOT' => $_SERVER['DOCUMENT_ROOT'] ?? '', // 文档根目录
];
?>