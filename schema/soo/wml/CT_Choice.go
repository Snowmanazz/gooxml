package wml

import (
	"encoding/xml"
	"github.com/Snowmanazz/gooxml"
)

type CT_Choice struct {
	Drawing *CT_Drawing
}

func NewCT_Choice() *CT_Choice {
	return &CT_Choice{}
}

func (m *CT_Choice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Drawing != nil {
		sedrawing := xml.StartElement{Name: xml.Name{Local: "w:drawing"}}
		e.EncodeElement(m.Drawing, sedrawing)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Choice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
CT_Choice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "drawing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "drawing"}:
				m.Drawing = NewCT_Drawing()
				if err := d.DecodeElement(m.Drawing, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on EG_RunInnerContent %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break CT_Choice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Drawing and its children
func (m *CT_Choice) Validate() error {
	return m.ValidateWithPath("CT_Choice")
}

// ValidateWithPath validates the CT_Drawing and its children, prefixing error messages with path
func (m *CT_Choice) ValidateWithPath(path string) error {
	if m.Drawing != nil {
		if err := m.Drawing.ValidateWithPath(path + "/Drawing"); err != nil {
			return err
		}
	}
	return nil
}
