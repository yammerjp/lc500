package main

// sample api server

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
)

type Tenant struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Hostname string    `json:"hostname"`
	Profile  Profile   `json:"profile"`
	Articles []Article `json:"articles"`
}

type Profile struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var tenants []Tenant

func main() {
	loadDatabase()

	r := http.NewServeMux()
	r.HandleFunc("/api/v1/profile", getProfile)
	r.HandleFunc("/api/v1/articles", getArticles)
	r.HandleFunc("/api/v1/articles/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/api/v1/articles/"):]
		getArticle(w, r, id)
	})

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	if port == 0 {
		port = 8082
	}

	fmt.Printf("Listening on port %d\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		panic(err)
	}
}

func loadDatabase() {
	path := os.Getenv("DATABASE_PATH")
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var dbData struct {
		Tenants map[string]struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			Articles    map[string]struct {
				Title string `json:"title"`
				Body  string `json:"body"`
			} `json:"articles"`
		} `json:"tenants"`
	} = struct {
		Tenants map[string]struct {
			Name        string "json:\"name\""
			Description string "json:\"description\""
			Articles    map[string]struct {
				Title string "json:\"title\""
				Body  string "json:\"body\""
			} "json:\"articles\""
		} "json:\"tenants\""
	}{}

	err = json.Unmarshal(data, &dbData)
	if err != nil {
		panic(err)
	}

	// Convert dbData to tenants slice
	tenants = make([]Tenant, 0, len(dbData.Tenants))
	for hostname, tenantData := range dbData.Tenants {
		tenant := Tenant{
			Hostname: hostname,
			Profile: Profile{
				Name:        tenantData.Name,
				Description: tenantData.Description,
			},
			Articles: make([]Article, 0, len(tenantData.Articles)),
		}

		for id, article := range tenantData.Articles {
			tenant.Articles = append(tenant.Articles, Article{
				ID:      id,
				Title:   article.Title,
				Content: article.Body,
			})
		}

		tenants = append(tenants, tenant)
	}
}

func getTenantFromHostname(hostname string) *Tenant {
	for _, tenant := range tenants {
		if tenant.Hostname == hostname {
			return &tenant
		}
	}
	return nil
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	hostname, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		panic(err)
	}
	tenant := getTenantFromHostname(hostname)
	if tenant == nil {
		http.Error(w, "Tenant not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tenant.Profile)
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	hostname, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		panic(err)
	}
	tenant := getTenantFromHostname(hostname)

	if tenant == nil {
		http.Error(w, "Tenant not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tenant.Articles)
}

func getArticle(w http.ResponseWriter, r *http.Request, id string) {
	hostname, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		panic(err)
	}
	tenant := getTenantFromHostname(hostname)

	if tenant == nil {
		http.Error(w, "Tenant not found", http.StatusNotFound)
		return
	}

	for _, article := range tenant.Articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	http.Error(w, "Article not found", http.StatusNotFound)
}
