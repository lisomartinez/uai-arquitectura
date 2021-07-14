import * as mongoose from 'mongoose'

const restaurantSchema = new mongoose.Schema({
  restaurant_id: {
    type: Number,
    required: true,
  },
  name: {
    type: String,
    required: true,
    trim: true,
  },
})

restaurantSchema.index({ restaurant_id: true }, { unique: true })

export const Restaurant = mongoose.model('restaurant', restaurantSchema)
