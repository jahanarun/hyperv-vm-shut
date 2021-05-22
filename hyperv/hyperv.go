package hyperv

import (
	"fmt"

	"github.com/jahanarun/hyperv-vm-shut/common"
	"github.com/jahanarun/hyperv-vm-shut/powershell"
)

type VirtualMachine struct {
	Name string
	State string
	Status string
}

func GetVM(vmname string) (VirtualMachine, error) {
	posh := powershell.New()
	cmd := fmt.Sprintf("Get-VM -VMName \"%s\" | fl -Property State, Status, Name", vmname)
	stdOut, stdErr, err := posh.Execute(cmd)
	fmt.Printf("stderr: %s, *****err:%s", stdErr, err)
	if stdErr != "" || err != nil {
		return VirtualMachine{}, err
	}

	vm := VirtualMachine{
		State: common.GetValue(stdOut, "state"),
		Status: common.GetValue(stdOut, "status"),
		Name: common.GetValue(stdOut, "name"),
	}	
	return vm, err
}

func SaveVM(vmname string) {
	cmd := fmt.Sprintf("Save-VM -VMName \"%s\"", vmname)
	posh := powershell.New()
	posh.Execute(cmd)
}
