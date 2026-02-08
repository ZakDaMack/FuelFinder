package dao

import (
	"database/sql/driver"
	"fmt"
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

	_, err := fmt.Sscanf(s, "POINT(%f %f)", &g.X, &g.Y)
	return err
}

// Value implements driver.Valuer
func (g GeometryPoint) Value() (driver.Value, error) {
	return fmt.Sprintf(
		"SRID=4326;POINT(%f %f)",
		g.X,
		g.Y,
	), nil
}
