<?php
// 类型检测函数
$vars = [
    123,               // integer
    45.67,             // float
    "hello",           // string
    true,              // boolean
    null,              // null
    [1, 2, 3],         // array
    new stdClass(),    // object
    fopen(__FILE__, 'r') // resource
];

foreach ($vars as $var) {
    echo gettype($var) . " - " . 
         (is_scalar($var) ? "标量" : "非标量") . "\n";
}

// 类型判断函数使用
$value = "123";

var_dump(is_int($value));        // false
var_dump(is_numeric($value));    // true
var_dump(isset($value));         // true
var_dump(empty($value));         // false
?>