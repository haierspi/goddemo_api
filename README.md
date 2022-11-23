Gin Gorm 演示项目
===


### DB模型自动生成工具

`gormgen` 介绍
通过脚本生成基于 gorm  的 CURD 方法。
参考： [https://github.com/MohamedBassem/gormgen](https://github.com/MohamedBassem/gormgen)

在根目录下执行脚本： ./scripts/gormgen.sh addr user pass name tables
- addr：数据库地址，例如：127.0.0.1:3306
- user：账号，例如：root
- pass：密码，例如：root
- name：数据库名称，例如：go\_gin\_api
- tables：表名，默认为 \*，多个表名可用“,”分割，例如：user\_demo
- prefix: 忽略的表名前缀
- savedir:  保存目录


```shell

```

### 打印工具

go get -u github.com/gookit/goutil 
