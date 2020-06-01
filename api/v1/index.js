var express = require('express');
var router = express.Router();

router.get('/', function(req, res) {
    res.json({ message: '(DEPRECATED) please make all requests to v2 of the Office api!' });
})

module.exports = router;