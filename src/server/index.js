const router = require('koa-router')();
const koaBody = require('koa-body');
const koa = require('koa');
const db = require('../db');

const Koa = require('koa');
const app = module.exports = new Koa();

app.use(koaBody());

router.get('/book/:type', async (ctx, next) => {

  ctx.body =db[ ctx.params.type ].get();

});

app.use(router.routes())

app.listen(8080);
