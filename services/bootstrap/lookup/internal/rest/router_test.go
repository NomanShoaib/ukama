package rest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"ukamaX/bootstrap/lookup/internal/db"
	"ukamaX/bootstrap/lookup/mocks"

	"github.com/ukama/ukama/services/common/ukama"

	"github.com/jackc/pgtype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var testNodeId = ukama.NewVirtualNodeId(ukama.NODE_ID_TYPE_HOMENODE)

const dummyOrgName = "some_org"

var testIp = pgtype.Inet{
	IPNet: &net.IPNet{
		IP: []byte{1, 2, 3, 4},
	},
}

func TestPingRoute(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	r := NewRouter(nil, nil, true).gin

	// act
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}

func TestGetDeviceRouteDeviceExist(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()

	nodeRepo := &mocks.NodeRepo{}
	orgRepo := &mocks.OrgRepo{}

	dbNode := &db.Node{
		Org:    db.Org{Name: "some_org", Certificate: "some_cert", Ip: testIp},
		NodeID: testNodeId.StringLowercase(),
	}
	nodeRepo.On("Get", testNodeId).Return(dbNode, nil).Once()

	r := NewRouter(nodeRepo, orgRepo, true).gin

	// act
	req, _ := http.NewRequest("GET", "/orgs/"+dummyOrgName+"/devices/"+testNodeId.String(), nil)
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	resp := GetDeviceResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	nodeRepo.AssertExpectations(t)
	assert.Equal(t, dbNode.Org.Name, resp.OrgName)
	assert.Equal(t, dbNode.Org.Certificate, resp.Certificate)
	assert.Equal(t, dbNode.Org.Ip.IPNet.IP.String(), resp.Ip)
}

func TestGetDeviceRouteDeviceDoesNotExist(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()

	nodeRepo := &mocks.NodeRepo{}
	orgRepo := &mocks.OrgRepo{}

	nodeRepo.On("Get", testNodeId).Return(nil, gorm.ErrRecordNotFound).Once()

	r := NewRouter(nodeRepo, orgRepo, true).gin

	// act
	req, _ := http.NewRequest("GET", "/orgs/"+dummyOrgName+"/devices/"+testNodeId.String(), nil)
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 404, w.Code)
	assert.Contains(t, w.Body.String(), "not found")

	nodeRepo.AssertExpectations(t)
}

func TestGetDeviceRouteRepoError(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()

	nodeRepo := &mocks.NodeRepo{}
	orgRepo := &mocks.OrgRepo{}

	nodeRepo.On("Get", testNodeId).Return(nil, fmt.Errorf("DB failed")).Once()

	r := NewRouter(nodeRepo, orgRepo, true).gin

	// act
	req, _ := http.NewRequest("GET", "/orgs/"+dummyOrgName+"/devices/"+testNodeId.String(), nil)
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Error getting the node")

	nodeRepo.AssertExpectations(t)
}

func TestGetDeviceRouteInvalidUuid(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()
	r := NewRouter(nil, nil, true).gin

	// act
	req, _ := http.NewRequest("GET", "/orgs/"+dummyOrgName+"/devices/123asadg", nil)
	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Error parsing NodeID")
}

func TestPostDeviceRoute(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()

	nodeRepo := &mocks.NodeRepo{}
	orgRepo := &mocks.OrgRepo{}
	const orgName = "some_org"
	const modelId = 5

	orgRepo.On("GetByName", orgName).Return(&db.Org{Name: orgName, Model: gorm.Model{ID: modelId}}, nil)
	nodeRepo.On("AddOrUpdate", mock.MatchedBy(func(n *db.Node) bool {
		return n.NodeID == testNodeId.StringLowercase() && n.OrgID == modelId
	})).Return(nil).Once()

	r := NewRouter(nodeRepo, orgRepo, true).gin

	// act
	req, _ := http.NewRequest("POST", "/orgs/"+dummyOrgName+"/devices/"+testNodeId.String(),
		nil)

	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	resp := GetDeviceResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)

	nodeRepo.AssertExpectations(t)
}

func TestPostOrgRoute(t *testing.T) {
	// arrange
	w := httptest.NewRecorder()
	orgRepo := &mocks.OrgRepo{}
	const orgName = "some_org"
	data := []byte("some certificate")
	cert := base64.StdEncoding.EncodeToString(data)

	orgRepo.On("Upsert", mock.MatchedBy(func(o *db.Org) bool {
		return o.Name == orgName && o.Ip.IPNet.IP.String() == testIp.IPNet.IP.String() && o.Certificate == cert
	})).Return(nil).Once()

	r := NewRouter(nil, orgRepo, true).gin

	// act
	req, _ := http.NewRequest("POST", "/orgs/"+orgName,
		strings.NewReader(fmt.Sprintf(`{
				"certificate":"%s",
				"ip": "%s"
			}`, cert, testIp.IPNet.IP.String())))

	r.ServeHTTP(w, req)

	// assert
	assert.Equal(t, 200, w.Code)
	orgRepo.AssertExpectations(t)
}

func TestPostOrgRouteOrgValidation(t *testing.T) {
	const orgName = "some_org"
	data := []byte("some certificate")
	cert := base64.StdEncoding.EncodeToString(data)

	tests := []struct {
		name            string
		request         string
		expectedMessage string
	}{
		{name: "cert-none-base64",
			request: fmt.Sprintf(`{
				"certificate":"%s",
				"ip": "%s"
			}`, "some none base 64", testIp.IPNet.IP.String()),
			expectedMessage: "Certificate",
		},
		{name: "cert-empty",
			request: fmt.Sprintf(`{
				"certificate":"%s",
				"ip": "%s"
			}`, "", testIp.IPNet.IP.String()),
			expectedMessage: "Certificate",
		},
		{name: "bad-ip",
			request: fmt.Sprintf(`{
				"certificate":"%s",
				"ip": "%s"
			}`, cert, "123423.324234.3.2"),
			expectedMessage: "Ip",
		},
		{name: "empty-ip",
			request: fmt.Sprintf(`{
				"certificate":"%s",
				"ip": "%s"
			}`, cert, ""),
			expectedMessage: "Ip",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := NewRouter(nil, nil, true).gin

			// act
			req, _ := http.NewRequest("POST", "/orgs/"+orgName,
				strings.NewReader(test.request))

			r.ServeHTTP(w, req)

			// assert
			assert.Equal(t, 400, w.Code)
			assert.Contains(t, w.Body.String(), test.expectedMessage)
		})
	}

}