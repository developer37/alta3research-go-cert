package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)
// Template of JSON for Current NEOS
type CurrentNeos struct {
	Links struct {
		Next string `json:"next"`
		Self string `json:"self"`
	} `json:"links"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"total_elements"`
		TotalPages    int `json:"total_pages"`
		Number        int `json:"number"`
	} `json:"page"`
	NearEarthObjects []struct {
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		ID                 string  `json:"id"`
		NeoReferenceID     string  `json:"neo_reference_id"`
		Name               string  `json:"name"`
		NameLimited        string  `json:"name_limited"`
		Designation        string  `json:"designation"`
		NasaJplURL         string  `json:"nasa_jpl_url"`
		AbsoluteMagnitudeH float64 `json:"absolute_magnitude_h"`
		EstimatedDiameter  struct {
			Kilometers struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"kilometers"`
			Meters struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"meters"`
			Miles struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"miles"`
			Feet struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"feet"`
		} `json:"estimated_diameter"`
		IsPotentiallyHazardousAsteroid bool `json:"is_potentially_hazardous_asteroid"`
		CloseApproachData              []struct {
			CloseApproachDate      string `json:"close_approach_date"`
			CloseApproachDateFull  string `json:"close_approach_date_full"`
			EpochDateCloseApproach int64  `json:"epoch_date_close_approach"`
			RelativeVelocity       struct {
				KilometersPerSecond string `json:"kilometers_per_second"`
				KilometersPerHour   string `json:"kilometers_per_hour"`
				MilesPerHour        string `json:"miles_per_hour"`
			} `json:"relative_velocity"`
			MissDistance struct {
				Astronomical string `json:"astronomical"`
				Lunar        string `json:"lunar"`
				Kilometers   string `json:"kilometers"`
				Miles        string `json:"miles"`
			} `json:"miss_distance"`
			OrbitingBody string `json:"orbiting_body"`
		} `json:"close_approach_data"`
		OrbitalData struct {
			OrbitID                   string `json:"orbit_id"`
			OrbitDeterminationDate    string `json:"orbit_determination_date"`
			FirstObservationDate      string `json:"first_observation_date"`
			LastObservationDate       string `json:"last_observation_date"`
			DataArcInDays             int    `json:"data_arc_in_days"`
			ObservationsUsed          int    `json:"observations_used"`
			OrbitUncertainty          string `json:"orbit_uncertainty"`
			MinimumOrbitIntersection  string `json:"minimum_orbit_intersection"`
			JupiterTisserandInvariant string `json:"jupiter_tisserand_invariant"`
			EpochOsculation           string `json:"epoch_osculation"`
			Eccentricity              string `json:"eccentricity"`
			SemiMajorAxis             string `json:"semi_major_axis"`
			Inclination               string `json:"inclination"`
			AscendingNodeLongitude    string `json:"ascending_node_longitude"`
			OrbitalPeriod             string `json:"orbital_period"`
			PerihelionDistance        string `json:"perihelion_distance"`
			PerihelionArgument        string `json:"perihelion_argument"`
			AphelionDistance          string `json:"aphelion_distance"`
			PerihelionTime            string `json:"perihelion_time"`
			MeanAnomaly               string `json:"mean_anomaly"`
			MeanMotion                string `json:"mean_motion"`
			Equinox                   string `json:"equinox"`
			OrbitClass                struct {
				OrbitClassType        string `json:"orbit_class_type"`
				OrbitClassDescription string `json:"orbit_class_description"`
				OrbitClassRange       string `json:"orbit_class_range"`
			} `json:"orbit_class"`
		} `json:"orbital_data"`
		IsSentryObject bool `json:"is_sentry_object"`
	} `json:"near_earth_objects"`
}

func GetNeoDetails(neoid string, url string) {
    // Template for a NEO detail JSON
	type NeodataTemplate struct {
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		ID                 string  `json:"id"`
		NeoReferenceID     string  `json:"neo_reference_id"`
		Name               string  `json:"name"`
		Designation        string  `json:"designation"`
		NasaJplURL         string  `json:"nasa_jpl_url"`
		AbsoluteMagnitudeH float64 `json:"absolute_magnitude_h"`
		EstimatedDiameter  struct {
			Kilometers struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"kilometers"`
			Meters struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"meters"`
			Miles struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"miles"`
			Feet struct {
				EstimatedDiameterMin float64 `json:"estimated_diameter_min"`
				EstimatedDiameterMax float64 `json:"estimated_diameter_max"`
			} `json:"feet"`
		} `json:"estimated_diameter"`
		IsPotentiallyHazardousAsteroid bool `json:"is_potentially_hazardous_asteroid"`
		CloseApproachData              []struct {
			CloseApproachDate      string `json:"close_approach_date"`
			CloseApproachDateFull  string `json:"close_approach_date_full"`
			EpochDateCloseApproach int64  `json:"epoch_date_close_approach"`
			RelativeVelocity       struct {
				KilometersPerSecond string `json:"kilometers_per_second"`
				KilometersPerHour   string `json:"kilometers_per_hour"`
				MilesPerHour        string `json:"miles_per_hour"`
			} `json:"relative_velocity"`
			MissDistance struct {
				Astronomical string `json:"astronomical"`
				Lunar        string `json:"lunar"`
				Kilometers   string `json:"kilometers"`
				Miles        string `json:"miles"`
			} `json:"miss_distance"`
			OrbitingBody string `json:"orbiting_body"`
		} `json:"close_approach_data"`
		OrbitalData struct {
			OrbitID                   string `json:"orbit_id"`
			OrbitDeterminationDate    string `json:"orbit_determination_date"`
			FirstObservationDate      string `json:"first_observation_date"`
			LastObservationDate       string `json:"last_observation_date"`
			DataArcInDays             int    `json:"data_arc_in_days"`
			ObservationsUsed          int    `json:"observations_used"`
			OrbitUncertainty          string `json:"orbit_uncertainty"`
			MinimumOrbitIntersection  string `json:"minimum_orbit_intersection"`
			JupiterTisserandInvariant string `json:"jupiter_tisserand_invariant"`
			EpochOsculation           string `json:"epoch_osculation"`
			Eccentricity              string `json:"eccentricity"`
			SemiMajorAxis             string `json:"semi_major_axis"`
			Inclination               string `json:"inclination"`
			AscendingNodeLongitude    string `json:"ascending_node_longitude"`
			OrbitalPeriod             string `json:"orbital_period"`
			PerihelionDistance        string `json:"perihelion_distance"`
			PerihelionArgument        string `json:"perihelion_argument"`
			AphelionDistance          string `json:"aphelion_distance"`
			PerihelionTime            string `json:"perihelion_time"`
			MeanAnomaly               string `json:"mean_anomaly"`
			MeanMotion                string `json:"mean_motion"`
			Equinox                   string `json:"equinox"`
			OrbitClass                struct {
				OrbitClassType        string `json:"orbit_class_type"`
				OrbitClassDescription string `json:"orbit_class_description"`
				OrbitClassRange       string `json:"orbit_class_range"`
			} `json:"orbit_class"`
		} `json:"orbital_data"`
		IsSentryObject bool `json:"is_sentry_object"`
	}
	path := url

	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var neodatatemplate NeodataTemplate
	jsonErr := json.Unmarshal(body, &neodatatemplate)
	if jsonErr != nil {
		log.Fatalf("unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}
	fmt.Println("NeoReferenceID:        " + neodatatemplate.NeoReferenceID)
	//fmt.Println(neodatatemplate.CloseApproachData)
	for _, e := range neodatatemplate.CloseApproachData {
		fmt.Println("CloseApproachDateFull: " + e.CloseApproachDateFull)
		fmt.Println("RelativeVelocity:      " + e.RelativeVelocity.MilesPerHour)
		fmt.Println("MissDistance:          " + e.MissDistance.Miles)
		fmt.Println(" ")
	}
	fmt.Println("GetNeoDetails Completed...")

}

func main() {

	keyvalue := "DEMO_KEY"
	params := "api_key=" + url.QueryEscape(keyvalue)

	path := fmt.Sprintf("https://api.nasa.gov/neo/rest/v1/neo/browse?%s", params)

	resp, err := http.Get(path)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(body))

	var currentneos CurrentNeos
	jsonErr := json.Unmarshal(body, &currentneos)
	if jsonErr != nil {
		log.Fatalf("unable to parse value: %q, error: %s", string(body), jsonErr.Error())
	}
	//fmt.Println(currentneos.NearEarthObjects)
	neos := currentneos.NearEarthObjects
	//Create map
	//key=id, url=self
	var neomap = make(map[string]string)
	var neoId string
	count := true
	for _, e := range neos {
		//fmt.Println(e.NeoReferenceID)
		//fmt.Println(e.Links.Self)
		neomap[e.NeoReferenceID] = e.Links.Self
		//get 1st Neo for subsequent REST GET call for its details
		if count {
			neoId = e.NeoReferenceID
			count = false

		}

	}
	//print list of neos ID and urls' for detail REST GET calls
	for key, value := range neomap {
		fmt.Println(key, value)
	}
	//Get 1st Neo details using its url for REST GET call
	GetNeoDetails(neoId, neomap[neoId])

	fmt.Println("Main Completed...")

}
