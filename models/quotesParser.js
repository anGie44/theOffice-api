var $  = require('cheerio')

const quotes = function(html) { 
    var parsedHTML = $.load(html)

    // get all div's w/ class: quote
    var scenes  = []
    parsedHTML('div.quote').filter(function(i, elm) { 
        return !parsedHTML(this).children().first().text().includes('Deleted');
        }).each(function(i, elm) {
                scenes[i] = parsedHTML(this).text().replace(/\t/g, '').split('\n').filter(quote => quote.length > 0);
    })
    return scenes;
}

const episodes = function(html, season) {
    var parsedHTML = $.load(html);
    var episodeUrls = [];
    parsedHTML('div.navEp').find('a').each(function(i,elm) {
        if (parsedHTML(this).attr('href').includes('no'+season)) {  
                episodeUrls.push(parsedHTML(this).attr('href'));
        }
    })
    return episodeUrls;
}
     
module.exports = {
    quotes,
    episodes,
}
