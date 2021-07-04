package com.github.lisomartinez.users.controller

import com.github.lisomartinez.users.model.User
import com.github.lisomartinez.users.service.UserService
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RestController

@RestController
class UserController(val userService: UserService) {

    @GetMapping("/users/{userId}")
    fun getUserById(@PathVariable userId: Long): User {
        return userService.getUser(userId)
    }
}