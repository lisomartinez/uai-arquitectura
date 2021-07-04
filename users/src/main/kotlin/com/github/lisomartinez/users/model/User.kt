package com.github.lisomartinez.users.model

import com.fasterxml.jackson.annotation.JsonIgnore
import javax.persistence.*

@Entity
class User(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @JsonIgnore
    val id: Long?,

    @Column(name = "userId", unique = true)
    val userId: Long,
    val firstName: String,
    val lastName: String,
    val address: Address
) {
     constructor(
        userId: Long,
        firstName: String,
        lastName: String,
        address: Address
    ) : this(null, userId, firstName, lastName, address)
}