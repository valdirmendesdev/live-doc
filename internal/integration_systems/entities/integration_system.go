package entities

import "errors"

type SystemType string

const (
	SAP   SystemType = "SAP"
	Other SystemType = "Other"
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
	EccHttp  CommunicationType = "Http"
	CPI      CommunicationType = "CPI"
	PI       CommunicationType = "PI"
	PO       CommunicationType = "PO"
	External CommunicationType = "External"
)

type IntegrationSystem struct {
	SystemType                    SystemType
	Name                          string
	CommunicationType             CommunicationType
	CommunicationName             string
	SAPVersion                    string
	SapBasis                      string
	SapAppl                       string
	RequestCreationAutonomy       AutonomyLevel
	RequestTransportationAutonomy AutonomyLevel
	HasSAPNfeImplementation       bool
	NfeBadiImplemented            bool
	DanfeLayoutImplemented        bool
}

func NewIntegrationSystem(integrationType SystemType, communicationType CommunicationType) IntegrationSystem {
	return IntegrationSystem{
		SystemType:        integrationType,
		CommunicationType: communicationType,
	}
}

func (is *IntegrationSystem) IsValid() (bool, error) {
	if is.SystemType == "" {
		return false, errors.New("system type must be filled")
	}

	if is.SystemType != SAP && is.SystemType != Other {
		return false, errors.New("system type is not valid")
	}

	if is.CommunicationType == "" {
		return false, errors.New("communication type must be filled")
	}

	if is.CommunicationType != EccHttp &&
		is.CommunicationType != CPI &&
		is.CommunicationType != PI &&
		is.CommunicationType != PO &&
		is.CommunicationType != External {
		return false, errors.New("communication type is not valid")
	}

	return true, nil
}
