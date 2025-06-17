// Install Docker

package cmd

import "fmt"

func installDocker() {
	fmt.Println("Removing old docker packages")
	runCommand("for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done")

	fmt.Println("Setting up Docker's apt repository")
	commands := []string{
		"sudo apt-get -y update",
		"sudo apt-get install ca-certificates curl",
		"sudo install -m 0755 -d /etc/apt/keyrings",
		"sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc",
		"sudo chmod a+r /etc/apt/keyrings/docker.asc",
		"echo \"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo \"${UBUNTU_CODENAME:-$VERSION_CODENAME}\") stable\" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null",
		"sudo apt-get -y update",
	}

	for _, cmd := range commands {
		runCommand(cmd)
	}

	fmt.Println("Installing Docker Packages")
	runCommand("sudo apt-get -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin")

	fmt.Println("Adding user to docker group")
	runCommand("sudo usermod -aG docker $USER")

	fmt.Println("Restarting docker service")
	runCommand("sudo systemctl restart docker")

	fmt.Println("You need to log out and log back in for the changes to take effect")
}
