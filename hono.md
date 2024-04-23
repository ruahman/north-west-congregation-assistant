# hono 

fast, light, built on web standards 
support for any javascript runtime 

Ultrafast and lightweight
The router *RegExRouter* is really fast 

small, simple and ultra fast 

create a hono project 
`bunx create-hono`

hono/tiny is only 14kb 

it uses only web standards

# Routers 

it has 5 routers 

# middleware 

```javascript
app.use(async (c, next) => {
  const start = Date.now()
  await next()
  const end = Date.now()
  c.res.headers.set('X-Response-Time', `${end - start}`)
})
```

# hono stack 

you can use both hono on the backend and frontend 

you can specify the route a a type that you can use on the frontend

```javascript
const route = app.get(
  '/hello',
  zValidator(
    'query',
    z.object({
      name: z.string(),
    })
  ),
  (c) => {
    const { name } = c.req.valid('query')
    return c.json({
      message: `Hello! ${name}`,
    })
  }
)

export type AppType = typeof route 
``` 

you can use hono on the frontend

```javascript 
import { AppType } from './server'
import { hc } from 'hono/client'

const client = hc<AppType>('/api')
const res = await client.hello.$get({
  query: {
    name: 'Hono',
  },
})

const data = await res.json()
console.log(`${data.message}`)
```

# App 

```javascript
import { Hono } from 'hono'

const app = new Hono()
//...

export default app // for Cloudflare Workers or Bundler
```

## not found 

```javascript
app.notFound((c) => {
  return c.text('Custom 404 Message', 404)
})

```

## Error Handling 

```javascript 
app.error((c, err) => {
  console.error(err)
  return c.text('Internal Server Error', 500)
})
```

## fetch 

app.fetch is the entry point of your application

```javascript
export default { 
  port: 3000, 
  fetch: app.fetch, 
} 
```


# request 

request is a useful method for testing 

```javascript
test('GET /hello is ok', async () => {
  const res = await app.request('/hello')
  expect(res.status).toBe(200)
})
```

```javascript
test('POST /message is ok', async () => {
  const req = new Request('Hello!', {
    method: 'POST',
  })
  const res = await app.request(req)
  expect(res.status).toBe(201)
})
```

# strict mode 

/hello and /hello/ are different 

```javascript
const app = new Hono({ strict: true })
```

# router option 

```javascript
import { RegExpRouter } from 'hono/router/reg-exp-router'

const app = new Hono({ router: new RegExpRouter() })
```


## Generics 

```javascript
type Bindings = {
  TOKEN: string
}

type Variables = {
  user: User
}

const app = new Hono<{ Bindings: Bindings; Variables: Variables }>()

app.use('/auth/*', async (c, next) => {
  const token = c.env.TOKEN // token is `string`
  // ...
  c.set('user', user) // user should be `User`
  await next()
})
```

# routing 

```javascript
// HTTP Methods
app.get('/', (c) => c.text('GET /'))
app.post('/', (c) => c.text('POST /'))
app.put('/', (c) => c.text('PUT /'))
app.delete('/', (c) => c.text('DELETE /'))

// Wildcard
app.get('/wild/*/card', (c) => {
  return c.text('GET /wild/*/card')
})

// Any HTTP methods
app.all('/hello', (c) => c.text('Any Method /hello'))

// Custom HTTP method
app.on('PURGE', '/cache', (c) => c.text('PURGE Method /cache'))

// Multiple Method
app.on(['PUT', 'DELETE'], '/post', (c) => c.text('PUT or DELETE /post'))

// Multiple Paths
app.on('GET', ['/hello', '/ja/hello', '/en/hello'], (c) => c.text('Hello'))
```

# path parameters 

```javascript
app.get('/user/:name', (c) => {
  const name = c.req.param('name')
  ...
})
```

```javascript
app.get('/posts/:id/comment/:comment_id', (c) => {
  const { id, comment_id } = c.req.param()
  ...
})
```

# Grouping 

you can group routes and add them to main app with the route method

```javascript
const book = new Hono()

book.get('/', (c) => c.text('List Books')) // GET /book
book.get('/:id', (c) => {
  // GET /book/:id
  const id = c.req.param('id')
  return c.text('Get Book: ' + id)
})
book.post('/', (c) => c.text('Create Book')) // POST /book

const app = new Hono()
app.route('/book', book)
```

# grouping base path 

```javascript
const book = new Hono()
book.get('/book', (c) => c.text('List Books')) // GET /book
book.post('/book', (c) => c.text('Create Book')) // POST /book

const user = new Hono().basePath('/user')
user.get('/', (c) => c.text('List Users')) // GET /user
user.post('/', (c) => c.text('Create User')) // POST /user

const app = new Hono()
app.route('/', book) // Handle /book 
app.route('/', user) // Handle /users
```

# base path 

```javascript
const api = new Hono().basePath('/api')
api.get('/book', (c) => c.text('List Books')) // GET /api/book
```

if you have middleware that you want to execute 

```javascript
app.use(logger())
app.get('/foo', (c) => c.text('foo'))
```

# context 

to handle Request and Response  

```javascript
app.get('/welcome', (c) => {
  // Set headers
  c.header('X-Message', 'Hello!')
  c.header('Content-Type', 'text/plain')

  // Set HTTP status code
  c.status(201)

  // Return the response body
  return c.body('Thank you for coming')
})
```

or you can write it this way 

```javascript
app.get('/welcome', (c) => {
  return c.body('Thank you for coming', 201, {
    'X-Message': 'Hello!',
    'Content-Type': 'text/plain',
  })
})
```

# text 

```javascript
app.get('/hello', (c) => {
  return c.text('Hello!')
})
```

# json 

```javascript
app.get('/hello', (c) => {
  return c.json({ message: 'Hello!' })
})
```

# html 

```javascript
app.get('/hello', (c) => {
  return c.html('<h1>Hello!</h1>')
})
```

# notFound 

```javascript
app.get('/notfound', (c) => {
  return c.notFound()
})
```

# redirect 

```javascript
app.get('/redirect', (c) => {
  return c.redirect('/')
})
app.get('/redirect-permanently', (c) => {
  return c.redirect('/', 301)
})
```

# set/get 

```javascript
app.use(async (c, next) => {
  c.set('message', 'Hono is cool!!')
  await next()
})

app.get('/', (c) => {
  const message = c.get('message')
  return c.text(`The message is "${message}"`)
})
```

if you want to make it type safe 

```javascript
type Variables = {
  message: string
}

const app = new Hono<{ Variables: Variables }>()
```

# middleware 

```javascript
const app = new Hono()

const echoMiddleware: MiddlewareHandler<{
  Variables: {
    echo: (str: string) => string
  }
}> = async (c, next) => {
  c.set('echo', (str) => str)
  await next()
}

app.get('/echo', echoMiddleware, (c) => {
  return c.text(c.var.echo('Hello!'))
})
```

# render / setRender 

you can set a layout using c.setRenderer

```javascript
app.use(async (c, next) => {
  c.setRenderer((content) => {
    return c.html(
      <html>
        <body>
          <p>{content}</p>
        </body>
      </html>
    )
  })
  await next()
})
```

you can then use render 

```javascript
app.get('/', (c) => {
  return c.render('Hello!')
})
```

```javascript
app.use('/pages/*', async (c, next) => {
  c.setRenderer((content, head) => {
    return c.html(
      <html>
        <head>
          <title>{head.title}</title>
        </head>
        <body>
          <header>{head.title}</header>
          <p>{content}</p>
        </body>
      </html>
    )
  })
  await next()
})

app.get('/pages/my-favorite', (c) => {
  return c.render(<p>Ramen and Sushi</p>, {
    title: 'My favorite',
  })
})

app.get('/pages/my-hobbies', (c) => {
  return c.render(<p>Watching baseball</p>, {
    title: 'My hobbies',
  })
})
```

# Request 

## param 

```javascript
// Captured params
app.get('/entry/:id', (c) => {
  const id = c.req.param('id')
  ...
})

// Get all params at once
app.get('/entry/:id/comment/:commentId', (c) => {
  const { id, commentId } = c.req.param()
})
```

# query 

```javascript
// Query params
app.get('/search', (c) => {
  const query = c.req.query('q')
  ...
})

// Get all params at once
app.get('/search', (c) => {
  const { q, limit, offset } = c.req.query()
  ...
})
```


# exception 

```javascript
import { HTTPException } from 'hono/http-exception'

// ...

app.post('/auth', async (c, next) => {
  // authentication
  if (authorized === false) {
    throw new HTTPException(401, { message: 'Custom error message' })
  }
  await next()
})
```

```javascript
import { HTTPException } from 'hono/http-exception'

// ...

app.onError((err, c) => {
  if (err instanceof HTTPException) {
    // Get the custom response
    return err.getResponse()
  }
  //...
})
```

# middleware 

middleware works before/after a handler. We can get Request before dispatching 
or manipulate Response after dispatching.

* Handler should return Response object.  only on handler will be called.
* Middleware returns nothing, will be proceeded to next middleware with `await next()`


you can register middleware using `app.use` 

```javascript
// match any method, all routes
app.use(logger())

// specify path
app.use('/posts/*', cors())

// specify method and path
app.post('/posts/*', basicAuth())
```

if the the handler returns Response it will stop the process 
```javascript   
app.post('/posts', (c) => c.text('Created!', 201))
```

## order of execution 

```javascript
app.use(async (_, next) => {
  console.log('middleware 1 start')
  await next()
  console.log('middleware 1 end')
})
app.use(async (_, next) => {
  console.log('middleware 2 start')
  await next()
  console.log('middleware 2 end')
})
app.use(async (_, next) => {
  console.log('middleware 3 start')
  await next()
  console.log('middleware 3 end')
})

app.get('/', (c) => {
  console.log('handler')
  return c.text('Hello!')
})
```

## custom middleware 

```javascript
// Custom logger
app.use(async (c, next) => {
  console.log(`[${c.req.method}] ${c.req.url}`)
  await next()
})

// Add a custom header
app.use('/message/*', async (c, next) => {
  await next()
  c.header('x-message', 'This is middleware!')
})

app.get('/message/hello', (c) => c.text('Hello Middleware!'))
```

# helpers 

helper help in assiting in developing your application.
they just provide usfull functions

```javascript
import { getCookie, setCookie } from 'hono/cookie'

const app = new Hono()

app.get('/cookie', (c) => {
  const yummyCookie = getCookie(c, 'yummy_cookie')
  // ...
  setCookie(c, 'delicious_cookie', 'macha')
  //
})
```









