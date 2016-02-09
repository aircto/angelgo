package angelco_testing

import (
    // "fmt"
    "testing"
)

func TestJobsList(t *testing.T) {
    _, err := api.JobsList()
    if err != nil {
        t.Error(err)
    }
}

