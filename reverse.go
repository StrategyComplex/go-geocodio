package geocodio

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	See: http://geocod.io/docs/#toc_16
*/
// Reverse does a reverse geocode look up for a single coordinate
func (g *Geocodio) Reverse(latitude, longitude float64) (GeocodeResult, error) {
	// if there is an address here, they should probably think about moving
	// regardless, we'll consider it an error
	if latitude == 0.0 && longitude == 0.0 {
		return GeocodeResult{}, ErrReverseGecodeMissingLatLng
	}

	latStr := strconv.FormatFloat(latitude, 'f', 9, 64)
	lngStr := strconv.FormatFloat(longitude, 'f', 9, 64)

	resp := GeocodeResult{}
	err := g.get("/reverse", map[string]string{"q": latStr + "," + lngStr}, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}

// GeocodeAndReturnTimezone will geocode and include Timezone in the fields response
func (g *Geocodio) ReverseAndReturnTimezone(latitude, longitude float64) (GeocodeResult, error) {
	return g.ReverseReturnFields(latitude, longitude, "timezone")
}

// GeocodeAndReturnZip4 will geocode and include zip4 in the fields response
func (g *Geocodio) ReverseAndReturnZip4(latitude, longitude float64) (GeocodeResult, error) {
	return g.ReverseReturnFields(latitude, longitude, "zip4")
}

// GeocodeAndReturnCongressionalDistrict will geocode and include Congressional District in the fields response
func (g *Geocodio) ReverseAndReturnCongressionalDistrict(latitude, longitude float64) (GeocodeResult, error) {
	return g.ReverseReturnFields(latitude, longitude, "cd")
}

// GeocodeAndReturnStateLegislativeDistricts will geocode and include State Legislative Districts in the fields response
func (g *Geocodio) ReverseAndReturnStateLegislativeDistricts(latitude, longitude float64) (GeocodeResult, error) {
	return g.ReverseReturnFields(latitude, longitude, "stateleg")
}

// GeocodeAndReturnCongressAndStateDistricts will geocode and include Congressional District and State Legislative Districts in the fields response
func (g *Geocodio) ReverseAndReturnCongressAndStateDistricts(latitude, longitude float64) (GeocodeResult, error) {
	return g.ReverseReturnFields(latitude, longitude, "cd,stateleg")
}

// GeocodeReturnFields will geocode and includes additional fields in response
/*
 	See: http://geocod.io/docs/#toc_22
	Note:
		Each field counts as an additional lookup each
*/
func (g *Geocodio) ReverseReturnFields(latitude, longitude float64, fields ...string) (GeocodeResult, error) {
	// if there is an address here, they should probably think about moving
	// regardless, we'll consider it an error
	if latitude == 0.0 && longitude == 0.0 {
		return GeocodeResult{}, ErrReverseGecodeMissingLatLng
	}

	latStr := strconv.FormatFloat(latitude, 'f', 9, 64)
	lngStr := strconv.FormatFloat(longitude, 'f', 9, 64)

	fieldsCommaSeparated := strings.Join(fields, ",")
	resp := GeocodeResult{}

	err := g.get("/reverse",
		map[string]string{
			"q":      latStr + "," + lngStr,
			"fields": fieldsCommaSeparated,
		}, &resp)

	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil
}

// ReverseBatch supports a batch lookup by lat/lng coordinate pairs
func (g *Geocodio) ReverseBatch(latlngs ...float64) (BatchResponse, error) {
	resp := BatchResponse{}
	if len(latlngs) == 0 {
		return resp, ErrReverseBatchMissingCoords
	}

	if len(latlngs)%2 == 1 {
		return resp, ErrReverseBatchInvalidCoordsPairs
	}

	var (
		payload = []string{}
		pair    string
	)

	for i := range latlngs {
		coord := strconv.FormatFloat(latlngs[i], 'f', 9, 64)
		if i == 0 {
			pair = coord
			continue
		}
		if i%2 == 0 {
			pair = fmt.Sprintf("%s,%s", pair, coord)
			payload = append(payload, pair)
			continue
		}
		pair = coord
	}

	err := g.post("/reverse", payload, nil, &resp)
	if err != nil {
		return resp, err
	}

	if len(resp.Results) == 0 {
		return resp, ErrNoResultsFound
	}

	return resp, nil

}
