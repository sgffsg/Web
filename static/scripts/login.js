const showPasswordButton = document.querySelector(".form-field__show-password");
const hidePasswordButton = document.querySelector(".form-field__hide-password");
const logInButton = document.querySelector(".form__button");

showPasswordButton.addEventListener(
    "click",
    () => {
        showPasswordButton.classList.add("form-field__show-password-replace");
        hidePasswordButton.classList.add("form-field__hide-password-replace");
        document.querySelector(".form-field__password").type = "text";
    }
)

hidePasswordButton.addEventListener(
    "click",
    () => {
        showPasswordButton.classList.remove("form-field__show-password-replace");
        hidePasswordButton.classList.remove("form-field__hide-password-replace");
        document.querySelector(".form-field__password").type = "password";
    }
)

logInButton.addEventListener(
    "click",
    () => {
        if (!formValidate())
            createWarning("critical", "A-Ah! Check all fields,");
        else
            console.log("Success");
    }    
)



function formValidate()
{
    isEmailValid();
    isPasswordValid();
    if (!isEmailValid() || !isPasswordValid())
        return false;
    return true;
}

function isEmailValid()
{
    let emailInput = document.getElementById("email");
    let emailText = emailInput.value;

    let dog = emailText.indexOf("@");
    let point = emailText.indexOf(".");
    let arr = emailText.split("");
    let currLength = arr.length;
    let minLength = 6;

    let result = true;

    if (emailText == "") 
    {
        result = false;
    }
    else
    {
        if (currLength <= minLength)
        {
            result = false;
        }
        else
        {
            if (point == -1) 
            {
                result = false;
            }
            if (dog == -1) 
            {
                result = false;
            }
            if (emailText.indexOf(",")>=0 || emailText.indexOf(";")>=0 || emailText.indexOf(" ")>=0)
            {
                result = false;
            }
            if (dog>point)
            {
                result = false;
            }
            if ((point-dog)==1 )  
            {
                result = false;
            }
            if ((dog<1)) 
            {
                result = false;
            }
            if (!(point < emailText.length - 2)) 
            {
                result = false;
            }
        }
    }
            
    if (result == false)
    {
        removeError(emailInput);
        createError(emailInput, "Incorrect email format. Correct format is ****@**.***");
    }
    else
        return true;
}

function isPasswordValid()
{
    let passwordInput = document.getElementById("pass");

    if (passwordInput.value == "")
    {
        removeError(passwordInput);
        createError(passwordInput, "Password is required.");
        return false;
    }
    return true;
}

function createError(input, text)
{
    const parent = input.parentNode;
    const errorLabel = document.createElement('div');

    input.classList.add("form-field__error");
    errorLabel.classList.add("field__error");
    errorLabel.textContent = text;
    
    parent.classList.add('error');
    parent.appendChild(errorLabel);
}

function removeError(input)
{
    var parent = input.parentNode;
    
    
    if (parent.classList.contains('error'))
    {
        input.classList.remove("form-field__error");   
        parent.querySelector('.field__error').remove();
        parent.classList.remove('error');
    }
}

function createWarning(warningType, text)
{
    let warning = document.querySelector(".form__warning");
    let warningIcon = document.getElementById("warning-icon");
    let warningText = document.getElementById("warning-text");

    warning.classList.remove("form__warning--hidden");

    switch(warningType)
    {
        case "success":
            warning.classList.add('form__warning--success');
            
            break;
        case "critical":
            warning.classList.add('form__warning--critical');
            
            break;
    }

    warningText.textContent = text;
}

function removeWarning()
{
    let warning = document.querySelector(".form__warning");
    let warningIcon = document.getElementById("warning-icon");
    let warningText = document.getElementById("warning-text");

    warning.classList.add("form__warning--hidden");
    warning.classList.remove('form__warning--critical');
    warning.classList.remove('form__warning--success');
}