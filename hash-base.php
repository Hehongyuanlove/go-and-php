<?php
$t1 = microtime(true);
$arr = array();
for($i = 2147483647;$i<2150483647;$i++){
    $value = time();
    $key = $i.'_'.$value;
    $arr[$key] = $value;
}

$t2 = microtime(true);
echo round(($t2-$t1)*1000) . "ms\n";
echo memory_get_peak_usage() ."B \n";
echo time();

/**
* PHPStudy 集成环境
* PHP 7.3.4 (cli) (built: Apr  2 2019 21:57:22) ( NTS MSVC15 (Visual C++ 2017) x64 )
* Copyright (c) 1997-2018 The PHP Group
* Zend Engine v3.3.4, Copyright (c) 1998-2018 Zend Technologies
*/

/**
* 1e6
* 297 ms
* 9033,6160 B
*/

/**
* 2e6
* 661ms
* 1,8027,9200 B
*/