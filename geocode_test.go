package geocodio_test

import (
	"testing"

	"github.com/stevepartridge/geocodio"
)

func TestGeocodeWithEmptyAddress(t *testing.T) {
	gio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	_, err = gio.Geocode("")
	if err == nil {
		t.Error("Error should not be nil.")
	}
}

func TestGeocodeDebugResponseAsString(t *testing.T) {
	gio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := gio.Geocode(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	if result.ResponseAsString() == "" {
		t.Error("Response should be a valid stringio.")
	}

}

func TestGeocodeFullAddress(t *testing.T) {
	gio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := gio.Geocode(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	// t.Log(result.ResponseAsString())

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Errorf("Location latitude %f does not match %f", result.Results[0].Location.Latitude, AddressTestOneLatitude)
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Errorf("Location longitude %f does not match %f", result.Results[0].Location.Longitude, AddressTestOneLongitude)
	}
}

func TestGeocodeFullAddressReturningTimezone(t *testing.T) {
	gio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := gio.GeocodeAndReturnTimezone(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	// t.Log(result.ResponseAsString())

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Errorf("Location latitude %f does not match %f", result.Results[0].Location.Latitude, AddressTestOneLatitude)
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Errorf("Location longitude %f does not match %f", result.Results[0].Location.Longitude, AddressTestOneLongitude)
	}

	if result.Results[0].Fields.Timezone.Name == "" {
		t.Error("Timezone field not found")
	}

	if !result.Results[0].Fields.Timezone.ObservesDST {
		t.Error("Timezone field does not match", result.Results[0].Fields.Timezone)
	}
}

func TestGeocodeFullAddressReturningZip4(t *testing.T) {
	Geocodio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := Geocodio.GeocodeAndReturnZip4(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(result.Debugio.RequestedURL)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Errorf("Location latitude %f does not match %f", result.Results[0].Location.Latitude, AddressTestOneLatitude)
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Errorf("Location longitude %f does not match %f", result.Results[0].Location.Longitude, AddressTestOneLongitude)
	}

	if len(result.Results[0].Fields.Zip4.Plus4) == 0 {
		t.Error("Zip4 field not found")
	}

	// if !result.Results[0].Fields.Timezone.ObservesDST {
	// 	t.Error("Zip4 field does not match", result.Results[0].Fields.Timezone)
	// }
}

func TestGeocodeFullAddressReturningCongressionalDistrict(t *testing.T) {
	Geocodio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}
	result, err := Geocodio.GeocodeAndReturnCongressionalDistrict(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
		t.Fail()
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, AddressTestOneLatitude)
		t.Fail()
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, AddressTestOneLongitude)
		t.Fail()
	}

	if len(result.Results[0].Fields.CongressionalDistricts) == 0 {
		t.Error("Congressional District field not found", result.Results[0].Fields.CongressionalDistrict)
		t.Fail()
	}

	if result.Results[0].Fields.CongressionalDistricts[0].Name == "" {
		t.Error("Congressional District field not found", result.Results[0].Fields.CongressionalDistricts[0])
		t.Fail()
	}

	if result.Results[0].Fields.CongressionalDistricts[0].DistrictNumber != 8 {
		t.Error("Congressional District field does not match", result.Results[0].Fields.CongressionalDistrict)
		t.Fail()
	}
}

func TestGeocodeFullAddressReturningStateLegislativeDistricts(t *testing.T) {
	Geocodio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}

	result, err := Geocodio.GeocodeAndReturnStateLegislativeDistricts(AddressTestOneFull)
	if err != nil {
		t.Error(err)
	}

	// t.Log(result.ResponseAsString())

	if len(result.Results) == 0 {
		t.Error("Results length is 0", result)
		t.Fail()
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Errorf("Location latitude %f does not match %f", result.Results[0].Location.Latitude, AddressTestOneLatitude)
		t.Fail()
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Errorf("Location longitude %f does not match %f", result.Results[0].Location.Longitude, AddressTestOneLongitude)
		t.Fail()
	}

	if result.Results[0].Fields.StateLegislativeDistricts.House.DistrictNumber != "47" {
		t.Error("State Legislative Districts house does not match", result.Results[0].Fields.StateLegislativeDistricts.House)
		t.Fail()
	}

	if result.Results[0].Fields.StateLegislativeDistricts.Senate.DistrictNumber != "31" {
		t.Error("State Legislative Districts senate does not match", result.Results[0].Fields.StateLegislativeDistricts.Senate)
		t.Fail()
	}
}

func TestGeocodeFullAddressReturningMultipleFields(t *testing.T) {
	Geocodio, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := Geocodio.GeocodeReturnFields(AddressTestOneFull, "timezone", "cd")
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(result.Debugio.RequestedURL)

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if result.Results[0].Location.Latitude != AddressTestOneLatitude {
		t.Error("Location latitude does not match", result.Results[0].Location.Latitude, AddressTestOneLatitude)
	}

	if result.Results[0].Location.Longitude != AddressTestOneLongitude {
		t.Error("Location longitude does not match", result.Results[0].Location.Longitude, AddressTestOneLongitude)
	}

	if result.Results[0].Fields.Timezone.Name == "" {
		t.Error("Timezone field not found")
	}

	if !result.Results[0].Fields.Timezone.ObservesDST {
		t.Error("Timezone field does not match", result.Results[0].Fields.Timezone)
	}

	congressionalDistrict := geocodio.CongressionalDistrict{}

	// check congressional district
	if result.Results[0].Fields.CongressionalDistrict.Name != "" {
		congressionalDistrict = result.Results[0].Fields.CongressionalDistrict
	} else if len(result.Results[0].Fields.CongressionalDistricts) > 0 {
		congressionalDistrict = result.Results[0].Fields.CongressionalDistricts[0]
	}

	if congressionalDistrict.Name == "" {
		t.Error("Congressional District field not found", congressionalDistrict)
	}

	if congressionalDistrict.DistrictNumber != 8 {
		t.Error("Congressional District field does not match", result.Results[0].Fields.CongressionalDistrict)
	}

}

// TODO: School District (school)
