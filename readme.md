写了几个月才想起来为什么不写一个工具把document里面的api参数扒拉下来免得一个个复制下来我是说我是弱智。

大约是
```txt
` document -> params' attributes -+-> swagger comments
`                                 +-> gin params get
`                                 +-> mastodon api entrance
```

## branches

main 主分支，先写js的好了，参考查漏补缺那个
py的re有点大便是说。

## js

总之也没想好是用js写还是用python写。
为什么不全都用这个方法呢。

能筛api的信息了
之后需要稍加处理，生成代码文件。
大概是`controller/mastodon`里的内容和`core/mastodon`里的内容(只有框架)
要点
- swagger
- gin get params
- pass to function
- 怎么给函数起名啊
  - 就变成大写好了。

## python

如果js写不出来就用python写写看（写得出就咕咕了，还是以实用为目的而不是学习，姑且python学过了）

# log

初始化npm
```sh
npm init
```
么有`.`啊

TODO: 调通 ts 的测试环境
## tsc
初始化
```sh
tsc --init
```

## install modules
会找不到类型的，先安装ts的类型
```sh
npm i -D @types/node
```
再安装module
```sh
npm i fs
```


# settings list



