<?php
$global_var = "我是全局变量";

function test_global_scope() {
    // 方法1：使用 global 关键字
    global $global_var;
    echo $global_var . "\n";
    
    // 方法2：使用 $GLOBALS 数组
    echo $GLOBALS['global_var'] . "\n";
}

test_global_scope();
?>