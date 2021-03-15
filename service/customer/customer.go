package customer

import (
	"math"
	"sort"
	"strconv"

	customerrepo "github.com/ahmedshaaban/intercom/repo/customer"
)

const (
	earthRaidusKm = 6371 // radius of the earth in kilometers.
	intercomLat   = 53.339428
	intercomLon   = -6.257664
)

type repo interface {
	GetCustomers() []customerrepo.Customer
}

type Service struct {
	repo repo
}

func New(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

type byID []customerrepo.Customer

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].ID < a[j].ID }

func (s *Service) SortedInvitees() []customerrepo.Customer {
	invitees := s.repo.GetCustomers()
	sortedInvitees := []customerrepo.Customer{}

	for _, v := range invitees {
		fLat, _ := strconv.ParseFloat(v.Latitude, 64)
		fLon, _ := strconv.ParseFloat(v.Longitude, 64)
		if calculateDistance(intercomLat, intercomLon, fLat, fLon) <= 100 {
			sortedInvitees = append(sortedInvitees, v)
		}
	}

	sort.Sort(byID(sortedInvitees))

	return sortedInvitees
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	rX1 := degreeToRadian(x1)
	rX2 := degreeToRadian(x2)
	rY1 := degreeToRadian(y1)
	rY2 := degreeToRadian(y2)

	diffLat := rX2 - rX1
	diffLon := rY2 - rY1

	a := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(rX1)*math.Cos(rX2)*
		math.Pow(math.Sin(diffLon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return c * earthRaidusKm
}

func degreeToRadian(d float64) float64 {
	return d * (math.Pi / 180)
}
