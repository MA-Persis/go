<?php
class Singleton {
    private static ?self $instance = null;
    
    // 私有构造函数防止外部实例化
    private function __construct() {}
    
    // 防止克隆
    private function __clone() {}
    
    // 防止反序列化
    public function __wakeup() {
        throw new Exception("Cannot unserialize singleton");
    }
    
    public static function getInstance(): static {
        if (self::$instance === null) {
            self::$instance = new static();
        }
        return self::$instance;
    }
}

$singleton1 = Singleton::getInstance();
$singleton2 = Singleton::getInstance();
var_dump($singleton1 === $singleton2); // true
?>