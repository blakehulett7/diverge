console.log("Jesus is Lord!")

const resumeInput = document.getElementById("resume-input")
console.log(uploadForm.files)

uploadForm.addEventListener("change", function() {
    console.log("files changed")

    document.forms["upload-form"].submit()
})
