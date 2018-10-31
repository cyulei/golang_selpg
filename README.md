# Go版本selpg

## 简介
该程序使用golang开发[开发Linux命令行实用程序][1]中的selpg

## 开发环境
- CentOS7
- go 1.9.4 linux/amd64

## 使用
### 下载
> go get github.com/cyulei/Golang_selpg 
### 参数
> selpg -s=Number -e=Number [options] [filename]


-    -s : Start page.
-   -e : End page.
-   -l : Determine the number of lines per page and default is 20.
-    -f : Determine the type and the way to be seprated.
-   -d : Determine the destination of output.
-   [filename] : Read input from this file.
-   If filename is not given, read input from stdin. and Ctrl+D to cut out.
### 示例

> selpg -s1 -e1 -l10 test.txt

> 将test.txt第1页的前10行打印到屏幕上
```
This is the line 1
This is the line 2
This is the line 3
This is the line 4
This is the line 5
This is the line 6
This is the line 7
This is the line 8
This is the line 9
This is the line 10
```

> selpg -s1 -e1 < test.txt

> selpg读取标准输入，而标准输入被shell/内核重定向为来自test.txt而不是显示命名的文件名参数。输入的第1页被写至屏幕。1页有20行。
```
This is the line 1
This is the line 2
This is the line 3
This is the line 4
This is the line 5
This is the line 6
This is the line 7
This is the line 8
This is the line 9
This is the line 10
This is the line 11
This is the line 12
This is the line 13
This is the line 14
This is the line 15
This is the line 16
This is the line 17
This is the line 18
This is the line 19
This is the line 20
```
> selpg -s2 -e2 test.txt > out.txt

> 将第2页写入out.txt中
```
out.txt：
This is the line 21
This is the line 22
This is the line 23
This is the line 24
This is the line 25
This is the line 26
This is the line 27
This is the line 28
This is the line 29
This is the line 30
```

>  selpg -s1 -e0 test.txt 2>error.txt

> 将错误消息写入test.txt中

```
test.txt:
: invalid end page 0
```

> selpg -s3 -e5 -l4 test.txt >out.txt 2>error.txt

> 将第3到5页写入out.txt中，标准错误将被写入error.txt中
```
out.txt:
This is the line 9
This is the line 10
This is the line 11
This is the line 12
This is the line 13
This is the line 14
This is the line 15
This is the line 16
This is the line 17
This is the line 18
This is the line 19
This is the line 20
```

> selpg -s1 -e1 -f test.txt

> 假定页由换页符界定，第1页被打印到屏幕（因为测试文档中没有换页符，所以文档中所有内容打印出）

```
This is the line 1
This is the line 2
This is the line 3
This is the line 4
This is the line 5
This is the line 6
This is the line 7
This is the line 8
This is the line 9
This is the line 10
This is the line 11
This is the line 12
This is the line 13
This is the line 14
This is the line 15
This is the line 16
This is the line 17
This is the line 18
This is the line 19
This is the line 20
This is the line 21
This is the line 22
This is the line 23
This is the line 24
This is the line 25
This is the line 26
This is the line 27
This is the line 28
This is the line 29
This is the line 30
```

> selpg -s1 -e1 -l5 -dcat test.txt

> 将第1页发送之命令cat，将前5行打印在屏幕上
```
This is the line 1
This is the line 2
This is the line 3
This is the line 4
This is the line 5
```


  [1]: https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html
