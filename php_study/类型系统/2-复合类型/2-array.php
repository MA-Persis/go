<?php
// 定义数组
$array1 = array(1, 2, 3);
$array2 = [1, 2, 3]; // 简写语法
$assoc = ['a' => 1, 'b' => 2];

// 数组操作
$array[] = 4;        // 追加
foreach ($array as $value) {
  echo "$value,";
}

unset($array[0]);    // 删除元素
foreach ($array as $value) {
  echo "$value,";
}

echo count($array);       // 获取长度

echo "\n";

// 遍历数组
foreach ($array1 as $value) {
  echo "$value,";
}

echo "\n";

foreach ($array2 as $value) {
  echo "$value,";
}

echo "\n";

foreach ($assoc as $key => $value) {
    echo "$key: $value\n";
}

// 注意事项
$sparse = [0 => 'a', 2 => 'c']; // 稀疏数组
isset($sparse[1]); // false

// 数组函数
$merged = array_merge($array1, $array2);

foreach ($merged as $value) {
  echo "$value,";
}

echo "\n";

$filtered = array_filter($array, function($v) { return $v > 1; });

foreach ($filtered as $value) {
  echo "$value,";
}

echo "\n";
?>