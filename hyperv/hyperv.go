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
	vm := new(VirtualMachine)
	vm.state = common.GetValue(stdOut, "state")
	vm.status = common.GetValue(stdOut, "status")
	return vm
}