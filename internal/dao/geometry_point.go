package dao

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strings"
)

type GeometryPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (g *GeometryPoint) Scan(val any) error {
	var s string
	switch v := val.(type) {
	case []byte:
		s = string(v)
	case string:
		s = v
	default:
		return fmt.Errorf("cannot convert %T to GeometryPoint", val)
	}

	if strings.HasPrefix(s, "SRID=") {
		var srid int
		_, err := fmt.Sscanf(s, "SRID=%d;POINT(%f %f)", &srid, &g.X, &g.Y)
		return err
	}

	if strings.HasPrefix(s, "POINT(") {
		_, err := fmt.Sscanf(s, "POINT(%f %f)", &g.X, &g.Y)
		return err
	}

	// raw column reads return hex-encoded (E)WKB
	return g.scanEWKB(s)
}

func (g *GeometryPoint) scanEWKB(s string) error {
	raw, err := hex.DecodeString(s)
	if err != nil {
		return fmt.Errorf("cannot parse %q as a geometry point", s)
	}
	if len(raw) < 5 {
		return fmt.Errorf("geometry too short: %q", s)
	}

	var order binary.ByteOrder = binary.LittleEndian
	if raw[0] == 0 {
		order = binary.BigEndian
	}

	const (
		ewkbSRIDFlag = 0x20000000
		wkbPointType = 1
	)

	geomType := order.Uint32(raw[1:5])
	body := raw[5:]
	if geomType&ewkbSRIDFlag != 0 {
		if len(body) < 4 {
			return fmt.Errorf("geometry too short: %q", s)
		}
		body = body[4:] // skip SRID
	}

	if geomType&^uint32(ewkbSRIDFlag) != wkbPointType {
		return fmt.Errorf("geometry is not a point: %q", s)
	}
	if len(body) < 16 {
		return fmt.Errorf("geometry too short: %q", s)
	}

	g.X = math.Float64frombits(order.Uint64(body[0:8]))
	g.Y = math.Float64frombits(order.Uint64(body[8:16]))
	return nil
}

// Value implements driver.Valuer
func (g GeometryPoint) Value() (driver.Value, error) {
	return fmt.Sprintf(
		"SRID=4326;POINT(%f %f)",
		g.X,
		g.Y,
	), nil
}
