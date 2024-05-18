import 'htmx.org'
import videojs from "video.js";
import "videojs-hls-quality-selector"
import "./post.config"
import "video.js/dist/video-js.css"
import "@videojs/http-streaming"
import {v4 as uuidv4} from "uuid"

import "@uppy/core"
import Uppy from "@uppy/core";
import AwsS3 from "@uppy/aws-s3";
import FileInput from '@uppy/file-input';
import ProgressBar from "@uppy/progress-bar";
import StatusBar from "@uppy/status-bar";

import '@uppy/core/dist/style.min.css';
import '@uppy/file-input/dist/style.css';
import '@uppy/progress-bar/dist/style.min.css';
import '@uppy/status-bar/dist/style.min.css';



var uppy = new Uppy()
    .use(AwsS3, {
        async getUploadParameters(file) {
            const response = await fetch("/api/upload", {
                method: "POST",
                body: formData
            });
            return {
                method: "PUT",
                url: response.data.url
            }
        }
    })
    .use(FileInput, { target: '#uppy-file-input' })
    .use(ProgressBar, { target: '#progress-bar' })
    .use(StatusBar, { target: '#status-bar' });






(function(window, document, undefined) {

    // code that should be taken care of right away

    window.onload = init;

    function init(){
        // the code to be called when the dom has loaded
        // #document has its nodes
        const uploadForm = document.getElementById("#uploadForm")
        console.log(uuidv4())

        uploadForm.

    }

})(window, document, undefined);

async function submitHandler(event) {
    event.preventDefault()

    const formData = new FormData(event.target)
    const clip = e.target.elements.file.files[0]
    await fileHandle(clip, formData)
    return false
}

async function fileHandle(clip, formData) {
    const objectName = uuidv4();

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
/*
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
 */