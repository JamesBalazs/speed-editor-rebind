<script setup>
import { ref, onMounted, computed } from "vue";
import { Events } from "@wailsio/runtime";
import { store } from "../store.js";
// import { GreetService } from "../../bindings/changeme";

const props = defineProps({
    led: Number,
    jogLed: Number,
});

const isActive = ref(false);

onMounted(() => {
    if (props.led > 0) {
        Events.On(`consolidateLeds`, (event) => {
            // Highlight the pressed keys
            consolidateLed(event.data, props.led);
        });
    }

    if (props.jogLed > 0) {
        Events.On(`consolidateJogLeds`, (event) => {
            // Highlight the pressed keys
            consolidateLed(event.data, props.jogLed);
        });
    }
});

function consolidateLed(data, id) {
    if (data.includes(id)) {
        isActive.value = true;
    } else {
        isActive.value = false;
    }
}
</script>

<template>
    <div class="led-container">
        <span class="led" :class="{ active: isActive }"></span>
    </div>
</template>

<style scoped>
.led {
    display: inline-block;
    border: 1px solid #333;
    padding: 2px 7px;
    background-color: #001;
}

.led-container {
    margin-top: -8px;
}

.active {
    background-color: rgba(255, 38, 54, 1);
}
</style>
