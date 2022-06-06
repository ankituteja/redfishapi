package redfishapi

//Dell Based Structs

//SysAttrDell ... System Attributes from the Redfish API
type SysAttrDell struct {
	RedfishSettings struct {
		odataContext   string
		odataType      string
		SettingsObject struct {
			odataID string
		} `json:"SettingsObject"`
	}
	odataContext      string
	odataID           string
	odataType         string
	AttributeRegistry string            `json:"AttributeRegistry"`
	Attributes        SysAttributesData `json:"Attributes"`
	Description       string            `json:"Description"`
	ID                string            `json:"Id"`
	Name              string            `json:"Name"`
}

type SysAttributesData struct {
	AcquisitionInfo_1_CostCenter                                         string `json:"AcquisitionInfo.1.CostCenter"`
	AcquisitionInfo_1_Expensed                                           string `json:"AcquisitionInfo.1.Expensed"`
	AcquisitionInfo_1_InstallDate                                        string `json:"AcquisitionInfo.1.InstallDate"`
	AcquisitionInfo_1_PONumber                                           string `json:"AcquisitionInfo.1.PONumber"`
	AcquisitionInfo_1_PurchaseCost                                       string `json:"AcquisitionInfo.1.PurchaseCost"`
	AcquisitionInfo_1_PurchaseDate                                       string `json:"AcquisitionInfo.1.PurchaseDate"`
	AcquisitionInfo_1_Vendor                                             string `json:"AcquisitionInfo.1.Vendor"`
	AcquisitionInfo_1_WayBill                                            string `json:"AcquisitionInfo.1.WayBill"`
	AcquisitionInfo_1_WhoSigned                                          string `json:"AcquisitionInfo.1.WhoSigned"`
	Backplane_1_BackplaneBusMode                                         string `json:"Backplane.1.BackplaneBusMode"`
	Backplane_1_BackplaneSplitMode                                       int    `json:"Backplane.1.BackplaneSplitMode"`
	CustomAsset_1_Label                                                  string `json:"CustomAsset.1.Label"`
	CustomAsset_1_Value                                                  string `json:"CustomAsset.1.Value"`
	CustomAsset_10_Label                                                 string `json:"CustomAsset.10.Label"`
	CustomAsset_10_Value                                                 string `json:"CustomAsset.10.Value"`
	CustomAsset_11_Label                                                 string `json:"CustomAsset.11.Label"`
	CustomAsset_11_Value                                                 string `json:"CustomAsset.11.Value"`
	CustomAsset_12_Label                                                 string `json:"CustomAsset.12.Label"`
	CustomAsset_12_Value                                                 string `json:"CustomAsset.12.Value"`
	CustomAsset_13_Label                                                 string `json:"CustomAsset.13.Label"`
	CustomAsset_13_Value                                                 string `json:"CustomAsset.13.Value"`
	CustomAsset_14_Label                                                 string `json:"CustomAsset.14.Label"`
	CustomAsset_14_Value                                                 string `json:"CustomAsset.14.Value"`
	CustomAsset_15_Label                                                 string `json:"CustomAsset.15.Label"`
	CustomAsset_15_Value                                                 string `json:"CustomAsset.15.Value"`
	CustomAsset_16_Label                                                 string `json:"CustomAsset.16.Label"`
	CustomAsset_16_Value                                                 string `json:"CustomAsset.16.Value"`
	CustomAsset_17_Label                                                 string `json:"CustomAsset.17.Label"`
	CustomAsset_17_Value                                                 string `json:"CustomAsset.17.Value"`
	CustomAsset_18_Label                                                 string `json:"CustomAsset.18.Label"`
	CustomAsset_18_Value                                                 string `json:"CustomAsset.18.Value"`
	CustomAsset_19_Label                                                 string `json:"CustomAsset.19.Label"`
	CustomAsset_19_Value                                                 string `json:"CustomAsset.19.Value"`
	CustomAsset_2_Label                                                  string `json:"CustomAsset.2.Label"`
	CustomAsset_2_Value                                                  string `json:"CustomAsset.2.Value"`
	CustomAsset_20_Label                                                 string `json:"CustomAsset.20.Label"`
	CustomAsset_20_Value                                                 string `json:"CustomAsset.20.Value"`
	CustomAsset_21_Label                                                 string `json:"CustomAsset.21.Label"`
	CustomAsset_21_Value                                                 string `json:"CustomAsset.21.Value"`
	CustomAsset_22_Label                                                 string `json:"CustomAsset.22.Label"`
	CustomAsset_22_Value                                                 string `json:"CustomAsset.22.Value"`
	CustomAsset_23_Label                                                 string `json:"CustomAsset.23.Label"`
	CustomAsset_23_Value                                                 string `json:"CustomAsset.23.Value"`
	CustomAsset_24_Label                                                 string `json:"CustomAsset.24.Label"`
	CustomAsset_24_Value                                                 string `json:"CustomAsset.24.Value"`
	CustomAsset_25_Label                                                 string `json:"CustomAsset.25.Label"`
	CustomAsset_25_Value                                                 string `json:"CustomAsset.25.Value"`
	CustomAsset_26_Label                                                 string `json:"CustomAsset.26.Label"`
	CustomAsset_26_Value                                                 string `json:"CustomAsset.26.Value"`
	CustomAsset_27_Label                                                 string `json:"CustomAsset.27.Label"`
	CustomAsset_27_Value                                                 string `json:"CustomAsset.27.Value"`
	CustomAsset_28_Label                                                 string `json:"CustomAsset.28.Label"`
	CustomAsset_28_Value                                                 string `json:"CustomAsset.28.Value"`
	CustomAsset_29_Label                                                 string `json:"CustomAsset.29.Label"`
	CustomAsset_29_Value                                                 string `json:"CustomAsset.29.Value"`
	CustomAsset_3_Label                                                  string `json:"CustomAsset.3.Label"`
	CustomAsset_3_Value                                                  string `json:"CustomAsset.3.Value"`
	CustomAsset_30_Label                                                 string `json:"CustomAsset.30.Label"`
	CustomAsset_30_Value                                                 string `json:"CustomAsset.30.Value"`
	CustomAsset_31_Label                                                 string `json:"CustomAsset.31.Label"`
	CustomAsset_31_Value                                                 string `json:"CustomAsset.31.Value"`
	CustomAsset_32_Label                                                 string `json:"CustomAsset.32.Label"`
	CustomAsset_32_Value                                                 string `json:"CustomAsset.32.Value"`
	CustomAsset_4_Label                                                  string `json:"CustomAsset.4.Label"`
	CustomAsset_4_Value                                                  string `json:"CustomAsset.4.Value"`
	CustomAsset_5_Label                                                  string `json:"CustomAsset.5.Label"`
	CustomAsset_5_Value                                                  string `json:"CustomAsset.5.Value"`
	CustomAsset_6_Label                                                  string `json:"CustomAsset.6.Label"`
	CustomAsset_6_Value                                                  string `json:"CustomAsset.6.Value"`
	CustomAsset_7_Label                                                  string `json:"CustomAsset.7.Label"`
	CustomAsset_7_Value                                                  string `json:"CustomAsset.7.Value"`
	CustomAsset_8_Label                                                  string `json:"CustomAsset.8.Label"`
	CustomAsset_8_Value                                                  string `json:"CustomAsset.8.Value"`
	CustomAsset_9_Label                                                  string `json:"CustomAsset.9.Label"`
	CustomAsset_9_Value                                                  string `json:"CustomAsset.9.Value"`
	DepreciationInfo_1_Duration                                          string `json:"DepreciationInfo.1.Duration"`
	DepreciationInfo_1_DurationUnit                                      string `json:"DepreciationInfo.1.DurationUnit"`
	DepreciationInfo_1_Method                                            string `json:"DepreciationInfo.1.Method"`
	DepreciationInfo_1_Percentage                                        string `json:"DepreciationInfo.1.Percentage"`
	Diagnostics_1_OSAppCollectionTime                                    string `json:"Diagnostics.1.OSAppCollectionTime"`
	ExtWarrantyInfo_1_Cost                                               string `json:"ExtWarrantyInfo.1.Cost"`
	ExtWarrantyInfo_1_EndDate                                            string `json:"ExtWarrantyInfo.1.EndDate"`
	ExtWarrantyInfo_1_Provider                                           string `json:"ExtWarrantyInfo.1.Provider"`
	ExtWarrantyInfo_1_StartDate                                          string `json:"ExtWarrantyInfo.1.StartDate"`
	LCD_1_ChassisIdentifyDuration                                        int    `json:"LCD.1.ChassisIdentifyDuration"`
	LCD_1_Configuration                                                  string `json:"LCD.1.Configuration"`
	LCD_1_CurrentDisplay                                                 string `json:"LCD.1.CurrentDisplay"`
	LCD_1_ErrorDisplayMode                                               string `json:"LCD.1.ErrorDisplayMode"`
	LCD_1_FrontPanelLocking                                              string `json:"LCD.1.FrontPanelLocking"`
	LCD_1_HideErrs                                                       string `json:"LCD.1.HideErrs"`
	LCD_1_LicenseMsgEnable                                               string `json:"LCD.1.LicenseMsgEnable"`
	LCD_1_NMIResetOverride                                               string `json:"LCD.1.NMIResetOverride"`
	LCD_1_NumberErrsHidden                                               int    `json:"LCD.1.NumberErrsHidden"`
	LCD_1_NumberErrsVisible                                              int    `json:"LCD.1.NumberErrsVisible"`
	LCD_1_QualifierTemp                                                  string `json:"LCD.1.QualifierTemp"`
	LCD_1_QualifierWatt                                                  string `json:"LCD.1.QualifierWatt"`
	LCD_1_UserDefinedString                                              string `json:"LCD.1.UserDefinedString"`
	LCD_1_vConsoleIndication                                             string `json:"LCD.1.vConsoleIndication"`
	LeaseInfo_1_Buyout                                                   string `json:"LeaseInfo.1.Buyout"`
	LeaseInfo_1_EndDate                                                  string `json:"LeaseInfo.1.EndDate"`
	LeaseInfo_1_FairValue                                                string `json:"LeaseInfo.1.FairValue"`
	LeaseInfo_1_Lessor                                                   string `json:"LeaseInfo.1.Lessor"`
	LeaseInfo_1_MultiSched                                               string `json:"LeaseInfo.1.MultiSched"`
	LeaseInfo_1_RateFactor                                               string `json:"LeaseInfo.1.RateFactor"`
	MaintenanceInfo_1_EndDate                                            string `json:"MaintenanceInfo.1.EndDate"`
	MaintenanceInfo_1_Provider                                           string `json:"MaintenanceInfo.1.Provider"`
	MaintenanceInfo_1_Restrictions                                       string `json:"MaintenanceInfo.1.Restrictions"`
	MaintenanceInfo_1_StartDate                                          string `json:"MaintenanceInfo.1.StartDate"`
	OutsourceInfo_1_ProbComp                                             string `json:"OutsourceInfo.1.ProbComp"`
	OutsourceInfo_1_ProviderFee                                          string `json:"OutsourceInfo.1.ProviderFee"`
	OutsourceInfo_1_SLALevel                                             string `json:"OutsourceInfo.1.SLALevel"`
	OutsourceInfo_1_ServiceFee                                           string `json:"OutsourceInfo.1.ServiceFee"`
	OutsourceInfo_1_SignedFor                                            string `json:"OutsourceInfo.1.SignedFor"`
	OwnerInfo_1_InsComp                                                  string `json:"OwnerInfo.1.InsComp"`
	OwnerInfo_1_OwnerName                                                string `json:"OwnerInfo.1.OwnerName"`
	OwnerInfo_1_Type                                                     string `json:"OwnerInfo.1.Type"`
	PCIeSlotLFM_1_3rdPartyCard                                           string `json:"PCIeSlotLFM.1.3rdPartyCard"`
	PCIeSlotLFM_1_CardType                                               string `json:"PCIeSlotLFM.1.CardType"`
	PCIeSlotLFM_1_CustomLFM                                              int    `json:"PCIeSlotLFM.1.CustomLFM"`
	PCIeSlotLFM_1_LFMMode                                                string `json:"PCIeSlotLFM.1.LFMMode"`
	PCIeSlotLFM_1_MaxLFM                                                 int    `json:"PCIeSlotLFM.1.MaxLFM"`
	PCIeSlotLFM_1_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.1.PCIeInletTemperature"`
	PCIeSlotLFM_1_SlotState                                              string `json:"PCIeSlotLFM.1.SlotState"`
	PCIeSlotLFM_1_TargetLFM                                              string `json:"PCIeSlotLFM.1.TargetLFM"`
	PCIeSlotLFM_10_3rdPartyCard                                          string `json:"PCIeSlotLFM.10.3rdPartyCard"`
	PCIeSlotLFM_10_CardType                                              string `json:"PCIeSlotLFM.10.CardType"`
	PCIeSlotLFM_10_CustomLFM                                             int    `json:"PCIeSlotLFM.10.CustomLFM"`
	PCIeSlotLFM_10_LFMMode                                               string `json:"PCIeSlotLFM.10.LFMMode"`
	PCIeSlotLFM_10_MaxLFM                                                int    `json:"PCIeSlotLFM.10.MaxLFM"`
	PCIeSlotLFM_10_PCIeInletTemperature                                  int    `json:"PCIeSlotLFM.10.PCIeInletTemperature"`
	PCIeSlotLFM_10_SlotState                                             string `json:"PCIeSlotLFM.10.SlotState"`
	PCIeSlotLFM_10_TargetLFM                                             string `json:"PCIeSlotLFM.10.TargetLFM"`
	PCIeSlotLFM_11_3rdPartyCard                                          string `json:"PCIeSlotLFM.11.3rdPartyCard"`
	PCIeSlotLFM_11_CardType                                              string `json:"PCIeSlotLFM.11.CardType"`
	PCIeSlotLFM_11_CustomLFM                                             int    `json:"PCIeSlotLFM.11.CustomLFM"`
	PCIeSlotLFM_11_LFMMode                                               string `json:"PCIeSlotLFM.11.LFMMode"`
	PCIeSlotLFM_11_MaxLFM                                                int    `json:"PCIeSlotLFM.11.MaxLFM"`
	PCIeSlotLFM_11_PCIeInletTemperature                                  int    `json:"PCIeSlotLFM.11.PCIeInletTemperature"`
	PCIeSlotLFM_11_SlotState                                             string `json:"PCIeSlotLFM.11.SlotState"`
	PCIeSlotLFM_11_TargetLFM                                             string `json:"PCIeSlotLFM.11.TargetLFM"`
	PCIeSlotLFM_12_3rdPartyCard                                          string `json:"PCIeSlotLFM.12.3rdPartyCard"`
	PCIeSlotLFM_12_CardType                                              string `json:"PCIeSlotLFM.12.CardType"`
	PCIeSlotLFM_12_CustomLFM                                             int    `json:"PCIeSlotLFM.12.CustomLFM"`
	PCIeSlotLFM_12_LFMMode                                               string `json:"PCIeSlotLFM.12.LFMMode"`
	PCIeSlotLFM_12_MaxLFM                                                int    `json:"PCIeSlotLFM.12.MaxLFM"`
	PCIeSlotLFM_12_PCIeInletTemperature                                  int    `json:"PCIeSlotLFM.12.PCIeInletTemperature"`
	PCIeSlotLFM_12_SlotState                                             string `json:"PCIeSlotLFM.12.SlotState"`
	PCIeSlotLFM_12_TargetLFM                                             string `json:"PCIeSlotLFM.12.TargetLFM"`
	PCIeSlotLFM_13_3rdPartyCard                                          string `json:"PCIeSlotLFM.13.3rdPartyCard"`
	PCIeSlotLFM_13_CardType                                              string `json:"PCIeSlotLFM.13.CardType"`
	PCIeSlotLFM_13_CustomLFM                                             int    `json:"PCIeSlotLFM.13.CustomLFM"`
	PCIeSlotLFM_13_LFMMode                                               string `json:"PCIeSlotLFM.13.LFMMode"`
	PCIeSlotLFM_13_MaxLFM                                                int    `json:"PCIeSlotLFM.13.MaxLFM"`
	PCIeSlotLFM_13_PCIeInletTemperature                                  int    `json:"PCIeSlotLFM.13.PCIeInletTemperature"`
	PCIeSlotLFM_13_SlotState                                             string `json:"PCIeSlotLFM.13.SlotState"`
	PCIeSlotLFM_13_TargetLFM                                             string `json:"PCIeSlotLFM.13.TargetLFM"`
	PCIeSlotLFM_14_3rdPartyCard                                          string `json:"PCIeSlotLFM.14.3rdPartyCard"`
	PCIeSlotLFM_14_CardType                                              string `json:"PCIeSlotLFM.14.CardType"`
	PCIeSlotLFM_14_CustomLFM                                             int    `json:"PCIeSlotLFM.14.CustomLFM"`
	PCIeSlotLFM_14_LFMMode                                               string `json:"PCIeSlotLFM.14.LFMMode"`
	PCIeSlotLFM_14_MaxLFM                                                int    `json:"PCIeSlotLFM.14.MaxLFM"`
	PCIeSlotLFM_14_PCIeInletTemperature                                  int    `json:"PCIeSlotLFM.14.PCIeInletTemperature"`
	PCIeSlotLFM_14_SlotState                                             string `json:"PCIeSlotLFM.14.SlotState"`
	PCIeSlotLFM_14_TargetLFM                                             string `json:"PCIeSlotLFM.14.TargetLFM"`
	PCIeSlotLFM_15_3rdPartyCard                                          string `json:"PCIeSlotLFM.15.3rdPartyCard"`
	PCIeSlotLFM_15_CardType                                              string `json:"PCIeSlotLFM.15.CardType"`
	PCIeSlotLFM_15_CustomLFM                                             int    `json:"PCIeSlotLFM.15.CustomLFM"`
	PCIeSlotLFM_15_LFMMode                                               string `json:"PCIeSlotLFM.15.LFMMode"`
	PCIeSlotLFM_15_MaxLFM                                                int    `json:"PCIeSlotLFM.15.MaxLFM"`
	PCIeSlotLFM_15_PCIeInletTemperature                                  int    `json:"PCIeSlotLFM.15.PCIeInletTemperature"`
	PCIeSlotLFM_15_SlotState                                             string `json:"PCIeSlotLFM.15.SlotState"`
	PCIeSlotLFM_15_TargetLFM                                             string `json:"PCIeSlotLFM.15.TargetLFM"`
	PCIeSlotLFM_2_3rdPartyCard                                           string `json:"PCIeSlotLFM.2.3rdPartyCard"`
	PCIeSlotLFM_2_CardType                                               string `json:"PCIeSlotLFM.2.CardType"`
	PCIeSlotLFM_2_CustomLFM                                              int    `json:"PCIeSlotLFM.2.CustomLFM"`
	PCIeSlotLFM_2_LFMMode                                                string `json:"PCIeSlotLFM.2.LFMMode"`
	PCIeSlotLFM_2_MaxLFM                                                 int    `json:"PCIeSlotLFM.2.MaxLFM"`
	PCIeSlotLFM_2_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.2.PCIeInletTemperature"`
	PCIeSlotLFM_2_SlotState                                              string `json:"PCIeSlotLFM.2.SlotState"`
	PCIeSlotLFM_2_TargetLFM                                              string `json:"PCIeSlotLFM.2.TargetLFM"`
	PCIeSlotLFM_3_3rdPartyCard                                           string `json:"PCIeSlotLFM.3.3rdPartyCard"`
	PCIeSlotLFM_3_CardType                                               string `json:"PCIeSlotLFM.3.CardType"`
	PCIeSlotLFM_3_CustomLFM                                              int    `json:"PCIeSlotLFM.3.CustomLFM"`
	PCIeSlotLFM_3_LFMMode                                                string `json:"PCIeSlotLFM.3.LFMMode"`
	PCIeSlotLFM_3_MaxLFM                                                 int    `json:"PCIeSlotLFM.3.MaxLFM"`
	PCIeSlotLFM_3_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.3.PCIeInletTemperature"`
	PCIeSlotLFM_3_SlotState                                              string `json:"PCIeSlotLFM.3.SlotState"`
	PCIeSlotLFM_3_TargetLFM                                              string `json:"PCIeSlotLFM.3.TargetLFM"`
	PCIeSlotLFM_4_3rdPartyCard                                           string `json:"PCIeSlotLFM.4.3rdPartyCard"`
	PCIeSlotLFM_4_CardType                                               string `json:"PCIeSlotLFM.4.CardType"`
	PCIeSlotLFM_4_CustomLFM                                              int    `json:"PCIeSlotLFM.4.CustomLFM"`
	PCIeSlotLFM_4_LFMMode                                                string `json:"PCIeSlotLFM.4.LFMMode"`
	PCIeSlotLFM_4_MaxLFM                                                 int    `json:"PCIeSlotLFM.4.MaxLFM"`
	PCIeSlotLFM_4_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.4.PCIeInletTemperature"`
	PCIeSlotLFM_4_SlotState                                              string `json:"PCIeSlotLFM.4.SlotState"`
	PCIeSlotLFM_4_TargetLFM                                              string `json:"PCIeSlotLFM.4.TargetLFM"`
	PCIeSlotLFM_5_3rdPartyCard                                           string `json:"PCIeSlotLFM.5.3rdPartyCard"`
	PCIeSlotLFM_5_CardType                                               string `json:"PCIeSlotLFM.5.CardType"`
	PCIeSlotLFM_5_CustomLFM                                              int    `json:"PCIeSlotLFM.5.CustomLFM"`
	PCIeSlotLFM_5_LFMMode                                                string `json:"PCIeSlotLFM.5.LFMMode"`
	PCIeSlotLFM_5_MaxLFM                                                 int    `json:"PCIeSlotLFM.5.MaxLFM"`
	PCIeSlotLFM_5_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.5.PCIeInletTemperature"`
	PCIeSlotLFM_5_SlotState                                              string `json:"PCIeSlotLFM.5.SlotState"`
	PCIeSlotLFM_5_TargetLFM                                              string `json:"PCIeSlotLFM.5.TargetLFM"`
	PCIeSlotLFM_6_3rdPartyCard                                           string `json:"PCIeSlotLFM.6.3rdPartyCard"`
	PCIeSlotLFM_6_CardType                                               string `json:"PCIeSlotLFM.6.CardType"`
	PCIeSlotLFM_6_CustomLFM                                              int    `json:"PCIeSlotLFM.6.CustomLFM"`
	PCIeSlotLFM_6_LFMMode                                                string `json:"PCIeSlotLFM.6.LFMMode"`
	PCIeSlotLFM_6_MaxLFM                                                 int    `json:"PCIeSlotLFM.6.MaxLFM"`
	PCIeSlotLFM_6_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.6.PCIeInletTemperature"`
	PCIeSlotLFM_6_SlotState                                              string `json:"PCIeSlotLFM.6.SlotState"`
	PCIeSlotLFM_6_TargetLFM                                              string `json:"PCIeSlotLFM.6.TargetLFM"`
	PCIeSlotLFM_7_3rdPartyCard                                           string `json:"PCIeSlotLFM.7.3rdPartyCard"`
	PCIeSlotLFM_7_CardType                                               string `json:"PCIeSlotLFM.7.CardType"`
	PCIeSlotLFM_7_CustomLFM                                              int    `json:"PCIeSlotLFM.7.CustomLFM"`
	PCIeSlotLFM_7_LFMMode                                                string `json:"PCIeSlotLFM.7.LFMMode"`
	PCIeSlotLFM_7_MaxLFM                                                 int    `json:"PCIeSlotLFM.7.MaxLFM"`
	PCIeSlotLFM_7_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.7.PCIeInletTemperature"`
	PCIeSlotLFM_7_SlotState                                              string `json:"PCIeSlotLFM.7.SlotState"`
	PCIeSlotLFM_7_TargetLFM                                              string `json:"PCIeSlotLFM.7.TargetLFM"`
	PCIeSlotLFM_8_3rdPartyCard                                           string `json:"PCIeSlotLFM.8.3rdPartyCard"`
	PCIeSlotLFM_8_CardType                                               string `json:"PCIeSlotLFM.8.CardType"`
	PCIeSlotLFM_8_CustomLFM                                              int    `json:"PCIeSlotLFM.8.CustomLFM"`
	PCIeSlotLFM_8_LFMMode                                                string `json:"PCIeSlotLFM.8.LFMMode"`
	PCIeSlotLFM_8_MaxLFM                                                 int    `json:"PCIeSlotLFM.8.MaxLFM"`
	PCIeSlotLFM_8_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.8.PCIeInletTemperature"`
	PCIeSlotLFM_8_SlotState                                              string `json:"PCIeSlotLFM.8.SlotState"`
	PCIeSlotLFM_8_TargetLFM                                              string `json:"PCIeSlotLFM.8.TargetLFM"`
	PCIeSlotLFM_9_3rdPartyCard                                           string `json:"PCIeSlotLFM.9.3rdPartyCard"`
	PCIeSlotLFM_9_CardType                                               string `json:"PCIeSlotLFM.9.CardType"`
	PCIeSlotLFM_9_CustomLFM                                              int    `json:"PCIeSlotLFM.9.CustomLFM"`
	PCIeSlotLFM_9_LFMMode                                                string `json:"PCIeSlotLFM.9.LFMMode"`
	PCIeSlotLFM_9_MaxLFM                                                 int    `json:"PCIeSlotLFM.9.MaxLFM"`
	PCIeSlotLFM_9_PCIeInletTemperature                                   int    `json:"PCIeSlotLFM.9.PCIeInletTemperature"`
	PCIeSlotLFM_9_SlotState                                              string `json:"PCIeSlotLFM.9.SlotState"`
	PCIeSlotLFM_9_TargetLFM                                              string `json:"PCIeSlotLFM.9.TargetLFM"`
	PowerHistorical_1_IntervalInSeconds                                  int    `json:"PowerHistorical.1.IntervalInSeconds"`
	QuickSync_1_Access                                                   string `json:"QuickSync.1.Access"`
	QuickSync_1_InactivityTimeout                                        int    `json:"QuickSync.1.InactivityTimeout"`
	QuickSync_1_InactivityTimerEnable                                    string `json:"QuickSync.1.InactivityTimerEnable"`
	QuickSync_1_Presence                                                 string `json:"QuickSync.1.Presence"`
	QuickSync_1_ReadAuthentication                                       string `json:"QuickSync.1.ReadAuthentication"`
	QuickSync_1_WifiEnable                                               string `json:"QuickSync.1.WifiEnable"`
	SC_BMC_1_PowerMonitoring                                             int    `json:"SC-BMC.1.PowerMonitoring"`
	ServerInfo_1_NodeID                                                  string `json:"ServerInfo.1.NodeID"`
	ServerInfo_1_RChassisServiceTag                                      string `json:"ServerInfo.1.RChassisServiceTag"`
	ServerInfo_1_ServerType                                              string `json:"ServerInfo.1.ServerType"`
	ServerInfo_1_ServiceTag                                              string `json:"ServerInfo.1.ServiceTag"`
	ServerOS_1_HostName                                                  string `json:"ServerOS.1.HostName"`
	ServerOS_1_OEMOSVersion                                              string `json:"ServerOS.1.OEMOSVersion"`
	ServerOS_1_OSName                                                    string `json:"ServerOS.1.OSName"`
	ServerOS_1_OSVersion                                                 string `json:"ServerOS.1.OSVersion"`
	ServerOS_1_ProductKey                                                string `json:"ServerOS.1.ProductKey"`
	ServerOS_1_ServerPoweredOnTime                                       int    `json:"ServerOS.1.ServerPoweredOnTime"`
	ServerPwr_1_ActivePolicyName                                         string `json:"ServerPwr.1.ActivePolicyName"`
	ServerPwr_1_ActivePowerCapVal                                        int    `json:"ServerPwr.1.ActivePowerCapVal"`
	ServerPwr_1_PSPFCEnabled                                             string `json:"ServerPwr.1.PSPFCEnabled"`
	ServerPwr_1_PSRapidOn                                                string `json:"ServerPwr.1.PSRapidOn"`
	ServerPwr_1_PSRedPolicy                                              string `json:"ServerPwr.1.PSRedPolicy"`
	ServerPwr_1_PowerCapMaxThres                                         int    `json:"ServerPwr.1.PowerCapMaxThres"`
	ServerPwr_1_PowerCapMinThres                                         int    `json:"ServerPwr.1.PowerCapMinThres"`
	ServerPwr_1_PowerCapSetting                                          string `json:"ServerPwr.1.PowerCapSetting"`
	ServerPwr_1_PowerCapValue                                            int    `json:"ServerPwr.1.PowerCapValue"`
	ServerPwr_1_RapidOnPrimaryPSU                                        string `json:"ServerPwr.1.RapidOnPrimaryPSU"`
	ServerPwrMon_1_AccumulativePower                                     int    `json:"ServerPwrMon.1.AccumulativePower"`
	ServerPwrMon_1_CumulativePowerStartTime                              int    `json:"ServerPwrMon.1.CumulativePowerStartTime"`
	ServerPwrMon_1_CumulativePowerStartTimeStr                           string `json:"ServerPwrMon.1.CumulativePowerStartTimeStr"`
	ServerPwrMon_1_MinPowerTime                                          int    `json:"ServerPwrMon.1.MinPowerTime"`
	ServerPwrMon_1_MinPowerTimeStr                                       string `json:"ServerPwrMon.1.MinPowerTimeStr"`
	ServerPwrMon_1_MinPowerWatts                                         int    `json:"ServerPwrMon.1.MinPowerWatts"`
	ServerPwrMon_1_PeakCurrentTime                                       int    `json:"ServerPwrMon.1.PeakCurrentTime"`
	ServerPwrMon_1_PeakCurrentTimeStr                                    string `json:"ServerPwrMon.1.PeakCurrentTimeStr"`
	ServerPwrMon_1_PeakPowerStartTime                                    int    `json:"ServerPwrMon.1.PeakPowerStartTime"`
	ServerPwrMon_1_PeakPowerStartTimeStr                                 string `json:"ServerPwrMon.1.PeakPowerStartTimeStr"`
	ServerPwrMon_1_PeakPowerTime                                         int    `json:"ServerPwrMon.1.PeakPowerTime"`
	ServerPwrMon_1_PeakPowerTimeStr                                      string `json:"ServerPwrMon.1.PeakPowerTimeStr"`
	ServerPwrMon_1_PeakPowerWatts                                        int    `json:"ServerPwrMon.1.PeakPowerWatts"`
	ServerPwrMon_1_PowerConfigReset                                      string `json:"ServerPwrMon.1.PowerConfigReset"`
	ServerTopology_1_AisleName                                           string `json:"ServerTopology.1.AisleName"`
	ServerTopology_1_DataCenterName                                      string `json:"ServerTopology.1.DataCenterName"`
	ServerTopology_1_RackName                                            string `json:"ServerTopology.1.RackName"`
	ServerTopology_1_RackSlot                                            int    `json:"ServerTopology.1.RackSlot"`
	ServerTopology_1_RoomName                                            string `json:"ServerTopology.1.RoomName"`
	ServerTopology_1_SizeOfManagedSystemInU                              int    `json:"ServerTopology.1.SizeOfManagedSystemInU"`
	ServiceContract_1_Renewed                                            string `json:"ServiceContract.1.Renewed"`
	ServiceContract_1_Type                                               string `json:"ServiceContract.1.Type"`
	ServiceContract_1_Vendor                                             string `json:"ServiceContract.1.Vendor"`
	Storage_1_AvailableSpareAlertThreshold                               int    `json:"Storage.1.AvailableSpareAlertThreshold"`
	Storage_1_RemainingRatedWriteEnduranceAlertThreshold                 int    `json:"Storage.1.RemainingRatedWriteEnduranceAlertThreshold"`
	SupportInfo_1_AutoFix                                                string `json:"SupportInfo.1.AutoFix"`
	SupportInfo_1_HelpDesk                                               string `json:"SupportInfo.1.HelpDesk"`
	SupportInfo_1_Outsourced                                             string `json:"SupportInfo.1.Outsourced"`
	SupportInfo_1_Type                                                   string `json:"SupportInfo.1.Type"`
	SystemInfo_1_BootTime                                                string `json:"SystemInfo.1.BootTime"`
	SystemInfo_1_PrimaryTelephone                                        string `json:"SystemInfo.1.PrimaryTelephone"`
	SystemInfo_1_PrimaryUser                                             string `json:"SystemInfo.1.PrimaryUser"`
	SystemInfo_1_SysLocation                                             string `json:"SystemInfo.1.SysLocation"`
	SystemInfo_1_SysTime                                                 string `json:"SystemInfo.1.SysTime"`
	ThermalConfig_1_CriticalEventGenerationInterval                      int    `json:"ThermalConfig.1.CriticalEventGenerationInterval"`
	ThermalConfig_1_EventGenerationInterval                              int    `json:"ThermalConfig.1.EventGenerationInterval"`
	ThermalConfig_1_FreshAirCompliantConfiguration                       string `json:"ThermalConfig.1.FreshAirCompliantConfiguration"`
	ThermalConfig_1_MaxCFM                                               int    `json:"ThermalConfig.1.MaxCFM"`
	ThermalConfig_1_ValidFanConfiguration                                string `json:"ThermalConfig.1.ValidFanConfiguration"`
	ThermalHistorical_1_IntervalInSeconds                                int    `json:"ThermalHistorical.1.IntervalInSeconds"`
	ThermalSettings_1_AirExhaustTemp                                     string `json:"ThermalSettings.1.AirExhaustTemp"`
	ThermalSettings_1_AirExhaustTempSupport                              string `json:"ThermalSettings.1.AirExhaustTempSupport"`
	ThermalSettings_1_AirTemperatureRiseLimit                            string `json:"ThermalSettings.1.AirTemperatureRiseLimit"`
	ThermalSettings_1_AirTemperatureRiseLimitSupport                     string `json:"ThermalSettings.1.AirTemperatureRiseLimitSupport"`
	ThermalSettings_1_CurrentSystemProfileValue                          string `json:"ThermalSettings.1.CurrentSystemProfileValue"`
	ThermalSettings_1_FanSpeedHighOffsetVal                              int    `json:"ThermalSettings.1.FanSpeedHighOffsetVal"`
	ThermalSettings_1_FanSpeedLowOffsetVal                               int    `json:"ThermalSettings.1.FanSpeedLowOffsetVal"`
	ThermalSettings_1_FanSpeedMaxOffsetVal                               int    `json:"ThermalSettings.1.FanSpeedMaxOffsetVal"`
	ThermalSettings_1_FanSpeedMediumOffsetVal                            int    `json:"ThermalSettings.1.FanSpeedMediumOffsetVal"`
	ThermalSettings_1_FanSpeedOffset                                     string `json:"ThermalSettings.1.FanSpeedOffset"`
	ThermalSettings_1_MFSMaximumLimit                                    int    `json:"ThermalSettings.1.MFSMaximumLimit"`
	ThermalSettings_1_MFSMinimumLimit                                    int    `json:"ThermalSettings.1.MFSMinimumLimit"`
	ThermalSettings_1_MaximumPCIeInletTemperatureLimit                   string `json:"ThermalSettings.1.MaximumPCIeInletTemperatureLimit"`
	ThermalSettings_1_MaximumPCIeInletTemperatureLimitSupport            string `json:"ThermalSettings.1.MaximumPCIeInletTemperatureLimitSupport"`
	ThermalSettings_1_MinimumFanSpeed                                    int    `json:"ThermalSettings.1.MinimumFanSpeed"`
	ThermalSettings_1_PCIeSlotLFMSupport                                 string `json:"ThermalSettings.1.PCIeSlotLFMSupport"`
	ThermalSettings_1_SetAirTemperatureRiseLimit                         string `json:"ThermalSettings.1.SetAirTemperatureRiseLimit"`
	ThermalSettings_1_SetMaximumExhaustTemperatureLimit                  string `json:"ThermalSettings.1.SetMaximumExhaustTemperatureLimit"`
	ThermalSettings_1_SystemCFMSupport                                   string `json:"ThermalSettings.1.SystemCFMSupport"`
	ThermalSettings_1_SystemExhaustTemperature                           int    `json:"ThermalSettings.1.SystemExhaustTemperature"`
	ThermalSettings_1_SystemInletTemperature                             int    `json:"ThermalSettings.1.SystemInletTemperature"`
	ThermalSettings_1_SystemInletTemperatureSupportLimitPerConfiguration int    `json:"ThermalSettings.1.SystemInletTemperatureSupportLimitPerConfiguration"`
	ThermalSettings_1_TargetExhaustTemperatureLimit                      int    `json:"ThermalSettings.1.TargetExhaustTemperatureLimit"`
	ThermalSettings_1_ThermalProfile                                     string `json:"ThermalSettings.1.ThermalProfile"`
	USBFront_1_Enable                                                    string `json:"USBFront.1.Enable"`
	WarrantyInfo_1_Cost                                                  string `json:"WarrantyInfo.1.Cost"`
	WarrantyInfo_1_Duration                                              string `json:"WarrantyInfo.1.Duration"`
	WarrantyInfo_1_EndDate                                               string `json:"WarrantyInfo.1.EndDate"`
	WarrantyInfo_1_UnitType                                              string `json:"WarrantyInfo.1.UnitType"`
}

//IDRACAttrDell ... IDRAC Attributes from the Redfish API
type IDRACAttrDell struct {
	RedfishSettings struct {
		odataContext   string
		odataType      string
		SettingsObject struct {
			_odata_id string
		} `json:"SettingsObject"`
	}
	odataContext      string
	odataID           string
	odataType         string
	AttributeRegistry string              `json:"AttributeRegistry"`
	Attributes        IDRACAttributesData `json:"Attributes"`
	Description       string              `json:"Description"`
	ID                string              `json:"Id"`
	Name              string              `json:"Name"`
}

type IDRACAttributesData struct {
	ADGroup_1_Domain                                                     string      `json:"ADGroup.1.Domain"`
	ADGroup_1_Name                                                       string      `json:"ADGroup.1.Name"`
	ADGroup_1_Privilege                                                  int         `json:"ADGroup.1.Privilege"`
	ADGroup_2_Domain                                                     string      `json:"ADGroup.2.Domain"`
	ADGroup_2_Name                                                       string      `json:"ADGroup.2.Name"`
	ADGroup_2_Privilege                                                  int         `json:"ADGroup.2.Privilege"`
	ADGroup_3_Domain                                                     string      `json:"ADGroup.3.Domain"`
	ADGroup_3_Name                                                       string      `json:"ADGroup.3.Name"`
	ADGroup_3_Privilege                                                  int         `json:"ADGroup.3.Privilege"`
	ADGroup_4_Domain                                                     string      `json:"ADGroup.4.Domain"`
	ADGroup_4_Name                                                       string      `json:"ADGroup.4.Name"`
	ADGroup_4_Privilege                                                  int         `json:"ADGroup.4.Privilege"`
	ADGroup_5_Domain                                                     string      `json:"ADGroup.5.Domain"`
	ADGroup_5_Name                                                       string      `json:"ADGroup.5.Name"`
	ADGroup_5_Privilege                                                  int         `json:"ADGroup.5.Privilege"`
	ASRConfig_1_Enable                                                   string      `json:"ASRConfig.1.Enable"`
	ActiveDirectory_1_AuthTimeout                                        int         `json:"ActiveDirectory.1.AuthTimeout"`
	ActiveDirectory_1_CertValidationEnable                               string      `json:"ActiveDirectory.1.CertValidationEnable"`
	ActiveDirectory_1_DCLookupByUserDomain                               string      `json:"ActiveDirectory.1.DCLookupByUserDomain"`
	ActiveDirectory_1_DCLookupDomainName                                 string      `json:"ActiveDirectory.1.DCLookupDomainName"`
	ActiveDirectory_1_DCLookupEnable                                     string      `json:"ActiveDirectory.1.DCLookupEnable"`
	ActiveDirectory_1_DomainController1                                  string      `json:"ActiveDirectory.1.DomainController1"`
	ActiveDirectory_1_DomainController2                                  string      `json:"ActiveDirectory.1.DomainController2"`
	ActiveDirectory_1_DomainController3                                  string      `json:"ActiveDirectory.1.DomainController3"`
	ActiveDirectory_1_Enable                                             string      `json:"ActiveDirectory.1.Enable"`
	ActiveDirectory_1_GCLookupEnable                                     string      `json:"ActiveDirectory.1.GCLookupEnable"`
	ActiveDirectory_1_GCRootDomain                                       string      `json:"ActiveDirectory.1.GCRootDomain"`
	ActiveDirectory_1_GlobalCatalog1                                     string      `json:"ActiveDirectory.1.GlobalCatalog1"`
	ActiveDirectory_1_GlobalCatalog2                                     string      `json:"ActiveDirectory.1.GlobalCatalog2"`
	ActiveDirectory_1_GlobalCatalog3                                     string      `json:"ActiveDirectory.1.GlobalCatalog3"`
	ActiveDirectory_1_RacDomain                                          string      `json:"ActiveDirectory.1.RacDomain"`
	ActiveDirectory_1_RacName                                            string      `json:"ActiveDirectory.1.RacName"`
	ActiveDirectory_1_SSOEnable                                          string      `json:"ActiveDirectory.1.SSOEnable"`
	ActiveDirectory_1_Schema                                             string      `json:"ActiveDirectory.1.Schema"`
	AutoOSLockGroup_1_AutoOSLockState                                    string      `json:"AutoOSLockGroup.1.AutoOSLockState"`
	Autodiscovery_1_EnableIPChangeAnnounce                               string      `json:"Autodiscovery.1.EnableIPChangeAnnounce"`
	Autodiscovery_1_EnableIPChangeAnnounceFromDHCP                       string      `json:"Autodiscovery.1.EnableIPChangeAnnounceFromDHCP"`
	Autodiscovery_1_EnableIPChangeAnnounceFromUnicastDNS                 string      `json:"Autodiscovery.1.EnableIPChangeAnnounceFromUnicastDNS"`
	Autodiscovery_1_EnableIPChangeAnnounceFrommDNS                       string      `json:"Autodiscovery.1.EnableIPChangeAnnounceFrommDNS"`
	Autodiscovery_1_UnsolicitedIPChangeAnnounceRate                      string      `json:"Autodiscovery.1.UnsolicitedIPChangeAnnounceRate"`
	CurrentIPv4_1_Address                                                string      `json:"CurrentIPv4.1.Address"`
	CurrentIPv4_1_DHCPEnable                                             string      `json:"CurrentIPv4.1.DHCPEnable"`
	CurrentIPv4_1_DNS1                                                   string      `json:"CurrentIPv4.1.DNS1"`
	CurrentIPv4_1_DNS2                                                   string      `json:"CurrentIPv4.1.DNS2"`
	CurrentIPv4_1_DNSFromDHCP                                            string      `json:"CurrentIPv4.1.DNSFromDHCP"`
	CurrentIPv4_1_DupAddrDetected                                        string      `json:"CurrentIPv4.1.DupAddrDetected"`
	CurrentIPv4_1_Enable                                                 string      `json:"CurrentIPv4.1.Enable"`
	CurrentIPv4_1_Gateway                                                string      `json:"CurrentIPv4.1.Gateway"`
	CurrentIPv4_1_Netmask                                                string      `json:"CurrentIPv4.1.Netmask"`
	CurrentIPv6_1_Address1                                               string      `json:"CurrentIPv6.1.Address1"`
	CurrentIPv6_1_Address10                                              string      `json:"CurrentIPv6.1.Address10"`
	CurrentIPv6_1_Address11                                              string      `json:"CurrentIPv6.1.Address11"`
	CurrentIPv6_1_Address12                                              string      `json:"CurrentIPv6.1.Address12"`
	CurrentIPv6_1_Address13                                              string      `json:"CurrentIPv6.1.Address13"`
	CurrentIPv6_1_Address14                                              string      `json:"CurrentIPv6.1.Address14"`
	CurrentIPv6_1_Address15                                              string      `json:"CurrentIPv6.1.Address15"`
	CurrentIPv6_1_Address2                                               string      `json:"CurrentIPv6.1.Address2"`
	CurrentIPv6_1_Address3                                               string      `json:"CurrentIPv6.1.Address3"`
	CurrentIPv6_1_Address4                                               string      `json:"CurrentIPv6.1.Address4"`
	CurrentIPv6_1_Address5                                               string      `json:"CurrentIPv6.1.Address5"`
	CurrentIPv6_1_Address6                                               string      `json:"CurrentIPv6.1.Address6"`
	CurrentIPv6_1_Address7                                               string      `json:"CurrentIPv6.1.Address7"`
	CurrentIPv6_1_Address8                                               string      `json:"CurrentIPv6.1.Address8"`
	CurrentIPv6_1_Address9                                               string      `json:"CurrentIPv6.1.Address9"`
	CurrentIPv6_1_AutoConfig                                             string      `json:"CurrentIPv6.1.AutoConfig"`
	CurrentIPv6_1_DNS1                                                   string      `json:"CurrentIPv6.1.DNS1"`
	CurrentIPv6_1_DNS2                                                   string      `json:"CurrentIPv6.1.DNS2"`
	CurrentIPv6_1_DNSFromDHCP6                                           string      `json:"CurrentIPv6.1.DNSFromDHCP6"`
	CurrentIPv6_1_DUID                                                   string      `json:"CurrentIPv6.1.DUID"`
	CurrentIPv6_1_Enable                                                 string      `json:"CurrentIPv6.1.Enable"`
	CurrentIPv6_1_Gateway                                                string      `json:"CurrentIPv6.1.Gateway"`
	CurrentIPv6_1_IPV6NumOfExtAddress                                    int         `json:"CurrentIPv6.1.IPV6NumOfExtAddress"`
	CurrentIPv6_1_LinkLocalAddress                                       string      `json:"CurrentIPv6.1.LinkLocalAddress"`
	CurrentIPv6_1_PrefixLength                                           int         `json:"CurrentIPv6.1.PrefixLength"`
	CurrentNIC_1_ActiveNIC                                               string      `json:"CurrentNIC.1.ActiveNIC"`
	CurrentNIC_1_ActiveSharedLOM                                         string      `json:"CurrentNIC.1.ActiveSharedLOM"`
	CurrentNIC_1_AutoDetect                                              string      `json:"CurrentNIC.1.AutoDetect"`
	CurrentNIC_1_Autoneg                                                 string      `json:"CurrentNIC.1.Autoneg"`
	CurrentNIC_1_DNSDomainFromDHCP                                       string      `json:"CurrentNIC.1.DNSDomainFromDHCP"`
	CurrentNIC_1_DNSDomainName                                           string      `json:"CurrentNIC.1.DNSDomainName"`
	CurrentNIC_1_DNSRacName                                              string      `json:"CurrentNIC.1.DNSRacName"`
	CurrentNIC_1_DNSRegister                                             string      `json:"CurrentNIC.1.DNSRegister"`
	CurrentNIC_1_DedicatedNICScanTime                                    int         `json:"CurrentNIC.1.DedicatedNICScanTime"`
	CurrentNIC_1_Duplex                                                  string      `json:"CurrentNIC.1.Duplex"`
	CurrentNIC_1_Enable                                                  string      `json:"CurrentNIC.1.Enable"`
	CurrentNIC_1_Failover                                                string      `json:"CurrentNIC.1.Failover"`
	CurrentNIC_1_LinkStatus                                              string      `json:"CurrentNIC.1.LinkStatus"`
	CurrentNIC_1_MACAddress                                              string      `json:"CurrentNIC.1.MACAddress"`
	CurrentNIC_1_MACAddress2                                             string      `json:"CurrentNIC.1.MACAddress2"`
	CurrentNIC_1_MTU                                                     int         `json:"CurrentNIC.1.MTU"`
	CurrentNIC_1_MgmtIfaceName                                           string      `json:"CurrentNIC.1.MgmtIfaceName"`
	CurrentNIC_1_NumberOfLOM                                             int         `json:"CurrentNIC.1.NumberOfLOM"`
	CurrentNIC_1_Selection                                               string      `json:"CurrentNIC.1.Selection"`
	CurrentNIC_1_SharedNICScanTime                                       int         `json:"CurrentNIC.1.SharedNICScanTime"`
	CurrentNIC_1_Speed                                                   string      `json:"CurrentNIC.1.Speed"`
	CurrentNIC_1_VLanEnable                                              string      `json:"CurrentNIC.1.VLanEnable"`
	CurrentNIC_1_VLanID                                                  int         `json:"CurrentNIC.1.VLanID"`
	CurrentNIC_1_VLanPriority                                            int         `json:"CurrentNIC.1.VLanPriority"`
	CurrentNIC_1_VLanSetting                                             string      `json:"CurrentNIC.1.VLanSetting"`
	DefaultCredentialMitigationConfigGroup_1_DefaultCredentialMitigation string      `json:"DefaultCredentialMitigationConfigGroup.1.DefaultCredentialMitigation"`
	EmailAlert_1_Address                                                 string      `json:"EmailAlert.1.Address"`
	EmailAlert_1_CustomMsg                                               string      `json:"EmailAlert.1.CustomMsg"`
	EmailAlert_1_Enable                                                  string      `json:"EmailAlert.1.Enable"`
	EmailAlert_2_Address                                                 string      `json:"EmailAlert.2.Address"`
	EmailAlert_2_CustomMsg                                               string      `json:"EmailAlert.2.CustomMsg"`
	EmailAlert_2_Enable                                                  string      `json:"EmailAlert.2.Enable"`
	EmailAlert_3_Address                                                 string      `json:"EmailAlert.3.Address"`
	EmailAlert_3_CustomMsg                                               string      `json:"EmailAlert.3.CustomMsg"`
	EmailAlert_3_Enable                                                  string      `json:"EmailAlert.3.Enable"`
	EmailAlert_4_Address                                                 string      `json:"EmailAlert.4.Address"`
	EmailAlert_4_CustomMsg                                               string      `json:"EmailAlert.4.CustomMsg"`
	EmailAlert_4_Enable                                                  string      `json:"EmailAlert.4.Enable"`
	GUI_1_SecurityPolicyMessage                                          string      `json:"GUI.1.SecurityPolicyMessage"`
	GroupManager_1_GroupName                                             string      `json:"GroupManager.1.GroupName"`
	GroupManager_1_GroupUUID                                             string      `json:"GroupManager.1.GroupUUID"`
	GroupManager_1_Status                                                string      `json:"GroupManager.1.Status"`
	IOIDOpt_1_IOIDOptEnable                                              string      `json:"IOIDOpt.1.IOIDOptEnable"`
	IOIDOpt_1_InitiatorPersistencePolicy                                 string      `json:"IOIDOpt.1.InitiatorPersistencePolicy"`
	IOIDOpt_1_StorageTargetPersistencePolicy                             string      `json:"IOIDOpt.1.StorageTargetPersistencePolicy"`
	IOIDOpt_1_VirtualAddressPersistencePolicyAuxPwrd                     string      `json:"IOIDOpt.1.VirtualAddressPersistencePolicyAuxPwrd"`
	IOIDOpt_1_VirtualAddressPersistencePolicyNonAuxPwrd                  string      `json:"IOIDOpt.1.VirtualAddressPersistencePolicyNonAuxPwrd"`
	IPBlocking_1_BlockEnable                                             string      `json:"IPBlocking.1.BlockEnable"`
	IPBlocking_1_FailCount                                               int         `json:"IPBlocking.1.FailCount"`
	IPBlocking_1_FailWindow                                              int         `json:"IPBlocking.1.FailWindow"`
	IPBlocking_1_PenaltyTime                                             int         `json:"IPBlocking.1.PenaltyTime"`
	IPBlocking_1_RangeAddr                                               string      `json:"IPBlocking.1.RangeAddr"`
	IPBlocking_1_RangeEnable                                             string      `json:"IPBlocking.1.RangeEnable"`
	IPBlocking_1_RangeMask                                               string      `json:"IPBlocking.1.RangeMask"`
	IPMILan_1_AlertEnable                                                string      `json:"IPMILan.1.AlertEnable"`
	IPMILan_1_CommunityName                                              string      `json:"IPMILan.1.CommunityName"`
	IPMILan_1_Enable                                                     string      `json:"IPMILan.1.Enable"`
	IPMILan_1_EncryptionKey                                              string      `json:"IPMILan.1.EncryptionKey"`
	IPMILan_1_PrivLimit                                                  string      `json:"IPMILan.1.PrivLimit"`
	IPMISOL_1_AccumulateInterval                                         int         `json:"IPMISOL.1.AccumulateInterval"`
	IPMISOL_1_BaudRate                                                   string      `json:"IPMISOL.1.BaudRate"`
	IPMISOL_1_Enable                                                     string      `json:"IPMISOL.1.Enable"`
	IPMISOL_1_MinPrivilege                                               string      `json:"IPMISOL.1.MinPrivilege"`
	IPMISOL_1_SendThreshold                                              int         `json:"IPMISOL.1.SendThreshold"`
	IPMISerial_1_BaudRate                                                string      `json:"IPMISerial.1.BaudRate"`
	IPMISerial_1_ChanPrivLimit                                           string      `json:"IPMISerial.1.ChanPrivLimit"`
	IPMISerial_1_ConnectionMode                                          string      `json:"IPMISerial.1.ConnectionMode"`
	IPMISerial_1_DeleteControl                                           string      `json:"IPMISerial.1.DeleteControl"`
	IPMISerial_1_EchoControl                                             string      `json:"IPMISerial.1.EchoControl"`
	IPMISerial_1_FlowControl                                             string      `json:"IPMISerial.1.FlowControl"`
	IPMISerial_1_HandshakeControl                                        string      `json:"IPMISerial.1.HandshakeControl"`
	IPMISerial_1_InputNewLineSeq                                         string      `json:"IPMISerial.1.InputNewLineSeq"`
	IPMISerial_1_LineEdit                                                string      `json:"IPMISerial.1.LineEdit"`
	IPMISerial_1_NewLineSeq                                              string      `json:"IPMISerial.1.NewLineSeq"`
	IPv4_1_Address                                                       string      `json:"IPv4.1.Address"`
	IPv4_1_DHCPEnable                                                    string      `json:"IPv4.1.DHCPEnable"`
	IPv4_1_DNS1                                                          string      `json:"IPv4.1.DNS1"`
	IPv4_1_DNS2                                                          string      `json:"IPv4.1.DNS2"`
	IPv4_1_DNSFromDHCP                                                   string      `json:"IPv4.1.DNSFromDHCP"`
	IPv4_1_Enable                                                        string      `json:"IPv4.1.Enable"`
	IPv4_1_Gateway                                                       string      `json:"IPv4.1.Gateway"`
	IPv4_1_Netmask                                                       string      `json:"IPv4.1.Netmask"`
	IPv4Static_1_Address                                                 string      `json:"IPv4Static.1.Address"`
	IPv4Static_1_DNS1                                                    string      `json:"IPv4Static.1.DNS1"`
	IPv4Static_1_DNS2                                                    string      `json:"IPv4Static.1.DNS2"`
	IPv4Static_1_DNSFromDHCP                                             string      `json:"IPv4Static.1.DNSFromDHCP"`
	IPv4Static_1_Gateway                                                 string      `json:"IPv4Static.1.Gateway"`
	IPv4Static_1_Netmask                                                 string      `json:"IPv4Static.1.Netmask"`
	IPv6_1_Address1                                                      string      `json:"IPv6.1.Address1"`
	IPv6_1_Address10                                                     string      `json:"IPv6.1.Address10"`
	IPv6_1_Address11                                                     string      `json:"IPv6.1.Address11"`
	IPv6_1_Address12                                                     string      `json:"IPv6.1.Address12"`
	IPv6_1_Address13                                                     string      `json:"IPv6.1.Address13"`
	IPv6_1_Address14                                                     string      `json:"IPv6.1.Address14"`
	IPv6_1_Address15                                                     string      `json:"IPv6.1.Address15"`
	IPv6_1_Address2                                                      string      `json:"IPv6.1.Address2"`
	IPv6_1_Address3                                                      string      `json:"IPv6.1.Address3"`
	IPv6_1_Address4                                                      string      `json:"IPv6.1.Address4"`
	IPv6_1_Address5                                                      string      `json:"IPv6.1.Address5"`
	IPv6_1_Address6                                                      string      `json:"IPv6.1.Address6"`
	IPv6_1_Address7                                                      string      `json:"IPv6.1.Address7"`
	IPv6_1_Address8                                                      string      `json:"IPv6.1.Address8"`
	IPv6_1_Address9                                                      string      `json:"IPv6.1.Address9"`
	IPv6_1_AddressState                                                  string      `json:"IPv6.1.AddressState"`
	IPv6_1_AutoConfig                                                    string      `json:"IPv6.1.AutoConfig"`
	IPv6_1_DNS1                                                          string      `json:"IPv6.1.DNS1"`
	IPv6_1_DNS2                                                          string      `json:"IPv6.1.DNS2"`
	IPv6_1_DNSFromDHCP6                                                  string      `json:"IPv6.1.DNSFromDHCP6"`
	IPv6_1_DUID                                                          string      `json:"IPv6.1.DUID"`
	IPv6_1_Enable                                                        string      `json:"IPv6.1.Enable"`
	IPv6_1_Gateway                                                       string      `json:"IPv6.1.Gateway"`
	IPv6_1_LinkLocalAddress                                              string      `json:"IPv6.1.LinkLocalAddress"`
	IPv6_1_PrefixLength                                                  int         `json:"IPv6.1.PrefixLength"`
	IPv6Static_1_Address1                                                string      `json:"IPv6Static.1.Address1"`
	IPv6Static_1_DNS1                                                    string      `json:"IPv6Static.1.DNS1"`
	IPv6Static_1_DNS2                                                    string      `json:"IPv6Static.1.DNS2"`
	IPv6Static_1_DNSFromDHCP6                                            string      `json:"IPv6Static.1.DNSFromDHCP6"`
	IPv6Static_1_Gateway                                                 string      `json:"IPv6Static.1.Gateway"`
	IPv6Static_1_PrefixLength                                            int         `json:"IPv6Static.1.PrefixLength"`
	IPv6URL_1_URL                                                        string      `json:"IPv6URL.1.URL"`
	Info_1_Build                                                         string      `json:"Info.1.Build"`
	Info_1_CPLDVersion                                                   string      `json:"Info.1.CPLDVersion"`
	Info_1_Description                                                   string      `json:"Info.1.Description"`
	Info_1_HWRev                                                         string      `json:"Info.1.HWRev"`
	Info_1_IPMIVersion                                                   string      `json:"Info.1.IPMIVersion"`
	Info_1_Name                                                          string      `json:"Info.1.Name"`
	Info_1_Product                                                       string      `json:"Info.1.Product"`
	Info_1_RollbackBuild                                                 string      `json:"Info.1.RollbackBuild"`
	Info_1_RollbackVersion                                               string      `json:"Info.1.RollbackVersion"`
	Info_1_ServerGen                                                     string      `json:"Info.1.ServerGen"`
	Info_1_Type                                                          string      `json:"Info.1.Type"`
	Info_1_Version                                                       string      `json:"Info.1.Version"`
	IntegratedDatacenter_1_DiscoveryEnable                               string      `json:"IntegratedDatacenter.1.DiscoveryEnable"`
	LDAP_1_BaseDN                                                        string      `json:"LDAP.1.BaseDN"`
	LDAP_1_BindDN                                                        string      `json:"LDAP.1.BindDN"`
	LDAP_1_BindPassword                                                  interface{} `json:"LDAP.1.BindPassword"`
	LDAP_1_CertValidationEnable                                          string      `json:"LDAP.1.CertValidationEnable"`
	LDAP_1_Enable                                                        string      `json:"LDAP.1.Enable"`
	LDAP_1_GroupAttribute                                                string      `json:"LDAP.1.GroupAttribute"`
	LDAP_1_GroupAttributeIsDN                                            string      `json:"LDAP.1.GroupAttributeIsDN"`
	LDAP_1_Port                                                          int         `json:"LDAP.1.Port"`
	LDAP_1_SearchFilter                                                  string      `json:"LDAP.1.SearchFilter"`
	LDAP_1_Server                                                        string      `json:"LDAP.1.Server"`
	LDAP_1_UserAttribute                                                 string      `json:"LDAP.1.UserAttribute"`
	LDAPRoleGroup_1_DN                                                   string      `json:"LDAPRoleGroup.1.DN"`
	LDAPRoleGroup_1_Privilege                                            int         `json:"LDAPRoleGroup.1.Privilege"`
	LDAPRoleGroup_2_DN                                                   string      `json:"LDAPRoleGroup.2.DN"`
	LDAPRoleGroup_2_Privilege                                            int         `json:"LDAPRoleGroup.2.Privilege"`
	LDAPRoleGroup_3_DN                                                   string      `json:"LDAPRoleGroup.3.DN"`
	LDAPRoleGroup_3_Privilege                                            int         `json:"LDAPRoleGroup.3.Privilege"`
	LDAPRoleGroup_4_DN                                                   string      `json:"LDAPRoleGroup.4.DN"`
	LDAPRoleGroup_4_Privilege                                            int         `json:"LDAPRoleGroup.4.Privilege"`
	LDAPRoleGroup_5_DN                                                   string      `json:"LDAPRoleGroup.5.DN"`
	LDAPRoleGroup_5_Privilege                                            int         `json:"LDAPRoleGroup.5.Privilege"`
	LocalSecurity_1_LocalConfig                                          string      `json:"LocalSecurity.1.LocalConfig"`
	LocalSecurity_1_PrebootConfig                                        string      `json:"LocalSecurity.1.PrebootConfig"`
	Lockdown_1_SystemLockdown                                            string      `json:"Lockdown.1.SystemLockdown"`
	Logging_1_SELBufferType                                              string      `json:"Logging.1.SELBufferType"`
	Logging_1_SELOEMEventFilterEnable                                    string      `json:"Logging.1.SELOEMEventFilterEnable"`
	NIC_1_AutoConfig                                                     string      `json:"NIC.1.AutoConfig"`
	NIC_1_AutoConfigIPV6                                                 string      `json:"NIC.1.AutoConfigIPV6"`
	NIC_1_AutoDetect                                                     string      `json:"NIC.1.AutoDetect"`
	NIC_1_Autoneg                                                        string      `json:"NIC.1.Autoneg"`
	NIC_1_DNSDomainFromDHCP                                              string      `json:"NIC.1.DNSDomainFromDHCP"`
	NIC_1_DNSDomainName                                                  string      `json:"NIC.1.DNSDomainName"`
	NIC_1_DNSDomainNameFromDHCP                                          string      `json:"NIC.1.DNSDomainNameFromDHCP"`
	NIC_1_DNSRacName                                                     string      `json:"NIC.1.DNSRacName"`
	NIC_1_DNSRegister                                                    string      `json:"NIC.1.DNSRegister"`
	NIC_1_DNSRegisterInterval                                            int         `json:"NIC.1.DNSRegisterInterval"`
	NIC_1_DedicatedNICScanTime                                           int         `json:"NIC.1.DedicatedNICScanTime"`
	NIC_1_DiscoveryLLDP                                                  string      `json:"NIC.1.DiscoveryLLDP"`
	NIC_1_Duplex                                                         string      `json:"NIC.1.Duplex"`
	NIC_1_Enable                                                         string      `json:"NIC.1.Enable"`
	NIC_1_Failover                                                       string      `json:"NIC.1.Failover"`
	NIC_1_MACAddress                                                     string      `json:"NIC.1.MACAddress"`
	NIC_1_MTU                                                            int         `json:"NIC.1.MTU"`
	NIC_1_Selection                                                      string      `json:"NIC.1.Selection"`
	NIC_1_SharedNICScanTime                                              int         `json:"NIC.1.SharedNICScanTime"`
	NIC_1_Speed                                                          string      `json:"NIC.1.Speed"`
	NIC_1_SwitchConnection                                               string      `json:"NIC.1.SwitchConnection"`
	NIC_1_SwitchPortConnection                                           string      `json:"NIC.1.SwitchPortConnection"`
	NIC_1_VLanEnable                                                     string      `json:"NIC.1.VLanEnable"`
	NIC_1_VLanID                                                         int         `json:"NIC.1.VLanID"`
	NIC_1_VLanPort                                                       string      `json:"NIC.1.VLanPort"`
	NIC_1_VLanPriority                                                   int         `json:"NIC.1.VLanPriority"`
	NICStatic_1_DNSDomainFromDHCP                                        string      `json:"NICStatic.1.DNSDomainFromDHCP"`
	NICStatic_1_DNSDomainName                                            string      `json:"NICStatic.1.DNSDomainName"`
	NTPConfigGroup_1_NTP1                                                string      `json:"NTPConfigGroup.1.NTP1"`
	NTPConfigGroup_1_NTP2                                                string      `json:"NTPConfigGroup.1.NTP2"`
	NTPConfigGroup_1_NTP3                                                string      `json:"NTPConfigGroup.1.NTP3"`
	NTPConfigGroup_1_NTPEnable                                           string      `json:"NTPConfigGroup.1.NTPEnable"`
	NTPConfigGroup_1_NTPMaxDist                                          int         `json:"NTPConfigGroup.1.NTPMaxDist"`
	OS_BMC_1_AdminState                                                  string      `json:"OS-BMC.1.AdminState"`
	OS_BMC_1_OsIPAddress                                                 string      `json:"OS-BMC.1.OsIpAddress"`
	OS_BMC_1_PTCapability                                                string      `json:"OS-BMC.1.PTCapability"`
	OS_BMC_1_PTMode                                                      string      `json:"OS-BMC.1.PTMode"`
	OS_BMC_1_UsbNicIPAddress                                             string      `json:"OS-BMC.1.UsbNicIpAddress"`
	OS_BMC_1_UsbNicIPV6Address                                           string      `json:"OS-BMC.1.UsbNicIpV6Address"`
	PlatformCapability_1_FrontPortUSBConfiguration                       string      `json:"PlatformCapability.1.FrontPortUSBConfiguration"`
	RFS_1_AttachMode                                                     string      `json:"RFS.1.AttachMode"`
	RFS_1_Enable                                                         string      `json:"RFS.1.Enable"`
	RFS_1_IgnoreCertWarning                                              string      `json:"RFS.1.IgnoreCertWarning"`
	RFS_1_Image                                                          string      `json:"RFS.1.Image"`
	RFS_1_MediaAttachState                                               string      `json:"RFS.1.MediaAttachState"`
	RFS_1_Password                                                       interface{} `json:"RFS.1.Password"`
	RFS_1_Status                                                         string      `json:"RFS.1.Status"`
	RFS_1_User                                                           string      `json:"RFS.1.User"`
	RFS_1_WriteProtected                                                 string      `json:"RFS.1.WriteProtected"`
	Racadm_1_Enable                                                      string      `json:"Racadm.1.Enable"`
	Racadm_1_MaxSessions                                                 int         `json:"Racadm.1.MaxSessions"`
	Racadm_1_Timeout                                                     int         `json:"Racadm.1.Timeout"`
	Redfish_1_Enable                                                     string      `json:"Redfish.1.Enable"`
	RedfishEventing_1_DeliveryRetryAttempts                              int         `json:"RedfishEventing.1.DeliveryRetryAttempts"`
	RedfishEventing_1_DeliveryRetryIntervalInSeconds                     int         `json:"RedfishEventing.1.DeliveryRetryIntervalInSeconds"`
	RedfishEventing_1_IgnoreCertificateErrors                            string      `json:"RedfishEventing.1.IgnoreCertificateErrors"`
	RemoteHosts_1_SMTPAuthentication                                     string      `json:"RemoteHosts.1.SMTPAuthentication"`
	RemoteHosts_1_SMTPPassword                                           interface{} `json:"RemoteHosts.1.SMTPPassword"`
	RemoteHosts_1_SMTPPort                                               int         `json:"RemoteHosts.1.SMTPPort"`
	RemoteHosts_1_SMTPServerIPAddress                                    string      `json:"RemoteHosts.1.SMTPServerIPAddress"`
	RemoteHosts_1_SMTPUserName                                           string      `json:"RemoteHosts.1.SMTPUserName"`
	RemoteHosts_1_SenderEmail                                            string      `json:"RemoteHosts.1.SenderEmail"`
	SNMP_1_AgentCommunity                                                string      `json:"SNMP.1.AgentCommunity"`
	SNMP_1_AgentEnable                                                   string      `json:"SNMP.1.AgentEnable"`
	SNMP_1_AlertPort                                                     int         `json:"SNMP.1.AlertPort"`
	SNMP_1_DiscoveryPort                                                 int         `json:"SNMP.1.DiscoveryPort"`
	SNMP_1_EngineID                                                      string      `json:"SNMP.1.EngineID"`
	SNMP_1_SNMPProtocol                                                  string      `json:"SNMP.1.SNMPProtocol"`
	SNMP_1_TrapFormat                                                    string      `json:"SNMP.1.TrapFormat"`
	SNMPAlert_1_Destination                                              string      `json:"SNMPAlert.1.Destination"`
	SNMPAlert_1_SNMPv3UserID                                             int         `json:"SNMPAlert.1.SNMPv3UserID"`
	SNMPAlert_1_SNMPv3Username                                           string      `json:"SNMPAlert.1.SNMPv3Username"`
	SNMPAlert_1_State                                                    string      `json:"SNMPAlert.1.State"`
	SNMPAlert_2_Destination                                              string      `json:"SNMPAlert.2.Destination"`
	SNMPAlert_2_SNMPv3UserID                                             int         `json:"SNMPAlert.2.SNMPv3UserID"`
	SNMPAlert_2_SNMPv3Username                                           string      `json:"SNMPAlert.2.SNMPv3Username"`
	SNMPAlert_2_State                                                    string      `json:"SNMPAlert.2.State"`
	SNMPAlert_3_Destination                                              string      `json:"SNMPAlert.3.Destination"`
	SNMPAlert_3_SNMPv3UserID                                             int         `json:"SNMPAlert.3.SNMPv3UserID"`
	SNMPAlert_3_SNMPv3Username                                           string      `json:"SNMPAlert.3.SNMPv3Username"`
	SNMPAlert_3_State                                                    string      `json:"SNMPAlert.3.State"`
	SNMPAlert_4_Destination                                              string      `json:"SNMPAlert.4.Destination"`
	SNMPAlert_4_SNMPv3UserID                                             int         `json:"SNMPAlert.4.SNMPv3UserID"`
	SNMPAlert_4_SNMPv3Username                                           string      `json:"SNMPAlert.4.SNMPv3Username"`
	SNMPAlert_4_State                                                    string      `json:"SNMPAlert.4.State"`
	SNMPAlert_5_Destination                                              string      `json:"SNMPAlert.5.Destination"`
	SNMPAlert_5_SNMPv3UserID                                             int         `json:"SNMPAlert.5.SNMPv3UserID"`
	SNMPAlert_5_SNMPv3Username                                           string      `json:"SNMPAlert.5.SNMPv3Username"`
	SNMPAlert_5_State                                                    string      `json:"SNMPAlert.5.State"`
	SNMPAlert_6_Destination                                              string      `json:"SNMPAlert.6.Destination"`
	SNMPAlert_6_SNMPv3UserID                                             int         `json:"SNMPAlert.6.SNMPv3UserID"`
	SNMPAlert_6_SNMPv3Username                                           string      `json:"SNMPAlert.6.SNMPv3Username"`
	SNMPAlert_6_State                                                    string      `json:"SNMPAlert.6.State"`
	SNMPAlert_7_Destination                                              string      `json:"SNMPAlert.7.Destination"`
	SNMPAlert_7_SNMPv3UserID                                             int         `json:"SNMPAlert.7.SNMPv3UserID"`
	SNMPAlert_7_SNMPv3Username                                           string      `json:"SNMPAlert.7.SNMPv3Username"`
	SNMPAlert_7_State                                                    string      `json:"SNMPAlert.7.State"`
	SNMPAlert_8_Destination                                              string      `json:"SNMPAlert.8.Destination"`
	SNMPAlert_8_SNMPv3UserID                                             int         `json:"SNMPAlert.8.SNMPv3UserID"`
	SNMPAlert_8_SNMPv3Username                                           string      `json:"SNMPAlert.8.SNMPv3Username"`
	SNMPAlert_8_State                                                    string      `json:"SNMPAlert.8.State"`
	SNMPTrapIPv4_1_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.1.DestIPv4Addr"`
	SNMPTrapIPv4_1_DestinationNum                                        int         `json:"SNMPTrapIPv4.1.DestinationNum"`
	SNMPTrapIPv4_1_State                                                 string      `json:"SNMPTrapIPv4.1.State"`
	SNMPTrapIPv4_2_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.2.DestIPv4Addr"`
	SNMPTrapIPv4_2_DestinationNum                                        int         `json:"SNMPTrapIPv4.2.DestinationNum"`
	SNMPTrapIPv4_2_State                                                 string      `json:"SNMPTrapIPv4.2.State"`
	SNMPTrapIPv4_3_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.3.DestIPv4Addr"`
	SNMPTrapIPv4_3_DestinationNum                                        int         `json:"SNMPTrapIPv4.3.DestinationNum"`
	SNMPTrapIPv4_3_State                                                 string      `json:"SNMPTrapIPv4.3.State"`
	SNMPTrapIPv4_4_DestIPv4Addr                                          string      `json:"SNMPTrapIPv4.4.DestIPv4Addr"`
	SNMPTrapIPv4_4_DestinationNum                                        int         `json:"SNMPTrapIPv4.4.DestinationNum"`
	SNMPTrapIPv4_4_State                                                 string      `json:"SNMPTrapIPv4.4.State"`
	SNMPTrapIPv6_1_DestIPv6Addr                                          string      `json:"SNMPTrapIPv6.1.DestIPv6Addr"`
	SNMPTrapIPv6_1_DestinationNum                                        int         `json:"SNMPTrapIPv6.1.DestinationNum"`
	SNMPTrapIPv6_1_State                                                 string      `json:"SNMPTrapIPv6.1.State"`
	SNMPTrapIPv6_2_DestIPv6Addr                                          string      `json:"SNMPTrapIPv6.2.DestIPv6Addr"`
	SNMPTrapIPv6_2_DestinationNum                                        int         `json:"SNMPTrapIPv6.2.DestinationNum"`
	SNMPTrapIPv6_2_State                                                 string      `json:"SNMPTrapIPv6.2.State"`
	SNMPTrapIPv6_3_DestIPv6Addr                                          string      `json:"SNMPTrapIPv6.3.DestIPv6Addr"`
	SNMPTrapIPv6_3_DestinationNum                                        int         `json:"SNMPTrapIPv6.3.DestinationNum"`
	SNMPTrapIPv6_3_State                                                 string      `json:"SNMPTrapIPv6.3.State"`
	SSH_1_Enable                                                         string      `json:"SSH.1.Enable"`
	SSH_1_MaxSessions                                                    int         `json:"SSH.1.MaxSessions"`
	SSH_1_Port                                                           int         `json:"SSH.1.Port"`
	SSH_1_Timeout                                                        int         `json:"SSH.1.Timeout"`
	Security_1_CsrCommonName                                             string      `json:"Security.1.CsrCommonName"`
	Security_1_CsrCountryCode                                            string      `json:"Security.1.CsrCountryCode"`
	Security_1_CsrEmailAddr                                              string      `json:"Security.1.CsrEmailAddr"`
	Security_1_CsrKeySize                                                string      `json:"Security.1.CsrKeySize"`
	Security_1_CsrLocalityName                                           string      `json:"Security.1.CsrLocalityName"`
	Security_1_CsrOrganizationName                                       string      `json:"Security.1.CsrOrganizationName"`
	Security_1_CsrOrganizationUnit                                       string      `json:"Security.1.CsrOrganizationUnit"`
	Security_1_CsrStateName                                              string      `json:"Security.1.CsrStateName"`
	Security_1_CsrSubjectAltName                                         string      `json:"Security.1.CsrSubjectAltName"`
	Security_1_FIPSMode                                                  string      `json:"Security.1.FIPSMode"`
	Serial_1_BaudRate                                                    string      `json:"Serial.1.BaudRate"`
	Serial_1_Command                                                     string      `json:"Serial.1.Command"`
	Serial_1_Enable                                                      string      `json:"Serial.1.Enable"`
	Serial_1_HistorySize                                                 int         `json:"Serial.1.HistorySize"`
	Serial_1_IdleTimeout                                                 int         `json:"Serial.1.IdleTimeout"`
	Serial_1_NoAuth                                                      string      `json:"Serial.1.NoAuth"`
	SerialRedirection_1_Enable                                           string      `json:"SerialRedirection.1.Enable"`
	SerialRedirection_1_QuitKey                                          string      `json:"SerialRedirection.1.QuitKey"`
	ServerBoot_1_BootOnce                                                string      `json:"ServerBoot.1.BootOnce"`
	ServerBoot_1_FirstBootDevice                                         string      `json:"ServerBoot.1.FirstBootDevice"`
	ServiceModule_1_HostSNMPAlert                                        string      `json:"ServiceModule.1.HostSNMPAlert"`
	ServiceModule_1_HostSNMPGet                                          string      `json:"ServiceModule.1.HostSNMPGet"`
	ServiceModule_1_LCLReplication                                       string      `json:"ServiceModule.1.LCLReplication"`
	ServiceModule_1_OMSAPresence                                         string      `json:"ServiceModule.1.OMSAPresence"`
	ServiceModule_1_OSInfo                                               string      `json:"ServiceModule.1.OSInfo"`
	ServiceModule_1_ServiceModuleEnable                                  string      `json:"ServiceModule.1.ServiceModuleEnable"`
	ServiceModule_1_ServiceModuleState                                   string      `json:"ServiceModule.1.ServiceModuleState"`
	ServiceModule_1_ServiceModuleVersion                                 string      `json:"ServiceModule.1.ServiceModuleVersion"`
	ServiceModule_1_WMIInfo                                              string      `json:"ServiceModule.1.WMIInfo"`
	ServiceModule_1_WatchdogRecoveryAction                               string      `json:"ServiceModule.1.WatchdogRecoveryAction"`
	ServiceModule_1_WatchdogResetTime                                    int         `json:"ServiceModule.1.WatchdogResetTime"`
	ServiceModule_1_WatchdogState                                        string      `json:"ServiceModule.1.WatchdogState"`
	ServiceModule_1_iDRACHardReset                                       string      `json:"ServiceModule.1.iDRACHardReset"`
	ServiceModule_1_iDRACSSOLauncher                                     string      `json:"ServiceModule.1.iDRACSSOLauncher"`
	SmartCard_1_SmartCardCRLEnable                                       string      `json:"SmartCard.1.SmartCardCRLEnable"`
	SmartCard_1_SmartCardLogonEnable                                     string      `json:"SmartCard.1.SmartCardLogonEnable"`
	SupportAssist_1_DefaultIPAddress                                     string      `json:"SupportAssist.1.DefaultIPAddress"`
	SupportAssist_1_DefaultPassword                                      interface{} `json:"SupportAssist.1.DefaultPassword"`
	SupportAssist_1_DefaultProtocol                                      string      `json:"SupportAssist.1.DefaultProtocol"`
	SupportAssist_1_DefaultProtocolPort                                  int         `json:"SupportAssist.1.DefaultProtocolPort"`
	SupportAssist_1_DefaultShareName                                     string      `json:"SupportAssist.1.DefaultShareName"`
	SupportAssist_1_DefaultUserName                                      string      `json:"SupportAssist.1.DefaultUserName"`
	SupportAssist_1_DefaultWorkgroupName                                 string      `json:"SupportAssist.1.DefaultWorkgroupName"`
	SupportAssist_1_EmailOptIn                                           string      `json:"SupportAssist.1.EmailOptIn"`
	SupportAssist_1_EventBasedAutoCollection                             string      `json:"SupportAssist.1.EventBasedAutoCollection"`
	SupportAssist_1_FilterAutoCollections                                string      `json:"SupportAssist.1.FilterAutoCollections"`
	SupportAssist_1_HostOSProxyAddress                                   string      `json:"SupportAssist.1.HostOSProxyAddress"`
	SupportAssist_1_HostOSProxyConfigured                                string      `json:"SupportAssist.1.HostOSProxyConfigured"`
	SupportAssist_1_HostOSProxyPassword                                  interface{} `json:"SupportAssist.1.HostOSProxyPassword"`
	SupportAssist_1_HostOSProxyPort                                      int         `json:"SupportAssist.1.HostOSProxyPort"`
	SupportAssist_1_HostOSProxyUserName                                  string      `json:"SupportAssist.1.HostOSProxyUserName"`
	SupportAssist_1_NativeOSLogsCollectionSupported                      string      `json:"SupportAssist.1.NativeOSLogsCollectionSupported"`
	SupportAssist_1_PreferredLanguage                                    string      `json:"SupportAssist.1.PreferredLanguage"`
	SupportAssist_1_ProSupportPlusRecommendationsReport                  string      `json:"SupportAssist.1.ProSupportPlusRecommendationsReport"`
	SupportAssist_1_RegistrationID                                       string      `json:"SupportAssist.1.RegistrationID"`
	SupportAssist_1_RequestTechnicianForPartsDispatch                    string      `json:"SupportAssist.1.RequestTechnicianForPartsDispatch"`
	SupportAssist_1_SupportAssistEnableState                             string      `json:"SupportAssist.1.SupportAssistEnableState"`
	SupportAssist_1_iDRACFirstPowerUpDateTime                            string      `json:"SupportAssist.1.iDRACFirstPowerUpDateTime"`
	SwitchConnectionView_1_Enable                                        string      `json:"SwitchConnectionView.1.Enable"`
	SysInfo_1_LocalConsoleLockOut                                        int         `json:"SysInfo.1.LocalConsoleLockOut"`
	SysInfo_1_SystemRev                                                  int         `json:"SysInfo.1.SystemRev"`
	SysLog_1_Port                                                        int         `json:"SysLog.1.Port"`
	SysLog_1_PowerLogEnable                                              string      `json:"SysLog.1.PowerLogEnable"`
	SysLog_1_PowerLogInterval                                            int         `json:"SysLog.1.PowerLogInterval"`
	SysLog_1_Server1                                                     string      `json:"SysLog.1.Server1"`
	SysLog_1_Server2                                                     string      `json:"SysLog.1.Server2"`
	SysLog_1_Server3                                                     string      `json:"SysLog.1.Server3"`
	SysLog_1_SysLogEnable                                                string      `json:"SysLog.1.SysLogEnable"`
	Telnet_1_Enable                                                      string      `json:"Telnet.1.Enable"`
	Telnet_1_MaxSessions                                                 int         `json:"Telnet.1.MaxSessions"`
	Telnet_1_Port                                                        int         `json:"Telnet.1.Port"`
	Telnet_1_Timeout                                                     int         `json:"Telnet.1.Timeout"`
	Time_1_DayLightOffset                                                int         `json:"Time.1.DayLightOffset"`
	Time_1_TimeZoneOffset                                                int         `json:"Time.1.TimeZoneOffset"`
	Time_1_Timezone                                                      string      `json:"Time.1.Timezone"`
	USB_1_ConfigurationXML                                               string      `json:"USB.1.ConfigurationXML"`
	USB_1_ManagementPortMode                                             string      `json:"USB.1.ManagementPortMode"`
	USB_1_PortStatus                                                     string      `json:"USB.1.PortStatus"`
	USB_1_ZipPassword                                                    interface{} `json:"USB.1.ZipPassword"`
	Update_1_FwUpdateIPAddr                                              string      `json:"Update.1.FwUpdateIPAddr"`
	Update_1_FwUpdatePath                                                string      `json:"Update.1.FwUpdatePath"`
	Update_1_FwUpdateTFTPEnable                                          string      `json:"Update.1.FwUpdateTFTPEnable"`
	UserDomain_1_Name                                                    string      `json:"UserDomain.1.Name"`
	UserDomain_10_Name                                                   string      `json:"UserDomain.10.Name"`
	UserDomain_11_Name                                                   string      `json:"UserDomain.11.Name"`
	UserDomain_12_Name                                                   string      `json:"UserDomain.12.Name"`
	UserDomain_13_Name                                                   string      `json:"UserDomain.13.Name"`
	UserDomain_14_Name                                                   string      `json:"UserDomain.14.Name"`
	UserDomain_15_Name                                                   string      `json:"UserDomain.15.Name"`
	UserDomain_16_Name                                                   string      `json:"UserDomain.16.Name"`
	UserDomain_17_Name                                                   string      `json:"UserDomain.17.Name"`
	UserDomain_18_Name                                                   string      `json:"UserDomain.18.Name"`
	UserDomain_19_Name                                                   string      `json:"UserDomain.19.Name"`
	UserDomain_2_Name                                                    string      `json:"UserDomain.2.Name"`
	UserDomain_20_Name                                                   string      `json:"UserDomain.20.Name"`
	UserDomain_21_Name                                                   string      `json:"UserDomain.21.Name"`
	UserDomain_22_Name                                                   string      `json:"UserDomain.22.Name"`
	UserDomain_23_Name                                                   string      `json:"UserDomain.23.Name"`
	UserDomain_24_Name                                                   string      `json:"UserDomain.24.Name"`
	UserDomain_25_Name                                                   string      `json:"UserDomain.25.Name"`
	UserDomain_26_Name                                                   string      `json:"UserDomain.26.Name"`
	UserDomain_27_Name                                                   string      `json:"UserDomain.27.Name"`
	UserDomain_28_Name                                                   string      `json:"UserDomain.28.Name"`
	UserDomain_29_Name                                                   string      `json:"UserDomain.29.Name"`
	UserDomain_3_Name                                                    string      `json:"UserDomain.3.Name"`
	UserDomain_30_Name                                                   string      `json:"UserDomain.30.Name"`
	UserDomain_31_Name                                                   string      `json:"UserDomain.31.Name"`
	UserDomain_32_Name                                                   string      `json:"UserDomain.32.Name"`
	UserDomain_33_Name                                                   string      `json:"UserDomain.33.Name"`
	UserDomain_34_Name                                                   string      `json:"UserDomain.34.Name"`
	UserDomain_35_Name                                                   string      `json:"UserDomain.35.Name"`
	UserDomain_36_Name                                                   string      `json:"UserDomain.36.Name"`
	UserDomain_37_Name                                                   string      `json:"UserDomain.37.Name"`
	UserDomain_38_Name                                                   string      `json:"UserDomain.38.Name"`
	UserDomain_39_Name                                                   string      `json:"UserDomain.39.Name"`
	UserDomain_4_Name                                                    string      `json:"UserDomain.4.Name"`
	UserDomain_40_Name                                                   string      `json:"UserDomain.40.Name"`
	UserDomain_5_Name                                                    string      `json:"UserDomain.5.Name"`
	UserDomain_6_Name                                                    string      `json:"UserDomain.6.Name"`
	UserDomain_7_Name                                                    string      `json:"UserDomain.7.Name"`
	UserDomain_8_Name                                                    string      `json:"UserDomain.8.Name"`
	UserDomain_9_Name                                                    string      `json:"UserDomain.9.Name"`
	Users_1_AuthenticationProtocol                                       string      `json:"Users.1.AuthenticationProtocol"`
	Users_1_Enable                                                       string      `json:"Users.1.Enable"`
	Users_1_IPMIKey                                                      string      `json:"Users.1.IPMIKey"`
	Users_1_IpmiLanPrivilege                                             string      `json:"Users.1.IpmiLanPrivilege"`
	Users_1_IpmiSerialPrivilege                                          string      `json:"Users.1.IpmiSerialPrivilege"`
	Users_1_MD5v3Key                                                     string      `json:"Users.1.MD5v3Key"`
	Users_1_Password                                                     interface{} `json:"Users.1.Password"`
	Users_1_PrivacyProtocol                                              string      `json:"Users.1.PrivacyProtocol"`
	Users_1_Privilege                                                    int         `json:"Users.1.Privilege"`
	Users_1_ProtocolEnable                                               string      `json:"Users.1.ProtocolEnable"`
	Users_1_SHA1v3Key                                                    string      `json:"Users.1.SHA1v3Key"`
	Users_1_SHA256Password                                               string      `json:"Users.1.SHA256Password"`
	Users_1_SHA256PasswordSalt                                           string      `json:"Users.1.SHA256PasswordSalt"`
	Users_1_SolEnable                                                    string      `json:"Users.1.SolEnable"`
	Users_1_UserName                                                     string      `json:"Users.1.UserName"`
	Users_10_AuthenticationProtocol                                      string      `json:"Users.10.AuthenticationProtocol"`
	Users_10_Enable                                                      string      `json:"Users.10.Enable"`
	Users_10_IPMIKey                                                     string      `json:"Users.10.IPMIKey"`
	Users_10_IpmiLanPrivilege                                            string      `json:"Users.10.IpmiLanPrivilege"`
	Users_10_IpmiSerialPrivilege                                         string      `json:"Users.10.IpmiSerialPrivilege"`
	Users_10_MD5v3Key                                                    string      `json:"Users.10.MD5v3Key"`
	Users_10_Password                                                    interface{} `json:"Users.10.Password"`
	Users_10_PrivacyProtocol                                             string      `json:"Users.10.PrivacyProtocol"`
	Users_10_Privilege                                                   int         `json:"Users.10.Privilege"`
	Users_10_ProtocolEnable                                              string      `json:"Users.10.ProtocolEnable"`
	Users_10_SHA1v3Key                                                   string      `json:"Users.10.SHA1v3Key"`
	Users_10_SHA256Password                                              string      `json:"Users.10.SHA256Password"`
	Users_10_SHA256PasswordSalt                                          string      `json:"Users.10.SHA256PasswordSalt"`
	Users_10_SolEnable                                                   string      `json:"Users.10.SolEnable"`
	Users_10_UserName                                                    string      `json:"Users.10.UserName"`
	Users_11_AuthenticationProtocol                                      string      `json:"Users.11.AuthenticationProtocol"`
	Users_11_Enable                                                      string      `json:"Users.11.Enable"`
	Users_11_IPMIKey                                                     string      `json:"Users.11.IPMIKey"`
	Users_11_IpmiLanPrivilege                                            string      `json:"Users.11.IpmiLanPrivilege"`
	Users_11_IpmiSerialPrivilege                                         string      `json:"Users.11.IpmiSerialPrivilege"`
	Users_11_MD5v3Key                                                    string      `json:"Users.11.MD5v3Key"`
	Users_11_Password                                                    interface{} `json:"Users.11.Password"`
	Users_11_PrivacyProtocol                                             string      `json:"Users.11.PrivacyProtocol"`
	Users_11_Privilege                                                   int         `json:"Users.11.Privilege"`
	Users_11_ProtocolEnable                                              string      `json:"Users.11.ProtocolEnable"`
	Users_11_SHA1v3Key                                                   string      `json:"Users.11.SHA1v3Key"`
	Users_11_SHA256Password                                              string      `json:"Users.11.SHA256Password"`
	Users_11_SHA256PasswordSalt                                          string      `json:"Users.11.SHA256PasswordSalt"`
	Users_11_SolEnable                                                   string      `json:"Users.11.SolEnable"`
	Users_11_UserName                                                    string      `json:"Users.11.UserName"`
	Users_12_AuthenticationProtocol                                      string      `json:"Users.12.AuthenticationProtocol"`
	Users_12_Enable                                                      string      `json:"Users.12.Enable"`
	Users_12_IPMIKey                                                     string      `json:"Users.12.IPMIKey"`
	Users_12_IpmiLanPrivilege                                            string      `json:"Users.12.IpmiLanPrivilege"`
	Users_12_IpmiSerialPrivilege                                         string      `json:"Users.12.IpmiSerialPrivilege"`
	Users_12_MD5v3Key                                                    string      `json:"Users.12.MD5v3Key"`
	Users_12_Password                                                    interface{} `json:"Users.12.Password"`
	Users_12_PrivacyProtocol                                             string      `json:"Users.12.PrivacyProtocol"`
	Users_12_Privilege                                                   int         `json:"Users.12.Privilege"`
	Users_12_ProtocolEnable                                              string      `json:"Users.12.ProtocolEnable"`
	Users_12_SHA1v3Key                                                   string      `json:"Users.12.SHA1v3Key"`
	Users_12_SHA256Password                                              string      `json:"Users.12.SHA256Password"`
	Users_12_SHA256PasswordSalt                                          string      `json:"Users.12.SHA256PasswordSalt"`
	Users_12_SolEnable                                                   string      `json:"Users.12.SolEnable"`
	Users_12_UserName                                                    string      `json:"Users.12.UserName"`
	Users_13_AuthenticationProtocol                                      string      `json:"Users.13.AuthenticationProtocol"`
	Users_13_Enable                                                      string      `json:"Users.13.Enable"`
	Users_13_IPMIKey                                                     string      `json:"Users.13.IPMIKey"`
	Users_13_IpmiLanPrivilege                                            string      `json:"Users.13.IpmiLanPrivilege"`
	Users_13_IpmiSerialPrivilege                                         string      `json:"Users.13.IpmiSerialPrivilege"`
	Users_13_MD5v3Key                                                    string      `json:"Users.13.MD5v3Key"`
	Users_13_Password                                                    interface{} `json:"Users.13.Password"`
	Users_13_PrivacyProtocol                                             string      `json:"Users.13.PrivacyProtocol"`
	Users_13_Privilege                                                   int         `json:"Users.13.Privilege"`
	Users_13_ProtocolEnable                                              string      `json:"Users.13.ProtocolEnable"`
	Users_13_SHA1v3Key                                                   string      `json:"Users.13.SHA1v3Key"`
	Users_13_SHA256Password                                              string      `json:"Users.13.SHA256Password"`
	Users_13_SHA256PasswordSalt                                          string      `json:"Users.13.SHA256PasswordSalt"`
	Users_13_SolEnable                                                   string      `json:"Users.13.SolEnable"`
	Users_13_UserName                                                    string      `json:"Users.13.UserName"`
	Users_14_AuthenticationProtocol                                      string      `json:"Users.14.AuthenticationProtocol"`
	Users_14_Enable                                                      string      `json:"Users.14.Enable"`
	Users_14_IPMIKey                                                     string      `json:"Users.14.IPMIKey"`
	Users_14_IpmiLanPrivilege                                            string      `json:"Users.14.IpmiLanPrivilege"`
	Users_14_IpmiSerialPrivilege                                         string      `json:"Users.14.IpmiSerialPrivilege"`
	Users_14_MD5v3Key                                                    string      `json:"Users.14.MD5v3Key"`
	Users_14_Password                                                    interface{} `json:"Users.14.Password"`
	Users_14_PrivacyProtocol                                             string      `json:"Users.14.PrivacyProtocol"`
	Users_14_Privilege                                                   int         `json:"Users.14.Privilege"`
	Users_14_ProtocolEnable                                              string      `json:"Users.14.ProtocolEnable"`
	Users_14_SHA1v3Key                                                   string      `json:"Users.14.SHA1v3Key"`
	Users_14_SHA256Password                                              string      `json:"Users.14.SHA256Password"`
	Users_14_SHA256PasswordSalt                                          string      `json:"Users.14.SHA256PasswordSalt"`
	Users_14_SolEnable                                                   string      `json:"Users.14.SolEnable"`
	Users_14_UserName                                                    string      `json:"Users.14.UserName"`
	Users_15_AuthenticationProtocol                                      string      `json:"Users.15.AuthenticationProtocol"`
	Users_15_Enable                                                      string      `json:"Users.15.Enable"`
	Users_15_IPMIKey                                                     string      `json:"Users.15.IPMIKey"`
	Users_15_IpmiLanPrivilege                                            string      `json:"Users.15.IpmiLanPrivilege"`
	Users_15_IpmiSerialPrivilege                                         string      `json:"Users.15.IpmiSerialPrivilege"`
	Users_15_MD5v3Key                                                    string      `json:"Users.15.MD5v3Key"`
	Users_15_Password                                                    interface{} `json:"Users.15.Password"`
	Users_15_PrivacyProtocol                                             string      `json:"Users.15.PrivacyProtocol"`
	Users_15_Privilege                                                   int         `json:"Users.15.Privilege"`
	Users_15_ProtocolEnable                                              string      `json:"Users.15.ProtocolEnable"`
	Users_15_SHA1v3Key                                                   string      `json:"Users.15.SHA1v3Key"`
	Users_15_SHA256Password                                              string      `json:"Users.15.SHA256Password"`
	Users_15_SHA256PasswordSalt                                          string      `json:"Users.15.SHA256PasswordSalt"`
	Users_15_SolEnable                                                   string      `json:"Users.15.SolEnable"`
	Users_15_UserName                                                    string      `json:"Users.15.UserName"`
	Users_16_AuthenticationProtocol                                      string      `json:"Users.16.AuthenticationProtocol"`
	Users_16_Enable                                                      string      `json:"Users.16.Enable"`
	Users_16_IPMIKey                                                     string      `json:"Users.16.IPMIKey"`
	Users_16_IpmiLanPrivilege                                            string      `json:"Users.16.IpmiLanPrivilege"`
	Users_16_IpmiSerialPrivilege                                         string      `json:"Users.16.IpmiSerialPrivilege"`
	Users_16_MD5v3Key                                                    string      `json:"Users.16.MD5v3Key"`
	Users_16_Password                                                    interface{} `json:"Users.16.Password"`
	Users_16_PrivacyProtocol                                             string      `json:"Users.16.PrivacyProtocol"`
	Users_16_Privilege                                                   int         `json:"Users.16.Privilege"`
	Users_16_ProtocolEnable                                              string      `json:"Users.16.ProtocolEnable"`
	Users_16_SHA1v3Key                                                   string      `json:"Users.16.SHA1v3Key"`
	Users_16_SHA256Password                                              string      `json:"Users.16.SHA256Password"`
	Users_16_SHA256PasswordSalt                                          string      `json:"Users.16.SHA256PasswordSalt"`
	Users_16_SolEnable                                                   string      `json:"Users.16.SolEnable"`
	Users_16_UserName                                                    string      `json:"Users.16.UserName"`
	Users_2_AuthenticationProtocol                                       string      `json:"Users.2.AuthenticationProtocol"`
	Users_2_Enable                                                       string      `json:"Users.2.Enable"`
	Users_2_IPMIKey                                                      string      `json:"Users.2.IPMIKey"`
	Users_2_IpmiLanPrivilege                                             string      `json:"Users.2.IpmiLanPrivilege"`
	Users_2_IpmiSerialPrivilege                                          string      `json:"Users.2.IpmiSerialPrivilege"`
	Users_2_MD5v3Key                                                     string      `json:"Users.2.MD5v3Key"`
	Users_2_Password                                                     interface{} `json:"Users.2.Password"`
	Users_2_PrivacyProtocol                                              string      `json:"Users.2.PrivacyProtocol"`
	Users_2_Privilege                                                    int         `json:"Users.2.Privilege"`
	Users_2_ProtocolEnable                                               string      `json:"Users.2.ProtocolEnable"`
	Users_2_SHA1v3Key                                                    string      `json:"Users.2.SHA1v3Key"`
	Users_2_SHA256Password                                               string      `json:"Users.2.SHA256Password"`
	Users_2_SHA256PasswordSalt                                           string      `json:"Users.2.SHA256PasswordSalt"`
	Users_2_SolEnable                                                    string      `json:"Users.2.SolEnable"`
	Users_2_UserName                                                     string      `json:"Users.2.UserName"`
	Users_3_AuthenticationProtocol                                       string      `json:"Users.3.AuthenticationProtocol"`
	Users_3_Enable                                                       string      `json:"Users.3.Enable"`
	Users_3_IPMIKey                                                      string      `json:"Users.3.IPMIKey"`
	Users_3_IpmiLanPrivilege                                             string      `json:"Users.3.IpmiLanPrivilege"`
	Users_3_IpmiSerialPrivilege                                          string      `json:"Users.3.IpmiSerialPrivilege"`
	Users_3_MD5v3Key                                                     string      `json:"Users.3.MD5v3Key"`
	Users_3_Password                                                     interface{} `json:"Users.3.Password"`
	Users_3_PrivacyProtocol                                              string      `json:"Users.3.PrivacyProtocol"`
	Users_3_Privilege                                                    int         `json:"Users.3.Privilege"`
	Users_3_ProtocolEnable                                               string      `json:"Users.3.ProtocolEnable"`
	Users_3_SHA1v3Key                                                    string      `json:"Users.3.SHA1v3Key"`
	Users_3_SHA256Password                                               string      `json:"Users.3.SHA256Password"`
	Users_3_SHA256PasswordSalt                                           string      `json:"Users.3.SHA256PasswordSalt"`
	Users_3_SolEnable                                                    string      `json:"Users.3.SolEnable"`
	Users_3_UserName                                                     string      `json:"Users.3.UserName"`
	Users_4_AuthenticationProtocol                                       string      `json:"Users.4.AuthenticationProtocol"`
	Users_4_Enable                                                       string      `json:"Users.4.Enable"`
	Users_4_IPMIKey                                                      string      `json:"Users.4.IPMIKey"`
	Users_4_IpmiLanPrivilege                                             string      `json:"Users.4.IpmiLanPrivilege"`
	Users_4_IpmiSerialPrivilege                                          string      `json:"Users.4.IpmiSerialPrivilege"`
	Users_4_MD5v3Key                                                     string      `json:"Users.4.MD5v3Key"`
	Users_4_Password                                                     interface{} `json:"Users.4.Password"`
	Users_4_PrivacyProtocol                                              string      `json:"Users.4.PrivacyProtocol"`
	Users_4_Privilege                                                    int         `json:"Users.4.Privilege"`
	Users_4_ProtocolEnable                                               string      `json:"Users.4.ProtocolEnable"`
	Users_4_SHA1v3Key                                                    string      `json:"Users.4.SHA1v3Key"`
	Users_4_SHA256Password                                               string      `json:"Users.4.SHA256Password"`
	Users_4_SHA256PasswordSalt                                           string      `json:"Users.4.SHA256PasswordSalt"`
	Users_4_SolEnable                                                    string      `json:"Users.4.SolEnable"`
	Users_4_UserName                                                     string      `json:"Users.4.UserName"`
	Users_5_AuthenticationProtocol                                       string      `json:"Users.5.AuthenticationProtocol"`
	Users_5_Enable                                                       string      `json:"Users.5.Enable"`
	Users_5_IPMIKey                                                      string      `json:"Users.5.IPMIKey"`
	Users_5_IpmiLanPrivilege                                             string      `json:"Users.5.IpmiLanPrivilege"`
	Users_5_IpmiSerialPrivilege                                          string      `json:"Users.5.IpmiSerialPrivilege"`
	Users_5_MD5v3Key                                                     string      `json:"Users.5.MD5v3Key"`
	Users_5_Password                                                     interface{} `json:"Users.5.Password"`
	Users_5_PrivacyProtocol                                              string      `json:"Users.5.PrivacyProtocol"`
	Users_5_Privilege                                                    int         `json:"Users.5.Privilege"`
	Users_5_ProtocolEnable                                               string      `json:"Users.5.ProtocolEnable"`
	Users_5_SHA1v3Key                                                    string      `json:"Users.5.SHA1v3Key"`
	Users_5_SHA256Password                                               string      `json:"Users.5.SHA256Password"`
	Users_5_SHA256PasswordSalt                                           string      `json:"Users.5.SHA256PasswordSalt"`
	Users_5_SolEnable                                                    string      `json:"Users.5.SolEnable"`
	Users_5_UserName                                                     string      `json:"Users.5.UserName"`
	Users_6_AuthenticationProtocol                                       string      `json:"Users.6.AuthenticationProtocol"`
	Users_6_Enable                                                       string      `json:"Users.6.Enable"`
	Users_6_IPMIKey                                                      string      `json:"Users.6.IPMIKey"`
	Users_6_IpmiLanPrivilege                                             string      `json:"Users.6.IpmiLanPrivilege"`
	Users_6_IpmiSerialPrivilege                                          string      `json:"Users.6.IpmiSerialPrivilege"`
	Users_6_MD5v3Key                                                     string      `json:"Users.6.MD5v3Key"`
	Users_6_Password                                                     interface{} `json:"Users.6.Password"`
	Users_6_PrivacyProtocol                                              string      `json:"Users.6.PrivacyProtocol"`
	Users_6_Privilege                                                    int         `json:"Users.6.Privilege"`
	Users_6_ProtocolEnable                                               string      `json:"Users.6.ProtocolEnable"`
	Users_6_SHA1v3Key                                                    string      `json:"Users.6.SHA1v3Key"`
	Users_6_SHA256Password                                               string      `json:"Users.6.SHA256Password"`
	Users_6_SHA256PasswordSalt                                           string      `json:"Users.6.SHA256PasswordSalt"`
	Users_6_SolEnable                                                    string      `json:"Users.6.SolEnable"`
	Users_6_UserName                                                     string      `json:"Users.6.UserName"`
	Users_7_AuthenticationProtocol                                       string      `json:"Users.7.AuthenticationProtocol"`
	Users_7_Enable                                                       string      `json:"Users.7.Enable"`
	Users_7_IPMIKey                                                      string      `json:"Users.7.IPMIKey"`
	Users_7_IpmiLanPrivilege                                             string      `json:"Users.7.IpmiLanPrivilege"`
	Users_7_IpmiSerialPrivilege                                          string      `json:"Users.7.IpmiSerialPrivilege"`
	Users_7_MD5v3Key                                                     string      `json:"Users.7.MD5v3Key"`
	Users_7_Password                                                     interface{} `json:"Users.7.Password"`
	Users_7_PrivacyProtocol                                              string      `json:"Users.7.PrivacyProtocol"`
	Users_7_Privilege                                                    int         `json:"Users.7.Privilege"`
	Users_7_ProtocolEnable                                               string      `json:"Users.7.ProtocolEnable"`
	Users_7_SHA1v3Key                                                    string      `json:"Users.7.SHA1v3Key"`
	Users_7_SHA256Password                                               string      `json:"Users.7.SHA256Password"`
	Users_7_SHA256PasswordSalt                                           string      `json:"Users.7.SHA256PasswordSalt"`
	Users_7_SolEnable                                                    string      `json:"Users.7.SolEnable"`
	Users_7_UserName                                                     string      `json:"Users.7.UserName"`
	Users_8_AuthenticationProtocol                                       string      `json:"Users.8.AuthenticationProtocol"`
	Users_8_Enable                                                       string      `json:"Users.8.Enable"`
	Users_8_IPMIKey                                                      string      `json:"Users.8.IPMIKey"`
	Users_8_IpmiLanPrivilege                                             string      `json:"Users.8.IpmiLanPrivilege"`
	Users_8_IpmiSerialPrivilege                                          string      `json:"Users.8.IpmiSerialPrivilege"`
	Users_8_MD5v3Key                                                     string      `json:"Users.8.MD5v3Key"`
	Users_8_Password                                                     interface{} `json:"Users.8.Password"`
	Users_8_PrivacyProtocol                                              string      `json:"Users.8.PrivacyProtocol"`
	Users_8_Privilege                                                    int         `json:"Users.8.Privilege"`
	Users_8_ProtocolEnable                                               string      `json:"Users.8.ProtocolEnable"`
	Users_8_SHA1v3Key                                                    string      `json:"Users.8.SHA1v3Key"`
	Users_8_SHA256Password                                               string      `json:"Users.8.SHA256Password"`
	Users_8_SHA256PasswordSalt                                           string      `json:"Users.8.SHA256PasswordSalt"`
	Users_8_SolEnable                                                    string      `json:"Users.8.SolEnable"`
	Users_8_UserName                                                     string      `json:"Users.8.UserName"`
	Users_9_AuthenticationProtocol                                       string      `json:"Users.9.AuthenticationProtocol"`
	Users_9_Enable                                                       string      `json:"Users.9.Enable"`
	Users_9_IPMIKey                                                      string      `json:"Users.9.IPMIKey"`
	Users_9_IpmiLanPrivilege                                             string      `json:"Users.9.IpmiLanPrivilege"`
	Users_9_IpmiSerialPrivilege                                          string      `json:"Users.9.IpmiSerialPrivilege"`
	Users_9_MD5v3Key                                                     string      `json:"Users.9.MD5v3Key"`
	Users_9_Password                                                     interface{} `json:"Users.9.Password"`
	Users_9_PrivacyProtocol                                              string      `json:"Users.9.PrivacyProtocol"`
	Users_9_Privilege                                                    int         `json:"Users.9.Privilege"`
	Users_9_ProtocolEnable                                               string      `json:"Users.9.ProtocolEnable"`
	Users_9_SHA1v3Key                                                    string      `json:"Users.9.SHA1v3Key"`
	Users_9_SHA256Password                                               string      `json:"Users.9.SHA256Password"`
	Users_9_SHA256PasswordSalt                                           string      `json:"Users.9.SHA256PasswordSalt"`
	Users_9_SolEnable                                                    string      `json:"Users.9.SolEnable"`
	Users_9_UserName                                                     string      `json:"Users.9.UserName"`
	VNCServer_1_ActiveSessions                                           int         `json:"VNCServer.1.ActiveSessions"`
	VNCServer_1_Enable                                                   string      `json:"VNCServer.1.Enable"`
	VNCServer_1_LowerEncryptionBitLength                                 string      `json:"VNCServer.1.LowerEncryptionBitLength"`
	VNCServer_1_MaxSessions                                              int         `json:"VNCServer.1.MaxSessions"`
	VNCServer_1_Password                                                 interface{} `json:"VNCServer.1.Password"`
	VNCServer_1_Port                                                     int         `json:"VNCServer.1.Port"`
	VNCServer_1_SSLEncryptionBitLength                                   string      `json:"VNCServer.1.SSLEncryptionBitLength"`
	VNCServer_1_Timeout                                                  int         `json:"VNCServer.1.Timeout"`
	VirtualConsole_1_AccessPrivilege                                     string      `json:"VirtualConsole.1.AccessPrivilege"`
	VirtualConsole_1_ActiveSessions                                      int         `json:"VirtualConsole.1.ActiveSessions"`
	VirtualConsole_1_AttachState                                         string      `json:"VirtualConsole.1.AttachState"`
	VirtualConsole_1_Enable                                              string      `json:"VirtualConsole.1.Enable"`
	VirtualConsole_1_EncryptEnable                                       string      `json:"VirtualConsole.1.EncryptEnable"`
	VirtualConsole_1_LocalDisable                                        string      `json:"VirtualConsole.1.LocalDisable"`
	VirtualConsole_1_LocalVideo                                          string      `json:"VirtualConsole.1.LocalVideo"`
	VirtualConsole_1_MaxSessions                                         int         `json:"VirtualConsole.1.MaxSessions"`
	VirtualConsole_1_PluginType                                          string      `json:"VirtualConsole.1.PluginType"`
	VirtualConsole_1_Port                                                int         `json:"VirtualConsole.1.Port"`
	VirtualConsole_1_Timeout                                             int         `json:"VirtualConsole.1.Timeout"`
	VirtualConsole_1_TimeoutEnable                                       string      `json:"VirtualConsole.1.TimeoutEnable"`
	VirtualMedia_1_ActiveSessions                                        int         `json:"VirtualMedia.1.ActiveSessions"`
	VirtualMedia_1_Attached                                              string      `json:"VirtualMedia.1.Attached"`
	VirtualMedia_1_BootOnce                                              string      `json:"VirtualMedia.1.BootOnce"`
	VirtualMedia_1_Enable                                                string      `json:"VirtualMedia.1.Enable"`
	VirtualMedia_1_EncryptEnable                                         string      `json:"VirtualMedia.1.EncryptEnable"`
	VirtualMedia_1_FloppyEmulation                                       string      `json:"VirtualMedia.1.FloppyEmulation"`
	VirtualMedia_1_KeyEnable                                             string      `json:"VirtualMedia.1.KeyEnable"`
	VirtualMedia_1_MaxSessions                                           int         `json:"VirtualMedia.1.MaxSessions"`
	WebServer_1_CustomCipherString                                       string      `json:"WebServer.1.CustomCipherString"`
	WebServer_1_Enable                                                   string      `json:"WebServer.1.Enable"`
	WebServer_1_HttpPort                                                 int         `json:"WebServer.1.HttpPort"`
	WebServer_1_HttpsPort                                                int         `json:"WebServer.1.HttpsPort"`
	WebServer_1_HttpsRedirection                                         string      `json:"WebServer.1.HttpsRedirection"`
	WebServer_1_LowerEncryptionBitLength                                 string      `json:"WebServer.1.LowerEncryptionBitLength"`
	WebServer_1_MaxNumberOfSessions                                      int         `json:"WebServer.1.MaxNumberOfSessions"`
	WebServer_1_SSLEncryptionBitLength                                   string      `json:"WebServer.1.SSLEncryptionBitLength"`
	WebServer_1_TLSProtocol                                              string      `json:"WebServer.1.TLSProtocol"`
	WebServer_1_Timeout                                                  int         `json:"WebServer.1.Timeout"`
	WebServer_1_TitleBarOption                                           string      `json:"WebServer.1.TitleBarOption"`
	WebServer_1_TitleBarOptionCustom                                     string      `json:"WebServer.1.TitleBarOptionCustom"`
	VFlashPartition_1_AccessType                                         string      `json:"vFlashPartition.1.AccessType"`
	VFlashPartition_1_AttachState                                        string      `json:"vFlashPartition.1.AttachState"`
	VFlashPartition_1_EmulationType                                      string      `json:"vFlashPartition.1.EmulationType"`
	VFlashPartition_1_FormatType                                         string      `json:"vFlashPartition.1.FormatType"`
	VFlashPartition_1_IsGroupInstanceValid                               string      `json:"vFlashPartition.1.IsGroupInstanceValid"`
	VFlashPartition_1_Size                                               string      `json:"vFlashPartition.1.Size"`
	VFlashPartition_1_VolumeLabel                                        string      `json:"vFlashPartition.1.VolumeLabel"`
	VFlashPartition_10_AccessType                                        string      `json:"vFlashPartition.10.AccessType"`
	VFlashPartition_10_AttachState                                       string      `json:"vFlashPartition.10.AttachState"`
	VFlashPartition_10_EmulationType                                     string      `json:"vFlashPartition.10.EmulationType"`
	VFlashPartition_10_FormatType                                        string      `json:"vFlashPartition.10.FormatType"`
	VFlashPartition_10_IsGroupInstanceValid                              string      `json:"vFlashPartition.10.IsGroupInstanceValid"`
	VFlashPartition_10_Size                                              string      `json:"vFlashPartition.10.Size"`
	VFlashPartition_10_VolumeLabel                                       string      `json:"vFlashPartition.10.VolumeLabel"`
	VFlashPartition_11_AccessType                                        string      `json:"vFlashPartition.11.AccessType"`
	VFlashPartition_11_AttachState                                       string      `json:"vFlashPartition.11.AttachState"`
	VFlashPartition_11_EmulationType                                     string      `json:"vFlashPartition.11.EmulationType"`
	VFlashPartition_11_FormatType                                        string      `json:"vFlashPartition.11.FormatType"`
	VFlashPartition_11_IsGroupInstanceValid                              string      `json:"vFlashPartition.11.IsGroupInstanceValid"`
	VFlashPartition_11_Size                                              string      `json:"vFlashPartition.11.Size"`
	VFlashPartition_11_VolumeLabel                                       string      `json:"vFlashPartition.11.VolumeLabel"`
	VFlashPartition_12_AccessType                                        string      `json:"vFlashPartition.12.AccessType"`
	VFlashPartition_12_AttachState                                       string      `json:"vFlashPartition.12.AttachState"`
	VFlashPartition_12_EmulationType                                     string      `json:"vFlashPartition.12.EmulationType"`
	VFlashPartition_12_FormatType                                        string      `json:"vFlashPartition.12.FormatType"`
	VFlashPartition_12_IsGroupInstanceValid                              string      `json:"vFlashPartition.12.IsGroupInstanceValid"`
	VFlashPartition_12_Size                                              string      `json:"vFlashPartition.12.Size"`
	VFlashPartition_12_VolumeLabel                                       string      `json:"vFlashPartition.12.VolumeLabel"`
	VFlashPartition_13_AccessType                                        string      `json:"vFlashPartition.13.AccessType"`
	VFlashPartition_13_AttachState                                       string      `json:"vFlashPartition.13.AttachState"`
	VFlashPartition_13_EmulationType                                     string      `json:"vFlashPartition.13.EmulationType"`
	VFlashPartition_13_FormatType                                        string      `json:"vFlashPartition.13.FormatType"`
	VFlashPartition_13_IsGroupInstanceValid                              string      `json:"vFlashPartition.13.IsGroupInstanceValid"`
	VFlashPartition_13_Size                                              string      `json:"vFlashPartition.13.Size"`
	VFlashPartition_13_VolumeLabel                                       string      `json:"vFlashPartition.13.VolumeLabel"`
	VFlashPartition_14_AccessType                                        string      `json:"vFlashPartition.14.AccessType"`
	VFlashPartition_14_AttachState                                       string      `json:"vFlashPartition.14.AttachState"`
	VFlashPartition_14_EmulationType                                     string      `json:"vFlashPartition.14.EmulationType"`
	VFlashPartition_14_FormatType                                        string      `json:"vFlashPartition.14.FormatType"`
	VFlashPartition_14_IsGroupInstanceValid                              string      `json:"vFlashPartition.14.IsGroupInstanceValid"`
	VFlashPartition_14_Size                                              string      `json:"vFlashPartition.14.Size"`
	VFlashPartition_14_VolumeLabel                                       string      `json:"vFlashPartition.14.VolumeLabel"`
	VFlashPartition_15_AccessType                                        string      `json:"vFlashPartition.15.AccessType"`
	VFlashPartition_15_AttachState                                       string      `json:"vFlashPartition.15.AttachState"`
	VFlashPartition_15_EmulationType                                     string      `json:"vFlashPartition.15.EmulationType"`
	VFlashPartition_15_FormatType                                        string      `json:"vFlashPartition.15.FormatType"`
	VFlashPartition_15_IsGroupInstanceValid                              string      `json:"vFlashPartition.15.IsGroupInstanceValid"`
	VFlashPartition_15_Size                                              string      `json:"vFlashPartition.15.Size"`
	VFlashPartition_15_VolumeLabel                                       string      `json:"vFlashPartition.15.VolumeLabel"`
	VFlashPartition_16_AccessType                                        string      `json:"vFlashPartition.16.AccessType"`
	VFlashPartition_16_AttachState                                       string      `json:"vFlashPartition.16.AttachState"`
	VFlashPartition_16_EmulationType                                     string      `json:"vFlashPartition.16.EmulationType"`
	VFlashPartition_16_FormatType                                        string      `json:"vFlashPartition.16.FormatType"`
	VFlashPartition_16_IsGroupInstanceValid                              string      `json:"vFlashPartition.16.IsGroupInstanceValid"`
	VFlashPartition_16_Size                                              string      `json:"vFlashPartition.16.Size"`
	VFlashPartition_16_VolumeLabel                                       string      `json:"vFlashPartition.16.VolumeLabel"`
	VFlashPartition_2_AccessType                                         string      `json:"vFlashPartition.2.AccessType"`
	VFlashPartition_2_AttachState                                        string      `json:"vFlashPartition.2.AttachState"`
	VFlashPartition_2_EmulationType                                      string      `json:"vFlashPartition.2.EmulationType"`
	VFlashPartition_2_FormatType                                         string      `json:"vFlashPartition.2.FormatType"`
	VFlashPartition_2_IsGroupInstanceValid                               string      `json:"vFlashPartition.2.IsGroupInstanceValid"`
	VFlashPartition_2_Size                                               string      `json:"vFlashPartition.2.Size"`
	VFlashPartition_2_VolumeLabel                                        string      `json:"vFlashPartition.2.VolumeLabel"`
	VFlashPartition_3_AccessType                                         string      `json:"vFlashPartition.3.AccessType"`
	VFlashPartition_3_AttachState                                        string      `json:"vFlashPartition.3.AttachState"`
	VFlashPartition_3_EmulationType                                      string      `json:"vFlashPartition.3.EmulationType"`
	VFlashPartition_3_FormatType                                         string      `json:"vFlashPartition.3.FormatType"`
	VFlashPartition_3_IsGroupInstanceValid                               string      `json:"vFlashPartition.3.IsGroupInstanceValid"`
	VFlashPartition_3_Size                                               string      `json:"vFlashPartition.3.Size"`
	VFlashPartition_3_VolumeLabel                                        string      `json:"vFlashPartition.3.VolumeLabel"`
	VFlashPartition_4_AccessType                                         string      `json:"vFlashPartition.4.AccessType"`
	VFlashPartition_4_AttachState                                        string      `json:"vFlashPartition.4.AttachState"`
	VFlashPartition_4_EmulationType                                      string      `json:"vFlashPartition.4.EmulationType"`
	VFlashPartition_4_FormatType                                         string      `json:"vFlashPartition.4.FormatType"`
	VFlashPartition_4_IsGroupInstanceValid                               string      `json:"vFlashPartition.4.IsGroupInstanceValid"`
	VFlashPartition_4_Size                                               string      `json:"vFlashPartition.4.Size"`
	VFlashPartition_4_VolumeLabel                                        string      `json:"vFlashPartition.4.VolumeLabel"`
	VFlashPartition_5_AccessType                                         string      `json:"vFlashPartition.5.AccessType"`
	VFlashPartition_5_AttachState                                        string      `json:"vFlashPartition.5.AttachState"`
	VFlashPartition_5_EmulationType                                      string      `json:"vFlashPartition.5.EmulationType"`
	VFlashPartition_5_FormatType                                         string      `json:"vFlashPartition.5.FormatType"`
	VFlashPartition_5_IsGroupInstanceValid                               string      `json:"vFlashPartition.5.IsGroupInstanceValid"`
	VFlashPartition_5_Size                                               string      `json:"vFlashPartition.5.Size"`
	VFlashPartition_5_VolumeLabel                                        string      `json:"vFlashPartition.5.VolumeLabel"`
	VFlashPartition_6_AccessType                                         string      `json:"vFlashPartition.6.AccessType"`
	VFlashPartition_6_AttachState                                        string      `json:"vFlashPartition.6.AttachState"`
	VFlashPartition_6_EmulationType                                      string      `json:"vFlashPartition.6.EmulationType"`
	VFlashPartition_6_FormatType                                         string      `json:"vFlashPartition.6.FormatType"`
	VFlashPartition_6_IsGroupInstanceValid                               string      `json:"vFlashPartition.6.IsGroupInstanceValid"`
	VFlashPartition_6_Size                                               string      `json:"vFlashPartition.6.Size"`
	VFlashPartition_6_VolumeLabel                                        string      `json:"vFlashPartition.6.VolumeLabel"`
	VFlashPartition_7_AccessType                                         string      `json:"vFlashPartition.7.AccessType"`
	VFlashPartition_7_AttachState                                        string      `json:"vFlashPartition.7.AttachState"`
	VFlashPartition_7_EmulationType                                      string      `json:"vFlashPartition.7.EmulationType"`
	VFlashPartition_7_FormatType                                         string      `json:"vFlashPartition.7.FormatType"`
	VFlashPartition_7_IsGroupInstanceValid                               string      `json:"vFlashPartition.7.IsGroupInstanceValid"`
	VFlashPartition_7_Size                                               string      `json:"vFlashPartition.7.Size"`
	VFlashPartition_7_VolumeLabel                                        string      `json:"vFlashPartition.7.VolumeLabel"`
	VFlashPartition_8_AccessType                                         string      `json:"vFlashPartition.8.AccessType"`
	VFlashPartition_8_AttachState                                        string      `json:"vFlashPartition.8.AttachState"`
	VFlashPartition_8_EmulationType                                      string      `json:"vFlashPartition.8.EmulationType"`
	VFlashPartition_8_FormatType                                         string      `json:"vFlashPartition.8.FormatType"`
	VFlashPartition_8_IsGroupInstanceValid                               string      `json:"vFlashPartition.8.IsGroupInstanceValid"`
	VFlashPartition_8_Size                                               string      `json:"vFlashPartition.8.Size"`
	VFlashPartition_8_VolumeLabel                                        string      `json:"vFlashPartition.8.VolumeLabel"`
	VFlashPartition_9_AccessType                                         string      `json:"vFlashPartition.9.AccessType"`
	VFlashPartition_9_AttachState                                        string      `json:"vFlashPartition.9.AttachState"`
	VFlashPartition_9_EmulationType                                      string      `json:"vFlashPartition.9.EmulationType"`
	VFlashPartition_9_FormatType                                         string      `json:"vFlashPartition.9.FormatType"`
	VFlashPartition_9_IsGroupInstanceValid                               string      `json:"vFlashPartition.9.IsGroupInstanceValid"`
	VFlashPartition_9_Size                                               string      `json:"vFlashPartition.9.Size"`
	VFlashPartition_9_VolumeLabel                                        string      `json:"vFlashPartition.9.VolumeLabel"`
	VFlashSD_1_AvailableSize                                             string      `json:"vFlashSD.1.AvailableSize"`
	VFlashSD_1_Bitmap                                                    string      `json:"vFlashSD.1.Bitmap"`
	VFlashSD_1_Enable                                                    string      `json:"vFlashSD.1.Enable"`
	VFlashSD_1_Health                                                    string      `json:"vFlashSD.1.Health"`
	VFlashSD_1_Initialized                                               string      `json:"vFlashSD.1.Initialized"`
	VFlashSD_1_Licensed                                                  string      `json:"vFlashSD.1.Licensed"`
	VFlashSD_1_Presence                                                  string      `json:"vFlashSD.1.Presence"`
	VFlashSD_1_Size                                                      string      `json:"vFlashSD.1.Size"`
	VFlashSD_1_WriteProtect                                              string      `json:"vFlashSD.1.WriteProtect"`
}

//LifeCycleAttrDell ... LifeCycle Controller Attributes from the Redfish API
type LifeCycleAttrDell struct {
	RedfishSettings struct {
		_odata_context string
		_odata_type    string
		SettingsObject struct {
			_odata_id string
		} `json:"SettingsObject"`
	} `json:"@Redfish.Settings"`
	_odata_context    string
	_odata_id         string
	_odata_type       string
	AttributeRegistry string `json:"AttributeRegistry"`
	Attributes        struct {
		AutoBackup                          string      `json:"LCAttributes.1.AutoBackup"`
		AutoDiscovery                       string      `json:"LCAttributes.1.AutoDiscovery"`
		AutoUpdate                          string      `json:"LCAttributes.1.AutoUpdate"`
		BIOSRTDRequested                    string      `json:"LCAttributes.1.BIOSRTDRequested"`
		CollectSystemInventoryOnRestart     string      `json:"LCAttributes.1.CollectSystemInventoryOnRestart"`
		DiscoveryFactoryDefaults            string      `json:"LCAttributes.1.DiscoveryFactoryDefaults"`
		IPAddress                           string      `json:"LCAttributes.1.IPAddress"`
		IPChangeNotifyPS                    string      `json:"LCAttributes.1.IPChangeNotifyPS"`
		IgnoreCertWarning                   string      `json:"LCAttributes.1.IgnoreCertWarning"`
		Licensed                            string      `json:"LCAttributes.1.Licensed"`
		LifecycleControllerState            string      `json:"LCAttributes.1.LifecycleControllerState"`
		PartConfigurationUpdate             string      `json:"LCAttributes.1.PartConfigurationUpdate"`
		PartFirmwareUpdate                  string      `json:"LCAttributes.1.PartFirmwareUpdate"`
		ProvisioningServer                  string      `json:"LCAttributes.1.ProvisioningServer"`
		StorageHealthRollupStatus           int         `json:"LCAttributes.1.StorageHealthRollupStatus"`
		SystemID                            string      `json:"LCAttributes.1.SystemID"`
		UserProxyPassword                   interface{} `json:"LCAttributes.1.UserProxyPassword"`
		UserProxyPort                       string      `json:"LCAttributes.1.UserProxyPort"`
		UserProxyServer                     string      `json:"LCAttributes.1.UserProxyServer"`
		UserProxyType                       string      `json:"LCAttributes.1.UserProxyType"`
		UserProxyUserName                   string      `json:"LCAttributes.1.UserProxyUserName"`
		VirtualAddressManagementApplication string      `json:"LCAttributes.1.VirtualAddressManagementApplication"`
	} `json:"Attributes"`
	Description string `json:"Description"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
}

//BiosDell ... Bios Settings from the Redfish API
type BiosAttrDell struct {
	Redfish_Settings struct {
		_odata_context string
		_odata_type    string
		SettingsObject struct {
			_odata_id string
		} `json:"SettingsObject"`
	} `json:"@Redfish.Settings"`
	_odata_context string
	_odata_id      string
	_odata_type    string
	Actions        struct {
		Bios_ChangePassword struct {
			Target string `json:"target"`
		} `json:"#Bios.ChangePassword"`
		Bios_ResetBios struct {
			Target string `json:"target"`
		} `json:"#Bios.ResetBios"`
	} `json:"Actions"`
	AttributeRegistry string             `json:"AttributeRegistry"`
	Attributes        BiosAttributesData `json:"Attributes"`
	Description       string             `json:"Description"`
	ID                string             `json:"Id"`
	Name              string             `json:"Name"`
}

type BiosAttributesData struct {
	AcPwrRcvry                   string      `json:"AcPwrRcvry"`
	AcPwrRcvryDelay              string      `json:"AcPwrRcvryDelay"`
	AcPwrRcvryUserDelay          int         `json:"AcPwrRcvryUserDelay"`
	AesNi                        string      `json:"AesNi"`
	AssetTag                     string      `json:"AssetTag"`
	AuthorizeDeviceFirmware      string      `json:"AuthorizeDeviceFirmware"`
	BootMode                     string      `json:"BootMode"`
	BootSeqRetry                 string      `json:"BootSeqRetry"`
	ConTermType                  string      `json:"ConTermType"`
	ControlledTurbo              string      `json:"ControlledTurbo"`
	CorrEccSmi                   string      `json:"CorrEccSmi"`
	CPUInterconnectBusLinkPower  string      `json:"CpuInterconnectBusLinkPower"`
	CPUInterconnectBusSpeed      string      `json:"CpuInterconnectBusSpeed"`
	CurrentEmbVideoState         string      `json:"CurrentEmbVideoState"`
	CurrentMemOpModeState        string      `json:"CurrentMemOpModeState"`
	DcuIPPrefetcher              string      `json:"DcuIpPrefetcher"`
	DcuStreamerPrefetcher        string      `json:"DcuStreamerPrefetcher"`
	DellAutoDiscovery            string      `json:"DellAutoDiscovery"`
	DellWyseP25BIOSAccess        string      `json:"DellWyseP25BIOSAccess"`
	DynamicCoreAllocation        string      `json:"DynamicCoreAllocation"`
	EmbSata                      string      `json:"EmbSata"`
	EmbVideo                     string      `json:"EmbVideo"`
	EnergyPerformanceBias        string      `json:"EnergyPerformanceBias"`
	ErrPrompt                    string      `json:"ErrPrompt"`
	ExtSerialConnector           string      `json:"ExtSerialConnector"`
	FailSafeBaud                 string      `json:"FailSafeBaud"`
	ForceInt10                   string      `json:"ForceInt10"`
	HddFailover                  string      `json:"HddFailover"`
	InBandManageabilityInterface string      `json:"InBandManageabilityInterface"`
	IntegratedNetwork1           string      `json:"IntegratedNetwork1"`
	IntelTxt                     string      `json:"IntelTxt"`
	InternalUsb                  string      `json:"InternalUsb"`
	IoatEngine                   string      `json:"IoatEngine"`
	LogicalProc                  string      `json:"LogicalProc"`
	MemFrequency                 string      `json:"MemFrequency"`
	MemOpMode                    string      `json:"MemOpMode"`
	MemPatrolScrub               string      `json:"MemPatrolScrub"`
	MemRefreshRate               string      `json:"MemRefreshRate"`
	MemTest                      string      `json:"MemTest"`
	MemoryMappedIOH              string      `json:"MemoryMappedIOH"`
	MmioAbove4Gb                 string      `json:"MmioAbove4Gb"`
	MonitorMwait                 string      `json:"MonitorMwait"`
	NodeInterleave               string      `json:"NodeInterleave"`
	NumLock                      string      `json:"NumLock"`
	NvmeMode                     string      `json:"NvmeMode"`
	OneTimeBootMode              string      `json:"OneTimeBootMode"`
	OneTimeBootSeqDev            string      `json:"OneTimeBootSeqDev"`
	OppSrefEn                    string      `json:"OppSrefEn"`
	OsWatchdogTimer              string      `json:"OsWatchdogTimer"`
	PasswordStatus               string      `json:"PasswordStatus"`
	PcieAspmL1                   string      `json:"PcieAspmL1"`
	PowerCycleRequest            string      `json:"PowerCycleRequest"`
	Proc1Brand                   string      `json:"Proc1Brand"`
	Proc1Id                      string      `json:"Proc1Id"`
	Proc1L2Cache                 string      `json:"Proc1L2Cache"`
	Proc1L3Cache                 string      `json:"Proc1L3Cache"`
	Proc1NumCores                int         `json:"Proc1NumCores"`
	Proc1TurboCoreNum            string      `json:"Proc1TurboCoreNum"`
	Proc2Brand                   string      `json:"Proc2Brand"`
	Proc2Id                      string      `json:"Proc2Id"`
	Proc2L2Cache                 string      `json:"Proc2L2Cache"`
	Proc2L3Cache                 string      `json:"Proc2L3Cache"`
	Proc2NumCores                int         `json:"Proc2NumCores"`
	Proc2TurboCoreNum            string      `json:"Proc2TurboCoreNum"`
	ProcAdjCacheLine             string      `json:"ProcAdjCacheLine"`
	ProcBusSpeed                 string      `json:"ProcBusSpeed"`
	ProcC1E                      string      `json:"ProcC1E"`
	ProcCStates                  string      `json:"ProcCStates"`
	ProcCoreSpeed                string      `json:"ProcCoreSpeed"`
	ProcCores                    string      `json:"ProcCores"`
	ProcHwPrefetcher             string      `json:"ProcHwPrefetcher"`
	ProcPwrPerf                  string      `json:"ProcPwrPerf"`
	ProcTurboMode                string      `json:"ProcTurboMode"`
	ProcVirtualization           string      `json:"ProcVirtualization"`
	ProcX2Apic                   string      `json:"ProcX2Apic"`
	PwrButton                    string      `json:"PwrButton"`
	RedirAfterBoot               string      `json:"RedirAfterBoot"`
	RedundantOsLocation          string      `json:"RedundantOsLocation"`
	SHA256SetupPassword          string      `json:"SHA256SetupPassword"`
	SHA256SetupPasswordSalt      string      `json:"SHA256SetupPasswordSalt"`
	SHA256SystemPassword         string      `json:"SHA256SystemPassword"`
	SHA256SystemPasswordSalt     string      `json:"SHA256SystemPasswordSalt"`
	SataPortA                    string      `json:"SataPortA"`
	SataPortACapacity            string      `json:"SataPortACapacity"`
	SataPortADriveType           string      `json:"SataPortADriveType"`
	SataPortAModel               string      `json:"SataPortAModel"`
	SataPortB                    string      `json:"SataPortB"`
	SataPortBCapacity            string      `json:"SataPortBCapacity"`
	SataPortBDriveType           string      `json:"SataPortBDriveType"`
	SataPortBModel               string      `json:"SataPortBModel"`
	SataPortC                    string      `json:"SataPortC"`
	SataPortCCapacity            string      `json:"SataPortCCapacity"`
	SataPortCDriveType           string      `json:"SataPortCDriveType"`
	SataPortCModel               string      `json:"SataPortCModel"`
	SataPortD                    string      `json:"SataPortD"`
	SataPortDCapacity            string      `json:"SataPortDCapacity"`
	SataPortDDriveType           string      `json:"SataPortDDriveType"`
	SataPortDModel               string      `json:"SataPortDModel"`
	SataPortE                    string      `json:"SataPortE"`
	SataPortECapacity            string      `json:"SataPortECapacity"`
	SataPortEDriveType           string      `json:"SataPortEDriveType"`
	SataPortEModel               string      `json:"SataPortEModel"`
	SataPortF                    string      `json:"SataPortF"`
	SataPortFCapacity            string      `json:"SataPortFCapacity"`
	SataPortFDriveType           string      `json:"SataPortFDriveType"`
	SataPortFModel               string      `json:"SataPortFModel"`
	SataPortG                    string      `json:"SataPortG"`
	SataPortGCapacity            string      `json:"SataPortGCapacity"`
	SataPortGDriveType           string      `json:"SataPortGDriveType"`
	SataPortGModel               string      `json:"SataPortGModel"`
	SataPortH                    string      `json:"SataPortH"`
	SataPortHCapacity            string      `json:"SataPortHCapacity"`
	SataPortHDriveType           string      `json:"SataPortHDriveType"`
	SataPortHModel               string      `json:"SataPortHModel"`
	SataPortI                    string      `json:"SataPortI"`
	SataPortICapacity            string      `json:"SataPortICapacity"`
	SataPortIDriveType           string      `json:"SataPortIDriveType"`
	SataPortIModel               string      `json:"SataPortIModel"`
	SataPortJ                    string      `json:"SataPortJ"`
	SataPortJCapacity            string      `json:"SataPortJCapacity"`
	SataPortJDriveType           string      `json:"SataPortJDriveType"`
	SataPortJModel               string      `json:"SataPortJModel"`
	SataPortK                    string      `json:"SataPortK"`
	SataPortKCapacity            string      `json:"SataPortKCapacity"`
	SataPortKDriveType           string      `json:"SataPortKDriveType"`
	SataPortKModel               string      `json:"SataPortKModel"`
	SataPortL                    string      `json:"SataPortL"`
	SataPortLCapacity            string      `json:"SataPortLCapacity"`
	SataPortLDriveType           string      `json:"SataPortLDriveType"`
	SataPortLModel               string      `json:"SataPortLModel"`
	SataPortM                    string      `json:"SataPortM"`
	SataPortMCapacity            string      `json:"SataPortMCapacity"`
	SataPortMDriveType           string      `json:"SataPortMDriveType"`
	SataPortMModel               string      `json:"SataPortMModel"`
	SataPortN                    string      `json:"SataPortN"`
	SataPortNCapacity            string      `json:"SataPortNCapacity"`
	SataPortNDriveType           string      `json:"SataPortNDriveType"`
	SataPortNModel               string      `json:"SataPortNModel"`
	SecureBoot                   string      `json:"SecureBoot"`
	SecureBootMode               string      `json:"SecureBootMode"`
	SecureBootPolicy             string      `json:"SecureBootPolicy"`
	SecurityFreezeLock           string      `json:"SecurityFreezeLock"`
	SerialComm                   string      `json:"SerialComm"`
	SerialPortAddress            string      `json:"SerialPortAddress"`
	SetBootOrderFqdd1            string      `json:"SetBootOrderFqdd1"`
	SetBootOrderFqdd10           string      `json:"SetBootOrderFqdd10"`
	SetBootOrderFqdd11           string      `json:"SetBootOrderFqdd11"`
	SetBootOrderFqdd12           string      `json:"SetBootOrderFqdd12"`
	SetBootOrderFqdd13           string      `json:"SetBootOrderFqdd13"`
	SetBootOrderFqdd14           string      `json:"SetBootOrderFqdd14"`
	SetBootOrderFqdd15           string      `json:"SetBootOrderFqdd15"`
	SetBootOrderFqdd16           string      `json:"SetBootOrderFqdd16"`
	SetBootOrderFqdd2            string      `json:"SetBootOrderFqdd2"`
	SetBootOrderFqdd3            string      `json:"SetBootOrderFqdd3"`
	SetBootOrderFqdd4            string      `json:"SetBootOrderFqdd4"`
	SetBootOrderFqdd5            string      `json:"SetBootOrderFqdd5"`
	SetBootOrderFqdd6            string      `json:"SetBootOrderFqdd6"`
	SetBootOrderFqdd7            string      `json:"SetBootOrderFqdd7"`
	SetBootOrderFqdd8            string      `json:"SetBootOrderFqdd8"`
	SetBootOrderFqdd9            string      `json:"SetBootOrderFqdd9"`
	SetLegacyHddOrderFqdd1       string      `json:"SetLegacyHddOrderFqdd1"`
	SetLegacyHddOrderFqdd10      string      `json:"SetLegacyHddOrderFqdd10"`
	SetLegacyHddOrderFqdd11      string      `json:"SetLegacyHddOrderFqdd11"`
	SetLegacyHddOrderFqdd12      string      `json:"SetLegacyHddOrderFqdd12"`
	SetLegacyHddOrderFqdd13      string      `json:"SetLegacyHddOrderFqdd13"`
	SetLegacyHddOrderFqdd14      string      `json:"SetLegacyHddOrderFqdd14"`
	SetLegacyHddOrderFqdd15      string      `json:"SetLegacyHddOrderFqdd15"`
	SetLegacyHddOrderFqdd16      string      `json:"SetLegacyHddOrderFqdd16"`
	SetLegacyHddOrderFqdd2       string      `json:"SetLegacyHddOrderFqdd2"`
	SetLegacyHddOrderFqdd3       string      `json:"SetLegacyHddOrderFqdd3"`
	SetLegacyHddOrderFqdd4       string      `json:"SetLegacyHddOrderFqdd4"`
	SetLegacyHddOrderFqdd5       string      `json:"SetLegacyHddOrderFqdd5"`
	SetLegacyHddOrderFqdd6       string      `json:"SetLegacyHddOrderFqdd6"`
	SetLegacyHddOrderFqdd7       string      `json:"SetLegacyHddOrderFqdd7"`
	SetLegacyHddOrderFqdd8       string      `json:"SetLegacyHddOrderFqdd8"`
	SetLegacyHddOrderFqdd9       string      `json:"SetLegacyHddOrderFqdd9"`
	SetupPassword                interface{} `json:"SetupPassword"`
	Slot1                        string      `json:"Slot1"`
	Slot1Bif                     string      `json:"Slot1Bif"`
	Slot2                        string      `json:"Slot2"`
	Slot2Bif                     string      `json:"Slot2Bif"`
	Slot3                        string      `json:"Slot3"`
	Slot3Bif                     string      `json:"Slot3Bif"`
	Slot4                        string      `json:"Slot4"`
	Slot4Bif                     string      `json:"Slot4Bif"`
	Slot5                        string      `json:"Slot5"`
	Slot5Bif                     string      `json:"Slot5Bif"`
	Slot6                        string      `json:"Slot6"`
	Slot6Bif                     string      `json:"Slot6Bif"`
	Slot7                        string      `json:"Slot7"`
	Slot7Bif                     string      `json:"Slot7Bif"`
	Slot8                        string      `json:"Slot8"`
	Slot8Bif                     string      `json:"Slot8Bif"`
	SriovGlobalEnable            string      `json:"SriovGlobalEnable"`
	SubNumaCluster               string      `json:"SubNumaCluster"`
	SysMemSize                   string      `json:"SysMemSize"`
	SysMemSpeed                  string      `json:"SysMemSpeed"`
	SysMemType                   string      `json:"SysMemType"`
	SysMemVolt                   string      `json:"SysMemVolt"`
	SysMfrContactInfo            string      `json:"SysMfrContactInfo"`
	SysPassword                  interface{} `json:"SysPassword"`
	SysProfile                   string      `json:"SysProfile"`
	SystemBiosVersion            string      `json:"SystemBiosVersion"`
	SystemCpldVersion            string      `json:"SystemCpldVersion"`
	SystemManufacturer           string      `json:"SystemManufacturer"`
	SystemMeVersion              string      `json:"SystemMeVersion"`
	SystemModelName              string      `json:"SystemModelName"`
	SystemServiceTag             string      `json:"SystemServiceTag"`
	TpmFirmware                  string      `json:"TpmFirmware"`
	TpmInfo                      string      `json:"TpmInfo"`
	TpmPpiBypassClear            string      `json:"TpmPpiBypassClear"`
	TpmPpiBypassProvision        string      `json:"TpmPpiBypassProvision"`
	TpmSecurity                  string      `json:"TpmSecurity"`
	UefiComplianceVersion        string      `json:"UefiComplianceVersion"`
	UefiVariableAccess           string      `json:"UefiVariableAccess"`
	UncoreFrequency              string      `json:"UncoreFrequency"`
	UpiPrefetch                  string      `json:"UpiPrefetch"`
	UsbManagedPort               string      `json:"UsbManagedPort"`
	UsbPorts                     string      `json:"UsbPorts"`
	VideoMem                     string      `json:"VideoMem"`
	WorkloadProfile              string      `json:"WorkloadProfile"`
	WriteCache                   string      `json:"WriteCache"`
	WriteDataCrc                 string      `json:"WriteDataCrc"`
}

//GetMacAddressDell ... Fetch the Mac Address Info from the Redfish API
type GetMacAddressDell struct {
	_odata_context                     string
	_odata_id                          string
	_odata_type                        string
	AutoNeg                            bool          `json:"AutoNeg"`
	Description                        string        `json:"Description"`
	FQDN                               interface{}   `json:"FQDN"`
	FullDuplex                         bool          `json:"FullDuplex"`
	HostName                           interface{}   `json:"HostName"`
	IPv4Addresses                      []interface{} `json:"IPv4Addresses"`
	IPv4Addresses_odata_count          int           `json:"IPv4Addresses@odata.count"`
	IPv6AddressPolicyTable             []interface{} `json:"IPv6AddressPolicyTable"`
	IPv6AddressPolicyTable_odata_count int           `json:"IPv6AddressPolicyTable@odata.count"`
	IPv6Addresses                      []interface{} `json:"IPv6Addresses"`
	IPv6Addresses_odata_count          int           `json:"IPv6Addresses@odata.count"`
	IPv6DefaultGateway                 interface{}   `json:"IPv6DefaultGateway"`
	IPv6StaticAddresses                []interface{} `json:"IPv6StaticAddresses"`
	IPv6StaticAddresses_odata_count    int           `json:"IPv6StaticAddresses@odata.count"`
	ID                                 string        `json:"Id"`
	InterfaceEnabled                   interface{}   `json:"InterfaceEnabled"`
	LinkStatus                         string        `json:"LinkStatus"`
	Links                              struct {
		Chassis struct {
			_odata_id string
		} `json:"Chassis"`
	} `json:"Links"`
	MACAddress              string        `json:"MACAddress"`
	MTUSize                 interface{}   `json:"MTUSize"`
	MaxIPv6StaticAddresses  interface{}   `json:"MaxIPv6StaticAddresses"`
	Name                    string        `json:"Name"`
	NameServers             []interface{} `json:"NameServers"`
	NameServers_odata_count int           `json:"NameServers@odata.count"`
	PermanentMACAddress     string        `json:"PermanentMACAddress"`
	SpeedMbps               int           `json:"SpeedMbps"`
	Status                  struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	UefiDevicePath string `json:"UefiDevicePath"`
	VLAN           string `json:"VLAN"`
}

// MACModelDell ...
type MACModelDell struct {
	MacName  string `json:"mac_name"`
	MacModel string `json:"mac_model"`
}

// NetworkPortsDell ...
type NetworkPortsDell struct {
	Odata_context              string      `json:"@odata.context"`
	Odata_id                   string      `json:"@odata.id"`
	Odata_type                 string      `json:"@odata.type"`
	ActiveLinkTechnology       string      `json:"ActiveLinkTechnology"`
	AssociatedNetworkAddresses []string    `json:"AssociatedNetworkAddresses"`
	CurrentLinkSpeedMbps       int64       `json:"CurrentLinkSpeedMbps"`
	Description                string      `json:"Description"`
	EEEEnabled                 interface{} `json:"EEEEnabled"`
	FlowControlConfiguration   string      `json:"FlowControlConfiguration"`
	FlowControlStatus          string      `json:"FlowControlStatus"`
	ID                         string      `json:"Id"`
	LinkStatus                 string      `json:"LinkStatus"`
	Name                       string      `json:"Name"`
	NetDevFuncMaxBWAlloc       []struct {
		MaxBWAllocPercent     interface{} `json:"MaxBWAllocPercent"`
		NetworkDeviceFunction struct {
			Odata_id string `json:"@odata.id"`
		} `json:"NetworkDeviceFunction"`
	} `json:"NetDevFuncMaxBWAlloc"`
	NetDevFuncMaxBWAlloc_odata_count int64 `json:"NetDevFuncMaxBWAlloc@odata.count"`
	NetDevFuncMinBWAlloc             []struct {
		MinBWAllocPercent     interface{} `json:"MinBWAllocPercent"`
		NetworkDeviceFunction struct {
			Odata_id string `json:"@odata.id"`
		} `json:"NetworkDeviceFunction"`
	} `json:"NetDevFuncMinBWAlloc"`
	NetDevFuncMinBWAlloc_odata_count int64 `json:"NetDevFuncMinBWAlloc@odata.count"`
	Oem                              struct {
		Dell struct {
			Odata_type             string `json:"@odata.type"`
			DellNetworkTransceiver struct {
				DeviceDescription string `json:"DeviceDescription"`
				FQDD              string `json:"FQDD"`
				Id                string `json:"Id"`
				IdentifierType    string `json:"IdentifierType"`
				InterfaceType     string `json:"InterfaceType"`
				Name              string `json:"Name"`
				PartNumber        string `json:"PartNumber"`
				Revision          string `json:"Revision"`
				SerialNumber      string `json:"SerialNumber"`
				VendorName        string `json:"VendorName"`
				Odata_context     string `json:"@odata.context"`
				Odata_type        string `json:"@odata.type"`
				Oddata_id         string `json:"@odata.id"`
			} `json:"DellNetworkTransceiver"`
		} `json:"Dell"`
	} `json:"Oem"`
	PhysicalPortNumber string `json:"PhysicalPortNumber"`
	Status             struct {
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
		State        string `json:"State"`
	} `json:"Status"`
	SupportedEthernetCapabilities             []string `json:"SupportedEthernetCapabilities"`
	SupportedEthernetCapabilities_odata_count int64    `json:"SupportedEthernetCapabilities@odata.count"`
	SupportedLinkCapabilities                 []struct {
		AutoSpeedNegotiation  bool   `json:"AutoSpeedNegotiation"`
		LinkNetworkTechnology string `json:"LinkNetworkTechnology"`
		LinkSpeedMbps         int64  `json:"LinkSpeedMbps"`
	} `json:"SupportedLinkCapabilities"`
	SupportedLinkCapabilities_odata_count int64  `json:"SupportedLinkCapabilities@odata.count"`
	VendorID                              string `json:"VendorId"`
	WakeOnLANEnabled                      bool   `json:"WakeOnLANEnabled"`
}

// NetworkDeviceDell ...
type NetworkDeviceDell struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Controllers    []struct {
		ControllerCapabilities struct {
			DataCenterBridging struct {
				Capable bool `json:"Capable"`
			} `json:"DataCenterBridging"`
			NPIV struct {
				MaxDeviceLogins int `json:"MaxDeviceLogins"`
				MaxPortLogins   int `json:"MaxPortLogins"`
			} `json:"NPIV"`
			NetworkDeviceFunctionCount int `json:"NetworkDeviceFunctionCount"`
			NetworkPortCount           int `json:"NetworkPortCount"`
			VirtualizationOffload      struct {
				SRIOV struct {
					SRIOVVEPACapable bool `json:"SRIOVVEPACapable"`
				} `json:"SRIOV"`
				VirtualFunction struct {
					DeviceMaxCount         int `json:"DeviceMaxCount"`
					MinAssignmentGroupSize int `json:"MinAssignmentGroupSize"`
					NetworkPortMaxCount    int `json:"NetworkPortMaxCount"`
				} `json:"VirtualFunction"`
			} `json:"VirtualizationOffload"`
		} `json:"ControllerCapabilities"`
		FirmwarePackageVersion string `json:"FirmwarePackageVersion"`
		Links                  struct {
			NetworkDeviceFunctions []struct {
				OdataId string `json:"@odata.id"`
			} `json:"NetworkDeviceFunctions"`
			NetworkDeviceFunctions_odata_count int `json:"NetworkDeviceFunctions@odata.count"`
			NetworkPorts                       []struct {
				OdataId string `json:"@odata.id"`
			} `json:"NetworkPorts"`
			NetworkPorts_odata_count int `json:"NetworkPorts@odata.count"`
		} `json:"Links"`
	} `json:"Controllers"`
	Description            string `json:"Description"`
	ID                     string `json:"Id"`
	Manufacturer           string `json:"Manufacturer"`
	Model                  string `json:"Model"`
	Name                   string `json:"Name"`
	NetworkDeviceFunctions struct {
		OdataId string `json:"@odata.id"`
	} `json:"NetworkDeviceFunctions"`
	NetworkPorts struct {
		OdataId string `json:"@odata.id"`
	} `json:"NetworkPorts"`
	PartNumber   string `json:"PartNumber"`
	SerialNumber string `json:"SerialNumber"`
	Status       struct {
		Health       interface{} `json:"Health"`
		HealthRollup interface{} `json:"HealthRollup"`
		State        string      `json:"State"`
	} `json:"Status"`
}

//SystemViewDell ... Fetch the System View Attributes from the Redfish API
type SystemViewDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Actions        struct {
		_ComputerSystem_Reset struct {
			ResetType_Redfish_AllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
			Target                            string   `json:"target"`
		}
	} `json:"Actions"`
	AssetTag string `json:"AssetTag"`
	Bios     struct {
		_odata_id string
	} `json:"Bios"`
	BiosVersion string `json:"BiosVersion"`
	Boot        struct {
		BootOptions struct {
			_odata_id string
		} `json:"BootOptions"`
		BootOrder                                        []string `json:"BootOrder"`
		BootOrder_odata_count                            int      `json:"BootOrder@odata.count"`
		BootSourceOverrideEnabled                        string   `json:"BootSourceOverrideEnabled"`
		BootSourceOverrideMode                           string   `json:"BootSourceOverrideMode"`
		BootSourceOverrideTarget                         string   `json:"BootSourceOverrideTarget"`
		BootSourceOverrideTarget_Redfish_AllowableValues []string `json:"BootSourceOverrideTarget@Redfish.AllowableValues"`
		UefiTargetBootSourceOverride                     string   `json:"UefiTargetBootSourceOverride"`
	} `json:"Boot"`
	Description        string `json:"Description"`
	EthernetInterfaces struct {
		_odata_id string
	} `json:"EthernetInterfaces"`
	HostName          string `json:"HostName"`
	HostWatchdogTimer struct {
		FunctionEnabled bool `json:"FunctionEnabled"`
		Status          struct {
			State string `json:"State"`
		} `json:"Status"`
		TimeoutAction string `json:"TimeoutAction"`
	} `json:"HostWatchdogTimer"`
	HostingRoles             []interface{} `json:"HostingRoles"`
	HostingRoles_odata_count int           `json:"HostingRoles@odata.count"`
	ID                       string        `json:"Id"`
	IndicatorLED             string        `json:"IndicatorLED"`
	Links                    struct {
		Chassis []struct {
			_odata_id string
		} `json:"Chassis"`
		Chassis_odata_count  int           `json:"Chassis@odata.count"`
		CooledBy             []interface{} `json:"CooledBy"`
		CooledBy_odata_count int           `json:"CooledBy@odata.count"`
		ManagedBy            []struct {
			_odata_id string
		} `json:"ManagedBy"`
		ManagedBy_odata_count int `json:"ManagedBy@odata.count"`
		Oem                   struct {
			Dell struct {
				BootOrder struct {
					_odata_id string
				} `json:"BootOrder"`
				DellNumericSensorCollection struct {
					_odata_id string
				} `json:"DellNumericSensorCollection"`
				DellOSDeploymentService struct {
					_odata_id string
				} `json:"DellOSDeploymentService"`
				DellPresenceAndStatusSensorCollection struct {
					_odata_id string
				} `json:"DellPresenceAndStatusSensorCollection"`
				DellRaidService struct {
					_odata_id string
				} `json:"DellRaidService"`
				DellSensorCollection struct {
					_odata_id string
				} `json:"DellSensorCollection"`
				DellSoftwareInstallationService struct {
					_odata_id string
				} `json:"DellSoftwareInstallationService"`
			} `json:"Dell"`
		} `json:"Oem"`
		PoweredBy []struct {
			_odata_id string
		} `json:"PoweredBy"`
		PoweredBy_odata_count int `json:"PoweredBy@odata.count"`
	} `json:"Links"`
	Manufacturer string `json:"Manufacturer"`
	Memory       struct {
		_odata_id string
	} `json:"Memory"`
	MemorySummary struct {
		MemoryMirroring string `json:"MemoryMirroring"`
		Status          struct {
			Health       interface{} `json:"Health"`
			HealthRollup interface{} `json:"HealthRollup"`
			State        string      `json:"State"`
		} `json:"Status"`
		TotalSystemMemoryGiB float32 `json:"TotalSystemMemoryGiB"`
	} `json:"MemorySummary"`
	Model             string `json:"Model"`
	Name              string `json:"Name"`
	NetworkInterfaces struct {
		_odata_id string
	} `json:"NetworkInterfaces"`
	Oem struct {
		Dell struct {
			DellSystem struct {
				_odata_context        string
				_odata_id             string
				_odata_type           string
				BIOSReleaseDate       string      `json:"BIOSReleaseDate"`
				BaseBoardChassisSlot  string      `json:"BaseBoardChassisSlot"`
				BatteryRollupStatus   string      `json:"BatteryRollupStatus"`
				BladeGeometry         string      `json:"BladeGeometry"`
				CMCIP                 interface{} `json:"CMCIP"`
				CPURollupStatus       string      `json:"CPURollupStatus"`
				ChassisServiceTag     string      `json:"ChassisServiceTag"`
				ExpressServiceCode    string      `json:"ExpressServiceCode"`
				FanRollupStatus       string      `json:"FanRollupStatus"`
				IntrusionRollupStatus string      `json:"IntrusionRollupStatus"`
				LicensingRollupStatus string      `json:"LicensingRollupStatus"`
				MaxDIMMSlots          int         `json:"MaxDIMMSlots"`
				MaxPCIeSlots          int         `json:"MaxPCIeSlots"`
				NodeID                string      `json:"NodeID"`
				PSRollupStatus        string      `json:"PSRollupStatus"`
				PowerCapEnabledState  string      `json:"PowerCapEnabledState"`
				StorageRollupStatus   string      `json:"StorageRollupStatus"`
				SysMemPrimaryStatus   string      `json:"SysMemPrimaryStatus"`
				SystemGeneration      string      `json:"SystemGeneration"`
				SystemID              int         `json:"SystemID"`
				TempRollupStatus      string      `json:"TempRollupStatus"`
				UUID                  string      `json:"UUID"`
				VoltRollupStatus      string      `json:"VoltRollupStatus"`
			} `json:"DellSystem"`
		} `json:"Dell"`
	} `json:"Oem"`
	PCIeDevices []struct {
		_odata_id string
	} `json:"PCIeDevices"`
	PCIeDevices_odata_count int `json:"PCIeDevices@odata.count"`
	PCIeFunctions           []struct {
		_odata_id string
	} `json:"PCIeFunctions"`
	PCIeFunctions_odata_count int    `json:"PCIeFunctions@odata.count"`
	PartNumber                string `json:"PartNumber"`
	PowerState                string `json:"PowerState"`
	ProcessorSummary          struct {
		Count                 int    `json:"Count"`
		LogicalProcessorCount int    `json:"LogicalProcessorCount"`
		Model                 string `json:"Model"`
		Status                struct {
			Health       interface{} `json:"Health"`
			HealthRollup interface{} `json:"HealthRollup"`
			State        string      `json:"State"`
		} `json:"Status"`
	} `json:"ProcessorSummary"`
	Processors struct {
		_odata_id string
	} `json:"Processors"`
	SKU        string `json:"SKU"`
	SecureBoot struct {
		_odata_id string
	} `json:"SecureBoot"`
	SerialNumber  string `json:"SerialNumber"`
	SimpleStorage struct {
		_odata_id string
	} `json:"SimpleStorage"`
	Status struct {
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
		State        string `json:"State"`
	} `json:"Status"`
	Storage struct {
		_odata_id string
	} `json:"Storage"`
	SystemType     string `json:"SystemType"`
	TrustedModules []struct {
		InterfaceType string `json:"InterfaceType"`
		Status        struct {
			State string `json:"State"`
		} `json:"Status"`
	} `json:"TrustedModules"`
	UUID string `json:"UUID"`
}

//BootOrderDell ... Fetch the Boot Order Info from the Refish API
type BootOrderDell struct {
	_Redfish_Settings struct {
		_odata_context string
		_odata_type    string
		SettingsObject struct {
			_odata_id string
		} `json:"SettingsObject"`
		SupportedApplyTimes []string `json:"SupportedApplyTimes"`
	}
	_odata_context    string
	_odata_id         string
	_odata_type       string
	AttributeRegistry string `json:"AttributeRegistry"`
	Attributes        struct {
		BootSeq []struct {
			Enabled bool   `json:"Enabled"`
			ID      string `json:"Id"`
			Index   int    `json:"Index"`
			Name    string `json:"Name"`
		} `json:"BootSeq"`
	} `json:"Attributes"`
	Description string `json:"Description"`
	ID          string `json:"Id"`
	Name        string `json:"Name"`
}

//FirmwareDataDell ...
type FirmwareDataDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	ID             string `json:"Id"`
	Name           string `json:"Name"`
	Oem            struct {
		Dell struct {
			DellSoftwareInventory struct {
				_odata_context   string
				_odata_id        string
				_odata_type      string
				ComponentID      string      `json:"ComponentID"`
				ComponentType    string      `json:"ComponentType"`
				DeviceID         string      `json:"DeviceID"`
				ElementName      string      `json:"ElementName"`
				HashValue        interface{} `json:"HashValue"`
				InstallationDate string      `json:"InstallationDate"`
				IsEntity         bool        `json:"IsEntity"`
				RevisionNumber   int         `json:"RevisionNumber"`
				Status           string      `json:"Status"`
				SubDeviceID      string      `json:"SubDeviceID"`
				SubVendorID      string      `json:"SubVendorID"`
				VendorID         string      `json:"VendorID"`
			} `json:"DellSoftwareInventory"`
		} `json:"Dell"`
	} `json:"Oem"`
	SoftwareID string `json:"SoftwareId"`
	Status     struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	Updateable bool   `json:"Updateable"`
	Version    string `json:"Version"`
}

//PowerDataDell ...
type PowerDataDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	ID             string `json:"Id"`
	Name           string `json:"Name"`
	PowerControl   []struct {
		_odata_context      string
		_odata_id           string
		_odata_type         string
		MemberID            string `json:"MemberId"`
		Name                string `json:"Name"`
		PowerAllocatedWatts int    `json:"PowerAllocatedWatts"`
		PowerAvailableWatts int    `json:"PowerAvailableWatts"`
		PowerCapacityWatts  int    `json:"PowerCapacityWatts"`
		PowerConsumedWatts  int    `json:"PowerConsumedWatts"`
		PowerLimit          struct {
			CorrectionInMs int    `json:"CorrectionInMs"`
			LimitException string `json:"LimitException"`
			LimitInWatts   int    `json:"LimitInWatts"`
		} `json:"PowerLimit"`
		PowerMetrics struct {
			AverageConsumedWatts int `json:"AverageConsumedWatts"`
			IntervalInMin        int `json:"IntervalInMin"`
			MaxConsumedWatts     int `json:"MaxConsumedWatts"`
			MinConsumedWatts     int `json:"MinConsumedWatts"`
		} `json:"PowerMetrics"`
		PowerRequestedWatts int `json:"PowerRequestedWatts"`
		RelatedItem         []struct {
			_odata_id string
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
	} `json:"PowerControl"`
	PowerControl_odata_count int `json:"PowerControl@odata.count"`
	PowerSupplies            []struct {
		_odata_context string
		_odata_id      string
		_odata_type    string
		Assembly       struct {
			_odata_id string
		} `json:"Assembly"`
		EfficiencyPercent float64 `json:"EfficiencyPercent"`
		FirmwareVersion   string  `json:"FirmwareVersion"`
		HotPluggable      bool    `json:"HotPluggable"`
		InputRanges       []struct {
			InputType          string `json:"InputType"`
			MaximumFrequencyHz int    `json:"MaximumFrequencyHz"`
			MaximumVoltage     int    `json:"MaximumVoltage"`
			MinimumFrequencyHz int    `json:"MinimumFrequencyHz"`
			MinimumVoltage     int    `json:"MinimumVoltage"`
			OutputWattage      int    `json:"OutputWattage"`
		} `json:"InputRanges"`
		InputRanges_odata_count int         `json:"InputRanges@odata.count"`
		LastPowerOutputWatts    interface{} `json:"LastPowerOutputWatts"`
		LineInputVoltage        int         `json:"LineInputVoltage"`
		LineInputVoltageType    string      `json:"LineInputVoltageType"`
		Manufacturer            string      `json:"Manufacturer"`
		MemberID                string      `json:"MemberId"`
		Model                   string      `json:"Model"`
		Name                    string      `json:"Name"`
		Oem                     struct {
			Dell struct {
				DellPowerSupply struct {
					_odata_context    string
					_odata_id         string
					_odata_type       string
					IsSwitchingSupply bool `json:"IsSwitchingSupply"`
					Links             struct {
						DellPSNumericSensorCollection []struct {
							_odata_id string
						} `json:"DellPSNumericSensorCollection"`
					} `json:"Links"`
				} `json:"DellPowerSupply"`
				DellPowerSupplyView struct {
					_odata_context           string
					_odata_id                string
					_odata_type              string
					DetailedState            string `json:"DetailedState"`
					Range1MaxInputPowerWatts int    `json:"Range1MaxInputPowerWatts"`
				} `json:"DellPowerSupplyView"`
			} `json:"Dell"`
		} `json:"Oem"`
		PartNumber         string `json:"PartNumber"`
		PowerCapacityWatts int    `json:"PowerCapacityWatts"`
		PowerInputWatts    int    `json:"PowerInputWatts"`
		PowerOutputWatts   int    `json:"PowerOutputWatts"`
		PowerSupplyType    string `json:"PowerSupplyType"`
		Redundancy         []struct {
			_odata_context  string
			_odata_id       string
			_odata_type     string
			MaxNumSupported int    `json:"MaxNumSupported"`
			MemberID        string `json:"MemberId"`
			MinNumNeeded    int    `json:"MinNumNeeded"`
			Mode            string `json:"Mode"`
			Name            string `json:"Name"`
			RedundancySet   []struct {
				_odata_id string
			} `json:"RedundancySet"`
			RedundancySet_odata_count int `json:"RedundancySet@odata.count"`
			Status                    struct {
				Health string `json:"Health"`
				State  string `json:"State"`
			} `json:"Status"`
		} `json:"Redundancy"`
		Redundancycount int `json:"Redundancy@odata.count"`
		RelatedItem     []struct {
			_odata_id string
		} `json:"RelatedItem"`
		RelatedItem_odata_count int    `json:"RelatedItem@odata.count"`
		SerialNumber            string `json:"SerialNumber"`
		SparePartNumber         string `json:"SparePartNumber"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"PowerSupplies"`
	PowerSuppliescount int `json:"PowerSupplies@odata.count"`
	Redundancy         []struct {
		_odata_context  string
		_odata_id       string
		_odata_type     string
		MaxNumSupported int    `json:"MaxNumSupported"`
		MemberID        string `json:"MemberId"`
		MinNumNeeded    int    `json:"MinNumNeeded"`
		Mode            string `json:"Mode"`
		Name            string `json:"Name"`
		RedundancySet   []struct {
			_odata_id string
		} `json:"RedundancySet"`
		Redundancycount int `json:"RedundancySet@odata.count"`
		Status          struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"Redundancy"`
	Redundancycount int `json:"Redundancy@odata.count"`
	Voltages        []struct {
		_odata_context            string
		_odata_id                 string
		_odata_type               string
		LowerThresholdCritical    interface{} `json:"LowerThresholdCritical"`
		LowerThresholdFatal       interface{} `json:"LowerThresholdFatal"`
		LowerThresholdNonCritical interface{} `json:"LowerThresholdNonCritical"`
		MaxReadingRange           int         `json:"MaxReadingRange"`
		MemberID                  string      `json:"MemberId"`
		MinReadingRange           int         `json:"MinReadingRange"`
		Name                      string      `json:"Name"`
		PhysicalContext           string      `json:"PhysicalContext"`
		ReadingVolts              int         `json:"ReadingVolts"`
		RelatedItem               []struct {
			_odata_id string
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
		SensorNumber            int `json:"SensorNumber"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		UpperThresholdCritical    interface{} `json:"UpperThresholdCritical"`
		UpperThresholdFatal       interface{} `json:"UpperThresholdFatal"`
		UpperThresholdNonCritical interface{} `json:"UpperThresholdNonCritical"`
	} `json:"Voltages"`
	Voltagescount int `json:"Voltages@odata.count"`
}

//AccountsInfoDell ...
type AccountsInfoDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	Enabled        bool   `json:"Enabled"`
	ID             string `json:"Id"`
	Links          struct {
		Role struct {
			_odata_id string
		} `json:"Role"`
	} `json:"Links"`
	Locked   bool        `json:"Locked"`
	Name     string      `json:"Name"`
	Password interface{} `json:"Password"`
	RoleID   string      `json:"RoleId"`
	UserName string      `json:"UserName"`
}

//SystemEventLogsV1Dell ...
type SystemEventLogsV1Dell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	Members        []struct {
		_odata_id   string
		_odata_type string
		Created     string `json:"Created"`
		Description string `json:"Description"`
		EntryCode   []struct {
			Member string `json:"Member"`
		} `json:"EntryCode"`
		EntryType               string        `json:"EntryType"`
		ID                      string        `json:"Id"`
		Links                   struct{}      `json:"Links"`
		Message                 string        `json:"Message"`
		MessageArgs             []interface{} `json:"MessageArgs"`
		MessageArgs_odata_count int           `json:"MessageArgs@odata.count"`
		MessageID               string        `json:"MessageID"`
		Name                    string        `json:"Name"`
		SensorNumber            int           `json:"SensorNumber"`
		SensorType              []struct {
			Member string `json:"Member"`
		} `json:"SensorType"`
		Severity string `json:"Severity"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}

//SystemEventLogsV2Dell ...
type SystemEventLogsV2Dell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	Members        []struct {
		_odata_id               string
		_odata_type             string
		Created                 string        `json:"Created"`
		Description             string        `json:"Description"`
		EntryCode               string        `json:"EntryCode"`
		EntryType               string        `json:"EntryType"`
		ID                      string        `json:"Id"`
		Links                   struct{}      `json:"Links"`
		Message                 string        `json:"Message"`
		MessageArgs             []interface{} `json:"MessageArgs"`
		MessageArgs_odata_count int           `json:"MessageArgs@odata.count"`
		MessageID               string        `json:"MessageId"`
		Name                    string        `json:"Name"`
		SensorNumber            int           `json:"SensorNumber"`
		SensorType              string        `json:"SensorType"`
		Severity                string        `json:"Severity"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}

//LifeCycleLogsV1Dell ...
type LifeCycleLogsV1Dell struct {
	Odata_context string
	Odata_id      string
	Odata_type    string
	Description   string `json:"Description"`
	Members       []struct {
		Odata_id    string
		Odata_type  string
		Created     string `json:"Created"`
		Description string `json:"Description"`
		EntryType   string `json:"EntryType"`
		ID          string `json:"Id"`
		Links       struct {
			OriginOfCondition struct {
				Odata_id string `json:"@odata.id"`
			} `json:"OriginOfCondition"`
		} `json:"Links"`
		Message                 string   `json:"Message"`
		MessageArgs             []string `json:"MessageArgs"`
		MessageArgs_odata_count int64    `json:"MessageArgs@odata.count"`
		MessageID               string   `json:"MessageId"`
		Name                    string   `json:"Name"`
		OemRecordFormat         string   `json:"OemRecordFormat"`
		Severity                string   `json:"Severity"`
	} `json:"Members"`
}

//ThermalHealthListDell ...
type ThermalHealthListDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	Fans           []struct {
		_odata_id   string
		_odata_type string
		Assembly    struct {
			_odata_id string
		} `json:"Assembly"`
		FanName                   string        `json:"FanName"`
		LowerThresholdCritical    int           `json:"LowerThresholdCritical"`
		LowerThresholdFatal       int           `json:"LowerThresholdFatal"`
		LowerThresholdNonCritical int           `json:"LowerThresholdNonCritical"`
		MaxReadingRange           interface{}   `json:"MaxReadingRange"`
		MemberID                  string        `json:"MemberId"`
		MinReadingRange           int           `json:"MinReadingRange"`
		Name                      string        `json:"Name"`
		PhysicalContext           string        `json:"PhysicalContext"`
		Reading                   int           `json:"Reading"`
		ReadingUnits              string        `json:"ReadingUnits"`
		Redundancy                []interface{} `json:"Redundancy"`
		Redundancy_odata_count    int           `json:"Redundancy@odata.count"`
		RelatedItem               []struct {
			_odata_id string
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		UpperThresholdCritical    interface{} `json:"UpperThresholdCritical"`
		UpperThresholdFatal       interface{} `json:"UpperThresholdFatal"`
		UpperThresholdNonCritical interface{} `json:"UpperThresholdNonCritical"`
	} `json:"Fans"`
	Fanscount  int    `json:"Fans@odata.count"`
	ID         string `json:"Id"`
	Name       string `json:"Name"`
	Redundancy []struct {
		_odata_id                 string
		_odata_type               string
		MaxNumSupported           int           `json:"MaxNumSupported"`
		MemberID                  string        `json:"MemberId"`
		MinNumNeeded              int           `json:"MinNumNeeded"`
		Mode                      string        `json:"Mode"`
		Name                      string        `json:"Name"`
		RedundancyEnabled         bool          `json:"RedundancyEnabled"`
		RedundancySet             []interface{} `json:"RedundancySet"`
		RedundancySet_odata_count int           `json:"RedundancySet@odata.count"`
		Status                    struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"Redundancy"`
	Redundancycount int `json:"Redundancy@odata.count"`
	Temperatures    []struct {
		_odata_context            string
		_odata_id                 string
		_odata_type               string
		LowerThresholdCritical    int         `json:"LowerThresholdCritical"`
		LowerThresholdFatal       int         `json:"LowerThresholdFatal"`
		LowerThresholdNonCritical interface{} `json:"LowerThresholdNonCritical"`
		MaxReadingRangeTemp       int         `json:"MaxReadingRangeTemp"`
		MemberID                  string      `json:"MemberId"`
		MinReadingRangeTemp       int         `json:"MinReadingRangeTemp"`
		Name                      string      `json:"Name"`
		PhysicalContext           string      `json:"PhysicalContext"`
		ReadingCelsius            int         `json:"ReadingCelsius"`
		RelatedItem               []struct {
			_odata_id string
		} `json:"RelatedItem"`
		RelatedItem_odata_count int `json:"RelatedItem@odata.count"`
		SensorNumber            int `json:"SensorNumber"`
		Status                  struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		UpperThresholdCritical    int         `json:"UpperThresholdCritical"`
		UpperThresholdFatal       int         `json:"UpperThresholdFatal"`
		UpperThresholdNonCritical interface{} `json:"UpperThresholdNonCritical"`
	} `json:"Temperatures"`
	Temperaturescount int `json:"Temperatures@odata.count"`
}

//MemberCountDell ...
type MemberCountDell struct {
	OdataContext        string    `json:"@odata.context"`
	OdataId             string    `json:"@odata.id"`
	OdataType           string    `json:"@odata.type"`
	Description         string    `json:"Description"`
	Members             []Members `json:"Members"`
	Members_odata_count int       `json:"Members@odata.count"`
	Name                string    `json:"Name"`
}

type Members struct {
	OdataId string `json:"@odata.id"`
}

//ProcessorDataDell ...
type ProcessorDataDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Assembly       struct {
		_odata_id string
	} `json:"Assembly"`
	Description    string `json:"Description"`
	ID             string `json:"Id"`
	InstructionSet string `json:"InstructionSet"`
	Links          struct {
		Chassis struct {
			_odata_id string
		} `json:"Chassis"`
	} `json:"Links"`
	Manufacturer string `json:"Manufacturer"`
	MaxSpeedMHz  int    `json:"MaxSpeedMHz"`
	Model        string `json:"Model"`
	Name         string `json:"Name"`
	Oem          struct {
		Dell struct {
			DellProcessor struct {
				_odata_context                  string
				_odata_id                       string
				_odata_type                     string
				CurrentClockSpeedMhz            int    `json:"CurrentClockSpeedMhz"`
				HyperThreadingCapable           string `json:"HyperThreadingCapable"`
				HyperThreadingEnabled           string `json:"HyperThreadingEnabled"`
				TurboModeCapable                string `json:"TurboModeCapable"`
				TurboModeEnabled                string `json:"TurboModeEnabled"`
				VirtualizationTechnologyCapable string `json:"VirtualizationTechnologyCapable"`
				VirtualizationTechnologyEnabled string `json:"VirtualizationTechnologyEnabled"`
			} `json:"DellProcessor"`
		} `json:"Dell"`
	} `json:"Oem"`
	ProcessorArchitecture string `json:"ProcessorArchitecture"`
	ProcessorID           struct {
		EffectiveFamily         string      `json:"EffectiveFamily"`
		EffectiveModel          string      `json:"EffectiveModel"`
		IdentificationRegisters string      `json:"IdentificationRegisters"`
		MicrocodeInfo           interface{} `json:"MicrocodeInfo"`
		Step                    string      `json:"Step"`
		VendorID                string      `json:"VendorId"`
	} `json:"ProcessorId"`
	ProcessorType string `json:"ProcessorType"`
	Socket        string `json:"Socket"`
	Status        struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	TotalCores   int `json:"TotalCores"`
	TotalThreads int `json:"TotalThreads"`
}

//ProcessorsListDataDell ...
type ProcessorsListDataDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	Members        []struct {
		OdataId string `json:"@odata.id"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}

//StorageCollectionDell ...
type StorageCollectionDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	Members        []struct {
		OdataId string `json:"@odata.id"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}

type StorageDetailsDell struct {
	_odata_context string
	_odata_id      string
	_odata_type    string
	Description    string `json:"Description"`
	Drives         []struct {
		OdataId string `json:"@odata.id"`
	} `json:"Drives"`
	Drivescount int    `json:"Drives@odata.count"`
	ID          string `json:"Id"`
	Links       struct {
		Enclosures []struct {
			_odata_id string
		} `json:"Enclosures"`
		Enclosures_odata_count int `json:"Enclosures@odata.count"`
	} `json:"Links"`
	Name string `json:"Name"`
	Oem  struct {
		Dell struct {
			DellController struct {
				_odata_context            string
				_odata_id                 string
				_odata_type               string
				CacheSizeInMB             int         `json:"CacheSizeInMB"`
				CachecadeCapability       string      `json:"CachecadeCapability"`
				ControllerFirmwareVersion string      `json:"ControllerFirmwareVersion"`
				DeviceCardSlotType        string      `json:"DeviceCardSlotType"`
				DriverVersion             interface{} `json:"DriverVersion"`
				EncryptionCapability      string      `json:"EncryptionCapability"`
				EncryptionMode            string      `json:"EncryptionMode"`
				PCISlot                   int         `json:"PCISlot"`
				PatrolReadState           string      `json:"PatrolReadState"`
				RollupStatus              string      `json:"RollupStatus"`
				SecurityStatus            string      `json:"SecurityStatus"`
				SlicedVDCapability        string      `json:"SlicedVDCapability"`
			} `json:"DellController"`
		} `json:"Dell"`
	} `json:"Oem"`
	Status struct {
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
		State        string `json:"State"`
	} `json:"Status"`
	StorageControllers []struct {
		_odata_id   string
		_odata_type string
		Assembly    struct {
			_odata_id string
		} `json:"Assembly"`
		FirmwareVersion string `json:"FirmwareVersion"`
		Identifiers     []struct {
			DurableName       string `json:"DurableName"`
			DurableNameFormat string `json:"DurableNameFormat"`
		} `json:"Identifiers"`
		Links        struct{} `json:"Links"`
		Manufacturer string   `json:"Manufacturer"`
		MemberID     string   `json:"MemberId"`
		Model        string   `json:"Model"`
		Name         string   `json:"Name"`
		SpeedGbps    int      `json:"SpeedGbps"`
		Status       struct {
			Health       string `json:"Health"`
			HealthRollup string `json:"HealthRollup"`
			State        string `json:"State"`
		} `json:"Status"`
		SupportedControllerProtocols []string `json:"SupportedControllerProtocols"`
		SupportedDeviceProtocols     []string `json:"SupportedDeviceProtocols"`
	} `json:"StorageControllers"`
	StorageControllers_odata_count int `json:"StorageControllers@odata.count"`
	Volumes                        struct {
		_odata_id string
	} `json:"Volumes"`
}

type StorageDriveDetailsDell struct {
	Actions struct {
		Drive_SecureErase struct {
			Target string `json:"target"`
		}
	} `json:"Actions"`
	Assembly struct {
		_odata_id string
	} `json:"Assembly"`
	BlockSizeBytes    int    `json:"BlockSizeBytes"`
	CapableSpeedGbs   int    `json:"CapableSpeedGbs"`
	CapacityBytes     int    `json:"CapacityBytes"`
	Description       string `json:"Description"`
	EncryptionAbility string `json:"EncryptionAbility"`
	EncryptionStatus  string `json:"EncryptionStatus"`
	FailurePredicted  bool   `json:"FailurePredicted"`
	HotspareType      string `json:"HotspareType"`
	ID                string `json:"Id"`
	Identifiers       []struct {
		DurableName       string `json:"DurableName"`
		DurableNameFormat string `json:"DurableNameFormat"`
	} `json:"Identifiers"`
	Links struct {
		Chassis struct {
			_odata_id string
		} `json:"Chassis"`
		Volumes             []interface{} `json:"Volumes"`
		Volumes_odata_count int           `json:"Volumes@odata.count"`
	} `json:"Links"`
	Location           []interface{} `json:"Location"`
	Manufacturer       string        `json:"Manufacturer"`
	MediaType          string        `json:"MediaType"`
	Model              string        `json:"Model"`
	Name               string        `json:"Name"`
	NegotiatedSpeedGbs int           `json:"NegotiatedSpeedGbs"`
	Oem                struct {
		Dell struct {
			DellPhysicalDisk struct {
				_odata_context         string
				_odata_id              string
				_odata_type            string
				Connector              int    `json:"Connector"`
				DriveFormFactor        string `json:"DriveFormFactor"`
				FreeSizeInBytes        int    `json:"FreeSizeInBytes"`
				ManufacturingDay       int    `json:"ManufacturingDay"`
				ManufacturingWeek      int    `json:"ManufacturingWeek"`
				ManufacturingYear      int    `json:"ManufacturingYear"`
				PPID                   string `json:"PPID"`
				PredictiveFailureState string `json:"PredictiveFailureState"`
				RaidStatus             string `json:"RaidStatus"`
				SASAddress             string `json:"SASAddress"`
				Slot                   int    `json:"Slot"`
				UsedSizeInBytes        int    `json:"UsedSizeInBytes"`
			} `json:"DellPhysicalDisk"`
		} `json:"Dell"`
	} `json:"Oem"`
	Operations                    []interface{} `json:"Operations"`
	PartNumber                    string        `json:"PartNumber"`
	PredictedMediaLifeLeftPercent interface{}   `json:"PredictedMediaLifeLeftPercent"`
	Protocol                      string        `json:"Protocol"`
	Revision                      string        `json:"Revision"`
	RotationSpeedRPM              int           `json:"RotationSpeedRPM"`
	SerialNumber                  string        `json:"SerialNumber"`
	Status                        struct {
		Health       string `json:"Health"`
		HealthRollup string `json:"HealthRollup"`
		State        string `json:"State"`
	} `json:"Status"`
}

//HP Based Structs

//FirmwareInventoryHP ...
type FirmwareInventoryHP struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Current      struct {
		One03c3239103c21c0 []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"103c3239103c21c0"`
		One4e41657103c22be []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"14e41657103c22be"`
		Eight08610fb103c17d0 []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			ResetRequired   bool     `json:"ResetRequired"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"808610fb103c17d0"`
		Eight08610fb103c17d3 []struct {
			ImageSizeBytes  int      `json:"ImageSizeBytes"`
			Key             string   `json:"Key"`
			Location        string   `json:"Location"`
			Name            string   `json:"Name"`
			ResetRequired   bool     `json:"ResetRequired"`
			UEFIDevicePaths []string `json:"UEFIDevicePaths"`
			Updateable      bool     `json:"Updateable"`
			VersionString   string   `json:"VersionString"`
		} `json:"808610fb103c17d3"`
		IntelligentProvisioning []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"IntelligentProvisioning"`
		PlatformDefinitionTable []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PlatformDefinitionTable"`
		PowerManagementController []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PowerManagementController"`
		PowerManagementControllerBootloader []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PowerManagementControllerBootloader"`
		PowerSupplies []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"PowerSupplies"`
		SASProgrammableLogicDevice []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SASProgrammableLogicDevice"`
		SPSFirmwareVersionData []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SPSFirmwareVersionData"`
		StorageBattery []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"StorageBattery"`
		SystemBMC []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemBMC"`
		SystemProgrammableLogicDevice []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemProgrammableLogicDevice"`
		SystemRomActive []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemRomActive"`
		SystemRomBackup []struct {
			Key           string `json:"Key"`
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"SystemRomBackup"`
		TrustedPlatformModule12 []struct {
			Location      string `json:"Location"`
			Name          string `json:"Name"`
			VersionString string `json:"VersionString"`
		} `json:"TrustedPlatformModule12"`
	} `json:"Current"`
	ID    string `json:"Id"`
	Name  string `json:"Name"`
	Type  string `json:"Type"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//SystemInfoHp is a struct which fetches the Overall System High Level info and its a Singleton Resource
type SystemInfoHP struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Actions      struct {
		ComputerSystemReset struct {
			ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
			Target                          string   `json:"target"`
		} `json:"#ComputerSystem.Reset"`
	} `json:"Actions"`
	AssetTag         string `json:"AssetTag"`
	AvailableActions []struct {
		Action       string `json:"Action"`
		Capabilities []struct {
			AllowableValues []string `json:"AllowableValues"`
			PropertyName    string   `json:"PropertyName"`
		} `json:"Capabilities"`
	} `json:"AvailableActions"`
	Bios struct {
		Current struct {
			VersionString string `json:"VersionString"`
		} `json:"Current"`
	} `json:"Bios"`
	BiosVersion string `json:"BiosVersion"`
	Boot        struct {
		BootSourceOverrideEnabled   string   `json:"BootSourceOverrideEnabled"`
		BootSourceOverrideSupported []string `json:"BootSourceOverrideSupported"`
		BootSourceOverrideTarget    string   `json:"BootSourceOverrideTarget"`
	} `json:"Boot"`
	Description     string `json:"Description"`
	HostCorrelation struct {
		HostMACAddress []string `json:"HostMACAddress"`
		HostName       string   `json:"HostName"`
		IPAddress      []string `json:"IPAddress"`
	} `json:"HostCorrelation"`
	HostName     string `json:"HostName"`
	ID           string `json:"Id"`
	IndicatorLED string `json:"IndicatorLED"`
	LogServices  struct {
		OdataID string `json:"@odata.id"`
	} `json:"LogServices"`
	Manufacturer string `json:"Manufacturer"`
	Memory       struct {
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
		TotalSystemMemoryGB float32 `json:"TotalSystemMemoryGB"`
	} `json:"Memory"`
	MemorySummary struct {
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
		TotalSystemMemoryGiB float32 `json:"TotalSystemMemoryGiB"`
	} `json:"MemorySummary"`
	Model string `json:"Model"`
	Name  string `json:"Name"`
	Oem   struct {
		Hp struct {
			OdataType string `json:"@odata.type"`
			Actions   struct {
				HpComputerSystemExtPowerButton struct {
					PushTypeRedfishAllowableValues []string `json:"PushType@Redfish.AllowableValues"`
					Target                         string   `json:"target"`
				} `json:"#HpComputerSystemExt.PowerButton"`
				HpComputerSystemExtSystemReset struct {
					ResetTypeRedfishAllowableValues []string `json:"ResetType@Redfish.AllowableValues"`
					Target                          string   `json:"target"`
				} `json:"#HpComputerSystemExt.SystemReset"`
			} `json:"Actions"`
			AvailableActions []struct {
				Action       string `json:"Action"`
				Capabilities []struct {
					AllowableValues []string `json:"AllowableValues"`
					PropertyName    string   `json:"PropertyName"`
				} `json:"Capabilities"`
			} `json:"AvailableActions"`
			Battery []struct {
				Condition       string `json:"Condition"`
				ErrorCode       int    `json:"ErrorCode"`
				FirmwareVersion string `json:"FirmwareVersion"`
				Index           int    `json:"Index"`
				MaxCapWatts     int    `json:"MaxCapWatts"`
				Model           string `json:"Model"`
				Present         string `json:"Present"`
				ProductName     string `json:"ProductName"`
				SerialNumber    string `json:"SerialNumber"`
				Spare           string `json:"Spare"`
			} `json:"Battery"`
			Bios struct {
				Backup struct {
					Date          string `json:"Date"`
					Family        string `json:"Family"`
					VersionString string `json:"VersionString"`
				} `json:"Backup"`
				Current struct {
					Date          string `json:"Date"`
					Family        string `json:"Family"`
					VersionString string `json:"VersionString"`
				} `json:"Current"`
				UefiClass int `json:"UefiClass"`
			} `json:"Bios"`
			DeviceDiscoveryComplete struct {
				AMSDeviceDiscovery  string `json:"AMSDeviceDiscovery"`
				DeviceDiscovery     string `json:"DeviceDiscovery"`
				SmartArrayDiscovery string `json:"SmartArrayDiscovery"`
			} `json:"DeviceDiscoveryComplete"`
			IntelligentProvisioningIndex    int      `json:"IntelligentProvisioningIndex"`
			IntelligentProvisioningLocation string   `json:"IntelligentProvisioningLocation"`
			IntelligentProvisioningVersion  string   `json:"IntelligentProvisioningVersion"`
			PostState                       string   `json:"PostState"`
			PowerAllocationLimit            int      `json:"PowerAllocationLimit"`
			PowerAutoOn                     string   `json:"PowerAutoOn"`
			PowerOnDelay                    string   `json:"PowerOnDelay"`
			PowerRegulatorMode              string   `json:"PowerRegulatorMode"`
			PowerRegulatorModesSupported    []string `json:"PowerRegulatorModesSupported"`
			TrustedModules                  []struct {
				FWVersion struct {
					Current struct {
						MajorVersion int `json:"MajorVersion"`
						MinorVersion int `json:"MinorVersion"`
					} `json:"Current"`
				} `json:"FWVersion"`
				ModuleType string `json:"ModuleType"`
				Status     string `json:"Status"`
			} `json:"TrustedModules"`
			Type           string `json:"Type"`
			VirtualProfile string `json:"VirtualProfile"`
			Links          struct {
				BIOS struct {
					Href string `json:"href"`
				} `json:"BIOS"`
				EthernetInterfaces struct {
					Href string `json:"href"`
				} `json:"EthernetInterfaces"`
				FirmwareInventory struct {
					Href string `json:"href"`
				} `json:"FirmwareInventory"`
				Memory struct {
					Href string `json:"href"`
				} `json:"Memory"`
				NetworkAdapters struct {
					Href string `json:"href"`
				} `json:"NetworkAdapters"`
				PCIDevices struct {
					Href string `json:"href"`
				} `json:"PCIDevices"`
				PCISlots struct {
					Href string `json:"href"`
				} `json:"PCISlots"`
				SecureBoot struct {
					Href string `json:"href"`
				} `json:"SecureBoot"`
				SmartStorage struct {
					Href string `json:"href"`
				} `json:"SmartStorage"`
				SoftwareInventory struct {
					Href string `json:"href"`
				} `json:"SoftwareInventory"`
			} `json:"links"`
		} `json:"Hp"`
	} `json:"Oem"`
	Power            string `json:"Power"`
	PowerState       string `json:"PowerState"`
	ProcessorSummary struct {
		Count  int    `json:"Count"`
		Model  string `json:"Model"`
		Status struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
	} `json:"ProcessorSummary"`
	Processors struct {
		Count           int    `json:"Count"`
		ProcessorFamily string `json:"ProcessorFamily"`
		Status          struct {
			HealthRollUp string `json:"HealthRollUp"`
		} `json:"Status"`
	} `json:"Processors"`
	SKU          string `json:"SKU"`
	SerialNumber string `json:"SerialNumber"`
	Status       struct {
		Health string `json:"Health"`
		State  string `json:"State"`
	} `json:"Status"`
	SystemType string `json:"SystemType"`
	Type       string `json:"Type"`
	UUID       string `json:"UUID"`
	Links      struct {
		Chassis []struct {
			Href string `json:"href"`
		} `json:"Chassis"`
		Logs struct {
			Href string `json:"href"`
		} `json:"Logs"`
		ManagedBy []struct {
			Href string `json:"href"`
		} `json:"ManagedBy"`
		Processors struct {
			Href string `json:"href"`
		} `json:"Processors"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//MemberCountHP ...
type MemberCountHP struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Description  string `json:"Description"`
	Members      []struct {
		OdataID string
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
}

//ThermalHealthListHP ...
type ThermalHealthListHP struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Fans         []struct {
		CurrentReading int    `json:"CurrentReading"`
		FanName        string `json:"FanName"`
		Oem            struct {
			Hp struct {
				_odata_type string
				Location    string `json:"Location"`
				Type        string `json:"Type"`
			} `json:"Hp"`
		} `json:"Oem"`
		Status struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		Units string `json:"Units"`
	} `json:"Fans"`
	ID           string `json:"Id"`
	Name         string `json:"Name"`
	Temperatures []struct {
		CurrentReading int    `json:"CurrentReading"`
		Name           string `json:"Name"`
		Number         int    `json:"Number"`
		Oem            struct {
			Hp struct {
				_odata_type string
				LocationXmm int    `json:"LocationXmm"`
				LocationYmm int    `json:"LocationYmm"`
				Type        string `json:"Type"`
			} `json:"Hp"`
		} `json:"Oem"`
		PhysicalContext string `json:"PhysicalContext"`
		ReadingCelsius  int    `json:"ReadingCelsius"`
		Status          struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		Units                  string `json:"Units"`
		UpperThresholdCritical int    `json:"UpperThresholdCritical"`
		UpperThresholdFatal    int    `json:"UpperThresholdFatal"`
	} `json:"Temperatures"`
	Type  string `json:"Type"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//PowerDataHP ...
type PowerDataHP struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	ID           string `json:"Id"`
	Name         string `json:"Name"`
	Oem          struct {
		Hp struct {
			_odata_type             string
			SNMPPowerThresholdAlert struct {
				DurationInMin  int    `json:"DurationInMin"`
				ThresholdWatts int    `json:"ThresholdWatts"`
				Trigger        string `json:"Trigger"`
			} `json:"SNMPPowerThresholdAlert"`
			Type  string `json:"Type"`
			Links struct {
				FastPowerMeter struct {
					Href string `json:"href"`
				} `json:"FastPowerMeter"`
				FederatedGroupCapping struct {
					Href string `json:"href"`
				} `json:"FederatedGroupCapping"`
				PowerMeter struct {
					Href string `json:"href"`
				} `json:"PowerMeter"`
			} `json:"links"`
		} `json:"Hp"`
	} `json:"Oem"`
	PowerCapacityWatts int `json:"PowerCapacityWatts"`
	PowerConsumedWatts int `json:"PowerConsumedWatts"`
	PowerControl       []struct {
		PowerCapacityWatts int `json:"PowerCapacityWatts"`
		PowerConsumedWatts int `json:"PowerConsumedWatts"`
		PowerLimit         struct {
			LimitInWatts interface{} `json:"LimitInWatts"`
		} `json:"PowerLimit"`
		PowerMetrics struct {
			AverageConsumedWatts int `json:"AverageConsumedWatts"`
			IntervalInMin        int `json:"IntervalInMin"`
			MaxConsumedWatts     int `json:"MaxConsumedWatts"`
			MinConsumedWatts     int `json:"MinConsumedWatts"`
		} `json:"PowerMetrics"`
	} `json:"PowerControl"`
	PowerLimit struct {
		LimitInWatts interface{} `json:"LimitInWatts"`
	} `json:"PowerLimit"`
	PowerMetrics struct {
		AverageConsumedWatts int `json:"AverageConsumedWatts"`
		IntervalInMin        int `json:"IntervalInMin"`
		MaxConsumedWatts     int `json:"MaxConsumedWatts"`
		MinConsumedWatts     int `json:"MinConsumedWatts"`
	} `json:"PowerMetrics"`
	PowerSupplies []struct {
		FirmwareVersion      string `json:"FirmwareVersion"`
		LastPowerOutputWatts int    `json:"LastPowerOutputWatts"`
		LineInputVoltage     int    `json:"LineInputVoltage"`
		LineInputVoltageType string `json:"LineInputVoltageType"`
		Model                string `json:"Model"`
		Name                 string `json:"Name"`
		Oem                  struct {
			Hp struct {
				_odata_type             string
				AveragePowerOutputWatts int  `json:"AveragePowerOutputWatts"`
				BayNumber               int  `json:"BayNumber"`
				HotplugCapable          bool `json:"HotplugCapable"`
				MaxPowerOutputWatts     int  `json:"MaxPowerOutputWatts"`
				Mismatched              bool `json:"Mismatched"`
				PowerSupplyStatus       struct {
					State string `json:"State"`
				} `json:"PowerSupplyStatus"`
				Type        string `json:"Type"`
				IPDUCapable bool   `json:"iPDUCapable"`
			} `json:"Hp"`
		} `json:"Oem"`
		PowerCapacityWatts int    `json:"PowerCapacityWatts"`
		PowerSupplyType    string `json:"PowerSupplyType"`
		SerialNumber       string `json:"SerialNumber"`
		SparePartNumber    string `json:"SparePartNumber"`
		Status             struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
	} `json:"PowerSupplies"`
	Redundancy []struct {
		MaxNumSupported int    `json:"MaxNumSupported"`
		MemberID        string `json:"MemberId"`
		MinNumNeeded    int    `json:"MinNumNeeded"`
		Mode            string `json:"Mode"`
		Name            string `json:"Name"`
		RedundancySet   []struct {
			_odata_id string
		} `json:"RedundancySet"`
	} `json:"Redundancy"`
	Type  string `json:"Type"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//EthernetInterfacesHealthHP ...
type EthernetInterfacesHP struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Description  string `json:"Description"`
	Items        []struct {
		_odata_context    string
		_odata_id         string
		_odata_type       string
		AutoNeg           bool   `json:"AutoNeg"`
		Autosense         bool   `json:"Autosense"`
		Description       string `json:"Description"`
		FQDN              string `json:"FQDN"`
		FactoryMacAddress string `json:"FactoryMacAddress"`
		FullDuplex        bool   `json:"FullDuplex"`
		HostName          string `json:"HostName"`
		IPv4Addresses     []struct {
			Address       string `json:"Address"`
			AddressOrigin string `json:"AddressOrigin"`
			Gateway       string `json:"Gateway"`
			SubnetMask    string `json:"SubnetMask"`
		} `json:"IPv4Addresses"`
		IPv6AddressPolicyTable []struct {
			Label      interface{} `json:"Label"`
			Precedence int         `json:"Precedence"`
			Prefix     string      `json:"Prefix"`
		} `json:"IPv6AddressPolicyTable"`
		IPv6Addresses []struct {
			Address       string `json:"Address"`
			AddressOrigin string `json:"AddressOrigin"`
			AddressState  string `json:"AddressState"`
			PrefixLength  int    `json:"PrefixLength"`
		} `json:"IPv6Addresses"`
		IPv6DefaultGateway  string `json:"IPv6DefaultGateway"`
		IPv6StaticAddresses []struct {
			Address      string      `json:"Address"`
			PrefixLength interface{} `json:"PrefixLength"`
		} `json:"IPv6StaticAddresses"`
		ID                     string `json:"Id"`
		LinkTechnology         string `json:"LinkTechnology"`
		MacAddress             string `json:"MacAddress"`
		MaxIPv6StaticAddresses int    `json:"MaxIPv6StaticAddresses"`
		Name                   string `json:"Name"`
		Oem                    struct {
			Hp struct {
				_odata_type           string
				ConfigurationSettings string `json:"ConfigurationSettings"`
				DHCPv4                struct {
					Enabled         bool `json:"Enabled"`
					UseDNSServers   bool `json:"UseDNSServers"`
					UseDomainName   bool `json:"UseDomainName"`
					UseGateway      bool `json:"UseGateway"`
					UseNTPServers   bool `json:"UseNTPServers"`
					UseStaticRoutes bool `json:"UseStaticRoutes"`
					UseWINSServers  bool `json:"UseWINSServers"`
				} `json:"DHCPv4"`
				DHCPv6 struct {
					StatefulModeEnabled  bool `json:"StatefulModeEnabled"`
					StatelessModeEnabled bool `json:"StatelessModeEnabled"`
					UseDNSServers        bool `json:"UseDNSServers"`
					UseDomainName        bool `json:"UseDomainName"`
					UseNTPServers        bool `json:"UseNTPServers"`
					UseRapidCommit       bool `json:"UseRapidCommit"`
				} `json:"DHCPv6"`
				DomainName string `json:"DomainName"`
				HostName   string `json:"HostName"`
				IPv4       struct {
					DDNSRegistration bool     `json:"DDNSRegistration"`
					DNSServers       []string `json:"DNSServers"`
					StaticRoutes     []struct {
						Destination string `json:"Destination"`
						Gateway     string `json:"Gateway"`
						SubnetMask  string `json:"SubnetMask"`
					} `json:"StaticRoutes"`
					WINSRegistration bool     `json:"WINSRegistration"`
					WINSServers      []string `json:"WINSServers"`
				} `json:"IPv4"`
				IPv6 struct {
					DDNSRegistration     bool     `json:"DDNSRegistration"`
					DNSServers           []string `json:"DNSServers"`
					SLAACEnabled         bool     `json:"SLAACEnabled"`
					StaticDefaultGateway string   `json:"StaticDefaultGateway"`
					StaticRoutes         []struct {
						Destination  string      `json:"Destination"`
						Gateway      string      `json:"Gateway"`
						PrefixLength interface{} `json:"PrefixLength"`
						Status       string      `json:"Status"`
					} `json:"StaticRoutes"`
				} `json:"IPv6"`
				NICEnabled           bool   `json:"NICEnabled"`
				NICSupportsIPv6      bool   `json:"NICSupportsIPv6"`
				PingGatewayOnStartup bool   `json:"PingGatewayOnStartup"`
				Type                 string `json:"Type"`
				Links                struct {
					DateTimeService struct {
						Href string `json:"href"`
					} `json:"DateTimeService"`
				} `json:"links"`
			} `json:"Hp"`
		} `json:"Oem"`
		PermanentMACAddress string `json:"PermanentMACAddress"`
		SpeedMbps           int    `json:"SpeedMbps"`
		Status              struct {
			Health string `json:"Health"`
			State  string `json:"State"`
		} `json:"Status"`
		Type  string `json:"Type"`
		Links struct {
			NetworkServices []struct {
				Href string `json:"href"`
			} `json:"NetworkServices"`
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"Items"`
	MemberType string `json:"MemberType"`
	Members    []struct {
		_odata_id string
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
	Total               int    `json:"Total"`
	Type                string `json:"Type"`
	Links               struct {
		Member []struct {
			Href string `json:"href"`
		} `json:"Member"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//ProcessorInfo4HP ...
type ProcessorInfoHP struct {
	OdataContext   string `json:"@odata.context"`
	OdataID        string `json:"@odata.id"`
	OdataType      string `json:"@odata.type"`
	ID             string `json:"Id"`
	InstructionSet string `json:"InstructionSet"`
	Manufacturer   string `json:"Manufacturer"`
	MaxSpeedMHz    int64  `json:"MaxSpeedMHz"`
	Model          string `json:"Model"`
	Name           string `json:"Name"`
	Oem            struct {
		Hp struct {
			OdataType string `json:"@odata.type"`
			AssetTag  string `json:"AssetTag"`
			Cache     []struct {
				Associativity     string   `json:"Associativity"`
				CacheSpeedns      int64    `json:"CacheSpeedns"`
				CurrentSRAMType   []string `json:"CurrentSRAMType"`
				EccType           string   `json:"EccType"`
				InstalledSizeKB   int64    `json:"InstalledSizeKB"`
				Location          string   `json:"Location"`
				MaximumSizeKB     int64    `json:"MaximumSizeKB"`
				Name              string   `json:"Name"`
				Policy            string   `json:"Policy"`
				Socketed          bool     `json:"Socketed"`
				SupportedSRAMType []string `json:"SupportedSRAMType"`
				SystemCacheType   string   `json:"SystemCacheType"`
			} `json:"Cache"`
			Characteristics []string `json:"Characteristics"`
			ConfigStatus    struct {
				Populated bool   `json:"Populated"`
				State     string `json:"State"`
			} `json:"ConfigStatus"`
			CoresEnabled     int64 `json:"CoresEnabled"`
			ExternalClockMHz int64 `json:"ExternalClockMHz"`
			MicrocodePatches []struct {
				CPUID   string `json:"CpuId"`
				Date    string `json:"Date"`
				PatchID string `json:"PatchId"`
			} `json:"MicrocodePatches"`
			PartNumber      string `json:"PartNumber"`
			RatedSpeedMHz   int64  `json:"RatedSpeedMHz"`
			SerialNumber    string `json:"SerialNumber"`
			Type            string `json:"Type"`
			VoltageVoltsX10 int64  `json:"VoltageVoltsX10"`
		} `json:"Hp"`
	} `json:"Oem"`
	ProcessorArchitecture string `json:"ProcessorArchitecture"`
	ProcessorID           struct {
		EffectiveFamily         string      `json:"EffectiveFamily"`
		EffectiveModel          string      `json:"EffectiveModel"`
		IdentificationRegisters string      `json:"IdentificationRegisters"`
		MicrocodeInfo           interface{} `json:"MicrocodeInfo"`
		Step                    string      `json:"Step"`
		VendorID                string      `json:"VendorId"`
	} `json:"ProcessorId"`
	ProcessorType string `json:"ProcessorType"`
	Socket        string `json:"Socket"`
	Status        struct {
		Health string `json:"Health"`
	} `json:"Status"`
	TotalCores   int64  `json:"TotalCores"`
	TotalThreads int64  `json:"TotalThreads"`
	Type         string `json:"Type"`
	Links        struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//ProcessorInfo5HP ...
type ProcessorInfo5HP struct {
	OdataContext   string `json:"@odata.context"`
	OdataEtag      string `json:"@odata.etag"`
	OdataID        string `json:"@odata.id"`
	OdataType      string `json:"@odata.type"`
	ID             string `json:"Id"`
	InstructionSet string `json:"InstructionSet"`
	Manufacturer   string `json:"Manufacturer"`
	MaxSpeedMHz    int64  `json:"MaxSpeedMHz"`
	Model          string `json:"Model"`
	Name           string `json:"Name"`
	Oem            struct {
		Hpe struct {
			OdataContext string `json:"@odata.context"`
			OdataType    string `json:"@odata.type"`
			AssetTag     string `json:"AssetTag"`
			Cache        []struct {
				Associativity     string   `json:"Associativity"`
				CacheSpeedns      int64    `json:"CacheSpeedns"`
				CurrentSRAMType   []string `json:"CurrentSRAMType"`
				EccType           string   `json:"EccType"`
				InstalledSizeKB   int64    `json:"InstalledSizeKB"`
				Location          string   `json:"Location"`
				MaximumSizeKB     int64    `json:"MaximumSizeKB"`
				Name              string   `json:"Name"`
				Policy            string   `json:"Policy"`
				Socketed          bool     `json:"Socketed"`
				SupportedSRAMType []string `json:"SupportedSRAMType"`
				SystemCacheType   string   `json:"SystemCacheType"`
			} `json:"Cache"`
			Characteristics []string `json:"Characteristics"`
			ConfigStatus    struct {
				Populated bool   `json:"Populated"`
				State     string `json:"State"`
			} `json:"ConfigStatus"`
			CoresEnabled     int64 `json:"CoresEnabled"`
			ExternalClockMHz int64 `json:"ExternalClockMHz"`
			MicrocodePatches []struct {
				CPUID   string `json:"CpuId"`
				Date    string `json:"Date"`
				PatchID string `json:"PatchId"`
			} `json:"MicrocodePatches"`
			PartNumber      string `json:"PartNumber"`
			RatedSpeedMHz   int64  `json:"RatedSpeedMHz"`
			SerialNumber    string `json:"SerialNumber"`
			VoltageVoltsX10 int64  `json:"VoltageVoltsX10"`
		} `json:"Hpe"`
	} `json:"Oem"`
	ProcessorArchitecture string `json:"ProcessorArchitecture"`
	ProcessorID           struct {
		EffectiveFamily         string      `json:"EffectiveFamily"`
		EffectiveModel          string      `json:"EffectiveModel"`
		IdentificationRegisters string      `json:"IdentificationRegisters"`
		MicrocodeInfo           interface{} `json:"MicrocodeInfo"`
		Step                    string      `json:"Step"`
		VendorID                string      `json:"VendorId"`
	} `json:"ProcessorId"`
	ProcessorType string `json:"ProcessorType"`
	Socket        string `json:"Socket"`
	Status        struct {
		Health string `json:"Health"`
	} `json:"Status"`
	TotalCores   int64 `json:"TotalCores"`
	TotalThreads int64 `json:"TotalThreads"`
}

//AccountsInfoHP ...
type AccountsInfoHP struct {
	OdataContext string
	OdataID      string
	OdataType    string
	Description  string `json:"Description"`
	Items        []struct {
		_odata_context string
		_odata_id      string
		_odata_type    string
		Description    string `json:"Description"`
		ID             string `json:"Id"`
		Name           string `json:"Name"`
		Oem            struct {
			Hp struct {
				_odata_type string
				LoginName   string `json:"LoginName"`
				Privileges  struct {
					LoginPriv                bool `json:"LoginPriv"`
					RemoteConsolePriv        bool `json:"RemoteConsolePriv"`
					UserConfigPriv           bool `json:"UserConfigPriv"`
					VirtualMediaPriv         bool `json:"VirtualMediaPriv"`
					VirtualPowerAndResetPriv bool `json:"VirtualPowerAndResetPriv"`
					ILOConfigPriv            bool `json:"iLOConfigPriv"`
				} `json:"Privileges"`
				Type string `json:"Type"`
			} `json:"Hp"`
		} `json:"Oem"`
		Password interface{} `json:"Password"`
		Type     string      `json:"Type"`
		UserName string      `json:"UserName"`
		Links    struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"Items"`
	MemberType string `json:"MemberType"`
	Members    []struct {
		_odata_id string
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
	Total               int    `json:"Total"`
	Type                string `json:"Type"`
	Links               struct {
		Member []struct {
			Href string `json:"href"`
		} `json:"Member"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//SystemEventLogsHP ...
type SystemEventLogsHP struct {
	OdataContext string
	OdataID      string
	OdataType    string
	Description  string `json:"Description"`
	Items        []struct {
		_odata_context string
		_odata_id      string
		_odata_type    string
		Created        string `json:"Created"`
		EntryType      string `json:"EntryType"`
		ID             string `json:"Id"`
		Message        string `json:"Message"`
		Name           string `json:"Name"`
		Number         int    `json:"Number"`
		Oem            struct {
			Hp struct {
				_odata_type string
				EventNumber int    `json:"EventNumber"`
				Type        string `json:"Type"`
				Updated     string `json:"Updated"`
			} `json:"Hp"`
		} `json:"Oem"`
		OemRecordFormat string `json:"OemRecordFormat"`
		RecordID        int    `json:"RecordId"`
		Severity        string `json:"Severity"`
		Type            string `json:"Type"`
		Links           struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"Items"`
	MemberType string `json:"MemberType"`
	Members    []struct {
		_odata_id string
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
	Total               int    `json:"Total"`
	Type                string `json:"Type"`
	Links               struct {
		Member []struct {
			Href string `json:"href"`
		} `json:"Member"`
		NextPage struct {
			Count int `json:"count"`
			Page  int `json:"page"`
		} `json:"NextPage"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//BiosAttrHP ... Bios Settings from the Redfish API
type BiosAttrHP struct {
	AcpiRootBridgePxm            string      `json:"AcpiRootBridgePxm"`
	AcpiSlit                     string      `json:"AcpiSlit"`
	AdjSecPrefetch               string      `json:"AdjSecPrefetch"`
	AdminEmail                   string      `json:"AdminEmail"`
	AdminName                    string      `json:"AdminName"`
	AdminOtherInfo               string      `json:"AdminOtherInfo"`
	AdminPassword                interface{} `json:"AdminPassword"`
	AdminPhone                   string      `json:"AdminPhone"`
	AdvancedMemProtection        string      `json:"AdvancedMemProtection"`
	AsrStatus                    string      `json:"AsrStatus"`
	AsrTimeoutMinutes            string      `json:"AsrTimeoutMinutes"`
	AssetTagProtection           string      `json:"AssetTagProtection"`
	AutoPowerOn                  string      `json:"AutoPowerOn"`
	BootMode                     string      `json:"BootMode"`
	BootOrderPolicy              string      `json:"BootOrderPolicy"`
	ChannelInterleaving          string      `json:"ChannelInterleaving"`
	CollabPowerControl           string      `json:"CollabPowerControl"`
	ConsistentDevNaming          string      `json:"ConsistentDevNaming"`
	CustomPostMessage            string      `json:"CustomPostMessage"`
	DaylightSavingsTime          string      `json:"DaylightSavingsTime"`
	DcuIPPrefetcher              string      `json:"DcuIpPrefetcher"`
	DcuStreamPrefetcher          string      `json:"DcuStreamPrefetcher"`
	Description                  string      `json:"Description"`
	Dhcpv4                       string      `json:"Dhcpv4"`
	DynamicPowerCapping          string      `json:"DynamicPowerCapping"`
	DynamicPowerResponse         string      `json:"DynamicPowerResponse"`
	EmbNicEnable                 string      `json:"EmbNicEnable"`
	EmbSas1Boot                  string      `json:"EmbSas1Boot"`
	EmbSasEnable                 string      `json:"EmbSasEnable"`
	EmbSata1Enable               string      `json:"EmbSata1Enable"`
	EmbSata2Enable               string      `json:"EmbSata2Enable"`
	EmbVideoConnection           string      `json:"EmbVideoConnection"`
	EmbeddedDiagnostics          string      `json:"EmbeddedDiagnostics"`
	EmbeddedDiagsMode            string      `json:"EmbeddedDiagsMode"`
	EmbeddedSata                 string      `json:"EmbeddedSata"`
	EmbeddedSerialPort           string      `json:"EmbeddedSerialPort"`
	EmbeddedUefiShell            string      `json:"EmbeddedUefiShell"`
	EmbeddedUserPartition        string      `json:"EmbeddedUserPartition"`
	EmsConsole                   string      `json:"EmsConsole"`
	EnergyPerfBias               string      `json:"EnergyPerfBias"`
	EraseUserDefaults            string      `json:"EraseUserDefaults"`
	ExtendedAmbientTemp          string      `json:"ExtendedAmbientTemp"`
	ExtendedMemTest              string      `json:"ExtendedMemTest"`
	F11BootMenu                  string      `json:"F11BootMenu"`
	FCScanPolicy                 string      `json:"FCScanPolicy"`
	FanFailPolicy                string      `json:"FanFailPolicy"`
	FanInstallReq                string      `json:"FanInstallReq"`
	FlexLom1Enable               string      `json:"FlexLom1Enable"`
	HwPrefetcher                 string      `json:"HwPrefetcher"`
	IntelDmiLinkFreq             string      `json:"IntelDmiLinkFreq"`
	IntelNicDmaChannels          string      `json:"IntelNicDmaChannels"`
	IntelPerfMonitoring          string      `json:"IntelPerfMonitoring"`
	IntelProcVtd                 string      `json:"IntelProcVtd"`
	IntelQpiFreq                 string      `json:"IntelQpiFreq"`
	IntelQpiLinkEn               string      `json:"IntelQpiLinkEn"`
	IntelQpiPowerManagement      string      `json:"IntelQpiPowerManagement"`
	IntelligentProvisioning      string      `json:"IntelligentProvisioning"`
	InternalSDCardSlot           string      `json:"InternalSDCardSlot"`
	IoNonPostedPrefetching       string      `json:"IoNonPostedPrefetching"`
	Ipv4Address                  string      `json:"Ipv4Address"`
	Ipv4Gateway                  string      `json:"Ipv4Gateway"`
	Ipv4PrimaryDNS               string      `json:"Ipv4PrimaryDNS"`
	Ipv4SecondaryDNS             string      `json:"Ipv4SecondaryDNS"`
	Ipv4SubnetMask               string      `json:"Ipv4SubnetMask"`
	Ipv6Duid                     string      `json:"Ipv6Duid"`
	MaxMemBusFreqMHz             string      `json:"MaxMemBusFreqMHz"`
	MaxPcieSpeed                 string      `json:"MaxPcieSpeed"`
	MemFastTraining              string      `json:"MemFastTraining"`
	MinProcIdlePkgState          string      `json:"MinProcIdlePkgState"`
	MinProcIdlePower             string      `json:"MinProcIdlePower"`
	MixedPowerSupplyReporting    string      `json:"MixedPowerSupplyReporting"`
	Modified                     string      `json:"Modified"`
	Name                         string      `json:"Name"`
	NetworkBootRetry             string      `json:"NetworkBootRetry"`
	NicBoot1                     string      `json:"NicBoot1"`
	NicBoot2                     string      `json:"NicBoot2"`
	NicBoot3                     string      `json:"NicBoot3"`
	NicBoot4                     string      `json:"NicBoot4"`
	NmiDebugButton               string      `json:"NmiDebugButton"`
	NodeInterleaving             string      `json:"NodeInterleaving"`
	NumaGroupSizeOpt             string      `json:"NumaGroupSizeOpt"`
	OldAdminPassword             interface{} `json:"OldAdminPassword"`
	OldPowerOnPassword           interface{} `json:"OldPowerOnPassword"`
	PciBusPadding                string      `json:"PciBusPadding"`
	PciSlot3Enable               string      `json:"PciSlot3Enable"`
	PciSlot4Enable               string      `json:"PciSlot4Enable"`
	PciSlot6Enable               string      `json:"PciSlot6Enable"`
	PcieExpressEcrcSupport       string      `json:"PcieExpressEcrcSupport"`
	PostF1Prompt                 string      `json:"PostF1Prompt"`
	PowerButton                  string      `json:"PowerButton"`
	PowerOnDelay                 string      `json:"PowerOnDelay"`
	PowerOnLogo                  string      `json:"PowerOnLogo"`
	PowerOnPassword              interface{} `json:"PowerOnPassword"`
	PowerProfile                 string      `json:"PowerProfile"`
	PowerRegulator               string      `json:"PowerRegulator"`
	PreBootNetwork               string      `json:"PreBootNetwork"`
	ProcAes                      string      `json:"ProcAes"`
	ProcCoreDisable              int         `json:"ProcCoreDisable"`
	ProcHyperthreading           string      `json:"ProcHyperthreading"`
	ProcNoExecute                string      `json:"ProcNoExecute"`
	ProcTurbo                    string      `json:"ProcTurbo"`
	ProcVirtualization           string      `json:"ProcVirtualization"`
	ProcX2Apic                   string      `json:"ProcX2Apic"`
	ProductID                    string      `json:"ProductId"`
	QpiBandwidthOpt              string      `json:"QpiBandwidthOpt"`
	QpiSnoopConfig               string      `json:"QpiSnoopConfig"`
	RedundantPowerSupply         string      `json:"RedundantPowerSupply"`
	RemovableFlashBootSeq        string      `json:"RemovableFlashBootSeq"`
	RestoreDefaults              string      `json:"RestoreDefaults"`
	RestoreManufacturingDefaults string      `json:"RestoreManufacturingDefaults"`
	RomSelection                 string      `json:"RomSelection"`
	SataSecureErase              string      `json:"SataSecureErase"`
	SaveUserDefaults             string      `json:"SaveUserDefaults"`
	SecureBootStatus             string      `json:"SecureBootStatus"`
	SerialConsoleBaudRate        string      `json:"SerialConsoleBaudRate"`
	SerialConsoleEmulation       string      `json:"SerialConsoleEmulation"`
	SerialConsolePort            string      `json:"SerialConsolePort"`
	SerialNumber                 string      `json:"SerialNumber"`
	ServerAssetTag               string      `json:"ServerAssetTag"`
	ServerName                   string      `json:"ServerName"`
	ServerOtherInfo              string      `json:"ServerOtherInfo"`
	ServerPrimaryOs              string      `json:"ServerPrimaryOs"`
	ServiceEmail                 string      `json:"ServiceEmail"`
	ServiceName                  string      `json:"ServiceName"`
	ServiceOtherInfo             string      `json:"ServiceOtherInfo"`
	ServicePhone                 string      `json:"ServicePhone"`
	Slot3NicBoot1                string      `json:"Slot3NicBoot1"`
	Slot3NicBoot2                string      `json:"Slot3NicBoot2"`
	Slot4NicBoot1                string      `json:"Slot4NicBoot1"`
	Slot4NicBoot2                string      `json:"Slot4NicBoot2"`
	Slot6NicBoot1                string      `json:"Slot6NicBoot1"`
	Slot6NicBoot2                string      `json:"Slot6NicBoot2"`
	Sriov                        string      `json:"Sriov"`
	ThermalConfig                string      `json:"ThermalConfig"`
	ThermalShutdown              string      `json:"ThermalShutdown"`
	TimeFormat                   string      `json:"TimeFormat"`
	TimeZone                     string      `json:"TimeZone"`
	TpmState                     string      `json:"TpmState"`
	TpmType                      string      `json:"TpmType"`
	Type                         string      `json:"Type"`
	UefiOptimizedBoot            string      `json:"UefiOptimizedBoot"`
	UefiPxeBoot                  string      `json:"UefiPxeBoot"`
	UefiShellBootOrder           string      `json:"UefiShellBootOrder"`
	UefiShellStartup             string      `json:"UefiShellStartup"`
	UefiShellStartupLocation     string      `json:"UefiShellStartupLocation"`
	UefiShellStartupURL          string      `json:"UefiShellStartupUrl"`
	URLBootFile                  string      `json:"UrlBootFile"`
	Usb3Mode                     string      `json:"Usb3Mode"`
	UsbBoot                      string      `json:"UsbBoot"`
	UsbControl                   string      `json:"UsbControl"`
	UtilityLang                  string      `json:"UtilityLang"`
	VirtualInstallDisk           string      `json:"VirtualInstallDisk"`
	VirtualSerialPort            string      `json:"VirtualSerialPort"`
	VlanControl                  string      `json:"VlanControl"`
	VlanID                       int         `json:"VlanId"`
	VlanPriority                 int         `json:"VlanPriority"`
	WakeOnLan                    string      `json:"WakeOnLan"`
	Links                        struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//LicenseInfoHP ... License Details from the Redfish API
type LicenseInfoHP struct {
	OdataContext string `json:"@odata.context"`
	OdataID      string `json:"@odata.id"`
	OdataType    string `json:"@odata.type"`
	Description  string `json:"Description"`
	Items        []struct {
		OdataContext string `json:"@odata.context"`
		OdataID      string `json:"@odata.id"`
		OdataType    string `json:"@odata.type"`
		Description  string `json:"Description"`
		ID           string `json:"Id"`
		License      string `json:"License"`
		LicenseKey   string `json:"LicenseKey"`
		LicenseType  string `json:"LicenseType"`
		Name         string `json:"Name"`
		Type         string `json:"Type"`
		Links        struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"Items"`
	MemberType string `json:"MemberType"`
	Members    []struct {
		OdataID string `json:"@odata.id"`
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
	Total               int    `json:"Total"`
	Type                string `json:"Type"`
	Links               struct {
		Member []struct {
			Href string `json:"href"`
		} `json:"Member"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//PCISlotsInfoHp ... PCI Slots Details from the Redfish API
type PCISlotsInfoHP struct {
	OdataContext string
	OdataID      string
	OdataType    string
	Description  string `json:"Description"`
	Items        []struct {
		OdataContext string `json:"@odata.context"`
		OdataID      string `json:"@odata.id"`
		OdataType    string `json:"@odata.type"`
		ID           string `json:"Id"`
		Length       string `json:"Length"`
		LinkLanes    string `json:"LinkLanes"`
		Name         string `json:"Name"`
		Status       struct {
			OperationalStatus []struct {
				Status string `json:"Status"`
			} `json:"OperationalStatus"`
		} `json:"Status"`
		SupportsHotPlug bool   `json:"SupportsHotPlug"`
		Technology      string `json:"Technology"`
		Type            string `json:"Type"`
		UEFIDevicePath  string `json:"UEFIDevicePath"`
		Links           struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"links"`
	} `json:"Items"`
	MemberType string `json:"MemberType"`
	Members    []struct {
		OdataID string
	} `json:"Members"`
	Members_odata_count int    `json:"Members@odata.count"`
	Name                string `json:"Name"`
	Total               int    `json:"Total"`
	Type                string `json:"Type"`
	Links               struct {
		Member []struct {
			Href string `json:"href"`
		} `json:"Member"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"links"`
}

//Custom Structs

//HealthList ...
type HealthList struct {
	Name   string `json:"name"`
	Health string `json:"health"`
	State  string `json:"state"`
}

//StorageHealthList ...
type StorageHealthList struct {
	Name   string `json:"name"`
	Health string `json:"health"`
	State  string `json:"state"`
	Space  int    `json:"space"`
}

//SystemEventLogRes ...
type SystemEventLogRes struct {
	EntryCode  string `json:"entry_code"`
	Message    string `json:"message"`
	Name       string `json:"name"`
	SensorType string `json:"sensor_type"`
	Severity   string `json:"severity"`
}

//LifeCycleEventLogRes ...
type LifeCycleEventLogRes struct {
	Created     string `json:"created"`
	Description string `json:"description"`
	EntryType   string `json:"entry_type"`
	ID          string `json:"id"`
	Message     string `json:"message"`
	MessageID   string `json:"message_id"`
	Name        string `json:"name"`
	Severity    string `json:"severity"`
}

//Accounts ...
type Accounts struct {
	Enabled  bool   `json:"enabled"`
	Locked   bool   `json:"locked"`
	Name     string `json:"name"`
	RoleId   string `json:"role_id"`
	Username string `json:"username"`
}

//MACData ...
type MACData struct {
	MacAddress   string `json:"macaddress"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Status       string `json:"status"`
	State        string `json:"state"`
	Vlan         string `json:"vlan"`
	PartNumber   string `json:"partnumber"`
	SerialNumber string `json:"serialnumber"`
	VendorName   string `json:"verndorname"`
}

//Check MACData struct for empty string
func (intInfo *MACData) UpdateEmpty() {
	if intInfo.PartNumber == "" {
		intInfo.PartNumber = "NULL"
	}
	if intInfo.SerialNumber == "" {
		intInfo.SerialNumber = "NULL"
	}
	if intInfo.VendorName == "" {
		intInfo.VendorName = "NULL"
	}
}

//BootOrderData ...
type BootOrderData struct {
	Enabled bool   `json:"enabled"`
	Index   int    `json:"index"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

//IDRACData ...
type IDRACData struct {
	VirtualConsoleMaxSessions int    `json:"virtualconsole_maxsessions"`
	VirtualConsolePluginType  string `json:"virtualconsole_plugintype"`
	WebServerSSLEncryption    string `json:"webserver_ssl_encryption"`
	IPMILanEnable             string `json:"ipmi_lan_enable"`
	DNSDomainName             string `json:"dns_domainname"`
	SnmpAgentStatus           string `json:"snmp_agent_status"`
	SnmpAgentCommunity        string `json:"snmp_agent_community"`
}

//BiosData ... Return Response for Bios Data for Dell
type BiosData struct {
	BootMode          string `json:"bootmode"`
	BootSeqRetry      string `json:"boot_sequence_retry"`
	InternalUsb       string `json:"internal_usb"`
	SriovGlobalEnable string `json:"sriov_enable"`
	SysProfile        string `json:"sys_profile"`
	AcPwrRcvry        string `json:"pwr_rcvry"`
	AcPwrRcvryDelay   string `json:"pwr_rcvry_delay"`
	Serial            string `json:"serial"`
}

// BiosDataHP ... ... Return Response for Bios Data for HP
type BiosDataHP struct {
	AcpiRootBridgePxm            string `json:"AcpiRootBridgePxm"`
	AcpiSlit                     string `json:"AcpiSlit"`
	AdjSecPrefetch               string `json:"AdjSecPrefetch"`
	AdminEmail                   string `json:"AdminEmail"`
	AdminName                    string `json:"AdminName"`
	AdminOtherInfo               string `json:"AdminOtherInfo"`
	AdminPhone                   string `json:"AdminPhone"`
	AdvancedMemProtection        string `json:"AdvancedMemProtection"`
	AsrStatus                    string `json:"AsrStatus"`
	AsrTimeoutMinutes            string `json:"AsrTimeoutMinutes"`
	AssetTagProtection           string `json:"AssetTagProtection"`
	AutoPowerOn                  string `json:"AutoPowerOn"`
	BootMode                     string `json:"BootMode"`
	BootOrderPolicy              string `json:"BootOrderPolicy"`
	ChannelInterleaving          string `json:"ChannelInterleaving"`
	CollabPowerControl           string `json:"CollabPowerControl"`
	ConsistentDevNaming          string `json:"ConsistentDevNaming"`
	CustomPostMessage            string `json:"CustomPostMessage"`
	DaylightSavingsTime          string `json:"DaylightSavingsTime"`
	DcuIPPrefetcher              string `json:"DcuIpPrefetcher"`
	DcuStreamPrefetcher          string `json:"DcuStreamPrefetcher"`
	Description                  string `json:"Description"`
	Dhcpv4                       string `json:"Dhcpv4"`
	DynamicPowerCapping          string `json:"DynamicPowerCapping"`
	DynamicPowerResponse         string `json:"DynamicPowerResponse"`
	EmbNicEnable                 string `json:"EmbNicEnable"`
	EmbSas1Boot                  string `json:"EmbSas1Boot"`
	EmbSasEnable                 string `json:"EmbSasEnable"`
	EmbSata1Enable               string `json:"EmbSata1Enable"`
	EmbSata2Enable               string `json:"EmbSata2Enable"`
	EmbVideoConnection           string `json:"EmbVideoConnection"`
	EmbeddedDiagnostics          string `json:"EmbeddedDiagnostics"`
	EmbeddedDiagsMode            string `json:"EmbeddedDiagsMode"`
	EmbeddedSata                 string `json:"EmbeddedSata"`
	EmbeddedSerialPort           string `json:"EmbeddedSerialPort"`
	EmbeddedUefiShell            string `json:"EmbeddedUefiShell"`
	EmbeddedUserPartition        string `json:"EmbeddedUserPartition"`
	EmsConsole                   string `json:"EmsConsole"`
	EnergyPerfBias               string `json:"EnergyPerfBias"`
	EraseUserDefaults            string `json:"EraseUserDefaults"`
	ExtendedAmbientTemp          string `json:"ExtendedAmbientTemp"`
	ExtendedMemTest              string `json:"ExtendedMemTest"`
	F11BootMenu                  string `json:"F11BootMenu"`
	FCScanPolicy                 string `json:"FCScanPolicy"`
	FanFailPolicy                string `json:"FanFailPolicy"`
	FanInstallReq                string `json:"FanInstallReq"`
	FlexLom1Enable               string `json:"FlexLom1Enable"`
	HwPrefetcher                 string `json:"HwPrefetcher"`
	IntelDmiLinkFreq             string `json:"IntelDmiLinkFreq"`
	IntelNicDmaChannels          string `json:"IntelNicDmaChannels"`
	IntelPerfMonitoring          string `json:"IntelPerfMonitoring"`
	IntelProcVtd                 string `json:"IntelProcVtd"`
	IntelQpiFreq                 string `json:"IntelQpiFreq"`
	IntelQpiLinkEn               string `json:"IntelQpiLinkEn"`
	IntelQpiPowerManagement      string `json:"IntelQpiPowerManagement"`
	IntelligentProvisioning      string `json:"IntelligentProvisioning"`
	InternalSDCardSlot           string `json:"InternalSDCardSlot"`
	IoNonPostedPrefetching       string `json:"IoNonPostedPrefetching"`
	Ipv4Address                  string `json:"Ipv4Address"`
	Ipv4Gateway                  string `json:"Ipv4Gateway"`
	Ipv4PrimaryDNS               string `json:"Ipv4PrimaryDNS"`
	Ipv4SecondaryDNS             string `json:"Ipv4SecondaryDNS"`
	Ipv4SubnetMask               string `json:"Ipv4SubnetMask"`
	Ipv6Duid                     string `json:"Ipv6Duid"`
	MaxMemBusFreqMHz             string `json:"MaxMemBusFreqMHz"`
	MaxPcieSpeed                 string `json:"MaxPcieSpeed"`
	MemFastTraining              string `json:"MemFastTraining"`
	MinProcIdlePkgState          string `json:"MinProcIdlePkgState"`
	MinProcIdlePower             string `json:"MinProcIdlePower"`
	MixedPowerSupplyReporting    string `json:"MixedPowerSupplyReporting"`
	Modified                     string `json:"Modified"`
	Name                         string `json:"Name"`
	NetworkBootRetry             string `json:"NetworkBootRetry"`
	NicBoot1                     string `json:"NicBoot1"`
	NicBoot2                     string `json:"NicBoot2"`
	NicBoot3                     string `json:"NicBoot3"`
	NicBoot4                     string `json:"NicBoot4"`
	NmiDebugButton               string `json:"NmiDebugButton"`
	NodeInterleaving             string `json:"NodeInterleaving"`
	NumaGroupSizeOpt             string `json:"NumaGroupSizeOpt"`
	PciBusPadding                string `json:"PciBusPadding"`
	PciSlot3Enable               string `json:"PciSlot3Enable"`
	PciSlot4Enable               string `json:"PciSlot4Enable"`
	PciSlot6Enable               string `json:"PciSlot6Enable"`
	PcieExpressEcrcSupport       string `json:"PcieExpressEcrcSupport"`
	PostF1Prompt                 string `json:"PostF1Prompt"`
	PowerButton                  string `json:"PowerButton"`
	PowerOnDelay                 string `json:"PowerOnDelay"`
	PowerOnLogo                  string `json:"PowerOnLogo"`
	PowerProfile                 string `json:"PowerProfile"`
	PowerRegulator               string `json:"PowerRegulator"`
	PreBootNetwork               string `json:"PreBootNetwork"`
	ProcAes                      string `json:"ProcAes"`
	ProcCoreDisable              int    `json:"ProcCoreDisable"`
	ProcHyperthreading           string `json:"ProcHyperthreading"`
	ProcNoExecute                string `json:"ProcNoExecute"`
	ProcTurbo                    string `json:"ProcTurbo"`
	ProcVirtualization           string `json:"ProcVirtualization"`
	ProcX2Apic                   string `json:"ProcX2Apic"`
	ProductID                    string `json:"ProductId"`
	QpiBandwidthOpt              string `json:"QpiBandwidthOpt"`
	QpiSnoopConfig               string `json:"QpiSnoopConfig"`
	RedundantPowerSupply         string `json:"RedundantPowerSupply"`
	RemovableFlashBootSeq        string `json:"RemovableFlashBootSeq"`
	RestoreDefaults              string `json:"RestoreDefaults"`
	RestoreManufacturingDefaults string `json:"RestoreManufacturingDefaults"`
	RomSelection                 string `json:"RomSelection"`
	SataSecureErase              string `json:"SataSecureErase"`
	SaveUserDefaults             string `json:"SaveUserDefaults"`
	SecureBootStatus             string `json:"SecureBootStatus"`
	SerialConsoleBaudRate        string `json:"SerialConsoleBaudRate"`
	SerialConsoleEmulation       string `json:"SerialConsoleEmulation"`
	SerialConsolePort            string `json:"SerialConsolePort"`
	SerialNumber                 string `json:"SerialNumber"`
	ServerAssetTag               string `json:"ServerAssetTag"`
	ServerName                   string `json:"ServerName"`
	ServerOtherInfo              string `json:"ServerOtherInfo"`
	ServerPrimaryOs              string `json:"ServerPrimaryOs"`
	ServiceEmail                 string `json:"ServiceEmail"`
	ServiceName                  string `json:"ServiceName"`
	ServiceOtherInfo             string `json:"ServiceOtherInfo"`
	ServicePhone                 string `json:"ServicePhone"`
	Slot3NicBoot1                string `json:"Slot3NicBoot1"`
	Slot3NicBoot2                string `json:"Slot3NicBoot2"`
	Slot4NicBoot1                string `json:"Slot4NicBoot1"`
	Slot4NicBoot2                string `json:"Slot4NicBoot2"`
	Slot6NicBoot1                string `json:"Slot6NicBoot1"`
	Slot6NicBoot2                string `json:"Slot6NicBoot2"`
	Sriov                        string `json:"Sriov"`
	ThermalConfig                string `json:"ThermalConfig"`
	ThermalShutdown              string `json:"ThermalShutdown"`
	TimeFormat                   string `json:"TimeFormat"`
	TimeZone                     string `json:"TimeZone"`
	TpmState                     string `json:"TpmState"`
	TpmType                      string `json:"TpmType"`
	Type                         string `json:"Type"`
	UefiOptimizedBoot            string `json:"UefiOptimizedBoot"`
	UefiPxeBoot                  string `json:"UefiPxeBoot"`
	UefiShellBootOrder           string `json:"UefiShellBootOrder"`
	UefiShellStartup             string `json:"UefiShellStartup"`
	UefiShellStartupLocation     string `json:"UefiShellStartupLocation"`
	UefiShellStartupURL          string `json:"UefiShellStartupUrl"`
	URLBootFile                  string `json:"UrlBootFile"`
	Usb3Mode                     string `json:"Usb3Mode"`
	UsbBoot                      string `json:"UsbBoot"`
	UsbControl                   string `json:"UsbControl"`
	UtilityLang                  string `json:"UtilityLang"`
	VirtualInstallDisk           string `json:"VirtualInstallDisk"`
	VirtualSerialPort            string `json:"VirtualSerialPort"`
	VlanControl                  string `json:"VlanControl"`
	VlanID                       int    `json:"VlanId"`
	VlanPriority                 int    `json:"VlanPriority"`
	WakeOnLan                    string `json:"WakeOnLan"`
}

//SysAttrData ... Return Response for System Attribute Data
type SysAttrData struct {
	ServerPwrPSRedPolicy string `json:"redundancy_policy"`
	ServerPwrPSRapidOn   string `json:"hot_spare"`
}

//LifeCycleData ...
type LifeCycleData struct {
	AutoBackup                          string      `json:"autobackup"`
	AutoDiscovery                       string      `json:"autodiscovery"`
	AutoUpdate                          string      `json:"autoupdate"`
	BIOSRTDRequested                    string      `json:"bios_rtd"`
	CollectSystemInventoryOnRestart     string      `json:"collect_system_inventory_on_restart"`
	DiscoveryFactoryDefaults            string      `json:"discovery_factory_defaults"`
	IPAddress                           string      `json:"ipaddress"`
	IPChangeNotifyPS                    string      `json:"ip_change_notify"`
	IgnoreCertWarning                   string      `json:"ignore_cert_warning"`
	Licensed                            string      `json:"licensed"`
	LifecycleControllerState            string      `json:"lifecycle_controller_state"`
	PartConfigurationUpdate             string      `json:"part_config_update"`
	PartFirmwareUpdate                  string      `json:"part_firmware_update"`
	ProvisioningServer                  string      `json:"provisioning_server"`
	StorageHealthRollupStatus           int         `json:"storage_health_rollup_status"`
	SystemID                            string      `json:"systemid"`
	UserProxyPassword                   interface{} `json:"user_proxy_password"`
	UserProxyPort                       string      `json:"user_proxy_port"`
	UserProxyServer                     string      `json:"user_proxy_server"`
	UserProxyType                       string      `json:"user_proxy_type"`
	UserProxyUserName                   string      `json:"user_proxy_username"`
	VirtualAddressManagementApplication string      `json:"virtual_addr_mgmt_application"`
}

//SystemData ...
type SystemData struct {
	PowerState      string  `json:"power_state"`
	SerialNumber    string  `json:"serial_number"`
	Health          string  `json:"health"`
	SystemType      string  `json:"system_type"`
	Model           string  `json:"model"`
	Memory          float32 `json:"memory"`
	Processors      int     `json:"processors"`
	ProcessorFamily string  `json:"processor_family"`
}

//FirmwareData ...
type FirmwareData struct {
	Name       string `json:"name"`
	Id         string `json:"id"`
	Version    string `json:"version"`
	Updateable bool   `json:"updateable"`
}

//LicenseInfo ...
type LicenseInfo struct {
	Name        string `json:"name"`
	LicenseKey  string `json:"license_key"`
	LicenseType string `json:"license_type"`
}

//PCISlotsInfo ...
type PCISlotsInfo struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	Technology string `json:"technology"`
	Type       string `json:"type"`
}

//ExportConfigStatus
type ExportConfigStatus struct {
	OdataContext string
	OdataID      string
	OdataType    string
	Description  string `json:"Description"`
	EndTime      string `json:"EndTime"`
	ID           string `json:"Id"`
	Messages     []struct {
		Message                 string        `json:"Message"`
		MessageArgs             []interface{} `json:"MessageArgs"`
		MessageArgs_odata_count int           `json:"MessageArgs@odata.count"`
		MessageID               string        `json:"MessageId"`
	} `json:"Messages"`
	Messages_odata_count int    `json:"Messages@odata.count"`
	Name                 string `json:"Name"`
	Oem                  struct {
		Dell struct {
			_odata_type       string
			CompletionTime    interface{}   `json:"CompletionTime"`
			Description       string        `json:"Description"`
			EndTime           interface{}   `json:"EndTime"`
			ID                string        `json:"Id"`
			JobState          string        `json:"JobState"`
			JobType           string        `json:"JobType"`
			Message           string        `json:"Message"`
			MessageArgs       []interface{} `json:"MessageArgs"`
			MessageID         string        `json:"MessageId"`
			Name              string        `json:"Name"`
			PercentComplete   int           `json:"PercentComplete"`
			StartTime         string        `json:"StartTime"`
			TargetSettingsURI interface{}   `json:"TargetSettingsURI"`
		} `json:"Dell"`
	} `json:"Oem"`
	StartTime  string `json:"StartTime"`
	TaskState  string `json:"TaskState"`
	TaskStatus string `json:"TaskStatus"`
}

type ExportConfigResponse struct {
	SystemConfiguration struct {
		Comments []struct {
			Comment string `json:"Comment"`
		} `json:"Comments"`
		Components []struct {
			Attributes []struct {
				Comment       string `json:"Comment"`
				Name          string `json:"Name"`
				Set_On_Import string `json:"Set On Import"`
				Value         string `json:"Value"`
			} `json:"Attributes"`
			FQDD string `json:"FQDD"`
		} `json:"Components"`
		Model      string `json:"Model"`
		ServiceTag string `json:"ServiceTag"`
		TimeStamp  string `json:"TimeStamp"`
	} `json:"SystemConfiguration"`
}

type JobResponseDell struct {
	MessageExtendedInfo []struct {
		Message                       string        `json:"Message"`
		MessageArgs                   []interface{} `json:"MessageArgs"`
		MessageArgs_odata_count       int           `json:"MessageArgs@odata.count"`
		MessageID                     string        `json:"MessageId"`
		RelatedProperties             []interface{} `json:"RelatedProperties"`
		RelatedProperties_odata_count int           `json:"RelatedProperties@odata.count"`
		Resolution                    string        `json:"Resolution"`
		Severity                      string        `json:"Severity"`
	} `json:"@Message.ExtendedInfo"`
}

type JobStatusDell struct {
	OdataContext      string
	OdataID           string
	OdataType         string
	CompletionTime    string        `json:"CompletionTime"`
	Description       string        `json:"Description"`
	EndTime           string        `json:"EndTime"`
	ID                string        `json:"Id"`
	JobState          string        `json:"JobState"`
	JobType           string        `json:"JobType"`
	Message           string        `json:"Message"`
	MessageArgs       []interface{} `json:"MessageArgs"`
	MessageID         string        `json:"MessageId"`
	Name              string        `json:"Name"`
	PercentComplete   int           `json:"PercentComplete"`
	StartTime         string        `json:"StartTime"`
	TargetSettingsURI interface{}   `json:"TargetSettingsURI"`
}

type UserListResponseDell struct {
	_odata_context string `json:"@odata.context"`
	_odata_id      string `json:"@odata.id"`
	_odata_type    string `json:"@odata.type"`
	Description    string `json:"Description"`
	Enabled        bool   `json:"Enabled"`
	ID             string `json:"Id"`
	Links          struct {
		Role struct {
			_odata_id string `json:"@odata.id"`
		} `json:"Role"`
	} `json:"Links"`
	Locked   bool        `json:"Locked"`
	Name     string      `json:"Name"`
	Password interface{} `json:"Password"`
	RoleID   string      `json:"RoleId"`
	UserName string      `json:"UserName"`
}

type UserListDell struct {
	UserName string `json:"UserName"`
	RoleID   string `json:"RoleId"`
	Enabled  bool   `json:"Enabled"`
	Locked   bool   `json:"Locked"`
}

//ImageStatusDell ...
type ImageStatusDell struct {
	Odata_context string `json:"@odata.context"`
	Odata_id      string `json:"@odata.id"`
	Odata_type    string `json:"@odata.type"`
	Actions       struct {
		VirtualMedia_EjectMedia struct {
			Target string `json:"target"`
		} `json:"#VirtualMedia.EjectMedia"`
		VirtualMedia_InsertMedia struct {
			Target string `json:"target"`
		} `json:"#VirtualMedia.InsertMedia"`
	} `json:"Actions"`
	ConnectedVia           string   `json:"ConnectedVia"`
	Description            string   `json:"Description"`
	ID                     string   `json:"Id"`
	Image                  string   `json:"Image"`
	ImageName              string   `json:"ImageName"`
	Inserted               bool     `json:"Inserted"`
	MediaTypes             []string `json:"MediaTypes"`
	MediaTypes_odata_count int64    `json:"MediaTypes@odata.count"`
	Name                   string   `json:"Name"`
	WriteProtected         bool     `json:"WriteProtected"`
}
