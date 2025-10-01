
package in

import (
    "time")

type WardesProfileDTOIn struct {
   Updatedat    time.Time `json:"updatedAt" required:"update,delete"`
}