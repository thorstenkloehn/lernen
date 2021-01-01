# Windows

## Einstellung

```

setx CARGO_HOME = e:\CARGO_HOME
setx RUSTUP_HOME = e:\RUSTUP_HOME


```


## Grundlagenn Programme

### Anti Virus Programmieren

* [Bitdefender](https://login.bitdefender.com/central/login.html?lang=de_DE&redirect_url=https:%2F%2Fcentral.bitdefender.com%2Factivity%3FbrowserLang%3Dde_DE)


### Browser


* [Microsoft Edge](https://www.microsoft.com/en-us/edge)
* [Google Chrome](https://www.google.de/chrome)
* [Firefox](https://www.mozilla.org/de/firefox/developer)

### Versionsverwaltungssoftware

* [Git Installieren](https://git-scm.com)

### Allgemeine Module


* [Msys2](https://www.msys2.org/)
* [MSYS2 Packages](https://packages.msys2.org/updates)

#### MSYS2 Console eingeben

``` 
pacman -S mingw-w64-x86_64-toolchain base-devel mingw-w64-x86_64-cmake 

```

## Programmiersprachen

* [Go Installieren](https://golang.org)
* [Nodjes Installieren](https://nodejs.org/en/download/)
* [Python](https://www.python.org/downloads/)
* [Java](https://aws.amazon.com/de/corretto/)


## Java

#### Java Build
##### Gradle Installieren

###### Test


```

java -v 


```

###### Herunderladen

* [Gradle Installieren Anleitung](https://gradle.org/install/)

###### Einstellung

```

set GRADLE_HOME = PATH



```

##### Maven Installieren

###### Test

```
java -v

```

###### Herunderladen

* [Maven Download](http://maven.apache.org/download.cgi)

###### Einstellung

```

set M2_HOME = PATH

```

## IDE

### Software


* [IntelliJ IDEA Ultimate](https://www.jetbrains.com/idea)
* [Clion](https://www.jetbrains.com/clion/)




#### IntelliJ IDEA Einstellung

```

# custom IntelliJ IDEA properties

idea.config.path = e:/work/idea/config/path
idea.system.path = e:/work/idea/system/path
idea.plugins.path = e:/work/idea/plugins/path
idea.log.path = e:/work/idea/log/path
```
#### Clion Einstellung

```

idea.config.path = e:/clion/work/idea/config/path
idea.system.path = e:/clion/work/idea/system/path
idea.plugins.path = e:/clion/work/idea/plugins/path
idea.log.path = e:/clion/work/idea/log/path

```



## Datenbank

* [postgresql](https://www.postgresql.org/download/)
* [Neo4j](https://neo4j.com/download-center/?ref=web-product-database/#community)
* [MongoDB]()
* [Apache Cassandra]()