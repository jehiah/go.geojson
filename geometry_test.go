package geojson

import (
	"bytes"
	"testing"
)

func TestGeometryMarshalJSONPoint(t *testing.T) {
	g := NewPointGeometry([]float64{1, 2})
	blob, err := g.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"type":"Point"`)) {
		t.Errorf("json should have type Point")
	}

	if !bytes.Contains(blob, []byte(`[1,2]`)) {
		t.Errorf("json should marshal coordinates correctly")
	}
}

func TestGeometryMarshalJSONMultiPoint(t *testing.T) {
	g := NewMultiPointGeometry([]float64{1, 2}, []float64{3, 4})
	blob, err := g.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"type":"MultiPoint"`)) {
		t.Errorf("json should have type MultiPoint")
	}

	if !bytes.Contains(blob, []byte(`[[1,2],[3,4]]`)) {
		t.Errorf("json should marshal coordinates correctly")
	}
}

func TestGeometryMarshalJSONLineString(t *testing.T) {
	g := NewLineStringGeometry([][]float64{[]float64{1, 2}, []float64{3, 4}})
	blob, err := g.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"type":"LineString"`)) {
		t.Errorf("json should have type LineString")
	}

	if !bytes.Contains(blob, []byte(`[[1,2],[3,4]]`)) {
		t.Errorf("json should marshal coordinates correctly")
	}
}

func TestGeometryMarshalJSONMultiLineString(t *testing.T) {
	g := NewMultiLineStringGeometry(
		[][]float64{[]float64{1, 2}, []float64{3, 4}},
		[][]float64{[]float64{5, 6}, []float64{7, 8}},
	)
	blob, err := g.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"type":"MultiLineString"`)) {
		t.Errorf("json should have type MultiLineString")
	}

	if !bytes.Contains(blob, []byte(`[[[1,2],[3,4]],[[5,6],[7,8]]]`)) {
		t.Errorf("json should marshal coordinates correctly")
	}
}

func TestGeometryMarshalJSONPolygon(t *testing.T) {
	g := NewPolygonGeometry([][][]float64{
		[][]float64{[]float64{1, 2}, []float64{3, 4}},
		[][]float64{[]float64{5, 6}, []float64{7, 8}},
	})
	blob, err := g.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"type":"Polygon"`)) {
		t.Errorf("json should have type Polygon")
	}

	if !bytes.Contains(blob, []byte(`[[[1,2],[3,4]],[[5,6],[7,8]]]`)) {
		t.Errorf("json should marshal coordinates correctly")
	}
}

func TestGeometryMarshalJSONMultiPolygon(t *testing.T) {
	g := NewMultiPolygonGeometry(
		[][][]float64{
			[][]float64{[]float64{1, 2}, []float64{3, 4}},
			[][]float64{[]float64{5, 6}, []float64{7, 8}},
		},
		[][][]float64{
			[][]float64{[]float64{8, 7}, []float64{6, 5}},
			[][]float64{[]float64{4, 3}, []float64{2, 1}},
		})
	blob, err := g.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"type":"MultiPolygon"`)) {
		t.Errorf("json should have type MultiPolygon")
	}

	if !bytes.Contains(blob, []byte(`[[[[1,2],[3,4]],[[5,6],[7,8]]],[[[8,7],[6,5]],[[4,3],[2,1]]]]`)) {
		t.Errorf("json should marshal coordinates correctly")
	}
}

func TestGeometryMarshalJSONCollection(t *testing.T) {
	g := NewCollectionGeometry(
		NewPointGeometry([]float64{1, 2}),
		NewMultiPointGeometry([]float64{1, 2}, []float64{3, 4}),
	)
	blob, err := g.MarshalJSON()

	if err != nil {
		t.Fatalf("should marshal to json just fine but got %v", err)
	}

	if !bytes.Contains(blob, []byte(`"type":"GeometryCollection"`)) {
		t.Errorf("json should have type GeometryCollection")
	}

	if !bytes.Contains(blob, []byte(`"geometries":`)) {
		t.Errorf("json should have geometries attribute")
	}
}

func TestUnmarshalGeometryPoint(t *testing.T) {
	rawJSON := `{"type": "Point", "coordinates": [102.0, 0.5]}`

	g, err := UnmarshalGeometry([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal geometry without issue, err %v", err)
	}

	if g.Type != "Point" {
		t.Errorf("incorrect type, got %v", g.Type)
	}

	if len(g.Point) != 2 {
		t.Errorf("should have 2 coordinate elements but got %d", len(g.Point))
	}
}

func TestUnmarshalGeometryMultiPoint(t *testing.T) {
	rawJSON := `{"type": "MultiPoint", "coordinates": [[1,2],[3,4]]}`

	g, err := UnmarshalGeometry([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal geometry without issue, err %v", err)
	}

	if g.Type != "MultiPoint" {
		t.Errorf("incorrect type, got %v", g.Type)
	}

	if len(g.MultiPoint) != 2 {
		t.Errorf("should have 2 coordinate elements but got %d", len(g.MultiPoint))
	}
}

func TestUnmarshalGeometryLineString(t *testing.T) {
	rawJSON := `{"type": "LineString", "coordinates": [[1,2],[3,4]]}`

	g, err := UnmarshalGeometry([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal geometry without issue, err %v", err)
	}

	if g.Type != "LineString" {
		t.Errorf("incorrect type, got %v", g.Type)
	}

	if len(g.LineString) != 2 {
		t.Errorf("should have 2 line string coordinates but got %d", len(g.LineString))
	}
}

func TestUnmarshalGeometryMultiLineString(t *testing.T) {
	rawJSON := `{"type": "MultiLineString", "coordinates": [[[1,2],[3,4]],[[5,6],[7,8]]]}`

	g, err := UnmarshalGeometry([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal geometry without issue, err %v", err)
	}

	if g.Type != "MultiLineString" {
		t.Errorf("incorrect type, got %v", g.Type)
	}

	if len(g.MultiLineString) != 2 {
		t.Errorf("should have 2 line strings but got %d", len(g.MultiLineString))
	}
}

func TestUnmarshalGeometryPolygon(t *testing.T) {
	rawJSON := `{"type": "Polygon", "coordinates": [[[1,2],[3,4]],[[5,6],[7,8]]]}`

	g, err := UnmarshalGeometry([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal geometry without issue, err %v", err)
	}

	if g.Type != "Polygon" {
		t.Errorf("incorrect type, got %v", g.Type)
	}

	if len(g.Polygon) != 2 {
		t.Errorf("should have 2 polygon paths but got %d", len(g.Polygon))
	}
}

func TestUnmarshalGeometryMultiPolygon(t *testing.T) {
	rawJSON := `{"type": "MultiPolygon", "coordinates": [[[[1,2],[3,4]],[[5,6],[7,8]]],[[[8,7],[6,5]],[[4,3],[2,1]]]]}`

	g, err := UnmarshalGeometry([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal geometry without issue, err %v", err)
	}

	if g.Type != "MultiPolygon" {
		t.Errorf("incorrect type, got %v", g.Type)
	}

	if len(g.MultiPolygon) != 2 {
		t.Errorf("should have 2 polygons but got %d", len(g.MultiPolygon))
	}
}

func TestUnmarshalGeometryCollection(t *testing.T) {
	rawJSON := `{"type": "GeometryCollection", "geometries": [
		{"type": "Point", "coordinates": [102.0, 0.5]},
		{"type": "MultiLineString", "coordinates": [[[1,2],[3,4]],[[5,6],[7,8]]]}
	]}`

	g, err := UnmarshalGeometry([]byte(rawJSON))
	if err != nil {
		t.Fatalf("should unmarshal geometry without issue, err %v", err)
	}

	if g.Type != "GeometryCollection" {
		t.Errorf("incorrect type, got %v", g.Type)
	}

	if len(g.Geometries) != 2 {
		t.Errorf("should have 2 geometries but got %d", len(g.Geometries))
	}
}