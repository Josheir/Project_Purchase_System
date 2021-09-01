

//the ajax returns a string which has the name of the keyword that needs to be removed and both alterd datas are used	
function deriveKeywordFromResult(result)
{

				
				var item = "";
				var item1 = "";
				var ourNewWord = "";
				item = result.match("\\[.*]");
				ourNewWord = result.replaceAll(item, '');
				alert(ourNewWord);
				var x = typeof(item);
				item1 = JSON.stringify(item);
				item1 = item1.slice(5);
				item1 = item1.substring(0, item1.length - 5);
				

return [item1, ourNewWord, item];				
}
