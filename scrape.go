package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/urlfetch"
)

func scrape(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client := urlfetch.Client(ctx)

	resp, err := client.Get("https://www.google.com/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	# 適当な値を入れている
	title := string(byteArray[0])
	date := string(byteArray[1])
	text := string(byteArray[2])

	cnt := &Content{
		Title:     title,
		Date:      date,
		Text:      text,
		CreatedAt: time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, "Content", nil)
	if _, err := datastore.Put(ctx, key, cnt); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
