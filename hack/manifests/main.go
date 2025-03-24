package main

import (
	"os"
)

func main() {
	switch os.Args[1] {
	case "cleancrd":
		cleanCRD(os.Args[2])
<<<<<<< HEAD:hack/main.go
	case "removecrdvalidation":
		removeCRDValidation(os.Args[2])
	case "docgen":
		generateDocs()
=======
	case "minimizecrd":
		minimizeCRD(os.Args[2])
>>>>>>> draft-3.6.5:hack/manifests/main.go
	default:
		panic(os.Args[1])
	}
}
