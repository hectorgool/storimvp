package schema

type (
	Document struct {
		Id, Date, Transaction string
	}
	DBDocument struct {
		Id          uint   `gorm:"primaryKey"`
		Date        string `gorm:"type:date"`
		Transaction string `json:"transaction"`
	}
)
