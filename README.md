# theOffice-api
a REST api to retrieve the Office quotes when needed (ALWAYS 💁‍)

Currently hosted at: https://the-office-api.herokuapp.com :warning: **JK it's been a minute and the heroku app is 💀...:cry::rofl:** :warning:

Currently used for: https://angie44.github.io/theOffice

![](https://media.giphy.com/media/MaItK5SUgStdm/giphy.gif)


# API Reference

## GET 

* Get quotes or nodes/links by season number
    * **URL:**           _/season/:season/format/:format_
    * **Method:**       `GET`
    * **URL Params:**
    
         **Required:**
         
         * `season=[integer] // season number [1-9], inclusive`
         
         * `format=[string] // "quotes" or "connections"`

     * **Success Response:**
       * **Code:** 200
       * **Content [Quotes]:** `[{ "season": seasonNumber, "episode" : episodeNumber, "scene": sceneNumber, "episode_name": episodeName, "character": character, "quote" : quote}]`
       * **Content [Connections]:** `[{ "episode": episodeNumber, "name": episodeName, "links" : [{ "source" : characterName, "target": characterName, "value" : numberOfCoOccurencesInEpisode }], "nodes" : [{ "id" : characterName }]`
        
   
* Get quotes for a specific season and episode
    * **URL:**          _/season/:season/episode/:episode_
    * **Method:**       `GET`
    * **URL Params:**
    
         **Required:** 
         
         * `season=[integer] // season number [1-9], inclusive`
         
         * `episode=[integer] // episode number within season (indexing begins at 1)`
    * **Success Response:**
      * **Code:** 200
      * **Content:** `[{ "season": seasonNumber, "episode" : episodeNumber, "scene": sceneNumber, "episode_name": episodeName, "character": character, "quote" : quote }]`
    
                
