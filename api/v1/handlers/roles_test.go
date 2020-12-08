package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"pp/api/models"
	"testing"
)

func TestParseCreateRolePayloadPositive(t *testing.T) {
	expectedRoleName := "test"
	sentRole := models.Role{
		Name: expectedRoleName,
	}

	jsonValue, _ := json.Marshal(sentRole)
	request, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonValue))

	if err != nil {
		t.Fatal(err)
	}

	result, err := parseCreateRolePayload(request)

	if err != nil {
		t.Fatal(err)
	}

	if result.Name != sentRole.Name {
		t.Errorf("Failed, expected: %v, got: %v", expectedRoleName, result)
	}
}

func TestParseCreateRolePayloadNegative(t *testing.T) {
	request, err := http.NewRequest(http.MethodPost, "", nil)

	if err != nil {
		t.Fatal(err)
	}

	_, err = parseCreateRolePayload(request)

	if err != nil {
		t.Fatal(err)
	}
}
