#! /bin/bash

# 变量
a=1
b=2
# 使用一个定义过的变量，只要在变量名前面加美元符号即可
echo $a
echo ${b} #变量名外面的花括号是可选的，加不加都行，加花括号是为了帮助解释器识别变量的边界

# 字符串

# 字符串可以用单引号，也可以用双引号，也可以不用引号

str='this is a string'

echo $str



date
date "+%Y-%m-%d %H:%M:%S"