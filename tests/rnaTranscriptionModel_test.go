package tests

import (
	"../service/rnaTranscriptionService"
	"testing"
)

func TestConvertDNAToRNA(t *testing.T) {
	var DNAString = "CTAG"
	var convertedRNA = rnaTranscriptionService.ConvertDNAToRNA(DNAString)
	if convertedRNA != "GAUC" {
		t.Errorf("Returned: %s, Expected: %s", convertedRNA, DNAString)
	}
}

