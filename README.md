# Sainsbury's GopherCon Challenge 2018

One of our delivery riders has been given some very confusing instructions. On a (seemingly endless) grid of houses, he has to deliver to each house he arrives at.

His instructions are given as a string consisting of one of four characters:
* ^ to move north one square
* \> to move east one square
* v to move south one square
* < to move west one square

To the bemusement of the rider, this means some houses get more than one delivery.

## The challenge

Complete the `CalculateHouses(input string) int` function to work out how many **unique** houses the rider visits.

**This also includes the house he starts at.**

## Completion

The challenge is complete when all the unit tests pass when run with `ginkgo` or `go test` (both work).

## Examples

If our rider is given `>>`, he visits three houses:
* The house he starts at, before moving east
* The next house, before moving east again
* The final house

If our rider is given `^>v<`, he visits four houses:
* The house he starts at, before moving north
* The next house, before moving east
* The next house, before moving south
* The next house, before moving west 
* This brings him back to the house he started at, where he has already made a delivery
