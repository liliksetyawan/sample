
package in

import (
    "time")

type WardesProfileDTOIn struct {
   Name    string `json:"name" required:"insert,update" min:"1" max:"255"`
   Address    string `json:"address" required:"insert,update" min:"1" max:"500"`
   Phone    string `json:"phone" required:"insert,update" min:"1" max:"20"`
   Email    string `json:"email" required:"insert,update" min:"1" max:"255"`
   RecordStatus    string `json:"record_status" required:"insert,update" min:"1" max:"1"`
   UpdatedAt    time.Time `json:"updated_at" required:"update"`
}