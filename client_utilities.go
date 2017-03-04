package main

import (
    "fmt"
    "encoding/json"
)

type ServerState struct{
	SchemaName string `json:"schemaName"`
	TenantId string  `json:"tenantId"`
}

type Body struct{
	ServerState ServerState `json:"serverState"`
	Platform string`json:"platform"`
	DeviceId string `json:"deviceId"`
	DeviceType string `json:"deviceType"`
	PushToken string `json:"pushToken"`
	Version string `json:"version"`
	PublicKey string `json:"publicKey"`
}

type RegistrationResponse struct{
	VocId string `json:"vocId"`
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	DailyDownloadCellular int `json:"dailyDownloadCellular"`
	DailyDownloadWifi int `json:"dailyDownloadWifi"`
	PlayAds bool `json:"playAds"`
	SkipPolicyFirstTime bool `json:"skipPolicyFirstTime"`
	DisplayInProgressVideos bool `json:"displayInProgressVideos"`
	SdkParameters struct {
		Type string `json:"type"`
	} `json:"sdkParameters"`
}

type VocInfo struct{
  VocId string
  AccessToken string
  RefreshToken string
}

type StatusBody struct{
	ServerState ServerState `json:"serverState"`
    AccessToken string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
	VocId string `json:"vocId"`
    DeviceStatus struct{
      Charger bool `json:"charger"`
	} `json:"deviceStatus"`
}

type ManifestBody struct{
    ServerState ServerState `json:"serverState"`
    AccessToken string `json:"accessToken"`
    RefreshToken string `json:"refreshToken"`
    VocId string `json:"vocId"`
}

func build_reg_json(schema string, tenant string, pubkey string) string{

    bodyD := &Body{
		PublicKey : pubkey,
		Platform: "Linux",
		DeviceId: "623bce38-a1f4-11e6-bb6c-3417eb9985a6",
		DeviceType: "pc",
		PushToken: "tt",
		Version: "18.1.2",
		ServerState: ServerState{
		    SchemaName:   schema,
		    TenantId: tenant},
		}
    bodyB, _ := json.Marshal(bodyD)
    fmt.Println(string(bodyB))
    return  string(bodyB)
}

func build_status_json(schema string, tenant string, voc_info VocInfo)string{

    bodyD := &StatusBody{
		VocId: voc_info.VocId,
        AccessToken : voc_info.AccessToken,
		ServerState: ServerState{
		    SchemaName:   schema,
		    TenantId: tenant},
		}
    bodyB, _ := json.Marshal(bodyD)
    fmt.Println(string(bodyB))
    return  string(bodyB)
}

func build_manifest_json(schema string, tenant string, voc_info VocInfo) string{
     bodyD := &ManifestBody{
               VocId: voc_info.VocId,
	       AccessToken : voc_info.AccessToken,
	               ServerState: ServerState{
                           SchemaName:   schema,
                           TenantId: tenant},
                       }
    bodyB, _ := json.Marshal(bodyD)
    fmt.Println(string(bodyB))
    return  string(bodyB)                 
}