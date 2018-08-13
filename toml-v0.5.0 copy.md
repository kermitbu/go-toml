![TOML Logo](../../logos/toml-200.png)

TOML v0.5.0
===========

全称: Tom的（语义）明显、（配置）最小化的语言。 (Tom's Obvious, Minimal Language. By Tom Preston-Werner.)

TOML从0.5.0版本开始，已经非常稳定了，强烈建议所有实现兼容0.5.0版本，以方便后续切换到1.0.0版本。

Objectives
----------

TOML的目标是成为一个有明显语义而容易去阅读的最小化配置文件格式。 TOML被设计成可以无歧义地被映射为哈希表，从而很容易的被解析成各种语言中的数据结构。
Table of contents
-------

- [示例](#user-content-example)
- [规格](#user-content-spec)
- [注释](#user-content-comment)
- [键/值对](#user-content-keyvalue-pair)
- [键](#user-content-keys)
- [字符串](#user-content-string)
- [整数](#user-content-integer)
- [浮点数](#user-content-float)
- [布尔值](#user-content-boolean)
- [偏移日期时间](#user-content-offset-date-time)
- [本地日期时间](#user-content-local-date-time)
- [本地日期](#user-content-local-date)
- [本地时间](#user-content-local-time)
- [数组](#user-content-array)
- [表](#user-content-table)
- [内联表](#user-content-inline-table)
- [表数组](#user-content-array-of-tables)
- [文件名扩展](#user-content-filename-extension)
- [与其他格式的比较](#user-content-comparison-with-other-formats)
- [参与](#user-content-get-involved)
- [维基](#user-content-wiki)

示例
-------

```toml
# This is a TOML document.

title = "TOML Example"

[owner]
name = "Tom Preston-Werner"
dob = 1979-05-27T07:32:00-08:00 # First class dates

[database]
server = "192.168.1.1"
ports = [ 8001, 8001, 8002 ]
connection_max = 5000
enabled = true

[servers]

  # Indentation (tabs and/or spaces) is allowed but not required
  [servers.alpha]
  ip = "10.0.0.1"
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"

[clients]
data = [ ["gamma", "delta"], [1, 2] ]

# Line breaks are OK when inside arrays
hosts = [
  "alpha",
  "omega"
]
```

规格
----

* TOML是大小写敏感的。
* TOML文件必须只包含UTF-8编码的Unicode字符。
* 空格是指制表符(0x09) 或空格 (0x20)。
* 换行符是指LF(0x0A)或CRLF (0x0D0A)。

注释
-------

用符号#来表示注释：

```toml
# This is a full-line comment
key = "value" # This is a comment at the end of a line
```

键/值对
--------------

TOML文档的主要由键/值对组成。
键位于等号的左侧，值位于右侧。键名和值周围的空格被忽略。键，等号和值必须在同一行（尽管某些值可以在多行上分解）。
```toml
key = "value"
```

值必须是以下类型：String，Integer，Float，Boolean，Datetime，Array或Inline Table。未指定的值无效。

```toml
key = # INVALID
```

Keys
----

键值可以是裸键，带引号的键或点。

**裸键**可能只包含ASCII字母，ASCII数字，下划线和短划线（A-Za-z0-9_-）。
请注意，允许裸键仅由ASCII数字组成，例如`1234`，但始终被解释为字符串。

```toml
key = "value"
bare_key = "value"
bare-key = "value"
1234 = "value"
```

**引号键**遵循与基本字符串或文字字符串完全相同的规则，并允许您使用更广泛的键名称集。除非有特别的必要，否则最佳做法是使用裸键。

```toml
"127.0.0.1" = "value"
"character encoding" = "value"
"ʎǝʞ" = "value"
'key2' = "value"
'quoted "value"' = "value"
```

裸键必须非空，但允许使用空引号键（不鼓励）。

```toml
= "无键名"  # 无效
"" = "空"     # 有效，但不鼓励
'' = '空'     # 有效，但不鼓励
```

**点键** 是使用点(.)连接的一系列裸键和引号键，可以把类似的属性分组在一起:

```toml
name = "Orange"
physical.color = "orange"
physical.shape = "round"
site."google.com" = true
```

如果用JSON来描述这个信息，等同于如下格式：

```json
{
  "name": "Orange",
  "physical": {
    "color": "orange",
    "shape": "round"
  },
  "site": {
    "google.com": true
  }
}
```

被点分隔部分周围的空格都会被忽略，但是最好不要使用任何多余的空格。
多次重复定义键值是无效的

```
# 不要这么做
name = "Tom"
name = "Pradyun"
```

尽管没有直接定义键值，任然可以在其他键值内部使用这个键值的名称。
只要没有直接定义密钥，您仍然可以写入密钥及其中的名称。

```
a.b.c = 1
a.d = 2
```

```
# 这么写是无效的
a.b = 1
a.b.c = 2
```

字符串
------

有四种方法来表示字符串：基本字符串、多行基本字符串、字面量和多行字面量。所有的字符串必须只包含有效的UTF-8字符。

**基本字符串** 是由引号括起来的任意字符串，除了那些必须要转义的，比如：双引号、反斜杠和控制字符(U+0000到U+001F, U+007F)。

```toml
str = "I'm a string. \"You can quote me\". Name\tJos\u00E9\nLocation\tSF."
```

常用的转义序列：

```
\b         - backspace       (U+0008)
\t         - tab             (U+0009)
\n         - linefeed        (U+000A)
\f         - form feed       (U+000C)
\r         - carriage return (U+000D)
\"         - quote           (U+0022)
\\         - backslash       (U+005C)
\uXXXX     - unicode         (U+XXXX)
\UXXXXXXXX - unicode         (U+XXXXXXXX)
```
任意Unicode字符都可能被转义为`\uXXXX` 或 `\UXXXXXXXX`的形式。这些转义代码必须是有效的Unicode[标量值](http://unicode.org/glossary/#unicode_scalar_value).。

所有未出现在上面名单中的转义序列是保留的，如果使用，TOML会产生错误。

有时你需要表达一段文本（比如，翻译文件），或者是将很长的字符串分成多行。TOML很容易处理这种情况。 

**多行基本字符串** 是被三引号括起来的字符串，并且允许换行。紧跟起始界定符后面的换行符会被剪掉，而其他的所有空格和换行字符仍然被保留。

```toml
str1 = """
Roses are red
Violets are blue"""
```

TOML解析器应该能正常处理不同平台下的换行符。

```toml
# On a Unix system, the above multi-line string will most likely be the same as:
str2 = "Roses are red\nViolets are blue"

# On a Windows system, it will most likely be equivalent to:
str3 = "Roses are red\r\nViolets are blue"
```
在行尾使用`\`，可以避免在写长字符串的时候引入多余的空格。 `\` 将会删除当前位置到下个非空字符或结束界定符之间的所有空格（包括换行符）。 如果在起始界定符之后的第一个字符是反斜杠和一个换行符，那么从此位置到下个非空白字符或结束界定符之间的所有空格和换行符都会被剪掉。 所有的转义序列对基本字符串都有效，也对多行基本字符串有效。

```toml
# The following strings are byte-for-byte equivalent:
str1 = "The quick brown fox jumps over the lazy dog."

str2 = """
The quick brown \


  fox jumps over \
    the lazy dog."""

str3 = """\
       The quick brown \
       fox jumps over \
       the lazy dog.\
       """
```

任何Unicode字符都可能被用到，除了那些可能需要转义的字符：反斜杠和控制字符(U+0000 到 U+001F, U+007F)。 引号不需要转义，除非它们的存在可能会造成提前关闭界定符。

如果你需要频繁的指定Windows的路径或正则表达式，那么不得不添加转义符就会变的繁琐和容易出错。 TOML支持完全不允许转义的字面量字符串来帮助你解决此类问题。

**字面量字符串** 是被单引号包含的字符串，跟基本字符串一样，它们一定是以单行出现:
```toml
# What you see is what you get.
winpath  = 'C:\Users\nodejs\templates'
winpath2 = '\\ServerX\admin$\system32\'
quoted   = 'Tom "Dubs" Preston-Werner'
regex    = '<\i\c*\s*>'
```

因为没有转义，所以在一个被单引号封闭的字面量字符串里面没有办法写单引号。 幸运的是，TOML支持多行版本的字面量字符串来解决这个问题。

**多行字面量字符串** 是被三个单引号括起来的字符串，并且允许换行。 跟字面量字符串一样，这也没有任何转义。 紧跟起始界定符的换行符会被剪掉。界定符之间的所有其他内容都会被按照原样解释而无需修改。

```toml
regex2 = '''I [dw]on't need \d{2} apples'''
lines  = '''
The first newline is
trimmed in raw strings.
   All other whitespace
   is preserved.
'''
```

对于二进制数据，建议你使用Base64或其他适合的编码，比如ASCII或UTF-8编码。具体的处理取决于特定的应用。

整数
-------

整数就是没有小数点的数字。正数前面也可以用加号，负数需要用负号前缀表示。

```toml
int1 = +99
int2 = 42
int3 = 0
int4 = -17
```

对于大整数，你可以用下划线提高可读性。每个下划线两边至少包含一个数字。

```toml
int5 = 1_000
int6 = 5_349_221
int7 = 1_2_3_4_5     # 有效，但不鼓励
```

前导零是不允许的。整数值' -0 '和' +0 '是有效的并且与无前缀的零相同。

非负整数值也可以用十六进制、八进制或二进制来表示。在这些格式中，允许前导零(在前缀之后)。十六进制值是不区分大小写。在数字之间可以使用下划线(但不能在前缀和值之间使用下划线)。

```toml
# 十六进制使用`0x`作为前缀
hex1 = 0xDEADBEEF
hex2 = 0xdeadbeef
hex3 = 0xdead_beef

# 八进制使用`0o`作为前缀
oct1 = 0o01234567
oct2 = 0o755 # useful for Unix file permissions

# 二进制使用`0b`作为前缀
bin1 = 0b11010110
```

预期的范围是64位 (signed long)(−9,223,372,036,854,775,808 到 9,223,372,036,854,775,807).

浮点数
-----

浮点数遵循IEEE 754 标准，预期精度为64位

浮点数由一个整数部分(它遵循与整数值相同的规则)和一个小数部分和/或一个指数部分组成。如果同时存在小数部分和指数部分，小数部分必须在指数部分之前。

```toml
# fractional
flt1 = +1.0
flt2 = 3.1415
flt3 = -0.01

# exponent
flt4 = 5e+22
flt5 = 1e6
flt6 = -2E-2

# both
flt7 = 6.626e-34
```

小数部分是小数点后跟一个或多个数字。

指数部分是E（大写或小写），后跟整数部分（遵循与整数值相同的规则）。

与整数类似，您可以使用下划线来增强可读性。每个下划线必须至少包含一个数字。

```toml
flt8 = 9_224_617.445_991_228_313
```

浮点值`-0.0`并且`+0.0`是有效的，应根据IEEE 754进行映射。

也可以表示特殊浮点值。它们总是小写的。

```toml
# 无穷大
sf1 = inf  # 正无穷大
sf2 = +inf # 正无穷大
sf3 = -inf # 负无穷大

# 不是数字
sf4 = nan  # 实际的 sNaN/qNaN 的编码是特定于实现
sf5 = +nan # same as `nan`
sf6 = -nan # 有效的，实际的编码是特定于实现
```

布尔值
-------

布尔值是小写的true和false。

```toml
bool1 = true
bool2 = false
```

Offset Date-Time
---------------

时间日期是RFC 3339中的时间格式。
要明确地表示特定的时间点，您可以使用一个带有偏移量的格式化日期[RFC 3339](http://tools.ietf.org/html/rfc3339)

```toml
odt1 = 1979-05-27T07:32:00Z
odt2 = 1979-05-27T00:32:00-07:00
odt3 = 1979-05-27T00:32:00.999999-07:00
```

为了便于阅读，您可以将日期和时间之间的T分隔符替换为空格(RFC 3339第5.6节)。

```toml
odt4 = 1979-05-27 07:32:00Z
```

小数秒的精度是特定于实现的，但是至少是毫秒级的精度。如果值包含的精度超过实现所支持的精度，则必须截断附加的精度，而不是四舍五入。


Local Date-Time
--------------

If you omit the offset from an [RFC 3339](http://tools.ietf.org/html/rfc3339)
formatted date-time, it will represent the given date-time without any relation
to an offset or timezone. It cannot be converted to an instant in time without
additional information. Conversion to an instant, if required, is implementation
specific.

```toml
ldt1 = 1979-05-27T07:32:00
ldt2 = 1979-05-27T00:32:00.999999
```

The precision of fractional seconds is implementation specific, but at least
millisecond precision is expected. If the value contains greater precision than
the implementation can support, the additional precision must be truncated, not
rounded.

Local Date
----------

If you include only the date portion of an
[RFC 3339](http://tools.ietf.org/html/rfc3339) formatted date-time, it will
represent that entire day without any relation to an offset or timezone.

```toml
ld1 = 1979-05-27
```

Local Time
----------

If you include only the time portion of an [RFC
3339](http://tools.ietf.org/html/rfc3339) formatted date-time, it will represent
that time of day without any relation to a specific day or any offset or
timezone.

```toml
lt1 = 07:32:00
lt2 = 00:32:00.999999
```

The precision of fractional seconds is implementation specific, but at least
millisecond precision is expected. If the value contains greater precision than
the implementation can support, the additional precision must be truncated, not
rounded.

Array
-----

数组是由方括号包括的基本单元。空格会被忽略。数组中的元素由逗号分隔。数据类型不能混用（所有的字符串均为同一类型）。

```toml
arr1 = [ 1, 2, 3 ]
arr2 = [ "red", "yellow", "green" ]
arr3 = [ [ 1, 2 ], [3, 4, 5] ]
arr4 = [ "all", 'strings', """are the same""", '''type''']
arr5 = [ [ 1, 2 ], ["a", "b", "c"] ]

arr6 = [ 1, 2.0 ] # INVALID
```

数组也可以多行。所以，除了忽略空格之外，数组也忽略了括号之间的换行符。在结束括号之前存在逗号是可以的。

```toml
arr7 = [
  1, 2, 3
]

arr8 = [
  1,
  2, # this is ok
]
```

Table
-----
表（也被称为哈希表或字典）是键值对集合。表格名由方括号包裹，自成一行。注意和数组相区分，数组里只有值。

```toml
[table]
```
在表名之下，直到下一个表或文件尾（EOF）之间都是该表的键值对。键是等号符左边的值，值是等号符右边的值。

```toml
[table-1]
key1 = "some string"
key2 = 123

[table-2]
key1 = "another string"
key2 = 456
```

表的明明规则与键相同（详见上面关于键的定义）

```toml
[dog."tater.man"]
type.name = "pug"
```

如果用JSON来描述这个信息，等同于如下格式：

```json
{ "dog": { "tater.man": { "type": { "name": "pug" } } } }
```

被点分隔部分周围的空格都会被忽略，但是最好不要使用任何多余的空格。

```toml
[a.b.c]            # this is best practice
[ d.e.f ]          # same as [d.e.f]
[ g .  h  . i ]    # same as [g.h.i]
[ j . "ʞ" . 'l' ]  # same as [j."ʞ".'l']
```

如果你不想，你可以完全不去指定父表（super-tables）。TOML知道该如何处理。

```toml
# [x] you
# [x.y] don't
# [x.y.z] need these
[x.y.z.w] # for this to work
```

空表是允许的，其中没有键值对。
像键一样，不能够多次定义同一个表。

```
# 不要这么做

[a]
b = 1

[a]
c = 2
```

```
# 也不要这么做

[a]
b = 1

[a.b]
c = 2
```

内联表
------------
内联表为表示表提供了更紧凑的语法。它们对于分组数据特别有用，否则这些数据会很快变得冗长。内联表包含在大括号“{”和“}”中。在大括号中，可能出现0或更多逗号分隔的键/值对。键/值对采用与标准表中的键/值对相同的形式。所有值类型都是允许的，包括内联表。

内联表将出现在一行中。在花括号之间不允许换行，除非它们在值中是有效的。即便如此，强烈建议将内联表分割为多行。如果你发现自己被这种欲望所吸引，这意味着你应该使用标准的表格。

```toml
name = { first = "Tom", last = "Preston-Werner" }
point = { x = 1, y = 2 }
animal = { type.name = "pug" }
```

上面的内联表与下面的标准表相同:

```toml
[name]
first = "Tom"
last = "Preston-Werner"

[point]
x = 1
y = 2

[animal]
type.name = "pug"

```

表数组
---------------
最后一个还没有表示的类型是一个表数组。可以通过使用双括号中的表名来表示。每个具有相同双括号名称的表都是数组中的一个元素。按遇到的顺序插入表。没有键/值对的双括号表将被视为空表。

```toml
[[products]]
name = "Hammer"
sku = 738594937

[[products]]

[[products]]
name = "Nail"
sku = 284758393
color = "gray"
```

如果用JSON来描述这个信息，等同于如下格式：

```json
{
  "products": [
    { "name": "Hammer", "sku": 738594937 },
    { },
    { "name": "Nail", "sku": 284758393, "color": "gray" }
  ]
}
```

您还可以创建嵌套的表数组。只需在子表上使用相同的双括号语法。每个双括号子表都属于它上面最近定义的表元素。

```toml
[[fruit]]
  name = "apple"

  [fruit.physical]
    color = "red"
    shape = "round"

  [[fruit.variety]]
    name = "red delicious"

  [[fruit.variety]]
    name = "granny smith"

[[fruit]]
  name = "banana"

  [[fruit.variety]]
    name = "plantain"
```

如果用JSON来描述这个信息，等同于如下格式：

```json
{
  "fruit": [
    {
      "name": "apple",
      "physical": {
        "color": "red",
        "shape": "round"
      },
      "variety": [
        { "name": "red delicious" },
        { "name": "granny smith" }
      ]
    },
    {
      "name": "banana",
      "variety": [
        { "name": "plantain" }
      ]
    }
  ]
}
```
尝试向静态定义的数组追加，即使该数组是空的或兼容的类型，也必须在解析时产生错误。

```toml
# 无效的TOML文档
fruit = []

[[fruit]] # 不允许
```

试图定义与已建立数组同名的普通表，在解析时必须产生错误。

```
# INVALID TOML DOC
[[fruit]]
  name = "apple"

  [[fruit.variety]]
    name = "red delicious"

  # This table conflicts with the previous table
  [fruit.variety]
    name = "granny smith"
```

可以在适当的情况下使用内联表:

```toml
points = [ { x = 1, y = 2, z = 3 },
           { x = 7, y = 8, z = 9 },
           { x = 2, y = 4, z = 8 } ]
```

文件扩展名
------------------
TOML文件的扩展名应该使用`.toml`

MIME 类型
---------
在网络上传输TOML文件时，合适的MIME类型是`application/toml`.

与其他格式比较
-----------------------------
在某些方面，TOML与JSON非常相似:简单、指定良好，并且很容易映射到无处不在的数据类型。JSON对于主要由计算机程序读写的序列化数据非常有用。TOML与JSON的不同之处在于，它强调的是便于人类阅读和书写。注释是一个很好的例子:当数据从一个程序发送到另一个程序时，注释没有任何作用，但是对于可以手工编辑的配置文件来说，注释非常有用。

YAML格式是面向配置文件的，就像TOML一样。然而，对于许多目的来说，YAML是一个过于复杂的解决方案。TOML的目标是简单，而这个目标在YAML规范中并不明显http://www.yaml.org/spec/1.2/spec.html。

INI格式也经常用于配置文件。但是，这种格式没有标准化，并且通常不处理一个或两个以上的嵌套。

参与
------------
欢迎贡献

Wiki
----------------------------------------------------------------------

We have an [Official TOML Wiki](https://github.com/toml-lang/toml/wiki) that
catalogs the following:

* Projects using TOML
* Implementations
* Validators
* Language agnostic test suite for TOML decoders and encoders
* Editor support
* Encoders
* Converters

Please take a look if you'd like to view or add to that list. Thanks for being
a part of the TOML community!
