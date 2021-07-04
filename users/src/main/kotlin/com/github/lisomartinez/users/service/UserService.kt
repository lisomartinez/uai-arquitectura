package com.github.lisomartinez.users.service

import com.github.lisomartinez.users.exceptions.UserNotFoundException
import com.github.lisomartinez.users.model.User
import com.github.lisomartinez.users.repository.UserRepository
import org.springframework.stereotype.Service

@Service
class UserService(private val userRepository: UserRepository) {
    fun getUser(userId: Long): User {
        return userRepository.findByUserId(userId) ?: throw UserNotFoundException(userId)
    }
}