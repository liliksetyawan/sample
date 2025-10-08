
package in

import (
    "time")

type WardesProfileDTOIn struct {
   NexchiefAccountId    int64 `json:"nexchief_account_id" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"9223372036854775807"`
   UserId    int64 `json:"user_id" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"9223372036854775807"`
   Username    string `json:"username" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"255" regex:"username"`
   PersonalName    string `json:"personal_name" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" auto_fix:"title" min:"1" max:"255" regex:"name"`
   Gender    string `json:"gender" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"10" regex:"text_only"`
   Phone    string `json:"phone" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"8" max:"20" regex:"phone"`
   Email    string `json:"email" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" auto_fix:"lowercase" min:"5" max:"255" regex:"email"`
   ProvinceId    int64 `json:"province_id" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"9223372036854775807"`
   DistrictId    int64 `json:"district_id" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"9223372036854775807"`
   SubDistrictId    int64 `json:"sub_district_id" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"9223372036854775807"`
   UrbanVillageId    int64 `json:"urban_village_id" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"9223372036854775807"`
   Nik    string `json:"nik" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"16" max:"16" regex:"nik"`
   Age    string `json:"age" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"3" regex:"numeric"`
   Hamlet    string `json:"hamlet" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   Neighbour    string `json:"neighbour" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   VillageFamily    string `json:"village_family" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   VillagePopulation    string `json:"village_population" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   BuildingLocation    string `json:"building_location" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"255" regex:"text_only"`
   BuildingOwnership    string `json:"building_ownership" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   BuildingType    string `json:"building_type" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   BuildingForm    string `json:"building_form" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   BuildingWidthFirstFloor    float64 `json:"building_width_first_floor" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   BuildingLengthFirstFloor    float64 `json:"building_length_first_floor" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   BuildingHeightFirstFloor    float64 `json:"building_height_first_floor" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   BuildingWidthSecondFloor    float64 `json:"building_width_second_floor" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   BuildingLengthSecondFloor    float64 `json:"building_length_second_floor" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   BuildingHeightSecondFloor    float64 `json:"building_height_second_floor" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   ParkingLotWidth    float64 `json:"parking_lot_width" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   ParkingLotLength    float64 `json:"parking_lot_length" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" max:"999999.99"`
   VillageInternetConnection    string `json:"village_internet_connection" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   InvestmentCapital    string `json:"investment_capital" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   InvestmentCapitalSource    string `json:"investment_capital_source" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   OwnershipStatus    string `json:"ownership_status" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   OwnershipFixedAsset    string `json:"ownership_fixed_asset" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   PermissionCover    string `json:"permission_cover" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   ProfileCompletionStatus    string `json:"profile_completion_status" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"text_only"`
   CompanyProfileId    int64 `json:"company_profile_id" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"9223372036854775807"`
   Code    string `json:"code" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"alphanumeric"`
   Schema    string `json:"schema" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"100" regex:"alphanumeric"`
   Status    string `json:"status" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"50" regex:"text_only"`
   ActiveDate    time.Time `json:"active_date" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e"`
   ResignDate    time.Time `json:"resign_date" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e"`
   IsNexwise    bool `json:"is_nexwise" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e"`
   DeadlineDate    time.Time `json:"deadline_date" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e"`
   NewProfileApprovalStatus    string `json:"new_profile_approval_status" required:"i,n,s,e,r,t,,, ,u,p,d,a,t,e" min:"1" max:"50" regex:"text_only"`
}