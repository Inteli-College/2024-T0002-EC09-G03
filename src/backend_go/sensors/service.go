package sensors

import "fmt"

func getReadings(body *ServiceBody) error {
  fmt.Printf("Reading: %s \n", body.Reading)

  return nil
}
