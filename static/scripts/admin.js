const title = document.getElementById('title');
const subTitle = document.getElementById('description');
const author = document.getElementById('author');
const publishDate = document.getElementById('publishDate');
const content = document.getElementById('content');

const uploadAuthorPhoto = document.querySelector(".author-photo-form__input");
const uploadTinyPostIMG = document.querySelector(".input-hero-image-tiny__input");
const uploadPostIMG = document.querySelector(".input-hero-image__input");

const removeAuthorPhoto = document.querySelector(".author-photo-form__remove-button");
const removeTinyPostIMG = document.querySelector(".remove-tiny-hero");
const removePostIMG = document.querySelector(".remove-hero");

const publishButton = document.querySelector(".form-head__button");

let postImgInput;
let authorIMGInput;
let previewIMGInput;

let postImgInputName;
let authorIMGInputName;
let previewIMGInputName;


title.addEventListener(
    "input" , 
    () => {
        let title = document.getElementById('title').value;
        let defaultTitle = 'New Post';
        if (title !== '' && title.length < 25) {
            document.getElementById('titleVisual').innerHTML = title;
            document.getElementById('titleVisualTiny').innerHTML = title;
        }
        else {
            document.getElementById('titleVisual').innerHTML = defaultTitle;
            document.getElementById('titleVisualTiny').innerHTML = defaultTitle;
        }
    }
)

subTitle.addEventListener(
    "input",  
    () => {
        let subtitle = document.getElementById('description').value;
        let defaultSubtitle = 'Please, enter any description';
        if (subtitle !== '' && subtitle.length < 60){
            document.getElementById('subtitleVisual').innerHTML = subtitle;
            document.getElementById('subtitleVisualTiny').innerHTML = subtitle;
        } 
        else {
            document.getElementById('subtitleVisual').innerHTML = defaultSubtitle;
            document.getElementById('subtitleVisualTiny').innerHTML = defaultSubtitle;
        }   
    }
)

author.addEventListener(
    "input",   
    () => {
        let name = document.getElementById('author').value;
        let defaultName = 'Enter author name';
        if (name !== '' && name.length < 25){
            document.getElementById('authorNameVisualTiny').innerHTML = name;
        }    
        else {
            document.getElementById('authorNameVisualTiny').innerHTML = defaultName;
        }
    }
)

publishDate.addEventListener(
    "input", 
    () => {
        let date = document.getElementById('publishDate').value;
        let defaultDate = '18.04.2024';
        if (date !== ''){
            document.getElementById('dateVisuality').innerHTML = date;
        }    
        else {
            document.getElementById('dateVisuality').innerHTML = defaultDate;
        }
    }
)

uploadAuthorPhoto.addEventListener(
    "input",  
    () => {
        const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
        const previewInput = document.querySelector(".preview-author-photo");
        const file = document.querySelector(".author-photo-form__input").files[0];
        const reader = new FileReader();
        reader.addEventListener(
            "load",
            () => {
                previewPostCardAuthorPhoto.src = reader.result;
                previewInput.src = reader.result;
                authorIMGInput = reader.result;
                authorIMGInput = file.name;
            },
            false
        );
        document.querySelector(".author-photo-form__remove-button").classList.add("remove-button__remove-button-show");
        uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
        uploadAuthorPhotoButton.innerHTML = 'Upload New';
        uploadAuthorPhotoButton.classList.add("author-photo-form__upload-button-view");
        document.querySelector(".upload-button__icon").classList.add("upload-button__icon-view");
        if (file) {
            reader.readAsDataURL(file);
        }
    }
)

removeAuthorPhoto.addEventListener(
    "click",
    () => {
        document.querySelector(".author-photo-form__remove-button").classList.remove("remove-button__remove-button-show");
        uploadAuthorPhotoButton = document.getElementById("uploadAuthorPhotoButton");
        uploadAuthorPhotoButton.innerHTML = 'Upload';
        uploadAuthorPhotoButton.classList.remove("author-photo-form__upload-button-view");
        document.querySelector(".upload-button__icon").classList.remove("upload-button__icon-view");
        const previewPostCardAuthorPhoto = document.querySelector(".post-card-info__photo");
        const previewInput = document.querySelector(".preview-author-photo");
        defaultAuthorPhoto = '../static/img/page/author_background.png';
        previewPostCardAuthorPhoto.src = defaultAuthorPhoto;
        previewInput.src = defaultAuthorPhoto;
    }
)

uploadTinyPostIMG.addEventListener(
    "input",
    () => {
        const previewPostCardPhoto = document.querySelector(".post-card__photo");
        const previewInput = document.querySelector(".upload-place-tiny__img");
        const file = document.querySelector(".input-hero-image-tiny__input").files[0];
        const reader = new FileReader();
        reader.addEventListener(
            "load",
            () => {
                previewPostCardPhoto.src = reader.result;
                previewInput.src = reader.result;
                previewIMGInput = reader.result;
                previewIMGInput = file.name;
                document.querySelector(".tiny-img-buttons").classList.add("tiny-img-buttons-show");
                document.querySelector(".input-hero-image-tiny__sign").classList.add("input-hero-image-tiny__sign-remove");
            },
            false
        );

        if (file) {
            reader.readAsDataURL(file);
        }
    }
)

removeTinyPostIMG.addEventListener(
    "click",
    () => {
        const defaultPostTinyIMG = "../static/img/page/card_background.png";
        const previewPostCardPhoto = document.querySelector(".post-card__photo");
        const previewInput = document.querySelector(".upload-place-tiny__img");
        previewPostCardPhoto.src = defaultPostTinyIMG;
        previewInput.src = defaultPostTinyIMG;
        document.querySelector(".tiny-img-buttons").classList.remove("tiny-img-buttons-show");
        document.querySelector(".input-hero-image-tiny__sign").classList.remove("input-hero-image-tiny__sign-remove");
    }
)

uploadPostIMG.addEventListener(
    "input",
    () => {
        const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
        const previewInput = document.querySelector(".upload-place__img");
        const file = document.querySelector(".input-hero-image__input").files[0];
        const reader = new FileReader();
        reader.addEventListener(
            "load",
            () => {
                previewPostCardPhoto.src = reader.result;
                previewInput.src = reader.result;
                postImgInput = reader.result;
                postImgInputName = file.name;
                document.querySelector(".img-buttons").classList.add("img-buttons-show");
                document.querySelector(".input-hero-image__sign").classList.add("input-hero-image__sign-remove");
            },
            false
        );

        if (file) {
            reader.readAsDataURL(file);
        }
    }
)

removePostIMG.addEventListener(
    "click",
    () => {
        const defaultPostIMG = "../static/img/page/cover_background.png";
        const previewPostCardPhoto = document.querySelector(".article-preview-post-visual__photo");
        const previewInput = document.querySelector(".upload-place__img");
        document.querySelector(".img-buttons").classList.remove("img-buttons-show");
        document.querySelector(".input-hero-image__sign").classList.remove("input-hero-image__sign-remove");
        previewPostCardPhoto.src = defaultPostIMG;
        previewInput.src = defaultPostIMG;
    }
)

// function localPrint()
// {
//     console.log("Title:", title.value);
//     console.log("Short Description:", subTitle.value);
//     console.log("Author Name:", author.value);
//     console.log("Publish Date:", publishDate.value);
//     console.log("Author Photo:", authorIMGInput);
//     console.log("Cover Photo:", previewIMGInput);
//     console.log("Card Photo:", postImgInput);
// }

// function Publish()
// {
//     localPrint();
//     const data = {
//         title: title.value,
//         subtitle: subTitle.value,
//         author: author.value, 
//         authorIMG: authorIMGInput,
//         previewIMG: previewIMGInput,
//         postIMG: postImgInput, 
//         publishDate: publishDate.value,
//     }

//     const json = JSON.stringify(data);
//     console.log("json:", json);
// }


publishButton.addEventListener(
    "click",
    () => {
        removeWarning();
        if (validate())
        {
            const data = {
                title: title.value,
                subtitle: subTitle.value,
                authorName: author.value, 
                authorIMG: authorIMGInput,
                authorIMGName: authorIMGInputName,
                postIMG: postImgInput, 
                postIMGName: postImgInputName, 
                previewIMG: previewIMGInput,
                previewIMGName: previewIMGInputName,
                publishDate: publishDate.value,
                content: content.value,
            }
            createWarning("success", "Publish Complete!");
            doPost(data);
        }
        else
        {
            createWarning("critical", "Whoops! Some fields need your attention :o");
        }
    }    
)

function doPost(data)
{
    let newpost = JSON.stringify(data);

    let XHR = new XMLHttpRequest();

    XHR.open("POST", "/api/post");
    XHR.send(newpost);
}

// async function doPost(data) {
//     const response = await fetch("/api/post", {
//         method: "POST",
//         headers: {
//             "Content-Type": "application/json;charset=utf-8"
//         },
//         body: JSON.stringify(data, null, "\t")
//     });
//     if (!response.ok) {
//         alert("Ошибка HTTP: " + response.status);
//     }
// }

function validate()
{
    let result = true;

    if (title.value == "")
    {
        removeError(title);
        createError(title, "Title is required.");
        if (result)
        {
            result = false;
        }
    }

    if (subTitle.value == "")
    {
        removeError(subTitle);
        createError(subTitle, "Subtitle is required.");
        if (result)
        {
            result = false;
        }
    }

    if (author.value == "")
    {
        removeError(author);
        createError(author, "Author Name is required.");
        if (result)
        {
            result = false;
        }
    }

    if (publishDate.value == "")
    {
        removeError(publishDate);
        createError(publishDate, "Publish Date is required.");
        if (result)
        {
            result = false;
        }
    }

    return result;
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