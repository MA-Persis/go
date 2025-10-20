<?php
// 变量声明和赋值
$variable = "值";
$userName = "张三";    // 驼峰命名
$user_age = 25;       // 蛇形命名

// 变量命名规则
$valid_names = [
    '$name',          // 正确
    '$user_name',     // 正确
    '$userName',      // 正确
    '$_temp',         // 正确
    '$name1',         // 正确
];

$invalid_names = [
    '$123name',       // 错误：不能以数字开头
    '$user-name',     // 错误：不能包含连字符
    '$user name',     // 错误：不能包含空格
    '$_#temp',        // 错误：不能包含特殊字符
];

// 变量引用
$a = 1;
$b = &$a;     // $b 是 $a 的引用
$a = 2;
echo $b;      // 输出 2

// 变量销毁
unset($a);    // 销毁变量 $a
?>