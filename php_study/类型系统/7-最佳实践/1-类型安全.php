<?php
declare(strict_types=1); // 严格类型模式

// 1. 总是使用类型声明
function calculate(int $a, int $b): int {
    return $a + $b;
}

// 2. 使用适当的类型检查
function safeDivide(float $a, float $b): float {
    if ($b == 0) {
        throw new InvalidArgumentException("除数不能为零");
    }
    return $a / $b;
}

// 3. 处理可能为null的情况
function getUsername(?int $userId): string {
    return $userId ? "user_$userId" : 'guest';
}

// 4. 使用枚举代替常量
enum UserRole: string {
    case ADMIN = 'admin';
    case USER = 'user';
    case GUEST = 'guest';
}

// 5. 避免意外的类型转换
$secureCompare = "123" === 123; // false
$insecureCompare = "123" == 123; // true

// 6. 使用类型安全的数组操作
class StringCollection {
    /** @var string[] */
    private array $items = [];
    
    public function add(string $item): void {
        $this->items[] = $item;
    }
    
    /** @return string[] */
    public function all(): array {
        return $this->items;
    }
}
?>