package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"pp/api/utils"
	"testing"
)

func TestParseCreateRolePayloadPositive(t *testing.T) {
	expectedRoleName := "test"
	rcr := RoleCreateRequest{Name: expectedRoleName}
	jsonValue, _ := json.Marshal(rcr)
	payload := bytes.NewBuffer(jsonValue)
	request, err := http.NewRequest(http.MethodPost, "", payload)

	if err != nil {
		t.Fatal(err)
	}

	result, err := parseCreateRolePayload(request)

	if err != nil {
		t.Fatal(err)
	}

	if result.Name != rcr.Name {
		t.Errorf("Failed, expected: %v, got: %v", expectedRoleName, result)
	}
}

func TestParseCreateRolePayloadNegative(t *testing.T) {
	payload := new(bytes.Buffer)
	request, err := http.NewRequest(http.MethodPost, "", payload)

	if err != nil {
		t.Fatal(err)
	}

	_, err = parseCreateRolePayload(request)

	if !errors.Is(utils.InvalidRequestPayload, err) {
		t.Errorf("Failed, expected: %s, got: %s", utils.InvalidRequestPayload.Error(), err.Error())
	}
}

func TestCheckLimitOffset(t *testing.T) {
	var limit, offset = 0, 0
	var expectedLimit, expectedOffset = 10, 0
	checkLimitOffset(&limit, &offset)

	if limit != expectedLimit {
		t.Errorf("Failed, expected: %d, got: %d", expectedLimit, limit)
	}

	if offset != expectedOffset {
		t.Errorf("Failed, expected: %d, got: %d", expectedOffset, offset)
	}

	limit, offset = 10, 0
	checkLimitOffset(&limit, &offset)

	if limit != expectedLimit {
		t.Errorf("Failed, expected: %d, got: %d", expectedLimit, limit)
	}

	if offset != expectedOffset {
		t.Errorf("Failed, expected: %d, got: %d", expectedOffset, offset)
	}

	limit, offset = 0, -1

	checkLimitOffset(&limit, &offset)

	if limit != expectedLimit {
		t.Errorf("Failed, expected: %d, got: %d", expectedLimit, limit)
	}

	if offset != expectedOffset {
		t.Errorf("Failed, expected: %d, got: %d", expectedOffset, offset)
	}

	limit, offset = 1, 0
	expectedLimit, expectedOffset = 1, 0
	checkLimitOffset(&limit, &offset)

	if limit != expectedLimit {
		t.Errorf("Failed, expected: %d, got: %d", expectedLimit, limit)
	}

	if expectedOffset != 0 {
		t.Errorf("Failed, expected: %d, got: %d", expectedOffset, offset)
	}

	limit, offset = 11, 0
	expectedLimit, expectedOffset = 10, 0
	checkLimitOffset(&limit, &offset)

	if limit != expectedLimit {
		t.Errorf("Failed, expected: %d, got: %d", expectedLimit, limit)
	}

	if offset != expectedOffset {
		t.Errorf("Failed, expected: %d, got: %d", expectedOffset, offset)
	}
}

func TestGetLimitOffset(t *testing.T) {
	payload := new(bytes.Buffer)
	request, err := http.NewRequest(http.MethodPost, "", payload)

	if err != nil {
		t.Fatal(err)
	}

	var expectedLimit, expectedOffset = 0, 0
	limit, offset := getLimitOffset(request)

	if limit != expectedLimit {
		t.Errorf("Failed, expected: %d, got: %d", expectedLimit, limit)
	}

	if offset != expectedOffset {
		t.Errorf("Failed, expected: %d, got: %d", expectedOffset, offset)
	}

	payload = &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("limit", "1")
	_ = writer.WriteField("offset", "1")
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	request, err = http.NewRequest(http.MethodGet, "", payload)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	expectedLimit, expectedOffset = 1, 1
	limit, offset = getLimitOffset(request)

	if limit != expectedLimit {
		t.Errorf("Failed, expected: %d, got: %d", expectedLimit, limit)
	}

	if offset != expectedOffset {
		t.Errorf("Failed, expected: %d, got: %d", expectedOffset, offset)
	}
}
