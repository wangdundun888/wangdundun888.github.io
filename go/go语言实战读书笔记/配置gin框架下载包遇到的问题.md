# 配置gin框架下载包遇到的问题
## 问题
### 问题1:
    开始学习gin框架,在GitHub上准备根据教程运行第一个demo,使用命令 *go get -u github.com/gin-gonic/gin* 下载相应的包,
    下载过程中,一些包下载失败,运行项目也找不到某些包。
### 问题2:
    问题2是问题1解决之后衍生的,相应的包下载完成后运行项目,编译器提示在$GOOT/src和$GOPATH/src下找不到某个包,但是我确认已经下载
    完毕。
## 原因及解决方法
### 原因及解决方法1:
    问题1的出现是因为国外某些网站被墙了,所以包下载失败。经提示,我选择了换镜像源,Windows下换源:
        ① $env:GO111MODULE="on"  //开启Go Modules功能 
        ② $env:GOPROXY="https://mirrors.aliyun.com/goproxy/,direct" //配置镜像源,这里我选择了阿里云
    其他系统的看[这里](https://learnku.com/go/wikis/38122)
    配置完毕后,重新使用go get命令下载包,非常流畅!
### 原因及解决方法2:
    然后我在debug的过程中,在$GOPATH/pkg目录下发现我下载的包,并把编译器提示缺少的包复制到$GOPATH/src对应的目录下,运行项目,
    成功。
    然后在搜索问题的过程中,我发现了原因,因为在配置镜像源的过程中,将GO111MODULE="on",设置为on之后会使用1.13的mod(即 Go
    Module,官方包管理方式,我暂时还不了解)包特性,下载的包不会在src目录下,而是在$GOPATH/pkg目录下。
    [解决方法](https://blog.csdn.net/qq_43442524/article/details/104906475)
    
     *之后需要补充关于Go Module的知识* 