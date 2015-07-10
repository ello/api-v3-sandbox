/* jshint esnext:true, strict:true, node:true, mocha:true */
"use strict";
let assert = require('assert');
let expect = require('chai').expect;
let sinon = require('sinon');

let server = require(__dirname + "/../server");


describe("API", function () {
    const options = {
        method: "GET",
        url: "/users",
        headers: {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        },
    };


    describe("users#GET", function () {
        let sandbox;
        let stub;

        beforeEach(function () {
            sandbox = sinon.sandbox.create();
        });

        afterEach(function () {
            sandbox.restore();
        });

        it("should return status code 200", function (done) {
            server.inject(
                options,
                function (response) {
                    expect(response.statusCode).to.equal(200);
                    done();
                }
            );
        });
    })
});