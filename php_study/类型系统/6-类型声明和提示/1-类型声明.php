<?php
// 标量类型声明
function add(int $a, int $b): int {
    return $a + $b;
}

// 联合类型（PHP 8.0+）
function process(int|string $input): int|string {
    return is_int($input) ? $input * 2 : strtoupper($input);
}

// 混合类型
function handle(mixed $data): void {
    // 可以接受任何类型
    var_dump($data);
}

// 返回void和never
function logMessage(string $msg): void {
    file_put_contents('log.txt', $msg, FILE_APPEND);
    // 不能有return值
}

function redirect(string $url): never {
    header("Location: $url");
    exit; // 必须终止执行
}

// 可空类型
function findUser(?int $id): ?string {
    return $id ? "用户$id" : null;
}
?>