<?php
// 类型声明可以提升性能（JIT编译优化）
function optimizedAdd(int $a, int $b): int {
    return $a + $b;
}

// 避免不必要的类型转换
// 慢
function slowAdd($a, $b) {
    return (int)$a + (int)$b;
}

// 快
function fastAdd(int $a, int $b): int {
    return $a + $b;
}
?>