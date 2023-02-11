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
		Type    string `json:"type"`
		Devices []struct {
			Id                       int     `json:"id"`
			UserId                   int     `json:"user_id"`
			SerialNumber             string  `json:"serial_number"`
			AquariumTankId           int     `json:"aquarium_tank_id"`
			FriendlyName             string  `json:"friendly_name"`
			FirmwareVersion          string  `json:"firmware_version"`
			NextFirmwareVersion      string  `json:"next_firmware_version"`
			McuVersion               *string `json:"mcu_version"`
			BatchNo                  int     `json:"batch_no"`
			IsAutoUpdate             int     `json:"is_auto_update,omitempty"`
			LastOnline               int     `json:"last_online"`
			IsActive                 int     `json:"is_active"`
			IsBetaTest               bool    `json:"is_beta_test,omitempty"`
			LocalIpAddress           string  `json:"local_ip_address"`
			LastCalibrateA           int     `json:"last_calibrate_a,omitempty"`
			LastCalibrateC           int     `json:"last_calibrate_c,omitempty"`
			LastCalibrateD           int     `json:"last_calibrate_d,omitempty"`
			LastCalibrateP           int     `json:"last_calibrate_p,omitempty"`
			LastResetTestCounterTime int     `json:"last_reset_test_counter_time,omitempty"`
			LastCloudAdvRecordTime   int     `json:"last_cloud_adv_record_time,omitempty"`
			LastCloudAdvUploadTime   int     `json:"last_cloud_adv_upload_time,omitempty"`
			Settings                 struct {
				NextTestTime            int    `json:"next_test_time,omitempty"`
				UpperKh                 int    `json:"upper_kh,omitempty"`
				LowerKh                 int    `json:"lower_kh,omitempty"`
				MeasureInterval         int    `json:"measure_interval,omitempty"`
				AquariumVolume          int    `json:"aquarium_volume,omitempty"`
				ReagentAlert            int    `json:"reagent_alert,omitempty"`
				DoseBufferAlert         int    `json:"dose_buffer_alert,omitempty"`
				KhAlert                 int    `json:"kh_alert,omitempty"`
				BaselineCalibration     int    `json:"baseline_calibration,omitempty"`
				IsActionModeOn          int    `json:"is_action_mode_on,omitempty"`
				IsWashoutModeOn         int    `json:"is_washout_mode_on,omitempty"`
				IsDosetronicModeOn      int    `json:"is_dosetronic_mode_on,omitempty"`
				IsFastModeOn            int    `json:"is_fast_mode_on,omitempty"`
				CorrectionMethod        int    `json:"correction_method,omitempty"`
				RetestDuration          int    `json:"retest_duration,omitempty"`
				PumpDAction             string `json:"pump_d_action,omitempty"`
				SettingsUpdateTime      int    `json:"settings_update_time,omitempty"`
				CurrentTestCount        int    `json:"current_test_count,omitempty"`
				TestCountLimit          int    `json:"test_count_limit,omitempty"`
				AlkatronicSerialNumber  string `json:"alkatronic_serial_number,omitempty"`
				ContinuousModeDoseLimit int    `json:"continuous_mode_dose_limit,omitempty"`
				Pumps                   []struct {
					Id                        int    `json:"id"`
					Name                      string `json:"name"`
					MaxVolume                 int    `json:"max_volume"`
					RemainingVolume           int    `json:"remaining_volume"`
					IsAlkatronicModeOn        int    `json:"is_alkatronic_mode_on"`
					IsContinuousModeOn        int    `json:"is_continuous_mode_on"`
					IsDoseAlertOn             int    `json:"is_dose_alert_on"`
					IsRefillAlertOn           int    `json:"is_refill_alert_on"`
					LastReplaceHoseTime       int    `json:"last_replace_hose_time"`
					DailyVolume               int    `json:"daily_volume"`
					Last24HourTotalDoseVolume int    `json:"last_24_hour_total_dose_volume"`
					HasSchedule               int    `json:"has_schedule"`
				} `json:"pumps,omitempty"`
			} `json:"settings"`
			Versions struct {
				IsAllowCloudSettings                   bool `json:"is_allow_cloud_settings"`
				IsAllowDosetronicSettings              bool `json:"is_allow_dosetronic_settings"`
				IsAllowKBncSettings                    bool `json:"is_allow_k_bnc_settings"`
				IsAllowMtestPhtestRefillSettings       bool `json:"is_allow_mtest_phtest_refill_settings"`
				IsAllowSdaOdaRetestSettings            bool `json:"is_allow_sda_oda_retest_settings"`
				IsAllowPlugSettings                    bool `json:"is_allow_plug_settings"`
				IsUseRecordFormatV2                    bool `json:"is_use_record_format_v2"`
				IsAllowResetPhProbeCalibrationSettings bool `json:"is_allow_reset_ph_probe_calibration_settings"`
				IsAllowSuperLowReferenceSettings       bool `json:"is_allow_super_low_reference_settings"`
			} `json:"versions,omitempty"`
			EligibleActions []struct {
				Key     string `json:"key"`
				Options struct {
					AquariumVolumeUpperLimit int    `json:"aquarium_volume_upper_limit,omitempty"`
					LowReferenceRangeFrom    int    `json:"low_reference_range_from,omitempty"`
					LowReferenceRangeTo      int    `json:"low_reference_range_to,omitempty"`
					HighReferenceRangeFrom   int    `json:"high_reference_range_from,omitempty"`
					HighReferenceRangeTo     int    `json:"high_reference_range_to,omitempty"`
					HasKValueSettings        int    `json:"has_k_value_settings,omitempty"`
					IsRequirePasscode        int    `json:"is_require_passcode,omitempty"`
					TargetVersion            string `json:"target_version,omitempty"`
					InstructionLink          string `json:"instruction_link,omitempty"`
				} `json:"options,omitempty"`
			} `json:"eligible_actions"`
			LastAlkatronicAdvTime int    `json:"last_alkatronic_adv_time,omitempty"`
			IsAdvActive           int    `json:"is_adv_active,omitempty"`
			NextMcuVersion        string `json:"next_mcu_version,omitempty"`
			Parameters            []struct {
				Parameter              string `json:"parameter"`
				LatestRecord           int    `json:"latest_record"`
				LowReference           int    `json:"low_reference"`
				HighReference          int    `json:"high_reference"`
				MultiplyFactor         int    `json:"multiply_factor"`
				IsActionModeOn         int    `json:"is_action_mode_on"`
				IsAutomaticModeOn      int    `json:"is_automatic_mode_on"`
				BaselineDeviceValue    int    `json:"baseline_device_value"`
				BaselineReferenceValue int    `json:"baseline_reference_value"`
				RecordTime             int    `json:"record_time"`
			} `json:"parameters,omitempty"`
		} `json:"devices"`
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
		RecordID             int         `json:"record_id"`
		DeviceID             int         `json:"device_id"`
		KhValue              float64     `json:"kh_value"`
		PhValue              int         `json:"ph_value"`
		SolutionAdded        float64     `json:"solution_added"`
		AcidUsed             int         `json:"acid_used"`
		IsPowerPlugOn        int         `json:"is_power_plug_on"`
		Indicator            int         `json:"indicator"`
		RemainingReagent     int         `json:"remaining_reagent"`
		RemainingDoseBuffer  int         `json:"remaining_dose_buffer"`
		TestCount            int         `json:"test_count"`
		IsHidden             int         `json:"is_hidden"`
		Note                 interface{} `json:"note"`
		IsDeleted            int         `json:"is_deleted"`
		LocalHour            interface{} `json:"local_hour"`
		LocalMinute          interface{} `json:"local_minute"`
		RecordTime           int64       `json:"record_time"`
		CreateTime           int64       `json:"create_time"`
		UpdateTime           int         `json:"update_time"`
		Db1                  interface{} `json:"db1"`
		HumanRecordTime      string      `json:"human_record_time"`
		Id                   int         `json:"id"`
		Parameter            string      `json:"parameter"`
		TestProfileId        int         `json:"test_profile_id"`
		TestProfileIndicator string      `json:"test_profile_indicator"`
		LowerBound           int         `json:"lower_bound"`
		UpperBound           int         `json:"upper_bound"`
		Value                float64     `json:"value"`
		BaselinedValue       int         `json:"baselined_value"`
		MultiplyFactor       int         `json:"multiply_factor"`
	} `json:"data"`
}

type Record struct {
	RecordID            int         `json:"record_id"`
	DeviceID            int         `json:"device_id"`
	KhValue             float64     `json:"kh_value"`
	PhValue             int         `json:"ph_value"`
	SolutionAdded       float64     `json:"solution_added"`
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
	RecordTime          int64       `json:"record_time"`
	CreateTime          int64       `json:"create_time"`
	UpdateTime          int         `json:"update_time"`
	Db1                 interface{} `json:"db1"`
	HumanRecordTime     string      `json:"human_record_time"`
}

type AlkatronicDevice struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    struct {
		ID                       int         `json:"id"`
		UserID                   int         `json:"user_id"`
		SerialNumber             string      `json:"serial_number"`
		AquariumTankID           int         `json:"aquarium_tank_id"`
		FriendlyName             string      `json:"friendly_name"`
		FirmwareVersion          string      `json:"firmware_version"`
		NextFirmwareVersion      string      `json:"next_firmware_version"`
		McuVersion               interface{} `json:"mcu_version"`
		BatchNo                  int         `json:"batch_no"`
		IsAutoUpdate             int         `json:"is_auto_update"`
		LastOnline               int         `json:"last_online"`
		IsActive                 int         `json:"is_active"`
		IsBetaTest               bool        `json:"is_beta_test"`
		LocalIPAddress           string      `json:"local_ip_address"`
		MacAddress               interface{} `json:"mac_address"`
		LastCalibrateA           int         `json:"last_calibrate_a"`
		LastCalibrateC           int         `json:"last_calibrate_c"`
		LastCalibrateD           int         `json:"last_calibrate_d"`
		LastCalibrateP           int         `json:"last_calibrate_p"`
		LastResetTestCounterTime int         `json:"last_reset_test_counter_time"`
		LastCloudAdvRecordTime   int         `json:"last_cloud_adv_record_time"`
		LastCloudAdvUploadTime   int         `json:"last_cloud_adv_upload_time"`
		Settings                 struct {
			NextTestTime        int    `json:"next_test_time"`
			UpperKh             int    `json:"upper_kh"`
			LowerKh             int    `json:"lower_kh"`
			MeasureInterval     int    `json:"measure_interval"`
			AquariumVolume      int    `json:"aquarium_volume"`
			ReagentAlert        int    `json:"reagent_alert"`
			DoseBufferAlert     int    `json:"dose_buffer_alert"`
			KhAlert             int    `json:"kh_alert"`
			BaselineCalibration int    `json:"baseline_calibration"`
			IsActionModeOn      int    `json:"is_action_mode_on"`
			IsWashoutModeOn     int    `json:"is_washout_mode_on"`
			IsDosetronicModeOn  int    `json:"is_dosetronic_mode_on"`
			IsFastModeOn        int    `json:"is_fast_mode_on"`
			CorrectionMethod    int    `json:"correction_method"`
			RetestDuration      int    `json:"retest_duration"`
			PumpDAction         string `json:"pump_d_action"`
			SettingsUpdateTime  int    `json:"settings_update_time"`
			CurrentTestCount    int    `json:"current_test_count"`
			TestCountLimit      int    `json:"test_count_limit"`
		} `json:"settings"`
		Versions struct {
			IsAllowCloudSettings                   bool `json:"is_allow_cloud_settings"`
			IsAllowDosetronicSettings              bool `json:"is_allow_dosetronic_settings"`
			IsAllowKBncSettings                    bool `json:"is_allow_k_bnc_settings"`
			IsAllowMtestPhtestRefillSettings       bool `json:"is_allow_mtest_phtest_refill_settings"`
			IsAllowSdaOdaRetestSettings            bool `json:"is_allow_sda_oda_retest_settings"`
			IsAllowPlugSettings                    bool `json:"is_allow_plug_settings"`
			IsUseRecordFormatV2                    bool `json:"is_use_record_format_v2"`
			IsAllowResetPhProbeCalibrationSettings bool `json:"is_allow_reset_ph_probe_calibration_settings"`
			IsAllowSuperLowReferenceSettings       bool `json:"is_allow_super_low_reference_settings"`
		} `json:"versions"`
		EligibleActions []struct {
			Key     string `json:"key"`
			Options struct {
				AquariumVolumeUpperLimit int `json:"aquarium_volume_upper_limit"`
				LowReferenceRangeFrom    int `json:"low_reference_range_from"`
				LowReferenceRangeTo      int `json:"low_reference_range_to"`
				HighReferenceRangeFrom   int `json:"high_reference_range_from"`
				HighReferenceRangeTo     int `json:"high_reference_range_to"`
				HasKValueSettings        int `json:"has_k_value_settings"`
			} `json:"options,omitempty"`
		} `json:"eligible_actions"`
	} `json:"data"`
}

type MastertronicDevice struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
	Data    struct {
		ID                         int         `json:"id"`
		UserID                     int         `json:"user_id"`
		SerialNumber               string      `json:"serial_number"`
		AquariumTankID             int         `json:"aquarium_tank_id"`
		FriendlyName               string      `json:"friendly_name"`
		FirmwareVersion            string      `json:"firmware_version"`
		NextFirmwareVersion        string      `json:"next_firmware_version"`
		McuVersion                 string      `json:"mcu_version"`
		NextMcuVersion             string      `json:"next_mcu_version"`
		BatchNo                    int         `json:"batch_no"`
		LastOnline                 int         `json:"last_online"`
		IsActive                   int         `json:"is_active"`
		LocalIPAddress             string      `json:"local_ip_address"`
		MacAddress                 interface{} `json:"mac_address"`
		LifetimeTestCount          int         `json:"lifetime_test_count"`
		LastResetHoseCounterTime   int         `json:"last_reset_hose_counter_time"`
		LastResetNeedleCounterTime int         `json:"last_reset_needle_counter_time"`
		McuStatus                  string      `json:"mcu_status"`
		Parameters                 []struct {
			Parameter              string `json:"parameter"`
			LatestRecord           int    `json:"latest_record"`
			LowReference           int    `json:"low_reference"`
			HighReference          int    `json:"high_reference"`
			MultiplyFactor         int    `json:"multiply_factor"`
			IsActionModeOn         int    `json:"is_action_mode_on"`
			IsAutomaticModeOn      int    `json:"is_automatic_mode_on"`
			BaselineDeviceValue    int    `json:"baseline_device_value"`
			BaselineReferenceValue int    `json:"baseline_reference_value"`
			RecordTime             int    `json:"record_time"`
		} `json:"parameters"`
		Settings struct {
			CurrentHoseCount   int `json:"current_hose_count"`
			CurrentNeedleCount int `json:"current_needle_count"`
			HoseCountLimit     int `json:"hose_count_limit"`
			NeedleCountLimit   int `json:"needle_count_limit"`
		} `json:"settings"`
	} `json:"data"`
}

type DosetronicRecord struct {
	PumpID          int     `json:"pump_id"`
	DoseVolume      float64 `json:"dose_volume"`
	RemainingVolume float64 `json:"remaining_volume"`
	DoseMode        int     `json:"dose_mode"`
	RecordTime      int64   `json:"record_time"`
}

type DosetronicRecords struct {
	Result  bool                 `json:"result"`
	Message string               `json:"message"`
	Data    [][]DosetronicRecord `json:"data"`
}

type MastertronicRecords struct {
	Result  bool                 `json:"result"`
	Message string               `json:"message"`
	Data    []MastertronicRecord `json:"data"`
}

type MastertronicRecord struct {
	ID                   int         `json:"id"`
	Parameter            string      `json:"parameter"`
	TestProfileID        int         `json:"test_profile_id"`
	TestProfileIndicator string      `json:"test_profile_indicator"`
	LowerBound           int         `json:"lower_bound"`
	UpperBound           int         `json:"upper_bound"`
	Value                float64     `json:"value"`
	BaselinedValue       int         `json:"baselined_value"`
	MultiplyFactor       int         `json:"multiply_factor"`
	Indicator            int         `json:"indicator"`
	Note                 interface{} `json:"note"`
	IsHidden             int         `json:"is_hidden"`
	RecordTime           int64       `json:"record_time"`
}

type AlkatronicRecords struct {
	Result  bool               `json:"result"`
	Message string             `json:"message"`
	Data    []AlkatronicRecord `json:"data"`
}

type AlkatronicRecord struct {
	Type          string      `json:"type"`
	KhValue       float64     `json:"kh_value"`
	PhValue       int         `json:"ph_value"`
	SolutionAdded float64     `json:"solution_added"`
	AcidUsed      int         `json:"acid_used"`
	IsPowerPlugOn int         `json:"is_power_plug_on"`
	Indicator     int         `json:"indicator"`
	IsHidden      int         `json:"is_hidden"`
	Note          interface{} `json:"note"`
	RecordTime    int64       `json:"record_time"`
	CreateTime    int64       `json:"create_time"`
}
