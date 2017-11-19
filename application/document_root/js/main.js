/* On startup get all old messages in db */
function refreshChat() {

   var xhr = new XMLHttpRequest();
   xhr.open('GET', 'api/messenger/get', true);
   xhr.send();

   xhr.onloadend = function() {
      if(xhr.status == 200) {
         if(xhr.response != undefined && xhr.response.length != 0) {
            resetChat();
            var jsonResponse = JSON.parse(xhr.responseText);
            var timestamp, username, message;

            $.each(jsonResponse, function(key, item) {
               timestamp = key;
               $.each(item, function(innerKey, innerItem) {
                  if (innerKey == "username") username = innerItem;
                  else if(innerKey == "message") message = innerItem;
               });
               insertChat(username, message, timestamp)
            });
         }
      }
      else {
         if(xhr.response != undefined && xhr.response.length != 0) {
            alert(xhr.response);
         }
         else {
            alert('An error has occurred, please contact your webhost administrator.');
         }
      }
   }
}

refreshChat();
var chatLoop = setInterval(refreshChat, 1500);
