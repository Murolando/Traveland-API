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

func (r ReviewBD) DeleteReview(id int)(bool,error){
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id = $1`,reviewTable)
	res1,err := r.db.Exec(query,id)
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
func (r ReviewBD) GetAllReview(placeId int,guideId int, offset int)([]ent.Review,error){
	reviews := make([]ent.Review,0)
	if placeId != 0{
		query := fmt.Sprintf(`SELECT id,user_id,rating,review_text,review_datetime FROM "%s" WHERE place_id = $1 LIMIT $2 OFFSET $3`, reviewTable)
		rows,err := r.db.Query(query,placeId,limit,offset)
		if err!=nil{
			return nil,err
		}
		for rows.Next(){
			var review ent.Review
	
			if err := rows.Scan(&review.ReviewId,&review.UserId,&review.Rating,&review.ReviewText,&review.ReviewTime);err!=nil{
				return nil, err
			}
			reviews = append(reviews, review)
		}
	}else{
		query := fmt.Sprintf(`SELECT id,user_id,rating,review_text,review_datetime FROM "%s" WHERE guide_id = $1 LIMIT $2 OFFSET $3`, reviewTable)		
		rows,err := r.db.Query(query,guideId,limit,offset)
		if err!=nil{
			return nil,err
		}
		for rows.Next(){
			var review ent.Review
	
			if err := rows.Scan(&review.ReviewId,&review.UserId,&review.Rating,&review.ReviewText,&review.ReviewTime);err!=nil{
				return nil, err
			}
			reviews = append(reviews, review)
		}
	}
	
	return reviews,nil
}

