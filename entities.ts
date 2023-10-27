import path from "path";
import { LineReader, toPascalCase, trimPrefix, trimSurfix } from "./LineReader";
import fs, { read } from 'fs';

// html documet to raw json
// entities/documents/*.txt => entities/raw/*.json

const docPath = path.join('entities', 'documents')
const rawPath = path.join('entities', 'raw')
const cookedPath = path.join('entities', 'cooked')
const filelist = [
  "Account",
  "Admin_Account",
  "Admin_CanonicalEmailBlock",
  "Admin_Cohort",
  "Admin_Dimension",
  "Admin_DomainAllow",
  "Admin_DomainBlock",
  "Admin_EmailDomainBlock",
  "Admin_Ip",
  "Admin_IpBlock",
  "Admin_Measure",
  "Admin_Report",
  "Announcement",
  "Application",
  "Context",
  "Conversation",
  "CustomEmoji",
  "DomainBlock",
  "Error",
  "ExtendedDescription",
  "FamiliarFollowers",
  "FeaturedTag",
  "Filter",
  "FilterKeyword",
  "FilterResult",
  "FilterStatus",
  "IdentityProof",
  "Instance",
  "List",
  "Marker",
  "MediaAttachment",
  "Notification",
  "Poll",
  "Preferences",
  "PreviewCard",
  "Reaction",
  "Relationship",
  "Report",
  "Role",
  "Rule",
  "ScheduledStatus",
  "Search",
  "Status",
  "StatusEdit",
  "StatusSource",
  "Suggestion",
  "Tag",
  "Token",
  "Translation",
  "V1_Filter",
  "V1_Instance",
  "WebPushSubscription",
]
// those points should be careful
// hash
// array of xxx

type EntityRaw = {
  name : string // the name of value
  type : string // type of value
  description : string // description
}

// @input fn: file name (also the object name).
// @input text: the file contents.
function textToJson(fn: string) {
  // 请先筛出然后到打印,察觉所有情况都没漏掉之后进行自动化
  const src = fs.readFileSync(path.join(docPath, `${fn}.txt`), 'utf8')
  const dst : EntityRaw[] = []
  const reader = new LineReader(src)
  while (reader.readline() != 'Attributes');
  while (reader.readline() != 'Attributes');
  while (!reader.isEmpty()) {
    const name = reader.readline()
    const description = reader.readline()
    const type = reader.readline()
    while (reader.readline() != '');
    if (!(description.startsWith('Description: ') && type.startsWith('Type: '))) break;
    dst.push({
      name,
      type,
      description,
    } as EntityRaw)
  }
  const fileContent = JSON.stringify(dst, null, "  ")
  fs.writeFileSync(path.join(rawPath, `${fn}.json`), fileContent, 'utf-8') 
}

function textToJsonAll() {
  filelist.forEach(fn => textToJson(fn))
}

type Entity = {
  name : string // the raw name
  jsonName : string // the name used in json 
  valueName : string // the name used in Golang 
  optional : boolean // is OPTIONAL
  type : string // the raw type
  valueType : string // type used in golang, not include []
  nullable : boolean // is nullable
  isArray : boolean // is array
  description : string // description
}

function toGoType(type: string) {
  if (type == 'Boolean') {
    return 'bool'
  } else if (type == 'Integer') {
    return 'int'
  } else if (type == 'String') {
    return 'string'
  } else if (type == 'Hash') {
    return 'map[string]any'
  }
  if (type.indexOf('::') > -1) {
    return type.split('::').join('')
  }
  return type
}

function rawToEntity(raw: EntityRaw) {
  const name = raw.name
  const jsonName = trimSurfix(name, ' OPTIONAL').match(/\w+/gi)!.pop()!
  console.log(jsonName)
  const valueName = toPascalCase(jsonName).replace('Id', 'ID')
  const optional = name.endsWith(' OPTIONAL')
  const type = raw.type
  console.log(type)
  let valueType = trimPrefix(type, 'Type: ')
  let nullable = false
  let isArray = false
  if (valueType.startsWith('NULLABLE ')) {
    nullable = true
    valueType = trimPrefix(valueType, 'NULLABLE ')
  }
  if (valueType.startsWith('Array of ')) {
    isArray = true
    valueType = trimPrefix(valueType, 'Array of ')
  }
  valueType = toGoType(trimSurfix(valueType.split(' ')[0], ','))
  const description = raw.description
  return {
    name,
    jsonName,
    valueName,
    optional,
    type,
    valueType,
    nullable,
    isArray,
    description,
  }
}

function rawToEntityAll() {
  filelist.forEach(fn => {
    const src = fs.readFileSync(path.join(rawPath, `${fn}.json`), 'utf8')
    const rawArray: EntityRaw[] = JSON.parse(src)
    const dstArray = rawArray.map(raw => rawToEntity(raw))
    const fileContent = JSON.stringify(dstArray, null, "  ")
    fs.writeFileSync(path.join(cookedPath, `${fn}.json`), fileContent, 'utf-8') 
  })
}

// step 1:
// parse text document to raw json
// textToJsonAll()

// step 2:
// parse raw json to entity
rawToEntityAll()