package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var cnts []*Content
	q := datastore.NewQuery("Content").Order("-createdAt").Limit(10)
	if _, err := q.GetAll(ctx, &cnts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
