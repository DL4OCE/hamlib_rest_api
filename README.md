# Hamlib REST API

A modern, high-performance Go-based REST API and management GUI designed to orchestrate, monitor, and control multiple Hamlib `rigctld` and `rotctld` instances across distributed environments. 

This project serves as a lightweight, robust control layer for amateur radio stations, allowing seamless integration of transceivers and antenna rotators into web applications, automated logging software, or site-wide control dashboards.

---

## What it Does & Key Features

* **Foundation for Distributed Architecture:** Run your transceivers and antenna rotators wherever you want on your network (e.g., across multiple distant Raspberry Pis or even on the other side of the Earth by tunneling through the internet) and interact with all of them in an incredibly simple, centralized way.
* **Unified Device Management:** Provides a centralized HTTP REST API to control transceiver parameters (Frequency, Mode, VFO, Split, PTT, Levels, Tuning steps) and Rotator orientations independant from the used hardware (hamlib does that tbh, the API just wraps around it's functionality).
* **Service Orchestration:** Dynamically interacts with systemd to `start` or `stop` specific `rigctld` and `rotctld` background services directly from the web GUI or via API calls.
* **Live Status Monitoring:** Periodically polls and reflects the exact live state (`RUNNING` / `STOPPED`) of individual device daemons.
* **Raw Command Passthrough:** Supports raw execution (`/raw` and `/raw_rx`) for advanced, rig-specific Hamlib sequences while protecting regular workflows.
* **Embedded Web GUI:** Features a clean, lightweight responsive control dashboard baked directly into the Go binary via `go:embed`—no external web server (like Apache or Nginx) required.

---

## Advantages Over the Legacy PHP Implementation

Upgrading from the traditional PHP-based architecture to this native Go implementation brings substantial operational benefits:

| Feature / Metric | Legacy PHP Implementation | Modern Go Architecture |
| :--- | :--- | :--- |
| **Concurrency & Speed** | Blocking I/O per request; slow multi-device polling. | Highly concurrent, non-blocking asynchronous execution. |
| **Deployment** | Requires a full LAMP/LEMP stack (PHP-FPM, Webserver, Permissions). | Single, self-contained statically linked binary. Zero external runtime dependencies. |
| **System Integration** | Complex `exec()` or `sudo` wrappers prone to string-injection or hanging timeouts. | Clean, type-safe OS process execution utilizing native channels and structured timing. |
| **Resource Footprint** | High memory usage per worker process. | Minimal CPU/RAM footprint—ideal for running continuously on a Raspberry Pi. |

---

## Distributed & Scalable Architecture

The system is designed natively for decentralized, scalable, and geographically distributed station setups. The Hamlib REST API Manager runs directly on the edge device (e.g., a Raspberry Pi) hosting the physical hardware connections.

* **Single-Node Scaling:** A single API instance on a single Raspberry Pi can manage $n$ parallel transceivers and antenna rotators by managing independent systemd template instances (`rigctld@1`, `rigctld@2`, `rotctld@1`, etc.).
* **Multi-Node Distribution:** For complex or geographically separated stations, you simply deploy additional autonomous Raspberry Pis—each running its own REST API instance. A central control software or dashboard can then aggregate and control all nodes over the network via simple HTTP calls.

```text
==================================================================================
GEOGRAPHICALLY DISTRIBUTED NETWORK (e.g., LAN, VPN, or Remote Site)
==================================================================================

       +------------------------------------------------------------------+
       |                     Central Control Software                     |
       |              (Contest Logger, Node-RED, Custom GUI)              |
       +-------+--------------------------+------------------------+------+
               |                          |                        |
        HTTP REST Calls            HTTP REST Calls          HTTP REST Calls
               |                          |                        |
  +------------v------------+ +-----------v------------+ +-----------v-------------+
  |   SITE A: Raspberry Pi  | |   SITE B: Raspberry Pi | |   SITE C: Raspberry Pi  |
  | (Hamlib REST API        | | (Hamlib REST API       | | (Hamlib REST API        |
  +----+---------------+----+ +------------+-----------+ +----+---+---+---+---+----+
       |               |                   |                  |   |   |   |   |
       |               |                   |                  |   |   |   |   |
 +-----v-----+   +-----v-----+       +-----v-----+       +----v+--v--+v---+---v--+
 | rigctld@1 |   | rotctld@1 |       | rigctld@1 |       |rig@1|rig@2|rig@3|rot@1|
 +-----+-----+   +-----+-----+       +-----+-----+       +-----+-----+-----+-----+
       |               |                   |               |     |     |     |
 (/dev/ttyUSB0)  (/dev/ttyUSB1)      (/dev/ttyUSB0)        |     |     |     |
       |               |                   |               |     |     |     |
 +-----v-----+   +-----v-----+       +-----v-----+       +-v-+ +-v-+ +-v-+ +-v-+
 |  Rig #1   |   | Rotator #1|       |  Rig #2   |       |Rig| |Rig| |Rig| |Rot|
 | (IC-7300) |   |  (G-601)  |       |  (K3S)    |       |#3 | |#4 | |#5 | |#2 |
 +-----------+   +-----------+       +-----------+       +---+ +---+ +---+ +---+

 ```

## Configuration Manual

The application relies on configuration files located in `/etc/hamlib_rest_api/`. These files tell the Go backend which devices exist and provide the parameters used by the underlying `systemd` template units (`rigctld@.service` and `rotctld@.service`) to spin up the daemons.

### 1. Transceiver Configuration (`rigctld.json`)

This file holds an array of JSON objects representing your transceivers. 

#### Production Configuration Example
```json
[
  {
    "id": 1,
    "model": "1",
    "device": "/dev/ttyUSB0",
    "baudrate": "19200",
    "port": "4500"
  },
  {
    "id": 2,
    "model": "1022",
    "device": "/dev/ttyUSB1",
    "baudrate": "9600",
    "port": "4501"
  }
]
```

#### Field Descriptions:
* **`id`** *(Integer)*: A unique internal identifier for the rig (e.g., `1`, `2`). This matches the systemd template suffix (e.g., `rigctld@1.service`) and anchors the URL route structure (`/trx/1/...`).
* **`model`** *(String)*: The Hamlib ID code matching your radio. 
  * Use `"1"` for a standard Hamlib dummy transceiver.
  * To lookup your physical device ID, check the official [Hamlib Supported Rig List Matrix](https://github.com/Hamlib/Hamlib/wiki/Supported-Radios) or run `rigctl -l` on your system.
* **`device`** *(String)*: Absolute path to the hardware interface TTY line (e.g., `/dev/ttyUSB0`). 
* **`baudrate`** *(String)*: Serial speed matching the radio's VFO interface menu settings.
* **`port`** *(String)*: The TCP port assigned to this daemon process instance (e.g., `"4500"`). Every rig instance must route to its own port block.

---

### 2. Rotator Configuration (`rotctld.json`)

This file configures individual array objects tracking antenna heading controllers.

#### Production Configuration Example
```json
[
  {
    "id": 1,
    "model": "1",
    "device": "dummy",
    "port": "4500",
    "baudrate": "9600"
  },
  { 
    "id": 2,
    "model": "1",
    "device": "dummy",
    "port": "4500",
    "baudrate": "19200"
  }
]
```

#### Field Descriptions:
* **`id`** *(Integer)*: A unique internal ID for the rotor unit, anchoring systemd tracking (`rotctld@1.service`) and UI controls (`/rotator/1/...`).
* **`model`** *(String)*: The exact ID mapping your position controller model.
  * Use `"1"` for the standard testing dummy driver.
  * To locate your specific controller type, review the [Hamlib Rotator Listing Wiki](https://github.com/Hamlib/Hamlib/wiki/Supported-Rotators) or execute `rotctl -l` on the host machine.
* **`device`** *(String)*: The hardware system path (e.g., `/dev/ttyUSB2`) or `"dummy"`.
* **`port`** *(String)*: Local TCP listening network socket port assigned to the `rotctld` process daemon.
* **`baudrate`** *(String)*: Physical serial interface speed for target system controller link adjustments.

---

## Runtime Application Behavior

Because the Hamlib REST API Manager evaluates configuration contents dynamically from storage media directly during processing requests, **restarting the central Go API daemon after editing JSON maps is not necessary**. The integrated UI view dashboard will pick up configuration layout mutations cleanly on your next browser reload.

*(Operational Note: If adjusting configurations for an actively running daemon instance, stop the target driver service via the UI action element first, adjust the values inside the target /etc/hamlib_rest_api/ JSON map, and select Start to re-initialize the unit.)*

```