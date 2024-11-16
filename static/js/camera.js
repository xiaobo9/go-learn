// const logger = el("#log");
// console.log = function (message) {
//     if (typeof message == 'object' && JSON && JSON.stringify) {
//         message = JSON.stringify(message)
//     }
//     logger.innerHtml += message + '<br />';
// }
let stream;
function video_scan() {
    navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia
    let constraints = {
        video: {
            width: { min: 1280, ideal: 1920, max: 2560, },
            height: { min: 720, ideal: 1080, max: 1440, },
            // facingMode: "environment"
        }
    }

    try {
        let video = el("#video");
        // let videoStream = await navigator.mediaDevices.getUserMedia(constraints);
        navigator.mediaDevices.getUserMedia({
            video: true
        }).then(function (videoStream) {
            video.srcObject = videoStream;
            stream = videoStream;

            video.onloadedmetadata = function (e) {
                // scan()
                console.log("onloadedmetadata")
            }
        });
    } catch (error) {
        mylog(error)
    }
}
function video_close() {
    console.log("close")
    stream.getTracks().forEach((track) => {
        console.log("close track")
        track.stop();
    })
    let video = el("#video");
    video.srcObject = null;
}
function scan() {
    const video = el("#video");
    const canvas = el("#canvas");
    canvas.width = video.videoWidth;
    canvas.height = video.videoHeight;
    canvas.getContext("2d").drawImage(video, 0, 0);
    const image = document.createElement("img");
    // base64 图片，可以上传到后端
    const imgData = canvas.toDataURL("image/png")
    image.src = imgData;
    el("#screenshotsContainer").prepend(image);
}


(
    function () {



    })();
