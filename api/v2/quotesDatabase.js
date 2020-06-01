const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');

var config = {
    url: process.env.CONNECTION_URL,
    username: process.env.USERNAME,
    password: process.env.PASSWORD,
    dbName: process.env.DATABASE_NAME,
    collectionName: process.env.COLLECTION_NAME
}

const mongoClient = function() {
    const uri = `mongodb+srv://${config.username}:${config.password}@${config.url}`;
    return new MongoClient(uri, { useNewUrlParser: true });
}

async function quotesByEpisode(season, episode) {
    const client = mongoClient();
    var quotes = [];
    try {
        await client.connect();
        const collection = client.db(config.dbName).collection(config.collectionName);
        quotes = await collection.find({ season: Number(season), episode: Number(episode) }).sort({scene : 1, line: 1}).toArray();
        
    } catch (err) {
        console.log(err.stack);
    }

    client.close();
    return {"quotes":quotes}

}


module.exports = {
    quotesByEpisode,
}