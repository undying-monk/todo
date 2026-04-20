# API DESIGN

## Context
Use to define concrete APIs for system design after modeling and collecting requirements

## Sample
Let say you have a system design for OMS, that requies a few abilities for ordering, upsert card,
we need a several entities such as Orders, OrderItems, Products, Inventories, Cards,

We can define API design as following:
- POST /orders
- GET /orders/:id
- PUT /cards
- GET /cards
