import { Hono } from 'hono';
import { serveStatic } from 'hono/bun';
import { logger } from 'hono/logger';

const app = new Hono();

app.use(logger());

app.use('/*', serveStatic({ root: './public/' }));

// app.get('/', c => {
//   return c.text('Hello Hono, dev!');
// });

// return json
app.get('/api/hello', c => {
  return c.json({
    ok: true,
    message: 'Hello Hono!',
  });
});

// response and reqeust
app.get('/posts/:id', c => {
  const page = c.req.query('page');
  const id = c.req.param('id');
  c.header('X-Message', 'Hi!');
  return c.text(`You want see ${page} of ${id}`);
});

// post, put and delete if you want
app.post('/posts', c => c.text('Created!', 201));
app.delete('/posts/:id', c => c.text(`${c.req.param('id')} is deleted!`));

// return html
const View = () => {
  return (
    <html>
      <body>
        <h1>Hello Hono!</h1>
      </body>
    </html>
  );
};

app.get('/page', c => {
  return c.html(<View />);
});

// return raw response
// app.get('/', c => {
//   return new Response('Good morning!');
// });

console.log('process.env.PORT: ', process.env.PORT);

export default {
  port: process.env.PORT,
  fetch: app.fetch,
};
