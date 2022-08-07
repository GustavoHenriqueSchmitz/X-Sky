// Function to make the code sleep in milisseconds
function Sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

/*========================================================
                    Main Function
========================================================*/
(function ($) {
    'use strict';
    try {
        $('.js-datepicker').daterangepicker({
            "singleDatePicker": true,
            "showDropdowns": true,
            "autoUpdateInput": false,
            locale: {
                format: 'DD/MM/YYYY'
            },
        });
    
        var myCalendar = $('.js-datepicker');
        var isClick = 0;
    
        $(window).on('click',function(){
            isClick = 0;
        });
    
        $(myCalendar).on('apply.daterangepicker',function(ev, picker){
            isClick = 0;
            $(this).val(picker.startDate.format('DD/MM/YYYY'));
    
        });
    
        $('.js-btn-calendar').on('click',function(e){
            e.stopPropagation();
    
            if(isClick === 1) isClick = 0;
            else if(isClick === 0) isClick = 1;
    
            if (isClick === 1) {
                myCalendar.focus();
            }
        });
    
        $(myCalendar).on('click',function(e){
            e.stopPropagation();
            isClick = 1;
        });
    
        $('.daterangepicker').on('click',function(e){
            e.stopPropagation();
        });
    
    
    } catch(er) {console.log(er);}
    
    try {
        var selectSimple = $('.js-select-simple');
    
        selectSimple.each(function () {
            var that = $(this);
            var selectBox = that.find('select');
            var selectDropdown = that.find('.select-dropdown');
            selectBox.select2({
                dropdownParent: selectDropdown
            });
        });
    
    } catch (err) {
        console.log(err);
    }
})(jQuery);

/*========================================================
                    Submit Form
========================================================*/
async function postFormFieldsAsJson({url, formData}) {
  //Create an object from the form data entries
  let formDataObject = Object.fromEntries(formData.entries());
  //Enconde image in base64
  var reader = new FileReader()
  reader.readAsDataURL(formDataObject.perfil_photo)
  reader.onload = function () {
    formDataObject.perfil_photo = reader.result;
  };
  reader.onerror = function (error) {
    console.log('Error: ', error);
  };
  // Sleep function to wait for the above functions to be executed
  await Sleep(500)
  // Format the plain form data as JSON
  let formDataJsonString = JSON.stringify(formDataObject);  // Set the fetch options (headers, body)
  let fetchOptions = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: formDataJsonString,
  };

  //Get the response body as JSON.
  //If the response was not OK, throw an error.
  let res = await fetch(url, fetchOptions);

  //If the response is not ok throw an error (for debugging)
  if (!res.ok) {
    let error = await res.text();
    throw new Error(error);
  }
  //If the response was OK, return the response body.
  window.location.replace("http://localhost:3000/login")
  return res.json();
}

// Get the form element by id
const sampleForm = document.getElementById("signup-form");

// Add an event listener to the form element and handler for the submit an event.
sampleForm.addEventListener("submit", async (e) => {
  /*Prevent the default browser behaviour of submitting
   the form so that you can handle this instead.*/
  e.preventDefault();

  // Get the element attached to the event handler.
  let form = e.currentTarget;
  // Take the URL from the form's `action` attribute.
  let url = form.action;

  try {
    /*Takes all the form fields and make the field values
    available through a `FormData` instance.*/
    let formData = new FormData(form);
    formData.set.name.perfil
    
    // The `postFormFieldsAsJson()` function in the next step.
    let responseData = await postFormFieldsAsJson({ url, formData });

    //Destructure the response data
    let { serverDataResponse } = responseData;

    //Display the response data in the console (for debugging)
    console.log(serverDataResponse);
  } catch (error) {
    //If an error occurs display it in the console (for debugging)
    console.error(error);
  }
});

/*========================================================
                    Select Button
========================================================*/
$("#selectfile").bind('change',function(){
    var filename = $("#selectfile").val();
    if(/^s*$/.test(filename)){
        $("#blankFile").text("No File Chosen..");
        $("#success").hide();
    }else{
        $("#blankFile").text(filename.replace("C:\\fakepath\\",""));
        $("#success").show();
    }
})
