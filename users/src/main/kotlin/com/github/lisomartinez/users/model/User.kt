package com.github.lisomartinez.users.model

import javax.persistence.*

@Entity
class User(
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    val id: Long,

    @Column(name = "userId", unique = true)
    val userId: Long,
    val firstName: String,
    val lastName: String,
    val address: Address
) {

}