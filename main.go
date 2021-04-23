package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

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

/////////
func display1(w http.ResponseWriter, r *http.Request) {

	string1 = ""

	/*
		err := r.ParseForm()
		if err != nil {
			fmt.Println("error")

		}

		fmt.Fprintf(w, "got here1!")
		fmt.Fprintln(w, "search :", r.Form.Get("search"))
	*/
	fmt.Println("in display 1")
	//fmt.Fprintf(w, "got here1!")
	//Var1 := "apple1"
	db := dbConn()

	///////////////////////////////////

	//https://dwahyudi.github.io/2020/05/15/mysql-where-in-query-with-golang.html
	//https://stackoverflow.com/questions/59005026/how-to-make-an-sql-query-in-golang-with-multiple-values-in-the-where-clause

	/*
	   	for results.Next() {
	   		var vehicle Vehicle
	   		err = results.Scan(&vehicle.ID, &vehicle.ProductionYear, &vehicle.Brand)
	   		panicError(err)
	   		fmt.Println(vehicle.ProductionYear)
	   		fmt.Println(vehicle.Brand)
	   		fmt.Println("=================")
	   	}
	   stmt, err := db.Prepare("select * from someTable where age = ? and hairColor = ?")
	   rows, err := stmt.Query(age,hairColor)
	*/
	////////////////////////////////////////

	//add products.ProductCatTitle = \"$titleOfSelectedDropDown\"
	//var q = "SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
	//	"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
	//	"FROM products WHERE " +
	//	"((products.ProductKeyWord1 = " + Var1 + ") OR " +
	//	"(products.ProductKeyWord2 = " + Var1 + ") OR (products.ProductKeyWord3 = " + Var1 + " ))"
	//add products.ProductCatTitle = \"$titleOfSelectedDropDown\"

	//var q = "SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
	//	"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle, products.ProductFilename " +
	//	//	"FROM products WHERE " +
	//	"FROM products WHERE " +
	//	"((products.ProductKeyWord1 = \"apple1\") OR " +
	//	"(products.ProductKeyWord2 = \"apple1\") OR (products.ProductKeyWord3 = \"apple1\" ))"

	globKeyword = "apple1"

	stmt, err := db.Prepare("SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
		"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
		"FROM products WHERE " +
		"((products.ProductKeyWord1 = ?) OR " +
		"(products.ProductKeyWord2 = ?) OR (products.ProductKeyWord3 = ? ))")
	if err != nil {
		panic(err.Error())
	}
	//globKeyword
	rows, err := stmt.Query(globKeyword, globKeyword, globKeyword)

	if err != nil {
		panic(err.Error())
	}

	counter := 0

	for rows.Next() {

		var ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename string

		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		if err != nil {
			panic(err.Error())
		}

		///////////////////
		counter = counter + 1
		str := strconv.Itoa(counter)

		var inputID = "inputID" + str
		var mainDiv = "mainDiv" + str
		var titleID = "titleID" + str
		var descID = "descID" + str
		var costID = "costID" + str
		var quantityID = "quantityID" + str
		var key1ID = "key1ID" + str
		var key2ID = "key2ID" + str
		var key3ID = "key3ID" + str

		string1 = string1 + "<p id = \"link1\">product id   : " + strconv.Itoa(ProductID) + " </p>" +
			"<p>category id  : " + ProductCatTitle + "</p>" +

			"<div id = \"mainDiv>\"" +
			"<div class=\"row\" > " +
			"<div class=\"col\">" +

			"<h4><center><p id = \"\">Imagejjjjjjjjjjjjjjjjjjjj</p></center></h4>" +

			"<img  height=\"100\" width=\"100\"  src= /golangproj/uploads/" + ProductFilename + " alt=\"no product image\">" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id =\"\"  >Title</p></center></h4>" +

			"<center>      <p  >      <input id = " + titleID + " value = " + ProductName + " type=\"text\" name=\"title\" placeholder=\"\"></p></center>" +
			"</div>" +

			"<div class=\"col\">" +

			"<h4><center><p id = \"\">Desc</p></center></h4>" +

			"<center><textarea wrap id = " + descID + "   value = " + ProductDescription + "  type=\"text\" rows=\"5\" cols=\"34\"></textarea></center>" +
			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Cost</p></center></h4>" +
			"<center><p>	<input id = " + costID + " value =  " + strconv.Itoa(ProductCost) + "   type=\"number\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"</div>" +

			"<div class=\"row\" >" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Quantity</p></center></h4>" +
			"<center><p> <input id = " + quantityID + " value = " + strconv.Itoa(ProductQuantity) + "  type=\"number\" name=\"title\" placeholder=\"\">	</p></center>" +
			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 1</p></center></h4>" +
			"<center><p>	<input id = " + key1ID + " value = " + gKeyword1 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 2</p></center></h4>" +
			"<center><p>	<input id = " + key2ID + " value = " + gKeyword2 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\" >Keyword 3</p></center></h4>" +
			"<center><p>	<input id =   " + key3ID + "  value = " + gKeyword3 + " type=\"text\" name=\"title\" placeholder=\"\">		</p></center>" +

			"</div>" +
			"</div>" +
			"<br><br>" +

			"<div class=\"row\" >" +

			"<div class=\"col\">" +

			"<br><br><br><br>" +
			//inputID is the quant amount  to purchase
			"<center><p>	<input id =   " + inputID + " type=\"number\" \" placeholder=\"\">		</p></center>" +
			"<center><button id = \"\" onclick = \"Purchase(" + inputID + "," + strconv.Itoa(ProductID) + "," + quantityID + "," + mainDiv + ")\">Purchase</button></center>" +

			"</div>" +

			" <br><br><br><br>" +
			"<hr>" +

			"</div>" +
			"</div>" //maindiv end tag

	} //for selDB.Next()

	receiveAjax(w, r)
}

func submitfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("aarg ")
	fmt.Println("aarg ")
}

/////////

//send from client to server and
//send form server to client
func getMessages(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	err := r.ParseForm()
	if err != nil {
		fmt.Println("error")

	}

	//doesn't work
	for key, values := range r.Form {
		fmt.Println(key, values)

	}

	//results := Results{Total: 100}
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//into json string

	//j, err := json.Marshal(a)
	//if err != nil {
	//	fmt.Printf("Error: %s", err.Error())
	//	fmt.Println("---qqq--")
	//}
	fmt.Println("--wwww---")

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

	mux.HandleFunc("/processSearch", processSearch)

	//button3 - just read session for right now
	mux.HandleFunc("/getMessages", getMessages)

	mux.HandleFunc("/display", display1)

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
