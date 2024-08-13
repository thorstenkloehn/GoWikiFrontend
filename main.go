package main

import (
    "encoding/json"
    "html/template"
    "io/ioutil"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/PuerkitoBio/goquery"
    "strings"
    "os"
)

type Config struct {
    APIUrl string `json:"APIUrl"`
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)
    r.HandleFunc("/{pageName}", PageHandler)
    http.Handle("/", r)
    http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/Hauptseite", http.StatusFound)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    pageName := vars["pageName"]

    // Konfigurationsdatei laden
    configFile, err := os.Open("config.json")
    if err != nil {
        http.Error(w, "Fehler beim Laden der Konfigurationsdatei: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer configFile.Close()

    var config Config
    err = json.NewDecoder(configFile).Decode(&config)
    if err != nil {
        http.Error(w, "Fehler beim Parsen der Konfigurationsdatei: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // URL der API-Abfrage
    url := config.APIUrl + pageName + "&disableeditsection=1&disabletoc=1"

    // API-Abfrage durchführen
    resp, err := http.Get(url)
    if err != nil {
        http.Error(w, "Fehler bei der API-Abfrage: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Fehler beim Lesen der Antwort: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // JSON-Antwort parsen
    var apiData map[string]interface{}
    err = json.Unmarshal(body, &apiData)
    if err != nil {
        http.Error(w, "Fehler beim Parsen der JSON-Antwort: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // HTML-Inhalt extrahieren
    parse, ok := apiData["parse"].(map[string]interface{})
    if !ok {
        http.Error(w, "Fehler beim Extrahieren der Parse-Daten", http.StatusInternalServerError)
        return
    }

    text, ok := parse["text"].(map[string]interface{})
    if !ok {
        http.Error(w, "Fehler beim Extrahieren des Textes", http.StatusInternalServerError)
        return
    }

    html, ok := text["*"].(string)
    if !ok {
        http.Error(w, "Fehler beim Extrahieren des HTML-Inhalts", http.StatusInternalServerError)
        return
    }

    // HTML-Inhalt mit goquery parsen
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
    if err != nil {
        http.Error(w, "Fehler beim Parsen des HTML-Inhalts: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Entfernen des <div id="map_leaflet_1">-Inhalts
    doc.Find("#map_leaflet_1").Remove()

    // Bereinigten HTML-Inhalt zurückgeben
    cleanedHTML, err := doc.Html()
    if err != nil {
        http.Error(w, "Fehler beim Generieren des bereinigten HTML-Inhalts: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // HTML-Vorlage parsen
    tmpl, err := template.ParseFiles("index.html")
    if err != nil {
        http.Error(w, "Fehler beim Laden der Vorlage: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Daten für die Vorlage
    data := struct {
        Title   string
        Content template.HTML
    }{
        Title:   pageName,
        Content: template.HTML(cleanedHTML),
    }

    // HTML-Inhalt mit der Vorlage ausgeben
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, "Fehler beim Ausführen der Vorlage: "+err.Error(), http.StatusInternalServerError)
    }
}