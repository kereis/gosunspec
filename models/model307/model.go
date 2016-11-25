// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model307

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block307 - Base Met - Base Meteorolgical Model

const (
	ModelID = 307
)

const (
	ElecFld = "ElecFld"
	PPT     = "PPT"
	Pres    = "Pres"
	RH      = "RH"
	Rain    = "Rain"
	Snw     = "Snw"
	SoilWet = "SoilWet"
	SurWet  = "SurWet"
	TmpAmb  = "TmpAmb"
	WndDir  = "WndDir"
	WndSpd  = "WndSpd"
)

type Block307 struct {
	TmpAmb  int16 `sunspec:"offset=0,sf=-1"`
	RH      int16 `sunspec:"offset=1"`
	Pres    int16 `sunspec:"offset=2"`
	WndSpd  int16 `sunspec:"offset=3"`
	WndDir  int16 `sunspec:"offset=4"`
	Rain    int16 `sunspec:"offset=5"`
	Snw     int16 `sunspec:"offset=6"`
	PPT     int16 `sunspec:"offset=7"`
	ElecFld int16 `sunspec:"offset=8"`
	SurWet  int16 `sunspec:"offset=9"`
	SoilWet int16 `sunspec:"offset=10"`
}

func (self *Block307) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "base_met",
		Length: 11,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 11,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: TmpAmb, Offset: 0, Type: typelabel.Int16, ScaleFactor: "-1", Units: "C"},
					smdx.PointElement{Id: RH, Offset: 1, Type: typelabel.Int16, Units: "Pct"},
					smdx.PointElement{Id: Pres, Offset: 2, Type: typelabel.Int16, Units: "HPa"},
					smdx.PointElement{Id: WndSpd, Offset: 3, Type: typelabel.Int16, Units: "mps"},
					smdx.PointElement{Id: WndDir, Offset: 4, Type: typelabel.Int16, Units: "deg"},
					smdx.PointElement{Id: Rain, Offset: 5, Type: typelabel.Int16, Units: "mm"},
					smdx.PointElement{Id: Snw, Offset: 6, Type: typelabel.Int16, Units: "mm"},
					smdx.PointElement{Id: PPT, Offset: 7, Type: typelabel.Int16},
					smdx.PointElement{Id: ElecFld, Offset: 8, Type: typelabel.Int16, Units: "Vm"},
					smdx.PointElement{Id: SurWet, Offset: 9, Type: typelabel.Int16, Units: "kO"},
					smdx.PointElement{Id: SoilWet, Offset: 10, Type: typelabel.Int16, Units: "Pct"},
				},
			},
		}})
}
