package main

import(
    "os"
    "fmt"

    "github.com/go-git/go-git/v5"
)

//  ?
func main() {
//  Command
    cmd = os.Args[1]
//  Location of the repository
    directory := os.Args[len(os.Args) - 1]

// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
//	CheckIfError(err)

	w, err := r.Worktree()
//	CheckIfError(err)

    switch cmd {
        case 'status':
            status, err := w.Status()
//	          CheckIfError(err)

            fmt.Println(status)
        default:
            Warning("got: %s is not a got command!", cmd)

            os.Exit(1);
    }
}
