<?php
// 接受数组或实现了Traversable的对象
function processItems(iterable $items): void {
    foreach ($items as $item) {
        echo $item . "\n";
    }
}

// 生成器也是iterable
function generateNumbers(int $max): Generator {
    for ($i = 1; $i <= $max; $i++) {
        yield $i;
    }
}

processItems([1, 2, 3]);
processItems(generateNumbers(5));

// 实现Iterator接口
class NumberIterator implements Iterator {
    private $position = 0;
    private $array = [1, 2, 3, 4, 5];
    
    public function current(): mixed { return $this->array[$this->position]; }
    public function key(): mixed { return $this->position; }
    public function next(): void { $this->position++; }
    public function rewind(): void { $this->position = 0; }
    public function valid(): bool { return isset($this->array[$this->position]); }
}
?>


<!-- 这段 PHP 代码定义了一个 NumberIterator 类，它实现了 PHP 内置的 Iterator 接口。这个类的目的是让开发者能够通过迭代器（Iterator）的方式遍历一个固定的数组 [1, 2, 3, 4, 5]，而不是直接操作数组。

代码解析

1. 类定义

class NumberIterator implements Iterator { ... }

• implements Iterator 表示这个类实现了 Iterator 接口，必须实现以下 5 个方法：

  1. current() – 返回当前元素
  2. key() – 返回当前元素的键
  3. next() – 移动到下一个元素
  4. rewind() – 重置迭代器到第一个元素
  5. valid() – 检查当前位置是否有效

2. 类属性

private $position = 0;          // 当前迭代位置（索引）
private $array = [1, 2, 3, 4, 5]; // 要遍历的数组

• $position 记录当前遍历的位置（从 0 开始）。

• $array 是要遍历的数组 [1, 2, 3, 4, 5]。

3. 接口方法实现

(1) current() – 返回当前元素的值

public function current(): mixed {
    return $this->array[$this->position];
}

• 例如，当 $position = 2 时，返回 3。

(2) key() – 返回当前元素的键

public function key(): mixed {
    return $this->position;
}

• 例如，当 $position = 2 时，返回 2（数组索引）。

(3) next() – 移动到下一个元素

public function next(): void {
    $this->position++;
}

• 每次调用后 $position 增加 1，指向下一个元素。

(4) rewind() – 重置迭代器到第一个元素

public function rewind(): void {
    $this->position = 0;
}

• 将 $position 重置为 0，回到数组开头。

(5) valid() – 检查当前位置是否有效

public function valid(): bool {
    return isset($this->array[$this->position]);
}

• 如果 $position 超出数组范围（如 5），返回 false，表示遍历结束。

如何使用这个迭代器？

示例 1：foreach 自动调用迭代器

$iterator = new NumberIterator();
foreach ($iterator as $key => $value) {
    echo "$key: $value\n";
}

输出：

0: 1
1: 2
2: 3
3: 4
4: 5

• foreach 会自动调用 rewind() → valid() → current() → key() → next() 循环执行。

示例 2：手动控制迭代

$iterator = new NumberIterator();
$iterator->rewind(); // 重置到开头

while ($iterator->valid()) {
    $key = $iterator->key();
    $value = $iterator->current();
    echo "$key: $value\n";
    $iterator->next(); // 移动到下一个
}

输出：

0: 1
1: 2
2: 3
3: 4
4: 5


为什么用迭代器（Iterator）？

1. 封装遍历逻辑：隐藏底层数据结构（如数组、数据库结果集）的实现细节。
2. 惰性加载：可以逐项处理数据，而不是一次性加载全部（适合大数据集）。
3. 统一接口：所有迭代器都遵循 Iterator 接口，兼容 foreach 和 PHP 内置函数。

总结

方法 作用 示例调用

current() 返回当前元素值 $value = $iterator->current();

key() 返回当前键 $key = $iterator->key();

next() 移动到下一个元素 $iterator->next();

rewind() 重置到第一个元素 $iterator->rewind();

valid() 检查是否还有元素 if ($iterator->valid()) { ... }

这段代码实现了一个简单的数组迭代器，可以用 foreach 或手动方式遍历数组 [1, 2, 3, 4, 5]。🚀 -->