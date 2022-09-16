package main

import "fmt"

// general interface for cilent to use
type computerAdapter interface {
	getOS() string
}

// first instance that needs "adapting"
type windowsComputer struct{}

func (winCom windowsComputer) getWindows() string {
	return "windows"
}

// first adapter, implements "computerAdapter" interface
type WindowsAdapter struct {
	winComp windowsComputer
}

func NewWindowsAdapter() *WindowsAdapter {
	return &WindowsAdapter{
		winComp: windowsComputer{},
	}
}

func (winAdapt WindowsAdapter) getOS() string {
	return winAdapt.winComp.getWindows()
}

// second instance that needs "adapting"
type linuxComputer struct{}

func (linCom linuxComputer) getLinux() string {
	return "linux"
}

// second adapter, implements "computerAdapter" interface
type LinuxAdapter struct {
	linCom linuxComputer
}

func (linAdapt LinuxAdapter) getOS() string {
	return linAdapt.linCom.getLinux()
}

func NewLinuxAdapter() *LinuxAdapter {
	return &LinuxAdapter{
		linCom: linuxComputer{},
	}
}

// example func, that client would use
func printOS(comp computerAdapter) {
	fmt.Println(comp.getOS())
}

func main() {
	win := NewWindowsAdapter()
	lin := NewLinuxAdapter()

	fmt.Print("Windows getOS: ")
	printOS(win)

	fmt.Print("Linux getOS: ")
	printOS(lin)
}
