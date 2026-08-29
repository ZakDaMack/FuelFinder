package dao

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ewkbPointHex(littleEndian bool, srid uint32, x, y float64) string {
	var order binary.AppendByteOrder = binary.BigEndian
	buf := make([]byte, 0, 25)
	if littleEndian {
		order = binary.LittleEndian
		buf = append(buf, 1)
	} else {
		buf = append(buf, 0)
	}

	geomType := uint32(1) // point
	if srid != 0 {
		geomType |= 0x20000000
	}
	buf = order.AppendUint32(buf, geomType)
	if srid != 0 {
		buf = order.AppendUint32(buf, srid)
	}
	buf = order.AppendUint64(buf, math.Float64bits(x))
	buf = order.AppendUint64(buf, math.Float64bits(y))

	return hex.EncodeToString(buf)
}

func TestGeometryPointScan(t *testing.T) {
	cases := map[string]string{
		"ewkt":             "SRID=4326;POINT(-0.070600 51.801800)",
		"wkt":              "POINT(-0.070600 51.801800)",
		"ewkb little":      ewkbPointHex(true, 4326, -0.0706, 51.8018),
		"ewkb big":         ewkbPointHex(false, 4326, -0.0706, 51.8018),
		"wkb without srid": ewkbPointHex(true, 0, -0.0706, 51.8018),
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			var point GeometryPoint
			assert.NoError(t, point.Scan(input))
			assert.InDelta(t, -0.0706, point.X, 0.0001)
			assert.InDelta(t, 51.8018, point.Y, 0.0001)
		})
	}
}

func TestGeometryPointScanInvalid(t *testing.T) {
	var point GeometryPoint
	assert.Error(t, point.Scan("not a point"))
	assert.Error(t, point.Scan(12345))
	// a linestring, not a point
	assert.Error(t, point.Scan("010200000000000000"))
}
