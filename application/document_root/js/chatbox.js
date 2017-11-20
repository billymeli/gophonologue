function convertUnix(unix_timestamp) {
   var date = new Date(unix_timestamp*1000);
   var hours = date.getHours();
   var minutes = "0" + date.getMinutes();
   var seconds = "0" + date.getSeconds();

   var formattedTime = hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2);
   return formattedTime;
}

function insertChat(username, message, timestamp) {
    var control = "";
    var currentUser = getCookie("username");

   if (currentUser == username){
      control =
      '<li>' +
         '<div class="msj macro">' +
            '<div class="username mr-2">' + username + ' - ' + convertUnix(timestamp) + '</div>' +
            '<div class="text text-l">' +
               '<p>' + message + '</p>' +
            '</div>' +
         '</div>' +
      '</li>';
   }
   else {
      control =
      '<li>' +
         '<div class="msj-rta macro">' +
            '<div class="text text-r">' +
               '<p>' + message + '</p>' +
            '</div>' +
         '<div class="username mr-2">' + username + ' - ' + convertUnix(timestamp) + '</div>' +
      '</li>';
   }

   $("#messageBox").append(control);
   $('#messageBox').scrollTop($('#messageBox')[0].scrollHeight);
}

function resetChat() {
    $("#messageBox").empty();
}
