
- Although ***type checking is enforced by the compiler on assignments to struct
fields, there’s no built‐in way to validate what we assign to them***.

  - So if we want ***to protect users from accidentally assigning an invalid value (for example,
  a negative price), we need to provide some accessor method that can check
  the value before assigning it***.

  - ***This would have to be a pointer method***, since without a pointer anything it assigned
  to the struct field wouldn’t persist: it would just be modifying its own local
  copy of the struct that then gets thrown away when the method returns.

  - ***in general a reciever object should be a pointer, it allows the method to modify actual value***

  - Even ***if there’s a validating accessor method, that doesn’t actually prevent users
  from setting any value they like on an exported struct field, bypassing the method
  altogether***.

  - So if we want ***to ensure that certain fields of the struct are always valid, we need
  to make those fields unexported, and only available via accessor methods***: 

    - in this case, we’ll need to provide a method to get the value as well as one to set it.

- However, ***cmp.Equal can’t look at unexported fields (nor can anyone else outside
the package), so we can pass the option cmpopts.IgnoreUnexported to tell it
not to even try.***

  - Just as you can create always valid fields by unexporting the field names, you can
  ***create always valid structs by unexporting the name of the struct itself***.

  - That implies some kind of ***constructor function is needed, because users can’t create literals of the struct (they’d need to use its name, and they’re not allowed to)***.

  - **It’s fine for constructors to return types with unexported names (indeed, it’s essential
  for always valid structs): it’s only the name users can’t refer to, not the value**.


- When you want ***to restrict values to one of an allowed set, a map of bool is a nice way to do it***, because you can write something like if ```validCategory[category]```, and the map lookup itself gives you the right answer, with no extra code.

  - Another way is to ***use the const keyword to define some named constants that users can refer to: if they misspell the name, or refer to one that doesn’t exist, the compiler will catch this***.

    - A good example is the set of HTTP status constants defined in net/http: it’s not
    inherently obvious without context what the integer literal 200 means, and you
    could get it wrong, but ***http.StatusOK*** is unequivocal, and the compiler will tell
    you if you mistyped it.

  - ***Often it doesn’t matter what the values of the constants are, only that they’re distinct,
  and in that case we can use integer constants and iota to auto‐assign values to them***.