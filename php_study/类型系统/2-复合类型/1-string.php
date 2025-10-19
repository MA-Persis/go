<?php
// 四种定义方式
$single = '单引号';
$double = "双引号";
$heredoc = <<<EOT
heredoc语法
 可以包含变量: $single
EOT;
$nowdoc = <<<'EOT'
 nowdoc语法
 不解析变量: $single
EOT;

echo $single. PHP_EOL .$double;
echo "\n";
echo $heredoc. PHP_EOL ;
echo $nowdoc;

// 字符串操作
$str = "Hello";
echo $str[1];        // e
echo substr($str, 1, 3); // ell
echo strlen($str);   // 5

// 数字字符串
$num_str = "123";
$not_num = "123abc";

var_dump(is_numeric($num_str));    // bool(true)
var_dump(is_numeric($not_num));    // bool(false)

// 自动转换
$sum = "123" + "456"; // 579 (整数)
var_dump($sum); // int(579)
?>