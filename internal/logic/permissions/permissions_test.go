package permissions

import (
	"context"
	"testing"
)

func Test_get_permissions(Test *testing.T) {
	sPermissions := New()
	name,permissions,err := sPermissions.GetRoleAndPermissions(context.Background(), 1)
	if err != nil {
		Test.Error(err)
		return
	}
	Test.Log(name,permissions)
}
