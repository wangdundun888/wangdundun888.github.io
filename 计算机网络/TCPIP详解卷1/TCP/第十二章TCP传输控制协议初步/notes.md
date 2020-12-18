*TCP:*
Transmission  Control Protocol 传输控制协议

---

该章对TCP进行了一个概要讲述,对一些概念,例如ARQ,超时重传,分组窗口,滑动窗口,流量控制等进行了一个简单的描述,让我大致对TCP有了一个简单的了解.

---

*TCP头部*

|0                  15|16                    31|
---|---
源端口(16位)|目的端口(16位)
序列号(32位)|
确认号(32位)|
头部长度4位|保留(0)4位
<table>
    <tr align="center">
        <td>第0-15位</td>
        <td>第16-31位</td>
    </tr>
    <tr>
        <td>源端口(16位)</td>
        <td>目的端口(16位)</td>
    </tr>
    <tr align="center">
        <td colspan="20" >序列号(32位)</td>
    </tr>
    <tr align="center">
        <td colspan="2" >确认号(32位)</td>
    </tr>
</table>