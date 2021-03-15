package customer

import (
	"reflect"
	"testing"

	customerrepo "github.com/ahmedshaaban/intercom/repo/customer"
)

type mockRepo struct {
	customers []customerrepo.Customer
}

func (m *mockRepo) GetCustomers() []customerrepo.Customer {
	return m.customers
}

func TestService_SortedInvitees(t *testing.T) {
	type fields struct {
		repo repo
	}
	tests := []struct {
		name   string
		fields fields
		want   []customerrepo.Customer
	}{
		{
			name:   "happy case",
			fields: fields{repo: &mockRepo{customers: []customerrepo.Customer{{ID: 1, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}, {ID: 1, Name: "test", Latitude: "73.341017", Longitude: "-6.256419"}}}},
			want:   []customerrepo.Customer{{ID: 1, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}},
		},
		{
			name:   "sorted case",
			fields: fields{repo: &mockRepo{customers: []customerrepo.Customer{{ID: 1, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}, {ID: 2, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}}}},
			want:   []customerrepo.Customer{{ID: 1, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}, {ID: 2, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}},
		},
		{
			name:   "reverse sorted case",
			fields: fields{repo: &mockRepo{customers: []customerrepo.Customer{{ID: 2, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}, {ID: 1, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}}}},
			want:   []customerrepo.Customer{{ID: 1, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}, {ID: 2, Name: "test", Latitude: "53.341017", Longitude: "-6.256419"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.fields.repo)
			if got := s.SortedInvitees(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.SortedInvitees() = %v, want %v", got, tt.want)
			}
		})
	}
}
