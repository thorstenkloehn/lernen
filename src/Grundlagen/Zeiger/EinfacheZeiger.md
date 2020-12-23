# Einfache Zeiger

## C++ 

```cpp

#include <iostream>
int main() {

    int number = 200;
    int * adresse_number = &number;
    std::cout << *adresse_number;
    return 0;
}
 


```


## Rust

```
fn main() {

    let wert = 100;
    let  zahl = &wert;
    println!("Zahl {}",zahl);
}

```

## Go  

```go

package main

import "fmt"

func main() {


	zahl := 100
	adresse := &zahl

	fmt.Println("Hallo Leute",*adresse)
}



```