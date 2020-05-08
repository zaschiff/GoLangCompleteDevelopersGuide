package main

import "fmt"

func main() {
	fmt.Println("Hi there!")
}

/*
   five basic questions:

        How do we run the code in our project?
            using the terminal/command prompt
                1. Navigate to the project directory.
                2. Use the go command; its a portal to all things Go. Gives us the ability to compile and execute projects we put together.
                3. To run the file, type go run 'file name'; ie. go run hello.go

                Go Commands -> CLI (Command Line Interface)
                    go build   -> compiles a bunch of go source code files
                    go run     -> compiles and executes one or more files
                    go fmt     -> formats all the code in each file in the current directory
                    go install -> compiles and "installs" a package
                    go get     -> downloads the raw source code of someone else's package
                    go test    -> runs any tests associated with the current project.

        What does 'package main' mean?

            package is a project or workspace; Collection of common source code files.
            First Line of file must declare the package that the file belongs to.

            2 different types of packages in go
                Executable -> generates a file we can run when compiled; used for doing something.
                Reusable   -> code used as 'helpers'. Good place to put reusable logic; dependencies or libraries

            how do we know it is one or the other?
                Name of package determines the type of package. Compiling a non-main package gives nothing.
                "main" is special! Any package called main must have a function called main.

        What does 'import "fmt"' mean?

            import is used to give our package access to some code written in another package.
            ftm is shortened form of format, ie the format package.
            import essentialy forms a link form our main package, which by defaults has not connecitons, to the one we need access to.
            import can give access to packages made by other authors, not just those in the standard library packages.
            https://golang.org/pkg has a list of all packages included in the standard library.

        What does 'func' mean?

            func is short for function. Functions is just like functions in other languages.
                using func main() {}:

                    func -> keyword to declare that the following code is a function
                    main -> sets the name of the function
                    ()   -> list of arguments to pass the function
                    {}   -> function body, calling the function runs the code between the curly braces.

        How is the hello.go file organized?

            always the same pattern inside of any go file.
                package declaration                     -> package main
                import other packages that we need      -> import "fmt"
                declare functions, tell Go to do things -> func main() { fmt.Println("hi there!")}

*/
