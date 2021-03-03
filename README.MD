# OnQueue
### basic implementation of a message queue using rabbitmq

##### Goal: Produce an email message that is published to the message broker (queue) where it is safely stored until a Consumer is available to take it off the queue and work it


# Architecture/Patterns
***
##### this is section is used to describe patterns/architecture that is commonly used throughout the project.
* `seperate models for database entities and HTTP responses` keeps things DRY and reduces risk of exposing private data not intended for the outside world.

* **Repository Pattern**
  * separate concrete implementation with an interface! Define the contract that ANY database can use, and once a database (like Postgres) has been decided on, use it by implementing the interface
* **Port** is an ***interface***  an input that defines how something from the outside can use the things on the inside of the application. Typically, inside the business logic
* **Adapter** is how the application talks to the external world. `you have to adapt your applications structures to what the external world expects i.e (pub/sub, sql queries, stripe, etc`
* **Application Logic** is a small layer "the glue" of the layers. In ***Clean Architecture*** this is the `Use Case` Layer
* **Domain** business logic **NO DEPENDENCIES**
  * example: `You can control your smartphone using the physical buttons, the touchscreen, or voice assistant. Whether you press the “volume up” button, swipe the volume bar up, or say “Siri, volume up”, the effect is the same. There are several entry points (ports) to the “change volume” logic.
    When you play some music, you can hear it coming from the speaker. If you plug in headphones, the audio will automatically change to them. Your music app doesn’t care. It’s not talking with the hardware directly, but using one of the adapters the OS provides.
    Can you imagine creating a mobile app that has to be aware of the headphones model connected to the smartphone? Including SQL queries directly inside the application logic is similar: it exposes the implementation details.`

### Terms
***
**Domain First**
: Depending on the complexity of the project, weeks can be spent ironing out the domain layer.
I always recommend using just basic inmem database implementations because it is easy and just an implementation detail

**Domain Layer**
: knows nothing about any layer above it, the innermost circle in the onion, the core of ports and adapters.
Pure business logic. Think of the objects here not as ORM or model specific data but **types with behavior**!
i.e. `my stove is set to turn on at 5` ***NOT*** `setting the startAt attribute on the oven to 5`


