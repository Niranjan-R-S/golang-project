package tests

import (
	"../model/rnaTranscriptionModel"
	"testing"
)

func TestConvertDNAToRNA(t *testing.T) {
	var DNAString = "CTAG"
	var convertedRNA = rnaTranscriptionModel.ConvertDNAToRNA(DNAString)
	if convertedRNA != "GAUC" {
		t.Errorf("Returned: %s, Expected: %s", convertedRNA, DNAString)
	}
}

