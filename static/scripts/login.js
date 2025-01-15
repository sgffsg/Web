const showPasswordButton = document.querySelector(".form-field__show-password");
const hidePasswordButton = document.querySelector(".form-field__hide-password");

const emailInputField = document.querySelector(".form-field__email");
const passwordInputField = document.querySelector(".form-field__password");

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

emailInputField.addEventListener(
    "input",   
    () => {
        isValidEmail(emailInputField, emailInputField.value);
    }
)


passwordInputField.addEventListener(
    "input",   
    () => {
        formValidation();
    }
)

function formValidation()
{
    if (!isValidEmail(emailInputField, emailInputField.value))
        return false;
    if (!isValidPassword(passwordInputField))
        return false;

    return true;
}

function isValidEmail(input, text)
{
    var dog = text.indexOf("@");
    var point = text.indexOf(".");
    var arr = text.split('');
    var currLength = arr.length;
    var minLength = 6;
    
    
    if (text == "") 
    {
        removeError(input);
        createError(input, "Введен пустой адрес эл. почты");
        return false
    }
    
    if (currLength <= minLength)
    {
        removeError(input);
        createError(input, "Vинимальный размер адреса эл. почты – 6 символов");
        return false
    }
    else
    {
        if (point == -1) 
        {
            removeError(input);
            createError(input, "Отсутствует символ .");
            return false
        }
        if (dog == -1) 
        {
            removeError(input);
            createError(input, "Отсутствует символ @");
            return false
        }
        if (text.indexOf(",")>=0 || text.indexOf(";")>=0 || text.indexOf(" ")>=0)
        {
            removeError(input);
            createError(input, "В адресе эл.почты не может быть запятых,точек с запятой и пробелов.");
            return false
        }
        if (dog>point)
        {
            removeError(input);
            createError(input, "Cправа от символа «@» должна быть, как минимум, одна точка");
            return false
        }
        if ((point-dog)==1 )  
        {
            removeError(input);
            createError(input, "Между символом «@» и следующей за ним точкой должен быть, как минимум, один символ.");
            return false
        }
        if ((dog<1)) 
        {
            removeError(input);
            createError(input, "Слева от символа @, должен быть как минимум 1 символ.");
            return false
        }
          //5
        if (!(point < text.length - 2)) 
        {
            removeError(input);
            createError(input, "Справа от последней точки должно быть, как минимум, 2 символа.");
            return false
        }
    }
    removeError(input);
    return true;
}

function isValidPassword(input)
{
    var value = input.value;
    var result = true;
    if (value == "")
    {
        result = false;
        removeError(input);
        createError(input, input.name + " is required.");
    }
    return result;
}

function createError(input, text)
{
    const parent = input.parentNode;
    const errorLabel = document.createElement('div');

    errorLabel.classList.add('auth-form__input-description');
    errorLabel.textContent = text;

    parent.classList.add('error');
    parent.appendChild(errorLabel);
}


function removeError(input)
{
    var parent = input.parentNode;
    
    
    if (parent.classList.contains('error'))
    {   
        parent.querySelector('.auth-form__input-description').remove();
        parent.classList.remove('error');
    }
}