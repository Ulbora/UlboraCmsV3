package services

import (
	"reflect"
	"testing"
)

func TestContentService_CachePage(t *testing.T) {
	type fields struct {
		Token    string
		ClientID string
		APIKey   string
		UserID   string
		Hashed   string
		Host     string
	}
	type args struct {
		clientID string
		pageName string
		page     PageCache
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &ContentService{
			Token:    tt.fields.Token,
			ClientID: tt.fields.ClientID,
			APIKey:   tt.fields.APIKey,
			UserID:   tt.fields.UserID,
			Hashed:   tt.fields.Hashed,
			Host:     tt.fields.Host,
		}
		c.CachePage(tt.args.clientID, tt.args.pageName, tt.args.page)
	}
}

func TestContentService_ReadPage(t *testing.T) {
	type fields struct {
		Token    string
		ClientID string
		APIKey   string
		UserID   string
		Hashed   string
		Host     string
	}
	type args struct {
		clientID string
		pageName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *PageHead
		want1  *[]Content
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &ContentService{
			Token:    tt.fields.Token,
			ClientID: tt.fields.ClientID,
			APIKey:   tt.fields.APIKey,
			UserID:   tt.fields.UserID,
			Hashed:   tt.fields.Hashed,
			Host:     tt.fields.Host,
		}
		got, got1 := c.ReadPage(tt.args.clientID, tt.args.pageName)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ContentService.ReadPage() got = %v, want %v", tt.name, got, tt.want)
		}
		if !reflect.DeepEqual(got1, tt.want1) {
			t.Errorf("%q. ContentService.ReadPage() got1 = %v, want %v", tt.name, got1, tt.want1)
		}
	}
}

func TestContentService_RemovePage(t *testing.T) {
	type fields struct {
		Token    string
		ClientID string
		APIKey   string
		UserID   string
		Hashed   string
		Host     string
	}
	type args struct {
		clientID string
		pageName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &ContentService{
			Token:    tt.fields.Token,
			ClientID: tt.fields.ClientID,
			APIKey:   tt.fields.APIKey,
			UserID:   tt.fields.UserID,
			Hashed:   tt.fields.Hashed,
			Host:     tt.fields.Host,
		}
		c.RemovePage(tt.args.clientID, tt.args.pageName)
	}
}

func TestContentService_DeletePage(t *testing.T) {
	type fields struct {
		Token    string
		ClientID string
		APIKey   string
		UserID   string
		Hashed   string
		Host     string
	}
	type args struct {
		clientID string
		pageName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		c := &ContentService{
			Token:    tt.fields.Token,
			ClientID: tt.fields.ClientID,
			APIKey:   tt.fields.APIKey,
			UserID:   tt.fields.UserID,
			Hashed:   tt.fields.Hashed,
			Host:     tt.fields.Host,
		}
		c.DeletePage(tt.args.clientID, tt.args.pageName)
	}
}
