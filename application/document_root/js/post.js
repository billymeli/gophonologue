jQuery('._js_form-post').on('submit', function(e){
  e.preventDefault();

  var requestData = {};
  requestData["username"] = "george"
  requestData["message"] = "im a sneaky hacker? <script>alert('test');</script>"

  var xhr = new XMLHttpRequest();
  xhr.open('POST', 'api/messenger/post', true);
  xhr.send(JSON.stringify(requestData));

  xhr.onloadend = function() {
    if(xhr.status == 200) {
      if(xhr.response != undefined && xhr.response.length != 0) {
       console.log(xhr.response);
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
