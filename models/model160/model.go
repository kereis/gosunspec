// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model160

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block160 - Multiple MPPT Inverter Extension Model -

const (
	ModelID = 160
)

const (
	DCA     = "DCA"
	DCA_SF  = "DCA_SF"
	DCEvt   = "DCEvt"
	DCSt    = "DCSt"
	DCV     = "DCV"
	DCV_SF  = "DCV_SF"
	DCW     = "DCW"
	DCWH    = "DCWH"
	DCWH_SF = "DCWH_SF"
	DCW_SF  = "DCW_SF"
	Evt     = "Evt"
	ID      = "ID"
	IDStr   = "IDStr"
	N       = "N"
	Tmp     = "Tmp"
	Tms     = "Tms"
	TmsPer  = "TmsPer"
)

type Block160Repeat struct {
	ID    uint16             `sunspec:"offset=0"`
	IDStr string             `sunspec:"offset=1,len=8"`
	DCA   uint16             `sunspec:"offset=9,sf=DCA_SF"`
	DCV   uint16             `sunspec:"offset=10,sf=DCV_SF"`
	DCW   uint16             `sunspec:"offset=11,sf=DCW_SF"`
	DCWH  sunspec.Acc32      `sunspec:"offset=12,sf=DCWH_SF"`
	Tms   uint32             `sunspec:"offset=14"`
	Tmp   int16              `sunspec:"offset=16"`
	DCSt  sunspec.Enum16     `sunspec:"offset=17"`
	DCEvt sunspec.Bitfield32 `sunspec:"offset=18"`
}

type Block160 struct {
	DCA_SF  sunspec.ScaleFactor `sunspec:"offset=0"`
	DCV_SF  sunspec.ScaleFactor `sunspec:"offset=1"`
	DCW_SF  sunspec.ScaleFactor `sunspec:"offset=2"`
	DCWH_SF sunspec.ScaleFactor `sunspec:"offset=3"`
	Evt     sunspec.Bitfield32  `sunspec:"offset=4"`
	N       sunspec.Count       `sunspec:"offset=6"`
	TmsPer  uint16              `sunspec:"offset=7"`

	Repeats []Block160Repeat
}

func (self *Block160) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "mppt",
		Length: 28,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 8,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: DCA_SF, Offset: 0, Type: typelabel.Sunssf},
					smdx.PointElement{Id: DCV_SF, Offset: 1, Type: typelabel.Sunssf},
					smdx.PointElement{Id: DCW_SF, Offset: 2, Type: typelabel.Sunssf},
					smdx.PointElement{Id: DCWH_SF, Offset: 3, Type: typelabel.Sunssf},
					smdx.PointElement{Id: Evt, Offset: 4, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: N, Offset: 6, Type: typelabel.Count},
					smdx.PointElement{Id: TmsPer, Offset: 7, Type: typelabel.Uint16},
				},
			},
			smdx.BlockElement{Name: "module",
				Length: 20,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ID, Offset: 0, Type: typelabel.Uint16},
					smdx.PointElement{Id: IDStr, Offset: 1, Type: typelabel.String, Length: 8},
					smdx.PointElement{Id: DCA, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "DCA_SF", Units: "A"},
					smdx.PointElement{Id: DCV, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "DCV_SF", Units: "V"},
					smdx.PointElement{Id: DCW, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "DCW_SF", Units: "W"},
					smdx.PointElement{Id: DCWH, Offset: 12, Type: typelabel.Acc32, ScaleFactor: "DCWH_SF", Units: "Wh"},
					smdx.PointElement{Id: Tms, Offset: 14, Type: typelabel.Uint32, Units: "Secs"},
					smdx.PointElement{Id: Tmp, Offset: 16, Type: typelabel.Int16, Units: "C"},
					smdx.PointElement{Id: DCSt, Offset: 17, Type: typelabel.Enum16},
					smdx.PointElement{Id: DCEvt, Offset: 18, Type: typelabel.Bitfield32},
				},
			},
		}})
}
