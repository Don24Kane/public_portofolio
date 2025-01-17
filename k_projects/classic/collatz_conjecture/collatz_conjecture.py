# by Don24Kane

# Python

n = int(input("enter your number: "))

steps = 0

while n > 1:
    if n % 2 == 0:
        n //= 2
    else:
        n = n * 3 + 1
    steps += 1

print("the number of steps it took for n to reach the value of 1 is:", steps)
