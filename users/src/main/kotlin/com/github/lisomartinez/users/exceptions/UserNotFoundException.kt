package com.github.lisomartinez.users.exceptions

import org.springframework.http.HttpStatus
import org.springframework.http.HttpStatus.NOT_FOUND
import org.springframework.web.bind.annotation.ResponseStatus

@ResponseStatus(code = NOT_FOUND)
class UserNotFoundException(userId: Long) : RuntimeException("user not found $userId")