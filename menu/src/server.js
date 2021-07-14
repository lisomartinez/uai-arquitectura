import express from 'express'
import { json, urlencoded } from 'body-parser'
import restaurantRoutes from './routes/restaurant'
import { connect } from './utils/db'

export const app = express()

app.use(json())
app.use(urlencoded({ extended: true }))

app.get('/', (req, res) => {
  res.sendStatus(200)
})

app.use('/restaurants', restaurantRoutes)

export const start = async () => {
  try {
    await connect()
    app.listen(3000, () => {
      console.log(`REST API on http://localhost:3000/`)
    })
  } catch (e) {
    console.error(e)
  }
}
