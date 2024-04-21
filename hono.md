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
