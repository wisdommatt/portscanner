package portscanner

import "fmt"

func main() {
	fmt.Println("i am ready")

	portResult := ScanPort("tcp", "localhost", 80)
	// portResult := ScanPort("udp", "localhost", 519999)
	fmt.Println(portResult)

	fmt.Println("")

	portResults := ScanPorts("udp", "localhost", 5000, 5009)
	fmt.Println(portResults)
}
