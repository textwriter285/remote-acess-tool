// go get github.com/shirou/gopsutil/v3/...

package main

import (
	"fmt"
	"os"
	"os/user"
	"runtime"

	// Importando módulos da gopsutil
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func main() {

	// ============================
	//  SISTEMA OPERACIONAL / CPU
	// ============================

	// runtime.GOOS retorna o sistema operacional
	fmt.Println("Sistema Operacional:", runtime.GOOS)

	// runtime.GOARCH retorna arquitetura (amd64, arm64, etc.)
	fmt.Println("Arquitetura:", runtime.GOARCH)

	// runtime.NumCPU() retorna número de CPUs lógicas
	fmt.Println("CPUs lógicas:", runtime.NumCPU())

	// ============================
	//  HOSTNAME e USUÁRIO
	// ============================

	// Os.Hostname pega o hostname da máquina
	hostname, _ := os.Hostname()
	fmt.Println("Hostname:", hostname)

	// Pegando o usuário atual que está rodando o programa
	currentUser, _ := user.Current()
	fmt.Println("Usuário atual:", currentUser.Username)
	fmt.Println("Home do usuário:", currentUser.HomeDir)

	// ============================
	//  INFORMAÇÕES DO SISTEMA (host)
	// ============================

	// host.Info traz dados avançados do sistema operacional
	hostInfo, _ := host.Info()

	fmt.Println("\n--- Informações do Sistema ---")
	fmt.Println("SO:", hostInfo.Platform)                 // ex: ubuntu / fedora / windows
	fmt.Println("Versão:", hostInfo.PlatformVersion)      // versão completa
	fmt.Println("Kernel:", hostInfo.KernelVersion)        // versão do kernel
	fmt.Println("Tempo ligado (uptime):", hostInfo.Uptime, "segundos")

	// ============================
	//  CPU REAL (gopsutil)
	// ============================

	cpuInfo, _ := cpu.Info()

	fmt.Println("\n--- CPU ---")
	// cpuInfo pode ter vários "cores físicos", então iteramos
	for i, c := range cpuInfo {
		fmt.Printf("CPU %d:\n", i)
		fmt.Println("  Modelo:", c.ModelName)
		fmt.Println("  Núcleos físicos:", c.Cores)
		fmt.Println("  Velocidade (MHz):", c.Mhz)
	}

	// % de uso da CPU no momento
	percent, _ := cpu.Percent(0, false)
	fmt.Println("Uso atual da CPU:", percent[0], "%")

	// ============================
	//  MEMÓRIA RAM
	// ============================

	memInfo, _ := mem.VirtualMemory()

	fmt.Println("\n--- Memória RAM ---")
	fmt.Println("Total:", memInfo.Total/1024/1024, "MB")
	fmt.Println("Usada:", memInfo.Used/1024/1024, "MB")
	fmt.Printf("Uso: %.2f%%\n", memInfo.UsedPercent)

	// ============================
	//  DISCO
	// ============================

	// "/" no Linux e macOS / "C:/" no Windows
	diskInfo, _ := disk.Usage("/")

	fmt.Println("\n--- Disco ---")
	fmt.Println("Total:", diskInfo.Total/1024/1024/1024, "GB")
	fmt.Println("Usado:", diskInfo.Used/1024/1024/1024, "GB")
	fmt.Printf("Uso: %.2f%%\n", diskInfo.UsedPercent)

	// ============================
	//  REDE (IPs, interfaces, MAC)
	// ============================

	interfaces, _ := net.Interfaces()

	fmt.Println("\n--- Interfaces de Rede ---")
	for _, iface := range interfaces {
		fmt.Println("Nome:", iface.Name)
		fmt.Println("  MAC:", iface.HardwareAddr)

		// Listando IPs da interface
		for _, addr := range iface.Addrs {
			fmt.Println("  IP:", addr.Addr)
		}
		fmt.Println()
	}

	// Fim
	fmt.Println("Coleta finalizada.")
}
