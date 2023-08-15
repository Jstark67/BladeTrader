// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type CarbonType string

const (
	CarbonTypeWood  CarbonType = "Wood"
	CarbonTypeInner CarbonType = "Inner"
	CarbonTypeOuter CarbonType = "Outer"
)

func (e *CarbonType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CarbonType(s)
	case string:
		*e = CarbonType(s)
	default:
		return fmt.Errorf("unsupported scan type for CarbonType: %T", src)
	}
	return nil
}

type NullCarbonType struct {
	CarbonType CarbonType `json:"carbon_type"`
	Valid      bool       `json:"valid"` // Valid is true if CarbonType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCarbonType) Scan(value interface{}) error {
	if value == nil {
		ns.CarbonType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CarbonType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCarbonType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CarbonType), nil
}

type GripType string

const (
	GripTypePenhold   GripType = "Penhold"
	GripTypeHandshake GripType = "Handshake"
)

func (e *GripType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = GripType(s)
	case string:
		*e = GripType(s)
	default:
		return fmt.Errorf("unsupported scan type for GripType: %T", src)
	}
	return nil
}

type NullGripType struct {
	GripType GripType `json:"grip_type"`
	Valid    bool     `json:"valid"` // Valid is true if GripType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGripType) Scan(value interface{}) error {
	if value == nil {
		ns.GripType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.GripType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGripType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.GripType), nil
}

type Status string

const (
	StatusSold    Status = "sold"
	StatusNotSold Status = "not_sold"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status `json:"status"`
	Valid  bool   `json:"valid"` // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// Positive or Negative
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type RacketInventory struct {
	ID              int64         `json:"id"`
	RacketID        int64         `json:"racket_id"`
	OwnerID         sql.NullInt64 `json:"owner_id"`
	PurchasedOnline bool          `json:"purchased_online"`
	PurchaseDate    time.Time     `json:"purchase_date"`
}

type RacketsForSale struct {
	ID         int64         `json:"id"`
	CarbonType CarbonType    `json:"carbon_type"`
	GripType   GripType      `json:"grip_type"`
	Price      int64         `json:"price"`
	SellerID   int64         `json:"seller_id"`
	PostedTime time.Time     `json:"posted_time"`
	Status     Status        `json:"status"`
	BuyerID    sql.NullInt64 `json:"buyer_id"`
}

type Transfer struct {
	ID              int64 `json:"id"`
	SourceAccountID int64 `json:"source_account_id"`
	TargetAccountID int64 `json:"target_account_id"`
	// Positive Only
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}