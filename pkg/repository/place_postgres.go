package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type PlaceBD struct {
	db *sqlx.DB
}

func NewPlaceBD(db *sqlx.DB) *PlaceBD {
	return &PlaceBD{
		db: db,
	}
}

func (r PlaceBD) GetPlaceByID(id int) (interface{}, error) {
	// take place types
	query := fmt.Sprintf("SELECT type_id FROM \"%s\" WHERE place_id = $1", placeTypeTable)
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	// looking for housing or events or others
	ind := 3
	for rows.Next() {
		tp := -1
		if err := rows.Scan(&tp); err != nil {
			return nil, err
		}
		if tp == 1 {
			ind = 1
			break
		}
		if tp == 2 {
			ind = 2
			break
		}
	}
	// parse in struct
	// return struct
	switch ind {
	case 1:
		var house ent.Housing
		query = fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,house_price,house_type_id,count_room,square,mail,site_url FROM \"%s\" WHERE id = $1", placeTable)
		row := r.db.QueryRow(query, id)
		if err := row.Scan(&house.PlaceInfo.PlaceId, &house.PlaceInfo.Name, &house.PlaceInfo.Description, &house.PlaceInfo.Longitude, &house.PlaceInfo.Latitude, &house.PlaceInfo.Adress, &house.PlaceInfo.Number, &house.HousePrice, &house.HouseTypeId, &house.CountRoom, &house.Square, &house.PlaceInfo.Mail, &house.PlaceInfo.Url); err != nil {
			return ent.Housing{}, err
		}
		// house.PlaceInfo.Photos,err = r.getAllPhotos(house.PlaceInfo.PlaceId)
		// if err != nil{
		// 	return ent.Housing{},err
		// }
		return house, nil
	case 2:
		var event ent.Event
		query = fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,pushkin,min_price FROM \"%s\" WHERE id = $1", placeTable)

		row := r.db.QueryRow(query, id)
		if err := row.Scan(&event.PlaceInfo.PlaceId, &event.PlaceInfo.Name, &event.PlaceInfo.Description, &event.PlaceInfo.Longitude, &event.PlaceInfo.Latitude, &event.PlaceInfo.Adress, &event.PlaceInfo.Number, &event.Pushkin, &event.Price); err != nil {
			return ent.Event{}, err
		}
		// event.PlaceInfo.Photos,err = r.getAllPhotos(event.PlaceInfo.PlaceId)
		// if err != nil{
		// 	return ent.Event{},err
		// }
		return event, nil
	default:
		var location ent.Location
		query = fmt.Sprintf("SELECT id,name,description,location_long,location_lat,address,numbers,pushkin,min_price FROM \"%s\" WHERE id = $1", placeTable)
		row := r.db.QueryRow(query, id)
		if err := row.Scan(&location.PlaceInfo.PlaceId, &location.PlaceInfo.Name, &location.PlaceInfo.Description, &location.PlaceInfo.Longitude, &location.PlaceInfo.Latitude, &location.PlaceInfo.Adress, &location.PlaceInfo.Number, &location.Pushkin, &location.MinPrice); err != nil {
			return ent.Location{}, err
		}
		// location.PlaceInfo.Photos,err = r.getAllPhotos(location.PlaceInfo.PlaceId)
		// if err != nil{
		// 	return ent.Location{},err
		// }
		return location, nil
	}
}

func (r PlaceBD) GetAllPlaces(placeInd int, offset int) (interface{}, error) {
	switch placeInd {
	case 1:
		houses, err := r.getAllHousing(limit, offset)
		if err != nil {
			return nil, err
		}
		return houses, nil
	case 2:
		events, err := r.getAllEvents(limit, offset)
		if err != nil {
			return nil, err
		}
		return events, nil
	default:
		locals, err := r.getAllLocations(limit, offset)
		if err != nil {
			return nil, err
		}
		return locals, nil
	}

}

func (r PlaceBD) getAllHousing(limit int, offset int) (*[]ent.Housing, error) {
	houses := make([]ent.Housing, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.location_long,place.location_lat,place.address,place.house_price,place.house_type_id FROM "%s" INNER JOIN "%s" ON place.id = place_type.place_id WHERE place_type.type_id = $1 LIMIT $2 OFFSET $3`, placeTable, placeTypeTable)
	rows, err := r.db.Query(query, 1, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var house ent.Housing

		if err := rows.Scan(&house.PlaceInfo.PlaceId, &house.PlaceInfo.Name, &house.PlaceInfo.Longitude, &house.PlaceInfo.Latitude, &house.PlaceInfo.Adress, &house.HousePrice, &house.HouseTypeId); err != nil {
			return nil, err
		}

		query = fmt.Sprintf(`SELECT COUNT(id) FROM "%s" WHERE place_id = $1`, reviewTable)
		row := r.db.QueryRow(query, house.PlaceInfo.PlaceId)
		if err := row.Scan(&house.PlaceInfo.NumberOfRating); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT AVG(rating) FROM "%s" WHERE place_id = $1`, reviewTable)
		row = r.db.QueryRow(query, house.PlaceInfo.PlaceId)
		if err := row.Scan(&house.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}
		// pht,err := r.getPhoto(house.PlaceInfo.PlaceId)
		// if err!=nil{
		// 	return nil, err
		// }
		// house.PlaceInfo.Photos = append(house.PlaceInfo.Photos,pht)
		houses = append(houses, house)
	}
	return &houses, nil
}

func (r PlaceBD) getAllEvents(limit int, offset int) (*[]ent.Event, error) {
	events := make([]ent.Event, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.location_long,place.location_lat,place.address,place.min_price FROM "%s" INNER JOIN "%s" ON place.id = place_type.place_id WHERE place_type.type_id = $1 LIMIT $2 OFFSET $3`, placeTable, placeTypeTable)
	rows, err := r.db.Query(query, 2, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var event ent.Event

		if err := rows.Scan(&event.PlaceInfo.PlaceId, &event.PlaceInfo.Name, &event.PlaceInfo.Longitude, &event.PlaceInfo.Latitude, &event.PlaceInfo.Adress, &event.Price); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT COUNT(id) FROM "%s" WHERE place_id = $1`, reviewTable)
		row := r.db.QueryRow(query, event.PlaceInfo.PlaceId)
		if err := row.Scan(&event.PlaceInfo.NumberOfRating); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT AVG(rating) FROM "%s" WHERE place_id = $1`, reviewTable)
		row = r.db.QueryRow(query, event.PlaceInfo.PlaceId)
		if err := row.Scan(&event.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}
		// pht,err := r.getPhoto(event.PlaceInfo.PlaceId)
		// if err!=nil{
		// 	return nil, err
		// }
		// event.PlaceInfo.Photos = append(event.PlaceInfo.Photos,pht)
		events = append(events, event)
	}
	return &events, nil
}

func (r PlaceBD) getAllLocations(limit int, offset int) (*[]ent.Location, error) {
	locations := make([]ent.Location, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.location_long,place.location_lat,place.address,place.min_price FROM "%s" INNER JOIN "%s" ON place.id = place_type.place_id WHERE NOT place_type.type_id = $1 and NOT place_type.type_id = $2 LIMIT $3 OFFSET $4`, placeTable, placeTypeTable)
	rows, err := r.db.Query(query, 2, 1, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var location ent.Location

		if err := rows.Scan(&location.PlaceInfo.PlaceId, &location.PlaceInfo.Name, &location.PlaceInfo.Longitude, &location.PlaceInfo.Latitude, &location.PlaceInfo.Adress, &location.MinPrice); err != nil {
			return nil, err
		}

		query = fmt.Sprintf(`SELECT COUNT(id) FROM "%s" WHERE place_id = $1`, reviewTable)
		row := r.db.QueryRow(query, location.PlaceInfo.PlaceId)
		if err := row.Scan(&location.PlaceInfo.NumberOfRating); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT AVG(rating) FROM "%s" WHERE place_id = $1`, reviewTable)
		row = r.db.QueryRow(query, location.PlaceInfo.PlaceId)
		if err := row.Scan(&location.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}

		// pht,err := r.getPhoto(location.PlaceInfo.PlaceId)
		// if err!=nil{
		// 	return nil, err
		// }
		// location.PlaceInfo.Photos = append(location.PlaceInfo.Photos,pht)
		locations = append(locations, location)
	}
	return &locations, nil
}

func (r PlaceBD) GetLocalByType(placeType int, offset int) (*[]ent.Location, error) {
	places := make([]ent.Location, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.location_long,place.location_lat,place.address,place.min_price FROM "%s" INNER JOIN "%s" ON place.id = place_type.place_id WHERE place_type.type_id = $1 LIMIT $2 OFFSET $3`, placeTable, placeTypeTable)
	rows, err := r.db.Query(query, placeType, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var location ent.Location

		if err := rows.Scan(&location.PlaceInfo.PlaceId, &location.PlaceInfo.Name, &location.PlaceInfo.Longitude, &location.PlaceInfo.Latitude, &location.PlaceInfo.Adress, &location.MinPrice); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT COUNT(id) FROM "%s" WHERE place_id = $1`, reviewTable)
		row := r.db.QueryRow(query, location.PlaceInfo.PlaceId)
		if err := row.Scan(&location.PlaceInfo.NumberOfRating); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT AVG(rating) FROM "%s" WHERE place_id = $1`, reviewTable)
		row = r.db.QueryRow(query, location.PlaceInfo.PlaceId)
		if err := row.Scan(&location.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}
		// pht,err := r.getPhoto(location.PlaceInfo.PlaceId)
		// if err!=nil{
		// 	return nil, err
		// }
		// location.PlaceInfo.Photos = append(location.PlaceInfo.Photos,pht)
		places = append(places, location)
	}
	return &places, nil
}

func (r PlaceBD) GetHouseByType(houseType int, offset int) (*[]ent.Housing, error) {
	houses := make([]ent.Housing, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.location_long,place.location_lat,place.address,place.house_price FROM "%s" WHERE place.house_type_id = $1 LIMIT $2 OFFSET $3`, placeTable)
	rows, err := r.db.Query(query, houseType, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var house ent.Housing

		if err := rows.Scan(&house.PlaceInfo.PlaceId, &house.PlaceInfo.Name, &house.PlaceInfo.Longitude, &house.PlaceInfo.Latitude, &house.PlaceInfo.Adress, &house.HousePrice); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT COUNT(id) FROM "%s" WHERE place_id = $1`, reviewTable)
		row := r.db.QueryRow(query, house.PlaceInfo.PlaceId)
		if err := row.Scan(&house.PlaceInfo.NumberOfRating); err != nil {
			return nil, err
		}
		query = fmt.Sprintf(`SELECT AVG(rating) FROM "%s" WHERE place_id = $1`, reviewTable)
		row = r.db.QueryRow(query, house.PlaceInfo.PlaceId)
		if err := row.Scan(&house.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}
		// pht,err := r.getPhoto(house.PlaceInfo.PlaceId)
		// if err!=nil{
		// 	return nil, err
		// }
		// house.PlaceInfo.Photos = append(house.PlaceInfo.Photos,pht)

		houses = append(houses, house)
	}
	return &houses, nil
}
func (r PlaceBD) GetLocalTypes() (*[]ent.LocalType, error) {
	localTypes := make([]ent.LocalType, 0)
	query := fmt.Sprintf(`SELECT id,name FROM "%s"`, typeTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var localType ent.LocalType

		if err := rows.Scan(&localType.TypeId, &localType.Name); err != nil {
			return nil, err
		}

		localTypes = append(localTypes, localType)
	}
	return &localTypes, nil
}
func (r PlaceBD) GetHouseTypes() (*[]ent.HouseType, error) {
	houseTypes := make([]ent.HouseType, 0)
	query := fmt.Sprintf(`SELECT id,name FROM "%s"`, houseTypeTable)
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var houseType ent.HouseType

		if err := rows.Scan(&houseType.HouseTypeId, &houseType.Name); err != nil {
			return nil, err
		}

		houseTypes = append(houseTypes, houseType)
	}
	return &houseTypes, nil
}

func (r PlaceBD) getAllPhotos(placeId int) ([]string, error) {
	var photos []string
	err := filepath.Walk(fmt.Sprintf("./storage/place/%d",placeId),
	// err := filepath.Walk("./storage/place",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				photos = append(photos, path)
			}
			return nil
		})
	if err != nil {
		return nil,err
	}
	return photos, nil
}

func (r PlaceBD) getPhoto(placeId int) (string, error) {
	var photo string
	err := filepath.Walk(fmt.Sprintf("./storage/place/%d",placeId),
	// err := filepath.Walk("./storage/place",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				photo = path
				return nil
			}
			return nil
		})
	if err != nil {
		return "",err
	}
	return photo, nil
}
