//Package to process the information of the database
package db

import(
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"

	"weather_app/internal/models"
	"weather_app/config"
)

//Interface for working with a mockdb
type DBInterface interface{
	SaveQuery(query *models.WeatherQuery) error 
	GetWeatherHistory()([]models.WeatherQuery,error)
}
//Database structure
type Database struct{
	Conn *sql.DB
}
//Connects to the DB
func New(cfg *config.Config)(*Database,error){
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",cfg.DBUser,cfg.DBPass,cfg.DBHost,cfg.DBName)
	db, err := sql.Open("mysql",dns)
	if err != nil{
		return nil, fmt.Errorf("could not open DB: %w",err)
	}

	if err := db.Ping();err != nil{
		return nil, fmt.Errorf("could not ping DB: %w", err)
	}

	log.Println("weather_app - db: Database connected")
	return &Database{Conn: db},nil
}
//Closes the connection to the database
func (d *Database) Close() error{
	return d.Conn.Close()
}
//Insert the request in the database
func (d *Database) SaveQuery(query *models.WeatherQuery) error{
	stmt := `INSERT INTO weather_queries (city, temperature, description, queried_at, ip_address)
	         VALUES (?, ?, ?, ?, ?)`

	_, err := d.Conn.Exec(stmt, query.City, query.Temperature, query.Description, query.Timestamp, query.IP)
	if err != nil{
		return fmt.Errorf("Could not save query: %w",err)
	}
	return nil
}
//Gets the history od queries from the database as a list of object of weatherquery
func (d *Database) GetWeatherHistory() ([]models.WeatherQuery, error) {
	rows, err := d.Conn.Query(`
		SELECT id, city, temperature, description, queried_at, ip_address 
		FROM weather_queries 
		ORDER BY queried_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.WeatherQuery

	for rows.Next() {
		var q models.WeatherQuery

		err := rows.Scan(&q.ID, &q.City, &q.Temperature, &q.Description, &q.Timestamp, &q.IP)
		if err != nil {
			log.Println("weather_app - db: Error scanning row:", err)
			continue
		}

		history = append(history, q)
	}

	return history, nil
}
