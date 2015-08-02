package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintln(os.Stderr, "Usage: dir605reboot ip username password")
		os.Exit(1)
	}
	ip := os.Args[1]
	username := os.Args[2]
	password := os.Args[3]

	fmt.Println("Login...")
	data := url.Values{
		"ACTION_POST":       {"LOGIN"},
		"FILECODE":          {""},
		"VERIFICATION_CODE": {""},
		"LOGIN_USER":        {username},
		"LOGIN_PASSWD":      {password},
		"VER_CODE":          {""},
	}
	resp1, err := http.PostForm("http://"+ip+"/login.php", data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to login:", err)
		os.Exit(1)
	}
	defer resp1.Body.Close()
	if resp1.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "Failed to login:", resp1.Status)
		os.Exit(1)
	}

	fmt.Println("Reboot...")
	resp2, err := http.Get("http://" + ip + "/sys_cfg_valid.xgi?&exeshell=submit%20REBOOT")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to reboot:", err)
		os.Exit(1)
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "Failed to reboot:", resp2.Status)
		os.Exit(1)
	}
}
