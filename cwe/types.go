package cwe

import (
	"encoding/json"
	"fmt"
)

const (
	//URL is the base URL for CWE definitions
	URL = "https://cwe.mitre.org/data/definitions/"
	//Acronym is the acronym of CWE
	Acronym = "CWE"
)

// Weakness defines a CWE weakness based on http://cwe.mitre.org/data/xsd/cwe_schema_v6.4.xsd
type Weakness struct {
	ID          string
	Name        string
	Description string
}

//SprintURL format the CWE URL
func (w *Weakness) SprintURL() string {
	return fmt.Sprintf("%s%s.html", URL, w.ID)
}

//SprintID format the CWE ID
func (w *Weakness) SprintID() string {
	return fmt.Sprintf("%s-%s", Acronym, w.ID)
}

//MarshalJSON print only id and URL
func (w *Weakness) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	}{
		ID:  w.ID,
		URL: w.SprintURL(),
	})
}
