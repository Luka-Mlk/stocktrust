package hrecordlist

import (
	"log"
	"stocktrust/pkg/hrecord"
)

type HRecordList struct {
	Length  int
	Records []hrecord.HRecord

	persistances []Persistence
}

type Persistence interface {
	Save(h HRecordList) error
}

type HRLParam func(*HRecordList) error

func NewHRecordList(o ...HRLParam) (*HRecordList, error) {
	hrl := &HRecordList{}
	for _, option := range o {
		err := option(hrl)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return hrl, nil
}

func WithHRecord(r *hrecord.HRecord) HRLParam {
	return func(hl *HRecordList) error {
		hl.Records = append(hl.Records, *r)
		hl.Length++
		return nil
	}
}

func (hl *HRecordList) Save() error {
	for _, persistence := range hl.persistances {
		err := persistence.Save(*hl)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
