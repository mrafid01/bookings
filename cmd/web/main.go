package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrafid01/bookings/internals/config"
	"github.com/mrafid01/bookings/internals/handlers"
	"github.com/mrafid01/bookings/internals/models"
	"github.com/mrafid01/bookings/internals/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("Starting application on port", portNumber)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	gob.Register(models.Reservation{})
	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	return nil
}

// func run() (*driver.DB, error) {
// 	// what am I going to put in the session
// 	gob.Register(models.Reservation{})
// 	gob.Register(models.User{})
// 	gob.Register(models.Room{})
// 	gob.Register(models.Restriction{})
// 	gob.Register(map[string]int{})

// 	// read flags
// 	inProduction := flag.Bool("production", true, "Application is in production")
// 	useCache := flag.Bool("cache", true, "Use template cache")
// 	dbHost := flag.String("dbhost", "localhost", "Database host")
// 	dbName := flag.String("dbname", "", "Database name")
// 	dbUser := flag.String("dbuser", "", "Database user")
// 	dbPass := flag.String("dbpass", "", "Database password")
// 	dbPort := flag.String("dbport", "5432", "Database port")
// 	dbSSL := flag.String("dbssl", "disable", "Database ssl settings (disable, prefer, require)")

// 	flag.Parse()

// 	if *dbName == "" || *dbUser == "" {
// 		fmt.Println("Missing required flags")
// 		os.Exit(1)
// 	}

// 	mailChan := make(chan models.MailData)
// 	app.MailChan = mailChan

// 	// change this to true when in production
// 	app.InProduction = *inProduction
// 	app.UseCache = *useCache

// 	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// 	app.InfoLog = infoLog

// 	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// 	app.ErrorLog = errorLog

// 	session = scs.New()
// 	session.Lifetime = 24 * time.Hour
// 	session.Cookie.Persist = true
// 	session.Cookie.SameSite = http.SameSiteLaxMode
// 	session.Cookie.Secure = app.InProduction

// 	app.Session = session

// 	// connect to database
// 	log.Println("Connecting to database...")
// 	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)
// 	db, err := driver.ConnectSQL(connectionString)
// 	if err != nil {
// 		log.Fatal("Cannot connect to database! Dying...")
// 	}
// 	log.Println("Connected to database!")

// 	tc, err := render.CreateTemplateCache()
// 	if err != nil {
// 		log.Fatal("cannot create template cache")
// 		return nil, err
// 	}

// 	app.TemplateCache = tc

// 	repo := handlers.NewRepo(&app, db)
// 	handlers.NewHandlers(repo)
// 	render.NewRenderer(&app)
// 	helpers.NewHelpers(&app)

// 	return db, nil
// }
