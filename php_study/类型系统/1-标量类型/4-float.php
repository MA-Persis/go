<?php
$a = 1.234;
$b = 1.2e3;     // 1200
$c = 7E-10;     // 0.0000000007

// 浮点数精度问题
echo (0.1 + 0.2) == 0.3; // false
echo bccomp(0.1 + 0.2, 0.3, 1); // 0 (相等)

// NAN和INF
$nan = acos(8);  // NAN
$inf = log(0);   // -INF

// 比较浮点数
function float_equal($a, $b, $epsilon = 1e-10) {
    return abs($a - $b) < $epsilon;
}
?>