/*========================================================
                    Submit Form
========================================================*/
async function postFormFieldsAsJson({url, formData}) {

    // Create an object from the form data entries
    let formDataObject = Object.fromEntries(formData.entries());
    // Format the plain form data as JSON
    let formDataJsonString = JSON.stringify(formDataObject);
    // Set the fetch options (headers, body)
    let fetchOptions = {
            method: "POST",
            headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
        },
        body: formDataJsonString,
    };

    let res = await fetch(url, fetchOptions);
    // Return the server response.
    return res;
}

// Get the form element by id
const sampleForm = document.querySelector(".login-form")

// Add an event listener to the form element and handler for the submit an event.
sampleForm.addEventListener("submit", async (e) => {
/*Prevent the default browser behaviour of submitting
the form so that you can handle this instead.*/
e.preventDefault();

// Get the element attached to the event handler.
let form = e.currentTarget
// Take the URL from the form's `action` attribute.
let url = form.action

try {
    /*Takes all the form fields and make the field values
    available through a `FormData` instance.*/
    let formData = new FormData(form)
    formData.set.name.perfil
    
    // The `postFormFieldsAsJson()` function in the next step.
    let res = await postFormFieldsAsJson({ url, formData });

    if (res.status == 200) {
        window.location.replace("/")
    } else {
        window.location.reload()
    }
} catch (error) {
    //If an error occurs display it in the console (for debugging)
    console.error(error);
    }
});
