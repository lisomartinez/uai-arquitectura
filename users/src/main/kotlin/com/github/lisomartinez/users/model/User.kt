package com.github.lisomartinez.users.model

import com.fasterxml.jackson.annotation.JsonIgnore
import javax.persistence.*

@Entity
@Table(name = "users")
class User(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @JsonIgnore
    @Column(name = "id")
    val id: Long?,

    @Column(name = "user_id", unique = true)
    val userId: Long,

    @Column(name = "first_name")
    val firstName: String,

    @Column(name = "last_name")
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