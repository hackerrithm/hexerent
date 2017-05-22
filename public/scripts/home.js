// More or less revealer for long posts
jQuery(function(){

    var minimized_elements = $('p.minimize');
    
    minimized_elements.each(function(){    
        var t = $(this).text();        
        if(t.length < 100) return;
        
        $(this).html(
            t.slice(0,100)+'<span>... </span><a href="#" class="more">More</a>'+
            '<span style="display:none;">'+ t.slice(100,t.length)+' <a href="#" class="less">Less</a></span>'
        );
        
    }); 
    
    $('a.more', minimized_elements).click(function(event){
        event.preventDefault();
        $(this).hide().prev().hide();
        $(this).next().show();        
    });
    
    $('a.less', minimized_elements).click(function(event){
        event.preventDefault();
        $(this).parent().hide().prev().show().prev().show();    
    });

});

// toggle comment box
$(document).ready(function(){

            var value;
            var commentValue;
            var index = 0;
    
            var arr = jQuery.map(jQuery('.commentButtonToggler'),function(n,i){
                console.log("this is id:", i);
                return jQuery(n).attr('id');
            });
            console.log("this is list:", arr);


            var commentArr = jQuery.map(jQuery('.commentBoxToggler'), function (n, i) {
              console.log("this is comment id:", i)
              return jQuery(n).attr('id');
            })
            console.log("this is the commentBox list:", commentArr);

           
            Object.keys(arr).forEach(function(key) {
                value = arr[key];
                console.log("these are the values at: ", value);
            });

            Object.keys(commentArr).forEach(function(key) {
                commentValue = commentArr[key];
                console.log("these are the comment values at: ", commentValue);
            });

            Object.keys(arr, commentArr).forEach(function(key) {
            var id = "#".concat((arr[key]).toString())
            var commentId = "#".concat((commentArr[key]).toString())
                  
                  
                  /*$(id).off().on("click", function() {
                  console.log(arr[key], " pressed", "testing the value of: ", (arr[key]).toString())
                          $(id).off().on("click", function() {
                          $(commentId).toggle();
                        });
                        $(document).click(function(e) {
                          if ($(e.target).closest('commentId, id').length == 0) {
                            $(commentId).hide();
                          }
                        })
                });*/
                    $(commentId).hide();
                    $(id).click(function(){
                        $(commentId).toggle();
                    });
            });

            console.log(arr.length-1)
});



// maintain scroll position
$(window).scroll(function() {
  sessionStorage.scrollTop = $(this).scrollTop();
});

$(document).ready(function() {
  if (sessionStorage.scrollTop != "undefined") {
    $(window).scrollTop(sessionStorage.scrollTop);
  }
});

$('#btnHomePost').attr('disabled', true);
$('#homeFeedPost:text').keyup(function () {
   var disable = false;
       $('#homeFeedPost:text').each(function(){
            if($(this).val()==""){
                 disable = true;      
            }
       });
  $('#btnHomePost').prop('disabled', disable);
});


// left-panel sticks to top of page on scroll
$(window).scroll(function(e){ 
  var $el = $('#usefull-links'); 
  var isPositionFixed = ($el.css('position') == 'fixed');
  if ($(this).scrollTop() > 200 && !isPositionFixed){ 
    $('#usefull-links').css({'position': 'fixed', 'top': '40px'}); 
  }
  if ($(this).scrollTop() < 200 && isPositionFixed)
  {
    $('#usefull-links').css({'position': 'static', 'top': '50px'}); 
  } 
});