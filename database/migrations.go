package database

import (
	"context"
	"fmt"
	"strings"
)

func (connection *DB) CreateScheduleTable() error {
	createEnumType := func(typeName string, values []string) error {
		checkQuery := fmt.Sprintf("SELECT 1 FROM pg_type WHERE typname = '%s'", typeName)
		var exists int
		err := connection.QueryRow(checkQuery).Scan(&exists)
		if err == nil {
			return nil
		}

		valueList := "'" + strings.Join(values, "', '") + "'"
		statement := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s)", typeName, valueList)

		_, err = connection.Exec(statement)
		return err
	}

	err := createEnumType("task_status", []string{
		string(TaskStatusInactive),
		string(TaskStatusInProgress),
		string(TaskStatusCompleted),
		string(TaskStatusAbandoned),
		string(TaskStatusFailed),
	})
	if err != nil {
		return err
	}

	err = createEnumType("schedule_type", []string{
		string(ScheduleTypeOneTime),
		string(ScheduleTypeRecurring),
	})
	if err != nil {
		return err
	}

	err = createEnumType("schedule_frequency", []string{
		string(ScheduleFrequencyNow),
		string(ScheduleFrequencyLater),
		string(scheduleFrequencyDaily),
		string(scheduleFrequencyWeekly),
		string(scheduleFrequencyMonthly),
		string(scheduleFrequencyQuarterly),
		string(scheduleFrequencyYearly),
	})
	if err != nil {
		return err
	}

	err = createEnumType("end_type", []string{
		string(EndTypeOn),
		string(EndTypeAfter),
		string(EndTypeNever),
	})
	if err != nil {
		return err
	}

	statement := `
        CREATE TABLE IF NOT EXISTS Schedules (
            Task_ID VARCHAR(50) UNIQUE NOT NULL,
            Status task_status NOT NULL DEFAULT 'INACTIVE',
            Schedule_Type schedule_type NOT NULL DEFAULT 'ONE_TIME',
            Schedule_Frequency schedule_frequency NOT NULL DEFAULT 'NOW',
            Schedule_Details VARCHAR(250) CHECK (Schedule_Details IS NULL OR Schedule_Details <> ''),
            Start_Date TIMESTAMP,
            End_Type end_type NOT NULL DEFAULT 'AFTER',
            End_Date DATE CHECK (End_Date IS NULL OR End_Date >= Start_Date),
            End_Occurrences INT DEFAULT 1 CHECK (End_Occurrences IS NULL OR End_Occurrences > 0)
        );
    `
	_, err = connection.Exec(statement)
	if err != nil {
		return err
	}

	comments := []string{
		"COMMENT ON TABLE Schedules IS 'The table that stores task information, including schedule details'",
		"COMMENT ON COLUMN Schedules.Task_ID IS 'Unique identifier used to identify the task.'",
		"COMMENT ON COLUMN Schedules.Status IS 'Represents the status of the task (INACTIVE, ACTIVE, COMPLETED, ABANDONED, FAILED).'",
		"COMMENT ON COLUMN Schedules.Schedule_Type IS 'Represents the type of the task (ONE_TIME, RECURRING).'",
		"COMMENT ON COLUMN Schedules.Schedule_Frequency IS 'Represents the execution frequency of the task (NOW, LATER, DAILY, WEEKLY, MONTHLY, QUARTERLY, YEARLY).'",
		"COMMENT ON COLUMN Schedules.Schedule_Details IS 'Additional details about the recurrence, e.g., specific days of the week, quarter of the year.'",
		"COMMENT ON COLUMN Schedules.Start_Date IS 'The date and time when the task is scheduled to start.'",
		"COMMENT ON COLUMN Schedules.End_Type IS 'Represents the end schedule of the task (ON, AFTER, NEVER).'",
		"COMMENT ON COLUMN Schedules.End_Date IS 'The date when the schedule should end for ON schedules.'",
		"COMMENT ON COLUMN Schedules.End_Occurrences IS 'The number of occurrences after which the schedule should end for AFTER schedules.'",
	}
	for _, comment := range comments {
		_, err = connection.Exec(comment)
		if err != nil {
			return err
		}
	}

	return nil
}

func (connection *DB) Setup() error {
	ctx := context.Background()
	tx, err := connection.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = connection.CreateScheduleTable()
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
