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

function CreateTableFromJSON(table, items) {
    // var table = document.createElement("table");                             // the table elements

    // var col = Object.keys(array[0]);                                         // the columns names (I think taking the keys of the first object will suffice)
    var col = items;
    // HEADER:
    var tr = table.insertRow(-1);                                            // the header row
    col.forEach(function(key) {                                              // for each key in col
      var th = document.createElement("th");                                 // create a header cell
      th.textContent = key;                                                  // use textContent instead of innerHTML (it's better)
      tr.appendChild(th);
    });

    // ROWS:
    array.forEach(function(obj) {                                            // for each object obj in company_info
      var tr = table.insertRow(-1);                                          // create a row for it
      col.forEach(function(key) {                                            // and for each key in col
        var tabCell = tr.insertCell(-1);                                     // create a cell
        if (Array.isArray(obj[key])) {                                       // if the current value is an array, then
          obj[key].forEach(function(contact) {                               // for each entry in that array
            var div = document.createElement("div");                         // create a div and fill it
            div.textContent = contact.first_name + " " + contact.last_name + ", " + contact.position;
            tabCell.appendChild(div);                                        // then add the div to the current cell
          });
        } else {                                                             // otherwise, if the value is not an array (it's a string)
          tabCell.textContent = obj[key];                                    // add it as text
        }
      });
    });

    var divContainer = document.getElementById("showData");
    divContainer.innerHTML = "";
    divContainer.appendChild(table);
  }

$(document).ready(function () {
    let tableEle = document.getElementById("dmp");
    let items = JSON.parse($('#items').val());
    (document.getElementById("title")).innerHTML = "dmp.json (" + items.schema + ")";
    CreateTableFromJSON(tableEle, items)
});
