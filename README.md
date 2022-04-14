# gods
Library of generic data structures for Go.


## Install
```
$ go get github.com/MehdiEidi/gods
```

## Import
```
import "github.com/MehdiEidi/gods"
```

## Usage
Either use convenience helper functions in gods package like this:<br>
Example: Creating a new set:
```
s := gods.NewSet[string]()

s.Add("awesome :D")
```
or directly import the intended data structure and use its factory function:
```
import "github.com/MehdiEidi/set"
.
.
.

s := set.New[string]()
s.Add("again awesome :D")
```

## Supported Data Structures
* [bitset](./bitset/)
* Linked List
  * [singly linked list](./linkedlist/doubly/)
  * [doubly linked list](./linkedlist/doubly/)
  * [circularly linked list](./linkedlist/circularly/)
* [set](./set/)
  
  --- more coming soon :D ---