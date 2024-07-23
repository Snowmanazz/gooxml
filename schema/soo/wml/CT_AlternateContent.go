package wml

import (
	"encoding/xml"
	"github.com/Snowmanazz/gooxml"
)

type CT_AlternateContent struct {
	Choice   *CT_Choice
	Fallback *gooxml.XSDAny
}

func NewCT_AlternateContent() *CT_AlternateContent {
	ret := &CT_AlternateContent{}
	return ret
}
func (m *CT_AlternateContent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Choice != nil {
		seanchor := xml.StartElement{Name: xml.Name{Local: "mc:Choice"}}
		e.EncodeElement(m.Choice, seanchor)
	}
	if m.Fallback != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "mc:Fallback"}}
		e.EncodeElement(m.Fallback, sedrawing)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AlternateContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AlternateContent:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/markup-compatibility/2006", Local: "Choice"}:
				m.Choice = NewCT_Choice()
				if err := d.DecodeElement(m.Choice, &el); err != nil {
					return err
				}
			default:
				m.Fallback = new(gooxml.XSDAny)
				if err := d.DecodeElement(m.Fallback, &el); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AlternateContent
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AlternateContent and its children
func (m *CT_AlternateContent) Validate() error {
	return m.ValidateWithPath("CT_AlternateContent")
}

// ValidateWithPath validates the CT_AlternateContent and its children, prefixing error messages with path
func (m *CT_AlternateContent) ValidateWithPath(path string) error {
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	return nil
}
