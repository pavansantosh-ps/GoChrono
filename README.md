# GoChrono Scheduler

GoChrono is a scalable, open-source task scheduler built with Go, designed to support various one-time and recurring schedules. It uses NATS for communication and includes a fail-safe mechanism with database persistence, ensuring reliable and efficient task execution.

## Prerequisites

- **NATS**: Make sure NATS is running as GoChrono relies on it for cron-like scheduling and worker coordination, enabling message-based task management.
- **Database**: GoChrono supports PostgreSQL or Oracle for storing schedule data and ensuring persistence.

## Setup

1. **Configuration**: Use the `.env` file to configure environment variables. Required variables are specified in this file.
2. **Run GoChrono**: Start the scheduler with:
   ```bash
   go run gochrono/main.go
   ```
   This command performs database migrations and starts the NATS worker.

## Supported Schedules

GoChrono supports the following scheduling options:

### One-Time Schedules

- **Now**: Execute immediately.
- **Later**: Execute at a specified start datetime.

### Recurring Schedules

- **Daily**: Requires a start datetime and executes daily.
- **Weekly**: Requires a start datetime and specified days of the week.
- **Monthly**: Requires a start datetime and executes on a specific date each month.
- **Quarterly**: Requires a start datetime, quarter, and specific date.
- **Yearly**: Requires a start datetime and executes annually on a specified date.

Each schedule can end based on:

- **End On**: Ends on a specific date.
- **After Occurrences**: Ends after a specified number of executions.
- **Never**: Continues indefinitely.

## Contributing

Contributions are welcome. Please submit issues or pull requests to help improve GoChrono.

## License

GoChrono is released under the MIT License.

---
