package wms

import (
	"encoding/xml"
)

type Capabilities interface {
	Version() string
	Layers() Layer
}

type Capabilities1_1_1 struct {
}

type Capabilities1_3_0 struct {
	XMLName    xml.Name   `xml:"WMS_Capabilities" json:"-"`
	Service    Service    `xml:"Service" json:"service"`
	Capability Capability `xml:"Capability" json:"capability"`
}

//func (c130 *Capabilities1_3_0) Version() string {
//	return string(Version130)
//}
//
//func (c130 *Capabilities1_3_0) Layers() Layer {
//	return c130.Capability.Layer
//}

type Service struct {
	Name                         string         `xml:"Name" json:"name"`
	Title                        string         `xml:"Title" json:"title"`
	Abstract                     string         `xml:"Abstract" json:"abstract"`
	KeywordList                  []string       `xml:"KeywordList>Keyword" json:"keyword_list"`
	OnlineResource               OnlineResource `xml:"OnlineResource" json:"online_resource"`
	ContactPerson                string         `xml:"ContactInformation>ContactPersonPrimary>ContactPerson" json:"contact_person"`
	ContactOrganization          string         `xml:"ContactInformation>ContactPersonPrimary>ContactOrganization" json:"contact_organization"`
	ContactPosition              string         `xml:"ContactInformation>ContactPosition" json:"position"`
	ContactAddress               Address        `xml:"ContactInformation>ContactAddress" json:"address"`
	ContactVoiceTelephone        string         `xml:"ContactInformation>ContactVoiceTelephone" json:"phone"`
	ContactFacsimileTelephone    string         `xml:"ContactInformation>ContactFacsimileTelephone" json:"fax"`
	ContactElectronicMailAddress string         `xml:"ContactInformation>ContactElectronicMailAddress" json:"email"`
	Fees                         string         `xml:"Fees" json:"fees"`
	AccessConstraints            string         `xml:"AccessConstraints" json:"access_constraints"`
}

type Capability struct {
	GetMapFormat         []string `xml:"Request>GetMap>Format" json:"get_map_format"`
	GetFeatureInfoFormat []string `xml:"Request>GetFeatureInfo>Format" json:"get_feature_info_format"`
	ExceptionFormat      []string `xml:"Exception>Format" json:"exception_format"`
	Layer                Layer    `xml:"Layer" json:"layer"`
}

type Layer struct {
	Queryable               string                  `xml:"queryable,attr" json:"queryable,omitempty"`
	Opaque                  string                  `xml:"opaque,attr" json:"opaque,omitempty"`
	Name                    string                  `xml:"Name" json:"name"`
	Title                   string                  `xml:"Title" json:"title"`
	Abstract                string                  `xml:"Abstract" json:"abstract"`
	CRS                     []string                `xml:"CRS" json:"CRS"`
	EXGeographicBoundingBox EXGeographicBoundingBox `xml:"EX_GeographicBoundingBox" json:"ex_geographic_bounding_box"`
	BoundingBox             []BoundingBox           `xml:"BoundingBox" json:"bounding_box"`
	Layers                  []Layer                 `xml:"Layer" json:"layer"`
	Style                   []Style                 `xml:"Style" json:"style"`
	KeywordList             []string                `xml:"KeywordList>Keyword" json:"keyword_list"`
}

type BoundingBox struct {
	MinX float64 `xml:"minx,attr"  json:"minx"`
	MaxX float64 `xml:"maxx,attr" json:"maxx"`
	MinY float64 `xml:"miny,attr" json:"miny"`
	MaxY float64 `xml:"maxy,attr" json:"maxy"`
	CRS  string  `xml:"CRS,attr" json:"crs"`
}

type Style struct {
	Name      string    `xml:"Name" json:"name"`
	Title     string    `xml:"Title" json:"title"`
	Abstract  string    `xml:"Abstract" json:"abstract"`
	LegendURL LegendURL `xml:"LegendURL" json:"legent_url"`
}

type LegendURL struct {
	Width          string         `xml:"width,attr" json:"width"`
	Height         string         `xml:"height,attr" json:"height"`
	OnlineResource OnlineResource `xml:"OnlineResource" json:"online_resource"`
}

type OnlineResource struct {
	Href string `xml:"href,attr" json:"href"`
}

type Address struct {
	AddressType     string `xml:"AddressType" json:"address_type"`
	Address         string `xml:"Address" json:"address"`
	City            string `xml:"City" json:"city"`
	StateOrProvince string `xml:"StateOrProvince" json:"state_or_province"`
	PostCode        string `xml:"PostCode" json:"post_code"`
	Country         string `xml:"Country" json:"country"`
}

type EXGeographicBoundingBox struct {
	WestBoundLongitude float64 `xml:"westBoundLongitude" json:"west_bound_longitude"`
	EastBoundLongitude float64 `xml:"eastBoundLongitude" json:"east_bound_longitude"`
	NorthBoundLatitude float64 `xml:"northBoundLatitude" json:"north_bound_latitude"`
	SouthBoundLatitude float64 `xml:"southBoundLatitude" json:"south_bound_latitude"`
}
