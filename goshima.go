package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"goshima/database"
	"goshima/mapping"

	"github.com/joho/godotenv"
)

const (
    SHORT = "short"
    LONG = "long"
)

var BASE_URL string = "https://ddh.in/"

func main() {
    // Loading .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Generating DSN string for GORM initialization
    configString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"
    config := fmt.Sprintf(configString, os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"), os.Getenv("PORT"), os.Getenv("SSLMODE"))

    // Getting command-line argument
    mode, url := getArgs()

    // Creating a map containing mapping for a bijective function
    bijectiveMap := make(map[int32]string, 62)
    mapping.EncodeMap(bijectiveMap, mapping.Charset{Start: 97, Iterations: 25})
    mapping.EncodeMap(bijectiveMap, mapping.Charset{Start: 65, Iterations: 25})
    mapping.EncodeMap(bijectiveMap, mapping.Charset{Start: 48, Iterations: 9})

    // Initializing the database 
    db, err := database.InitializeGorm(config)
    if err != nil {
        panic("Failed to connect to database")
    }

    db.AutoMigrate(&database.UrlStore{})

    var urlStore database.UrlStore
    
    switch mode {
    case SHORT:
        db.Create(&database.UrlStore{OriginalUrl: url})

        db.Find(&urlStore, "original_url = ?", url)

        shortenedUrl := BASE_URL + mapping.GetShortPath(urlStore.ID, bijectiveMap)
        fmt.Printf("Shortened URL: %s\n", shortenedUrl)

    case LONG:
        shortPath := strings.SplitAfter(url, BASE_URL)[1]
        primaryKey := mapping.GetUrlId(shortPath, bijectiveMap)

        db.Find(&urlStore, "id = ?", primaryKey)
        originalUrl := urlStore.OriginalUrl
        fmt.Printf("Original URL: %s\n", originalUrl)
    default:
        fmt.Println("Invalid choice")
    }
}

// Gets the first command line argument
func getArgs() (string, string) {
    var url string
    var mode string
    args := os.Args[1:]

    if len(args) > 0 {
        mode = args[0]
        url = args[1]
        
        fmt.Printf("Entered URL is %s\nMode: %s\n", url, mode)
        
        if url != "" {
            return strings.TrimSpace(mode), strings.TrimSpace(url)
        }
    }

    return mode, url
}

