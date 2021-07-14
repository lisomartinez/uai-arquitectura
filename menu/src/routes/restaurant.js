import { Router } from 'express'
import restaurantController from '../controllers/restaurant'
import menuController from '../controllers/menu'

const router = Router()

router.route('/:id').get(restaurantController.getOne)

router.route('/:id/menu').get(menuController.getOne)

export default router
