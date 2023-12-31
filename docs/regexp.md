```
go doc regexp/syntax
```    
    
### 单一
```
.                   匹配任意一个字符，如果设置 s = true，则可以匹配换行符

[字符类]            匹配“字符类”中的一个字符，“字符类”见后面的说明
[^字符类]           匹配“字符类”外的一个字符，“字符类”见后面的说明

\小写Perl标记       匹配“Perl类”中的一个字符，“Perl类”见后面的说明
\大写Perl标记       匹配“Perl类”外的一个字符，“Perl类”见后面的说明

[:ASCII类名:]       匹配“ASCII类”中的一个字符，“ASCII类”见后面的说明
[:^ASCII类名:]      匹配“ASCII类”外的一个字符，“ASCII类”见后面的说明

\pUnicode普通类名   匹配“Unicode类”中的一个字符(仅普通类)，“Unicode类”见后面的说明
\PUnicode普通类名   匹配“Unicode类”外的一个字符(仅普通类)，“Unicode类”见后面的说明

\p{Unicode类名}     匹配“Unicode类”中的一个字符，“Unicode类”见后面的说明
\P{Unicode类名}     匹配“Unicode类”外的一个字符，“Unicode类”见后面的说明
```
复核
```
    xy             匹配 xy（x 后面跟随 y）
    x|y            匹配 x 或 y (优先匹配 x)
```

重复
```
    x*             匹配零个或多个 x，优先匹配更多(贪婪)
    x+             匹配一个或多个 x，优先匹配更多(贪婪)
    x?             匹配零个或一个 x，优先匹配一个(贪婪)
    x{n,m}         匹配 n 到 m 个 x，优先匹配更多(贪婪)
    x{n,}          匹配 n 个或多个 x，优先匹配更多(贪婪)
    x{n}           只匹配 n 个 x
    x*?            匹配零个或多个 x，优先匹配更少(非贪婪)
    x+?            匹配一个或多个 x，优先匹配更少(非贪婪)
    x??            匹配零个或一个 x，优先匹配零个(非贪婪)
    x{n,m}?        匹配 n 到 m 个 x，优先匹配更少(非贪婪)
    x{n,}?         匹配 n 个或多个 x，优先匹配更少(非贪婪)
    x{n}?          只匹配 n 个 x

```

> 贪婪模式在整个表达式匹配成功的前提下，尽可能多的匹配，而非贪婪模式在整个表达式匹配成功的前提下，尽可能少的匹配。

### 分组
```
 (子表达式)            被捕获的组，该组被编号 (子匹配)
    (?P<命名>子表达式)    被捕获的组，该组被编号且被命名 (子匹配)
    (?:子表达式)          非捕获的组 (子匹配)
    (?标记)               在组内设置标记，非捕获，标记影响当前组后的正则表达式
    (?标记:子表达式)      在组内设置标记，非捕获，标记影响当前组内的子表达式

    标记的语法是：
    xyz  (设置 xyz 标记)
    -xyz (清除 xyz 标记)
    xy-z (设置 xy 标记, 清除 z 标记)

    可以设置的标记有：
    i              不区分大小写 (默认为 false)
    m              多行模式：让 ^ 和 $ 匹配整个文本的开头和结尾，而非行首和行尾(默认为 false)
    s              让 . 匹配 \n (默认为 false)
    U              非贪婪模式：交换 x* 和 x*? 等的含义 (默认为 false)
```

### 位置标记：
```
    ^              如果标记 m=true 则匹配行首，否则匹配整个文本的开头（m 默认为 false）
    $              如果标记 m=true 则匹配行尾，否则匹配整个文本的结尾（m 默认为 false）
    \A             匹配整个文本的开头，忽略 m 标记
    \b             匹配单词边界
    \B             匹配非单词边界
    \z             匹配整个文本的结尾，忽略 m 标记
```

### 转义序列：
```
   \a             匹配响铃符    （相当于 \x07）
                   注意：正则表达式中不能使用 \b 匹配退格符，因为 \b 被用来匹配单词边界，
                   可以使用 \x08 表示退格符。
    \f             匹配换页符    （相当于 \x0C）
    \t             匹配横向制表符（相当于 \x09）
    \n             匹配换行符    （相当于 \x0A）
    \r             匹配回车符    （相当于 \x0D）
    \v             匹配纵向制表符（相当于 \x0B）
    \123           匹配 8  進制编码所代表的字符（必须是 3 位数字）
    \x7F           匹配 16 進制编码所代表的字符（必须是 3 位数字）
    \x{10FFFF}     匹配 16 進制编码所代表的字符（最大值 10FFFF  ）
    \Q...\E        匹配 \Q 和 \E 之间的文本，忽略文本中的正则语法

    \\             匹配字符 \
    \^             匹配字符 ^
    \$             匹配字符 $
    \.             匹配字符 .
    \*             匹配字符 *
    \+             匹配字符 +
    \?             匹配字符 ?
    \{             匹配字符 {
    \}             匹配字符 }
    \(             匹配字符 (
    \)             匹配字符 )
    \[             匹配字符 [
    \]             匹配字符 ]
    \|             匹配字符 |
```

### 可以将“命名字符类”作为“字符类”的元素

```
 [\d]           匹配数字 (相当于 \d)
    [^\d]          匹配非数字 (相当于 \D)
    [\D]           匹配非数字 (相当于 \D)
    [^\D]          匹配数字 (相当于 \d)
    [[:name:]]     命名的“ASCII 类”包含在“字符类”中 (相当于 [:name:])
    [^[:name:]]    命名的“ASCII 类”不包含在“字符类”中 (相当于 [:^name:])
    [\p{Name}]     命名的“Unicode 类”包含在“字符类”中 (相当于 \p{Name})
    [^\p{Name}]    命名的“Unicode 类”不包含在“字符类”中 (相当于 \P{Name})
```


> “字符类”取值如下（“字符类”包含“Perl类”、“ASCII类”、“Unicode类”）：
```
    x                    单个字符
    A-Z                  字符范围(包含首尾字符)
    \小写字母            Perl类
    [:ASCII类名:]        ASCII类
    \p{Unicode脚本类名}  Unicode类 (脚本类)
    \pUnicode普通类名    Unicode类 (普通类)
```
    
> “Perl 类”取值如下：

```
\d             数字 (相当于 [0-9])
\D             非数字 (相当于 [^0-9])
\s             空白 (相当于 [\t\n\f\r ])
\S             非空白 (相当于[^\t\n\f\r ])
\w             单词字符 (相当于 [0-9A-Za-z_])
\W             非单词字符 (相当于 [^0-9A-Za-z_])
```

### “ASCII 类”取值如下
```
[:alnum:]      字母数字 (相当于 [0-9A-Za-z])
[:alpha:]      字母 (相当于 [A-Za-z])
[:ascii:]      ASCII 字符集 (相当于 [\x00-\x7F])
[:blank:]      空白占位符 (相当于 [\t ])
[:cntrl:]      控制字符 (相当于 [\x00-\x1F\x7F])
[:digit:]      数字 (相当于 [0-9])
[:graph:]      图形字符 (相当于 [!-~])
[:lower:]      小写字母 (相当于 [a-z])
[:print:]      可打印字符 (相当于 [ -~] 相当于 [ [:graph:]])
[:punct:]      标点符号 (相当于 [!-/:-@[-反引号{-~])
[:space:]      空白字符(相当于 [\t\n\v\f\r ])
[:upper:]      大写字母(相当于 [A-Z])
[:word:]       单词字符(相当于 [0-9A-Za-z_])
[:xdigit:]     16 進制字符集(相当于 [0-9A-Fa-f])
```
