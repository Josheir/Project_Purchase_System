package main

//https://github.com/go-session/session

//https://stackoverflow.com/questions/32087233/how-does-mysql-handle-concurrent-inserts
//http://go-database-sql.org/prepared.html
//https://stackoverflow.com/questions/37404989/whats-the-difference-between-db-query-and-db-preparestmt-query-in-golang
//https://golangdocs.com/mysql-golang-crud-example

import (
	"database/sql"
	"encoding/json"

	//"errors"
	"fmt"
	"html/template"
	//"io/ioutil"

	//"log"
	//"strings"

	//"bytes"
	"net/http"

	"context"
	"math"

	//"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var globalFlagisVariable = "no"

type product struct {
	ProductQuantity int
	ProductName     string
	ProductCatTitle string
	ProductCost     int
}

type Product1 struct {
	CondYellow         int
	ProductIDID        string
	RemoveRecordDivID  string
	GrandTotalStringID string

	GrandTotalString string
	BoughtID         string
	Bought           int
	TotalCost        string
	TotalCostID      string
	CostID           string
	AmountToBuyID    string
	Condition        int
	Condition2       int
	ProductID        int
	ProductQuantity  int
	ProductName      string
	DivID            string
	ProductCatTitle  string
	ProductCost      string
}

//spit back to last html page
type Product2 struct {
	ID                int
	QuantityAvailable int
	IsEnoughQuantity  bool
}

type HoldsFlag struct {
	Flag string
}
type User1 struct {
	Text   string
	UserID int
}

//used in createtemplate2
var ProductList = []Product1{}

//var ProductList2 = []Product2{}

//used in spitback
var ProductList2A = []Product2{}
var User = []User1{}

//https://www.bing.com/videos/search?q=youtbe+golang+template&refig=e742578f4d004a2b8a5bd1f28849eb0f&ru=%2fsearch%3fq%3dyoutbe%2bgolang%2btemplate%26form%3dANNTH1%26refig%3de742578f4d004a2b8a5bd1f28849eb0f&view=detail&mmscn=vwrc&mid=BD040005A2743ACB801ABD040005A2743ACB801A&FORM=WRVORC
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

//prod := Product2{
//
//	IsEnoughQuantity:  enough,
//	QuantityAvailable: quant,
//	ID:                id,
//}

func MakeUser(text string, userid int) {

	user := User1{
		Text:   text,
		UserID: userid,
	}

	User = append(User, user)

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

// USD represents US dollar amount in terms of cents
type USD int64

// Float64 converts a USD to float64
func (m USD) Float64() float64 {
	x := float64(m)
	x = x / 100
	return x
}

// Multiply safely multiplies a USD value by a float64, rounding
// to the nearest cent.
func (m USD) Multiply(f float64) USD {
	x := (float64(m) * f) + 0.5
	return USD(x)
}

// String returns a formatted USD value
func (m USD) String() string {
	x := float64(m)
	x = x / 100
	return fmt.Sprintf("$%.2f", x)
}

////////////
// ToUSD converts a float64 to USD
// e.g. 1.23 to $1.23, 1.345 to $1.35
func ToUSD(f float64) USD {
	return USD((f * 100) + 0.5)
}

//////////

//https://www.bing.com/search?q=receiver%20int%20golang&qs=n&form=QBRE&sp=-1&pq=receiver%20int%20golang&sc=0-19&sk=&cvid=14C3226BD73C46F09A57AA46291441EA
func addElement(var1 int, var2 string, var3 string, var4 int) {

	var element product
	element.ProductQuantity = var1
	element.ProductName = var2
	element.ProductCatTitle = var3
	element.ProductCost = var4

}

type Product3 struct {
	ID    int
	Quant int
}

func makeListForLastpageA(enough bool, id int, quant int) {

	//to spit back to html
	prod := Product2{

		IsEnoughQuantity:  enough,
		QuantityAvailable: quant,
		ID:                id,
	}
	//list to spit back to html for rewriting all the quant
	ProductList2A = append(ProductList2A, prod)
}

//this last page is where the data is spat back to html to note any database changes that cause purchase impossible
//func makeListForLastpage(id int, quant int) {
//
//	//to spit back to html
//	prod := Product2{
//
//		QuantityAvailable: quant,
//		ID:                id,
//	}
//	//list to spit back to html for rewriting all the quant
//	ProductList2 = append(ProductList2, prod)
//}

var orderid1 = 100

//https://www.geeksforgeeks.org/how-to-get-current-time-in-golang/
func processLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	/*
		err = r.ParseForm()

		array := r.Form["var"][1]
		fmt.Println(array)

		value := r.Form["var"]

		a := value[0]

	*/

	User = nil

	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, err)
	}

	//query := r.URL.Query()

	//userid, present := query["userid"]
	userid := r.Form["userid"][0]

	//if !present || len(userid) == 0 {
	//	fmt.Println("filters not present")
	//}

	//string to int
	userid1, err := (strconv.Atoi(userid))

	if err != nil {
		fmt.Fprint(w, err)
	}

	//pass, present := query["pass"]

	//if !present || len(pass) == 0 {
	//	fmt.Println("filters not present")
	//}
	pass := r.Form["pass"][0]

	db := dbConn()

	stmt, err := db.Prepare("SELECT customers.Password FROM customers WHERE customers.CustomerID = ?")

	if err != nil {
		fmt.Fprint(w, err)
	}

	rows, err := stmt.Query(userid1)

	if err != nil {
		fmt.Fprint(w, err)
	}

	var PasswordID string

	for rows.Next() {

		err = rows.Scan(&PasswordID)
		if err != nil {
			fmt.Println(err)
		}

	}

	passFlag := "no"

	if PasswordID == "" {
		passFlag = "password wrong"
	} else if PasswordID == pass {

		passFlag = "password correct"

		var UserID = 1
		//var userID = 1
		//DOES THIS PRODUCT RECORD ALREADY EXIST
		stmt, err := db.Prepare("DELETE FROM savedtext WHERE savedtext.UserID = ?")

		if err != nil {
			fmt.Println(err)
		}

		stmt.Exec(UserID)

		////////////

		stmt2, err := db.Prepare("INSERT INTO savedtext(Text, UserID) VALUES(?,?)")
		if err != nil {
			fmt.Println(err)
		}
		stmt2.Exec("[1]", UserID)

		/////////////

	} else {

		passFlag = "password wrong"
	}

	MakeUser(passFlag, userid1)

	json.NewEncoder(w).Encode(User)

}

/////////////////////////////////////////
/////////////////////////////////////////

/////////////////////////////////////////
////////////////////////////////////////

func spitBackAmounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	ctx := context.Background()
	ProductList2A = nil

	query := r.URL.Query()

	allIds, present := query["id"]

	if !present || len(allIds) == 0 {
		fmt.Println("filters not present")
	}

	allQuants, present := query["quant"]
	//in template 2 bought column
	if !present || len(allQuants) == 0 {
		fmt.Println("filters not present")
	}

	userID1, present := query["userid"]

	if !present || len(userID1) == 0 {
		fmt.Println("filters not present")
	}

	db := dbConn()

	var thisProductID = 0
	DatabaseQuantity := 0

	var haveWrittenOrder bool = false
	var j = 0
	var didRollback = false

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	for j = 0; j < len(allIds); j++ {

		thisProductID, _ = strconv.Atoi(allIds[j])

		var enough bool = false

		val1, err1 := strconv.Atoi(allIds[j])
		if err1 != nil {
			fmt.Println(err)
		}

		//checks if enough product to remove the quantity in database

		err := tx.QueryRowContext(ctx, "SELECT (products.ProductQuantity >= 0)  FROM products WHERE products.ProductID = ? AND products.ProductStatus = 'ready' ", allIds[j]).Scan(&enough)

		if err != nil {
			fmt.Println(err)
		}

		if !enough {

			didRollback = true
			err := tx.Rollback()
			if err != nil {
				fmt.Println(err)
			}

		}

		//after rollback finish with making list to pass to template2
		if didRollback {

			var k = 0
			for k = 0; k < len(allIds); k++ {
				enough = false
				stmt, err := db.Prepare("SELECT products.ProductQuantity FROM products WHERE products.ProductID = ? AND products.ProductStatus = 'ready'")

				if err != nil {
					fmt.Println(err)
				}

				rows, err := stmt.Query(allIds[k])

				if err != nil {
					fmt.Println(err)
				}

				var prodQuant int

				for rows.Next() {

					//database hold this
					err = rows.Scan(&prodQuant)

					if err != nil {
						fmt.Println(err)
					}

					quantPurchasing, err := strconv.Atoi(allQuants[k])
					if err != nil {
						fmt.Println(err)
					}

					//in database for product -
					if DatabaseQuantity-quantPurchasing >= 0 {
						enough = true
					}
					makeListForLastpageA(enough, (val1), quantPurchasing)

				}
			}

			return
		}

		intQuant, err := strconv.Atoi(allQuants[j])
		if err != nil {
			fmt.Println(err)
		}
		//in database for product -
		if DatabaseQuantity-intQuant >= 0 {
			enough = true
		}

		//intQuant is amount purchasing
		makeListForLastpageA(enough, (val1), intQuant)
		enough = false

		//this gets the record for insert of quant

		var ProductCost float64
		var ProductQuantity, ProductID, AdminID, CustomerID, OrderID, ID int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename, ProductStatus string

		//if client b is passed this than quantity will be the same as client a, so whole thing needs to be transaction
		//because productquantity is used

		//gets all the fields of data from  a particular productid and ready status
		err = tx.QueryRowContext(ctx, "SELECT * FROM products WHERE products.ProductID = ? and products.ProductStatus = 'ready' ", allIds[j]).Scan(
			&ProductFilename, &ProductName, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &gKeyword1, &gKeyword2, &gKeyword3, &CustomerID,
			&OrderID, &ProductStatus, &AdminID, &ProductID, &ID)

		if err != nil {
			fmt.Println(err)
		}

		ProductID = thisProductID
		var thisQuant = ProductQuantity - intQuant
		//updates productid fields to its quantity minus int-quant
		_, err = tx.ExecContext(ctx, "Update products SET ProductQuantity = ? WHERE products.ProductID = ? and products.ProductStatus = 'ready' ", thisQuant, allIds[j])
		if err != nil {
			fmt.Println(err)
		}

		datetime := time.Now()

		var id1 = 0

		var productQuant int64
		var order_ID int64
		//check if there is a product record ctreated to store in order table, if there is, than an order record has been created too.
		err = tx.QueryRowContext(ctx, "SELECT products.OrderID, products.ProductQuantity  FROM products WHERE products.ProductID =  ? and  products.ProductStatus = 'purchased'", allIds[j]).Scan(&id1, &productQuant)
		if err != nil {
			fmt.Println(err)
		}
		//no record of product created to store in order, so create both

		if err == sql.ErrNoRows {

			if !haveWrittenOrder {
				res, err := tx.ExecContext(ctx, "INSERT INTO orders (OrderDate) values(?)", datetime)

				if err != nil {
					fmt.Println(err)
				}

				order_ID, err = res.LastInsertId()

				if err != nil {
					fmt.Println(err)
				}

				//lastID++

				haveWrittenOrder = true
			}

			//insert product to store in table - change quantity and status

			ProductStatus = "purchased"
			_, err = tx.ExecContext(ctx, "INSERT INTO products (ProductFilename, ProductName, ProductDescription, ProductCost, ProductQuantity, ProductCatTitle,ProductKeyword1,ProductKeyword2 , ProductKeyword3, CustomerID, OrderID, ProductStatus, AdminID, ProductID) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", ProductFilename, ProductName, ProductDescription, ProductCost, (int64(intQuant) + productQuant), ProductCatTitle, gKeyword1, gKeyword2, gKeyword3, CustomerID, order_ID, ProductStatus, AdminID, ProductID)

			if err != nil {
				fmt.Println(err)
			}

			//there is a product created to store in order table so than the order table record has also been created, so no need to do anything
		} else {

			//update product with status of purchased from product table:  original quantity + intQuant
			//productquant is quantity of product
			//intquant is database quantity taken from the allquant array
			_, err = tx.ExecContext(ctx, "Update products SET ProductQuantity = ?, OrderID = ?  WHERE products.ProductID = ? and products.ProductStatus = 'purchased' ", (int64(intQuant) + productQuant), int64(order_ID), allIds[j])
			if err != nil {
				fmt.Println(err)
			}

		}

	} //for

	err5 := tx.Commit()
	if err5 != nil {
		fmt.Println(err5)
	}

	//if !didRollback {
	json.NewEncoder(w).Encode(ProductList2A)
	//}

}

//https://github.com/strongo/decimal
//https://programming.guide/go/convert-int64-to-string.html

//stackoverflow.com/questions/54362751/how-can-i-truncate-float64-number-to-a-particular-precision
//stackoverflow.com/questions/4187146/truncate-number-to-two-decimal-places-without-rounding#:~:text=General%20solution%20to%20truncate%20%28no%20rounding%29%20a%20number,with%20exactly%20n%20decimal%20digits%2C%20for%20any%20n%E2%89%A50.

var Condition = 0

func createTemplate2(w http.ResponseWriter, r *http.Request) {

	ProductList = nil

	w.Header().Set("Access-Control-Allow-Origin", "*")

	//fmt.Println("+++++++++++++++++++++++++++++++++++++AAAAAAAAAAAAAAAAAAAAAAAAAA++++++++++++++++")
	fmt.Println("GET params were:", r.URL.Query())

	query := r.URL.Query()

	//filters=["color", "price", "brand"]
	allIds, present := query["id"]
	fmt.Println("allIds")
	fmt.Println(allIds)

	if !present || len(allIds) == 0 {

		fmt.Println("filters not present")

	}

	//fmt.Println(allIds[0])

	allQuants, present := query["quant"]

	if !present || len(allQuants) == 0 {
		fmt.Println("filters not present")
	}

	/////////////

	//////////
	/////////

	string1 = ""

	db := dbConn()

	var var1 = "D"
	var var2 = "A"
	var var3 = "C"
	var var4 = "TC"
	var var5 = "B"
	var var6 = "GT"
	var var7 = "V"
	var var8 = "P"
	//yes this is right product starts at one

	//var j = 1
	//var ProductID = 2
	var i = 0

	//fmt.Println("how many times!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	Condition++
	var Condition2 = 0

	//error when making two product and pressing checkout displays all table data below main table with another checkout button
	//6/26/21

	ID := 0
	bought := 0
	numTotal := 0.0

	for i = 0; i < len(allIds); i++ {

		fmt.Println("length")
		fmt.Println(len(allIds))

		//Condition++
		Condition2++

		DivID := var1 + (strconv.Itoa(i))
		AmountToBuyID := var2 + (strconv.Itoa(i))
		CostID := var3 + (strconv.Itoa(i))
		TotalCostID := var4 + (strconv.Itoa(i))
		BoughtID := var5 + (strconv.Itoa(i))
		GrandTotalStringID := var6 + (strconv.Itoa(i))
		RemoveRecordDivID := var7 + (strconv.Itoa(i))
		ProductIDID := var8 + (strconv.Itoa(i))
		//ID := var8 + (strconv.Itoa(i))

		var prodid, err = (strconv.Atoi(allIds[i]))
		if err != nil {
			fmt.Println(err)
		}

		//Total := ""
		stmt, err := db.Prepare("SELECT products.ProductQuantity,products.ProductName,products.ProductCatTitle, products.ProductCost  " +
			"FROM products WHERE " +
			"products.ProductID = ? AND products.ProductStatus = 'ready'")

		if err != nil {
			fmt.Println(err)
		}

		rows, err := stmt.Query(prodid)

		if err != nil {
			fmt.Println(err)
		}

		//var counter = 0

		//var templ1 product

		var ProductQuantity int

		var ProductName, ProductCatTitle, ProductCost, TotalCost string

		//fmt.Println("ProductList")
		//fmt.Println(fmt.Sprintf("%+v", ProductList))

		//defer rows.Close()

		//jumps past this, first run through
		for rows.Next() {

			//fmt.Println("ProductList1")
			//fmt.Println(fmt.Sprintf("%+v", ProductList))

			//copies from database row to these variables
			err = rows.Scan(&ProductQuantity, &ProductName, &ProductCatTitle, &ProductCost)
			if err != nil {
				fmt.Println(err)
			}

			var j = 0
			for j = 0; j < len(allIds); j++ {

				bought, err = (strconv.Atoi(allQuants[j]))
				if err != nil {
					fmt.Println(err)
				}

				ID, err = strconv.Atoi(allIds[j])
				if err != nil {
					fmt.Println(err)
				}

				//?????????????
				//there is boutght total that goes with this id
				if prodid == ID {

					//subtract bought from quantity
					ProductQuantity = ProductQuantity - bought
					break
				}

			} //for

			//value1 is product ID
			if i == 0 {
				Condition = 1

				//fmt.Println("++++++++++++++++++this is important, here.++++++++++++++++++++++++++++++")
			} else {
				Condition = 0

			}
			if i == (len(allIds) - 1) {
				Condition2 = -1
			}

			aQuant, err3 := strconv.Atoi(allQuants[i])
			if err3 == nil {
				fmt.Println(var2)
			}

			//https://yourbasic.org/golang/round-float-2-decimal-places/
			//https://stackoverflow.com/questions/20596428/how-to-represent-currency-in-go
			//https://www.bing.com/search?q=put%20commas%20in%20string%20golang&qs=n&form=QBRE&sp=-1&pq=put%20commas%20in%20string%20golang&sc=0-27&sk=&cvid=D3A2A7E4E0E141BCAA5BA7E7EE279532
			//quantity
			//However, whole numner
			var QuantityFloat float64 = float64(aQuant)
			//in cents, no decimal
			ProductCostString := ProductCost

			//cents
			ProductCostFloat, err := strconv.ParseFloat(ProductCostString, 64)
			if err != nil {
				fmt.Println(err)
			}
			//cents
			ProductCostFloat2 := ProductCostFloat

			//move decimal
			ProductCostFloat = ProductCostFloat / (math.Pow(10, 2))
			//take care of decimals - ready for display
			ProductCostString = fmt.Sprintf("%.2f", ProductCostFloat)

			//in cents and wholenumber
			TotalCostFloat := QuantityFloat * ProductCostFloat2

			//cents
			TotalCostFloat2 := TotalCostFloat
			//move decimal
			TotalCostFloat = TotalCostFloat / (math.Pow(10, 2))
			//take care of decimals
			TotalCost = fmt.Sprintf("%.2f", TotalCostFloat)

			//cents
			numTotal = numTotal + TotalCostFloat2

			tax := 0.0

			//numTotal = numTotal  * .05
			//this is the tax  amount
			var GrandTotalString = "this text doesnt display"
			if i == (len(allIds) - 1) {
				tax = numTotal * 5
				numTotal = numTotal * 100
				numTotal = numTotal + tax

				numTotal = numTotal / math.Pow(10, 4)
				GrandTotalString = fmt.Sprintf("%.2f", numTotal)
			}

			addProduct(ProductIDID, RemoveRecordDivID, GrandTotalStringID, GrandTotalString, BoughtID, bought, TotalCost, TotalCostID, ProductQuantity, CostID, AmountToBuyID, Condition, Condition2, prodid, ProductQuantity, ProductName, DivID, ProductCatTitle, ProductCostString)

		}
		//fmt.Println("ProductListXXX")
		//fmt.Println(fmt.Sprintf("%+v", ProductList))

	} //for next loop

	///////////

	//https://stackoverflow.com/questions/24755509/using-conditions-inside-templates
	globt = template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template2.html"))
	//fmt.Println("ProductList2")
	//fmt.Println(fmt.Sprintf("%+v", ProductList))

	err1 := globt.Execute(w, ProductList)

	if err1 != nil {
		//fmt.Println("CC---------------")
		fmt.Println(err1.Error())

		//panic(err1.Error())

	}

	///////////

}

//func removeAllProducts() {
//
//	if globalFlagisVariable == "yes" {
//
//		ProductList = ProductList[:0]
//
//	}
//}

func addProduct(productidid string, removerecorddivID string, totalID string, total string, boughtid string, bought int, totalcost string, totalcostid string, ProductQuantity int, costid string, amountid string, condition int, condition2 int, prodid int, quant int, name string, div string, cat string, cost string) {

	prod := Product1{
		ProductIDID:        productidid,
		RemoveRecordDivID:  removerecorddivID,
		GrandTotalStringID: totalID,
		GrandTotalString:   total,
		BoughtID:           boughtid,
		Bought:             bought,
		TotalCost:          totalcost,
		TotalCostID:        totalcostid,
		CostID:             costid,
		AmountToBuyID:      amountid,
		Condition:          condition,
		Condition2:         condition2,
		ProductID:          prodid,
		ProductQuantity:    quant,
		ProductName:        name,
		DivID:              div,
		ProductCatTitle:    cat,
		ProductCost:        cost,
	}
	flag := "nonefound"

	//ProductList = append(ProductList, prod)

	//could be done better, if time allows
	for i := 0; i < len(ProductList); i++ {
		if (ProductList[i].ProductID) == prodid {
			//ProductQuantityFloat = float64(ProductQuantity)
			//prod.Bought = prod.Bought + bought
			ProductList[i].ProductQuantity = ProductQuantity
			ProductList[i].Bought = bought
			ProductList[i].TotalCost = totalcost
			//break out
			globalFlagisVariable = "yes"
			flag = "found"
			i = 100
		}
	}

	if flag != "found" {
		//prod.ProductQuantity = prod.ProductQuantity - prod.Bought
		ProductList = append(ProductList, prod)
		globalFlagisVariable = "yes"
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Rrrrrrraarg ")
}

////////example:
/*

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
*/

type forTemplate struct {
	CondYellow      int
	Link            string
	Condition       int
	AmountPurchased int
	ProductID       int
	ProductCatTitle string
	//MainDiv         string
	TitleID string
	//ProductFilename    string
	ProductName        string
	DescID             string
	ProductDescription string
	CostID             string
	ProductCost        float64
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

	globt := template.Must(template.ParseFiles("C:/wamp64/www/golangproj/twemplate1.html"))

	err1 := globt.Execute(w, var2)

	if err1 != nil {
		fmt.Println("B---------------")
		fmt.Println(err1.Error())

		panic(err1.Error())

	}

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

//var counter1 = -1

//////////

//this function displays all the keywords in the url, these records are the records already been displayed in display1
func display2(w http.ResponseWriter, r *http.Request) {

	//ARRAY OF INTS  [3,4,7]
	//thesse ints are kept in database and changed to an array to see if they have aleady been
	//displayed so continue.  Does not effect the client side is an int
	//product is checked with array and if exists is contniues

	GlobCounter++
	//store, err := session.Start(context.Background(), w, r)

	fmt.Println("+++++++++++++++++")

	query := r.URL.Query()

	//this is the searchterm in order from first to last now
	key1, present := query["var"]

	if !present || len(key1) == 0 {
		fmt.Println("filters not present1")
	}

	keyTotalAmountBought, present2 := query["quant"]
	if !present2 || len(keyTotalAmountBought) == 0 {
		fmt.Println("filters not present2")
	}
	ProdID, present3 := query["id"]
	if !present3 || len(ProdID) == 0 {
		fmt.Println("filters not present3")
	}

	UserIDstring, present4 := query["uid"]
	if !present4 || len(UserIDstring) == 0 {
		fmt.Println("filters not present4")

	}

	//var UID = 0
	//var err= ""
	//if len(UserIDstring) != 0 {

	//only one  // UserID =
	UID, err := strconv.Atoi(UserIDstring[0])
	if err != nil {
		fmt.Println(err)
	}
	//} else {
	//isnt used (is still in dispaly1, for int[] array saved in db as string):
	//UserID = 1
	//}

	globKeyword := key1[0]

	w.Header().Set("Access-Control-Allow-Origin", "*")

	string1 = ""

	fmt.Println("in display 1")

	db := dbConn()

	/*
		_, ok := store.Get("isFirstUse")
		if !ok {

			store.Set("isFirstUse", "yes")
			err = store.Save()
			if err != nil {
				fmt.Fprint(w, err)
				//	return
			}

		}
	*/

	var m = 0
	//var lastUseOf = false
	var numRecords1 = 0
	var oneTime = true
	var recordCounter = 0
	for m = 0; m < len(key1); m++ {

		recordCounter = 0
		//////////

		//get amount of records for the last keyword, one time
		if m == (len(key1)-1) && oneTime == true {
			oneTime = false
			db4 := dbConn()

			stmt, err := db4.Prepare("SELECT COUNT(*) FROM products WHERE ((products.ProductKeyWord1 = ?) OR (products.ProductKeyWord2 = ?) OR " +
				"(products.ProductKeyWord3 = ? )) AND products.ProductStatus = 'ready'")

			if err != nil {
				fmt.Println(err)
			}

			rows, err := stmt.Query(globKeyword, globKeyword, globKeyword)

			if err != nil {
				fmt.Println(err)
			}

			for rows.Next() {

				err = rows.Scan(&numRecords1)

				if err != nil {
					fmt.Println(err)
				}

			}

		}

		///////////

		globKeyword = key1[m]

		//get records that use keywords

		stmt, err := db.Prepare("SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
			"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
			"FROM products WHERE " +
			"((products.ProductKeyWord1 = ?) OR " +
			"(products.ProductKeyWord2 = ?) OR (products.ProductKeyWord3 = ? )) AND products.ProductStatus = 'ready'")
		if err != nil {
			fmt.Println(err)
		}

		rows, err := stmt.Query(globKeyword, globKeyword, globKeyword)

		if err != nil {
			fmt.Fprint(w, err)
		}

		//var templ1 forTemplate

		var Link = globKeyword

		var Condition = 0

		var lastProductID = -1

		var stringText = ""
		var ints []int
		for rows.Next() {

			/////////////

			db5 := dbConn()
			stmt1, err := db5.Prepare("SELECT savedtext.Text FROM savedtext WHERE savedText.UserID = ?")

			if err != nil {
				fmt.Println(err)
			}

			rows1, err := stmt1.Query(UID)

			if err != nil {
				fmt.Println(err)
			}

			//get string from database
			for rows1.Next() {

				err = rows1.Scan(&stringText)
				if err != nil {
					fmt.Println(err)
				}

				//change string to array
				err := json.Unmarshal([]byte(stringText), &ints) //
				if err != nil {
					fmt.Println(err)
				}

			}

			////////////

			var j = 0
			//check if productID was already created in display1
			for j = 0; j < len(ints); j++ {
				if ProductID == ints[j] {

					continue
				}

			}
			/////////////

			recordCounter++

			//marshalFlag = "no"

			Condition++
			var ProductCost float64
			var ProductQuantity, CondYellow int
			var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename string

			CondYellow = 0
			err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

			if ProductID == lastProductID {
				continue
			}

			lastProductID = ProductID

			if err != nil {
				fmt.Println(err)
			}

			///////////////
			var i = 0
			var continueFlag = "no"
			counter1 = counter1 + 1
			//if finds a record that aleady exists as a passed in url parameter, than create the template with the value and continue the main for.
			for i = 0; i < len(ProdID); i++ {

				prodID, err := strconv.Atoi(ProdID[i])
				if err != nil {
					fmt.Println(err)
				}

				if prodID == ProductID {
					continueFlag = "continue"

					prodBoughtStr := keyTotalAmountBought[i]
					prodBoughtInt, err := strconv.Atoi(prodBoughtStr)
					if err != nil {
						fmt.Println(err)
					}

					CondYellow = 1
					AmountPurchased := prodBoughtInt
					sendToTemplate(&globKeyword, &counter1, &w, &CondYellow, &Link, &Condition, &AmountPurchased, &ProductID, &ProductCatTitle, &ProductName, &ProductDescription, &ProductCost, &ProductQuantity,
						&gKeyword1, &gKeyword2, &gKeyword3, &ProductFilename)

					break
				} //if

			} //for i

			//productid found a record that already existed, so get next productid
			if continueFlag == "continue" {
				continue
			}

			///////////////the productid record does not exist yet, create the record which is part of the current keyword

			AmountPurchased := 0
			sendToTemplate(&globKeyword, &counter1, &w, &CondYellow, &Link, &Condition, &AmountPurchased, &ProductID, &ProductCatTitle, &ProductName, &ProductDescription, &ProductCost, &ProductQuantity,
				&gKeyword1, &gKeyword2, &gKeyword3, &ProductFilename)

		} //row

		/////////

	} //main loop

}

/////////

var GlobCounter = -1
var counter1 = 0

////////
func sendToTemplate(globKeyword *string, counter1 *int, w *http.ResponseWriter, CondYellow *int, Link *string, Condition *int, AmountPurchased *int, ProductID *int, ProductCatTitle *string, ProductName *string, ProductDescription *string, ProductCost *float64, ProductQuantity *int,
	gKeyword1 *string, gKeyword2 *string, gKeyword3 *string, ProductFilename *string) {
	*counter1++
	//counter1 = 0
	str := strconv.Itoa(*counter1)

	//var inputID = "inputID" + str
	var mainDivID = "mainDivID" + str
	var titleID = "titleID" + str
	var descID = "descID" + str
	var costID = "costID" + str
	var quantityID = "quantityID" + str
	var key1ID = "key1ID" + str
	var key2ID = "key2ID" + str
	var key3ID = "key3ID" + str
	var AmountToPurchaseID = "amountID" + str
	var AmountPurchasedID = "amountPID" + str

	//AmountPurchased = 120
	//json.NewEncoder(*w).Encode(globKeyword)

	//AmountPurchased = prodBoughtInt
	templ1 := forTemplate{*CondYellow, *Link, *Condition, *AmountPurchased, *ProductID, *ProductCatTitle, titleID, *ProductName, descID, *ProductDescription, costID, *ProductCost, quantityID, *ProductQuantity,
		key1ID, *gKeyword1, key2ID, *gKeyword2, key3ID, *gKeyword3, *ProductFilename, AmountToPurchaseID, AmountPurchasedID, mainDivID}

	fmt.Println(templ1)

	globt = template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template1.html"))

	//err1 := globt.Execute(w, testvar)
	var err1 = globt.Execute(*w, templ1)

	if err1 != nil {
		fmt.Println("---------------")
		fmt.Println(err1.Error())

	}
}

////////

//the purpose of this function is to display the information of the keyword sent here.
//the actual ids are stored in a database when  they have been used
//if there are no ID/Quantity ordered url parameters than the function creates a new
//record with zero value for AmountPurchased.  Otherwise there is an array of ids and
//quants at top of function.  A for loop loops through all the ids and creates displayed
//records to be displayed after the execution at end.

//this function is used when search is pressed in the index.html
type geoData struct {
	Var   []string
	Id    []int
	Quant []int
	UID   int `json:"a4"`
}

type try1 struct {
}

//type geoData[4]

type display5 struct {
	Var   string `json:"var"`
	Id    string `json:"id"`
	Quant string `json:"quant"`
	Uid   string `json:"uid"`
}

type Display3 struct {
	Var int `json:"var"`
}

func display1(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "application/json")

	var myArray display5

	c := json.NewDecoder(r.Body)
	var err80 = c.Decode(&myArray)
	if err80 != nil {
		fmt.Println(err80)
		fmt.Println(c)
		fmt.Println(myArray)

	}

	//err78 := json.Unmarshal([]byte(c), &myArray)
	//if err78 != nil {
	//	fmt.Println(c)
//
//	}

	/*
	//req := endpoint.AddRequest{}

	//req := endpoint.AddRequest{}

	//req, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))

	//var unknown map[string]interface{}

	var myArray []display5

	b, err77 := ioutil.ReadAll(r.Body)
	var c = string(b)
	fmt.Println(c)

	//r.ParseForm()
	//err79 := json.NewDecoder(r.Body).Decode(&myArray)
	//if err79 != nil {
	//	fmt.Print(err79)
	//	fmt.Println(c)
	//
	//}

	err78 := json.Unmarshal([]byte(c), &myArray)
	if err78 != nil {
		fmt.Println(c)

	}

	fmt.Println(c)
	fmt.Println(myArray)

	for index, element := range myArray {
		fmt.Println(index, "=>", element)
	}

	//unknown[]
	//	fmt.Println(unknown[0])

	fmt.Println(b)
	fmt.Println(err77)

	c = string(b)
	fmt.Println(string(b))
	fmt.Println(c)
*/
	r.ParseForm()
	var body1 string
	for key, _ := range r.Form {
		body1 = string(key)
		fmt.Println(body1)
		fmt.Println(string(body1))
		break
	}

	array := r.Form["var"]
	fmt.Println(array)

	var a = r.FormValue("var")
	fmt.Println(a)
	queries := r.URL.Query()
	fmt.Println(queries)

	//var display2 Display2
	//r.ParseForm()
	//err3 := json.NewDecoder(r.Body).Decode(&display2)
	//if err3 != nil {
	//	fmt.Println(err3)
	//	fmt.Println(r.Body)
	//	//fmt.Println(r.)
	//
	//	}

	//val12, err100 := r.GetBody()

	//fmt.Println(val12)
	//fmt.Println(err100)

	////////1591
	/*r.ParseForm()
	userid := r.Form["Var"]
	fmt.Println(userid)

	fmt.Println(r.Body)
	//var p display__
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	//err9 := dec.Decode(&p)
	if err9 != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err9, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err9, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(w, msg, http.StatusBadRequest)

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err9, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err9.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err9.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err9, io.EOF):
			msg := "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err9.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			log.Println(err9.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	//err3 := json.NewDecoder(r.Body).Decode(&p)
	//if err3 != nil {
	//	fmt.Println(err3)
	//	fmt.Println(r.Body)
	//	//fmt.Println(r.)
	//
	//}

	//err5 := r.ParseForm()
	//if err5 != nil {
	//	fmt.Println(err5)
	//}

	//userid := r.Form["Id"][0]
	//fmt.Println(userid)

	//val3 := r.FormValue("mydata")
	//fmt.Println(val3)

	var result interface{}
	err7 := json.NewDecoder(r.Body).Decode(&result)
	if err7 != nil {

	}
	//err5 := r.ParseForm()
	//if err5 != nil {
	//	fmt.Println(err5)
	//}
	//userid := r.Form["var"][0]
	//fmt.Println(userid)

	var j = `{"var":"test","id":10,"quant":1,"uid":1}`

	var it display_

	val11 := json.Unmarshal([]byte(j), &it)
	fmt.Println(val11)

	if err7 != nil {
		panic(err7)
	}
	//fmt.Println(it)

	//	body, err6 := ioutil.ReadAll(r.Body)
	//if err6 != nil {
	//    panic(err6)
	//}
	//
	//fmt.Println(body)

	//fmt.Println(key)

	//	var p display_

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	//  err3 := json.NewDecoder(r.Body).Decode(&p)
	//	if err3 != nil {
	//		fmt.Println(w, err3)
	//	}

	//fmt.Println(p)

	err4 := r.ParseForm()
	if err4 != nil {
		fmt.Println(w, err4)
	}

	//query := r.URL.Query()

	//r.Body
	//userid, present := query["userid"]
	//userid := r.Form["Var"][0]
	//userid = r.Form["Id"][1]
	//userid = r.Form["Id"][2]

	//for key, val := range userid {
	///
	//		fmt.Println(key)
	//		fmt.Println(val)
	//	}
	//
	//	fmt.Println(userid)

	/////////////
	//queries := r.URL.Query()
	fmt.Println(queries)
	val2 := ""
	for key, val := range queries {

		fmt.Println(key)
		fmt.Println(val)

		val2 = val[0]
		fmt.Println(val2)

	}

	//var biArray [][]string

	//json.Unmarshal([]byte(` [["keyword1","keyword2"],["1","2"],["2","2"],["2"]]`), &biArray)
	//"[{\"var\":\"test\",\"id\":10,\"quant\":1,\"uid\":1},{\"var\":\"test2\",\"id\":100,\"quant\":2,\"uid\":2}]"
	//b := json.Unmarshal([]byte(`[{\"var\":\"test\",\"id\":10,\"quant\":1,\"uid\":1},{\"var\":\"test2\",\"id\":100,\"quant\":2,\"uid\":2}]`), &biArray )
	//c := json.Unmarshal([]byte(`[{"var":"test","id":10,"quant":1,"uid":1},{"var":"test2","id":100,"quant":2,"uid":2}]`), &biArray)
	//d := json.Unmarshal([]byte(val2), &biArray)

	//for index, element := range d {
	//	fmt.Println(index, "=>", element)
	//	//fmt.Println(d)
	//	//fmt.Println(c)

	//}

	//ARRAY OF INTS  [3,4,7]
	//thesse ints are kept in database and changed to an array to see if they have aleady been
	//displayed so continue.  Does not effect the client side is an int
	//product is checked with array and if exists is contniues

	GlobCounter++
	//store, err := session.Start(context.Background(), w, r)
	*/
	/*
		err = r.ParseForm()

		array := r.Form["var"][1]
		fmt.Println(array)

		value := r.Form["var"]

		a := value[0]
	*/

	query := r.URL.Query()

	//this is the searchterm in order from first to last now
	key1, present := query["var"]

	if !present || len(key1) == 0 {
		fmt.Println("filters not present1")
	}

	keyTotalAmountBought, present2 := query["quant"]
	if !present2 || len(keyTotalAmountBought) == 0 {
		fmt.Println("filters not present2")
	}
	ProdID, present3 := query["id"]
	if !present3 || len(ProdID) == 0 {
		fmt.Println("filters not present3")

	}

	UserIDstring, present4 := query["uid"]
	if !present4 || len(UserIDstring) == 0 {
		fmt.Println("filters not present4")

	}

	var val1 = ""
	val1 = UserIDstring[0]
	//var err1 = ""
	var UserID int
	var err error
	if len(UserIDstring) != 0 {

		//only one
		UserID, err = strconv.Atoi(val1)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		UserID = 1
	}

	globKeyword := key1[0]

	w.Header().Set("Access-Control-Allow-Origin", "*")

	string1 = ""

	fmt.Println("in display 1")

	db := dbConn()

	//////

	var numRecords = 0
	stmt, err := db.Prepare("SELECT COUNT(*) FROM products WHERE ((products.ProductKeyWord1 = ?) OR (products.ProductKeyWord2 = ?) OR " +
		"(products.ProductKeyWord3 = ? )) AND products.ProductStatus = 'ready'")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := stmt.Query(globKeyword, globKeyword, globKeyword)

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {

		err = rows.Scan(&numRecords)

		if err != nil {
			fmt.Println(err)
		}

	}

	//////

	stmt, err = db.Prepare("SELECT products.ProductKeyword1, products.ProductKeyword2, products.ProductKeyword3, products.ProductName, products.ProductID, " +
		"products.ProductDescription, products.ProductCost, products.ProductQuantity, products.ProductCatTitle , products.ProductFilename " +
		"FROM products WHERE " +
		"((products.ProductKeyWord1 = ?) OR " +
		"(products.ProductKeyWord2 = ?) OR (products.ProductKeyWord3 = ? )) AND products.ProductStatus = 'ready'")
	if err != nil {
		//	//panic(err.Error())
	}

	rows, err = stmt.Query(globKeyword, globKeyword, globKeyword)

	if err != nil {
		fmt.Fprint(w, err)
	}

	var templ1 forTemplate

	var Link = globKeyword

	var Condition = 0
	//saved text product ids :  1,11,5,7
	var ints []int
	//var keywords []string

	var marshalFlag = "no"

	//var lastProductID = -1
	//counter1 = -1

	//counter1 = 0
	var counterOfRecords = 0
	for rows.Next() {

		counterOfRecords++
		//counter1++

		marshalFlag = "no"
		//counter1 = counter1 + 1
		Condition++
		var ProductCost float64
		var ProductQuantity, CondYellow int
		var gKeyword1, gKeyword2, gKeyword3, ProductName, ProductDescription, ProductCatTitle, ProductFilename, AmountToPurchaseID, AmountPurchasedID string

		CondYellow = 0
		err = rows.Scan(&gKeyword1, &gKeyword2, &gKeyword3, &ProductName, &ProductID, &ProductDescription, &ProductCost, &ProductQuantity, &ProductCatTitle, &ProductFilename)

		//if ProductID == lastProductID {
		//	continue
		//}

		//lastProductID = ProductID

		if err != nil {
			fmt.Fprint(w, err)
		}

		db3 := dbConn()
		//get from dbase

		textstring := ""
		//selects all the product ids with quantity
		stmt1, err := db3.Prepare("SELECT savedtext.Text FROM savedtext WHERE savedtext.UserID = ?")

		if err != nil {
			fmt.Println(err)
		}

		rows1, err := stmt1.Query(UserID)

		if err != nil {
			fmt.Println(err)
		}

		//get string from database
		for rows1.Next() {
			marshalFlag = "yes"
			err = rows1.Scan(&textstring)
			if err != nil {
				fmt.Println(err)
			}

		}

		//change string to array

		if marshalFlag == "yes" && textstring != "" {
			err = json.Unmarshal([]byte(textstring), &ints)
			if err != nil {
				fmt.Println(err)
			}
		}

		//check for duplicates, that is if productID already has been displayed don't display again
		var flag1 = 0
		var j = 0
		for j = 0; j < len(ints); j++ {
			if ProductID == ints[j] {
				flag1 = 1
				break
			}

		}
		if flag1 == 1 {

			continue
		}

		//creates and sets records :  K1, K2...THIS IS FOR LOOKING AT GLOB WORD AND IS NOW ONLY ONE ELEMENT READ BELOW
		/*var index = "K" + strconv.Itoa(GlobCounter)
		store.Set(index, globKeyword)
		err = store.Save()
		if err != nil {
			fmt.Fprint(w, err)
			//	return
		}
		*/

		var stringText = ""

		//DOES THIS PRODUCT RECORD ALREADY EXIST
		stmt1, err = db3.Prepare("SELECT savedtext.Text FROM savedtext WHERE savedtext.UserID = ?")

		if err != nil {
			fmt.Println(err)
		}

		rows1, err = stmt1.Query(UserID)

		if err != nil {
			fmt.Println(err)
		}

		var flag = 0

		//get string from database - is at least one record
		for rows1.Next() {
			flag = 1
			err = rows1.Scan(&stringText)
			if err != nil {
				fmt.Println(err)
			}

			//change string to array
			err := json.Unmarshal([]byte(stringText), &ints)
			if err != nil {
				fmt.Println(err)
			}

			//push to array
			ints = append(ints, ProductID)

		}

		//no database entry yet, so insert
		if flag == 0 {

			//pass in array and get string back
			//var textstring, err = json.Marshal(ints)

			stmt2, err := db3.Prepare("INSERT INTO savedtext(Text) VALUES(?)")
			if err != nil {
				fmt.Println(err)
			}

			ints = append(ints, ProductID)

			var textstring, err1 = json.Marshal(ints)
			if err1 != nil {
				fmt.Println(err1)
			}

			stmt2.Exec(textstring)

			//there is/are database entries, so update
		} else {

			var textstring, err1 = json.Marshal(ints)
			if err1 != nil {
				fmt.Println(err)
			}

			//update string

			stmt1, err := db3.Prepare("UPDATE savedtext SET Text=? WHERE UserID=?")
			if err != nil {
				fmt.Println(err)
			}
			stmt1.Exec(textstring, UserID)

		}

		/////////

		////////

		stmt1, err = db3.Prepare("SELECT savedtext.Text FROM savedtext WHERE savedText.UserID = ?")

		if err != nil {
			fmt.Println(err)
		}

		rows1, err = stmt1.Query(UserID)

		if err != nil {
			fmt.Println(err)
		}

		//get string from database
		for rows1.Next() {

			err = rows1.Scan(&stringText)
			if err != nil {
				fmt.Println(err)
			}

			//change string to array
			err := json.Unmarshal([]byte(stringText), &ints) //
			if err != nil {
				fmt.Println(err)
			}

		}

		i := 0
		prodBoughtInt := 0

		AmountPurchased := 0

		var flagProductIDHasBeenTemplated = false
		//is records with product amounts already
		for i = 0; i < len(ProdID); i++ {
			prodIDStr := ProdID[i]

			prodIDInt, err := strconv.Atoi(prodIDStr)
			if err != nil {
				fmt.Println(err)
			}

			prodBoughtStr := keyTotalAmountBought[i]
			prodBoughtInt, err = strconv.Atoi(prodBoughtStr)
			if err != nil {
				fmt.Println(err)

			}

			productIDInt, err := strconv.Atoi(ProdID[i])

			//write over with values, already existed with url parameters
			if ProductID == productIDInt {

				flagProductIDHasBeenTemplated = true
				counter1++

				str := strconv.Itoa(counter1)

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

				AmountPurchased = prodBoughtInt

				templ1 = forTemplate{CondYellow, Link, Condition, AmountPurchased, prodIDInt, ProductCatTitle, titleID, ProductName, descID, ProductDescription, costID, ProductCost, quantityID, ProductQuantity,
					key1ID, gKeyword1, key2ID, gKeyword2, key3ID, gKeyword3, ProductFilename, AmountToPurchaseID, AmountPurchasedID, mainDivID}

				fmt.Println(templ1)

				globt = template.Must(template.ParseFiles("C:/wamp64/www/golangproj/template1.html"))

				var err1 = globt.Execute(w, templ1)

				if err1 != nil {
					fmt.Println("---------------")
					fmt.Println(err.Error())
				}

				//create a displayed record wiht zero amountPurchased

				//create a zero amount new productID record

				sendToTemplate(&globKeyword, &counter1, &w, &CondYellow, &Link, &Condition, &AmountPurchased, &ProductID, &ProductCatTitle, &ProductName, &ProductDescription, &ProductCost, &ProductQuantity,
					&gKeyword1, &gKeyword2, &gKeyword3, &ProductFilename)

				break
			}

		}
		if !flagProductIDHasBeenTemplated {

			AmountPurchased = 0

			sendToTemplate(&globKeyword, &counter1, &w, &CondYellow, &Link, &Condition, &AmountPurchased, &ProductID, &ProductCatTitle, &ProductName, &ProductDescription, &ProductCost, &ProductQuantity,
				&gKeyword1, &gKeyword2, &gKeyword3, &ProductFilename)

		}

	}
}

//////////

/////////////

//send from client to server and
//send form server to client
//this is a good example
func getMessages(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	type User struct {
		Name string
		Age  int
		City string
	}

	//	a := User{Name:"a" , Age: 10 , City:"s" };

	var user = []User{{

		Name: "John Doe",
		Age:  10,
		City: "richmond"}}

	var msg = new(User)
	msg.Name = "Test namee"
	msg.Age = 30
	msg.City = "here"
	user = append(user, *msg)

	msg = new(User)
	msg.Name = "namee"
	msg.Age = 20
	msg.City = "here2"
	user = append(user, *msg)

	json.NewEncoder(w).Encode(user)

	//w.Header().Set("Content-Type", "application/json")
	//w.Write(j)
	fmt.Println("--wwww---")
	fmt.Println(user)
}

////////////

/////////////

func main() {

	one := http.NewServeMux()

	//mux := http.NewServeMux()

	//has an id value passed in url
	one.HandleFunc("/updateForm/", updateForm)

	//button3 - just read session for right now
	one.HandleFunc("/getMessages", getMessages)

	one.HandleFunc("/display", display1)

	one.HandleFunc("/display2", display2)

	one.HandleFunc("/HelloWorld", HelloWorld)
	one.HandleFunc("/processLogin", processLogin)

	//two := http.NewServeMux()

	//
	one.HandleFunc("/template2", createTemplate2)
	one.HandleFunc("/spitBackAmounts", spitBackAmounts)

	//go func() {
	//
	//		http.ListenAndServe(":8080", one)
	//	}()

	http.ListenAndServe(":8080", one)

}
