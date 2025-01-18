package compfmt

import (
	"pages/pkg/company"
	"strings"
)

func Company(cmp *company.Company) {
	if cmp.Address == "" {
		cmp.Address = "/"
		cmp.City = "/"
		cmp.Country = "/"
		cmp.Email = "/"
		cmp.Website = "/"
		cmp.ContactName = "/"
		cmp.ContactPhone = "/"
		cmp.ContactEmail = "/"
		cmp.Phone = "/"
		cmp.Fax = "/"
		cmp.Prospect = "/"
		return
	}
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
