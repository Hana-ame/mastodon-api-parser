
export class LineReader {
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
  lastline() {
    if (this.idx == 0) return ""
    const r = this.lines[this.idx-1]
    return r.trim()
  }
  isEmpty() {
    return this.idx >= this.lines.length
  }
}

function capitalizeFirstLetter(word: string) {
  return word.charAt(0).toUpperCase() + word.slice(1);
}

export function toPascalCase(name: string) {
  return name
          .match(/[a-zA-Z]+/gi)!
          .map(w => capitalizeFirstLetter(w))
          .join('')
}

export function trimSurfix(s: string, surfix: string) {
  if (s.endsWith(surfix)) {
    return s.substring(0, s.length-surfix.length)
  }
  return s
}

export function trimPrefix(s: string, prefix: string) {
  if (s.startsWith(prefix)) {
    return s.substring(prefix.length) 
  }
  return s
}