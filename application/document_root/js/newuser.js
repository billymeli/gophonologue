function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(',');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function setCookie(cname, cvalue, time) {
    var d = new Date();
    d.setTime(d.getTime() + time);
    var expires = "expires="+ d.toUTCString();
    document.cookie = cname + "=" + cvalue + "," + expires + ",path=/";
}

jQuery('._js_new-user-form').on('submit', function(e){
   e.preventDefault();

   var newuser = jQuery(this).find('input[name="screen-name"]').val();
   setCookie("username", newuser, 60*60*2);
   jQuery('.new-user-modal').modal('hide');
});

var user = getCookie("username");
if (user == "") {
   jQuery('.new-user-modal').modal({
      keyboard: false,
      backdrop:'static'
   });
}
