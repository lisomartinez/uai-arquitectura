import { Restaurant } from '../models/restaurant'

const getOne = async (req, res) => {
  try {
    const doc = await Restaurant.findOne({ restaurant_id: req.params.id })
      .lean()
      .exec()

    if (!doc) {
      return res.status(400).end()
    }

    res.status(200).json({ restaurant_id: doc.restaurant_id, name: doc.name })
  } catch (e) {
    console.error(e)
    res.status(400).end()
  }
}

export default { getOne }
