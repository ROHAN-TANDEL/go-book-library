package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand/v2"
	"net/http"
	"time"
)

func getBooksViaAPI(c *gin.Context) (map[string]any, error) {

	var data map[string]any

	var url = getBookUrl(c)

	data, err := callHttpsUrl(c, url)

	if err != nil {
		return data, err
	}

	field := data["results"].([]any)

	for key, val := range field {

		bookMap, ok := val.(map[string]any)

		if !ok {
			fmt.Println(key, "is not map[string]any")
			continue
		}

		err := aggregate(c, bookMap)
		if err != nil {
			return data, err
		}
	}

	return data, nil
}

func getBookUrl(c *gin.Context) string {
	return "https://gutendex.com/books/?page=2"
}

func callHttpsUrl(c *gin.Context, url string) (map[string]any, error) {

	var data map[string]any

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Get(url)

	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, err
	}

	er := json.NewDecoder(res.Body).Decode(&data)

	if er != nil {
		return data, er
	}

	return data, nil
}

func aggregate(c *gin.Context, data map[string]any) error {

	authorData, ok := data["authors"].([]any)
	if !ok {
		fmt.Println(" author is not []any")
	}

	var authors = storeAuthor(c, authorData)

	var cats = storeCategory(c, data["subjects"].([]any))

	var err = storeBook(c, data, authors, cats)
	if err != nil {
		return err
	}

	return nil
}

func storeBook(c *gin.Context, data map[string]any, authors []Author, cats []Category) error {
	var bookData newBook
	var language string
	languageList, ok := data["languages"].([]any)
	if !ok {
		fmt.Println(" languages is not []any")
	}
	for _, val := range languageList {
		language = fmt.Sprintf("%s %s", language, val)
	}

	sumArr, ok := data["summaries"].([]any)
	if !ok || len(sumArr) == 0 {
		fmt.Println(" summary is not []any")
	}

	if len(sumArr) > 0 {
		if sumText, isString := sumArr[0].(string); isString {
			data["summary"] = sumText
		} else {
			data["summary"] = "internal response summary"
		}
	} else {
		fmt.Println(" summary is not []any")
		data["summary"] = "internal response summary"
	}

	todayStr := time.Now().UTC().Format("2006-01-02")
	bookData = newBook{
		Title:           strPtr(data["title"].(string)),
		Language:        &language,
		Publisher:       strPtr(data["summary"].(string)[:8]),
		PublicationDate: &todayStr,
		Isbn:            strPtr(generateRandomISBN()),
		Summary:         strPtr(data["summary"].(string)),
		Categories:      cats,
		Authors:         authors,
	}

	_, err := addBookRepo(&bookData)

	if err != nil {
		fmt.Println(" addBookRepo err:", err)
	}
	return nil
}

func strPtr(s string) *string {
	return &s
}

func generateRandomISBN() string {
	minA := int64(1000000000)
	maxA := int64(9999999999)

	// rand.Int64N(N) returns [0, N). We shift the range mathematically:
	randomBody := rand.Int64N(maxA-minA+1) + minA
	return fmt.Sprintf("978%d", randomBody)

}

func storeCategory(c *gin.Context, data []any) []Category {
	if len(data) < 1 {
		return nil
	}
	var cats []Category
	for _, val := range data {
		category := Category{
			Name:        val.(string),
			Description: val.(string),
		}
		cat, err := addCategoryRepo(category)
		if err != nil {
			fmt.Println("category failed to add category")
			fmt.Println(err)
		} else {
			cats = append(cats, cat)
		}
	}
	return cats
}

func storeAuthor(c *gin.Context, data []any) []Author {

	var author Author
	fmt.Println("printing data value")
	fmt.Println(data)
	var authors []Author

	for _, bookAuthor := range data {
		authorData, ok := bookAuthor.(map[string]any)
		if !ok {
			fmt.Println(" author is not map[string]any")
			continue
		}

		var name = authorData["name"].(string)
		birth, ok := authorData["birth_year"].(float64)
		death, ok := authorData["death_year"].(float64)
		var bio = fmt.Sprintf("Born and died in respective years : %v and %v", birth, death)

		author = Author{
			Name:      name,
			Biography: bio,
		}

		res, err := addAuthorRepo(author)
		if err != nil {
			fmt.Println(" addAuthorRepo err:", err)
		} else {
			authors = append(authors, res)
		}
	}

	return authors
}
