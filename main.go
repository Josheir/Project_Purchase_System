package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 // 1MB

var string1 = ""

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
func index(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rrrrrrrrg ")
	db := dbConn()

	//add products.ProductCatTitle = \"$titleOfSelectedDropDown\"
	var q = "SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
		"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle " +
		"FROM products WHERE " +
		"((products.ProductKeyWord1 = \"apple1\") OR " +
		"(products.ProductKeyWord2 = \"apple1\") or (products.ProductKeyWord3 = \"apple1\" ))"

	selDB, err := db.Query(q)
	if err != nil {
		panic(err.Error())
	}

	counter := 0

	for selDB.Next() {

		var ProductID, ProductCost, ProductQuantity int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle string

		err = selDB.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle)

		if err != nil {
			panic(err.Error())
		}

		///////////////////
		string1 += "<form" +

			"id=\"form\"" +
			"enctype=\"multipart/form-data\"" +
			"action=\"http://localhost:8080/upload\" " +
			"method=\"POST\">" +
			"<input class=\"input file-input\" type=\"file\" name=\"file\" />" +
			"<button class=\"button\" type=\"submit\">Submit</button>" +
			"</form>" +

			"<p style=\"color:Tomato;\" ><b>Images can not exceed 50 megabytes</p>" +
			"<input id='button2' type='button' value='Confirm Image'>" +
			"<br><br><br>" +
			"<input id='button1' type='button' value='Search Database'>"

			///////////////////
		counter = counter + 1
		str := strconv.Itoa(counter)

		var titleID = "titleID" + str
		var descID = "descID" + str
		var costID = "costID" + str
		var quantityID = "quantityID" + str
		var key1ID = "key1ID" + str
		var key2ID = "key2ID" + str
		var key3ID = "key3ID" + str

		string1 += "<p id = \"link1\">product id   : " + strconv.Itoa(ProductID) + " </p>" +
			"<p>category id  : " + ProductCatTitle + "</p>" +

			"<div class=\"container\">" +
			"<div class=\"row\" >" +

			"<div class=\"col\">" +
			"<h4><center><p id = \"\">Image</p></center></h4>" +
			"<center><p id = \"\"> image </p></center>" +
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
			"</div>" +

			"<div class=\"container\">" +
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

			"<br><br>" +
			"</div>" +
			"</div>" +

			"<div class=\"container\">" +
			"<div class=\"row\" >" +

			"<div class=\"col\">" +

			//<!--THIS LINE IS THE MAIN PROBLEM, THE FIRST PARAMETER WORKS BUT THE IDS HAVE A VALUE OF HTMLOBJECT
			//deleteFlag + "); , mainDiv , strconv.Itoa(ProductID) , titleID , descID,costID,quantityID ,key1ID,key2ID,key3ID
			"<center><button id = \"\" onclick = \"SaveProductItems(" + strconv.Itoa(ProductID) + ", " + titleID + ", " + descID + " )\">Submit</button></center>" +

			//<!--flag for determining if record delete will effect" sessioncount, 0 is no.-->
			//"<center><button id = \"\" onclick = \"deleteRecord( 1,  mainDiv, strconv.Itoa(ProductID) )\">Delete</button></center>" +
			//"<p><a href=\"#add\">To Add</a></p>" +
			//-->
			"</div>" +
			"<br><br>" +

			"</div>" +
			"</div>" +

			" <br><br><br><br>"

	} //for selDB.Next()

	receiveAjax(w, r)
}

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

////////

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("aaaa ")
	db := dbConn()

	////////////
	//	ProdID, ok := r.URL.Query()["ProductID"]
	//
	//	if !ok || len(ProdID[0]) < 1 {
	//		log.Println("Url Param 'key' is missing")
	//		return
	//	}
	//
	//	var productID = ProdID[0]

	var productID = 100

	var productID2 string
	productID2 = strconv.Itoa(productID)

	////////////

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}
	////////////
	// The argument to FormFile must match the name attribute
	// of the file input on the frontend
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	//////////
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filetype := http.DetectContentType(buff)
	if filetype != "image/jpeg" && filetype != "image/png" {
		http.Error(w, "The provided file format is not allowed. Please upload a JPEG or PNG image", http.StatusBadRequest)
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/////////

	////////
	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(fileHeader.Filename)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//image is saved, now save to database
	fmt.Println("Received data string ")

	////////////

	//db := dbConn()

	var counter = 0

	//////////////////////////////////////////////////////////
	//DATABASE BELOW HERE, ALREADY SAVED FILE
	//////////////////////////////////////////////////////////

	var ProductFilename string
	var q0 = "SELECT ProductFilename FROM products WHERE ProductID = " + (productID2)
	selDB, err := db.Query(q0)
	if err != nil {
		panic(err.Error())
	}
	//gets last number
	for selDB.Next() {

		counter = counter + 1
		_ = selDB.Scan(ProductFilename)
	}

	if counter > 0 {

		//if is record
		//delete

		var path = "/uploads" + ProductFilename
		//err := os.Remove(path)
		os.Remove(path)
	}

	//get latest number
	var Number int
	var q3 = "SELECT Number FROM  numbers"
	selDB1, err := db.Query(q3)
	if err != nil {
		panic(err.Error())
	}
	//gets last number
	for selDB1.Next() {

		_ = selDB1.Scan(Number)
	}

	var value1 = Number + 1
	var value2 = strconv.Itoa(value1)
	var Filename = "A" + (value2) + "." + filetype

	//put in database
	//var ProductID string
	var q1 = "UPDATE products SET  ProductFilename = '" + Filename + "' WHERE ProductID = " + productID2
	//selDB2, err := db.Query(q1)
	_, err = db.Query(q1)

	if err != nil {
		panic(err.Error())
	}

	//iterate - okay
	var q2 = "UPDATE numbers SET Number = " + (value2)

	//selDB3, err := db.Query(q2)
	_, err = db.Query(q2)

	if err != nil {
		panic(err.Error())
	}
}

func submitfunc(w http.ResponseWriter, r *http.Request) {

	fmt.Println("aarg ")
	fmt.Println("aarg ")
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/upload", uploadHandler)
	//form printout
	mux.HandleFunc("/index", index)
	http.ListenAndServe(":8080", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
