package models

import (
	"database/sql"
	"time"
)

type UserSession struct {
	Id                int64               `db:"id" fieldtag:"pk"`
	CreatedAt         sql.NullTime        `db:"created_at"`
	UpdatedAt         sql.NullTime        `db:"updated_at"`
	SessionIdentifier string              `db:"session_identifier"`
	Started           time.Time           `db:"started"`
	LastAccessed      time.Time           `db:"last_accessed"`
	AuthMethods       string              `db:"auth_methods"`
	AcrLevel          string              `db:"acr_level"`
	AuthTime          time.Time           `db:"auth_time"`
	IpAddress         string              `db:"ip_address"`
	DeviceName        string              `db:"device_name"`
	DeviceType        string              `db:"device_type"`
	DeviceOS          string              `db:"device_os"`
	UserId            int64               `db:"user_id"`
	User              User                `db:"-"`
	Clients           []UserSessionClient `db:"-"`
}

type UserSessionClient struct {
	Id            int64        `db:"id" fieldtag:"pk"`
	CreatedAt     sql.NullTime `db:"created_at"`
	UpdatedAt     sql.NullTime `db:"updated_at"`
	UserSessionId int64        `db:"user_session_id"`
	ClientId      int64        `db:"client_id"`
	Client        Client       `db:"-"`
	Started       time.Time    `db:"started"`
	LastAccessed  time.Time    `db:"last_accessed"`
}