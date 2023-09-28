package application

import (
	"github.com/earelin/pega/pkg/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetComunidadesAutonomas(t *testing.T) {
	router := gin.Default()
	repository := new(EntidadesAdministrativasRepositoryMock)
	NewEntidadesAdministrativasController(router, repository)

	repository.On("FindAllComunidadesAutonomas").
		Return(comunidadesAutonomas)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/comunidades-autonomas", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, comunidadesAutonomasResponse, w.Body.String())
}

func TestGetProvincias(t *testing.T) {
	router := gin.Default()
	repository := new(EntidadesAdministrativasRepositoryMock)
	NewEntidadesAdministrativasController(router, repository)

	repository.On("FindAllProvincias").
		Return(provincias)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/provincias", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, provinciasResponse, w.Body.String())
}

func TestGetComunidadesAutonomaProvincias(t *testing.T) {
	router := gin.Default()
	repository := new(EntidadesAdministrativasRepositoryMock)
	NewEntidadesAdministrativasController(router, repository)

	repository.On("FindAllProvinciasByComunidadeAutonoma", 1).
		Return(provincias)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/comunidade-autonoma/1/provincias", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, provinciasResponse, w.Body.String())
}

func TestGetComunidadesAutonomaProvincias_ComunidadeAutonomaNotFound(t *testing.T) {
	router := gin.Default()
	repository := new(EntidadesAdministrativasRepositoryMock)
	NewEntidadesAdministrativasController(router, repository)

	repository.On("FindAllProvinciasByComunidadeAutonoma", 1).
		Return([]domain.EntidadeAdministrativa{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/comunidade-autonoma/1/provincias", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestGetConcellosProvincia(t *testing.T) {
	router := gin.Default()
	repository := new(EntidadesAdministrativasRepositoryMock)
	NewEntidadesAdministrativasController(router, repository)

	repository.On("FindAllConcellosByProvincia", 1).
		Return(concellos)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/provincia/1/concellos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, concellosResponse, w.Body.String())
}

func TestGetConcellosProvincia_ProvinciaNotFound(t *testing.T) {
	router := gin.Default()
	repository := new(EntidadesAdministrativasRepositoryMock)
	NewEntidadesAdministrativasController(router, repository)

	repository.On("FindAllConcellosByProvincia", 1).
		Return([]domain.EntidadeAdministrativa{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/provincia/1/concellos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

var comunidadesAutonomas = []domain.EntidadeAdministrativa{
	{
		Id:   1,
		Nome: "Galicia",
	},
	{
		Id:   2,
		Nome: "Asturias",
	},
}

type EntidadesAdministrativasRepositoryMock struct {
	mock.Mock
}

func (m *EntidadesAdministrativasRepositoryMock) FindAllComunidadesAutonomas() []domain.EntidadeAdministrativa {
	args := m.Called()
	return args.Get(0).([]domain.EntidadeAdministrativa)
}

func (m *EntidadesAdministrativasRepositoryMock) FindAllProvincias() []domain.EntidadeAdministrativa {
	args := m.Called()
	return args.Get(0).([]domain.EntidadeAdministrativa)
}

func (m *EntidadesAdministrativasRepositoryMock) FindAllProvinciasByComunidadeAutonoma(comunidadeAutonomaId int) []domain.EntidadeAdministrativa {
	args := m.Called(comunidadeAutonomaId)
	return args.Get(0).([]domain.EntidadeAdministrativa)
}

func (m *EntidadesAdministrativasRepositoryMock) FindAllConcellosByProvincia(provinciaId int) []domain.EntidadeAdministrativa {
	args := m.Called(provinciaId)
	return args.Get(0).([]domain.EntidadeAdministrativa)
}

func (m *EntidadesAdministrativasRepositoryMock) FindAllConcellosByName(name string) []domain.EntidadeAdministrativa {
	args := m.Called(name)
	return args.Get(0).([]domain.EntidadeAdministrativa)
}

var comunidadesAutonomasResponse = `[{"id":1,"nome":"Galicia"},{"id":2,"nome":"Asturias"}]`

var provincias = []domain.EntidadeAdministrativa{
	{
		Id:   1,
		Nome: "A Coruña",
	},
	{
		Id:   2,
		Nome: "Lugo",
	},
}

var provinciasResponse = `[{"id":1,"nome":"A Coruña"},{"id":2,"nome":"Lugo"}]`

var concellos = []domain.EntidadeAdministrativa{
	{
		Id:   1,
		Nome: "A Coruña",
	},
	{
		Id:   2,
		Nome: "Santiago de Compostela",
	},
}

var concellosResponse = `[{"id":1,"nome":"A Coruña"},{"id":2,"nome":"Santiago de Compostela"}]`
