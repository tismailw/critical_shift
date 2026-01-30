package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/d", getDevices)
	router.GET("/d/:id", getDevicesByDevice_id)
	router.GET("/d/t/:id", getDevicesByTenant_id)
	router.POST("/d", postDevice)

	router.Run("localhost:8080")
}

// iot devices would probably have fields like, deviceID,
type DeviceCheckIn struct {
	DeviceID string `json:"device_id"`
	TenantID string `json:"tenant_id"` // An ID that references the user, x data belongs to tenent_id (compnay x)
	//SerialNumber string `json:"serial_number"`
	Firmware string `json:"firmware"`

	BOM map[string]any `json:"bom"` //bill of materials

	Timestamp string `json:"timestamp"`
}

// data generated via llm
var devices = []DeviceCheckIn{
	{
		DeviceID: "dev-001",
		TenantID: "tenant-alpha",
		Firmware: "v1.0.3",
		BOM: map[string]any{
			"cpu":     "ARM Cortex-A53",
			"ram_mb":  1024,
			"storage": "16GB eMMC",
			"radio":   "LTE-M",
		},
		Timestamp: "2026-01-30T14:05:00Z",
	},
	{
		DeviceID: "dev-002",
		TenantID: "tenant-alpha",
		Firmware: "v1.1.0",
		BOM: map[string]any{
			"cpu":     "ARM Cortex-A72",
			"ram_mb":  2048,
			"storage": "32GB eMMC",
			"radio":   "WiFi",
		},
		Timestamp: "2026-01-30T14:12:00Z",
	},
	{
		DeviceID: "dev-101",
		TenantID: "tenant-bravo",
		Firmware: "v0.9.8",
		BOM: map[string]any{
			"cpu":     "ARM Cortex-M7",
			"ram_mb":  512,
			"storage": "8GB eMMC",
			"radio":   "LoRa",
		},
		Timestamp: "2026-01-30T13:58:00Z",
	},
}

// get Devices - get a list of all devices
/*
curl http://localhost:8080/d
*/
func getDevices(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, devices)
}

// post Devices - add another device
func postDevice(c *gin.Context) {

/*
curl http://localhost:8080/d \
  --header "Content-Type: application/json" \
  --data '{
    "device_id": "me_101",
    "tenant_id": "tenant-alpha",
    "firmware": "v1.2.0",
    "bom": {
      "cpu": "ARM Cortex-A76",
      "ram_mb": 4096,
      "storage": "64GB eMMC",
      "radio": "5G"
    },
    "timestamp": "2026-01-30T15:00:00Z"
  }'
*/

	var newDevice DeviceCheckIn

	//check if newAlbum is null or not
	err := c.BindJSON(&newDevice)
	if err != nil { // if the error is not nothing then we return bc that means theres an error
		return
	}

	devices = append(devices, newDevice)

	c.IndentedJSON(http.StatusCreated, newDevice)
}

// get Devices{id} - getting the device with the id specified
/*
curl http://localhost:8080/d/{id}
*/
func getDevicesByDevice_id(c *gin.Context) {

	id := c.Param("id")

	for _, x := range devices {
		if x.DeviceID == id {
			c.IndentedJSON(http.StatusOK, x)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": " Device ID does not exist",
		"ID": id,
	})

}
func getDevicesByTenant_id(c *gin.Context) {
/*
curl http://localhost:8080/d/t/{id}
*/
	id := c.Param("id")

	var results[]DeviceCheckIn

	for _, x := range devices {
		if x.TenantID == id {
			results = append(results, x)
		}
	}
	if len(results) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": " Device ID does not exist",
			"ID": id,
		})
		return


	}
	c.IndentedJSON(http.StatusOK, results)


}
