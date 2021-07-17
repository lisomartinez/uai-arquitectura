# HACER EXPORTS EN LA TERMINAL (en la misma que vas a ejecutar el comando)


# SQL Usuarios

``` bash
export AZ_DATABASE_NAME=
export AZ_SQL_SERVER_PASSWORD=
```

# COSMOS - Order

```bash
export AZ_COSMOS_ORDER_CONNECTION_STRING=
```

# COSMOS - MENU / RESTAURANTS
```bash
export AZ_COSMOS_MENU_CONNECTION_STRING=
```
# EJECUTAR LOCAL

menu: npm run dev
orders: go run ./cmd/main.go
users: ./gradlew bootRun

# DOCKER


```bash
docker login --username uaiarch --password W7Yt+8qXrK6JEssZMrHH6XGlfDB14gPc uaiarch.azurecr.io


docker build --build-arg AZ_COSMOS_MENU_CONNECTION_STRING=$AZ_COSMOS_MENU_CONNECTION_STRING  --tag uaiarch.azurecr.io/menu:1 .

docker push uaiarch.azurecr.io/menu:1


docker build --build-arg AZ_COSMOS_ORDER_CONNECTION_STRING=$AZ_COSMOS_ORDER_CONNECTION_STRING --tag uaiarch.azurecr.io/orders:1 .


docker push uaiarch.azurecr.io/orders:1


docker build --build-arg AZ_DATABASE_NAME=$AZ_DATABASE_NAME --build-arg AZ_SQL_SERVER_PASSWORD=$AZ_SQL_SERVER_PASSWORD --tag uaiarch.azurecr.io/users:1 .

docker push uaiarch.azurecr.io/users:1

```

# TABLA

``` sql
create table users (
    id bigint IDENTITY(1,1) PRIMARY KEY,
    user_id bigint unique,
    first_name varchar(255),
    last_name varchar(255),
    address_street varchar(255),
    address_number varchar(255),
    address_city varchar(255),
    address_country varchar(255)
);

SELECT * from users;

INSERT INTO users(user_id, first_name, last_name, address_street, address_number, address_city, address_country)
VALUES (1, 'Pedro', 'Sanchez', 'arenales', '2112', 'caba', 'argentina')
```

# RESTAURANT - COSMOS


## restaurants

``` json
{
  "restaurant_id": 1,
  "name": "lo de pepe"
}
```

## menu
``` json
{
    "restaurant_id" : 1,
    "items" : [
        {
            "id" : 1,
            "name" : "Sopa huertana",
            "ingredients" : [
                "Zanahoria",
                "Puerro",
                "Patata mediana",
                "Cebolla"
            ]
        },
        {
            "id" : 2,
            "name" : "Croquetas de pollo y jamón",
            "ingredients" : [
                "Harina de trigo",
                "Pechuga de pollo",
                "Jamón serrano",
                "Fondo de carne"
            ]
        },
        {
            "id" : 3,
            "name" : "Ratatouille",
            "ingredients" : [
                "Berenjena",
                "Calabacín",
                "Pimiento rojo",
                "Cebolla"
            ]
        },
        {
            "id" : 4,
            "name" : "Albóndigas en salsa española",
            "ingredients" : [
                "Carne picada de ternera",
                "Tomate frito casero",
                "Extracto de carne",
                "Zanahoria rallada"
            ]
        }
    ]
}
```

