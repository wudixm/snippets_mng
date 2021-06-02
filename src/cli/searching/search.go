package searching

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gin-gonic/gin"
	"log"
	. "main/src/models"
	"strconv"
	"strings"
	"time"
)

// M is an alias for map[string]interface{}
type M map[string]interface{}

func searchByKeyword(kw string) []interface{} {
	var (
		r M
	)
	log.Printf("kw is %s", kw)
	log.Printf("kw is \"\" %s", kw == "")
	var buf bytes.Buffer
	query := M{
		"query": M{
			"multi_match": M{
				"query":  kw,
				"fields": []string{"content", "title"},
			},
		},
	}
	if kw == "" {

		sorted_list := [1]M{{"_id": M{"order": "desc"}}}
		query = M{
			"query": M{
				"match_all": M{
				},
			},
			"sort": sorted_list,
			"size": 100,
		}
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}
	log.Printf(buf.String())
	explain := true
	//analyze_wildcard := true

	req := esapi.SearchRequest{
		Index: []string{"test"},
		//Query: "content:"+kw,
		Body:    &buf,
		Explain: &explain,
		//Analyzer: "simple", 用Query 类型的请求的时候可以添加analyzer, 不能是body类型的请求
		//AnalyzeWildcard: &analyze_wildcard, 同上
		Human:  true,
		Pretty: true,
	}
	res, err := req.Do(context.Background(), ES)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf("%s", res.Status())
	//log.Printf("%s", res.Body)
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	result := []interface{}{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		_id := hit.(map[string]interface{})["_id"]
		_source := hit.(map[string]interface{})["_source"]

		log.Printf(" * ID=%s, %s", _id, _source)
		m := make(map[string]interface{})
		m["_id"] = _id
		m["_source"] = _source

		result = append(result, m)
	}

	log.Println(strings.Repeat("=", 37))
	return result
}
func searchByID(_id string) []interface{} {
	var (
		r map[string]interface{}
	)
	req := esapi.SearchRequest{
		Index: []string{"test"},
		Query: "_id:" + _id,
	}
	res, err := req.Do(context.Background(), ES)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	log.Printf("%s", res.Status())
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.
	result := []interface{}{}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		_id := hit.(map[string]interface{})["_id"]
		_source := hit.(map[string]interface{})["_source"]
		m := make(map[string]interface{})
		m["_id"] = _id
		m["_source"] = _source

		result = append(result, m)
	}

	log.Println(strings.Repeat("=", 37))
	return result
}

func SearchByKeyWord(c *gin.Context) []interface{} {
	kw := c.DefaultQuery("keyword", "")
	this_id := c.DefaultQuery("this_id", "")
	if this_id == "" {
		return searchByKeyword(kw)
	} else {
		return searchByID(this_id)
	}

}

func IndexDocument(c *gin.Context) {
	kw := c.DefaultPostForm("content", "")
	code := strconv.QuoteToASCII(kw)
	//title := c.DefaultQuery("title", "")
	tag := c.DefaultPostForm("tag", "")
	lang := c.DefaultPostForm("lang", "")
	title := c.DefaultPostForm("title", "")
	fmt.Printf("params============begin")
	fmt.Printf("%s", kw)
	fmt.Printf("%s", code)
	fmt.Printf("%s", tag)
	fmt.Printf("%s", lang)
	fmt.Printf("%s", title)
	fmt.Printf("params============end")

	//var buf bytes.Buffer
	//cond_length := len(kw)
	//for i := 0;i < cond_length;i++ {
	//	buf.Write(kw[i])
	//}
	//kw_byte := []byte(kw)

	now := time.Now()
	var b strings.Builder
	fmt.Printf("builder============begin")
	b.WriteString(`{"content":`)
	b.WriteString(code)
	b.WriteString(`,`)
	b.WriteString(`"title":"`)
	b.WriteString(title)
	b.WriteString(`",`)
	b.WriteString(`"tag":"`)
	b.WriteString(tag)
	b.WriteString(`",`)
	b.WriteString(`"language":"`)
	b.WriteString(lang)
	b.WriteString(`"}`)
	// Set up the request object.
	print(b.String())
	fmt.Printf("builder============end")
	req := esapi.IndexRequest{
		Index:      "test",
		DocumentID: strconv.Itoa(int(now.Unix())),
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}
	res, err := req.Do(context.Background(), ES)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%d", res.Status(), now.Minute())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}

func DeleteIndex(c *gin.Context) {
	kw := c.DefaultQuery("id", "")

	// Set up the request object.
	req := esapi.DeleteRequest{
		Index:      "test",
		DocumentID: kw,
	}
	res, err := req.Do(context.Background(), ES)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	now := time.Now()
	if res.IsError() {
		log.Printf("[%s] Error delete document ID=%d", res.Status(), now.Minute())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}

func TestIndex() {

	res, err := ES.Index(
		"my-index",
		strings.NewReader(`{"title":"Test"}`),
		ES.Index.WithDocumentID("1"))
	fmt.Println(res, err)
}
