Factory method is a creational design pattern that provides an interface for creating objects in a supperclass,
but allows subclasses alter the type of objects that will be created. 

* Problem: Logistics management application.
    First version can only handle transportation by trucks => bulk of code live inside `Truck` class.
    After a while, you want to support more requests of transportation by sea, air => adding `Ship`, `AirPlan`.
=> your codebase already coupled to `Truck` class, this require making change to entirely codebase.
=> As a result, you will end up with pretty nasty code, riddle with conditionals that switch app's behavior depending 
pn the class of transportation objects.

* Solution:

