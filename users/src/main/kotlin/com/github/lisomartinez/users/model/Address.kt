package com.github.lisomartinez.users.model

import javax.persistence.Column
import javax.persistence.Embeddable

@Embeddable
class Address(
        @Column(name = "address_street") val street: String,
        @Column(name = "address_number") val number: String,
        @Column(name = "address_city") val city: String,
        @Column(name = "address_country") val country: String
) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as Address

        if (street != other.street) return false
        if (number != other.number) return false
        if (city != other.city) return false
        if (country != other.country) return false

        return true
    }

    override fun hashCode(): Int {
        var result = street.hashCode()
        result = 31 * result + number.hashCode()
        result = 31 * result + city.hashCode()
        result = 31 * result + country.hashCode()
        return result
    }
}
