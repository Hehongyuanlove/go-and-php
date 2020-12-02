



> 知乎交流



#### 测试环境

```

/**
* win10   E5-2678 v3  32G DDR4
*
* PHPStudy 集成环境
* PHP 7.3.4 (cli) (built: Apr  2 2019 21:57:22) ( NTS MSVC15 (Visual C++ 2017) x64 )
* Copyright (c) 1997-2018 The PHP Group
* Zend Engine v3.3.4, Copyright (c) 1998-2018 Zend Technologies
*
* go version go1.15.5 windows/amd64
*
*/

```



#### 测试说明



- hash-base.php  截图代码

- hash-base.go 截图代码

- hash-changed.go  go 自己优化过代码（刚学，有问题请指正）



| 代码 循环 1E6   | 时间（ms） | 内存（B）   |
| --------------- | ---------- | ----------- |
| hash-base.php   | 297        | 9033,6160   |
| hash-base.go    | 350.0016   | 1,3828,9792 |
| hash-changed.go | 135.1354   | 1,0627,1344 |



| 代码 循环 2E6   | 时间        | 内存（B）   |
| --------------- | ----------- | ----------- |
| hash-base.php   | 661 ms      | 1,8027,9200 |
| hash-base.go    | 1.0810163 s | 3,4213,1952 |
| hash-changed.go | 248.9743 ms | 1,5427,8416 |



| 代码 循环 3E6   | 时间（ms） | 内存（B）   |
| --------------- | ---------- | ----------- |
| hash-base.php   | 未测       | 未测        |
| hash-base.go    | 未测       | 未测        |
| hash-changed.go | 362.0003   | 2,0228,1328 |

