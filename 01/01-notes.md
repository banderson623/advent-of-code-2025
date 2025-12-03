Oh i like how golang does input/string parsing with 
```go
	lines := strings.Split(strings.TrimSpace(input), "\n")
  for _, line := range lines {
    _, err := fmt.Sscanf(line, "%1s%d", &direction, &steps)
  }
```



Okay step 2 got me pretty good. I was trying to be too fancy with the counts of the dial.
The strategy I picked was simple, it was two step:
1. for numbers larger than 100, divide it by 100 to see the whole number spins passed zero
2. take that remainder and see if that number lands on zero or passes zero.

Once I got this approach it was easy to get the correct answer
