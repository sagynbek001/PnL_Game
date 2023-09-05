const http = require('http');
const express = require('express');
const config = require('./framework/config');

const app = express();
const server = http.createServer(app);

app.use(express.json());
app.use(express.urlencoded({extended: true}));

app.get('/', (req,res) => {

});

server.listen(config.get('app.port'), () => {
	console.log('APP Server ===> %s is listening on the port %d', config.getPNLHost(),
		config.get('app.port'));
})
