package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// func main() {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("69"))
// 	})

// 	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello Lala Cantik!"))
// 	})

// 	dat, err := os.ReadFile("kocak.html")

// 	if err != nil {
// 		panic(err)
// 	}

// 	r.Get("/nano", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write(dat)
// 	})
// 	bocil, e := os.ReadFile("bocil.html")
// 	if e != nil {
// 		panic(e)
// 	}
// 	r.Get("/bocil",func(w http.ResponseWriter, r *http.Request) {
// 		w.Write(bocil)
// 	})

// 	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("{lala: cantik}"))
// 	})
// 	http.ListenAndServe(":3001", r)
// }
 
/*func main(): This defines the main part of the program. When the program runs, it starts here.
  r := chi.NewRouter(): We create a new router (called r) using chi.NewRouter(). A router is like a traffic controller for the internet, deciding what to do when someone visits a specific address.
  r.Use(middleware.Logger): This tells the router to use the Logger, which records all the visits and actions on the website (like a security camera).*/

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

/*r.Use(cors.Handler(...)): This tells the router to use the CORS handler. CORS lets the server allow or block requests from different origins (websites).
  AllowedOrigins: This says which websites are allowed to ask for information from this server. In this case, http://localhost:3001 is allowed (your Next.js site).
  AllowedMethods: These are the actions the server allows, like:
   - GET: Asking for information.
   - POST: Sending new information.
   - PUT: Updating information.
   - DELETE: Removing information.
   - OPTIONS: Checking what actions are allowed.
  AllowedHeaders: These are the kinds of information (headers) the server accepts in a request, like:
  Accept: What kind of data the client can understand.
  Authorization: Checking if the client is allowed to access data.
  Content-Type: What kind of data is being sent (e.g., JSON).
  X-CSRF-Token: Extra security token to prevent bad actions.
  ExposedHeaders: These are headers the server is allowed to send back.
  AllowCredentials: This allows the server to accept credentials (like login information) from the client.
  MaxAge: This sets how long (in seconds) the browser can remember these CORS rules.*/

	// Enable CORS
    r.Use(cors.Handler(cors.Options{
        AllowedOrigins: []string{"http://localhost:3001"}, // Allow requests from Next.js
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders: []string{"Link"},
        AllowCredentials: true,
        MaxAge: 300,
    }))

	/*r.Get("/"): This sets up a route (or path) for the website. The / means it’s the homepage. So when someone visits http://localhost:3000/, the server responds.
      handleFile("book.json"): This calls a function that reads and sends the contents of the book.json file when someone visits the homepage.*/
	r.Get("/", handleFile("book.json"))

	//This starts the web server on port 3000. The r is the router, which controls what happens when someone visits the website.
	http.ListenAndServe(":3000", r)
}

/*func handleFile(book string) func(w http.ResponseWriter, r *http.Request): This creates a function called handleFile. This function takes the file name (book) as input and returns another function that responds to requests.
  w http.ResponseWriter: This is how we send the response back to the visitor (like writing a reply letter).
  *r http.Request: This is the visitor’s request (like a letter asking for information).*/

// Function untuk baca file dan mengembalikan return sebuah function "Handler" sebagai
// callback untuk router
func handleFile(book string) func(w http.ResponseWriter, r *http.Request) {
	// Ambil nama file
	data, e := os.ReadFile("book.json")

	if e != nil {
		panic(e)
	}

	/*return func(w http.ResponseWriter, r *http.Request): This returns a new function. This function will send the book data to the visitor when they ask for it.
     w.Header().Set("Content-Type", "application/json;charset=utf-8"): This sets the header of the response. The header tells the browser that we are sending JSON data and that it uses UTF-8 characters (a common way of encoding text).
     w.Write(data): This writes (or sends) the data (which contains the content of book.json) to the visitor.*/
	 
	// Melakukan return sebuah function
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.Write(data)
	}
}
