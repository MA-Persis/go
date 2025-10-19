<?php
$var = null;
var_dump($var); // NULL

// 判断是否为null
if ($var === null) {
    echo "变量为null";
}

// 注意事项
$unset_var; // 未赋值的变量会产生警告
var_dump($unset_var); // NULL + 警告
?>

<!-- 
NULL
变量为nullPHP Warning:  Undefined variable $unset_var in /home/mashen/pnp_study/类型系统/1-标量类型/1-NULL.php on line 12
NULL 
-->