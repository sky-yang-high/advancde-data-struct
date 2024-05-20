# Bloom Filter - 布隆过滤器

参考：[布隆过滤器技术原理及应用实战](https://mp.weixin.qq.com/s?__biz=MzkxMjQzMjA0OQ==&mid=2247484535&idx=1&sn=bf906823a8ac5efd1a7c567622ff8968)

bloom filter是一种数据结构，和bitmap相似，都是处理集合中是否存在某元素这类问题催生的。

在说bloom filter之前，下面先说说bitmap

## bitmap

很简单，就是用数组来标识该位是否存在。

而它和数组的区别也就在于bit，即它用位来标识，而不是一个字节(甚至一个数据宽度)，从而降低内存占用。

因此其set,clear,和get操作都是使用位运算来实现：

- 记输入数为num，则在bitmap中的字节为num/8，在该字节中的位为num%8

- set: 使用 |= 1<<(num%8) 来把该位置1
- clear: 使用 &^= 1<<(num%8) 来把该位置零
- get: 使用 == 1<<(num%8) 来判断

更进一步，对于字符串类型，我们则可以使用hash，先把该字符串hash化，然后把hash值对 bitmap的总size(即占用字节数\*8) 取模即可。

## bloom filter

上面的bitmap由于只有1次hash，因此如果数据量大时，很可能发生碰撞。

另外，我们可以发现：

- 若一个字符串存在，bitmap 不会 将其误判为不存在
- 若一个字符串不存在，bitmap 可能会 将其误判为存在

而bloom filter则是增加hash函数的个数，来尽量减少碰撞的概率。

具体怎么做：

- 对于一个字符串，分别对k个hash函数取hash值
- 把bitmap中的k个hash值对应的位置均置1
- 判断该hash值是否存在，也就需要同时判断这k个位置是否都为1

但是有一些缺点：

- 由于hash次数更多，比bitmap更慢
- 极端情况，所有位置都被置为1，那么任意字符串都会被误判为存在(通过及时更换一个新的bloom filter解决)
- 无法删除数据，由于1个位置上还可能同时被其他字符串置1，因此不能通过把对应的k个位置置零来删除数据(可以通过添加计数来解决(解决了吗?))

有另一款过滤器，叫布谷鸟过滤器，能够在一定程度上支持 bitmap 中的数据删除操作，链接见：[seiflotfy/cuckoofilter: Cuckoo Filter: Practically Better Than Bloom ](https://github.com/seiflotfy/cuckoofilter)

### 误判率分析

参数如下：

- bitmap的长度(按位)：m
- hash函数个数：k
- bitmap中已输入元素个数：n

这里不把推演过程展示出来了，就是简单的概率计算和微积分等价无穷小的计算，最终得到：

> 当 m->∞ 时，误判概率可以简化表示为—— p = [1-e^(-k·n/m)^]^k^.

以及，通过对上面的函数求导，得到：

> 当 k·n/m=ln2 时，误判概率 f(k) 取到极小值.

这个网站，提供了直观的可视化bloom filter调优界面，可以尝试[bloom filter](https://hur.st/bloomfilter)

根据上面的式子，k = m/n*ln2，即若m,n量级相当时，k通常只需要常数级，2/3就足够了。

### hash算法选取

由于我们不需要加密算法那么严格的安全性，因此不选择像md5, sha1这些加密hash算法。

选择一款性能相对比较高的非加密算法即可，例如murmur3, cityhash。

### 代码实现

代码实现上，很简单，和bitmap的实现没有太大的差别。

- set: 先把输入的字符串经过hash得到k个hash值(我们的实现是把本次hash值作为输入再次hash，重复K次来模拟K个hash函数)，然后把对应位置都置1即可
- check: 同样，经过hash得到k个hash值，若这k个位置都为1则视为存在，否则视为不存在





















