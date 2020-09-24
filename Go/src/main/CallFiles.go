package main

import "Files"

func main() {
	Files.WriteIntoFile("temp.dat")
	Files.WriteIntoFileUsingScan("temp.dat")
	Files.ArgparamSum()
}
