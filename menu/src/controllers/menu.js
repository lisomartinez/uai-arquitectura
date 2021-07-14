import { Menu } from '../models/menu'

const getOne = async (req, res) => {
  try {
    const doc = await Menu.findOne({ restaurant_id: req.params.id })
      .lean()
      .exec()

    if (!doc) {
      return res.status(400).end()
    }

    res.status(200).json({ restaurant_id: doc.restaurant_id, items: doc.items })
  } catch (e) {
    console.error(e)
    res.status(400).end()
  }
}

export default { getOne }
