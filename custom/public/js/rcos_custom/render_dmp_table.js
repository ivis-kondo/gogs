// createTable is RCOS specific code.
// This generates a table to check the information of the DMP
// created by the user and displays it on the screen.
// function createTable(element, items) {
//     for (const item in items) {
//         let tr = document.createElement("tr");

//         // add key
//         let field = document.createElement("td");
//         field.innerHTML = item
//         tr.appendChild(field);

//         // add value
//         if (typeof items[item] === "object") {
//             createTable(tr, items[item])
//         } else {
//             let value = document.createElement("td");
//             value.innerHTML = items[item]
//             tr.appendChild(value);
//         }
//         element.appendChild(tr);
//     }
// }



// $(document).ready(function () {
//     let tableEle = document.getElementById("dmp");
//     let items = JSON.parse($('#items').val());
//     (document.getElementById("title")).innerHTML = "dmp.json (" + items.schema + ")";
//     createTable(tableEle, items)
// });

function outputLog(msg) {
    const fs = require("fs");
    log_msg = String(msg) + '\n'
    fs.appendFile("./log.txt",log_msg, (err) => {
        if (err) throw err;
        // console.log('正常に書き込みが完了しました');
      });
}

function IsArrayObject(data)  {
    return Array.isArray(data)
}

function IsObject(data)  {
    return (typeof data === "object")
}

function IsStringOrBoolOrNumber(data) {
    type = typeof data
    if (type === 'string') {
        return true
    } else if (type === 'number') {
        return true
    } else if (type === 'boolean') {
        return true
    } else {
        return false
    }
}

function createDmp2dArr(rowNum, colNum) {
    // create dmp2DArr 2d array
    var dmp2DArr = [] //DMP data 2D array
    for (let i =0; i < rowNum; i++)  {
        let dmpArrChild = []
        for (let j = 0; j < colNum; j++) {
            dmpArrChild.push("")
        }
        dmp2DArr.push(dmpArrChild)
    }
    return dmp2DArr
}

function fillDmp2dArr(dmp2DArr, dmp, raw_index, col_index) {
    console.log('fillDmp2dArr start')
    console.log('raw_index :' + raw_index)
    console.log('col_index :' + col_index)
    // fill DMP data 2d Array
    for (const key in dmp) {
        console.log('key_1 :' + key)
        const value = dmp[key]
        console.log('value :' + value)
        console.log('value type :' + typeof value)
        const record_num =  getLeafNum(value)
        console.log('record_num :' + record_num)
        console.log('raw_index :' + raw_index)
        endPoint = raw_index + record_num
        console.log('endPoint :' + endPoint)

        for (i =raw_index; i< endPoint; i++) {
            console.log('dmp2DArr[' + i + '][' + col_index + '] = ' + key)
            dmp2DArr[i][col_index] = key
            console.log('dmp2DArr['+ i +'] :' + dmp2DArr[i])
            if (record_num == 1) {
                console.log('dmp2DArr[' + i + '][' + (col_index+1) + '] = ' + String(value))
                dmp2DArr[i][col_index+1] = String(value)
                console.log('dmp2DArr['+ i +'] :' + dmp2DArr[i])
            }else{
                if (IsArrayObject(value)) {
                    // is array
                    console.log('Is Array!!!!!')
                    console.log('value.length :' + value.length)
                    console.log('value :' + value)
                    console.log('value :' + JSON.stringify(value))
                    let arr_start_index = i
                    console.log('PRE arr_start_index :' + arr_start_index)
                    for (let nested_i = 0; nested_i < value.length; ++nested_i) {
                        console.log('nested_i :' + nested_i)
                        console.log('value.length :' + value.length)
                        const arr_data = value[nested_i]
                        console.log('arr_data :' + arr_data)
                        console.log('arr_data :' + JSON.stringify(arr_data))
                        if (IsObject(arr_data)) {
                            console.log('Is Obj!!!!!')
                            const arr_data_record_num =  getLeafNum(arr_data)
                            console.log('arr_data_record_num :' + arr_data_record_num)
                            console.log('dmp2DArr[' + arr_start_index + '][' + (col_index+1) + '] = ' + nested_i)
                            dmp2DArr[arr_start_index][col_index+1] = nested_i
                            next_raw_index = arr_start_index
                            next_col_index = col_index+2
                            dmp2DArr = fillDmp2dArr(dmp2DArr, arr_data, next_raw_index, next_col_index)
                            console.log('ADD arr_start_index :' + arr_start_index + '+' + arr_data_record_num)
                            arr_start_index = arr_start_index + arr_data_record_num
                            console.log('CHANGE arr_start_index :' + arr_start_index)
                        } else {
                            console.log('Is Not Obj!!!!!')
                            dmp2DArr[arr_start_index][col_index+1] = nested_i
                            dmp2DArr[arr_start_index][col_index+2] = arr_data
                            console.log('ADD arr_start_index :' + arr_start_index + '+' + 1)
                            arr_start_index = arr_start_index + 1
                            console.log('CHANGE arr_start_index :' + arr_start_index)
                        }
                    }

                } else {
                    next_raw_index = raw_index
                    next_col_index = col_index+1
                    dmp2DArr = fillDmp2dArr(dmp2DArr, value, next_raw_index, next_col_index)
                }

            }
        }

        raw_index = raw_index + record_num
        console.log('changed raw_index :' + raw_index)
    }

    console.log('dmp2DArr.length :' + dmp2DArr.length)
    for (let i = 0; i < dmp2DArr.length; ++i) {
        console.log(dmp2DArr[i]);
    }
    return dmp2DArr
}

function createRowspanOrColspan2DArr(rowNum, colNum) {
        // create resed each 2d array
        var rowspanOrColspan2DArr = [] //rowspan info 2D array
        for (let i =0; i < rowNum; i++)  {
            let rowspanArrChild = []
            for (let j = 0; j < colNum; j++) {
                rowspanArrChild.push(0)
            }
            rowspanOrColspan2DArr.push(rowspanArrChild)
        }
    // outputLog('rowspanOrColspan2DArr.length :' + rowspanOrColspan2DArr.length)
    // for (let i = 0; i < rowspanOrColspan2DArr.length; ++i) {
    //     outputLog(rowspanOrColspan2DArr[i]);
    // }
    return rowspanOrColspan2DArr
}



function getLeafNum(data) {
    if (IsStringOrBoolOrNumber(data)) {
        return 1
    } else {
        return getRowNum(data)
    }
}

function getRowNum(data) {
    //get creating row num
    let row_num = 0
    for (const key in data) {
        //console.log("key : " + key)
        let value = data[key]
        if (IsArrayObject(value)) {
            // is array
            for (let i = 0; i < value.length; ++i) {
                let arr_val = value[i]
                if  (IsObject(arr_val)) {
                    row_num = row_num + getRowNum(value[i])
                    //console.log("CHANGE OBJECT-ARRY row_num : " + row_num)
                }else{
                    row_num = row_num + 1
                   // console.log("CHANGE STRING-ARRY row_num : " + row_num)
                }
            }
        } else if (IsObject(value)) {
            // is json object
            row_num = row_num + getRowNum(value)
            //console.log("CHANGE OBJECT row_num : " + row_num)
        } else {
            // is str, num, bool...
            row_num = row_num + 1
            //console.log("CHANGE STRING row_num : " + row_num)
        }
    }
    //console.log("RETTURN row_num : " + row_num)
    return row_num
}

function getColNum(data) {
    //get creating col num
    return getNestedNum(data, 0) + 2 //列数
}

function getNestedNum(data, pre_nested_num) {
    // get nested num in dmp json
    let nested_num = pre_nested_num

    for (key in data) {
        //console.log("key : " + key)
        let value = data[key]
        if (IsArrayObject(value)) {
            //console.log("is array")
            //console.log("value.length : " + value.length)
            // is array
            for (let i = 0; i < value.length; ++i) {
                let arr_val = value[i]
                //console.log("arr_val : " + arr_val)
                if (IsObject(arr_val)){
                    arr_tmp = getNestedNum(arr_val, pre_nested_num+1) + 1
                    //console.log("COMPARE [ arr_tmp : " + arr_tmp + " now nested_num : " +nested_num + "]")
                    if (arr_tmp > nested_num) {
                        nested_num = arr_tmp
                        //console.log("CHANGE nested_num ARRAY : " + nested_num)
                    }
                }
            }
        } else if (IsObject(value)) {
            // is json object
            tmp = getNestedNum(value, pre_nested_num+1)
            //console.log("COMPARE [ tmp : " + tmp + " now nested_num : " +nested_num + "]")
            if (tmp > nested_num) {
                nested_num = tmp
                //console.log("CHANGE nested_num OBJECT : " + nested_num)
            }
        } else {
            // is str, num, bool...
        }
    }
    //console.log("RETURN nested_num : " + nested_num)
    return nested_num
}


function CreateTableFromJSON(tbodyEle, dmp) {
    let rowNum = getRowNum(dmp) //行数
    let colNum = getColNum(dmp) //列数

    // create reseted 3
    dmpData2DArr = createDmp2dArr(rowNum, colNum)

}

// $(document).ready(function () {
//     let tbodyEle = document.getElementById("dmp");
//     let items = JSON.parse($('#items').val());
//     (document.getElementById("title")).innerHTML = "dmp.json (" + items.schema + ")";
//     CreateTableFromJSON(tbodyEle, items)
// });



// outputLog('テスト')
// var amed_data = JSON.parse(amed);
// let rowNum = getRowNum(amed_data) //行数
// outputLog('rowNum :' + rowNum)
// let colNum = getColNum(amed_data) //列数
// outputLog('colNum :' + colNum)
// let dmpData2dArr = createDmp2dArr(rowNum, colNum)
// outputLog('dmpData2dArr.length :' + dmpData2dArr.length)
// for (let i = 0; i < dmpData2dArr.length; ++i) {
//     outputLog(dmpData2dArr[i]);
// }
// fillDmp2dArr(dmpData2dArr, amed_data, 0, 0)


// var min = '{ "workflowIdentifier": "basic","project": {"fiscalYear": 2021,"title": "The Project","problemName": "aaaa","representative": {"belongTo": "NII","post": "aaaaaa","name": "John Doe"}}}'
// outputLog('テスト')
// var min_data = JSON.parse(min);
// let rowNum = getRowNum(min_data) //行数
// outputLog('rowNum :' + rowNum)
// let colNum = getColNum(min_data) //列数
// outputLog('colNum :' + colNum)
// let dmpData2dArr_1 = createDmp2dArr(rowNum, colNum)
// outputLog('dmpData2dArr.length :' + dmpData2dArr_1.length)
// for (let i = 0; i < dmpData2dArr_1.length; ++i) {
//     outputLog(dmpData2dArr_1[i]);
// }
// dmptable = fillDmp2dArr(dmpData2dArr_1, min_data, 0, 0)
// outputLog('finish fillDmp2dArr')
// for (let i = 0; i < dmptable.length; ++i) {
//     outputLog(dmptable[i]);
// }

// var min = '{ "workflowIdentifier": "basic","project": {"fiscalYear": 2021,"title": "The Project","problemName": "","representative": {"belongTo": "NII","post": "a","name": "John Doe"}}, "researches": {"description": "a","data": [{"title": "The Data","releasePolicy": "a","concealReason": "a","repositoryType": "a","repositoryName": "a","dataAmount": "a"}]}}'

// amed
// rowNum :38　行　OK
// colNum :5　列　OK
// OK
// var min = '{"workflowIdentifier": "basic","contentSize": "1GB","datasetStructure": "with_code","useDocker": "YES","schema": "amed","createDate": "2021/07/21","project": {"fiscalYear": 2021,"title": "The Project","problemName": "hogeho","representative": {"belongTo": "NII","post": "hogeho","name": "John Doe"}},"required": {"hasRegistNecessity": false,"noRegistReason": "hogeho"},"researches": {"description": "hogeho","data": [{"title": "The Data","releasePolicy": "hogeho","concealReason": "hogeho","repositoryType": "hogeho","repositoryName": "hogeho","dataAmount": "hogeho"}]},"forPublication": {"hasOfferPolicy": false,"policyName": "This is just content when \'hasOfferPolicy\' is true"},"researchers": {"numberOfPeople": 1,"manager": {"isConcurrent": false,"personal": {"belongTo": "NII","post": "hogeho","name": "Bill Doe"}},"staff": [{"belongTo": "NII","post": "hogeho","name": "Mary Doe","e-Rad": "0000","canPublished": false,"postType": "Professor","financialResource": "AMED","employmentStatus": "full-time","roles": "data curation","remarks": "hogeho"}]}}'
// outputLog('テスト AMED')
// var min_data = JSON.parse(min);
// var rowNum = getRowNum(min_data) //行数
// outputLog('rowNum :' + rowNum)
// console.log('rowNum :' + rowNum)
// var colNum = getColNum(min_data) //列数
// outputLog('colNum :' + colNum)
// console.log('colNum :' + colNum)
// var dmpData2dArr_1 = createDmp2dArr(rowNum, colNum)
// outputLog('dmpData2dArr.length :' + dmpData2dArr_1.length)
// for (let i = 0; i < dmpData2dArr_1.length; ++i) {
//     outputLog(dmpData2dArr_1[i]);
// }
// dmptable = fillDmp2dArr(dmpData2dArr_1, min_data, 0, 0)
// outputLog('dmptable.length :' + dmptable.length)
// console.log('dmptable.length :' + dmptable.length)
// outputLog('finish fillDmp2dArr')
// for (let i = 0; i < dmptable.length; ++i) {
//     outputLog(dmptable[i]);
//     console.log(dmptable[i]);
// }

//meti ダメ
// 32 行OK
// 4列　OK
// tabel OK
// var min = '{"workflowIdentifier": "basic","contentSize": "1GB","datasetStructure": "with_code","useDocker": "YES","schema": "meti","dmpType": "New","agreementTitle": "The Data Management Plan","agreementDate": "2021-09-20","submitDate": "2021-10-01","corporateName": "The Corporate","researches": [{"index": 1,"title": "The Research Data","description": "This is description.","manager": "John Doe","dataType": "My Data","releaseLevel": 4,"concealReason": "nothing","concealPeriod": "hogehoe","acquirer": "John Lab","acquireMethod": "by download link","remarks": "hogehoe"},{"index": 2,"title": "The Research Data2","description": "This is description.","manager": "Jumpei Kuwata","dataType": "My Data","releaseLevel": 4,"concealReason": "nothing","concealPeriod": "hogehoe","acquirer": "IVIS Lab","acquireMethod": "by download link","remarks": "hogehoe"}]}'
// outputLog('テスト METI')
// var min_data = JSON.parse(min);
// var rowNum = getRowNum(min_data) //行数 OK
// outputLog('rowNum :' + rowNum)
// console.log('rowNum :' + rowNum)
// var colNum = getColNum(min_data) //列数 OK
// outputLog('colNum :' + colNum)
// console.log('colNum :' + colNum)
// var dmpData2dArr_1 = createDmp2dArr(rowNum, colNum)
// outputLog('dmpData2dArr.length :' + dmpData2dArr_1.length)
// console.log('dmpData2dArr.length :' + dmpData2dArr_1.length)
// for (let i = 0; i < dmpData2dArr_1.length; ++i) {
//     outputLog(dmpData2dArr_1[i]);
//     console.log('dmpData2dArr_1[i].length :' + dmpData2dArr_1[i].length)
// } //OK
// dmptable = fillDmp2dArr(dmpData2dArr_1, min_data, 0, 0)
// outputLog('dmptable.length :' + dmptable.length)
// console.log('dmptable.length :' + dmptable.length)
// outputLog('finish fillDmp2dArr')
// for (let i = 0; i < dmptable.length; ++i) {
//     outputLog(dmptable[i]);
//     console.log(dmptable[i]);
// }

// jst
// 44行 OK
// 5列 OK
// table OK
var min = '{"workflowIdentifier": "basic","contentSize": "1GB","datasetStructure": "with_code","useDocker": "YES","schema": "jst","createDate": "2021/07/21","amedNumber": "0000","project": {"fiscalYear": 2021,"title": "The Project","problemName": "hogehoge","representative": {"belongTo": "NII","post": "hogehoge","name": "John Doe"}},"forPublication": {"hasUsed": false,"unwriteReason": "This is just content when \'hasUsed\' is true"},"researches": [{"title": "The Title","type": ["ヒト個人（研究参加者及びヒト試料由来のデータ）", "hoge_type"],"description": "hogehoge","releasePolicy": "hogehoge","concealReason": "hogehoge","hasOfferPolicy": false,"policyName": "This is just content when \'hasOfferPolicy\' is true","repositoryType": "hogehoge","repositoryName": "hogehoge","dataAmount": "hogehoge","dataSchema": "Excel","processingPolicy": "hogehoge","isRegistered": false,"registeredInfo": "hogehoge"}],"researchers": {"numberOfPeople": 1,"manager": {"isConcurrent": false,"personal": {"belongTo": "NII","post": "hogehoge","name": "Bill Doe"}},"staff": [{"belongTo": "NII","post": "hogehoge","name": "Mary Doe","e-Rad": "0000","canPublished": false,"postType": "Professor","financialResource": "AMED","employmentStatus": "full-time","roles": "data curation","remarks": "hogehoge"}]}}'
outputLog('テスト JST')
var min_data = JSON.parse(min);
var rowNum = getRowNum(min_data) //行数
outputLog('rowNum :' + rowNum)
console.log('rowNum :' + rowNum);
var colNum = getColNum(min_data) //列数
outputLog('colNum :' + colNum)
console.log('colNum :' + colNum);
var dmpData2dArr_1 = createDmp2dArr(rowNum, colNum)
outputLog('dmpData2dArr.length :' + dmpData2dArr_1.length)
for (let i = 0; i < dmpData2dArr_1.length; ++i) {
    outputLog(dmpData2dArr_1[i]);
}
dmptable = fillDmp2dArr(dmpData2dArr_1, min_data, 0, 0)
outputLog('dmptable.length :' + dmptable.length)
console.log('dmptable.length :' + dmptable.length);
outputLog('finish fillDmp2dArr')
for (let i = 0; i < dmptable.length; ++i) {
    console.log(dmptable[i]);
    outputLog(dmptable[i]);
}

// moon
// 29行
// 3列
// OK
// var min = '{"workflowIdentifier": "basic","contentSize": "1GB","datasetStructure": "with_code","useDocker": "YES","schema": "moonshot","funder": "AMED","eRadProjectID": "hogehoge","projectName": "誰もが自在に活躍できるアバター共生社会の実現","dataNo": "hogehoge","dataTitle": "〇〇実証においてセンサより撮像したデータ及び関連データ","dataPublished": "2021-04-26","dataDescription": "〇〇実証においてセンサより撮像したデータであり、道路の画像データ","researchField": "ライフサイエンス (Life Science)","dataType": "dataset","fileSize": "<1GB","dataPolicy": "一定期間後に事業の実施上有益なものに対して有償又は無償で提供を開始。但しデータのクレジット標記を条件とする。なおサンプルデータを公開している。","accessRights": "公開 (open access)","availableDate": "2021-04-26","repositoryInformation": "〇〇大学学術機関リポジトリ","repositoryURL": "hogehoge","creatorName": "John Doe","eRadCreatorIdentifer": "hogehoge","hostingInstitution": "hogehoge","dataManager": "John Doe","eRadManagerIdentifer": "hogehoge","dataManagerContact": {"affiliation": "〇〇研究所〇〇部門〇〇課","TEL": "00-000-0000","eMail": "user@example.com"},"remarks": "hogehoge"}'
// outputLog('テスト moon')
// var min_data = JSON.parse(min);
// var rowNum = getRowNum(min_data) //行数
// console.log('rowNum :' + rowNum);
// outputLog('rowNum :' + rowNum)
// var colNum = getColNum(min_data) //列数
// outputLog('colNum :' + colNum)
// console.log('colNum :' + colNum);
// var dmpData2dArr_1 = createDmp2dArr(rowNum, colNum)
// outputLog('dmpData2dArr.length :' + dmpData2dArr_1.length)
// for (let i = 0; i < dmpData2dArr_1.length; ++i) {
//     outputLog(dmpData2dArr_1[i]);
// }
// dmptable = fillDmp2dArr(dmpData2dArr_1, min_data, 0, 0)
// outputLog('dmptable.length :' + dmptable.length)
// console.log('dmptable.length :' + dmptable.length)
// outputLog('finish fillDmp2dArr')
// for (let i = 0; i < dmptable.length; ++i) {
//     outputLog(dmptable[i]);
//     console.log(dmptable[i]);
// }
