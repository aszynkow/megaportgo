// Copyright 2020 Megaport Pty Ltd
//
// Licensed under the Mozilla Public License, Version 2.0 (the
// "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//       https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

type VXCOrder struct {
	AssociatedVXCs []VXCOrderConfiguration `json:"associatedVxcs"`
	PortID         string                  `json:"productUid"`
}
type VXCOrderConfiguration struct {
	Name      string                    `json:"productName"`
	RateLimit int                       `json:"rateLimit"`
	AEnd      VXCOrderAEndConfiguration `json:"aEnd"`
	BEnd      VXCOrderBEndConfiguration `json:"bEnd"`
}

type VXCOrderAEndConfiguration struct {
	VLAN          int                       `json:"vlan"`
	PartnerConfig VXCOrderAEndPartnerConfig `json:"partnerConfig,omitempty"`
}

type VXCOrderAEndPartnerConfig struct {
	Interfaces []PartnerConfigInterface `json:"interfaces,omitempty"`
}

type VXCOrderBEndConfiguration struct {
	VLAN       int    `json:"vlan,omitempty"`
	ProductUID string `json:"productUid"`
}

type VXCOrderConfirmation struct {
	TechnicalServiceUID string `json:"vxcJTechnicalServiceUid"`
}

// BGP CONFIG STUFF

type PartnerConfigInterface struct {
	IpAddresses    []string              `json:"ipAddresses,omitempty"`
	IpRoutes       []IpRoute             `json:"ipRoutes,omitempty"`
	NatIpAddresses []string              `json:"natIpAddresses,omitempty"`
	Bfd            BfdConfig             `json:"bfd,omitempty"`
	BgpConnections []BgpConnectionConfig `json:"bgpConnections,omitempty"`
}

type IpRoute struct {
	Prefix      string `json:"prefix"`
	Description string `json:"description,omitempty"`
	NextHop     string `json:"nextHop"`
}

type BfdConfig struct {
	TxInterval int `json:"txInterval,omitempty"`
	RxInterval int `json:"rxInterval,omitempty"`
	Multiplier int `json:"multiplier,omitempty"`
}

type BgpConnectionConfig struct {
	PeerAsn         int      `json:"peerAsn"`
	LocalIpAddress  string   `json:"localIpAddress"`
	PeerIpAddress   string   `json:"peerIpAddress"`
	Password        string   `json:"password,omitempty"`
	Shutdown        bool     `json:"shutdown"`
	Description     string   `json:"description,omitempty"`
	MedIn           int      `json:"medIn,omitempty"`
	MedOut          int      `json:"medOut,omitempty"`
	BfdEnabled      bool     `json:"bfdEnabled"`
	ExportPolicy    string   `json:"exportPolicy,omitempty"`
	PermitExportTo  []string `json:"permitExportTo,omitempty"`
	DenyExportTo    []string `json:"denyExportTo,omitempty"`
	ImportWhitelist int      `json:"importWhitelist,omitempty"`
	ImportBlacklist int      `json:"importBlacklist,omitempty"`
	ExportWhitelist int      `json:"exportWhitelist,omitempty"`
	ExportBlacklist int      `json:"exportBlacklist,omitempty"`
}

// AWS STUFF

type AWSVXCOrder struct {
	AssociatedVXCs []AWSVXCOrderConfiguration `json:"associatedVxcs"`
	PortID         string                     `json:"productUid"`
}

type AWSVXCOrderConfiguration struct {
	Name      string                       `json:"productName"`
	RateLimit int                          `json:"rateLimit"`
	AEnd      VXCOrderAEndConfiguration    `json:"aEnd"`
	BEnd      AWSVXCOrderBEndConfiguration `json:"bEnd"`
}

type AWSVXCOrderBEndConfiguration struct {
	ProductUID    string                       `json:"productUid"`
	PartnerConfig AWSVXCOrderBEndPartnerConfig `json:"partnerConfig"`
}

type AWSVXCOrderBEndPartnerConfig struct {
	ConnectType       string `json:"connectType"`
	Type              string `json:"type"`
	OwnerAccount      string `json:"ownerAccount"`
	ASN               int    `json:"asn,omitempty"`
	AmazonASN         int    `json:"amazonAsn,omitempty"`
	AuthKey           string `json:"authKey,omitempty"`
	Prefixes          string `json:"prefixes,omitempty"`
	CustomerIPAddress string `json:"customerIpAddress,omitempty"`
	AmazonIPAddress   string `json:"amazonIpAddress,omitempty"`
	ConnectionName    string `json:"name,omitempty"`
}

// Partner

type PartnerOrder struct {
	PortID         string                 `json:"productUid"`
	AssociatedVXCs []PartnerOrderContents `json:"associatedVxcs"`
}

type PartnerOrderContents struct {
	Name      string                        `json:"productName"`
	RateLimit int                           `json:"rateLimit"`
	AEnd      VXCOrderAEndConfiguration     `json:"aEnd"`
	BEnd      PartnerOrderBEndConfiguration `json:"bEnd"`
}

type PartnerOrderBEndConfiguration struct {
	PartnerPortID string      `json:"productUid"`
	PartnerConfig interface{} `json:"partnerConfig"`
}

type PartnerOrderAzurePartnerConfig struct {
	ConnectType string                           `json:"connectType"`
	ServiceKey  string                           `json:"serviceKey"`
	Peers       []PartnerOrderAzurePeeringConfig `json:"peers"`
}

type PartnerOrderAzurePeeringConfig struct {
	Type            string `json:"type"`
	PeerASN         string `json:"peer_asn"`
	PrimarySubnet   string `json:"primary_subnet"`
	SecondarySubnet string `json:"secondary_subnet"`
	Prefixes        string `json:"prefixes,omitempty"`
	SharedKey       string `json:"shared_key,omitempty"`
	VLAN            int    `json:"vlan"`
}

type PartnerOrderGooglePartnerConfig struct {
	ConnectType string `json:"connectType"`
	PairingKey  string `json:"pairingKey"`
}

type PartnerOrderOciPartnerConfig struct {
	ConnectType string `json:"connectType"`
	VirtualCircutId  string `json:"virtualCircuitId"`
}
