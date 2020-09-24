package api

//L2ConnectionResponse get connection by uuid response
type L2ConnectionResponse struct {
	UUID                string                       `json:"uuid,omitempty"`
	Name                string                       `json:"name,omitempty"`
	SellerServiceUUID   string                       `json:"sellerServiceUUID,omitempty"`
	Speed               int                          `json:"speed,omitempty"`
	SpeedUnit           string                       `json:"speedUnit,omitempty"`
	Status              string                       `json:"status,omitempty"`
	ProviderStatus      string                       `json:"providerStatus,omitempty"`
	Notifications       []string                     `json:"notifications"`
	PurchaseOrderNumber string                       `json:"purchaseOrderNumber"`
	PortUUID            string                       `json:"portUUID,omitempty"`
	VirtualDeviceUUID   string                       `json:"virtualDeviceUUID,omitempty"`
	VlanSTag            int                          `json:"vlanSTag,omitempty"`
	VlanCTag            int                          `json:"vlanCTag,omitempty"`
	NamedTag            string                       `json:"namedTag,omitempty"`
	AdditionalInfo      []L2ConnectionAdditionalInfo `json:"additionalInfo,omitempty"`
	ZSidePortUUID       string                       `json:"zSidePortUUID,omitempty"`
	ZSideVlanCTag       int                          `json:"zSideVlanCTag,omitempty"`
	ZSideVlanSTag       int                          `json:"zSideVlanSTag,omitempty"`
	SellerRegion        string                       `json:"sellerRegion,omitempty"`
	SellerMetroCode     string                       `json:"sellerMetroCode,omitempty"`
	AuthorizationKey    string                       `json:"authorizationKey,omitempty"`
	RedundantUUID       string                       `json:"redundantUUID,omitempty"`
	RedundancyType      string                       `json:"redundancyType,omitempty"`
}

//DeleteL2ConnectionResponse l2 connection delete response
type DeleteL2ConnectionResponse struct {
	Message             string `json:"message,omitempty"`
	PrimaryConnectionID string `json:"primaryConnectionId,omitempty"`
}

//L2ConnectionRequest post l2 connections request
type L2ConnectionRequest struct {
	PrimaryName                string                       `json:"primaryName,omitempty"`
	ProfileUUID                string                       `json:"profileUUID,omitempty"`
	Speed                      int                          `json:"speed,omitempty"`
	SpeedUnit                  string                       `json:"speedUnit,omitempty"`
	Notifications              []string                     `json:"notifications"`
	PurchaseOrderNumber        string                       `json:"purchaseOrderNumber"`
	PrimaryPortUUID            string                       `json:"primaryPortUUID,omitempty"`
	VirtualDeviceUUID          string                       `json:"virtualDeviceUUID,omitempty"`
	PrimaryVlanSTag            int                          `json:"primaryVlanSTag,omitempty"`
	PrimaryVlanCTag            int                          `json:"primaryVlanCTag,omitempty"`
	NamedTag                   string                       `json:"namedTag,omitempty"`
	AdditionalInfo             []L2ConnectionAdditionalInfo `json:"additionalInfo,omitempty"`
	PrimaryZSidePortUUID       string                       `json:"primaryZSidePortUUID,omitempty"`
	PrimaryZSideVlanSTag       int                          `json:"primaryZSideVlanSTag,omitempty"`
	PrimaryZSideVlanCTag       int                          `json:"primaryZSideVlanCTag,omitempty"`
	SecondaryName              string                       `json:"secondaryName,omitempty"`
	SecondaryPortUUID          string                       `json:"secondaryPortUUID,omitempty"`
	SecondaryVirtualDeviceUUID string                       `json:"secondaryVirtualDeviceUUID,omitempty"`
	SecondaryVlanSTag          int                          `json:"secondaryVlanSTag,omitempty"`
	SecondaryVlanCTag          int                          `json:"secondaryVlanCTag,omitempty"`
	SecondaryZSidePortUUID     string                       `json:"secondaryZSidePortUUID,omitempty"`
	SecondaryZSideVlanSTag     int                          `json:"secondaryZSideVlanSTag,omitempty"`
	SecondaryZSideVlanCTag     int                          `json:"secondaryZSideVlanCTag,omitempty"`
	SellerRegion               string                       `json:"sellerRegion,omitempty"`
	SellerMetroCode            string                       `json:"sellerMetroCode,omitempty"`
	AuthorizationKey           string                       `json:"authorizationKey,omitempty"`
}

//CreateL2ConnectionResponse post l2 connection response
type CreateL2ConnectionResponse struct {
	Message               string `json:"message,omitempty"`
	PrimaryConnectionID   string `json:"primaryConnectionId,omitempty"`
	SecondaryConnectionID string `json:"secondaryConnectionId,omitempty"`
	Status                string `json:"status,omitempty"`
}

//L2ConnectionAdditionalInfo additional info object used in L2 connections
type L2ConnectionAdditionalInfo struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

//L2ConnectionUpdateRequest describes layer2 connection update request
type L2ConnectionUpdateRequest struct {
	Name      string `json:"connectionNewName,omitempty"`
	Speed     int    `json:"speed,omitempty"`
	SpeedUnit string `json:"speedUnit,omitempty"`
}

//L2ConnectionUpdateRequest describes layer2 connection update response
type L2ConnectionUpdateResponse struct {
	Message             string `json:"message,omitempty"`
	PrimaryConnectionID string `json:"primaryConnectionId,omitempty"`
	Status              string `json:"status,omitempty"`
}
