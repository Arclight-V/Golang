# Golang

Learning Golang through the execution of buildings.

## module00/ex00
Introduction to the golang syntax using the example of mathematical calculations

### make rules

```
run    - Launch the application. Flags can be passed through argv (Mean, Median, Mode, SD) one or more
test   - Runing tests
build  - Build end Launch an applications with the transmission of test data

```

## module01/ex00
Implementation of interfaces and unmarshalling .xml, .json files

### make rules

```
help    - Оutput a list of all rules
build   - Build the ex00 application
input1  - Build the ex00 application and execute with -f inputDB/stolen_database.json
input2  - Build the ex00 application and execute with -f inputDB/stolen_database.json-f inputDB/original_database.xml
test    - Runing tests
clean   - Delete executable file

```

## module01/ex01
Comparison of structures. continued ex00

The application compares with each other .xml or .json files
### make rules

```
help    - Оutput a list of all rules
run     - Launch the application with --old inputDB/original_database.xml --new inputDB/stolen_database.json 
build   - Build the ex00 application
clean   - Delete executable file
```

## module01/ex02
Comparison of files. continued ex00

The application compares with each other huger .txt files
### make rules

```
help    - Оutput a list of all rules
run     - Launch the application with --old --old snapshot/snapshot1.txt --new snapshot/snapshot2.txt
build   - Build the compareFS application
test    - Runing tests
clean   - Delete executable file
```