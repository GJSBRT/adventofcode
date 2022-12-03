const fs = require('fs');

fs.readFile('data.txt', 'utf8', (err, data) => {
    if (err) throw err;

    let newData = [];
    const splitted = data.split(/\n\s*\n/);
    splitted.forEach((capture, i) => {
        let rowAmount = 0;
        const row = capture.split('\r\n');
        row.forEach(value => {
            if (value == '') return;

            rowAmount = rowAmount + parseInt(value.replace('\r', ''));
        })
        newData.push(rowAmount);
    })

    const sortedData = newData.sort((a, b) => b - a);
    console.log(sortedData[0]);
});