// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model302

import (
	"github.com/andig/gosunspec"
	"github.com/andig/gosunspec/smdx"
	"github.com/andig/gosunspec/typelabel"
)

// Block302 - Irradiance Model - Include to support various irradiance measurements

const (
	ModelID = 302
)

const (
	DFI  = "DFI"
	DNI  = "DNI"
	GHI  = "GHI"
	OTI  = "OTI"
	POAI = "POAI"
)

type Block302Repeat struct {
	GHI  uint16 `sunspec:"offset=0"`
	POAI uint16 `sunspec:"offset=1"`
	DFI  uint16 `sunspec:"offset=2"`
	DNI  uint16 `sunspec:"offset=3"`
	OTI  uint16 `sunspec:"offset=4"`
}

type Block302 struct {
	Repeats []Block302Repeat
}

func (self *Block302) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "irradiance",
		Length: 5,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 5,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: GHI, Offset: 0, Type: typelabel.Uint16, Units: "W/m2"},
					smdx.PointElement{Id: POAI, Offset: 1, Type: typelabel.Uint16, Units: "W/m2"},
					smdx.PointElement{Id: DFI, Offset: 2, Type: typelabel.Uint16, Units: "W/m2"},
					smdx.PointElement{Id: DNI, Offset: 3, Type: typelabel.Uint16, Units: "W/m2"},
					smdx.PointElement{Id: OTI, Offset: 4, Type: typelabel.Uint16, Units: "W/m2"},
				},
			},
		}})
}
