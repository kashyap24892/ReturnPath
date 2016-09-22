$(document).ready(function(){
  $("#Country").change(function(){

    $("#State").empty();
    var n = $(this).val();
    $.post("/states",{Country:n},function(data){
      var states = data.toString().split(",");
      for(var i=0; i<states.length;i++){
        $("#State").append("<option value=\"" + states[i] + "\"> " + states[i] + "</option>");
      }
      // alert("Returned " + states );
    });
  });
  
  $("#State").change(function(){
    $("#City").empty();
    var n = $(this).val();
    $.post("/cities",{State:n},function(data){
      var cities = data.toString().split(",");
      for(var i=0; i<cities.length;i++){
        $("#City").append("<option value=\"" + cities[i] + "\"> " + cities[i] + "</option>");
      }
      //  alert("Returned " + data);
    });
  });
});
