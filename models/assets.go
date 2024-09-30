package models

import "gorm.io/gorm"

// AssetClassification defines the enumeration for asset classification
type AssetClassification string

const (
	// Workstation classifications
	Laptop      AssetClassification = "Laptop"
	Desktop     AssetClassification = "Desktop"
	Peripherals AssetClassification = "Peripherals"
	// Server
	Server AssetClassification = "Server"
	// Storage classifications
	Storage        AssetClassification = "Storage"
	TableLibraries AssetClassification = "Table Libraries"
	SanSwitches    AssetClassification = "SAN Switches"
	// Network classifications
	Firewall            AssetClassification = "Firewall"
	L3Switches          AssetClassification = "L3 Switches"
	L2SwitchesManaged   AssetClassification = "L2 Switches (Managed)"
	L2SwitchesUnmanaged AssetClassification = "L2 Switches (Unmanaged)"
	POEManaged          AssetClassification = "POE (Managed)"
	POEUnmanaged        AssetClassification = "POE (Unmanaged)"
)

// AllocationStatus defines the enumeration for asset allocation status
type AllocationStatus string

const (
	InUse    AllocationStatus = "In Use"
	InStore  AllocationStatus = "In Store"
	Scrap    AllocationStatus = "Scrap"
	InRepair AllocationStatus = "In Repair"
)

// Asset represents the asset model
type Asset struct {
	gorm.Model

	Classification   AssetClassification `json:"classification"`
	Make             string              `json:"make"`
	AssetModel       string              `json:"asset-model"`
	PurchaseDate     string              `json:"purchase_date"`
	Warranty         string              `json:"warranty"`
	ServiceTag       string              `json:"service_tag"`
	AssetCode        string              `json:"asset_code"`
	Location         string              `json:"location"`
	CostCenterCode   string              `json:"cost_center_code"`
	AllocationStatus AllocationStatus    `json:"allocation_status"`
	EmpID            string              `json:"emp_id"` // Foreign key to Employee
}
