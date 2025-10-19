<?php
$true = true;
$false = false;

// 转换为布尔值为false的情况
$false_cases = [
    0,        // 整型0
    0.0,      // 浮点0.0
    "",       // 空字符串
    "0",      // 字符串"0"
    [],       // 空数组
    null      // null
];

// 严格比较
if ($true === true) {
    echo "严格为真";
}
?>