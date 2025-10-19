<?php
// 不同形式的callable
function normalFunction($arg) {
    return $arg * 2;
}

class MyClass {
    public static function staticMethod($arg) {
        return $arg * 3;
    }
    
    public function instanceMethod($arg) {
        return $arg * 4;
    }
}

// 调用方式
$callable1 = 'normalFunction';
$callable2 = ['MyClass', 'staticMethod'];
$callable3 = [new MyClass(), 'instanceMethod'];
$callable4 = function($arg) { return $arg * 5; };

echo $callable1(2); // 4
echo $callable2(2); // 6

// 检查是否可调用
if (is_callable($callable1)) {
    call_user_func($callable1, 10);
}
?>