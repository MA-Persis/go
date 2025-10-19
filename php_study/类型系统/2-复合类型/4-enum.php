<?php
// 基本枚举
enum Status: string {
    case PENDING = 'pending';
    case APPROVED = 'approved';
    case REJECTED = 'rejected';
}

// 使用方法枚举
enum Color {
    case RED;
    case GREEN;
    case BLUE;
    
    public function getHex(): string {
        return match($this) {
            self::RED => '#FF0000',
            self::GREEN => '#00FF00',
            self::BLUE => '#0000FF',
        };
    }
}

$status = Status::PENDING;
echo $status->value; // 'pending'

$color = Color::RED;
echo $color->getHex(); // #FF0000
?>