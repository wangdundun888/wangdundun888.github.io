*TCP:*
Transmission  Control Protocol 传输控制协议

---

该章对TCP进行了一个概要讲述,对一些概念,例如ARQ,超时重传,分组窗口,滑动窗口,流量控制等进行了一个简单的描述,让我大致对TCP有了一个简单的了解.

---

*TCP头部*

<table>
    <tr align="center" >
        <td colspan="16">源端口(16位)</td>
        <td colspan="16">目的端口(16位)</td>
    </tr>
    <tr align="center">
        <td colspan="32" >序列号(32位)</td>
    </tr>
    <tr align="center">
        <td colspan="32" >确认号(32位)</td>
    </tr>
    <tr>
        <td colspan="4">头部长度(4位)</td>
        <td colspan="4">保留(0)(4位)</td>
        <td colspan="1">CWR(1位)</td>
        <td colspan="1">ECE(1位)</td>
        <td colspan="1">URG(1位)</td>
        <td colspan="1">ACK(1位)</td>
        <td colspan="1">PSH(1位)</td>
        <td colspan="1">RST(1位)</td>
        <td colspan="1">SYN(1位)</td>
        <td colspan="1">FIN(1位)</td>
        <td>窗口大小(16位)</td>
    </tr>
    <tr align="center" >
        <td colspan="16">TCP校验和(16位)</td>
        <td colspan="16">紧急指针(16位)</td>
    </tr>
    <tr align="center">
        <td colspan="32" >选项(可变量)</td>
    </tr>
</table>

