package src

import (
	"flag"
	"fmt"
	"os"
)

func InitCli() (string, int, int, bool, bool, bool) {

	var (
		routines        int
		maxRequestSize  int
		smooth          bool
		detach          bool
		commonPathsOnly bool
		help            bool
	)

	flag.IntVar(&routines, "r", 50, "Number of concurrent requests to make")
	flag.IntVar(&maxRequestSize, "m", 10, "Set the maximum size of the request tested")
	flag.BoolVar(&smooth, "s", false, "Enable smooth request rate by adding a delay between requests")
	flag.BoolVar(&detach, "d", false, "Detaches the output to a file and releases the terminal")
	flag.BoolVar(&commonPathsOnly, "c", false, "Limit search to common paths in the words list")
	flag.BoolVar(&help, "h", false, "Show help message")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	args := flag.Args()

	if len(args) < 1 {
		printErrorUrlNotSpecified()
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("Mais parametros que esperado")
		os.Exit(1)
	}

	setAlphabet("")

	urlFlag := args[0]

	if detach {
		fmt.Print("")
	}

	return urlFlag, routines, maxRequestSize, smooth, detach, commonPathsOnly
}

/*func StartDetachedProcess(cmd string, args []string, uid, gid int) error {
	cred := &syscall.Credential{Uid: uint32(uid), Gid: uint32(gid), Groups: []uint32{}}
	sysproc := &syscall.SysProcAttr{Credential: cred, Noctty: true}
	attr := os.ProcAttr{
		Dir:   ".",
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, nil, nil},
		Sys:   sysproc,
	}

	process, err := os.StartProcess(cmd, args, &attr)
	if err != nil {
		return err
	}

	if err := syscall.Setpgid(process.Pid, process.Pid); err != nil {
		return err
	}

	if err := process.Release(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := StartDetachedProcess("/bin/sleep", []string{"/bin/sleep", "100"}, 501, 100); err != nil {
		fmt.Println(err.Error())
	}
}
*/
