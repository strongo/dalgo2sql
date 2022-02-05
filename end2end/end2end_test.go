package end2end

import "testing"

func TestEndToEnd(t *testing.T) {
	t.Run("RAMSQLDB", func(t *testing.T) {
		testEndToEndRAMSQLDB(t)
	})
}
