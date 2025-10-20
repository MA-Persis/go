在 PHP 中，`empty()`、`isset()`、`is_int()` 和 `is_numeric()` 都是用于变量检查的函数，但它们的功能和用途有所不同。让我详细解释每个函数：

## 1. empty()

检查变量是否为空（为空返回 `true`，否则返回 `false`）

```php
<?php
// 以下情况 empty() 返回 true：
$var1 = "";          // 空字符串
$var2 = 0;           // 数字 0
$var3 = "0";         // 字符串 "0"
$var4 = null;        // null
$var5 = array();     // 空数组
$var6 = false;       // false
$var7;              // 未赋值的变量

var_dump(empty($var1)); // bool(true)
var_dump(empty($var2)); // bool(true)
var_dump(empty($var3)); // bool(true)
var_dump(empty($var4)); // bool(true)
var_dump(empty($var5)); // bool(true)
var_dump(empty($var6)); // bool(true)
var_dump(empty($var7)); // bool(true)

// 以下情况 empty() 返回 false：
$var8 = "Hello";
$var9 = 1;
$var10 = array(1,2,3);

var_dump(empty($var8));  // bool(false)
var_dump(empty($var9));  // bool(false)
var_dump(empty($var10)); // bool(false)
?>
```

## 2. isset()

检查变量是否已设置且不为 null

```php
<?php
$var1 = "Hello";
$var2 = 0;
$var3 = null;
$var4; // 未赋值

var_dump(isset($var1)); // bool(true) - 已设置且不为null
var_dump(isset($var2)); // bool(true) - 已设置且不为null
var_dump(isset($var3)); // bool(false) - 为null
var_dump(isset($var4)); // bool(false) - 未设置

// 检查多个变量
var_dump(isset($var1, $var2)); // bool(true) - 两个都设置
var_dump(isset($var1, $var3)); // bool(false) - 有一个未设置
?>
```

## 3. is_int()

检查变量是否为整数类型

```php
<?php
$var1 = 42;
$var2 = "42";
$var3 = 42.5;
$var4 = 0;

var_dump(is_int($var1)); // bool(true) - 整数
var_dump(is_int($var2)); // bool(false) - 字符串
var_dump(is_int($var3)); // bool(false) - 浮点数
var_dump(is_int($var4)); // bool(true) - 整数

// 别名：is_integer()
var_dump(is_integer($var1)); // bool(true)
?>
```

## 4. is_numeric()

检查变量是否为数字或数字字符串

```php
<?php
$var1 = 42;
$var2 = "42";
$var3 = "42.5";
$var4 = "42abc";
$var5 = "Hello";
$var6 = 42.5;

var_dump(is_numeric($var1)); // bool(true) - 整数
var_dump(is_numeric($var2)); // bool(true) - 数字字符串
var_dump(is_numeric($var3)); // bool(true) - 浮点数字符串
var_dump(is_numeric($var4)); // bool(false) - 包含非数字字符
var_dump(is_numeric($var5)); // bool(false) - 非数字字符串
var_dump(is_numeric($var6)); // bool(true) - 浮点数
?>
```

## 综合比较示例

```php
<?php
function checkVariables($value, $name) {
    echo "检查 $name ('$value'):\n";
    echo "empty(): " . (empty($value) ? 'true' : 'false') . "\n";
    echo "isset(): " . (isset($value) ? 'true' : 'false') . "\n";
    echo "is_int(): " . (is_int($value) ? 'true' : 'false') . "\n";
    echo "is_numeric(): " . (is_numeric($value) ? 'true' : 'false') . "\n";
    echo "---\n";
}

// 测试各种值
checkVariables(0, "0");
checkVariables("0", "\"0\"");
checkVariables(42, "42");
checkVariables("42", "\"42\"");
checkVariables("", "空字符串");
checkVariables(null, "null");
?>
```

## 实际应用场景

```php
<?php
// 表单验证示例
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $username = $_POST['username'] ?? '';
    $age = $_POST['age'] ?? '';
    
    // 检查用户名是否为空
    if (empty($username)) {
        echo "用户名不能为空";
    }
    
    // 检查年龄是否为有效数字
    if (!isset($age) || !is_numeric($age)) {
        echo "年龄必须是数字";
    } elseif (!is_int($age + 0)) { // 转换为数字后检查是否为整数
        echo "年龄必须是整数";
    } elseif ($age < 0 || $age > 150) {
        echo "年龄必须在0-150之间";
    } else {
        echo "验证通过";
    }
}

// 安全地访问数组元素
$config = [
    'debug' => '1',
    'port' => '8080'
];

// 安全的检查方式
if (isset($config['debug']) && is_numeric($config['debug'])) {
    $debug = (int)$config['debug'];
} else {
    $debug = 0;
}
?>
```

## 总结

| 函数 | 用途 | 返回值 |
|------|------|--------|
| `empty()` | 检查是否为空值 | 空值返回 true |
| `isset()` | 检查变量是否存在且不为 null | 存在且不为 null 返回 true |
| `is_int()` | 检查是否为整数类型 | 整数返回 true |
| `is_numeric()` | 检查是否为数字或数字字符串 | 数字或数字字符串返回 true |

这些函数在表单验证、数据过滤和变量检查中非常有用，选择合适的函数可以提高代码的健壮性和安全性。