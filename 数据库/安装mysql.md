### windows下安装mysql
step1: 
    在[官网](https://dev.mysql.com/downloads/mysql/)下载并解压

step2:
    打开所解压文件,创建**my.ini**文件,该文件具体内容入下
```
[client]
# 设置mysql客户端默认字符集
default-character-set=utf8

[mysqld]
# 设置3306端口
port = 3306
# 设置mysql的安装目录
basedir=C:\\web\\mysql-8.0.11
# 设置 mysql数据库的数据的存放目录，MySQL 8+ 不需要以下配置，系统自己生成即可，否则有可能报错
# data存放目录与所解压文件在同一目录
# datadir=C:\\web\\sqldata
# 允许最大连接数
max_connections=20
# 服务端使用的字符集默认为8比特编码的latin1字符集
character-set-server=utf8
# 创建新表时将使用的默认存储引擎
default-storage-engine=INNODB
```

step3: 
    如果懒得将该文件的bin目录配置为环境变量,则打开bin目录,在此打开cmd,然后初始化数据库:
```
mysqld --initialize --console
```
执行完成后,会输出root用户的初始密码,如:
```
2018-04-20T02:35:05.464644Z 5 [Note] [MY-010454] [Server]
A temporary password is generated for root@localhost: APWCY5ws&hjQ
```
**APWCY5ws&hjQ8**就是初始密码,初次登陆会用到.
然后,进行使用命令进行安装:
```
mysqld install //注意,需要使用管理员身份运行cmd,此命令才会成功运行
```
step4:启动服务
```
net start mysql
```
step5:启动mysql,登陆并修改密码:
```
    输入mysql -u root -p,然后输入初始密码登陆
    最后使用 ALTER USER 'root'@'localhost' IDENTIFIED  BY 'newPassword';修改初始密码
```