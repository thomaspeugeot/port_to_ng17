package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/tidwall/sjson"
)

var (
	allCommands = flag.Bool("allCommands", true, "generates allCommands")
)

func main() {

	flag.Parse()
	if len(flag.Args()) > 1 {
		log.Fatal("surplus arguments")
	}

	if false {

		{
			// Read the contents of the tsconfig.json file
			file, err := ioutil.ReadFile("tsconfig.json")
			if err != nil {
				fmt.Println(err)
				return
			}

			// Modify the skipLibCheck field
			file, err = sjson.SetBytes(file, "compilerOptions.skipLibCheck", true)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Write the updated JSON back to the tsconfig.json file
			err = ioutil.WriteFile("tsconfig.json", file, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	}
	if *allCommands {
		cmd := exec.Command("rm", "-rf", "package-lock.json", "node_modules", "dist")

		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))

	}
	if *allCommands {
		cmd := exec.Command("npm", "install")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}
	if *allCommands {
		cmd := exec.Command("npx", "@angular/cli@16", "update", "@angular/core@16", "@angular/cli@16", "--allow-dirty", "--force")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}
	if *allCommands {
		cmd := exec.Command("npm", "install", "--save", "@angular-material-components/datetime-picker@16", "--force")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}
	if *allCommands {
		cmd := exec.Command("npx", "@angular/cli@16", "update", "@angular/material@16", "--allow-dirty")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}
	if *allCommands {
		cmd := exec.Command("npm", "install", "typescript@5", "--save-dev")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}
	if *allCommands {
		cmd := exec.Command("npm", "install", "angular-split@16")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}

	if *allCommands {
		cmd := exec.Command("ng", "generate", "@angular/material:mdc-migration", "--interactive=false", "--components=all")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}

	if *allCommands {
		cmd := exec.Command("ng", "build")
		var stdBuffer bytes.Buffer
		mw := io.MultiWriter(os.Stdout, &stdBuffer)

		cmd.Stdout = mw
		cmd.Stderr = mw

		log.Println(cmd.String())
		log.Println(stdBuffer.String())

		// Execute the command
		start := time.Now()
		if err := cmd.Run(); err != nil {
			log.Panic(err)
		}
		log.Printf("command %s over and took %s", cmd, time.Since(start))
	}
}
