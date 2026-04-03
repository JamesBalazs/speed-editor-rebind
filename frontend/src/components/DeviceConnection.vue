<script setup>
import { ref, onMounted } from "vue";
import { Events } from "@wailsio/runtime";
// import { GreetService } from "../../bindings/changeme";

const statusWaiting = "Waiting for connection...";

const connectionString = ref(statusWaiting);
const connected = ref(false);

onMounted(() => {
    Events.On("heartbeat", (event) => {
        console.log("got heartbeat: ", event.data);

        if (event.data.Error != "") {
            connectionString.value = event.data.Error;
        } else if (connected.value === false && event.data.Connected == true) {
            connected.value = true;
            connectionString.value = `Speed Editor connected. Serial number: ${event.data.Serial}`;

            // Initialize key elements map from static HTML
            // initKeyElements();
        } else if (connected.value === true && event.data.Connected === false) {
            connected.value = false;
            connectionString.value = statusWaiting;

            // Clear highlights
            // Object.values(keyElements).forEach((el) => {
            //     el.classList.remove("active");
            // });
        }
    });
});
</script>

<template>
    <div>
        <p>{{ connectionString }}</p>
    </div>
</template>
