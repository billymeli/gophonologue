jQuery('._js_form-post').on('submit', function(e){
  e.preventDefault();

  var requestData = {};
  requestData["username"] = getCookie("username");
  requestData["message"] = jQuery(this).find('input[name="message"]').val();

  jQuery(this).find('input[name="message"]').val('');

  var xhr = new XMLHttpRequest();
  xhr.open('POST', 'api/messenger/post', true);
  xhr.send(JSON.stringify(requestData));

  xhr.onloadend = function() {
    if(xhr.status == 200) {
      if(xhr.response != undefined && xhr.response.length != 0) {
         var jsonResponse = JSON.parse(xhr.responseText);
         var timestamp, username, message;

         $.each(jsonResponse, function(key, item) {
            timestamp = key;
            $.each(item, function(innerKey, innerItem) {
               if (innerKey == "username") username = innerItem;
               else if(innerKey == "message") message = innerItem;
            });
         });

         insertChat(username, message, timestamp)
      }
      else {
       alert('An error has occurred, please contact your webhost administrator.');
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
});
