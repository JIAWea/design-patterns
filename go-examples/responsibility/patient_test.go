package responsibility

import (
	"testing"
)

func TestResponsibility(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cashier := &cashier{}

			//Set next for medical department
			medical := &medical{}
			medical.setNext(cashier)

			//Set next for doctor department
			doctor := &doctor{}
			doctor.setNext(medical)

			//Set next for reception department
			reception := &reception{}
			reception.setNext(doctor)

			patient := &patient{name: tt.name}
			//Patient visiting
			reception.execute(patient)
		})
	}
}
