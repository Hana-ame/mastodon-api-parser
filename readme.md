好,我已经完全不记得进度了
目前大概是

- [ ] 从api的json搓出controll下的go文件
  - [ ] swagger
    - [ ] 没测试啊

好,放弃下面的了。逻辑写controller里面



放弃了
  - [ ] ~~gin~~
    - [ ] ~~如何取得param~~
  - [ ] ~~api function~~
    - [ ] ~~如何传入param~~

基本形态

controller
```go
package mastodon

import (
  mastodon "mastodon/[folder]"
)


// swag
// ..
func FuncNmae(c *gin.Context) {
  
  // get params
  ...

  resp, err := mastodon.FuncName(..params)

  // error handle
  ....

  return
}
```


写了几个月才想起来为什么不写一个工具把document里面的api参数扒拉下来免得一个个复制下来我是说我是弱智。

大约是
```txt
` document -> params' attributes -+-> swagger comments
`                                 +-> gin params get
`                                 +-> mastodon api entrance
```

**请关注**
[TODO](#todo)
[known issue](#known-issue)

请同时关注 readme 中的 gin.md

很有必要出swagger的api嘛。

## **TODO**
 


~~接着复制document~~
~~去复制entities~~
~~复制完了再去生成代码~~


看了swaggo源码，mf啥来着，accept要改成
否则的话是用www-encoding post过来的

## file struct
- `html.ts`
  - api documents (not saved) to divided api documents in `text`
- `index.ts`
  - documents in `text` to raw json in `json`
- `entities.ts`
  - documents in `entities/documents` to raw json in `entities/raw`
  - raw json in `entities/raw` to cooked json `entities/cooked`
    - TODO: maybe not `map[string]any` but other struct
    - TODO: update special keys
- `api.ts`
  - TODO: raw json to cooked son
- `db.ts`
  - general db settings (CURD) of entities.
  - sould handwrite some keys
- `swagger.ts`
  - generate swagger infos
- TODO
  - `gin.ts`
    - generate gin infos
  - `core.ts`
    - generate `core/mastodon` entrances.
    - please leave a link on its comment.

## dev log

### [swagger](api.ts#swagger)

从`index.ts`里面复制过来了.

写了下swagger,现在开始degug

- [ ] params, type不在其中的情况


### html.tx
```sh
npm run html
```

### index.ts
```sh
npm run test
```

### entities.ts
```sh
npm run entities
```
改回 Hash ，算了不改了

生成`.go`文件？
`.go`文件要什么
```go
type EntityName struct {
  Key type `json:"key"`
}
```
要记得自己建文件夹。
TODO: 补一下 `Status::Xxx` 的部分
TODO: 需要db关照一下many to many的部分，primarykey的部分，需要做index的部分

~~### db.ts~~
不用新文件，还在entities里面弄
~~```sh~~
~~npm run db~~
~~```~~
- methods:
  - CRUD
    - create one
    - read one
    - update one
    - delete one

## note

需要筛出
- swagger 需要的
- 从gin parse变量需要的
- 传参需要的(golang 类型)
- 返回需要的entities类型
- gin router

TODO:看到的话先去写entities的东西好了.

entities做好了cooked json
接下来请做好golang的表示
hash可能要手动修
请以尽快为目的写出代码，不必纠结正确与否，做了再说。

api部分可能需要从raw json思考一下如何制作cooked json

PS：一些工具在`LineReader.ts`里面

## known issue

1. [ ] Admin_Account类型不统一，有下划线和没下划线
2. [ ] ::会有错误。
   1. [ ] 不能多段用`.`吧

# branches

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

v1/v2被冲了

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



