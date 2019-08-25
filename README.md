[![Go Report Card](https://goreportcard.com/badge/github.com/cnnrrss/peter-piper)](https://goreportcard.com/report/github.com/cnnrrss/peter-piper)
[![GoDoc](https://godoc.org/github.com/cnnrrss/peter-piper?status.svg)](https://godoc.org/github.com/cnnrrss/peter-piper)


# Peter Piper

Implement the Knuth-Morris-Prat string matching algorithm in Go to complete the match subtext to a body of text.

### Installation

##### Install package
`go get -u github.com/cnnrrss/peter-piper`

##### Verify test cases pass
`go test`
**or** for verbose output..
`go test -v`

##### Verify all code is covered
`go test -cover`

### Implementation
The easiest way to solve this task would probably be some combination of the features from Go standard packages such as `strings` or `regexp`, but I assumed that was not the intent of the exercise.

Therefore I thought about the available string matching algorithms to choose from: **Robin-Karp**, **Knuth-Morris-Pratt**, and **Boyer-Moore**. I chose to implement a modified version of the **Knuth-Morris-Pratt** because it provides a Time Complexity of **O(n + k)** and was easier to implement in the 1 - 2 hr activity window. Another potential solution would have been to implement a **Trie** / **Radix Tree** data structure.

#### Dependencies
Although I did not import any external packages to complete the task, I did rely on the **strings** package to update the input text to lowercase, satisfying the case insensitive requirement. I also imported the **reflect** and **testing** std packages to facilitate the test cases.
