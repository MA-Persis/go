<?php
namespace MyApp;

class Controller {
    public function render(self $controller): static {
        // self: 当前类
        // static: 调用者类（后期静态绑定）
        return new static();
    }
}

class AdminController extends Controller {
    // 继承父类方法，但返回AdminController实例
}

$admin = new AdminController();
$result = $admin->render($admin); // 返回AdminController实例
?>