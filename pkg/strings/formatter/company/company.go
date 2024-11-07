package compfmt

import (
	"stocktrust/pkg/company"
	"strings"
)

func Company(cmp *company.Company) {
	scn := strings.Split(cmp.ContactName, "\n")
	for i, v := range scn {
		scn[i] = strings.TrimSpace(v)
	}
	cmp.ContactName = strings.Join(scn, ",")
	scp := strings.Split(cmp.ContactPhone, "\n")
	for i, v := range scp {
		scp[i] = strings.TrimSpace(v)
	}
	cmp.ContactPhone = strings.Join(scp, ",")
	sce := strings.Split(cmp.ContactEmail, "\n")
	for i, v := range sce {
		sce[i] = strings.TrimSpace(v)
	}
	cmp.ContactEmail = strings.Join(sce, ",")
}
