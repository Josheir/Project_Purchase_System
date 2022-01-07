/*
			//template number one presses this and changes values in quantity and amount purchased
			function submitChangesOnOrderTemplates(mainDiv, originalQuantity, productID, quantityID, amountPurchasedID, amountToPurchaseID)
			{
				

				
				alert("zzz");
				//alert(quantityID);

				var makeYellow = document.getElementById(mainDiv);
				makeYellow.style["background-color"] = "yellow";
				var amountToBePurchased  = 0;
				var  amountInPurchasedField = 0;
				var amountPurchasedSoFar = 0;
				var value1 = 0;
				var value2 = 0;
				//total of product with server
				amountInPurchasedField = document.getElementById(quantityID).value;
				//already purchased
				amountPurchasedSoFar = document.getElementById(amountPurchasedID).value;
				//alert("in the amount purchased");
				//alert(amountPurchasedSoFar);
				
				//amount held for subtraction when submit button pressed
				amountToBePurchased = document.getElementById(amountToPurchaseID).value;
				//alert(amountToBePurchased);
				
				//alert(amountToBePurchased);
				//document.getElementById(amountPurchasedID).value="";

				
				//if there is enough product go in here - take from quantity amount
				if((parseInt(amountInPurchasedField) >= parseInt(amountToBePurchased))   )//&& parseInt(amountToBePurchased) > 0  )
				{

					
					//
					value1 = document.getElementById(amountToPurchaseID).value; 
					//value2 = amountopurchase + amountpurchaseedsofar
					value2  = parseInt(value1) + parseInt(amountPurchasedSoFar);
					if(value2 < 0)
					{
						document.getElementById(amountToPurchaseID).value = 0;
						return;
					}
					//alert("values");
					//alert(value1);
					//alert(value2);
					//alert(amountPurchasedSoFar)
					
					//puts total in here
					document.getElementById(amountPurchasedID).value = value2 ;
					
					if(value1 > 0){

						//alert("a");
					document.getElementById(quantityID).value = parseInt(amountInPurchasedField) - parseInt(amountToBePurchased) ;
					}else if (value1 == 0)
					{
						//alert("b");
					}

					//negative value for amount to purchase
					//value
					else if ( (parseInt(amountPurchasedSoFar) + parseInt(value1)) > -1){
						//alert("in neg");
						//alert(amountInPurchasedField);
						//alert( amountToBePurchased);
						document.getElementById(quantityID).value = parseInt(amountInPurchasedField) - parseInt(amountToBePurchased) ;
						//
						document.getElementById(amountToPurchaseID).value = 0;
					}


					//alert(amountToBePurchased);
					
					gTotalOfPurchases =  parseInt(gTotalOfPurchases) +  parseInt(amountToBePurchased); 
					
					//alert(amountToBePurchased);
					
					if(gTotalOfPurchases > 0){
                        document.getElementById(amountToPurchaseID).value = 0;
                        //alert("here");
						//alert(amountToBePurchased);
						
						//so far, so good is still in database!
						//SetDataFromOrderTemplateOnSubmit(originalQuantity, productID, amountToBePurchased, 0)
						SetDataFromOrderTemplateOnSubmit(amountPurchasedSoFar, productID, amountToBePurchased, 0)



                        return (1);
                    }

					///////////


				

					///////////



				}
      }


      //called from checkout button to setitem for template2 to use
      function  SetDataFromOrderTemplateOnSubmit(quantityOfProduct, productID, quantityTryingToSet, reduceBool)
			{
				//if(firstTime == "yes")
				//{
				//	//sessionStorage.clear();
				//	firstTime = "no";
				//}
				
				var string = "";
				
				var keyFlag = ""
				var bought, totalquantity = 0;

				//alert("before");




				alert("000");
				let keys = Object.keys(sessionStorage);
				alert(keys);
				//console.log(keys);
				for(let key of keys) {
					//alert("key");
					//alert(key)
					//alert("prodcutid")
					//alert(productID)
				
					alert("here1");
					alert(key);
				//total the quantity and overwrite - is already a record
				if(key == (productID.toString())){
					alert("here2");
					alert(key);
				
					//alert("in this place");
					//alert(productID);
					//alert("SetDataFromOrderTemplateOnSubmit");
					keyFlag = "found";
					//sofar
					//alert(sessionStorage.getItem(key));
					
					//bought
					bought = sessionStorage.getItem(key);
					//alert("bought");
					//alert(bought);
					//document.getElementById(amountPurchasedID).value = bought;
					//alert(quantityOfProduct);
					//quantity of records is decreasing
					var totalquantity = 0;
					
					//if (reduceBool)
					//{
					//	//alert("1a");
					//	totalquantity =  parseInt(bought) -  parseInt(quantityTryingToSet);
					//	//alert(quantityOfProduct);
					//}
					//quantity of records is acummulating
					//else 
					{
						//alert("1b");
						//alert(quantityOfProduct);
						//alert(quantityTryingToSet);
						//alert(totalquantity);
						//alert(quantityTryingToSet);
						totalquantity = parseInt(bought) + parseInt(quantityTryingToSet);
						//console.log(document.getElementById(amountPurchasedID).value);
						//document.getElementById(amountPurchasedID).value = 579;//'//totalquantity;
						//alert("quantityOfProduct");
						//alert(quantityOfProduct);
						//alert("quantityTryingToSet");
						//alert(quantityTryingToSet);
						//alert("totalquantity");
						//alert(totalquantity);
						
					}
					
					
					//overwrites
					sessionStorage.removeItem(productID);
					//alert("set");
					//alert(productID);
					//alert(totalquantity);
					
					


					//sessionStorage.setItem(productID, totalquantity);
					sessionStorage.setItem(productID, parseInt(quantityTryingToSet));
					alert("set1");
					alert(quantityTryingToSet);

					//alert("overwrites")
					break;
				
					}
				}
				//there is no productID record yet, make one here 
				if (keyFlag !=  "found")
				{
					



					//sessionStorage.setItem( parseInt(productID),  parseInt(quantityTryingToSet));
					sessionStorage.setItem( parseInt(productID),  parseInt(quantityTryingToSet));
					
					alert("set2");
					alert(quantityTryingToSet);

				}
				
                

			}*/