var express = require('express')
var app = express();
var request = require('request')
var office_quote_extractor = require('./models/quotesParser.js')

var port = process.env.POR || 8080;
var router = express.Router();
var domain = 'http://officequotes.net/';

router.get('/', function(req, res) {
        res.json({ message: 'welcome to the Office api!' });
})

router.get('/season/:season', function(req, res) {
        var seasonKey = 's' + req.params['season'];
        var episodeKey = 'ALL';
        var q_data = [];

        request(domain, function(err, resp, html) {
            if (!err && resp.statusCode == 200) {       
                var urls = office_quote_extractor.episodes(html, req.params['season']);
                for (i in urls) {
                        request(domain + urls[i], function(error, response, html) {
                                if (!error && response.statusCode == 200){
                                        q_data.push(office_quote_extractor.quotes(html));
                                }
                                if (q_data.length == urls.length) {
                                        quote_data = { "season" : seasonKey, "episode" : q_data.reduce(function(acc, cur, i) { acc['e'+(i+1)] = cur; return acc;}, {}) };
                                        res.json({data: quote_data});
                                }

                        })
               }
            }
        })
})


router.get('/season/:season/episode/:episode', function(req, res) {
        var episode = parseInt(req.params['episode']);
        var seasonKey = 's' + req.params['season'];
        var episodeKey = 'e' + req.params['episode'];
        if (parseInt(episode) < 10) episode = '0' + episode;
        var quotePage = 'no' + req.params['season'] + '-' + episode + '.php';
        request(domain + quotePage, function(error, response, html) {
            if (!error && response.statusCode == 200){
                quote_data = { "season" : seasonKey, "episode" : episodeKey, "quotes": office_quote_extractor.quotes(html) };        
                res.json({data: quote_data});

            }
        })
})

router.get('/search/season/:season/:key', function(req, res) {
        res.send(req.params);
})

app.use('/', router);

app.listen(port);
console.log('Magic happens on port ' + port);


