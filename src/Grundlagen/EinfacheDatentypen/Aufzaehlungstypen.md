# Aufzählungstypen

## Enum

C++

```cpp

#include <iostream>

enum JAHRZEITEN {FRUEJAHR,SOMMER,HERBST,WINTER};

using namespace std;
int main() {

    JAHRZEITEN jetzt;
    jetzt = SOMMER;
    switch (jetzt) {
        case  FRUEJAHR: {
             cout << "Frühjahr" << endl;
             break;
        }
        case SOMMER: {
            cout << "Sommer" << endl;
            break;
        }

        case HERBST: {
            cout << "Herbst" << endl;
        }

        case WINTER: {
            cout << "Winter" << endl;
        }

        default: {
            cout << "Keine Daten Vorhanden" << endl;
        }

    }
    return 0;
}
```

### JAVA

```java
package com.company;

public class Main {

    public static enum JAHREZEIT {
        FRUEJAHR,
        SOMMER,
        HERBST,
        WINTER
    }
    public static void main(String[] args) {

        JAHREZEIT jetzt;
        jetzt = JAHREZEIT.HERBST;
        switch (jetzt) {

            case FRUEJAHR: {
           System.out.println("Frühjahr");
           break;
            }
            case SOMMER: {

                System.out.println("Sommer");
                break;
            }
            case HERBST: {
                System.out.println("Herbst");
                break;
            }

            case WINTER: {
                System.out.println("Winter");

            }
            default:{
                System.out.println("Keine Daten");
                break;
            }

        }
    }
}


```

