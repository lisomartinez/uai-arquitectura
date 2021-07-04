package com.github.lisomartinez.users

import com.github.lisomartinez.users.model.Address
import com.github.lisomartinez.users.model.User
import com.github.lisomartinez.users.repository.UserRepository
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.CommandLineRunner
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class UsersApplication : CommandLineRunner {
    @Autowired
    private lateinit var userRepository: UserRepository

    override fun run(vararg args: String?) {
        val user = User(
            userId = 1,
            firstName = "juan",
            lastName = "perez",
            address = Address("arenales", "355", "CABA", "Argentina")
        )
        userRepository.save(user)
    }
}

fun main(args: Array<String>) {
    runApplication<UsersApplication>(*args)
}
