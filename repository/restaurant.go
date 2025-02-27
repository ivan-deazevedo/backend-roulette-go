package repository

import (
	"database/sql"
	"log"
	"lunch-be/model"
)

type RestaurantRepository struct {
	Db *sql.DB
}

func NewRestaurantRepository(db *sql.DB) RestaurantRepositoryInterface {
	return &RestaurantRepository{Db: db}
}

func (r *RestaurantRepository) DeleteRestaurant(id uint) bool {
	_, err := r.Db.Exec("DELETE FROM restaurants WHERE restaurant_id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (r *RestaurantRepository) GetOneRestaurant(id uint) model.Restaurant {
	query, err := r.Db.Query("SELECT * FROM restaurants WHERE restaurant_id = $1", id)
	if err != nil {
		log.Println(err)
		return model.Restaurant{}
	}
	var restaurant model.Restaurant
	if query != nil {
		for query.Next() {
			var (
				id     uint
				naam   string
				teller uint
				weburl string
			)
			err := query.Scan(&id, &naam, &teller, &weburl)
			if err != nil {
				log.Println(err)
			}
			restaurant = model.Restaurant{Id: id, Naam: naam, Teller: teller, WebUrl: weburl}
		}
	}
	return restaurant
}

func (r *RestaurantRepository) GetAllRestaurants() []model.Restaurant {
	query, err := r.Db.Query("SELECT * FROM restaurants ORDER BY picked DESC, restaurant_id ")
	if err != nil {
		log.Println(err)
		return nil
	}
	var restaurants []model.Restaurant
	if query != nil {
		for query.Next() {
			var (
				id     uint
				naam   string
				teller uint
				weburl string
			)
			err := query.Scan(&id, &naam, &teller, &weburl)
			if err != nil {
				log.Println(err)
			}
			resto := model.Restaurant{Id: id, Naam: naam, Teller: teller, WebUrl: weburl}
			restaurants = append(restaurants, resto)
		}
	}
	return restaurants
}

func (r *RestaurantRepository) InsertRestaurant(post model.PostRestaurant) bool {
	statement, err := r.Db.Prepare("INSERT INTO restaurants(naam, webUrl) VALUES($1, $2)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer statement.Close()

	_, err2 := statement.Exec(post.Naam, post.WebUrl)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}

func (r *RestaurantRepository) UpdateRestaurant(id uint, post model.UpdateRestaurant) model.Restaurant {
	_, err := r.Db.Exec("UPDATE restaurants SET picked = $1, webUrl = $2 WHERE restaurant_id = $3", post.Teller, post.WebUrl, id)
	if err != nil {
		log.Println(err)
		return model.Restaurant{}
	}
	return r.GetOneRestaurant(id)
}
