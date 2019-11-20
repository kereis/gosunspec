// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model64111

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block64111 - Basic Charge Controller -

const (
	ModelID = 64111
)

const (
	AH_SF           = "AH_SF"
	A_SF            = "A_SF"
	ArrayV          = "ArrayV"
	BattV           = "BattV"
	ChargerSt       = "ChargerSt"
	InputA          = "InputA"
	KWH_SF          = "KWH_SF"
	LifeTimeAHOut   = "LifeTimeAHOut"
	LifeTimeKWHOut  = "LifeTimeKWHOut"
	LifeTimeMaxBatt = "LifeTimeMaxBatt"
	LifeTimeMaxOut  = "LifeTimeMaxOut"
	LifeTimeMaxVOC  = "LifeTimeMaxVOC"
	OutputA         = "OutputA"
	OutputW         = "OutputW"
	P_SF            = "P_SF"
	Port            = "Port"
	TodayAHOutput   = "TodayAHOutput"
	TodayMaxBatV    = "TodayMaxBatV"
	TodayMaxVOC     = "TodayMaxVOC"
	TodayMinBatV    = "TodayMinBatV"
	TodaykWhOutput  = "TodaykWhOutput"
	VOCV            = "VOCV"
	V_SF            = "V_SF"
)

type Block64111 struct {
	Port            uint16              `sunspec:"offset=0"`
	V_SF            sunspec.ScaleFactor `sunspec:"offset=1"`
	A_SF            sunspec.ScaleFactor `sunspec:"offset=2"`
	P_SF            sunspec.ScaleFactor `sunspec:"offset=3"`
	AH_SF           sunspec.ScaleFactor `sunspec:"offset=4"`
	KWH_SF          sunspec.ScaleFactor `sunspec:"offset=5"`
	BattV           uint16              `sunspec:"offset=6,sf=V_SF"`
	ArrayV          uint16              `sunspec:"offset=7,sf=V_SF"`
	OutputA         uint16              `sunspec:"offset=8,sf=A_SF"`
	InputA          uint16              `sunspec:"offset=9,sf=P_SF"`
	ChargerSt       sunspec.Enum16      `sunspec:"offset=10"`
	OutputW         uint16              `sunspec:"offset=11,sf=P_SF"`
	TodayMinBatV    uint16              `sunspec:"offset=12,sf=V_SF"`
	TodayMaxBatV    uint16              `sunspec:"offset=13,sf=V_SF"`
	VOCV            uint16              `sunspec:"offset=14,sf=V_SF"`
	TodayMaxVOC     uint16              `sunspec:"offset=15,sf=V_SF"`
	TodaykWhOutput  uint16              `sunspec:"offset=16,sf=KWH_SF"`
	TodayAHOutput   uint16              `sunspec:"offset=17,sf=AH_SF"`
	LifeTimeKWHOut  uint16              `sunspec:"offset=18,sf=P_SF"`
	LifeTimeAHOut   uint16              `sunspec:"offset=19,sf=KWH_SF"`
	LifeTimeMaxOut  uint16              `sunspec:"offset=20,sf=P_SF"`
	LifeTimeMaxBatt uint16              `sunspec:"offset=21,sf=V_SF"`
	LifeTimeMaxVOC  uint16              `sunspec:"offset=22,sf=V_SF"`
}

func (self *Block64111) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 23,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 23,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Port, Offset: 0, Type: typelabel.Uint16, Mandatory: true},
					smdx.PointElement{Id: V_SF, Offset: 1, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: A_SF, Offset: 2, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: P_SF, Offset: 3, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: AH_SF, Offset: 4, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: KWH_SF, Offset: 5, Type: typelabel.ScaleFactor, Mandatory: true},
					smdx.PointElement{Id: BattV, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: ArrayV, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: OutputA, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: InputA, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "P_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: ChargerSt, Offset: 10, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: OutputW, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "P_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: TodayMinBatV, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: TodayMaxBatV, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: VOCV, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: TodayMaxVOC, Offset: 15, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: TodaykWhOutput, Offset: 16, Type: typelabel.Uint16, ScaleFactor: "KWH_SF", Units: "kWh", Mandatory: true},
					smdx.PointElement{Id: TodayAHOutput, Offset: 17, Type: typelabel.Uint16, ScaleFactor: "AH_SF", Units: "AH", Mandatory: true},
					smdx.PointElement{Id: LifeTimeKWHOut, Offset: 18, Type: typelabel.Uint16, ScaleFactor: "P_SF", Units: "kWh", Mandatory: true},
					smdx.PointElement{Id: LifeTimeAHOut, Offset: 19, Type: typelabel.Uint16, ScaleFactor: "KWH_SF", Units: "kAH", Mandatory: true},
					smdx.PointElement{Id: LifeTimeMaxOut, Offset: 20, Type: typelabel.Uint16, ScaleFactor: "P_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: LifeTimeMaxBatt, Offset: 21, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: LifeTimeMaxVOC, Offset: 22, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
				},
			},
		}})
}
