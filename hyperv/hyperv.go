package hyperv

import (
	"fmt"

	"github.com/jahanarun/hyperv-vm-shut/common"
	"github.com/jahanarun/hyperv-vm-shut/powershell"
)

type VirtualMachine struct {
	state string
	status string
}

func GetVM(vmname string) *VirtualMachine {
	posh := powershell.New()
	cmd := fmt.Sprintf("Get-VM -VMName \"%s\" | fl -Property State, Status", vmname)
	stdOut, _, _ := posh.Execute(cmd)
	// Get-VM -VMName dex-adfs | fl -Property state,status
	// fmt.Printf("\nEnableHyperV:\nStdOut : '%s'\nStdErr: '%s'\nErr: %s", strings.TrimSpace(stdOut), stdErr, err)
	cleanedResult := common.ResponseCleanup(stdOut)
	fmt.Println("cleanedResult:", cleanedResult)
	// vm := VirtualMachine{}
	vm := new(VirtualMachine)
	vm.state = common.GetValue(cleanedResult, `state:(?P<state>.*)`)
	vm.status = common.GetValue(cleanedResult, `status:(?P<status>.*)`)
	
	fmt.Println("vm.state:", vm.state)
	fmt.Println("vm.status:", vm.status)
	fmt.Println("vm:", vm)
	return vm
}