# theOffice-api
a REST api to retrieve the Office quotes when needed (ALWAYS...obvio 💁‍)

Currently hosted at: <a href=https://the-office-api.herokuapp.com target="_blank">https://the-office-api.herokuapp.com</a>

![](https://media.giphy.com/media/MaItK5SUgStdm/giphy.gif)


# API Reference

## GET 

* Get quotes or nodes/links by season number 
    * URL       :   /season/:season/format/:format
    * Method    :   GET
    * Request   : 
                { body:
                {
                season: int // season number [1-9],
                format: string // format can be either "quotes" or "connections"
                }
                }
* Get quotes for a specific season and episode
    * URL       :   /season/:season/episode/:episode
    * Method    :   GET
    * Request   : 
                { body:
                {
                season: int // season number [1-9],
                episode: int // episode number within season (indexing begins at 1)
                }
