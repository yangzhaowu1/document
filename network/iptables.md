# iptables
## Netfilter框架
* 数据包处理框架
* 功能：数据包过滤、修改、SNAT、DNAT等
* 实现：netfilter在内核协议栈的不同位置实现了5个hook点，其它内核模块(比如ip_tables)可以向这些hook点注册处理函数，这样当数据包经过这些hook点时，其上注册的处理函数就被依次调用

![netfilter框架](vx_images/217654418232181.png =1180x)

* bridge level：ebtables的表和链，只工作在链路层，处理的是以太网帧(比如修改源目mac地址)
* network level：iptables的表和链,只处理IP数据包
* conntrack：connection tracking，netfilter提供的连接跟踪机制，此机制允许内核”审查”通过此处的所有网络数据包，并能识别出此数据包属于哪个网络连接，使内核能够跟踪并记录通过此处的所有网络连接及其状态
* bridge check：数据包从某个网络接口进入(ingress),检查此接口是否属于某个Bridge的port，如果是就会进入Bridge代码处理逻辑(下方蓝色区域bridge level), 否则就会送入网络层Network Layer处理
* bridging decision类似普通二层交换机的查表转发功能，根据数据包目的MAC地址判断此数据包是转发还是交给上层处理
    1. 包目的MAC为Bridge本身MAC地址(当br0设置有IP地址)，从MAC地址这一层来看，收到发往主机自身的数据包，交给上层协议栈(D –> J)
    2. 广播包，转发到Bridge上的所有接口(br0,tap0,tap1,tap…)
    3. 单播&&存在于MAC端口映射表，查表直接转发到对应接口(比如 D –> E)
    4. 单播&&不存在于MAC端口映射表，泛洪到Bridge连接的所有接口(br0,tap0,tap1,tap…)
    5. 数据包目的地址接口不是网桥接口，桥不处理，交给上层协议栈(D –> J)
* routing decision：路由选择，根据系统路由表(ip route查看), 决定数据包是forward，还是交给本地处理



## conntrack
iptables实现状态匹配(-m state)以及NAT的基础，它由单独的内核模块nf_conntrack实现

## NAT

