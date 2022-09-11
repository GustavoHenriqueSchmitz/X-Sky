/*========================================================
                    Submit Form
========================================================*/
async function postFormFieldsAsJson({url, formData}) {

  //Create an object from the form data entries
  let formDataObject = Object.fromEntries(formData.entries());

  //Enconde image in base64
  var reader = new FileReader()
  reader.readAsDataURL(formDataObject.perfil_photo)
  reader.onload = () => {
    formDataObject.perfil_photo = reader.result;
  }
  reader.onerror = (error) => {
    console.log('Error: ', error);
  }

  // Sleep for 10ms.
  await new Promise(resolve => setTimeout(resolve, 100));

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

  let res = await fetch(url, fetchOptions);

  //Return the response.
  return res;
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
    let res = await postFormFieldsAsJson({ url, formData });

    if (res.status == 201) {
      window.location.replace("/front-end/templates/login/login.html")
    } else {
      window.location.replace("/front-end/templates/login/login.html")
    }

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
