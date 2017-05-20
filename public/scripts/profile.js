$(document).ready(function(){
  $('.materialboxed').materialbox();
});

$(document).ready(function(){
  $('ul.tabs').tabs();
});

$(document).ready(function(){
  $('#btnHomePost').on('click',function(){
  $('#home-feeds').prepend('<li>Kemzi</li>')
});
});

$('.datepicker').pickadate({
  selectMonths: true, // Creates a dropdown to control month
  selectYears: 15, // Creates a dropdown of 15 years to control year
  format: 'yyyy-mm-dd'
});



