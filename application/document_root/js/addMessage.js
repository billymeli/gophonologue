
function formatAMPM(date) {
    var hours = date.getHours();
    var minutes = date.getMinutes();
    var ampm = hours >= 12 ? 'PM' : 'AM';
    hours = hours % 12;
    hours = hours ? hours : 12; // the hour '0' should be '12'
    minutes = minutes < 10 ? '0'+minutes : minutes;
    var strTime = hours + ':' + minutes + ' ' + ampm;
    return strTime;
}

//-- No use time. It is a javaScript effect.
function insertChat(who, text, time = 0){
    var control = "";
    var date = formatAMPM(new Date());

    if (who == "me"){

        control = '<li style="width:85%">' +
                        '<div class="msj macro">' +
                        '<div class="avatar"></div>' +
                            '<div class="text text-l">' +
                                '<p>'+ text +'</p>' +
                                '<p style="color:#424141"><small>'+date+'</small></p>' +
                            '</div>' +
                        '</div>' +
                    '</li>';
    }else{
        control = '<li style="width:85%;">' +
                        '<div class="msj-rta macro">' +
                            '<div class="text text-r">' +
                                '<p>'+text+'</p>' +
                                '<p style="color:#424141"><small>'+date+'</small></p>' +
                            '</div>' +
                        '<div class="avatar" style="padding:0px 0px 0px 10px !important"></div>' +
                  '</li>';
    }
    setTimeout(
        function(){
            $("#messageBox").append(control);

        }, time);

}

function resetChat(){
    $("#messageBox").empty();
}

//-- Clear Chat
resetChat();

//-- Print Messages
insertChat("me", "Hello Ennovar!", 0);
insertChat("others", "How you doing?", 1500);
insertChat("me", "How is your weekend going so far?", 3500);
insertChat("others", "Pretty good",7000);
insertChat("me", "How is the coding challenge goin on?", 9500);
insertChat("others", "LOL", 12000);


//-- NOTE: No use time on insertChat.
