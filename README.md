# 说明

这个仓库是临时起意建的，也就是把之前自己实现的一些高级数据结构整合在一起。

## trie - 24.04

前缀树实现，包括普通前缀树(trie)和基数前缀树(radix-trie)。

这俩的作用主要是在 http 的 url 前缀匹配上面，像 gin 就使用了 radix-trie。golang 的 内存模型里面，heap 也借助了 radix-trie 来快速查找空闲页。

## bitmap && bloom filter - 24.04

bitmap 即位图，通过 比特位来标识数据的存在情况(更宽泛的说，元素只有两种状态，元素集是连续的，就可以用 bitmap 来维护)。在 golang 内存模型就使用了 bitmap 来标识 page 的 空闲状态，linux 同样使用 bitmap.

bloom filter 则是对 bitmap 的一种升级，通过对输入的元素做 k 次 hash，来降低冲突率，更大限度的利用空间。

还有另一种 filter 叫 布谷鸟过滤器(cuckoo filter)，在布隆过滤器的基础上，还支持删除元素这一操作，有兴趣可以了解。

## B+ tree - 24.05

B+ 树 实现，主要使用是在数据库底层，用于维护和保存数据。mysql 的 索引就是通过 b+ 树实现的。

## skiplist - 24.05

跳表，是对于单向链表的一种改进，也是一种有序表，其大多数操作(增删改查)都是O(logN)，还支持范围查询和近似查询。



