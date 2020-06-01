var express = require('express');
var router = express.Router();

var office_quote_extractor_v2 = require('./quotesDatabase.js')

router.get('/', function(req, res) {
    res.json({ message: 'welcome to v2.0 of the Office api!' });
})

router.get('/season/:season/format/:format', function(req, res, next) {
    season = req.params.season      
    format = req.params.format
    quote_data = office_quote_extractor_v2.quotes(season, format);   
    res.json({data: quote_data}); 
})

router.get('/season/:season/episode/:episode', function(req, res, next) {
    season = req.params.season
    episode = req.params.episode
    office_quote_extractor_v2.quotesByEpisode(season, episode)
    .then(result => {
        res.json({data: result});
    });
})

module.exports = router;

