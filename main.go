package main

import (
    "fmt"
    "html/template"
    "net/http"
    "sync"
)

var (
    count int
    imageShow = false
    mutex sync.Mutex // Para evitar condições de corrida
)

// Handler para a página principal
func homeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.Execute(w, nil)
}

// Handler para a atualização de conteúdo via HTMX
func updateHandler(w http.ResponseWriter, r *http.Request) {
    mutex.Lock() // Bloqueia para evitar condições de corrida
    count++
    updatedContent := fmt.Sprintf("<p>O botão foi clicado %d vezes!</p>", count)
    mutex.Unlock() // Desbloqueia após incrementar

    w.Write([]byte(updatedContent))
}

// Handler para uma segunda rota de teste
func updateHandlerFigure(w http.ResponseWriter, r *http.Request) {
    mutex.Lock() // Bloqueia para evitar condições de corrida
    defer mutex.Unlock()
	imageShow = !imageShow
    var updatedContent string
	if imageShow {
		   updatedContent = `
		       <figure>
		           <img src="https://p15-kimg.kwai.net/kimg/EKzM1y8qmgEKAnMzEg1waG90by1vdmVyc2VhGoQBdXBpYy8yMDIzLzA0LzAyLzIxL0JNakF5TXpBME1ESXlNVEUwTWpaZk1UVXdNREF3TVRnMU1EQTVNemd5WHpFMU1ERXdNalUyT1RRNU9EYzVNMTh5WHpNPV9vZmZuX0I5Njc4NDA1OWE0ZDQ1M2RlYzcwYmU5ZTlhYzVlNTYwZi53ZWJw.webp" alt="Minha Figura">
		           <figcaption>Informações da Figura</figcaption>
		       </figure>`
    } else {
    	updatedContent = ``
    }
    w.Write([]byte(updatedContent))
}

func main() {
    // Rota para a página principal
    http.HandleFunc("/", homeHandler)

    // Rota para a atualização do conteúdo via HTMX
    http.HandleFunc("/atualizar", updateHandler)
    http.HandleFunc("/figure", updateHandlerFigure)

    // Inicia o servidor na porta 8080
    http.ListenAndServe(":8080", nil)
}
