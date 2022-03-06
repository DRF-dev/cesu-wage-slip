package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"sync"
	"time"
)

type Data struct {
 	Object struct {
		Numero string `json:"numero"`
	} `json:"objet"`
}

type Data2 struct {
	ListObj []struct {
		Ref string `json:"referenceDocumentaire"`
	} `json:"listeObjets"`
}

func main() {
	username := flag.String("u", "username", "Username used for login")
	password := flag.String("p", "password", "Password used for login")
	folder := flag.String("f", "folder", "Folder where download all wage slips")
	flag.Parse()

	var wg sync.WaitGroup

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}

	client := http.Client{
		Timeout: time.Second * 40,
		Jar: cookieJar,
	}
	urlLogin := "https://www.cesu.urssaf.fr/info/accueil.login.do"

	data := url.Values{}
	data.Set("username", *username)
	data.Set("password", *password)

	res, err := client.PostForm(urlLogin, data)
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}
 	res.Body.Close()

	res, err = client.Get("https://www.cesu.urssaf.fr/cesuwebdec/status")
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}
	var httpData Data
	d, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}
	json.Unmarshal(d, &httpData)
 	res.Body.Close()

	year, _, _ := time.Now().Date()
	url := fmt.Sprintf("https://www.cesu.urssaf.fr/cesuwebdec/salaries/%v/bulletinsSalaire?pseudoSiret=&dtDebutRecherche=%v0101&dtFinRecherche=%v1231&numStart=0&nbAffiche=100000&numeroOrdre=0&orderBy=orderByRefDoc", httpData.Object.Numero, year - 5, year)
	res, err = client.Get(url)
	if err != nil {
		log.Fatalf("Error %v\n", err)
	}

	var httpData2 Data2
	r, err := io.ReadAll(res.Body)
	json.Unmarshal(r, &httpData2)
	res.Body.Close()

	for _, s0 := range httpData2.ListObj {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			a := fmt.Sprintf("https://www.cesu.urssaf.fr/cesuwebdec/salaries/%v/editions/bulletinSalairePE?refDoc=%v", httpData.Object.Numero, s)

			res, _ = client.Get(a)
			if err != nil {
				log.Fatalf("Error %v\n", err)
			}

			b := fmt.Sprintf("%v/%v.pdf", *folder, s)
			out, _ := os.Create(b)

			io.Copy(out, res.Body)

			res.Body.Close()
			out.Close()
		}(s0.Ref)
	}

	wg.Wait()
}
