package main

// AkamaiNetworkLists object format
type AkamaiNetworkLists struct {
	NetworkLists []struct {
		UpdateEpoch                int64    `json:"updateEpoch"`
		CreateEpoch                int64    `json:"createEpoch"`
		CreateDate                 int64    `json:"createDate"`
		UpdatedBy                  string   `json:"updatedBy"`
		UpdateDate                 int64    `json:"updateDate"`
		CreatedBy                  string   `json:"createdBy"`
		ProductionActivationStatus string   `json:"productionActivationStatus"`
		StagingActivationStatus    string   `json:"stagingActivationStatus"`
		Name                       string   `json:"name"`
		Type                       string   `json:"type"`
		UniqueID                   string   `json:"unique-id"`
		List                       []string `json:"list"`
		Links                      []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		SyncPoint  int `json:"sync-point"`
		NumEntries int `json:"numEntries"`
	} `json:"network_lists"`
}

// AkamaiNetworkList single object of list
type AkamaiNetworkList struct {
	CreateEpoch                int    `json:"createEpoch"`
	UpdateEpoch                int    `json:"updateEpoch"`
	CreateDate                 int64  `json:"createDate"`
	CreatedBy                  string `json:"createdBy"`
	UpdatedBy                  string `json:"updatedBy"`
	UpdateDate                 int64  `json:"updateDate"`
	StagingActivationStatus    string `json:"stagingActivationStatus"`
	ProductionActivationStatus string `json:"productionActivationStatus"`
	Name                       string `json:"name"`
	Type                       string `json:"type"`
	UniqueID                   string `json:"unique-id"`
	Account                    string `json:"account"`
	ReadOnly                   bool   `json:"readOnly"`
	SyncPoint                  int    `json:"sync-point"`
	Links                      []struct {
		Rel  string `json:"rel"`
		Href string `json:"href"`
	} `json:"links"`
	List       []string `json:"list"`
	NumEntries int      `json:"numEntries"`
}

// SingleAkamaiNetworkList object to create new network list
type SingleAkamaiNetworkList struct {
	Name        string   `json:"name,omitempty"`
	Type        string   `json:"type,omitempty"`
	AcgID       string   `json:"acgId,omitempty"`
	ContractID  string   `json:"contractId,omitempty"`
	GroupID     int64    `json:"groupId,omitempty"`
	Description string   `json:"description,omitempty"`
	List        []string `json:"list"`
}

type ActivationNetworkList struct {
	SiebelTicketID         string   `json:"siebel-ticket-id"`
	NotificationRecipients []string `json:"notification-recipients"`
	Comments               string   `json:"comments"`
}
