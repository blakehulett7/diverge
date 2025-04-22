console.log("Jesus is Lord!")

const uploadForm = document.getElementById("resume-input")
console.log(uploadForm.files)

uploadForm.addEventListener("change", function() {
    console.log(uploadForm.submit())
})
