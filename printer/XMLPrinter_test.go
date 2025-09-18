package printer

import (
	"DocumentConverter/models"
	"testing"
)

func TestXMLPrinter_PrintNode(t *testing.T) {
	printer := NewXMLPrinter()

	result := printer.Print(models.DocumentNode{
		Name:       "root",
		Attributes: []models.Attribute{{Name: "format", Value: "v2"}},
		Children: []models.DocumentNode{{
			Name:       "user",
			Attributes: []models.Attribute{{Name: "Name", Value: "Mr"}, {Name: "lastName", Value: "Tester"}},
			Children: []models.DocumentNode{{
				Name:       "phone",
				Attributes: []models.Attribute{{Name: "landline", Value: "024-23123"}, {Name: "mobile", Value: "073-1232434"}},
				Children:   nil,
			},
				{
					Name:       "address",
					Attributes: []models.Attribute{{Name: "street", Value: "testgatan2"}, {Name: "zip", Value: "15"}},
					Children:   nil,
				},
			},
		},
			{
				Name:       "user",
				Attributes: []models.Attribute{{Name: "Name", Value: "Ms"}, {Name: "lastName", Value: "Tester"}},
				Children: []models.DocumentNode{{
					Name:       "phone",
					Attributes: []models.Attribute{{Name: "landline", Value: "024-23123"}, {Name: "mobile", Value: "076-8123453"}},
					Children:   nil,
				},
				},
			},
		},
	})

	t.Logf("\n%s", result)
}
