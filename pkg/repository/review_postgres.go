package repository

import (
	"fmt"
	"time"
	"traveland/ent"

	"github.com/jmoiron/sqlx"
)

type ReviewBD struct {
	db *sqlx.DB
}

func NewReviewBD(db *sqlx.DB) *ReviewBD {
	return &ReviewBD{
		db: db,
	}
}

func (r ReviewBD) AddReview(review ent.Review) (int, error) {
	var id int
	t := time.Now()
	review.ReviewTime = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	if review.PlaceId != 0 {
		query := fmt.Sprintf(`INSERT INTO "%s" (user_id, rating, review_text, review_datetime, place_id) values ($1, $2, $3, $4, $5) RETURNING id`,reviewTable)
		row := r.db.QueryRow(query,review.UserId,review.Rating,review.ReviewText,review.ReviewTime,review.PlaceId)
		if err := 	row.Scan(&id);err!=nil{
			return 0,err
		}
	} else{
		query := fmt.Sprintf(`INSERT INTO "%s" (user_id, rating, review_text, review_datetime, guide_id) values ($1, $2, $3, $4, $5) RETURNING id`,reviewTable)
		row := r.db.QueryRow(query,review.UserId,review.Rating,review.ReviewText,review.ReviewTime,review.GuideId)
		if err := 	row.Scan(&id);err!=nil{
			return 0,err
		}
	}
	return id, nil
}

func (r ReviewBD) DeleteReview(id int,userId int)(bool,error){
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1 AND user_id = $2`,reviewTable)
	res1,err := r.db.Exec(query,id,userId)
	if err!=nil{
		return false,err
	}
    count, err := res1.RowsAffected()
    if err != nil {
        return false,err
    }
	if count != 0{
		return true,nil
	}
	return false,nil
}

func (r ReviewBD) GetAllReview(params *ent.ReviewQueryParams)(*ent.AllReviewResponce,error){
	placeId := params.PlaceId
	guideId := params.GuideId
	limit 	:= params.Limit
	offset  := params.Offset
	var reviews ent.AllReviewResponce
	reviews.Reviews = make([]ent.ReviewResponce,0)
	if placeId != 0{
		queryAvg := fmt.Sprintf(`
		SELECT 
			AVG(rating)
			FROM "%s" 
			WHERE place_id = $1
		`,reviewTable)
		row := r.db.QueryRow(queryAvg,placeId)
		if err := row.Scan(&reviews.MeanRating);err!=nil{
			return nil,err
		}

		query := fmt.Sprintf(`
		SELECT id,
		(SELECT name 
			FROM "user" 
			WHERE id = review.user_id) AS name,
		rating,review_text,review_datetime 
		FROM "%s" 
		WHERE place_id = $1 
		LIMIT $2 
		OFFSET $3`, reviewTable)
		rows,err := r.db.Query(query,placeId,limit,offset)
		if err!=nil{
			return nil,err
		}
		for rows.Next(){
			var review ent.ReviewResponce
	
			if err := rows.Scan(&review.ReviewId,&review.UserName,&review.Rating,&review.ReviewText,&review.ReviewTime);err!=nil{
				return nil, err
			}
			reviews.Reviews = append(reviews.Reviews, review)
		}
	}else{
		queryAvg := fmt.Sprintf(`
		SELECT 
			AVG(rating)
			FROM "%s" 
			WHERE guide_id = $1
		`,reviewTable)
		row := r.db.QueryRow(queryAvg,guideId)
		if err := row.Scan(&reviews.MeanRating);err!=nil{
			return nil,err
		}

		query := fmt.Sprintf(`SELECT id,
		(SELECT name 
			FROM "user" 
			WHERE id = review.user_id) AS name,
		,rating,review_text,review_datetime 
		FROM "%s" 
		WHERE guide_id = $1 
		LIMIT $2 
		OFFSET $3`, reviewTable)		
		rows,err := r.db.Query(query,guideId,limit,offset)
		if err!=nil{
			return nil,err
		}
		for rows.Next(){
			var review ent.ReviewResponce
	
			if err := rows.Scan(&review.ReviewId,&review.UserName,&review.Rating,&review.ReviewText,&review.ReviewTime);err!=nil{
				return nil, err
			}
			reviews.Reviews = append(reviews.Reviews, review)
		}
	}
	
	return &reviews,nil
}

func (r ReviewBD) UpdateReview(reviewId int,rating int, reviewText string) (bool,error){
	t := time.Now()
	time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	query := fmt.Sprintf(`UPDATE "%s" SET rating = $1, review_text = $2, review_datetime = $3 WHERE id = $4`,reviewTable)
	_,err := r.db.Exec(query,rating,reviewText,time,reviewId)
	if err!=nil{
		return false,err
	}
	return true,nil
}