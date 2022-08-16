# Ref from
    - https://levelup.gitconnected.com/
    - https://dave.cheney.net/2016/08/20/solid-go-design
# Single responsibility Principle
    - A piece of code / Class / module should have one and only one reason to change beacuse If the code has a single responsibility
    - It has the fewest reason to change
    One Class or Module should be responsible for a single stack-holder
# Example 1
    - Example 1 shows that how a simple program can voilating  the SRP
    - Because It is having number of reason to change or modify the existing code, means it is a bad design, and we have to go for implement using good design
        - 1. If we want to put in the another form lets say in JSON instead of Text we need to change
        - 2. If we want to find out the area of another shape still we need to change the code
    - So a piece of code which follows the SRP is implemented into example 2
