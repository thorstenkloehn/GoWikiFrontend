# Ahrensburg Wiki Viewer

Dieses Programm ist ein einfacher Webserver, der Inhalte von der Ahrensburg Wiki API abruft und auf einer Webseite anzeigt. Es verwendet Go, um die API-Abfragen durchzuführen und die HTML-Inhalte zu bereinigen, bevor sie dem Benutzer präsentiert werden.

## Funktionen

- **Startseite**: Leitet den Benutzer zur Hauptseite des Wikis weiter.
- **Seitenanzeige**: Ruft den Inhalt einer spezifischen Wiki-Seite ab und zeigt ihn an.
- **Bereinigung**: Entfernt spezifische HTML-Elemente (z.B. `<div id="map_leaflet_1">`) aus dem abgerufenen Inhalt.

## Installation

1. **Go installieren**: Stellen Sie sicher, dass Go auf Ihrem System installiert ist. Sie können es von [golang.org](https://golang.org/dl/) herunterladen.
2. **Abhängigkeiten installieren**: Führen Sie den folgenden Befehl aus, um die benötigten Go-Pakete zu installieren:
    ```sh
    go get -u github.com/gorilla/mux github.com/PuerkitoBio/goquery
    ```
3. **Konfigurationsdatei erstellen**: Erstellen Sie eine `config.json`-Datei im selben Verzeichnis wie `main.go` mit folgendem Inhalt:
    ```json
    {
        "APIUrl": "https://wiki.ahrensburg.city/api.php?action=parse&format=json&page="
    }
    ```

## Verwendung

1. **Server starten**: Führen Sie den folgenden Befehl aus, um den Webserver zu starten:
    ```sh
    go run main.go
    ```
2. **Webseite aufrufen**: Öffnen Sie Ihren Webbrowser und gehen Sie zu `http://localhost:8080`. Sie werden zur Hauptseite des Wikis weitergeleitet.
3. **Spezifische Seite anzeigen**: Um eine spezifische Wiki-Seite anzuzeigen, geben Sie die URL im Format `http://localhost:8080/{pageName}` ein, wobei `{pageName}` der Name der gewünschten Wiki-Seite ist.

## Beispiel

Um die Seite "Hauptseite" anzuzeigen, öffnen Sie `http://localhost:8080/Hauptseite` in Ihrem Webbrowser.

## Lizenz

Dieses Projekt ist unter der MIT-Lizenz lizenziert. Weitere Informationen finden Sie in der `LICENSE`-Datei.