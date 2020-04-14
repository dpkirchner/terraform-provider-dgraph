package resources

import (
	"log"
	"math"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Retry will run the specified Terraform resource function until it succeeds or
// hits a non-retryable error. This should only be used for idempotent alters.
func Retry(fn func(d *schema.ResourceData, m interface{}) error) func(*schema.ResourceData, interface{}) error {
	return func(d *schema.ResourceData, m interface{}) error {
		attempt := 0

		for {
			err := fn(d, m)
			if err == nil {
				return nil
			}

			// The full error string Dgraph returns is "errIndexingInProgress. Please
			// retry" and it means it's still modifying the schema from a previous
			// request.
			if !strings.Contains(err.Error(), "errIndexingInProgress") {
				return err
			}

			attempt++

			log.Printf("[DEBUG] Retry #%d for %s", attempt, d.Get("name").(string))

			delay := math.Min(500, math.Pow(2, float64(attempt))*100)
			time.Sleep(time.Duration(delay) * time.Millisecond)
		}
	}
}
