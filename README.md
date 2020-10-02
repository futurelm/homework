问题：
-------
有一个 100GB 的文件，里面内容是文本， 要求：
+ 假定：每行都是`,`分隔的字符串
+ 找出第一个不重复的词
+ 只允许扫一遍原文件
+ 尽量少的 IO
+ 内存限制 16G

思路：
--------
 1. 创建一个长度为`n`的有缓冲channel，缓冲长度`m`数组，开启一个协程，按行扫描文件，把每个字符串c存入根
据字符串哈希值对`n`取余得到的channel中。另一个协程遍历channel数组，分别从每个channel中过滤出非重复的字符串，统
 一写入另一个channel `final channel`,最后在`final channel`中就是所有不重复的词，然后挑选索引最小的返回
 2. 可以调整channel数组长度 `n`，channel 缓冲区大小 `m`的值, 在读文件使用scanner时指定buffer size。
 限制内存使用
 
 使用
 -------
 ```bash                             
  ## 创建一个随机字符串文件，words.txt
  go test -v createWorlds_test.go createWorlds.go  
  ## 扫描文件，输出结果
  go test -v find_non_repeated_string_test.go find_non_repeated_string.go 
 ```

