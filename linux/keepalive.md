# keepalive
提供一个虚拟IP供用户去使用和访问，当其中一台服务器出现宕机的情况时，另一台服务器能够在短时间内接替这个虚拟IP为用户提供服务
## VRRP
VRRP（Virtual Router RedundancyProtocol）：虚拟路由器冗余协议，解决静态路由单点故障问题的，它能够保证当个别节点宕机时，整个网络可以不间断地运行
## vip切换原理
1. 正常情况下，用户通过虚拟IP是直接访问到Keepalived-Master的（没有成为Master的就是Backup）；
2. 成为Master的Keepalived，会每秒向所有的Backup发送VRRP包，通告自己是主，且运行正常；
3. 当Master因为网络原因或者是别的原因导致与集群断开之后，Backup会在3.6秒左右（以优先级100为例，计算公式为3 × 1 + 256 × （256 - 100））认定Master宕机；
4. 如果是多播的情况下，Master宕机，那么剩余的Backup要通过选举产生新的Master；如果是单播，则由剩下的Bakcup直接作为新的Master
## 选举机制
keepalived中优先级高的节点为MASTER。MASTER其中一个职责就是响应VIP的arp包，将VIP和mac地址映射关系告诉局域网内其他主机，同时，它还会以多播的形式（默认目的地址224.0.0.18）向局域网中发送VRRP通告，告知自己的优先级。网络中的所有BACKUP节点只负责处理MASTER发出的多播包，在抢占模式下，当发现MASTER的优先级没自己高，或者没收到MASTER的VRRP通告时，BACKUP将自己切换到MASTER状态，然后做MASTER该做的事：响应arp包和发送VRRP通告

当一个激活了VRRP的接口Up之后，如果接口的VRRP优先级为255，那么其VRRP状态将直接从
Initialize切换到Master，而如果接口的VRRP优先级不为255，则首先切换到Backup状态，然后再
视竞争结果决定是否能够切换到Master状态

如果在同一个广播域的同一个VRRP组内出现了两台Master路由器，那么它们收到对方发送的VRRP通
告报文之后，将比较自己与对方的优先级，优先级的值更大的设备胜出，继续保持Master状态，而
竞争失败的路由器则切换到Backup状态。如果这两台Master路由器的优先级相等，那么接口IP地址
更大的路由器接口将会保持Master状态，而另一台设备则切换到Backup状态

## 抢占模式
针对主机崩溃，集群已经重新选出新的主机，并且原来的主机重新上线后重新争夺主机、成为主机的情况。这种情况比较适合需要崩溃的主机重新上线做主机的情况。也就是本模式是某主机节点在崩溃后，集群新选出了主机。但是崩溃的节点恢复后，又触发了选举，最终原节点再次成为主机的过程
## 非抢占模式
* 节点无主从之分，皆为backup
* 配置文件中添加nopreempt标识

针对主机崩溃，集群已经重新选出新的主机，并且原来的主机重新上线后并不争夺主机的情况。这种模式适合那些倾向于认为崩溃的主机即便上线还是会出现崩溃的场景
核心思想：将所有节点的优先级（priority）值设为相同，当两个节点的优先级相同时，以节点发送VRRP通告的IP作为比较对象，IP较大者为MASTER
当主机崩溃时，所有从机中IP地址最大的那个节点一定会成为主机，并且原来的主机
上线后不会再次触发选举，原来的主机也就成为了从机

##常见问题
### 脑裂
两台keepalived高可用服务器在指定时间内，无法检测到对方存活心跳信息，从而导致互相抢占对方的资源和服务所有权，然而此时两台高可用服务器有都还存活
可能出现的原因：

1. 服务器网线松动等网络故障；
2. 服务器硬件故障发生损坏现象而崩溃；
3. 主备都开启了firewalld 防火墙。
4. 在Keepalived+nginx 架构中，当Nginx宕机，会导致用户请求失败，但是keepalived不会进行切换，

## 安装配置
### 安装
```shell
apt install keepalived
cat > /etc/keepalived/keepalived.conf <<EOF
systemctl start keepalived.service
```
### 配置
```shell
! Configuration File for keepalived
global_defs {
   router_id node_A
   vrrp_iptables
   vrrp_version 3
}

vrrp_instance IP1 {
    state BACKUP  
    interface eth0
    
    #让master 和backup在同一个虚拟路由里，id 号必须相同
    virtual_router_id 10
    
    #优先级,谁的优先级高谁就是master，0~255
    priority 100
    
    #设置 MASTER 与 BACKUP 负载均衡之间同步即主备间通告时间检查的时间间隔, 单位为秒，默认 1s
    advert_int 0.1
    authentication {
        auth_type PASS  #认证
        auth_pass lSXDYtK0 #密码
    }
    
    unicast_src_ip 10.176.34.116
    unicast_peer {
        10.176.34.232
    }
    
    #vip
    virtual_ipaddress {
        10.176.34.253
    }
}
```