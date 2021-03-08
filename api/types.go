package api

type Response struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type Devices struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    []struct {
		DeviceID                               int         `json:"device_id"`
		SerialNumber                           string      `json:"serial_number"`
		BatchNo                                int         `json:"batch_no"`
		FirmwareVersion                        string      `json:"firmware_version"`
		NextFirmwareVersion                    interface{} `json:"next_firmware_version"`
		NextTestTime                           int         `json:"next_test_time"`
		UserID                                 int         `json:"user_id"`
		FriendlyName                           string      `json:"friendly_name"`
		UpperKh                                float64     `json:"upper_kh"`
		LowerKh                                float64     `json:"lower_kh"`
		MeasureInterval                        int         `json:"measure_interval"`
		AquariumVolume                         int         `json:"aquarium_volume"`
		ReagentAlert                           int         `json:"reagent_alert"`
		DoseBufferAlert                        int         `json:"dose_buffer_alert"`
		KhAlert                                int         `json:"kh_alert"`
		BaselineCalibration                    int         `json:"baseline_calibration"`
		IsAutoUpdate                           int         `json:"is_auto_update"`
		IsActionModeOn                         int         `json:"is_action_mode_on"`
		IsWashoutModeOn                        int         `json:"is_washout_mode_on"`
		IsDosetronicModeOn                     int         `json:"is_dosetronic_mode_on"`
		IsFastModeOn                           int         `json:"is_fast_mode_on"`
		CorrectionMethod                       int         `json:"correction_method"`
		RetestDuration                         int         `json:"retest_duration"`
		SettingsUpdateTime                     int         `json:"settings_update_time"`
		LastCalibrateA                         int         `json:"last_calibrate_a"`
		LastCalibrateC                         int         `json:"last_calibrate_c"`
		LastCalibrateD                         int         `json:"last_calibrate_d"`
		LastCalibrateP                         int         `json:"last_calibrate_p"`
		LastOnline                             int         `json:"last_online"`
		IsActive                               int         `json:"is_active"`
		IsBetaTest                             bool        `json:"is_beta_test"`
		TimeZone                               interface{} `json:"time_zone"`
		LocalIPAddress                         interface{} `json:"local_ip_address"`
		MonitorBy                              interface{} `json:"monitor_by"`
		UnlinkToken                            interface{} `json:"unlink_token"`
		CreateTime                             int         `json:"create_time"`
		UpdateTime                             int         `json:"update_time"`
		IsAllowCloudSettings                   bool        `json:"is_allow_cloud_settings"`
		IsAllowDosetronicSettings              bool        `json:"is_allow_dosetronic_settings"`
		IsAllowKBncSettings                    bool        `json:"is_allow_k_bnc_settings"`
		IsAllowMtestPhtestRefillSettings       bool        `json:"is_allow_mtest_phtest_refill_settings"`
		IsAllowSdaOdaRetestSettings            bool        `json:"is_allow_sda_oda_retest_settings"`
		IsAllowPlugSettings                    bool        `json:"is_allow_plug_settings"`
		IsUseRecordFormatV2                    bool        `json:"is_use_record_format_v2"`
		IsAllowResetPhProbeCalibrationSettings bool        `json:"is_allow_reset_ph_probe_calibration_settings"`
		IsAllowSuperLowReferenceSettings       bool        `json:"is_allow_super_low_reference_settings"`
	} `json:"data"`
}

type Device struct {
	DeviceID                               int         `json:"device_id"`
	SerialNumber                           string      `json:"serial_number"`
	BatchNo                                int         `json:"batch_no"`
	FirmwareVersion                        string      `json:"firmware_version"`
	NextFirmwareVersion                    interface{} `json:"next_firmware_version"`
	NextTestTime                           int         `json:"next_test_time"`
	UserID                                 int         `json:"user_id"`
	FriendlyName                           string      `json:"friendly_name"`
	UpperKh                                float64     `json:"upper_kh"`
	LowerKh                                float64     `json:"lower_kh"`
	MeasureInterval                        int         `json:"measure_interval"`
	AquariumVolume                         int         `json:"aquarium_volume"`
	ReagentAlert                           int         `json:"reagent_alert"`
	DoseBufferAlert                        int         `json:"dose_buffer_alert"`
	KhAlert                                int         `json:"kh_alert"`
	BaselineCalibration                    int         `json:"baseline_calibration"`
	IsAutoUpdate                           int         `json:"is_auto_update"`
	IsActionModeOn                         int         `json:"is_action_mode_on"`
	IsWashoutModeOn                        int         `json:"is_washout_mode_on"`
	IsDosetronicModeOn                     int         `json:"is_dosetronic_mode_on"`
	IsFastModeOn                           int         `json:"is_fast_mode_on"`
	CorrectionMethod                       int         `json:"correction_method"`
	RetestDuration                         int         `json:"retest_duration"`
	SettingsUpdateTime                     int         `json:"settings_update_time"`
	LastCalibrateA                         int         `json:"last_calibrate_a"`
	LastCalibrateC                         int         `json:"last_calibrate_c"`
	LastCalibrateD                         int         `json:"last_calibrate_d"`
	LastCalibrateP                         int         `json:"last_calibrate_p"`
	LastOnline                             int         `json:"last_online"`
	IsActive                               int         `json:"is_active"`
	IsBetaTest                             bool        `json:"is_beta_test"`
	TimeZone                               interface{} `json:"time_zone"`
	LocalIPAddress                         interface{} `json:"local_ip_address"`
	MonitorBy                              interface{} `json:"monitor_by"`
	UnlinkToken                            interface{} `json:"unlink_token"`
	CreateTime                             int         `json:"create_time"`
	UpdateTime                             int         `json:"update_time"`
	IsAllowCloudSettings                   bool        `json:"is_allow_cloud_settings"`
	IsAllowDosetronicSettings              bool        `json:"is_allow_dosetronic_settings"`
	IsAllowKBncSettings                    bool        `json:"is_allow_k_bnc_settings"`
	IsAllowMtestPhtestRefillSettings       bool        `json:"is_allow_mtest_phtest_refill_settings"`
	IsAllowSdaOdaRetestSettings            bool        `json:"is_allow_sda_oda_retest_settings"`
	IsAllowPlugSettings                    bool        `json:"is_allow_plug_settings"`
	IsUseRecordFormatV2                    bool        `json:"is_use_record_format_v2"`
	IsAllowResetPhProbeCalibrationSettings bool        `json:"is_allow_reset_ph_probe_calibration_settings"`
	IsAllowSuperLowReferenceSettings       bool        `json:"is_allow_super_low_reference_settings"`
}

type Records struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    []struct {
		RecordID            int         `json:"record_id"`
		DeviceID            int         `json:"device_id"`
		KhValue             float64     `json:"kh_value"`
		PhValue             int         `json:"ph_value"`
		SolutionAdded       int         `json:"solution_added"`
		AcidUsed            int         `json:"acid_used"`
		IsPowerPlugOn       int         `json:"is_power_plug_on"`
		Indicator           int         `json:"indicator"`
		RemainingReagent    int         `json:"remaining_reagent"`
		RemainingDoseBuffer int         `json:"remaining_dose_buffer"`
		TestCount           int         `json:"test_count"`
		IsHidden            int         `json:"is_hidden"`
		Note                interface{} `json:"note"`
		IsDeleted           int         `json:"is_deleted"`
		LocalHour           interface{} `json:"local_hour"`
		LocalMinute         interface{} `json:"local_minute"`
		RecordTime          int         `json:"record_time"`
		CreateTime          int64       `json:"create_time"`
		UpdateTime          int         `json:"update_time"`
		Db1                 interface{} `json:"db1"`
		HumanRecordTime     string      `json:"human_record_time"`
	} `json:"data"`
}

type Record struct {
	RecordID            int         `json:"record_id"`
	DeviceID            int         `json:"device_id"`
	KhValue             float64     `json:"kh_value"`
	PhValue             int         `json:"ph_value"`
	SolutionAdded       int         `json:"solution_added"`
	AcidUsed            int         `json:"acid_used"`
	IsPowerPlugOn       int         `json:"is_power_plug_on"`
	Indicator           int         `json:"indicator"`
	RemainingReagent    int         `json:"remaining_reagent"`
	RemainingDoseBuffer int         `json:"remaining_dose_buffer"`
	TestCount           int         `json:"test_count"`
	IsHidden            int         `json:"is_hidden"`
	Note                interface{} `json:"note"`
	IsDeleted           int         `json:"is_deleted"`
	LocalHour           interface{} `json:"local_hour"`
	LocalMinute         interface{} `json:"local_minute"`
	RecordTime          int         `json:"record_time"`
	CreateTime          int64       `json:"create_time"`
	UpdateTime          int         `json:"update_time"`
	Db1                 interface{} `json:"db1"`
	HumanRecordTime     string      `json:"human_record_time"`
}
