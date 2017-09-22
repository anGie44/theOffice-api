var rp = require('request-promise');
var $ = require('cheerio');
var request = require('request');

const urls = ["http://www.officequotes.net/no1-01.php"];
const promises = urls.map(url => rp(url));
// 	rp(url).then();
// console.log(promises);
// Promise.all(promises)
// 	.then(function(data) { console.log(data);
// 	})
// 	.catch(function(err) {

// 	});

// request(urls[0], function(error, response, html) {
// 	if (!error && response.statusCode == 200){
// 		 var parsedHTML = $.load(html)
// 		 console.log(html)
// 	}
// });
rp(urls[0])
.then(function(htmlString) {
	x.push(htmlString);
})
