package geocodio_test

import (
	"fmt"
	"testing"

	"github.com/strategycomplex/go-geocodio"
)

func TestReverseGeocodeLookup(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	result, err := gc.Reverse(AddressTestTwoLatitude, AddressTestTwoLongitude)
	if err != nil {
		t.Error(err)
	}

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

	if len(result.Results) < 3 {
		t.Error("Results found length is less than 3", len(result.Results))
	}

	if len(result.Results) == 0 {
		t.Error("No results were found.")
		return
	}

	if result.Results[0].Formatted != AddressTestTwoFull {
		t.Error("Location latitude does not match", result.Results[0].Formatted, AddressTestTwoFull)
	}
}

func TestReverseWithZeroLatLng(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	_, err = gc.Reverse(0.0, 0.0)
	if err != geocodio.ErrReverseGecodeMissingLatLng {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}
}

func TestReverseLookupReturningTimezone(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := gc.ReverseAndReturnTimezone(AddressTestOneLatitude, AddressTestOneLongitude)
	if err != nil {
		t.Error(err)
	}

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

func TestReverseLookupReturningZip4(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	result, err := gc.ReverseAndReturnZip4(AddressTestOneLatitude, AddressTestOneLongitude)
	if err != nil {
		t.Error(err)
	}

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

func TestReverseLookupReturningCongressionalDistrict(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}
	result, err := gc.ReverseAndReturnCongressionalDistrict(AddressTestOneLatitude, AddressTestOneLongitude)
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

func TestReverseLookupReturningStateLegislativeDistricts(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
		t.Fail()
	}

	result, err := gc.ReverseAndReturnStateLegislativeDistricts(AddressTestOneLatitude, AddressTestOneLongitude)
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

	if result.Results[0].Fields.StateLegislativeDistricts.House.DistrictNumber != "2" {
		t.Error("State Legislative Districts house does not match", result.Results[0].Fields.StateLegislativeDistricts.House)
		t.Fail()
	}

	if result.Results[0].Fields.StateLegislativeDistricts.Senate.DistrictNumber != "40" {
		t.Error("State Legislative Districts senate does not match", result.Results[0].Fields.StateLegislativeDistricts.Senate)
		t.Fail()
	}
}

func TestReverseLookupReturningMultipleFields(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	result, err := gc.GeocodeReturnFields(AddressTestOneFull, "timezone", "cd")
	if err != nil {
		t.Error(err)
	}

	// fmt.Println(result.Debugc.RequestedURL)

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

func TestReverseBatchLookup(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	result, err := gc.ReverseBatch(
		AddressTestOneLatitude, AddressTestOneLongitude,
		AddressTestTwoLatitude, AddressTestTwoLongitude,
		AddressTestThreeLatitude, AddressTestThreeLongitude,
	)
	fmt.Println(result.ResponseAsString())
	if err != nil {
		t.Error(err)
	}

	if len(result.Results) == 0 {
		t.Error("Results length is 0")
	}

}

func TestReverseBatchWithoutLatLng(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
	_, err = gc.ReverseBatch()
	if err != geocodio.ErrReverseBatchMissingCoords {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}
}

func TestReverseBatchWithInvalidLatLngPairs(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = gc.ReverseBatch(AddressTestOneLatitude)
	if err != geocodio.ErrReverseBatchInvalidCoordsPairs {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}

	_, err = gc.ReverseBatch(AddressTestOneLatitude, AddressTestOneLongitude, AddressTestTwoLatitude)
	if err != geocodio.ErrReverseBatchInvalidCoordsPairs {
		t.Errorf("Error should be '%s' not '%s'", geocodio.ErrReverseGecodeMissingLatLng.Error(), err.Error())
	}
}

func TestReverseWithInvalidLatLng(t *testing.T) {
	gc, err := geocodio.New()
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = gc.Reverse(1.234, 5.678)
	if err == nil {
		t.Error("Expected to see an error")
		return
	}
	if err != geocodio.ErrNoResultsFound {
		t.Error("Expected error", geocodio.ErrNoResultsFound.Error(), "but saw", err.Error())
	}

	_, err = gc.ReverseBatch(1.234, 5.678)
	if err == nil {
		t.Error("Expected to see an error")
		return
	}
	if err != geocodio.ErrNoResultsFound {
		t.Error("Expected error", geocodio.ErrNoResultsFound.Error(), "but saw", err.Error())
	}

}
