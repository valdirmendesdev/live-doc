package models

type Type string

const (
	SAP   Type = "SAP"
	Other Type = "Other"
)

type AutonomyLevel string

const (
	None            AutonomyLevel = "None"
	ExtremeControl  AutonomyLevel = "ExtremeControl"
	ModerateControl AutonomyLevel = "ModerateControl"
	FullAutonomy    AutonomyLevel = "FullAutonomy"
)

type CommunicationType string

const (
	EccHttp  CommunicationType = "ECC_HTTP"
	CPI      CommunicationType = "CPI"
	PI       CommunicationType = "PI"
	PO       CommunicationType = "PO"
	External CommunicationType = "External"
)

type IntegrationSystem struct {
	Type                          Type
	Name                          string
	SAPVersion                    string
	SapBasis                      string
	SapAppl                       string
	RequestCreationAutonomy       AutonomyLevel
	RequestTransportationAutonomy AutonomyLevel
	HasSAPNfeImplementation       bool
	NfeBadiImplemented            bool
	DanfeLayoutImplemented        bool
	CommunicationType             CommunicationType
	CommunicationName             string
}
