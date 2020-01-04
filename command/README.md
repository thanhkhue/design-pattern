### Command Design Pattern:

1. Definition.
- `Command` is a **behavioral** design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you parameterize methods with different requests. 

2. Discussion.
- `Command` decouples the object that invokes the operation from the one that knows how to perform it. To achieve this separation, the designer creates an abstract base class that maps a receiver (an object) with an action. The base class contains an execute() method that simply calls the action on the receiver.

3. Pros and Cons.
- Pros:
    + ***Single Responsibility Principle***: you can decouple classes that invoke operations from classes that perform these operations.
    + ***Open/Closed Principle***: You can introduce new commands into that app without breaking existing client code.
- Cons:
    + The code may become more complicated since you’re introducing a whole new layer between senders and receivers.
4. Relations with other patterns.

5. Source [Refactoring Guru](https://refactoring.guru/design-patterns/command)