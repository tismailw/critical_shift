# Critical Shift – Device Check-In API

This is a simple REST API built in Go using the Gin framework.
It simulates an IoT device check-in service where devices report metadata such as firmware version, tenant ownership, and bill of materials (BOM).

I built this as a small project to demonstrate basic API design, routing, and JSON handling in Go.

---

## Tech Stack

* Go
* Gin (HTTP web framework)

---

## Getting Started

### Prerequisites

* Go 1.20+ installed

### Install Dependencies

```bash
go mod init critical-shift
go get github.com/gin-gonic/gin
```

### Run the Server

```bash
go run critical_shift.go
```

The server will start on:

```
http://localhost:8080
```

---

## API Endpoints

### Get All Devices

```http
GET /d
```

Example:

```bash
curl http://localhost:8080/d
```

---

### Get Device by Device ID

```http
GET /d/:id
```

Example:

```bash
curl http://localhost:8080/d/dev-001
```

---

### Get Devices by Tenant ID

```http
GET /d/t/:id
```

Example:

```bash
curl http://localhost:8080/d/t/tenant-alpha
```

---

### Add a New Device

```http
POST /d
```

Example:

```bash
curl http://localhost:8080/d \
  --header "Content-Type: application/json" \
  --data '{
    "device_id": "dev-999",
    "tenant_id": "tenant-charlie",
    "firmware": "v1.2.0",
    "bom": {
      "cpu": "ARM Cortex-A76",
      "ram_mb": 4096,
      "storage": "64GB eMMC",
      "radio": "5G"
    },
    "timestamp": "2026-01-30T15:00:00Z"
  }'
```

---

## Data Model

Each device check-in contains:

* `device_id` – unique device identifier
* `tenant_id` – tenant/company the device belongs to
* `firmware` – firmware version
* `bom` – bill of materials (hardware metadata)
* `timestamp` – check-in time (ISO 8601)

---

## Notes

* Data is stored in memory (no database).
* This project focuses on API structure and request handling, not persistence or authentication.
* Initial device data is mocked for demonstration purposes.

---

Nice idea — interviewers *love* this section. Here’s a **clean, realistic “Things to Do in the Future”** section you can drop straight into the README. It sounds thoughtful without overengineering.

You can paste this at the bottom of your README.

---

## Things to Do in the Future

* **Device Online / Offline Tracking**
  Track the last check-in time for each device and automatically mark devices as *offline* if they have not reported in within a configurable time window.

* **Firmware Update Check on Reconnect**
  When a device comes back online after being offline, automatically compare its firmware version against the latest available version and flag whether an update is required.

* **Heartbeat Endpoint**
  Introduce a lightweight heartbeat endpoint that devices can call periodically to confirm connectivity without sending full metadata.

* **Persistent Storage**
  Replace in-memory storage with a database (PostgreSQL, MongoDB, etc.) to persist device state, history, and firmware versions across restarts.

* **Device Status Endpoint**
  Add an endpoint to retrieve current device status (online/offline), last seen timestamp, and update eligibility.

* **Authentication & Authorization**
  Secure endpoints using API keys or OAuth to ensure devices and tenants can only access their own data.

* **Event Logging**
  Store historical check-in events for monitoring, debugging, and analytics purposes.
