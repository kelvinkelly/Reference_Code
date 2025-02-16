### 1. Basic Hello World
```python
# This line prints "Hello, World!" to the console
print("Hello, World!")

# This function takes a name as an argument and returns a greeting
def greet(name):
    return f"Hello, {name}!"

# Calling the function with the argument "Alice"
print(greet("Alice"))

# This block opens a file named 'example.txt' in read mode
# and prints its content to the console
with open('example.txt', 'r') as file:
    content = file.read()
    print(content)

# This block opens a file named 'output.txt' in write mode
# and writes a sample text to it
with open('output.txt', 'w') as file:
    file.write("This is a sample text.")

# This class defines a Dog with a name and age
class Dog:
    def __init__(self, name, age):
        self.name = name
        self.age = age

    # This method returns a barking sound
    def bark(self):
        return f"{self.name} says woof!"

# Creating an instance of Dog named "Buddy" who is 3 years old
my_dog = Dog("Buddy", 3)
# Calling the bark method on the instance
print(my_dog.bark())

# This list comprehension creates a list of squares of numbers from 0 to 9
squares = [x**2 for x in range(10)]
print(squares)

# This dictionary stores student names as keys and their grades as values
student_grades = {"Alice": 85, "Bob": 92, "Charlie": 78}

# This loop iterates over the dictionary items and prints each student's grade
for student, grade in student_grades.items():
    print(f"{student}: {grade}")

