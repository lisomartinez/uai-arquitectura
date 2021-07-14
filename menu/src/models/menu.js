import mongoose from 'mongoose'

const menuItemSchema = new mongoose.Schema({
  item_id: {
    type: Number,
    required: true,
    trim: true,
  },
  name: {
    type: String,
    required: true,
    trim: true,
  },
  ingredients: {
    type: [String],
  },
})

const menuSchema = new mongoose.Schema({
  restaurant_id: {
    type: Number,
    required: true,
  },
  items: {
    type: [menuItemSchema],
  },
})

menuSchema.index({ restaurant_id: 1 }, { unique: true })

export const Menu = mongoose.model('menu', menuSchema)
