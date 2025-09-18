package printer

import "DocumentConverter/models"

type Printer interface {
	Print(node models.DocumentNode) string
}
