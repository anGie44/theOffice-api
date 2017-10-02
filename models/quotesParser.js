var $  = require('cheerio')
var d3 = require('d3')

/////////// Helper Functions ////////////

function pairwise(list) {
    if (list.length < 2) { return []; } 
    var first = list[0], rest = list.slice(1), pairs = rest.map(function(x) { return [first, x]; });
    return pairs.concat(pairwise(rest));
}

Set.prototype.union = function(setB) {
    var union = new Set(this);
    for (var elem of setB) {
        union.add(elem);
    }
    return union;
}

////////// Modules //////////////////

/* Return all quotes for a given episode */
const quotes = function(html) { 
    var parsedHTML = $.load(html)

    // get all div's w/ class: quote
    var scenes  = [], episode_num = '', episode_name = '';
    parsedHTML('div.quote').filter(function(i, elm) { 
        return !parsedHTML(this).children().first().text().includes('Deleted');
        }).each(function(i, elm) {
                scenes[i] = parsedHTML(this).text().replace(/\t/g, '').split('\n').filter(quote => quote.length > 0);
    });

    parsedHTML('b').filter(function(i, elm) {
        return parsedHTML(this).text().includes("Season") && parsedHTML(this).text().includes("Episode");  
    }).each(function(i, elm) {
        var textStr = parsedHTML(this).text().replace(/\n+\t+/g, ' ');
        episode_num = textStr.match(/Episode (\d+)/)[1];

        if (/"/.test(textStr)) {
            episode_name = textStr.match( /"(.*?)"/ )[1];
        }
        else { // episode_name not found w/in <b> tag
            episode_name = parsedHTML(this).next().next().text().match(/"(.*?)"/)[1];
        }

    })
    return {"episode" : episode_num, "name": episode_name, "quotes" : scenes};
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

const links_and_nodes = function(season_data) {
   var result = []
    season_data.data.forEach(item => {
        [links, nodes] = generateNodesAndLinks(item.quotes)
        result.push({"episode": item.episode, "links": links, "nodes": nodes})
    })
    return result;
}


const generateNodesAndLinks = function(episode_quotes) {
    var links = [];
    var nodes = new Set();

    Object.keys(episode_quotes).forEach(key => {
        var characters_in_scene = episode_quotes[key].filter(name => name.indexOf(':') != -1).map(name => name.split(":")[0]);
        nodes = nodes.union([...new Set(characters_in_scene)]);
        var char_links = pairwise([...new Set(characters_in_scene)].sort());
        for (i = 0; i < char_links.length; i++) {
            links.push({"source": char_links[i][0], "target": char_links[i][1]});
        }
    });

    var linkCounts = d3.nest()
        .key(function(d) { return d.source; })
        .key(function(d) { return d.target; })
        .rollup(function(v) { return v.length; })
        .entries(links);

    var final_links = [];
    for(var i = 0; i < linkCounts.length; i++) {
        for (var j = 0; j < linkCounts[i]["values"].length; j++) {
            final_links.push({"source": linkCounts[i].key, "target" : linkCounts[i]["values"][j].key, "value": linkCounts[i]["values"][j].value });
        }
    }
    
    nodes = [...nodes].map(item => ({ "id" : item }));
    return [final_links, nodes];
}

     
module.exports = {
    quotes,
    episodes,
    links_and_nodes
}
