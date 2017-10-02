var cors = require('cors')
var express = require('express')
var fs = require('fs')
var https = require('https')
var rp = require('request-promise')
var office_quote_extractor = require('./models/quotesParser.js')

var app = express();
var port = process.env.PORT || 8080;
var domain = 'http://officequotes.net/';
var router = express.Router(); 
router.use(cors())

/* Uncomment https-server code to run locally */

/*
https.createServer({
        key: fs.readFileSync('key.pem'),
        cert: fs.readFileSync('cert.pem')
}, app).listen(port);

*/

app.listen(port);

router.get('/', function(req, res) {
        res.json({ message: 'welcome to the Office api!' });
})

router.get('/season/:season/format/:format', function(req, res, next) {
        var seasonKey = 's' + req.params.season;
        var quote_data = [];

        rp(domain)
            .then(function(htmlString) {
                    var urls = office_quote_extractor.episodes(htmlString, req.params.season);
                    for (i in urls) {
                        rp(domain + urls[i])
                            .then(function(htmlString) {
                                    quote_data.push(office_quote_extractor.quotes(htmlString));
                                    if (quote_data.length == urls.length) {
                                            // quote_data = {"episode" : q_data.reduce(function (acc, curr) { acc[curr[1]] = {"name": curr[2], "quotes": curr[0]} ; return acc; }, {} )};
                                            if (req.params.format == "quotes") {
                                                res.json({data: quote_data.sort((a,b) => parseInt(a.episode) - parseInt(b.episode))});
                                            }
                                            else if (req.params.format == "connections") {
                                                res.locals = {data: quote_data};
                                                next()
                                            }
                                    }

                        })
               }
            })
        
}, function(req, res, next) {
    res.json({data: office_quote_extractor.links_and_nodes(res.locals)})
})


router.get('/season/:season/episode/:episode', function(req, res) {
        var episode = parseInt(req.params.episode);
        if (parseInt(episode) < 10) episode = '0' + episode;
        var quotePage = 'no' + req.params.season + '-' + episode + '.php';
        rp(domain + quotePage)
            .then(function(htmlString) {
                var quote_data = office_quote_extractor.quotes(htmlString);   

                res.json({data: {...{"season" : req.params.season}, ...quote_data}});
            })
})

router.get('/search/season/:season/:key', function(req, res) {
        res.send(req.params);
})

app.use('/', router);
console.log('Magic happens on port ' + port);


