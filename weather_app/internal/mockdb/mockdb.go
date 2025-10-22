package mockdb
// A fake package to simulate interaction with a db
import(
	"weather_app/internal/models"
)

type MockDB struct{}

func (m *MockDB) SaveQuery(query *models.WeatherQuery) error{
	println("Mock SaveQuery: ",query.City)
	return nil
}

func (m *MockDB) GetWeatherHistory()([]models.WeatherQuery,error){
	var history []models.WeatherQuery
	println("Mock HistoryQuery ")
	return history,nil
}