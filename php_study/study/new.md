在 PHP 中，new 关键字用于 实例化一个对象，即根据类（class）创建一个具体的对象实例。它是面向对象编程（OOP）的核心操作之一。

基本语法

$object = new ClassName();

• ClassName：要实例化的类名。

• ()：可传递参数给类的构造函数（如果类有定义 __construct()）。

示例代码

1. 定义一个类

class Car {
    public $color;

    // 构造函数（可选）
    public function __construct($color) {
        $this->color = $color;
    }

    public function drive() {
        echo "The {$this->color} car is driving.";
    }
}


2. 使用 new 实例化对象

$myCar = new Car("red");  // 创建 Car 类的实例，并传递参数给构造函数
$myCar->drive();          // 调用对象的方法

输出：

The red car is driving.


关键细节

1. 构造函数 __construct()

• 在实例化时自动调用。

• 用于初始化对象的属性或执行必要操作。

• 如果没有定义构造函数，可以省略 ()：
  $obj = new ClassName;  // 无括号（无构造函数时）
  

2. 匿名类（PHP 7+）

可以直接在 new 后定义一次性使用的类：
$obj = new class {
    public function sayHello() {
        echo "Hello!";
    }
};
$obj->sayHello();  // 输出 "Hello!"


3. 动态类名

类名可以用变量表示：
$className = "Car";
$myCar = new $className();  // 等效于 new Car()


4. 对象赋值是引用传递

PHP 中对象默认通过引用传递：
$car1 = new Car("blue");
$car2 = $car1;          // $car2 和 $car1 指向同一个对象
$car2->color = "green";
echo $car1->color;      // 输出 "green"（两者同步修改）

如需复制对象，用 clone：
$car2 = clone $car1;    // 创建新副本


与其他语言的对比

特性 PHP Java / C# JavaScript

实例化语法 new ClassName() new ClassName() new ClassName()

构造函数 __construct() ClassName() constructor()

匿名类 支持（PHP 7+） 支持 支持

动态类名 支持（new $className()） 反射实现 支持

常见问题

Q：new 和 ::class 是什么关系？

• ::class 用于获取类的完全限定名称（字符串），常与 new 配合：
  $className = Car::class;  // 返回 "Car"（或带命名空间的完整类名）
  $obj = new $className();
  

Q：new self() 和 new static() 的区别？

• new self()：始终实例化当前类（忽略继承）。

• new static()：延迟绑定，实例化实际调用的类（支持继承）。

示例：
class ParentClass {
    public static function createSelf() {
        return new self();  // 总是 ParentClass
    }
    public static function createStatic() {
        return new static(); // 可能是子类
    }
}

class ChildClass extends ParentClass {}

$obj1 = ChildClass::createSelf();    // ParentClass 实例
$obj2 = ChildClass::createStatic();  // ChildClass 实例


总结

• new 是 PHP 中创建对象的核心关键字。

• 结合构造函数可实现灵活的初始化逻辑。

• 注意对象赋值的引用特性，必要时用 clone。

• 匿名类和动态类名增强了灵活性。

掌握 new 的用法是 PHP 面向对象编程的基础！ 🚀