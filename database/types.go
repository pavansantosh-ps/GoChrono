package database

import (
	"database/sql"
	"time"
)

type Dialect string

const (
	PostgresDialect Dialect = "postgres"
	OracleDialect   Dialect = "oracle"
)

type Config struct {
	Dialect         Dialect
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	Schema          string // For PostgreSQL
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	SSLMode         string // For PostgreSQL
	ConnectTimeout  int    // Connection timeout in seconds
	ServiceName     string
}

type DB struct {
	*sql.DB
	config *Config
}

type taskStatus string

const (
	TaskStatusInactive   taskStatus = "INACTIVE"
	TaskStatusInProgress taskStatus = "ACTIVE"
	TaskStatusCompleted  taskStatus = "COMPLETED"
	TaskStatusAbandoned  taskStatus = "ABANDONED"
	TaskStatusFailed     taskStatus = "FAILED"
)

type scheduleType string

const (
	ScheduleTypeOneTime   scheduleType = "ONE_TIME"
	ScheduleTypeRecurring scheduleType = "RECURRING"
)

type scheduleFrequency string

const (
	ScheduleFrequencyNow       scheduleFrequency = "NOW"
	ScheduleFrequencyLater     scheduleFrequency = "LATER"
	scheduleFrequencyDaily     scheduleFrequency = "DAILY"
	scheduleFrequencyWeekly    scheduleFrequency = "WEEKLY"
	scheduleFrequencyMonthly   scheduleFrequency = "MONTHLY"
	scheduleFrequencyQuarterly scheduleFrequency = "QUARTERLY"
	scheduleFrequencyYearly    scheduleFrequency = "YEARLY"
)

type endType string

const (
	EndTypeOn    endType = "ON"
	EndTypeAfter endType = "AFTER"
	EndTypeNever endType = "NEVER"
)
