# go-challenges

## Type Assertions in Go

In this challenge, we are going to become familiar with the concept of Type Assertions in Go!

### Preface

If you are new to the language, then type assertions are a concept that can sometimes trip you up and appear a little tricky at first, but after overcoming the syntax it becomes far easier to understand.

Through using type assertions, we can retrieve the dynamic value of an interface. For example:

```GO
var myName interface{} = "Elliot"

name := myName.(string)
fmt.Println(name)
```

In this example, we have an interface which has a dynamic value of “Elliot”. We can then use a type assertion to retrieve this dynamic value and use the value just like we would any other string value in Go.

### Challenge

In this challenge, we are going to define a function that is called GetDeveloper which will take in 2 interface{} arguments.

Within this function, you will have to declare a new Developer instance and use type assertion to populate the values correctly before then returning this new Developer instance.

---

## Satisfying Interfaces in Go

In this challenge, you are going to implement the necessary methods needed to satisfy the provided Go interface.

On the left hand screen, you have a simple Go application that features an interface called Employee.

In order to complete this challenge, you will have to complete the code and satisfy this interface.

---

## Sorting Flights by Price

In this challenge, you are going to be attempting to sort a list of Flights based on their price from highest to lowest.

You will have to implement the SortByPrice function that takes in a slice of type Flight and returns the sorted list of Flights.

In order to help you see what is going on, you have been provided a very quick printFlights function which you can use to print the flights out.

---

##
