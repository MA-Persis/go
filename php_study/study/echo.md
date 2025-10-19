在 PHP 中，echo 用于输出内容，换行可以通过以下几种方式实现：

1. 在 HTML 中换行（适用于浏览器输出）

如果 PHP 输出到 HTML 页面，可以使用 <br> 标签换行：
echo "第一行<br>第二行";

浏览器显示效果：

第一行
第二行


2. 在命令行或纯文本中换行

如果 PHP 在 命令行 运行（如脚本、日志），使用 \n（Unix/Linux）或 \r\n（Windows）换行：
echo "第一行\n第二行";  // Linux/Mac/Unix
echo "第一行\r\n第二行"; // Windows

终端显示效果：

第一行
第二行


3. 使用 PHP 的 PHP_EOL 常量（跨平台兼容）

PHP_EOL 会根据操作系统自动选择 \n 或 \r\n：
echo "第一行" . PHP_EOL . "第二行";

适合需要兼容不同操作系统的脚本。

4. 多行字符串（Heredoc 语法）

如果需要输出多行文本，可以用 Heredoc 语法：
echo <<<EOD
第一行
第二行
第三行
EOD;

输出效果：

第一行
第二行
第三行


5. 连续使用多个 echo

每个 echo 默认不自动换行，但可以手动添加换行符：
echo "第一行";
echo "\n"; // 或 echo "<br>";
echo "第二行";


总结

场景 换行方式 示例

HTML 页面 <br> echo "A<br>B";

命令行/日志 \n 或 \r\n echo "A\nB";

跨平台脚本 PHP_EOL echo "A" . PHP_EOL . "B";

多行文本 Heredoc echo <<<EOD ... EOD;

根据你的需求选择合适的方式即可！ 🚀