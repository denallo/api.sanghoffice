package components

import (
	"reflect"
	"testing"
)

func TestIsValidSession(t *testing.T) {
	ClearCache()
	const TESTING_USER_ID = "____"
	sessID := GenerateSessionID(TESTING_USER_ID, true, "")

	type args struct {
		sessionID string
	}
	tests := []struct {
		name        string
		args        args
		wantIsValid bool
	}{
		// TODO: Add test cases.
		{
			name:        "case_1",
			args:        args{sessionID: sessID},
			wantIsValid: true,
		},
		{
			name:        "case_2",
			args:        args{sessionID: "you never found me"},
			wantIsValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotIsValid := IsValidSession(tt.args.sessionID); gotIsValid != tt.wantIsValid {
				t.Errorf("IsValidSession() = %v, want %v", gotIsValid, tt.wantIsValid)
			}
		})
	}
}

func TestGetUserID(t *testing.T) {
	ClearCache()
	const TESTING_USER_ID = "____"
	const TESTING_SESSION_ID_NOT_EXISTED = "you never found me."
	sessID := GenerateSessionID(TESTING_USER_ID, true, "")

	type args struct {
		sessionID string
	}
	tests := []struct {
		name        string
		args        args
		wantUserID  string
		wantSuccess bool
	}{
		// TODO: Add test cases.
		{
			name:        "case1",
			args:        args{sessionID: sessID},
			wantUserID:  TESTING_USER_ID,
			wantSuccess: true,
		},
		{
			name:        "case2",
			args:        args{sessionID: TESTING_SESSION_ID_NOT_EXISTED},
			wantUserID:  "",
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserID, gotSuccess := GetUserID(tt.args.sessionID)
			if gotUserID != tt.wantUserID {
				t.Errorf("GetUserID() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("GetUserID() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}

func TestGetUserData(t *testing.T) {
	ClearCache()
	const TESTING_USER_ID = "____"
	sessID := GenerateSessionID(TESTING_USER_ID, true, "")
	_ = SetUserData(sessID, "key1", "data1")
	_ = SetUserData(sessID, "keyCover", "data2")
	_ = SetUserData(sessID, "keyCover", "data3")
	type args struct {
		sessionID string
		key       string
	}
	tests := []struct {
		name        string
		args        args
		wantData    interface{}
		wantSuccess bool
	}{
		// TODO: Add test cases.
		{
			name:        "case1",
			args:        args{sessionID: sessID, key: "key1"},
			wantData:    "data1",
			wantSuccess: true,
		},
		{
			name:        "case2",
			args:        args{sessionID: sessID, key: "keyCover"},
			wantData:    "data3",
			wantSuccess: true,
		},
		{
			name:        "case3",
			args:        args{sessionID: sessID, key: "keyFaild"},
			wantData:    nil,
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotData, gotSuccess := GetUserData(tt.args.sessionID, tt.args.key)
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetUserData() gotData = %v, want %v", gotData, tt.wantData)
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("GetUserData() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}
