All go code could only put under a package.

Every go program should contain a main package and a main function. 

The main function is the entry for the whole program, and the will be first run.

Package name and function name won't conflict.

The unit of processing in go is package rather than file, that means we could 
separate a package into any number of files.From the compiler's view, these 
files belong to the same package if they have the same package declaration, it
the same as put all contents of these files in a single big file.