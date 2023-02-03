package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"traveland/ent"

	"github.com/gin-gonic/gin"
)
var sortList []string = []string{"name","price","avg_rating","rating_count"}
const (
	authHeader = "Authorization"
	
)


func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil || userId == 0 {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set("userId", userId)
}

// sort_by = str  name, price, avg_rating, rating_count
// sort_order = str asc,desc
// offset = int 0...n
// limit = int	0...n
// place_type_id = int (3...n)
// house_type_id = int (1...n)


func (h *Handler) placeQueryParams(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	placeQuery := ent.PlaceQueryParams{
		Limit: 20,
		Offset: 0,
		SortBy: "name",
		SortOrder: "ASC",
		SearchStr: "",
	}


	if queryParams["house_type_id"] != nil{
		houseType,err := strconv.Atoi(queryParams["house_type_id"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad house_type_id type")
			return
		}
		placeQuery.HouseTypeId = houseType
	}
		
	if queryParams["place_type_id"] != nil{
		placeType,err := strconv.Atoi(queryParams["place_type_id"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad place_type_id type")
			return
		}
		placeQuery.PlaceTypeId = placeType
	}

	if (queryParams["sort_by"] != nil){
		flag := false
		for _,val := range(sortList){
			if val == queryParams["sort_by"][0]{
				flag = true
				break
			}
		}
		if !flag{
			newErrorResponse(c, http.StatusUnauthorized, "bad sort_by type")
			return
		}
		if queryParams["sort_by"][0] == "price"{
			placeQuery.SortBy = "min_price"
		}else{
			placeQuery.SortBy = queryParams["sort_by"][0]
		}
		
	}	
	if queryParams["sort_order"] != nil{
		line := strings.ToUpper(queryParams["sort_order"][0])
		if line == "ASC" || line=="DESC"{
			placeQuery.SortOrder = line
		}else{
			newErrorResponse(c, http.StatusUnauthorized, "bad sort_order type")
			return
		}
	}

	if queryParams["offset"] != nil{
		offset,err := strconv.Atoi(queryParams["offset"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad offset type")
			return
		}
		placeQuery.Offset = offset
	}
	if queryParams["limit"] != nil{
		limit,err := strconv.Atoi(queryParams["limit"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad limit type")
			return
		}
		placeQuery.Limit = limit
	}

	if queryParams["search"] != nil{
		fmt.Println(queryParams["search"][0])
		searchStr := queryParams["search"][0]
		
		placeQuery.SearchStr = searchStr
	}


	c.Set("placeQueryParams", &placeQuery)
}

func (h *Handler) reviewQueryParams(c *gin.Context){
	queryParams := c.Request.URL.Query()
	reviewQuery := ent.ReviewQueryParams{
		Limit: 20,
		Offset: 0,
		PlaceId: -1,
		GuideId: -1,
	}
	if queryParams["place_id"] != nil{
		placeId,err := strconv.Atoi(queryParams["place_id"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad place_id type")
			return
		}
		if placeId <=0{
			newErrorResponse(c, http.StatusUnauthorized, "bad place_id type")
			return
		}
		reviewQuery.PlaceId = placeId
	}
	if queryParams["guide_id"] != nil{
		guideId,err := strconv.Atoi(queryParams["guide_id"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad guide_id type")
			return
		}
		if guideId <=0{
			newErrorResponse(c, http.StatusUnauthorized, "bad guide_id type")
			return
		}
		reviewQuery.PlaceId = guideId
	}
	if reviewQuery.PlaceId ==-1 && reviewQuery.GuideId ==-1{
		newErrorResponse(c, http.StatusUnauthorized, "not found place_id or guide_id")
		return
	}
	if queryParams["offset"] != nil{
		offset,err := strconv.Atoi(queryParams["offset"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad offset type")
			return
		}
		reviewQuery.Offset = offset
	}
	if queryParams["limit"] != nil{
		limit,err := strconv.Atoi(queryParams["limit"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad limit type")
			return
		}
		reviewQuery.Limit = limit
	}
	c.Set("reviewQueryParams", &reviewQuery)
}

func (h *Handler) tourQueryParams(c *gin.Context){
	queryParams := c.Request.URL.Query()
	tourQuery := ent.TourQueryParams{
		Offset: 0,
		Limit: 20,
	}
	if queryParams["offset"] != nil{
		offset,err := strconv.Atoi(queryParams["offset"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad offset type")
			return
		}
		tourQuery.Offset = offset
	}
	if queryParams["limit"] != nil{
		limit,err := strconv.Atoi(queryParams["limit"][0])
		if err!=nil{
			newErrorResponse(c, http.StatusUnauthorized, "bad limit type")
			return
		}
		tourQuery.Limit = limit
	}
	c.Set("tourQueryParams", &tourQuery)
}