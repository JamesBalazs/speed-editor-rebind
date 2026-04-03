import { Events } from "@wailsio/runtime";

const connectionElement = document.getElementById("connection");
const keyboardElement = document.getElementById("keyboard");

let connected = false;
let keyElements = {}; // Map key Id to DOM element

// Build keyElements map from static HTML
function initKeyElements() {
    if (!keyboardElement) return;

    keyElements = {};
    const keyDivs = keyboardElement.querySelectorAll(".key");
    keyDivs.forEach((keyDiv) => {
        const id = parseInt(keyDiv.dataset.id);
        if (!isNaN(id)) {
            keyElements[id] = keyDiv;
        }
    });
}

// Highlight pressed keys
function highlightKeys(pressedKeyIds) {
    // First, remove all active states
    Object.values(keyElements).forEach((el) => {
        el.classList.remove("active");
    });

    // Then, add active state to pressed keys
    pressedKeyIds.forEach((id) => {
        const keyEl = keyElements[id];
        if (keyEl) {
            keyEl.classList.add("active");
        }
    });
}

Events.On("heartbeat", (event) => {
    if (event.data.Error != "") {
        connectionElement.innerText = event.data.Error;
    } else if (connected === false && event.data.Connected == true) {
        connected = true;
        connectionElement.innerText = `Speed Editor connected. Serial number: ${event.data.Serial}`;

        // Initialize key elements map from static HTML
        initKeyElements();
    } else if (connected === true && event.data.Connected === false) {
        connected = false;
        connectionElement.innerText = "Waiting for connection...";

        // Clear highlights
        Object.values(keyElements).forEach((el) => {
            el.classList.remove("active");
        });
    }
});

Events.On("keyPress", (event) => {
    // Highlight the pressed keys
    highlightKeys(event.data);

    // Log for debugging
    console.log("Key pressed:", event.data);
});
