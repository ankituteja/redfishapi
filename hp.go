package redfishapi

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StartServerHP ...
// ResetType@Redfish.AllowableValues
// 0	"On"
// 1	"ForceOff"
// 2	"ForceRestart",
// 3	"Nmi",
// 4	"PushPowerButton"
// target: "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
func (c *IloClient) StartServerHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
	var jsonStr = []byte(`{"ResetType": "On"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Started", nil
}

// StopServerHP ... Will Request to stop the server
func (c *IloClient) StopServerHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/Actions/ComputerSystem.Reset/"
	var jsonStr = []byte(`{"ResetType": "ForceOff"}`)
	_, _, _, err := queryData(c, "POST", url, jsonStr)
	if err != nil {
		return "", err
	}

	return "Server Stopped", nil
}

// GetSystemInfoHP ... Will fetch the system info
func (c *IloClient) GetSystemInfoHP() (SystemData, error) {

	url := c.Hostname + "/redfish/v1/Systems/1"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return SystemData{}, err
	}

	var x SystemInfoHP

	json.Unmarshal(resp, &x)
	totalMemoryGB := x.Memory.TotalSystemMemoryGB
	if totalMemoryGB <= 0 && x.MemorySummary.TotalSystemMemoryGiB > 0 {
		totalMemoryGB = x.MemorySummary.TotalSystemMemoryGiB
	}

	_result := SystemData{Health: x.Status.Health,
		Memory:          totalMemoryGB,
		Model:           x.Model,
		PowerState:      x.PowerState,
		Processors:      x.ProcessorSummary.Count,
		ProcessorFamily: x.ProcessorSummary.Model,
		SerialNumber:    x.SerialNumber,
	}

	return _result, nil

}
// GetSystemMemoryInfoHP ... will fetch raw memory collection details
func (c *IloClient) GetSystemMemoryInfoHP() (MemoryCollectionHP, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/Memory/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return MemoryCollectionHP{}, err
	}

	var collection MemoryCollectionHP
	if err := json.Unmarshal(resp, &collection); err != nil {
		return MemoryCollectionHP{}, err
	}

	return collection, nil
}

// GetServerPowerStateHP ... Will fetch the current state of the Server
func (c *IloClient) GetServerPowerStateHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}

	var data SystemInfoHP

	json.Unmarshal(resp, &data)

	return data.Power, nil

}

// CheckLoginHP ... Will check the credentials of the Server
func (c *IloClient) CheckLoginHP() (string, error) {
	url := c.Hostname + "/redfish/v1/Systems/1"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return "", err
	}
	var data SystemInfoHP
	json.Unmarshal(resp, &data)
	return string(data.Status.Health), nil
}

// GetFirmwareHP ... will fetch the Firmware details
func (c *IloClient) GetFirmwareHP() ([]FirmwareData, error) {
	url := c.Hostname + "/redfish/v1/UpdateService/FirmwareInventory"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	var (
		x         MemberCountHP
		_firmdata []FirmwareData
	)
	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y FirmwareDataHP

		json.Unmarshal(resp, &y)

		deviceContext := y.Oem.Hpe.DeviceContext
		if deviceContext == "" {
			deviceContext = y.Name
		}
		deviceContext = strings.ReplaceAll(deviceContext, " ", ".")
		firmwareID := fmt.Sprintf("Installed-%s-%s__%s", y.ID, y.Version, deviceContext)

		_result := FirmwareData{
			Id:         firmwareID,
			Name:       y.Name,
			Updateable: y.Updateable,
			Version:    y.Version,
		}
		_firmdata = append(_firmdata, _result)
	}

	return _firmdata, nil
}

// GetThermalHealthHP ... will fetch the Thermal Health
func (c *IloClient) GetThermalHealthHP() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Chassis/1/Thermal/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       ThermalHealthListHP
		_health []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Fans {
		_result := HealthList{Name: x.Fans[i].FanName,
			Health: x.Fans[i].Status.Health,
			State:  x.Fans[i].Status.State}
		_health = append(_health, _result)
	}

	for i := range x.Temperatures {
		_result := HealthList{Name: x.Temperatures[i].Name,
			Health: x.Temperatures[i].Status.Health,
			State:  x.Temperatures[i].Status.State}
		_health = append(_health, _result)
	}

	return _health, nil
}

// GetPowerHealthHP ... will fetch the Power Health
func (c *IloClient) GetPowerHealthHP() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Chassis/1/Power/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       PowerDataHP
		_health []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.PowerSupplies {
		_name := fmt.Sprintf("%s_%d", x.PowerSupplies[i].Name, i)
		_result := HealthList{Name: _name,
			Health: x.PowerSupplies[i].Status.Health,
			State:  x.PowerSupplies[i].Status.State}
		_health = append(_health, _result)
	}

	return _health, nil
}

// GetInterfaceHealthHP ... will fetch the Interface Health
func (c *IloClient) GetInterfaceHealthHP() ([]HealthList, error) {
	url := c.Hostname + "/redfish/v1/Managers/1/EthernetInterfaces/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       EthernetInterfacesHP
		_health []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Items {
		_result := HealthList{Name: x.Items[i].Name,
			Health: x.Items[i].Status.Health,
			State:  x.Items[i].Status.State}
		_health = append(_health, _result)
	}

	return _health, nil
}

// GetProcessorHealthHP ... will Fetch the Processor Health Details
func (c *IloClient) GetProcessorInfoHP() ([]ProcessorInfoHP, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/Processors/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x           MemberCountHP
		processData []ProcessorInfoHP
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y ProcessorInfoHP

		json.Unmarshal(resp, &y)

		processData = append(processData, y)
	}

	return processData, nil

}

// GetProcessorHealthHP ... will Fetch the Processor Health Details
func (c *IloClient) GetProcessorHealthHP() ([]HealthList, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/Processors/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x             MemberCountHP
		processHealth []HealthList
	)

	json.Unmarshal(resp, &x)

	for i := range x.Members {
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}

		var y ProcessorInfoHP

		json.Unmarshal(resp, &y)

		procHealth := HealthList{
			Name:   y.ID,
			Health: y.Status.Health,
			State:  y.Oem.Hp.ConfigStatus.State,
		}
		processHealth = append(processHealth, procHealth)
	}

	return processHealth, nil

}

// GetUserAccountsHP ... will fetch the current User Accounts
func (c *IloClient) GetUserAccountsHP() ([]Accounts, error) {

	url := c.Hostname + "/redfish/v1/AccountService/Accounts"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var (
		x       AccountsInfoHP
		users   []Accounts
		_locked bool
	)

	json.Unmarshal(resp, &x)

	for i := range x.Items {

		if x.Items[i].Oem.Hp.Privileges.LoginPriv {
			_locked = false
		} else {
			_locked = true
		}

		user := Accounts{
			Name:     x.Items[i].Name,
			Enabled:  x.Items[i].Oem.Hp.Privileges.LoginPriv,
			Locked:   _locked,
			RoleId:   x.Items[i].ID,
			Username: x.Items[i].UserName,
		}
		users = append(users, user)

	}

	return users, nil

}

// GetSystemEventLogsHP ... will fetch the SystemEvent Logs
func (c *IloClient) GetSystemEventLogsHP() ([]SystemEventLogRes, error) {

	url := c.Hostname + "/redfish/v1/Managers/1/LogServices/IEL/Entries/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x SystemEventLogsHP

	json.Unmarshal(resp, &x)

	var _systemEventLogs []SystemEventLogRes

	for i := range x.Items {

		_result := SystemEventLogRes{
			EntryCode:  x.Items[i].EntryType,
			Message:    x.Items[i].Message,
			Name:       x.Items[i].Name,
			SensorType: x.Items[i].Type,
			Severity:   x.Items[i].Severity,
		}

		_systemEventLogs = append(_systemEventLogs, _result)
	}

	return _systemEventLogs, nil

}

// GetBiosDataHP ... will fetch BIOS settings and enrich with system memory/model, memory frequency, and processor brand/core details
func (c *IloClient) GetBiosDataHP() (BiosAttributesData, error) {

	url := c.Hostname + "/redfish/v1/systems/1/bios/settings/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return BiosAttributesData{}, err
	}

	var x BiosAttrHP

	err = json.Unmarshal(resp, &x)
	if err != nil {
		return BiosAttributesData{}, err
	}

	y := BiosAttributesData{
		BootMode:          x.BootMode,
		BootSeqRetry:      x.NetworkBootRetry,
		InternalUsb:       x.UsbControl,
		SriovGlobalEnable: x.Sriov,
		SysProfile:        x.PowerProfile,
		AcPwrRcvry:        x.AutoPowerOn,
		AcPwrRcvryDelay:   x.PowerOnDelay,
		SystemServiceTag:  x.SerialNumber,
	}

	// Enrich BIOS data with system memory speed, size, model, and processor details
	if memoryCollection, memErr := c.GetSystemMemoryInfoHP(); memErr == nil {
		if len(memoryCollection.Oem.Hpe.MemoryList) > 0 {
			y.SysMemSpeed = fmt.Sprintf("%d MHz", memoryCollection.Oem.Hpe.MemoryList[0].BoardOperationalFrequency)
		}
	}

	if sysInfo, sysErr := c.GetSystemInfoHP(); sysErr == nil {
		if sysInfo.Memory > 0 {
			y.SysMemSize = fmt.Sprintf("%.0f GB", sysInfo.Memory)
		}
		if sysInfo.Model != "" {
			y.SystemModelName = sysInfo.Model
		}
	}

	if processors, procErr := c.GetProcessorInfoHP(); procErr == nil {
		if len(processors) > 0 {
			y.Proc1Brand = processors[0].Model
			y.Proc1NumCores = int(processors[0].TotalCores)
		}
		if len(processors) > 1 {
			y.Proc2Brand = processors[1].Model
			y.Proc2NumCores = int(processors[1].TotalCores)
		}
	}

	return y, nil
}

// GetLicenseInfoHP ... will fetch the current License Details
func (c *IloClient) GetLicenseInfoHP() (LicenseInfo, error) {

	url := c.Hostname + "/redfish/v1/Managers/1/LicenseService/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return LicenseInfo{}, err
	}

	var x LicenseInfoHP

	json.Unmarshal(resp, &x)

	_result := LicenseInfo{
		Name:        x.Name,
		LicenseKey:  x.Items[0].LicenseKey,
		LicenseType: x.Items[0].LicenseType,
	}

	return _result, nil
}

// GetPCISlotsHp ... will fetch the PCI Slots Details
func (c *IloClient) GetPCISlotsHp() ([]PCISlotsInfo, error) {

	url := c.Hostname + "/redfish/v1/Systems/1/PCISlots/"

	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x PCISlotsInfoHP

	json.Unmarshal(resp, &x)

	var _pciSlots []PCISlotsInfo

	for i := range x.Items {
		_result := PCISlotsInfo{
			Name:   x.Items[i].Name,
			Status: x.Items[i].Status.OperationalStatus[0].Status,
		}
		_pciSlots = append(_pciSlots, _result)
	}

	return _pciSlots, nil

}

// GetStorageRaidHP ... will fetch HP SmartStorage logical drive details
func (c *IloClient) GetStorageRaidHP() ([]StorageRaidDetailsDell, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var controllers MemberCountHP
	json.Unmarshal(resp, &controllers)

	var _raiddata []StorageRaidDetailsDell
	for i := range controllers.Members {
		controllerID := strings.TrimSuffix(controllers.Members[i].OdataId, "/")
		controllerNum := controllerID[strings.LastIndex(controllerID, "/")+1:]
		logicalDrivesURL := c.Hostname + controllerID + "/LogicalDrives/"
		logicalDrivesResp, _, _, err := queryData(c, "GET", logicalDrivesURL, nil)
		if err != nil {
			return nil, err
		}

		var logicalCollection MemberCountHP
		json.Unmarshal(logicalDrivesResp, &logicalCollection)

		for j := range logicalCollection.Members {
			logicalDriveURL := c.Hostname + logicalCollection.Members[j].OdataId
			logicalDriveResp, _, _, err := queryData(c, "GET", logicalDriveURL, nil)
			if err != nil {
				return nil, err
			}

			var ld SmartStorageLogicalDriveHP
			json.Unmarshal(logicalDriveResp, &ld)

			layout := ld.RAIDType
			if layout == "" {
				layout = ld.Raid
			}
			if layout == "" {
				layout = ld.FaultTolerance
			}

			capacityBytes := ""
			if ld.CapacityBytes > 0 {
				capacityBytes = fmt.Sprintf("%d", ld.CapacityBytes)
			} else if ld.CapacityMiB > 0 {
				capacityBytes = fmt.Sprintf("%d", ld.CapacityMiB*1024*1024)
			} else if ld.CapacityGB > 0 {
				capacityBytes = fmt.Sprintf("%d", ld.CapacityGB*1000*1000*1000)
			}

			drivesCount := ""
			if ld.Links.DataDrivesCount > 0 {
				drivesCount = fmt.Sprintf("%d", ld.Links.DataDrivesCount)
			} else if ld.Links.DataDrives.OdataID != "" {
				dataDrivesURL := c.Hostname + ld.Links.DataDrives.OdataID
				dataDrivesResp, _, _, err := queryData(c, "GET", dataDrivesURL, nil)
				if err == nil {
					var dataDrivesCollection MemberCountHP
					json.Unmarshal(dataDrivesResp, &dataDrivesCollection)
					drivesCount = fmt.Sprintf("%d", len(dataDrivesCollection.Members))
				}
			}

			stripeSize := ld.StripeSize
			if stripeSize == "" && ld.StripeSizeBytes > 0 {
				stripeSize = fmt.Sprintf("%d", ld.StripeSizeBytes)
			}
			if stripeSize == "" && ld.StripSizeBytes > 0 {
				stripeSize = fmt.Sprintf("%d", ld.StripSizeBytes)
			}

			raidName := ld.Name
			if ld.LogicalDriveName != "" {
				raidName = ld.LogicalDriveName
			}

			// Build composite ID: LogicalDrive.<id>:ArrayController.<controllerNum>
			compositeID := fmt.Sprintf("LogicalDrive.%s:ArrayController.%s", ld.ID, controllerNum)

			raidDevice := StorageRaidDetailsDell{
				Name:             raidName,
				Id:               compositeID,
				Layout:           layout,
				MediaType:        ld.MediaType,
				DrivesCount:      drivesCount,
				ReadCachePolicy:  ld.ReadCachePolicy,
				CapacityBytes:    capacityBytes,
				StripeSize:       stripeSize,
				WriteCachePolicy: ld.WriteCachePolicy,
			}

			_raiddata = append(_raiddata, raidDevice)
		}
	}

	return _raiddata, nil
}

// GetStorageDriveDetailsHP ... will fetch disk drive details from all SmartStorage array controllers
// and map them to the same return type used by Dell storage drive details.
func (c *IloClient) GetStorageDriveDetailsHP() ([]StorageDriveDetailsDell, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/SmartStorage/ArrayControllers/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var controllers MemberCountHP
	json.Unmarshal(resp, &controllers)

	var _drives []StorageDriveDetailsDell
	for i := range controllers.Members {
		controllerID := strings.TrimSuffix(controllers.Members[i].OdataId, "/")
		drivesURL := c.Hostname + controllerID + "/DiskDrives/"
		drivesResp, _, _, err := queryData(c, "GET", drivesURL, nil)
		if err != nil {
			return nil, err
		}

		var driveCollection MemberCountHP
		json.Unmarshal(drivesResp, &driveCollection)

		controllerNum := controllerID[strings.LastIndex(controllerID, "/")+1:]

		for j := range driveCollection.Members {
			driveURL := c.Hostname + driveCollection.Members[j].OdataId
			driveResp, _, _, err := queryData(c, "GET", driveURL, nil)
			if err != nil {
				return nil, err
			}

			var raw SmartStorageDiskDriveHP
			json.Unmarshal(driveResp, &raw)

			// Build composite ID using sequential ID and controller port
			// HP Location field is formatted as ControllerPort:Box:Bay (e.g. "1I:3:4")
			// We use the sequential ID instead of physical bay number for consistency
			compositeID := raw.ID
			if raw.Location != "" {
				locationParts := strings.Split(raw.Location, ":")
				if len(locationParts) == 3 {
					// locationParts[0]=ControllerPort (e.g., "1I", "2I")
					compositeID = fmt.Sprintf("Disk.Bay.%s:ControllerPort.%s:ArrayController.%s", raw.ID, locationParts[0], controllerNum)
				}
			}

			capacityBytes := int(raw.CapacityLogicalBlocks * raw.BlockSizeBytes)
			drive := StorageDriveDetailsDell{
				ID:             compositeID,
				Name:           raw.Name,
				Description:    raw.Description,
				BlockSizeBytes: int(raw.BlockSizeBytes),
				CapacityBytes:  capacityBytes,
				MediaType:      raw.MediaType,
				Model:          raw.Model,
				PartNumber:     "",
				Revision:       raw.FirmwareVersion.Current.VersionString,
				Manufacturer:   "",
			}
			drive.Status.Health = raw.Status.Health
			drive.Status.State = raw.Status.State

			_drives = append(_drives, drive)
		}
	}

	return _drives, nil
}

// GetEthernetInterfacesHP ... will fetch the EthernetInterfaces Details
func (c *IloClient) GetEthernetInterfacesHP() ([]MACData, error) {
	url := c.Hostname + "/redfish/v1/Systems/1/EthernetInterfaces/"
	resp, _, _, err := queryData(c, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	var x MemberCountHP
	json.Unmarshal(resp, &x)

	var _macData []MACData
	for i := range x.Members {
		var y GetMacAddressHP
		_url := c.Hostname + x.Members[i].OdataId
		resp, _, _, err := queryData(c, "GET", _url, nil)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(resp, &y)

		_result := MACData{
			Name:        "EthernetInterface-" + y.Id,
			Description: "Null",
			MacAddress:  y.MACAddress,
			State:       y.Status.State,
			Status:      y.Status.Health,
			Vlan:        "Null",
		}

		_result.UpdateEmpty()
		_macData = append(_macData, _result)
	}

	return _macData, nil
}
