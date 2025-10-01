
package in

import (
    "time")

type PersonProfileDTOIn struct {
   UpdatedAtStr    string `json:"updated_at"`
   UpdatedAt    time.Time `required:"insert"`
   OtherField1    string `json:"other_field_1" required:"insert" min:"1" max:"100"`
   OtherField2    int64 `json:"other_field_2" required:"insert" max:"1000"`
}