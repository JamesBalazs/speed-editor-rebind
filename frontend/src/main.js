import { Events } from "@wailsio/runtime";
import { SpeedEditorService } from "../bindings/github.com/JamesBalazs/speed-editor-rebind";

const resultElement = document.getElementById("result");
const connectionElement = document.getElementById("connection");

let connected = false;

Events.On("heartbeat", (event) => {
    if (event.data.Error != "") {
        connectionElement.innerText = event.data.Error;
    } else if (connected === false && event.data.Connected == true) {
        connected = true;

        connectionElement.innerText = `Speed Editor connected. Serial number: ${event.data.Serial}`;
    } else if (connected === true && event.data.Connected === false) {
        connected = false;

        connectionElement.innerText = "Waiting for connection...";
    }
});

// document.addEventListener("DOMContentLoaded", () => {
//     SpeedEditorService.Connect().catch((err) => console.error(err));
// });
