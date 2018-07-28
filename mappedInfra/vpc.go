package mappedInfra

import (
	"github.com/thomas.obenaus/terrastate/aws"
	"github.com/thomas.obenaus/terrastate/terraform"
)

type vpc struct {
	awsVpc aws.Vpc
	tfVpc  terraform.Resource
}

func (v *vpc) Id() string {
	return v.awsVpc.Id()
}

func (v *vpc) Type() aws.Type {
	return v.awsVpc.Type()
}

func (v *vpc) String() string {
	hasTfStr := "-tf"
	if v.HasTerraform() {
		hasTfStr = "+tf"
	}
	return "[" + hasTfStr + "] " + v.awsVpc.String()
}

func (v *vpc) Aws() aws.Resource {
	return v
}

func (v *vpc) IsAws() bool {
	return true
}

func (v *vpc) HasTerraform() bool {
	return v.tfVpc != nil
}

func (v *vpc) Terraform() terraform.Resource {
	return v.tfVpc
}

func newVpc(awsVpc aws.Vpc, tfVpc terraform.Resource) MappedResource {
	return &vpc{awsVpc: awsVpc, tfVpc: tfVpc}
}
