package schema

type (
	Document struct {
		Id, Date, Transaction string
	}
	DBDocument struct {
		Id            uint    `gorm:"primaryKey"`
		IdTransaction uint    `json:"idTransaction"`
		Date          string  `gorm:"type:date"`
		Transaction   float64 `json:"transaction"`
	}
)
