# MQTT test

## install 

see `install.sh` and `test.sh`

## golang.org/x

```bash
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net.git
```
其它 golang.org/x 下的包获取皆可使用该方法。

例如，很多go的软件在编译时都要使用tools里面的内容，使用下面方法获取：

进入上面的x目录下，输入：

git clone https://github.com/golang/tools.git
注意，一定要保持与go get获取的目录结构是一致的，否则库就找不到了。
