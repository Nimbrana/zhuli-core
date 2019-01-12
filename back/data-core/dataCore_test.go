package datacore

import (
	"testing"
)

var dc = DataController{}

func TestInitDataControllerShouldFail(t *testing.T) {
	dc.UseMongoDriver      	= true
	dc.UseFileSystemDriver 	= false
	dc.Server				= ""
	dc.Database				= ""
	dc.Collection			= ""
	dc.User					= ""
	dc.Password				= ""
	
	if err := dc.Init(); err == nil {
		t.Errorf("%s",err)
	}

}