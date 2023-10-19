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
  path : string
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
  // 读取当前文件夹中的内容
  const files = fs.readdirSync(currentPath);

  // 遍历每个文件/文件夹
  files.forEach((file) => {
    const filePath = path.join(currentPath, file);
    const stats = fs.statSync(filePath);

    if (stats.isFile()) {
      // 处理文件
      console.log('File:', filePath);
      textToAPRraw(currentPath, file)
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

function textToAPRraw(dir :string, file: string) {
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
traverseText('text')