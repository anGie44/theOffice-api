# theOffice-api
a REST api to retrieve the Office quotes when needed (ALWAYS 💁‍)

Currently hosted at: https://the-office-api.herokuapp.com

Currently used for: https://angie44.github.io/theOffice

![](https://media.giphy.com/media/MaItK5SUgStdm/giphy.gif)


# API Reference

## GET 

* Get quotes or nodes/links by season number 
    * **URL:**           _/season/:season/format/:format_
    * **Method:**       `GET`
    * **URL Params:**
    
         **Required:**
         
         `season=[integer] // season number [1-9], inclusive`
         
         `format=[string] // "quotes" or "connections"`

     * **Success Response:**
      * **Code:** 200 <br />
      
        **Content [Quotes]:** { "data" : { "episode": _episodeNumber_, "name" : _episodeName_, "quotes" : [[]] }
        
        **Content [Connections]:** { "data" : [{ "episode": _episodeNumber_, "name": _episodeName_, "links" : [{ "source" : _characterName_, "target": _characterName_, "value" : _numberOfCoOccurencesInEpisode_ }], "nodes" : [{ "id" : _characterName_ }]}
        
   
* Get quotes for a specific season and episode
    * **URL:**          _/season/:season/episode/:episode_
    * **Method:**       `GET`
    * **URL Params:**
    
         **Required:** 
         
         `season=[integer] // season number [1-9], inclusive`
         
         `episode=[integer] // episode number within season (indexing begins at 1)`
    * **Success Response:**
     * **Code:** 200 <br />
       **Content:** { "data" : { "episode" : _episodeNumber_, "name": _episodeName_, "quotes" : [[]] }
    
                
