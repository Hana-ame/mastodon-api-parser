import fs from 'fs'
import path from 'path'
import { trimPrefix } from './LineReader';

// step 1. traverse the folder.
function traverse(currentPath: string) {
  // console.log('traverse')
  if (currentPath == 'text\\streaming') return
  // 读取当前文件夹中的内容
  const files = fs.readdirSync(currentPath);

  // 遍历每个文件/文件夹
  files.forEach((file) => {
    const filePath = path.join(currentPath, file);
    const stats = fs.statSync(filePath);

    if (stats.isFile()) {
      // 处理文件
      console.log('File:', filePath);
      handleFile(currentPath, file)
    } else if (stats.isDirectory()) {
      // 处理子文件夹
      console.log('Directory:', filePath);
      // 递归遍历子文件夹
      traverse(filePath);
    }
  });
}

// types for step 2

type ParamType = "formdata" | "path" | "header" | "query" | "unknown"

type ParamRaw = {
  name : string
  paramType : ParamType
  // required : boolean // add later
  // dataType : string // add later
  comment : string
}

type APIRaw = {
  summary : string
  method : string
  router : string
  description : string
  returns : string
  oauth : string
  request : ParamRaw[]
}

function summaryToParscal(summary: string) {
  const matches = summary.match(/\w+/g); 
  if (matches == null) {
    console.log(summary)
    console.log(matches)
    throw new Error('Variable is undefined or null');
  }
  const t = matches.join("_")
  return t
}

// swagger的逻辑在这里
class Swagger{
  raw: APIRaw
  group : string
  constructor(raw: APIRaw, tag : string) {
    this.raw = raw
    this.group = tag
  }
  get title() {
    const summary = this.raw.summary
    const t = summaryToParscal(summary)
    return Swagger.title(t)
  }
  get summary() {
    return Swagger.summary(this.raw.summary)
  }
  get description() {
    return Swagger.description(this.raw.description.trim())
  }
  get tag() {
    return Swagger.tag('mastodon', this.group)
  }
  get produce() {
    return Swagger.produce('json')
  }
  get params() {
    return this.raw.request.map(param => {
      let name = param.name
      name = trimPrefix(name, ':')
      const type = param.paramType
      let datatype = 'string'
      if (type !== 'header') {
        let typeComment = param.comment.split('.')[0]
        typeComment = trimPrefix(typeComment, "REQUIRED ")
        typeComment = trimPrefix(typeComment, "Array of ")
        switch (typeComment) {
          case "String":
            datatype = 'string'
            break
          case "Boolean":
            datatype = 'boolean'
            break
          case "Integer":
            datatype = 'int'
            break
          default:
            console.log(typeComment) // 其他的类型
        }
      }
      const isArray =  param.comment.includes("Array of")
      const isRequired = param.comment.includes("REQUIRED")
      const comment = param.comment
      return Swagger.param(name, type, datatype, isRequired, comment)
    })
  }
  get success() {
    const isArray = this.raw.returns.includes('Array')
    let type = trimPrefix(this.raw.returns, "Array of ").trim()
    const arrayPrefix = isArray ? '[]' : ''
    return Swagger.success(200, isArray,  `entities.${type}`)
  }
  get router() {
    const router = this.raw.router
    const method = this.raw.method.toLocaleLowerCase()
    return Swagger.router(router, method)
  }
  static title(title : string ){
    return `// ${title} godoc`
  }
  static summary(summary: string){
    return `// @Summary ${summary}`
  }
  static description(description: string){ // with \n s
    const arr = description.split('\n')
    let res = ''
    for (const line of arr) {
      res += '// @Description ' + line
      res += '\n'
    }
    return res.slice(0, -1);
  }
  static tag(...tag : string[]) { 
    return `// @Tags ${tag.join(',')}`
  }
  static produce(mime: string) {
    return `// @Produce ${mime}`
  }
  static param(name: string, type: string, dataType: string, isRequired: boolean, comment: string ) {
    return `// @Param ${name} ${type} ${dataType} ${isRequired} "${comment}"`
  }
  static success(code: number, isArray: boolean, type: string) {
    return `// @Success ${code} ${isArray? 'array' : 'object'} ${type}`
  }
  static router(router: string, method: string) {
    return `// @Router ${router} [${method}]`
  }
}

// gin的逻辑在这里。
class Gin{
  raw: APIRaw
  group : string
  constructor(raw: APIRaw, tag : string) {
    this.raw = raw
    this.group = tag
  }
  get parscalSummary() {
    const summary = this.raw.summary
    const t = summaryToParscal(summary)
    return t
  }

  get firstline() {
    return `func ${this.parscalSummary}(c *gin.Context){`
  }
  get endOfBody() {
    return `c.JSON(http.StatusNotImplemented, gin.H{"error":"Not Implemented"})` 
    
  }
  get lastline() {
    return '}'
  }
}

// 在这里
class API {
  raw : APIRaw
  tag : string
  constructor(raw : APIRaw, tag : string) {
    this.raw = raw
    this.tag = tag
  }

  get prevText() {
    let t = ''
    t += 'package ' + this.tag + '\n'
    // t += 'import ('
    // t += `  api "github.com/moonchan-xyz/moonchan/back/mastodon/${this.tag}"`
    // t += `  "github.com/moonchan-xyz/moonchan/back/controller"`
    // t += ')'

    return t
  }
  get swaggerText() {
    const lines : string[] = []
    const swagger = new Swagger(this.raw, this.tag)
    lines.push(swagger.title)
    lines.push(swagger.summary)
    lines.push(swagger.description)
    lines.push(swagger.tag)
    lines.push(swagger.produce)
    lines.push(...swagger.params)
    lines.push(swagger.success)
    lines.push(swagger.router)

    return lines.join('\n')
  }
  get ginText() {
    const lines : string[] = []
    const gin = new Gin(this.raw, this.tag)
    lines.push(gin.firstline)
    lines.push(gin.endOfBody)
    lines.push(gin.lastline)

    return lines.join("\n")
  }

}

// tool function for step 2.

function readFile(path : string) {
  try{
    const fileContents = fs.readFileSync(path, 'utf8')
    return fileContents
  } catch(e) {
    console.log(e)
  }
}

function readRawAPI(fn: string) {
  const content = readFile(fn)
  if (content) {
    const json = JSON.parse(content) as APIRaw
    return json
  }
  return null
}

function readAPI() {
  
}

// step 2. handle it 
function handleFile(currentPath: string, file: string){
  const filePath = path.join(currentPath, file);
  const tag = trimPrefix(currentPath, 'json\\')
  // read APIRaw
  const raw = readRawAPI(filePath)
  // here is the api
  if (raw == null) {
    console.log("error api is null")
    return 
  }
  const cooked = new API(raw, tag)
  // save
  saveToGofiles(tag, cooked)
}

function saveToGofiles(tag: string, cooked: API) {
  let newfilePath = path.join('controller', tag)
  if (!fs.existsSync(newfilePath)) {
    fs.mkdirSync(newfilePath, { recursive: true });
  }
  newfilePath = path.join(newfilePath,  summaryToParscal(cooked.raw.summary) + '.go')

  const fileContent = cooked.prevText + '\n' + cooked.swaggerText + '\n' + cooked.ginText

  fs.writeFileSync(newfilePath, fileContent, 'utf-8')  
}
// run
traverse("json")