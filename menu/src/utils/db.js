import mongoose from 'mongoose'

let connection = `${process.env['AZ_COSMOS_MENU_CONNECTION_STRING']}`

export const connect = (
  url = connection.replace('AZ_COSMOS_MENU_CONNECTION_STRING=', ''),
  opts = {}
) => {
  console.log('connection string:', url)
  return mongoose.connect(url, { ...opts, useNewUrlParser: true })
}
