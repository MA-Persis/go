在 PHP 中，?int 是一种 类型声明（Type Hinting），表示参数 $userId 可以是 int 类型 或 null。  

具体解释

• int：表示参数必须是整数类型。

• ?：表示该参数 允许为 null（即“可选的”）。

• 组合 ?int：表示 $userId 可以是 整数 或 null。

示例代码

function getUsername(?int $userId): string {
    return $userId ? "user_$userId" : 'guest';
}

调用方式

echo getUsername(123);   // 输出 "user_123"
echo getUsername(null);   // 输出 "guest"
echo getUsername("abc");  // 报错！必须是 int 或 null


为什么用 ?int？

1. 明确参数类型：  
   • 强制 $userId 只能是 int 或 null，避免意外类型（如字符串、数组）。

2. 替代旧版写法：  
   在 PHP 7.0 之前，需要用注释或手动检查：
   // PHP 5.x 的写法（无类型声明）
   function getUsername($userId) {
       if ($userId !== null && !is_int($userId)) {
           throw new InvalidArgumentException("必须是整数或 null");
       }
       return $userId ? "user_$userId" : 'guest';
   }
   
   PHP 7.1+ 的 ?int 让代码更简洁、安全。

3. 与返回类型结合：  
   函数返回值 : string 表示必须返回字符串，与参数类型声明形成完整约束。

其他类似用法

类型声明 含义 示例

?int 可空的整数 function foo(?int $num)

?string 可空的字符串 function bar(?string $name)

?array 可空的数组 function baz(?array $list)

int 非空整数 function qux(int $num)

string|null 同 ?string（PHP 8.0+） function quux(string|null $str)

常见问题

Q：如果传 string 或 float 会怎样？

PHP 会抛出 TypeError 异常：
getUsername("123"); // TypeError: 必须是 int 或 null，string 不行
getUsername(3.14);  // TypeError: float 也不行


Q：默认参数 null 可以省略吗？

可以，但需要显式声明默认值：
function getUsername(?int $userId = null): string {
    return $userId ? "user_$userId" : 'guest';
}

调用时：
getUsername(); // 合法，$userId 默认为 null


总结

• ?int 表示参数 可以是整数或 null。

• 目的是 增强代码健壮性，避免意外类型错误。

• PHP 7.1+ 支持，替代了旧版的手动类型检查。

这种语法在数据库查询、API 参数处理等场景中非常有用！ 🚀