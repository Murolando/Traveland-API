package ent

type PlaceQueryParams struct {
	// sortList   []string 
	SortBy      string // sort_by = str  name, price, avg_rating, rating_count
	SortOrder   string // sort_order = str asc,desc
	Offset      int    // offset = int
	Limit       int    // limit = int
	PlaceTypeId int    // place_type_id = int
	HouseTypeId int    // house_type_id = int
	SearchStr 	string
}
// sort_by = str  name, price, avg_rating, rating_count
// sort_order = str asc,desc
// offset = int
// limit = int
// place_type_id = int
// house_type_id = int


type ReviewQueryParams struct {
	// SortBy      string // sort_by = str  name, price, avg_rating, rating_count
	// SortOrder   string // sort_order = str asc,desc
	PlaceId int // place_id = int
	GuideId int	// guide_id = int
	Offset  int // offset = int
	Limit   int // limit = int
}

type TourQueryParams struct {
	// SortBy      string // sort_by = str  name, price, avg_rating, rating_count
	// SortOrder   string // sort_order = str asc,desc
	Offset  int // offset = int
	Limit   int // limit = int
}