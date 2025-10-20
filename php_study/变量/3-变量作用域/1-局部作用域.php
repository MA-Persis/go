<?php
function test_local() {
    $local_var = "我是局部变量";
    echo $local_var . "\n"; // 正常工作
}

test_local();
// echo $local_var; // 错误：未定义变量
?>