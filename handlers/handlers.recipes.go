package handlers

// func HandleRecipe(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		{
// 			urlPathSegments := strings.Split(r.URL.Path, "/")
// 			accountId, err := strconv.Atoi(urlPathSegments[3])
// 			if err != nil {
// 				log.Print(err)
// 				w.WriteHeader(http.StatusNotFound)
// 				return
// 			}
// 			recipeId, err := strconv.Atoi(urlPathSegments[5])
// 			if err != nil {
// 				log.Print(err)
// 				w.WriteHeader(http.StatusNotFound)
// 				return
// 			}

// 			rs, err := repository.GetRecipe(recipeId, accountId)

// 			if err != nil {
// 				log.Print(err)
// 				w.WriteHeader(http.StatusNotFound)
// 				return
// 			}

// 			j, err := json.Marshal(rs)

// 			if err != nil {
// 				log.Print(err)
// 				w.WriteHeader(http.StatusInternalServerError)
// 				return
// 			}

// 			w.Write(j)
// 		}
// 	case http.MethodPost:
// 		var nr models.Recipe

// 		err := json.NewDecoder(r.Body).Decode(&nr)

// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		b, _ := repository.SaveRecipe(&nr)

// 		if b {
// 			w.Write([]byte(strconv.FormatBool(b)))
// 		}
// 	case http.MethodDelete:
// 		urlPathSegments := strings.Split(r.URL.Path, "/")
// 		accountId, err := strconv.Atoi(urlPathSegments[3])
// 		if err != nil {
// 			log.Print(err)
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}
// 		recipeId, err := strconv.Atoi(urlPathSegments[5])
// 		if err != nil {
// 			log.Print(err)
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}

// 		b, err := repository.DeleteRecipe(recipeId, accountId)

// 		if err != nil {
// 			log.Print(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write([]byte(strconv.FormatBool(b)))
// 	case http.MethodPut:
// 		urlPathSegments := strings.Split(r.URL.Path, "/")
// 		accountId, err := strconv.Atoi(urlPathSegments[3])
// 		if err != nil {
// 			log.Print(err)
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}
// 		recipeId, err := strconv.Atoi(urlPathSegments[5])
// 		if err != nil {
// 			log.Print(err)
// 			w.WriteHeader(http.StatusNotFound)
// 			return
// 		}
// 		var er models.Recipe

// 		err = json.NewDecoder(r.Body).Decode(&er)

// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			return
// 		}

// 		b := repository.UpdateRecipe(&er, recipeId, accountId)

// 		if !b {
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}

// 		w.Write([]byte(strconv.FormatBool(b)))
// 	}
// }
