const router = require('koa-router')();
const koaBody = require('koa-body');
const db = require('../db');

const Koa = require('koa');
const app = module.exports = new Koa();

app.use(koaBody());

app.use((ctx, next) => {
  ctx.set('Access-Control-Allow-Origin', '*');
  next();
});

router.get('/book', async (ctx) => {
  ctx.body = {
    native: db.native.get(),
    translated: db.translated.get(),
  };
});

router.get('/book/:type', async (ctx, next) => {
  ctx.body = db[ ctx.params.type ].get();
});

app.use(router.routes());

app.listen(8080);
