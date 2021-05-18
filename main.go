package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"encoding/json"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

type product struct {
	Quantity int
	Title    string
	Category string
	Cost     int
}

//cited
//https://www.bing.com/videos/search?q=youtbe+golang+template&refig=e742578f4d004a2b8a5bd1f28849eb0f&ru=%2fsearch%3fq%3dyoutbe%2bgolang%2btemplate%26form%3dANNTH1%26refig%3de742578f4d004a2b8a5bd1f28849eb0f&view=detail&mmscn=vwrc&mid=BD040005A2743ACB801ABD040005A2743ACB801A&FORM=WRVORC
//http://localhost:8080/golangproj/

var globt *template.Template
var globKeyword = ""
var Test = 1

var ProductID = 0

type Rectangle struct {
	Length  int
	breadth int
	color   string
}

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

var string1 = ""

type App struct {
	Name string
}

var tpl *template.Template

type employee struct {
	gKeyword1           string
	gKeyword2           string
	gKeyword3           string
	ProductName         string
	ProductID           int
	ProductdDescription string
	ProductCost         int
	ProductQuantity     int
	ProductCatTitle     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "ecommerce"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//////////

//executes back to :  finalpage.html

//var list[]product

//var List1  []product
var prod []product

func addElement(var1 int, var2 string, var3 string, var4 int) {

	var element product
	element.Quantity = var1
	element.Title = var2
	element.Category = var3
	element.Cost = var4

	prod = append(prod, element)

}

//pass an array to here from index.html
func purgeHTML(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	///////////////

	////////////////
	//incoming : productid and quantity
	db := dbConn()
	//nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT products.ProductQuantity, products.ProductName, products.ProductCatTitle, products.ProductCost FROM products WHERE products.ProductCost=100")
	if err != nil {
		panic(err.Error())
	}

	//rows, err := db.Query()
	//if err != nil {
	//	panic(err.Error())
	//}

	for selDB.Next() {
		var quant, cost int
		var title, category string
		err = selDB.Scan(&quant, &title, &category, &cost)
		if err != nil {
			panic(err.Error())
		}

		//prod.Quantity = quant
		//prod.Title = title
		//prod.Category = category
		//prod.Cost = cost
		addElement(quant, title, category, cost)
	}
	//tmpl.ExecuteTemplate(w, "Show", emp)

	//templ1 = product{ProductID, ProductCatTitle, titleID, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity,
	//	key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword, ProductFilename, AmountToPurchaseID, AmountPurchasedID, mainDivID}

	//fmt.Println(templ1)

	globt := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template2.html"))

	globt.Execute(w, prod)

	defer db.Close()
}

//////////

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rrrrrrraarg ")
}

////////

func receiveAjax(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		//data := r.FormValue("post_data")
		r.FormValue("post_data")
		fmt.Println("Receive ajax post data string ")

		w.Header().Add("Content-Type", "application/html")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Write([]byte(string1))

	}

}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//this is for testin, not used anympre
func processSearch(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rrrrrrraarg ")
	fmt.Fprintf(w, "got here1!")

}

type forTemplate struct {
	ProductID       int
	ProductCatTitle string
	//MainDiv         string
	TitleID string
	//ProductFilename    string
	ProductName        string
	DescID             string
	ProductDescription string
	CostID             string
	ProductCost        int
	QuantityID         string
	ProductQuantity    int
	Key1ID             string
	GKeyword1          string
	Key2ID             string
	GKeyword2          string
	Key3ID             string
	GKeyword3          string
	ProductFilename    string
	AmountToPurchaseID string
	AmountPurchasedID  string
	MainDivID          string
}

type Name struct {
	FName string
	LName string
}

type VAR1 struct {
	Var1 string
}

//var templ1 = forTemplate{str3, var18, var2, var3, var4, var5, var6, var7, str4, var9, str2, var11, var12, var13, var14, var15, var16}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var Var1 = "abc"
	var var2 = VAR1{Var1}
	//w.Header().Add("Content-Type", "application/html")

	//w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//data := VAR{"testname"}
	//var var1 = VAR{var2}

	globt := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/twemplate1.html"))

	err1 := globt.Execute(w, var2)

	if err1 != nil {
		fmt.Println("B---------------")
		fmt.Println(err1.Error())

		panic(err1.Error())

	}

	//t, _ := template.ParseFiles("index1.html")
	//var t = template.Must(template.New("").Parse("index1.html"))
	//globt.Execute(w, "a")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	name := Name{"mindorks2", "Subject2"}
	template, _ := template.ParseFiles("index2.html")
	template.Execute(w, name)
}

func updateForm(w http.ResponseWriter, r *http.Request) {
	/*
		var var1 = r.FormValue("productID")
		var var18 = r.FormValue("ProductCatTitle")
		var var2 = r.FormValue("mainDiv")
		var var3 = r.FormValue("titleID")
		var var4 = r.FormValue("ProductName")
		var var5 = r.FormValue("descID")
		var var6 = r.FormValue("ProductDescription")
		var var7 = r.FormValue("costID")
		var var8 = r.FormValue("ProductCost")
		var var9 = r.FormValue("quantityID")
		var var10 = r.FormValue("ProductQuantity")
		var var11 = r.FormValue("key1ID")
		var var12 = r.FormValue("globKeyword")
		var var13 = r.FormValue("key2ID")
		var var14 = r.FormValue("globKeyword")
		var var15 = r.FormValue("key3ID")
		var var16 = r.FormValue("globKeyword")
		var var17 = r.FormValue("amountPurchased")
		var var19 = r.FormValue("ProductFilename")

		//enough product

		str, _ := strconv.Atoi(var17)
		str2, _ := strconv.Atoi(var10)
		str3, _ := strconv.Atoi(var1)
		str4, _ := strconv.Atoi(var8)

		if str <= str2 {

			var templ1 = forTemplate{str3, var18, var2, var3, var4, var5, var6, var7, str4, var9, str2, var11, var12, var13, var14, var15, var16}

			fmt.Println(templ1)

			_ = globt.Execute(w, templ1)

		} else {

			return
		}
	*/
}

/////////
func display1(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	string1 = ""

	fmt.Println("in display 1")

	db := dbConn()

	globKeyword = "apple1"

	stmt, err := db.Prepare("SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
		"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
		"FROM products WHERE " +
		"((products.ProductKeyWord1 = ?) OR " +
		"(products.ProductKeyWord2 = ?) OR (products.ProductKeyWord3 = ? ))")
	if err != nil {
		panic(err.Error())
	}

	rows, err := stmt.Query(globKeyword, globKeyword, globKeyword)

	if err != nil {
		panic(err.Error())
	}

	var counter = 0

	var templ1 forTemplate

	for rows.Next() {

		var ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename, AmountToPurchaseID, AmountPurchasedID string

		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		if err != nil {
			panic(err.Error())
		}

		counter = counter + 1
		str := strconv.Itoa(counter)

		//var inputID = "inputID" + str
		var mainDivID = "mainDivID" + str
		var titleID = "titleID" + str
		var descID = "descID" + str
		var costID = "costID" + str
		var quantityID = "quantityID" + str
		var key1ID = "key1ID" + str
		var key2ID = "key2ID" + str
		var key3ID = "key3ID" + str
		AmountToPurchaseID = "amountID" + str
		AmountPurchasedID = "amountPID" + str

		templ1 = forTemplate{ProductID, ProductCatTitle, titleID, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity,
			key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword, ProductFilename, AmountToPurchaseID, AmountPurchasedID, mainDivID}

		fmt.Println(templ1)

		globt := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template1.html"))

		err1 := globt.Execute(w, templ1)

		//My idea was to update the template this way, however creates a new record.
		//ProductName = "this is a test"

		//templ1 = forTemplate{ProductID, ProductCatTitle, mainDiv, titleID, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity, key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword}

		//err1 = globt.Execute(w, templ1)

		if err1 != nil {
			fmt.Println("B---------------")
			fmt.Println(err1.Error())

			panic(err1.Error())

		}

	}

}

func submitfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("aarg ")
	fmt.Println("aarg ")
}

//send from client to server and
//send form server to client
func getMessages(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		City string `json:"city"`
	}

	w.Header().Set("Content-Type", "application/json")

	user := User{

		Name: "John Doe",
		Age:  10,
		City: "richmond"}

	json.NewEncoder(w).Encode(user)

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(j)
	fmt.Println("--wwww---")

}

//////

//type one struct {
//}

//type two struct {
//}

//func (m *one) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Listening on 8080: foo "))
//}
//func (m *two) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte("Listening on 8081: foo "))
//}
func main() {

	//finish := make(chan bool)

	//wg := new(sync.WaitGroup)
	//wg.Add(2)

	one := http.NewServeMux()

	//mux := http.NewServeMux()

	//has an id value passed in url
	one.HandleFunc("/updateForm/", updateForm)
	one.HandleFunc("/processSearch", processSearch)

	//button3 - just read session for right now
	one.HandleFunc("/getMessages", getMessages)

	one.HandleFunc("/display", display1)

	one.HandleFunc("/HelloWorld", HelloWorld)

	two := http.NewServeMux()

	//
	two.HandleFunc("/template2", purgeHTML)

	go func() {

		http.ListenAndServe(":8080", one)
	}()

	http.ListenAndServe(":8081", two)

	//wg.Wait()
}
