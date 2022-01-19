package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type MoreInfo struct {
	Location OneLocation
	Date     Date
	Artist   Artists
	Linfo    []LocationInfo
}

//Structure artiste
type Artists struct {
	Id           int
	Name         string
	Members      []string
	CreationDate int
	Image        string
	Locations    string
	FirstAlbum   string
}

type OneLocation struct {
	Id        int
	Locations []string
	Dates     string
}

type LocationInfo struct {
	Name string
	Lat  float64
	Lon  float64
}

//Structure Location
type Location struct {
	Index []OneLocation
}

type Date struct {
	Index2 []struct {
		ID    int
		Dates []string
	}
}

//Page d'accueil
func indexpage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
}

// Page artist
func artistpage(w http.ResponseWriter, r *http.Request) {
	//recupe donné artiste via URL API
	var artist []Artists
	tmpl := template.Must(template.ParseFiles("templates/artist.html"))
	request, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(request.Body)
	request.Body.Close()
	json.Unmarshal(byteValue, &artist)
	tmpl.Execute(w, artist)

}

//Page formulaire artiste
func contactpage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("./templates/contact.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
}

//Page plus détails artistes
func mapage(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	var artist Artists

	//var Locationss []Location
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	_ = json.Unmarshal(b, &artist)

	var Locations OneLocation
	resp2, err := http.Get(artist.Locations)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := ioutil.ReadAll(resp2.Body)
	_ = resp2.Body.Close()
	_ = json.Unmarshal(b2, &Locations)

	linfo := []LocationInfo{}
	for _, v := range Locations.Locations {
		lf := LocationInfo{Name: v, Lat: 42.4668, Lon: -70.9495}
		linfo = append(linfo, lf)
	}

	var MoreInfos MoreInfo = MoreInfo{Location: Locations, Artist: artist, Linfo: linfo}
	tmpl := template.Must(template.ParseFiles("./templates/map.html"))
	_ = tmpl.Execute(writer, MoreInfos)

}

func main() {

	//Url Page

	http.HandleFunc("/", indexpage)
	http.HandleFunc("/artist", artistpage)
	http.HandleFunc("/contact", contactpage)
	http.HandleFunc("/map", mapage)

	//Integration CSS + IMG
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	//Demarage server message
	fmt.Printf("Démarage du serveur Go sur le port 8080 --> Groupie-Traker")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

/* efface / remplace en JavaScript
documentqueryselector : remove || display none // AJAX // route chercher // afficher le résultat JS puis insérer dans la page*/

/*Filtre : Formulaire / De la même façon que la barre de recherche (plus professionnel)*/

/* MVC = modèle Vue Controleur / séparer modèle (schéma de données= ici donnée récup d'une api) Controleur = influer sur ces données Vue = template
Faire un fichier server à la racine qui fait juste le lancement du serveur : register route package enregistrer toutes les routes / package controleur qui va contenir : Search/filtres...
(vidéo module à regarder) */
