package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/google/uuid"
)

const directoryPerms = fs.FileMode(0755)

func create_site(pdf []byte) {
	site_id := uuid.NewString()
	site_path := fmt.Sprintf("./sites/%v", site_id)
	os.Mkdir(site_path, directoryPerms)

	git.PlainClone(site_path, false, &git.CloneOptions{
		URL:      "https://github.com/blakehulett7/site_template",
		Progress: os.Stdout,
	})

	parse_r(pdf, site_path)
	os.Chdir(site_path)

	cmd := exec.Command("hugo", "mod", "tidy")
	cmd.Run()
	cmd = exec.Command("hugo", "mod", "npm", "pack")
	cmd.Run()
	cmd = exec.Command("npm", "install")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
