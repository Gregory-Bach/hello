package gitscripts

import (
	"fmt"
	"os/exec"
)

func GitUpdate(hostname string) int {

	dir := ""

	if hostname == "deathstar" {
		dir = "/mnt/c/Users/grego/Documents/my-notes"
	} else if hostname == "Gregorys-Air" {
		dir = "/Users/greg/Documents/my-notes"
	} else {
		fmt.Println("Unknown or no hostname: ", hostname)
		fmt.Println("Please config correctly...")
		return 2
	}

	fmt.Println("")
	fmt.Println("**********************")
	fmt.Println("***   GIT UDPATE   ***")
	fmt.Println("**********************")

	fmt.Println("Updating git repos...")

	errAdd := GitAdd(dir)
	if errAdd == 1 {
		return 1
	}

	errCommit := GitCommit(dir)
	if errCommit == 1 {
		return 1
	}

	errPull := GitPull(dir)
	if errPull == 1 {
		return 1
	}

	errPush := GitPush(dir)
	if errPush == 1 {
		return 1
	}

	return 0
}

func GitAdd(dir string) int {
	fmt.Println("Git Add...")
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error on git add: ", err)
		return 1
	} else {
		fmt.Println("Git Add successfully.")
		return 0
	}
}

func GitCommit(dir string) int {
	fmt.Println("Git Commit...")
	cmd := exec.Command("git", "commit", "-m", "Update from Hello")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error on git commit: ", err)
		return 1
	} else {
		fmt.Println("Git Commit successfully.")
		return 0
	}
}

func GitPull(dir string) int {
	fmt.Println("Git Pull...")
	cmd := exec.Command("git", "pull")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error on git pull: ", err)
		return 1
	} else {
		fmt.Println("Git Pull successfully.")
		return 0
	}
}

func GitPush(dir string) int {
	fmt.Println("Git Push...")
	cmd := exec.Command("git", "push")
	cmd.Dir = dir
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error on git push: ", err)
		return 1
	} else {
		fmt.Println("Git Push successfully.")
		return 0
	}
}
