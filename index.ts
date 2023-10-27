import fs from 'fs'
import path from 'path'
// api

// types
const lineFormData = "Form data parameters"
const lineHeader = "Headers"
const lineQuery = "Query parameters"
const linePath = "Path parameters"
type ParamType = "formData" | "path" | "header" | "query" | "unknown"
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

// traverse the folder
function traverseText(currentPath: string) {
  if (currentPath == 'text\\streaming') return
  // 读取当前文件夹中的内容
  const files = fs.readdirSync(currentPath);

  // 遍历每个文件/文件夹
  files.forEach((file) => {
    if (!file.endsWith('.txt')) return;
    const filePath = path.join(currentPath, file);
    const stats = fs.statSync(filePath);

    if (stats.isFile()) {
      // 处理文件
      console.log('File:', filePath);
      textToAPIraw(currentPath, file)
    } else if (stats.isDirectory()) {
      // 处理子文件夹
      console.log('Directory:', filePath);
      // 递归遍历子文件夹
      traverseText(filePath);
    }
  });
}

// read txt file synchronously
function readFile(path : string) {
  try{
    const fileContents = fs.readFileSync(path, 'utf8')
    return fileContents
  } catch(e) {
    console.log(e)
  }
}

class LineReader {
  private idx = 0;
  private data = "";
  private lines:string[] = []
  constructor(data: string){
    this.data = data
    this.lines = data.split("\n")
  }
  readline(idx?:number) {
    if (idx) {
      if (idx<this.lines.length)
        return this.lines[idx].trim()
      return ""
    }
    if (this.idx<this.lines.length) {
      const r = this.lines[this.idx]
      this.idx++
      return r.trim()
    }
    return ""
  }
}
// text to APIRaw
function parseAPIRaw(doc: string) {
  const reader = new LineReader(doc)
  let line = ""
  let arr = [""]
  const summary = reader.readline()

  line = reader.readline()
  arr = line.split(" ")
  const method = arr[0]
  const router = arr[1]

  let description = ""
  line = reader.readline()
  while (!line.startsWith("Returns: ")) {
    description += line + "\n"
    line = reader.readline()
  }
  
  let returns = ""
  // line = reader.readline()
  if (line.startsWith("Returns: ")) {
    returns = line.split(" ").slice(1).join(" ")
  } else console.log("error on Returns")

  let oauth = ""
  line = reader.readline()
  if (line.startsWith("OAuth: ")) {
    oauth = line.split(" ")[1]
  } else console.log("error on OAuth")

  while (line !== "") line = reader.readline()

  let request : ParamRaw[] = []
  line = reader.readline()
  if (line == 'Request') {
    let name = ""
    let paramType : ParamType = "unknown"
    // let required = false
    // let dataType = "unknown"
    let comment = "dummy"
    line = reader.readline()
    while (line != "Response") {
      let skip = true
      if (line == lineFormData) paramType = "formData"
      else if (line == lineHeader) paramType = "header"
      else if (line == linePath) paramType = "path"
      else if (line == lineQuery) paramType = "query"
      else skip = false
      if (skip) {
        line = reader.readline()
        continue
      }
      if (line.includes('.')) {
        comment = line
        request.push({
          name,
          paramType,
          comment,
        })
      } else {
        name = line
      }      
      line = reader.readline()
    }
  }
  
  return {
    summary ,
    method ,
    router ,
    description ,
    returns ,
    oauth ,
    request ,
  } as APIRaw
}

function textToAPIraw(dir :string, file: string) {
    const filePath = path.join(dir, file);
    const doc = readFile(filePath) ?? "error: "
  if (doc.startsWith("error: ")) {
    console.log(doc, filePath)
    return 
  }
  const api = parseAPIRaw(doc)
  const folderPath = 'json/' + dir.split('\\').slice(1).join('/')
  const newfilePath = folderPath + '/' + file.split('.').slice(0,-1).join('.') + '.json'
  // console.log(folderPath)
  // console.log(newfilePath)
  if (!fs.existsSync(folderPath)) {
    fs.mkdirSync(folderPath, { recursive: true });
  }
  const fileContent = JSON.stringify(api, null, "  ")
  fs.writeFileSync(newfilePath, fileContent, 'utf-8')  
}

// to Param and API

type Param = {
  name : string // name
  varName : string // use in go, as param name
  paramType : ParamType // which type of param
  dataType : string  // 
  isArray : boolean // is array or not
  required : boolean  // is required or not
  comment : string // comment used in swag
}


type API = {
  summary : string // first line
  name : string // name of function
  description : string
  tag : string // mastodon/accounts ...
  returns : string  // @Produce but all json
  request : Param[] // @Params
  path : string
  method : string // "GET" ...
  router : string // /path/:param  used in go
  swagPath : string // used in swag

  // oauth : string // not used
}


function ParamRawToParam(raw: ParamRaw) {
  // name is the raw name
  // eg: something_nothing, hash[key][another], hash[][arrkey]
  const name = raw.name 

  // this may do later. 
  // todo: array, hash
  let varName = ""
  const matches = name.match(/w+/g); 
  if (matches == null) {console.log(raw);}
  else varName = matches.join("_")

  // it's the type 
  // "formData" | "path" | "header" | "query" | "unknown"
  const paramType = raw.paramType

  // data type
  // string for header and path
  // query is all string but sometimes should turn to int
  // formData is complex which contains hash[] and array[]
  // and files.
  // find it in comment
  // is array 
  // required 
  // find them in comment
  // what's more, it should meet the golang types of varibles.
  let dataType : string = "undefined"
  let isArray : boolean = false
  if (paramType == "header" || paramType == "path") {
    dataType = "string"
    isArray = false
  } else if (paramType == "query" || paramType == "formData") {
    console.log(raw) // todo
  } else {
    console.log(raw)
  }
  const required = raw.comment.includes("REQUIRED") // required

  // comment is passed as it is
  // this line should present in swag comment
  const comment = raw.comment // as is
  return {
    name,
    varName,
    paramType,
    dataType,
    isArray,
    required,
    comment,
  } as Param
}

function APIRawToAPI(raw: APIRaw) {
  // summary is the title
  // which used in swagger
  const summary = raw.summary

  // the name used for the function name in golang.
  // anyway...
  const matches = summary.match(/w+/g); 
  if (matches == null) {console.log(raw); return null}
  const name = matches.join("_")

  // comment used in swagger
  const description = raw.description
  
  // topo, it should show folder's name
  const tag = 'mastodon'

  // dunno
  // this should used in swagger?
  // all apis return json though..
  const returns = "json"

  // parse requests, see the function above
  const request = raw.request.map(r => ParamRawToParam(r)).filter(r => r !== null)    

  // http method
  // used in gin router and swagger
  const method = raw.method

  // router
  // passed by raw
  const router = raw.router
  
  // the path used in swagger
  // not tested
  const swagPath = router.split('/')
  .map(v => {
    if (v.startsWith(':')) {
      return `{${v.substring(1)}}`
    } else {
      return v
    }
  }).join('/')

  // this is what in raw document
  // same as gin router. should use it
  const path = raw.router
  return {
    summary ,
    name ,
    description ,
    tag ,
    returns ,
    request ,
    path ,
    method ,
    router ,
    swagPath ,
  } as API
}


// entity (later)

type EntityRaw = {
  description : string
  name : string
  type : string
}

// for test

// step: 0
// paste all document in text file.

// step: 1
// convert text file to raw json file.
// traverseText('text')

// step: 2
// convert raw json file to cooked json file