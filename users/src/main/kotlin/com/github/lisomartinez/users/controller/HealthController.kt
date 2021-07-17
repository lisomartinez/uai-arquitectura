package com.github.lisomartinez.users.controller

import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class HealthController {

    @GetMapping("/health")
    fun getUserById(): ResponseEntity<Void> {
        return ResponseEntity.ok().build()
    }
}
