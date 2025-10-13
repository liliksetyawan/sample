
package in

import (
    "time"
)

type WardesProfileDTOIn struct {
   NexchiefAccountID    int64 `json:"nexchief_account_id" required:"insert,update" min:"1" max:"9223372036854775807"`
   UserID    int64 `json:"user_id" required:"insert,update" min:"1" max:"9223372036854775807"`
   Username    string `json:"username" required:"insert,update" min:"1" max:"255" regex:"username"`
   PersonalName    string `json:"personal_name" required:"insert,update" auto_fix:"title" min:"1" max:"255" regex:"name"`
   Gender    string `json:"gender" required:"insert,update" min:"1" max:"10" regex:"text_only"`
   Phone    string `json:"phone" required:"insert,update" min:"8" max:"20" regex:"phone"`
   Email    string `json:"email" required:"insert,update" auto_fix:"lowercase" min:"5" max:"255" regex:"email"`
   ProvinceID    int64 `json:"province_id" required:"insert,update" min:"1" max:"9223372036854775807"`
   DistrictID    int64 `json:"district_id" required:"insert,update" min:"1" max:"9223372036854775807"`
   SubDistrictID    int64 `json:"sub_district_id" required:"insert,update" min:"1" max:"9223372036854775807"`
   UrbanVillageID    int64 `json:"urban_village_id" required:"insert,update" min:"1" max:"9223372036854775807"`
   Nik    string `json:"nik" required:"insert,update" min:"16" max:"16" regex:"nik"`
   Age    string `json:"age" required:"insert,update" min:"1" max:"3" regex:"numeric"`
   Hamlet    string `json:"hamlet" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   Neighbour    string `json:"neighbour" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   VillageFamily    string `json:"village_family" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   VillagePopulation    string `json:"village_population" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   BuildingLocation    string `json:"building_location" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   BuildingOwnership    string `json:"building_ownership" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   BuildingType    string `json:"building_type" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   BuildingForm    string `json:"building_form" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   BuildingWidthFirstFloor    float64 `json:"building_width_first_floor" required:"insert,update" max:"999999.99"`
   BuildingLengthFirstFloor    float64 `json:"building_length_first_floor" required:"insert,update" max:"999999.99"`
   BuildingHeightFirstFloor    float64 `json:"building_height_first_floor" required:"insert,update" max:"999999.99"`
   BuildingWidthSecondFloor    float64 `json:"building_width_second_floor" required:"insert,update" max:"999999.99"`
   BuildingLengthSecondFloor    float64 `json:"building_length_second_floor" required:"insert,update" max:"999999.99"`
   BuildingHeightSecondFloor    float64 `json:"building_height_second_floor" required:"insert,update" max:"999999.99"`
   ParkingLotWidth    float64 `json:"parking_lot_width" required:"insert,update" max:"999999.99"`
   ParkingLotLength    float64 `json:"parking_lot_length" required:"insert,update" max:"999999.99"`
   VillageInternetConnection    string `json:"village_internet_connection" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   InvestmentCapital    string `json:"investment_capital" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   InvestmentCapitalSource    string `json:"investment_capital_source" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   OwnershipStatus    string `json:"ownership_status" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   OwnershipFixedAsset    string `json:"ownership_fixed_asset" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   PermissionCover    string `json:"permission_cover" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   ProfileCompletionStatus    string `json:"profile_completion_status" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   CompanyProfileID    int64 `json:"company_profile_id" required:"insert,update" min:"1" max:"9223372036854775807"`
   Code    string `json:"code" required:"insert,update" min:"1" max:"50" regex:"alphanumeric"`
   Schema    string `json:"schema" required:"insert,update" min:"1" max:"100" regex:"text_only"`
   Status    string `json:"status" required:"insert,update" min:"1" max:"50" regex:"text_only"`
   ActiveDate    time.Time `json:"active_date" required:"insert,update"`
   ResignDate    time.Time `json:"resign_date" required:"insert,update"`
   IsNexwise    bool `json:"is_nexwise" required:"insert,update"`
   DeadlineDate    time.Time `json:"deadline_date" required:"insert,update"`
   NewProfileApprovalStatus    string `json:"new_profile_approval_status" required:"insert,update" min:"1" max:"50" regex:"text_only"`
   UpdatedAt    time.Time `json:"updated_at" required:"update"`
}