import mongoose from 'mongoose'

export const connect = (
  url = process.env.AZ_COSMOS_MENU_CONNECTION_STRING,
  opts = {}
) => {
  return mongoose.connect(url, { ...opts, useNewUrlParser: true })
}
