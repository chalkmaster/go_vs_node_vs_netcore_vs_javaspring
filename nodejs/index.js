var express = require('express');
var cors = require('cors');
var app = express();

app.use(cors());

const database = [];

app.get('/', (req, res) => {
    console.log(1);
    res.send('ok').end();
});

app.get('/player/progress/:media_id', (req, res) => {
    if (!database[req.params.media_id]) {
        database[req.params.media_id] = 0;
    }
    res.send({progress: database[req.params.media_id]}).end();
});

app.patch('/player/progress/:media_id/:position', (req, res) => {
    database[req.params.media_id] = req.params.position;
    res.sendStatus(200).end();
});

app.listen(8123, () => { console.log('running'); });