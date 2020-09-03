package entities

import (
	"time"

	"github.com/mercadolibre/fury_go-nullables/pkg/nullable"
)

type Values struct {
	DNI      int64     `json:"dni"`
	Money    float64   `json:"money"`
	FullName string    `json:"full_name"`
	Working  bool      `json:"working"`
	JoinDate time.Time `json:"join_date"`
}

type Pointers struct {
	DNI      *int64     `json:"dni"`
	Money    *float64   `json:"money"`
	FullName *string    `json:"full_name"`
	Working  *bool      `json:"working"`
	JoinDate *time.Time `json:"join_date"`
}

type Nullables struct {
	DNI      nullable.Int    `json:"dni"`
	Money    nullable.Float  `json:"money"`
	FullName nullable.String `json:"full_name"`
	Working  nullable.Bool   `json:"working"`
	JoinDate nullable.Time   `json:"join_date"`
}
