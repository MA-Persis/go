<?php
// 显式转换（强制类型转换）
$int = (int) "123";       // 123
$float = (float) "123.45"; // 123.45
$string = (string) 123;   // "123"
$bool = (bool) 1;         // true
$array = (array) "hello"; // ["hello"]
$object = (object) ['a' => 1]; // stdClass对象

// 隐式转换
$result = "5" + "10";     // 15 (整数)
$concat = "5" . "10";     // "510" (字符串)

// 使用settype函数
$var = "123";
settype($var, "int");     // $var现在是整数123

// 注意事项
echo (int) "10abc";       // 10
echo (int) "abc10";       // 0
echo (bool) "0";          // false
echo (bool) "false";      // true (!)

// 严格比较避免意外转换
if ("123" === 123) {      // false
    // 不会执行
}
?>