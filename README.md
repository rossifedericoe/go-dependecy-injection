# go-dependecy-injection
Pet project, testing wire

# Using
* Gin Framework (https://github.com/gin-gonic/gin)
* Google Wire (https://github.com/google/wire)

# Domain
There's "Car" and "Bike" entities, and each domain defines an interface: CarService and BikeService.

# DI
The gin engine is wrapped into a struct that receives all the dependencies resolved from wire.
