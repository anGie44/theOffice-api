# theOffice-api
a REST api to retrieve the Office quotes when needed (ALWAYS...obvio 💁‍)

Currently hosted at: https://the-office-api.herokuapp.com

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
           **Content:** {"data" : {"season": _seasonNumber_ , "episode" : { "_episodeNumber_" : { "name" : "_episodeName_", "quotes" : [[]] }
   
* Get quotes for a specific season and episode
    * **URL:**          _/season/:season/episode/:episode_
    * **Method:**       `GET`
    * **URL Params:**
    
         **Required:** 
         
         `season=[integer] // season number [1-9], inclusive`
         
         `episode=[integer] // episode number within season (indexing begins at 1)`
    * **Success Response:**
     * **Code:** 200 <br />
       **Content:** JSON
    
                
