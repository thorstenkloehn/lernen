# Aufzählungstypen

## Enum



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

