- When we want to retrieve a specified element from a collection directly, without
looping to look for it, a map is a good choice of data structure.Declared like 

    ```go
        myMap := make(map[int]string)
        myMap :=map[int]string{
            1:"a",
            2:"b",
        } 
    ```

- A map links keys, such as book IDs, to elements, and a map literal comprises a list
of key‐element pairs.

- Looking up a map element by key uses the same square‐bracket notation as looking
up a slice element by index: for example catalog[ID].

- We can also assign a new element to a specified map key; if the key already exists,
its old element will be replaced with the new one, and if it doesn’t exist, the new
element will be added to the map.

- The result of looking up a key that doesn’t exist in the map is the zero value for the
element type: for example, with a map of Book, if you look up some key that’s not
in the map, you’ll get an empty Book value.

- If we receive a second ok value from the map lookup, it will tell us whether the
key actually existed (true), or whether we’re just getting a default zero value instead
(false).

    ```go
        book, ok := catalog[ID]
        if !ok {
            return Book{}, fmt.Errorf("no book with id %d found", id)
        }
    ```

- And to make the error message more helpful, we can use fmt.Errorf to interpolate
some data into it, such as the ID the user asked for.

- If we want a slice containing all the keys of some map, we can loop over the slice
with for ... range, appending each key to our result slice.

- By design, maps just don’t have any defined key order.

- To simplify testing functions that return items from a map, we can use
sort.Slice to sort the results so that we can compare them reliably.

    ```go
        sort.Slice(got, func(i, j int) bool {
                return got[i].ID < got[j].ID
        })
    ```