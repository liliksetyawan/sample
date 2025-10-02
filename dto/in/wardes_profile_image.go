
package in

import (
    "time")

type WardesProfileImageDTOIn struct {
   ProfileImage    string `json:"profile_image" required:"i,n,s,e,r,t" min:"1" max:"255"`
   Description    string `json:"description" required:"i,n,s,e,r,t" max:"500"`
   RecordStatus    string `json:"record_status" required:"i,n,s,e,r,t" min:"1" max:"1"`
}