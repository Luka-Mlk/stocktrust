package hrecordlist

import (
	"log"
	"stocktrust/pkg/hrecord"
)

type HRecordList struct {
	Length  int
	Records []hrecord.HRecord

	persistences []Persistence
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

func WithPersistence(ps Persistence) HRLParam {
	return func(h *HRecordList) error {
		h.persistences = append(h.persistences, ps)
		return nil
	}
}

func (hrl *HRecordList) Append(h hrecord.HRecord) {
	hrl.Records = append(hrl.Records, h)
}

func (hrl *HRecordList) AppendHRL(h HRecordList) {
	hrl.Records = append(hrl.Records, h.Records...)
}

func (hl *HRecordList) Save() error {
	for _, persistence := range hl.persistences {
		err := persistence.Save(*hl)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
