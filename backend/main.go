package main
import "strings"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)


type Apartment struct {
	ID          int
	Title       string
	ImageLink   string
	Description string

	Price        int
	Favourite    bool
}


var apartments = []Apartment{
	{ID: 1, Title: "Aghanim's Shard",  ImageLink: "https://liquipedia.net/commons/images/7/79/Aghanim%27s_Shard_itemicon_dota2_gameasset.png", Description: "With origins known only to a single wizard, fragments of this impossible crystal are nearly as coveted as the renowned scepter itself..", Price: 1400, Favourite: false},
    {ID: 2, Title: "Blink Dagger",  ImageLink: "https://liquipedia.net/commons/images/d/dc/Blink_Dagger_itemicon_dota2_gameasset.png", Description: "The fabled dagger used by the fastest assassin ever to walk the lands.", Price: 2250, Favourite: false},
    {ID: 3, Title: "Demon Edge",  ImageLink: "https://liquipedia.net/commons/images/4/44/Demon_Edge_itemicon_dota2_gameasset.png", Description: "One of the oldest weapons forged by the Demon-Smith Abzidian, it killed its maker when he tested its edge.", Price: 2200, Favourite: false},
    {ID: 4, Title: "Hand of Midas",  ImageLink: "https://liquipedia.net/commons/images/6/69/Hand_of_Midas_itemicon_dota2_gameasset.png", Description: "Preserved through unknown magical means, the Hand of Midas is a weapon of greed, sacrificing animals to line the owner's pockets.", Price: 2200, Favourite: false},
    {ID: 5, Title: "Boots of Bearing",  ImageLink: "https://liquipedia.net/commons/images/2/2f/Boots_of_Bearing_itemicon_dota2_gameasset.png", Description: "Resplendent footwear fashioned for the ancient herald that first dared spread the glory of Stonehall beyond the original borders of its nascent claim.", Price: 4275, Favourite: false},
    {ID: 6, Title: "Aghanim's Blessing",  ImageLink: "https://liquipedia.net/commons/images/1/1b/Aghanim%27s_Scepter_itemicon_dota2_gameasset.png", Description: "The scepter of a wizard with demigod-like powers.",  Price: 5800, Favourite: false},

}


func getApartmentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(apartments)
}


func createApartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newApartment Apartment
	err := json.NewDecoder(r.Body).Decode(&newApartment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newApartment.ID = len(apartments) + 1
	apartments = append(apartments, newApartment)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newApartment)
}


func getApartmentByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/apartments/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
		return
	}

	for _, apartment := range apartments {
		if apartment.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(apartment)
			return
		}
	}

	http.Error(w, "Apartment not found", http.StatusNotFound)
}


func deleteApartmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Path[len("/apartments/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
		return
	}

	for i, apartment := range apartments {
		if apartment.ID == id {
			apartments = append(apartments[:i], apartments[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Apartment not found", http.StatusNotFound)
}
func updateApartmentHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    fmt.Println("Received PUT request for updating apartment")


    idStr := strings.TrimPrefix(r.URL.Path, "/apartments/update/")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
        return
    }

    var updatedFields struct {
        Title       string  `json:"title"`
        Description string  `json:"description"`
        ImageLink   string  `json:"image_link"`
        Price       float64 `json:"price"`
    }

    err = json.NewDecoder(r.Body).Decode(&updatedFields)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }


    for i, apartment := range apartments {
        if apartment.ID == id {
            fmt.Println("Apartment found, updating fields...")
            if updatedFields.Title != "" {
                apartments[i].Title = updatedFields.Title
            }
            if updatedFields.Description != "" {
                apartments[i].Description = updatedFields.Description
            }
            if updatedFields.ImageLink != "" {
                apartments[i].ImageLink = updatedFields.ImageLink
            }
            if updatedFields.Price > 0 {
                apartments[i].Price = int(updatedFields.Price)
            }

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(apartments[i])
            fmt.Println("Apartment updated successfully")
            return
        }
    }

    fmt.Println("Apartment not found")
    http.Error(w, "Apartment not found", http.StatusNotFound)
}



func toggleFavouriteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}


	idStr := strings.TrimPrefix(r.URL.Path, "/apartments/favourite/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Apartment ID", http.StatusBadRequest)
		return
	}


	for i, apartment := range apartments {
		if apartment.ID == id {

			apartments[i].Favourite = !apartments[i].Favourite


			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(apartments[i])
			return
		}
	}

	http.Error(w, "Apartment not found", http.StatusNotFound)
}


func main() {
	http.HandleFunc("/apartments", getApartmentsHandler)             // Получить все квартиры
	http.HandleFunc("/apartments/create", createApartmentHandler)    // Создать квартиру
	http.HandleFunc("/apartments/", getApartmentByIDHandler)         // Получить квартиру по ID
	http.HandleFunc("/apartments/update/", updateApartmentHandler)   // Обновить квартиру
	http.HandleFunc("/apartments/delete/", deleteApartmentHandler)   // Удалить квартиру
    http.HandleFunc("/apartments/favourite/", toggleFavouriteHandler) // Изменить Favourite
	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}