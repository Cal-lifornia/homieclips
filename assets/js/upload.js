import {v4 as uuidv4} from "/assets/index";


(function(window, document, undefined) {

    // code that should be taken care of right away

    window.onload = init;

    function init(){
        // the code to be called when the dom has loaded
        // #document has its nodes
        const uploadForm = document.getElementById("#uploadForm")
        console.log(uuidv4())

        uploadForm.onsubmit =  async (e) => {
            await e.preventDefault()
            console.log(formData.values())
            const formData = new FormData(uploadForm);
            const clip = e.target.elements.file.files[0]
            await fileHandle(clip, formData)
            return false
        };

    }

})(window, document, undefined);



async function fileHandle(clip, formData) {
    //const objectName = uuidv4();

    const objectName = "018ea927-be9b-7a28-b6dd-b1a188b28c07"

    formData.append("object_name", objectName)

    const response = await fetch("/api/upload", {
        method: "POST",
        body: formData
    });

    const presignUrl = await response.json()

    await uploadFile(clip,presignUrl["presign_put_url"])

    /*
    const postClip = async () => {
        try {


            console.log("retrieved upload URL")

            let uploadData = {
                file: clip,
                contentType: clip.type,
            }
            const res = await fetch(presignUrl["presign_put_url"],{
                method: "POST",
                body: uploadData,
            })

            console.log("started upload")
        } catch (error) {
            console.error("Error: ", error)
        }
    }

    await postClip()
     */
}

function uploadFile(
    file: File,
    presignedUploadUrl: string,
    onProgress: (pct: number) => void,
): Promise<void> {
    return new Promise<void>((resolve, reject) => {
        const xhr = new XMLHttpRequest();
        xhr.upload.addEventListener("progress", (e) => {
            if (e.lengthComputable) {
                const pct = e.loaded / e.total;
                onProgress(pct * 100);
            }
        });
        xhr.upload.addEventListener("error", (e) => {
            reject(new Error("Upload failed: " + e.toString()));
        });
        xhr.upload.addEventListener("abort", (e) => {
            reject(new Error("Upload aborted: " + e.toString()));
        });
        xhr.addEventListener("load", (e) => {
            if (xhr.status === 200) {
                resolve();
            } else {
                reject(new Error("Upload failed " + xhr.status));
            }
        });
        xhr.open("PUT", presignedUploadUrl, true);
        try {
            xhr.send(file);
        } catch (e) {
            reject(new Error("Upload failed: " + e.toString()));
        }
    });
}