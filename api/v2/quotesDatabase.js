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
    var name = "";

    try {
        await client.connect();
        const collection = client.db(config.dbName).collection(config.collectionName);
        query = { season: Number(season), episode: Number(episode) }
        fields = { _id: 0, character: 1, quote: 1 }
        quotes = await collection.find(query)
            .sort({scene : 1, line: 1})
            .project(fields)
            .toArray();
        found = await collection.findOne(query);
        if (found != null) {
            name = found.episode_name;
        }

    } catch (err) {
        console.log(err.stack);
    }

    client.close();

    return {"episode_name" : name, "quotes" : quotes }

}

module.exports = {
    quotesByEpisode,
}