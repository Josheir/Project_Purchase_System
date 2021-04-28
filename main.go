package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"encoding/json"
	"html/template"

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

type forTemplate struct {
	ProductID       string
	ProductCatTitle string
	MainDiv         string
	TitleID         string
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
	//InputID            string
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

/////////
func display1(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	//temp := forTemplate{ProductName: "test"}

	//var templ1 = forTemplate{"a", "a", "a", "a", "a", "a", 1, "a", 1, "a", "a", "a", "a", "a", "a"}

	//fmt.Println(templ1)

	//w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//t, _ := template.New("test").Parse("{{.ProductName}} is name")

	//http.HandleFunc("/", HelloWorld)
	//name := Name{"hello"}

	//t := template.Must(template.ParseFiles("index2.html"))
	//if err1 != nil {
	//	fmt.Println("A--------------")
	//	fmt.Println(err1.Error())
	//
	//			panic(err1.Error())
	//		}

	//template.Execute(w , "")
	//_ = t.Execute(os.Stdout, name)

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

	for rows.Next() {

		var ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename string

		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		if err != nil {
			panic(err.Error())
		}

		///////////////////
		//	counter = counter + 1
		//	str := strconv.Itoa(counter)

		//var inputID = "inputID" + str
		//	var mainDiv = "mainDivID" + str
		//	var titleID = "titleID" + str
		//	var descID = "descID" + str
		//	var costID = "costID" + str
		//	var quantityID = "quantityID" + str
		//	var key1ID = "key1ID" + str
		//	var key2ID = "key2ID" + str
		//	var key3ID = "key3ID" + str

		// add:  ProductFilename
		//var templ1 = forTemplate{mainDiv, titleID, ProductName, descID, ProductDescription, costID, ProductCost,
		//	quantityID, ProductQuantity, key1ID, gKeyword1, key2ID, gKeyword2, key3ID, gKeyword3}

		var templ1 = forTemplate{"a", "a", "a", "a", "a", "a", "a", "a", 1, "a", 1, "a", "a", "a", "a", "a", "a"}

		fmt.Println(templ1)

		//w.Header().Set("Content-Type", "text/html; charset=utf-8")
		//t, _ := template.New("test").Parse("{{.ProductName}} is name")
		t := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/index1.html"))
		//if err1 != nil {
		//	fmt.Println("A--------------")
		//	fmt.Println(err1.Error())
		//
		//			panic(err1.Error())
		//		}

		//template.Execute(w , "")
		err1 := t.Execute(w, templ1)

		if err1 != nil {
			fmt.Println("B---------------")
			fmt.Println(err1.Error())

			panic(err1.Error())

		}

		/*
			string1 = string1 + "<p id = \"link1\">product id   : " + strconv.Itoa(ProductID) + " </p>" +
				"<p>category id  : " + ProductCatTitle + "</p>" +

				"<div id=" + m + " >" +
				"<div class=\"row\" > " +
				"<div class=\"col\">" +

				"<h4><center><p id = \"\">Image</p></center></h4>" +

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
				"<center><p> <input id = " + quantityID + " value = " + strconv.Itoa(ProductQuantity) + "  type=\"text\" name=\"title\" placeholder=\"\">	</p></center>" +
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
				"<center><p>	<input id =" + inputID + " type=\"number\" \" placeholder=\"\">		</p></center>" +
				//"<button id = \"\" onclick = \"Purchase(" + inputID + "," + strconv.Itoa(ProductID) + "," + quantityID + " )\">  Purchase</button>" +
				"<button id=\"\" onclick=\"Purchase(" + m + "," + strconv.Itoa(ProductID) + ",'" + key1ID + "'," + inputID + ")\">Purchase</button>" +

				//"," + quantityID + "," + mainDiv + "")\">Purchase</button></center>" + -->

				//"</div>"+

				" <br><br><br><br>" +
				"<hr>" +

				"</div></div></div>"
			//maindiv end tag


		*/

	} //for selDB.Next()

	//receiveAjax(w, r)
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

	//works
	//mux.HandleFunc("/Hello", Hello)
	mux.HandleFunc("/Hello", display1)

	//mux.HandleFunc("/", HelloWorld)

	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

//https://www.bing.com/videos/search?q=youtbe+golang+template&refig=e742578f4d004a2b8a5bd1f28849eb0f&ru=%2fsearch%3fq%3dyoutbe%2bgolang%2btemplate%26form%3dANNTH1%26refig%3de742578f4d004a2b8a5bd1f28849eb0f&view=detail&mmscn=vwrc&mid=BD040005A2743ACB801ABD040005A2743ACB801A&FORM=WRVORC
//http://localhost:8080/golangproj/
