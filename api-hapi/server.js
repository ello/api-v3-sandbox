/* jshint esnext:true, strict:true, node:true*/
"use strict";

let Hapi = require('hapi');
let config = require('config');

let server = new Hapi.Server();
server.connection({
    port: config.get('server.port')
});

let usersGetHandler = function (request, reply) {
    reply(`[{"username":"rtyer", "name":"Ryan Tyer"},{"username":"jayzes", "name":"Jay Zeschin"}]`)
        .code(200)
}

let usersGetConfig = {
    handler: usersGetHandler,
    description: 'Returns users!'
}

server.route([{
    method: ['GET'],
    path: '/users',
    config: usersGetConfig
}]);

// If being run standalone (not in a test)
if (!module.parent) {
    server.register({
        register: require('good'),
        options: {
            reporters: [{
                reporter: require('good-console'),
                events: {
                    log: '*',
                    response: '*'
                }
            }]
        }
    }, function (err) {
        if (err) {
            throw err; // something bad happened loading the plugin
        }

        server.start(function () {
            server.log('info', `Server running at: ${server.info.uri}`);
        });
    });
}

module.exports = server;