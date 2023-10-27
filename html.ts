// 把html里的两个文件
// `html/list.txt` 和 `html/document.txt` 表达为txt格式的可以喂给 index.ts 的文件.

import { exec } from 'child_process';
import fs, { read } from 'fs';

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
  isEmpty() {
    return this.idx >= this.lines.length
  }
}

const listfile = fs.readFileSync('html/list.txt', 'utf8')
const docfile = fs.readFileSync('html/document.txt', 'utf8')

// read all names of methods.
// ok
const nameList = listfile
                  .split('\n')
                  .map(l => l.trim())
                  .filter(l => l !== '')

// check line in name list
// ok
function checkStringInArray(s: string, list: string[]) {
  const arr = list.filter(l => l == s)
  return arr.length > 0
} 

// main
let nowName = ""
let nowDoc = ""
// files

function saveWithName(name: string, data: string){
  console.log('save', name)
  try {
    fs.mkdirSync(`text/${process.argv[2]}/`)
  } catch (e) {

  }
  fs.writeFileSync(`text/${process.argv[2]}/${name}.txt`, data)
}
const reader = new LineReader(docfile)
// console.log(reader, reader.isEmpty())
while (!reader.isEmpty()) {
  const line = reader.readline()
  // console.log(line)
  if (checkStringInArray(line, nameList)) {
    if (nowName !== '') {
      saveWithName(nowName, nowDoc)
      nowDoc = ''
    }    
    nowName = line
  }
  nowDoc = nowDoc + line + '\n'
}
saveWithName(nowName, nowDoc)

console.log(process.argv)
// exec(`./move_txtfiles.sh ${process.argv[2]}`)