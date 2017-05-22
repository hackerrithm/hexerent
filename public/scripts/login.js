console.log(6 * 6)


$(function() {
  var bool = true, flag = false;
  $('#btnLogin').prop('disabled', bool); // use prop to disable the button

  $(document).keyup(function() { // listen the keyup on the document or you can change to form in case if you have or you can try the closest div which contains the text inputs
    $('input:text').each(function() { // loop through each text inputs
      bool = $.trim(this.value) === "" ?  true :  false; // update the var bool with boolean values
      if(bool)
      return flag;
    });
    $('#btnLogin').prop('disabled', bool); // and apply the boolean here to enable
  });
});