
package in

import (
    "time"
)

type WardesProfileDTOIn struct {
   NexchiefAccountID    int64 `json:"nexchief_account_id" required:"insert,update" max:"999999999"`
   UserID    int64 `json:"user_id" required:"insert,update" max:"999999999"`
   Username    string `json:"username" required:"insert,update" auto_fix:"lowercase" min:"1" max:"100"`
   PersonalName    string `json:"personal_name" required:"insert,update" auto_fix:"title" min:"1" max:"100" regex:"name"`
   Gender    string `json:"gender" required:"insert,update" min:"1" max:"50"`
   Phone    string `json:"phone" required:"insert,update" min:"8" max:"20" regex:"phone"`
   Email    string `json:"email" required:"insert,update" auto_fix:"lowercase" min:"5" max:"255" regex:"email"`
   ProvinceID    int64 `json:"province_id" required:"insert,update" max:"999999999"`
   DistrictID    int64 `json:"district_id" required:"insert,update" max:"999999999"`
   SubDistrictID    int64 `json:"sub_district_id" required:"insert,update" max:"999999999"`
   UrbanVillageID    int64 `json:"urban_village_id" required:"insert,update" max:"999999999"`
   Nik    string `json:"nik" required:"insert,update" min:"1" max:"50" regex:"nik"`
   Age    string `json:"age" required:"insert,update" min:"1" max:"10"`
   Hamlet    string `json:"hamlet" required:"insert,update" auto_fix:"title" min:"1" max:"255"`
   Neighbour    string `json:"neighbour" required:"insert,update" auto_fix:"title" min:"1" max:"255"`
   VillageFamily    string `json:"village_family" required:"insert,update" min:"1" max:"255"`
   VillagePopulation    string `json:"village_population" required:"insert,update" min:"1" max:"255"`
   BuildingLocation    string `json:"building_location" required:"insert,update" min:"1" max:"255"`
   BuildingOwnership    string `json:"building_ownership" required:"insert,update" min:"1" max:"255"`
   BuildingType    string `json:"building_type" required:"insert,update" min:"1" max:"255"`
   BuildingForm    string `json:"building_form" required:"insert,update" min:"1" max:"255"`
   BuildingWidthFirstFloor    float64 `json:"building_width_first_floor" required:"insert,update" max:"999999999"`
   BuildingLengthFirstFloor    float64 `json:"building_length_first_floor" required:"insert,update" max:"999999999"`
   BuildingHeightFirstFloor    float64 `json:"building_height_first_floor" required:"insert,update" max:"999999999"`
   BuildingWidthSecondFloor    float64 `json:"building_width_second_floor" required:"insert,update" max:"999999999"`
   BuildingLengthSecondFloor    float64 `json:"building_length_second_floor" required:"insert,update" max:"999999999"`
   BuildingHeightSecondFloor    float64 `json:"building_height_second_floor" required:"insert,update" max:"999999999"`
   ParkingLotWidth    float64 `json:"parking_lot_width" required:"insert,update" max:"999999999"`
   ParkingLotLength    float64 `json:"parking_lot_length" required:"insert,update" max:"999999999"`
   VillageInternetConnection    string `json:"village_internet_connection" required:"insert,update" min:"1" max:"255"`
   InvestmentCapital    string `json:"investment_capital" required:"insert,update" min:"1" max:"255"`
   InvestmentCapitalSource    string `json:"investment_capital_source" required:"insert,update" min:"1" max:"255"`
   OwnershipStatus    string `json:"ownership_status" required:"insert,update" min:"1" max:"255"`
   OwnershipFixedAsset    string `json:"ownership_fixed_asset" required:"insert,update" min:"1" max:"255"`
   PermissionCover    string `json:"permission_cover" required:"insert,update" min:"1" max:"255"`
   ProfileCompletionStatus    string `json:"profile_completion_status" required:"insert,update" min:"1" max:"255" regex:"record_status"`
   CompanyProfileID    int64 `json:"company_profile_id" required:"insert,update" max:"999999999"`
   Code    string `json:"code" required:"insert,update" auto_fix:"uppercase" min:"1" max:"100"`
   Schema    string `json:"schema" required:"insert,update" min:"1" max:"255"`
   Status    string `json:"status" required:"insert,update" min:"1" max:"50" regex:"record_status"`
   ActiveDateStr    string `json:"active_date"`
   ActiveDate    time.Time `required:"insert,update"`
   ResignDateStr    string `json:"resign_date"`
   ResignDate    time.Time `required:"insert,update"`
   IsNexwise    bool `json:"is_nexwise" required:"insert,update"`
   DeadlineDateStr    string `json:"deadline_date"`
   DeadlineDate    time.Time `required:"insert,update"`
   UpdatedClient    string `json:"updated_client" required:"insert,update" min:"1" max:"100"`
   NewProfileApprovalStatus    string `json:"new_profile_approval_status" required:"insert,update" min:"1" max:"100"`
}