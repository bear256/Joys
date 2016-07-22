// Node.js version
var lineReader = require('readline').createInterface({
    input: require('fs').createReadStream('UCloud.txt')
});
var total = 0;
var i = 0;
lineReader.on('line', function (text) {
    var count = (text.match(/UCanUup/g) || []).length;
    console.log('Line', i, ':', count);
    total += count;
    i++;
}).on('close', function () {
    console.log('Total:', total);
});
