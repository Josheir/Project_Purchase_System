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

//cited
//https://www.bing.com/videos/search?q=youtbe+golang+template&refig=e742578f4d004a2b8a5bd1f28849eb0f&ru=%2fsearch%3fq%3dyoutbe%2bgolang%2btemplate%26form%3dANNTH1%26refig%3de742578f4d004a2b8a5bd1f28849eb0f&view=detail&mmscn=vwrc&mid=BD040005A2743ACB801ABD040005A2743ACB801A&FORM=WRVORC
//http://localhost:8080/golangproj/

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

////////https://stackoverflow.com/questions/21520244/how-to-simply-send-a-request-parameter-with-jquery-form-submit

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

//this is for testin, not used anympre
func processSearch(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rrrrrrraarg ")
	fmt.Fprintf(w, "got here1!")

	//parse like this not used with json - unmarshal instead?
	//err := r.ParseForm()
	//if err != nil {
	//	fmt.Println("error")

	//}
	//fmt.Fprintf(w, "got here2!")
	//fmt.Fprintln(w, "search :", r.Form.Get("search"))

	//globKeyword = r.Form.Get("a")
	//fmt.Println("-----")
	//fmt.Println("globKeyword")
	//fmt.Println("-----")

	//fmt.Println("here")
	//httpServletRequest.getParameter("myparam")

}

type forTemplate struct {
	ProductID          int
	ProductCatTitle    string
	MainDiv            string
	TitleID            string
	ProductFilename    string
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

	//ProductID int
}

type Name struct {
	FName string
	LName string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	name := Name{"mindorks", "Subject"}
	template, _ := template.ParseFiles("index2.html")
	template.Execute(w, name)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	name := Name{"mindorks2", "Subject2"}
	template, _ := template.ParseFiles("index2.html")
	template.Execute(w, name)
}
 
func updateForm(w http.ResponseWriter, r *http.Request) {

	//first, get the value of how many are available, if less than amount wanted create a message
	//parses all the values by id which is firtst value -  puts amount in form by parsing the amount with deduction
	//parses input as blank

	//first, get number available

	var ProductQuantity, amountPurchased int
	db := dbConn()

	//how many currently available in database
	stmt, err := db.Prepare("SELECT * FROM products WHERE ProductID = ?") 
		
	if err != nil {
		panic(err.Error())
	}


	//PRODUCT FILENAME LOOK AT THIS

	//var templ1 = forTemplate{ProductID, ProductCatTitle, mainDiv, titleID, ProductFilename, ProductName, descID, ProductDescription, 
	//costID, ProductCost, quantityID, ProductQuantity, key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword}

	rows, err := stmt.Query(r.FormValue("productID"));

	if err != nil {
		panic(err.Error())
	}
	var1 = r.FormValue("")
	var2 = r.FormValue("")
	var3 = r.FormValue("")
	var4 = r.FormValue("")
	var5 = r.FormValue("")
	var6 = r.FormValue("")
	var7 = r.FormValue("")
	var8 = r.FormValue("")
	var9 = r.FormValue("")
	var10 = r.FormValue("")
	var11 = r.FormValue("")
	var12 = r.FormValue("")
	var13 = r.FormValue("")
	var14 = r.FormValue("")
	var15 = r.FormValue("")
	var16 = r.FormValue("")
	var17 = r.FormValue("")
	var18 = r.FormValue("")


	err = rows.Scan(&ProductQuantity)

		if err != nil {
			panic(err.Error())
		}


		//enough product
		if (amountPurchased <=  ProductQuantity){
		
			productLeft = ProductQuantity - amounPurchased;
			
		}
		else 
		{
			return 1;
		}

		


	//r.ParseForm()
	//id := r.FormValue("ProductID")
	//row := db.QureryRow("SELECT * from eccomerce WHERE ProducID  = ?;", id);
	//var templ1 = forTemplate{ProductID, ProductCatTitle, mainDiv, titleID, ProductFilename, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity, key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword}

	var templ2 = forTemplate{0, "", "", "", "", "", "", "", "", 0, "", 0, "", "", "", "", "", ""}
	tpl.ExecuteTemplate(w, "deleteRecord.html", templ2)

}

/////////
func display1(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

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

	//counter := 0

	var counter = 0

	for rows.Next() {

		var ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename string

		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		if err != nil {
			panic(err.Error())
		}

		/////////////////

		counter = counter + 1
		str := strconv.Itoa(counter)

		//var inputID = "inputID" + str
		var mainDiv = "mainDivID" + str
		var titleID = "titleID" + str
		var descID = "descID" + str
		var costID = "costID" + str
		var quantityID = "quantityID" + str
		var key1ID = "key1ID" + str
		var key2ID = "key2ID" + str
		var key3ID = "key3ID" + str

		// add:  ProductFilename
		//var templ1 = forTemplate{mainDiv, titleID, ProductName, descID, ProductDescription, costID, ProductCost,
		//	quantityID, ProductQuantity, key1ID, gKeyword1, key2ID, gKeyword2, key3ID, gKeyword3}

		var templ1 = forTemplate{ProductID, ProductCatTitle, mainDiv, titleID, ProductFilename, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity, key1ID, globKeyword, key2ID, globKeyword, key3ID, globKeyword}

		//////////

		fmt.Println(templ1)

		t := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/index1.html"))

		err1 := t.Execute(w, templ1)

		if err1 != nil {
			fmt.Println("B---------------")
			fmt.Println(err1.Error())

			panic(err1.Error())

		}

	} //for selDB.Next()

	//display was here
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

func main() {

	mux := http.NewServeMux()

	//has an id value passed in url
	mux.HandleFunc("/updateForm/", updateForm)
	mux.HandleFunc("/processSearch", processSearch)

	//button3 - just read session for right now
	mux.HandleFunc("/getMessages", getMessages)

	mux.HandleFunc("/Hello", display1)

	http.ListenAndServe(":8080", mux)
}
