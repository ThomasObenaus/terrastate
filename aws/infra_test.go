package aws

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomas.obenaus/inframapper/trace"
)

func TestFindById(t *testing.T) {

	var vpcs []*Vpc
	vpc := &Vpc{
		CIDR:  "BLA",
		VpcId: "blubb",
	}
	vpcs = append(vpcs, vpc)

	data := &infraData{vpcs: vpcs}
	infra, err := newInfra(data)
	require.NotNil(t, infra)
	require.NoError(t, err)

	resource := infra.FindById("ABCD")
	assert.Nil(t, resource)

	resource = infra.FindById("blubb")
	assert.NotNil(t, resource)
}

func TestFindVpc(t *testing.T) {

	var vpcs []*Vpc
	vpc := &Vpc{
		CIDR:  "BLA",
		VpcId: "blubb",
	}
	vpcs = append(vpcs, vpc)

	data := &infraData{vpcs: vpcs}
	infra, err := newInfra(data)
	require.NotNil(t, infra)
	require.NoError(t, err)

	vpc = infra.FindVpc("ABCD")
	assert.Nil(t, vpc)

	vpc = infra.FindVpc("blubb")
	assert.NotNil(t, vpc)
}

func TestNew(t *testing.T) {
	infra, err := newInfra(nil)
	assert.Error(t, err)
	assert.Nil(t, infra)

	data := &infraData{}
	infra, err = newInfra(data)
	assert.NoError(t, err)
	assert.NotNil(t, infra)
}

func TestVpcs(t *testing.T) {

	data := &infraData{}
	infra := &infraImpl{tracer: trace.Off(), data: data}

	vpcs := infra.Vpcs()
	assert.Nil(t, vpcs)

	infra = &infraImpl{tracer: trace.Off()}
	vpcs = infra.Vpcs()
	assert.Nil(t, vpcs)
}
