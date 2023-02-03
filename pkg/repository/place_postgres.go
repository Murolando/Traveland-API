package repository

import (
	"fmt"
	"strconv"

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

// sort_by = str  name, price, avg_rating, rating_count
// sort_order = str asc,desc
// offset = int 0...n
// limit = int	0...n
// place_type_id = int (3...n)
// house_type_id = int (1...n)

func (r PlaceBD) likeStr(str string) string{
	if str == ""{
		return ``
	}else{
		line := "%%"+str+"%%"
		query := fmt.Sprintf(`AND (name LIKE '%s' OR description LIKE '%s') `,line,line)
		return query
	}
}
func (r PlaceBD) sortByOrder(srtBy string, srtOrder string) string {
	query := fmt.Sprintf(`ORDER BY %s %s `, srtBy, srtOrder)
	return query
}
func (r PlaceBD) getAllPhotos(placeId int) ([]string,error){
	var photos []string = make([]string,0 )
	query:= fmt.Sprintf(`
	SELECT image_src 
	FROM "%s" 
	WHERE place_id = $1`,placeSrcTable)
	rows,err := r.db.Query(query,placeId)
	if err!=nil{
		return nil,err
	}
	for rows.Next(){
		var photo string
		if err := rows.Scan(&photo);err!=nil{
			return nil,err
		}
		photos = append(photos, photo)
	}
	return photos,nil
}
func (r PlaceBD) houseType(houseType int) string {
	if houseType != 0 {
		query := fmt.Sprintf(`AND place.house_type_id = %s `, strconv.Itoa(houseType))
		return query
	}
	return ``
}
func (r PlaceBD) localType(placeType int) string {
	if placeType != 0 {
		query := fmt.Sprintf(`AND place_type.type_id =  %s `, strconv.Itoa(placeType))
		return query
	}
	return ``
}
func (r PlaceBD) getShedule(placeId int)([]ent.Shedule,error){
	query := fmt.Sprintf(`SELECT place_id,day_id,start_work,end_work,start_timeout,end_timeout
	FROM "%s"
	WHERE place_id = $1`,weekTable)
	newRows, err := r.db.Query(query,placeId)
	if err != nil {
		return nil, err
	}
	shedulers := make([]ent.Shedule,0)
	for newRows.Next(){
		var	shedule ent.Shedule
		if err := newRows.Scan(&shedule.PlaceId,
			&shedule.WeekDay,&shedule.StartWork,&shedule.EndWork,
			&shedule.StartTimeOut,&shedule.EndTimeOut); err != nil {
			return nil, err
		}
		shedulers = append(shedulers,shedule)
	}
	return shedulers,nil
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
		query = fmt.Sprintf(`SELECT id,name,
		description,location_long,
		location_lat,address,numbers,
		house_price,house_type_id,mail,site_url 
		FROM \"%s\"
		 WHERE id = $1`, placeTable)
		row := r.db.QueryRow(query, id)
		if err := row.Scan(&house.PlaceInfo.PlaceId, &house.PlaceInfo.Name, 
			&house.PlaceInfo.Description,
			 &house.PlaceInfo.Longitude, &house.PlaceInfo.Latitude,
			  &house.PlaceInfo.Adress, &house.PlaceInfo.Number, 
			  &house.HousePrice, &house.HouseTypeId,
			  &house.PlaceInfo.Mail, &house.PlaceInfo.Url); err != nil {
			return ent.Housing{}, err
		}
		// house.PlaceInfo.Photos,err = r.getAllPhotos(house.PlaceInfo.PlaceId)
		// if err != nil{
		// 	return ent.Housing{},err
		// }
		return house, nil
	case 2:
		var event ent.Event
		query = fmt.Sprintf(`
		SELECT id,name,description,
		location_long,location_lat,
		address,numbers,pushkin
		,min_price 
		FROM \"%s\" 
		WHERE id = $1`, placeTable)

		row := r.db.QueryRow(query, id)
		if err := row.Scan(&event.PlaceInfo.PlaceId, &event.PlaceInfo.Name,
			 &event.PlaceInfo.Description, &event.PlaceInfo.Longitude, 
			 &event.PlaceInfo.Latitude, &event.PlaceInfo.Adress,
			  &event.PlaceInfo.Number, &event.Pushkin, &event.Price); err != nil {
			return ent.Event{}, err
		}
		// event.PlaceInfo.Photos,err = r.getAllPhotos(event.PlaceInfo.PlaceId)
		// if err != nil{
		// 	return ent.Event{},err
		// }
		return event, nil
	default:
		var location ent.Location
		query = fmt.Sprintf(`
		SELECT id,name,description,location_long,
		location_lat,address,numbers,pushkin,min_price 
		FROM \"%s\" 
		WHERE id = $1`, placeTable)
		row := r.db.QueryRow(query, id)
		if err := row.Scan(&location.PlaceInfo.PlaceId, 
			&location.PlaceInfo.Name, &location.PlaceInfo.Description,
			 &location.PlaceInfo.Longitude, &location.PlaceInfo.Latitude,
			  &location.PlaceInfo.Adress, &location.PlaceInfo.Number, 
			  &location.Pushkin, &location.MinPrice); err != nil {
			return ent.Location{}, err
		}
		// location.PlaceInfo.Photos,err = r.getAllPhotos(location.PlaceInfo.PlaceId)
		// if err != nil{
		// 	return ent.Location{},err
		// }
		return location, nil
	}
}

func (r PlaceBD) GetAllPlaces(placeInd int, params *ent.PlaceQueryParams) (interface{}, error) {
	switch placeInd {
	case 1:
		houses, err := r.getAllHousing(params)
		if err != nil {
			return nil, err
		}
		return houses, nil
	case 2:
		events, err := r.getAllEvents(params)
		if err != nil {
			return nil, err
		}
		return events, nil
	default:
		locals, err := r.getAllLocations(params)
		if err != nil {
			return nil, err
		}
		return locals, nil
	}

}
func (r PlaceBD) getAllHousing(params *ent.PlaceQueryParams) (*[]ent.Housing, error) {
	houses := make([]ent.Housing, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.description,
	place.location_long,place.location_lat,place.address,place.numbers,
	place.mail,place.site_url,
	place.house_price,
	place.house_type_id,
	(SELECT 
	 COUNT(id) 
	 FROM review 
	 WHERE place_id = place.id) AS rating_count,
	 (SELECT 
	 AVG(rating)
	 FROM review 
	 WHERE place_id = place.id) AS avg_rating
	FROM "%s" 
	INNER JOIN "%s" ON place.id = place_type.place_id 
	WHERE place_type.type_id = $1 `+
		r.houseType(params.HouseTypeId)+
		r.likeStr(params.SearchStr)+
		r.sortByOrder(params.SortBy, params.SortOrder)+
		`LIMIT $2
	OFFSET $3`, placeTable, placeTypeTable)
	// fmt.Println(query)
	rows, err := r.db.Query(query, 1, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		var house ent.Housing
		if err := rows.Scan(&house.PlaceInfo.PlaceId, 
			&house.PlaceInfo.Name, &house.PlaceInfo.Description,
			&house.PlaceInfo.Longitude, 
			&house.PlaceInfo.Latitude, &house.PlaceInfo.Adress, 
			&house.PlaceInfo.Number,&house.PlaceInfo.Mail,&house.PlaceInfo.Url,
			&house.HousePrice, &house.HouseTypeId, &house.PlaceInfo.NumberOfRating,
			&house.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}

		house.Shedule,err = r.getShedule(house.PlaceInfo.PlaceId)
		if err!=nil{
			return nil,err
		}
		house.PlaceInfo.Photos,err = r.getAllPhotos(house.PlaceInfo.PlaceId)
		if err!=nil{
			return nil,err
		}

		houses = append(houses, house)
	}
	return &houses, nil
}

func (r PlaceBD) getAllEvents(params *ent.PlaceQueryParams) (*[]ent.Event, error) {
	events := make([]ent.Event, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.description,
	place.location_long,place.location_lat,place.address,place.numbers,
	place.mail,place.site_url,place.min_price,
	place.event_day,place.event_start_time,place.event_end_time,
	(SELECT 
	 COUNT(id) 
	 FROM review 
	 WHERE place_id = place.id) AS rating_count,
	 (SELECT 
	 AVG(rating)
	 FROM review 
	 WHERE place_id = place.id) AS avg_rating
	FROM "%s" 
	INNER JOIN "%s" ON place.id = place_type.place_id 
	WHERE place_type.type_id = $1 `+
		r.likeStr(params.SearchStr)+
		r.sortByOrder(params.SortBy, params.SortOrder)+
		`LIMIT $2
	OFFSET $3`, placeTable, placeTypeTable)
	rows, err := r.db.Query(query, 2,params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var event ent.Event

		if err := rows.Scan(&event.PlaceInfo.PlaceId, &event.PlaceInfo.Name, 
			&event.PlaceInfo.Description,&event.PlaceInfo.Longitude, 
			&event.PlaceInfo.Latitude,&event.PlaceInfo.Adress, 
			&event.PlaceInfo.Number,&event.PlaceInfo.Mail,&event.PlaceInfo.Url,
			&event.Price,&event.EventDay,&event.EventStartTime,
			&event.EventEndTime,&event.PlaceInfo.NumberOfRating,
			&event.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}
		event.PlaceInfo.Photos,err = r.getAllPhotos(event.PlaceInfo.PlaceId)
		if err!=nil{
			return nil,err
		}
		events = append(events, event)
	}
	return &events, nil
}

func (r PlaceBD) getAllLocations(params *ent.PlaceQueryParams) (*[]ent.Location, error) {
	locations := make([]ent.Location, 0)
	query := fmt.Sprintf(`SELECT place.id,
	place.name,place.description,place.location_long,
	place.location_lat,place.address,
	place.numbers,
	place.mail,place.site_url,
	place.min_price,
	(SELECT 
	 COUNT(id) 
	 FROM review 
	 WHERE place_id = place.id) AS rating_count,
	 (SELECT 
	 AVG(rating)
	 FROM review 
	 WHERE place_id = place.id) AS avg_rating
	FROM "%s" 
	INNER JOIN "%s" ON place.id = place_type.place_id 
	WHERE NOT place_type.type_id = $1 and NOT place_type.type_id = $2 `+
	r.localType(params.PlaceTypeId)+
	r.likeStr(params.SearchStr)+
	r.sortByOrder(params.SortBy, params.SortOrder)+
	`LIMIT $3
	OFFSET $4`, placeTable, placeTypeTable)
	rows, err := r.db.Query(query, 2, 1, params.Limit,params.Offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var location ent.Location

		if err := rows.Scan(&location.PlaceInfo.PlaceId, &location.PlaceInfo.Name,
			&location.PlaceInfo.Description, 
			&location.PlaceInfo.Longitude, &location.PlaceInfo.Latitude, 
			&location.PlaceInfo.Adress, &location.PlaceInfo.Number,
			&location.PlaceInfo.Mail,&location.PlaceInfo.Url,&location.MinPrice,
			&location.PlaceInfo.NumberOfRating,&location.PlaceInfo.MeanRating); err != nil {
			return nil, err
		}


		location.Shedule,err = r.getShedule(location.PlaceInfo.PlaceId)
		if err!=nil{
			return nil,err
		}
		location.PlaceInfo.Photos,err = r.getAllPhotos(location.PlaceInfo.PlaceId)
		if err!=nil{
			return nil,err
		}
		locations = append(locations, location)
	}
	
	return &locations, nil
}




func (r PlaceBD) GetLocalByType(placeType int, offset int) (*[]ent.Location, error) {
	places := make([]ent.Location, 0)
	query := fmt.Sprintf(`SELECT place.id,place.name,place.location_long,
	place.location_lat,place.address,place.min_price 
	FROM "%s" INNER JOIN "%s" ON place.id = place_type.place_id 
	WHERE place_type.type_id = $1 LIMIT $2 OFFSET $3`, placeTable, placeTypeTable)
	rows, err := r.db.Query(query, placeType, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var location ent.Location

		if err := rows.Scan(&location.PlaceInfo.PlaceId, 
			&location.PlaceInfo.Name, &location.PlaceInfo.Longitude, 
			&location.PlaceInfo.Latitude, &location.PlaceInfo.Adress, &location.MinPrice); err != nil {
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

func (r PlaceBD) AddFavoritePlace(userId int, placeId int) (bool, error) {
	query := fmt.Sprintf("INSERT INTO \"%s\" (user_id,place_id) values ($1,$2)", favoritePlaceTable)
	_, err := r.db.Exec(query, userId, placeId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r PlaceBD) GetAllUserFavoritePlaces(userId int) (*[]interface{}, error) {
	query := fmt.Sprintf("SELECT place_id FROM \"%s\" WHERE user_id = $1", favoritePlaceTable)
	var favPlaces []interface{} = make([]interface{}, 0)
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var placeId int
		if err := rows.Scan(&placeId); err != nil {
			return nil, err
		}
		plc, err := r.GetPlaceByID(placeId)
		if err != nil {
			return nil, err
		}
		favPlaces = append(favPlaces, plc)
	}
	return &favPlaces, nil
}

func (r PlaceBD) GetCountOfPlaceFavorites(placeId int) (int, error) {
	var counts int
	query := fmt.Sprintf("SELECT COUNT(id) FROM \"%s\" WHERE place_id = $1", favoritePlaceTable)
	row := r.db.QueryRow(query, placeId)
	if err := row.Scan(&counts); err != nil {
		return 0, err
	}
	return counts, nil
}
