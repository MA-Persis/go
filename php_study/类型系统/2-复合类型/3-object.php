<?php
class Person {
    public string $name;
    private int $age;
    
    public function __construct(string $name, int $age) {
        $this->name = $name;
        $this->age = $age;
    }
    
    public function getInfo(): string {
        return "{$this->name}, {$this->age}岁";
    }
}

$person = new Person("张三", 25);
echo $person->getInfo();

// 对象比较
$p1 = new Person("李四", 30);
$p2 = new Person("李四", 30);
$p3 = $p1;

var_dump($p1 == $p2);  // true (值相等)
var_dump($p1 === $p3); // true (同一实例)
?>