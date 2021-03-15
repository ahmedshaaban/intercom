package customer

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type mockHttp struct {
	resp *http.Response
	err  error
}

func (m *mockHttp) Get(url string) (resp *http.Response, err error) {
	return m.resp, m.err
}

func TestRepo_fillCustomers(t *testing.T) {
	type fields struct {
		client    client
		customers []Customer
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "happy case",
			fields: fields{client: &mockHttp{
				resp: &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`))),
				},
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				client:    tt.fields.client,
				customers: tt.fields.customers,
			}
			if err := r.fillCustomers(); (err != nil) != tt.wantErr {
				t.Errorf("Repo.fillCustomers() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	mClient := &mockHttp{
		resp: &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`))),
		},
	}

	type args struct {
		client client
	}
	tests := []struct {
		name    string
		args    args
		want    *Repo
		wantErr bool
	}{
		{
			name:    "happy path",
			args:    args{client: mClient},
			want:    &Repo{client: mClient, customers: []Customer{{ID: 12, Name: "Christina McArdle", Latitude: "52.986375", Longitude: "-6.043701"}}},
			wantErr: false,
		},
		{
			name:    "happy path",
			args:    args{client: &mockHttp{err: errors.New("mock err")}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New("", tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepo_GetCustomers(t *testing.T) {
	type fields struct {
		client    client
		customers []Customer
	}
	tests := []struct {
		name   string
		fields fields
		want   []Customer
	}{
		{
			name: "",
			fields: fields{
				client: &mockHttp{
					resp: &http.Response{
						StatusCode: 200,
						Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"latitude": "52.986375", "user_id": 12, "name": "Christina McArdle", "longitude": "-6.043701"}`))),
					},
				},
				customers: []Customer{{ID: 12, Name: "Christina McArdle", Latitude: "52.986375", Longitude: "-6.043701"}},
			},
			want: []Customer{{ID: 12, Name: "Christina McArdle", Latitude: "52.986375", Longitude: "-6.043701"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				client:    tt.fields.client,
				customers: tt.fields.customers,
			}
			if got := r.GetCustomers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.GetCustomers() = %v, want %v", got, tt.want)
			}
		})
	}
}
